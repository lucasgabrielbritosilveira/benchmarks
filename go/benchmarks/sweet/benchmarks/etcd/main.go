// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm && !plan9

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/benchmarks/sweet/benchmarks/internal/driver"
	"golang.org/x/benchmarks/sweet/benchmarks/internal/par"
	"golang.org/x/benchmarks/sweet/benchmarks/internal/server"
	"golang.org/x/benchmarks/sweet/common/diagnostics"
)

const (
	// Benchmark against multiple instances that form a cluster.
	// This helps us get more interesting behavior out of etcd vs.
	// just running a single instance because of intracluster
	// traffic.
	etcdInstances = 3
	basePort      = 2379
)

type config struct {
	host         string
	etcdBin      string
	benchmarkBin string
	tmpDir       string
	benchName    string
	short        bool
	procsPerInst int
	gomaxprocs   int
	bench        *benchmark
}

var cliCfg config

func init() {
	driver.SetFlags(flag.CommandLine)
	flag.StringVar(&cliCfg.host, "host", "127.0.0.1", "hostname of tile38 server")
	flag.StringVar(&cliCfg.etcdBin, "etcd-bin", "", "path to etcd binary")
	flag.StringVar(&cliCfg.benchmarkBin, "benchmark-bin", "", "path to benchmark binary")
	flag.StringVar(&cliCfg.tmpDir, "tmp", "", "path to temporary directory")
	flag.StringVar(&cliCfg.benchName, "bench", "", "name of the benchmark to run")
	flag.BoolVar(&cliCfg.short, "short", false, "whether to run a short version of this benchmark")

	// We're going to launch a bunch of etcd instances. Distribute
	// GOMAXPROCS between those and ourselves equally.
	procs := runtime.GOMAXPROCS(-1)
	procsPerInst := procs / (etcdInstances + 1)
	if procsPerInst == 0 {
		procsPerInst = 1
	}
	runtime.GOMAXPROCS(procsPerInst)
	cliCfg.procsPerInst = procsPerInst
	cliCfg.gomaxprocs = procs
}

type etcdInstance struct {
	name       string
	clientPort int
	peerPort   int
	cmd        *exec.Cmd
	output     bytes.Buffer
}

type portType int

const (
	badPort portType = iota
	clientPort
	peerPort
)

func clusterString(instances []*etcdInstance, typ portType) string {
	var s []string
	for _, inst := range instances {
		s = append(s, fmt.Sprintf("%s=http://%s", inst.name, inst.host(typ)))
	}
	return strings.Join(s, ",")
}

func launchEtcdCluster(cfg *config) ([]*etcdInstance, error) {
	var instances []*etcdInstance
	for i := 0; i < etcdInstances; i++ {
		instances = append(instances, &etcdInstance{
			name:       fmt.Sprintf("infra%d", i+1),
			clientPort: basePort + 2*i,
			peerPort:   basePort + 2*i + 1,
		})
	}
	initCluster := clusterString(instances, peerPort)
	for _, inst := range instances {
		inst.cmd = exec.Command(cfg.etcdBin,
			"--name", inst.name,
			"--listen-client-urls", "http://"+inst.host(clientPort),
			"--advertise-client-urls", "http://"+inst.host(clientPort),
			"--listen-peer-urls", "http://"+inst.host(peerPort),
			"--initial-advertise-peer-urls", "http://"+inst.host(peerPort),
			"--initial-cluster-token", "etcd-cluster-1",
			"--initial-cluster", initCluster,
			"--initial-cluster-state", "new",
			"--data-dir", filepath.Join(cfg.tmpDir, inst.name+".data"),
			"--enable-pprof",
			"--logger=zap",
			"--log-outputs=stderr",
		)
		inst.cmd.Env = append(os.Environ(),
			fmt.Sprintf("GOMAXPROCS=%d", cfg.procsPerInst),
		)
		inst.cmd.Stdout = &inst.output
		inst.cmd.Stderr = &inst.output
		if err := inst.cmd.Start(); err != nil {
			return nil, fmt.Errorf("failed to start instance %q: %v", inst.name, err)
		}
	}
	// Ping all the instances to make sure they're up and ready before continuing.
	//
	// If we don't do this, then benchmarks might have the first few request have
	// a really high latency as they block until the instances are set up.
	for _, inst := range instances {
		if err := inst.ping(); err != nil {
			return nil, err
		}
	}
	return instances, nil
}

func (i *etcdInstance) ping() error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{i.host(clientPort)},
	})
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Put(context.Background(), "sample_key", "sample_value")
	return err
}

func (i *etcdInstance) host(typ portType) string {
	var port int
	switch typ {
	case clientPort:
		port = i.clientPort
	case peerPort:
		port = i.peerPort
	}
	return fmt.Sprintf("127.0.0.1:%d", port)
}

func (i *etcdInstance) shutdown() error {
	if err := i.cmd.Process.Signal(os.Interrupt); err != nil {
		return err
	}
	if _, err := i.cmd.Process.Wait(); err != nil {
		return err
	}
	return nil
}

type benchmark struct {
	name       string
	reportName string
	args       []string
	longArgs   []string // if !config.short
	shortArgs  []string // if config.short
}

var benchmarks = []benchmark{
	// "put" comes from https://etcd.io/docs/v3.5/op-guide/performance/.
	{
		name:       "put",
		reportName: "EtcdPut",
		args: []string{
			"--precise",
			"--conns=100",
			"--clients=1000",
			"put",
			"--key-size=8",
			"--sequential-keys",
			"--val-size=256",
		},
		longArgs: []string{
			"--total=100000",
		},
		shortArgs: []string{
			"--total=1000",
		},
	},
	{
		name:       "stm",
		reportName: "EtcdSTM",
		args: []string{
			"--precise",
			"--conns=100",
			"--clients=1000",
			"stm",
			"--keys=1000",
			"--keys-per-txn=2",
			"--val-size=8",
		},
		longArgs: []string{
			"--total=100000",
		},
		shortArgs: []string{
			"--total=1000",
		},
	},
}

func runBenchmark(b *driver.B, cfg *config, instances []*etcdInstance) (err error) {
	var hosts []string
	for _, inst := range instances {
		hosts = append(hosts, inst.host(clientPort))
	}
	args := append([]string{"--endpoints", strings.Join(hosts, ",")}, cfg.bench.args...)
	if cfg.short {
		args = append(args, cfg.bench.shortArgs...)
	} else {
		args = append(args, cfg.bench.longArgs...)
	}
	cmd := exec.Command(cfg.benchmarkBin, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMAXPROCS=%d", cfg.procsPerInst))

	defer func() {
		if err != nil && stderr.Len() != 0 {
			fmt.Fprintln(os.Stderr, "=== Benchmarking tool stderr ===")
			fmt.Fprintln(os.Stderr, stderr.String())
		}
	}()

	b.ResetTimer()
	if err := cmd.Run(); err != nil {
		return err
	}
	b.StopTimer()

	return reportFromBenchmarkOutput(b, stdout.String())
}

func reportFromBenchmarkOutput(b *driver.B, output string) (err error) {
	defer func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, "=== Benchmarking tool output ===")
			fmt.Fprintln(os.Stderr, output)
		}
	}()

	p50, err := getQuantileLatency("50", output)
	if err != nil {
		return err
	}
	p90, err := getQuantileLatency("90", output)
	if err != nil {
		return err
	}
	p99, err := getQuantileLatency("99", output)
	if err != nil {
		return err
	}
	b.Report("p50-latency-ns", uint64(p50*1e9))
	b.Report("p90-latency-ns", uint64(p90*1e9))
	b.Report("p99-latency-ns", uint64(p99*1e9))

	tput, err := getSummaryField("Requests/sec", output)
	if err != nil {
		return err
	}
	avg, err := getSummaryField("Average", output)
	if err != nil {
		return err
	}
	total, err := getSummaryField("Total", output)
	if err != nil {
		return err
	}

	// Report throughput.
	b.Report("ops/s", uint64(tput))

	// Report the average request latency.
	b.Ops(int(tput * total))
	b.Report(driver.StatTime, uint64(avg*1e9))
	return nil
}

func getQuantileLatency(quantile, output string) (float64, error) {
	re := regexp.MustCompile(fmt.Sprintf(`%s%%\s*in\s*(?P<value>\d+\.\d+(e(\+|-)\d+)?)`, regexp.QuoteMeta(quantile)))
	vi := re.SubexpIndex("value")
	matches := re.FindStringSubmatch(output)
	if len(matches) <= vi {
		return 0, fmt.Errorf("failed to find quantile latency pattern in output")
	}
	return strconv.ParseFloat(matches[vi], 64)
}

func getSummaryField(field, output string) (float64, error) {
	re := regexp.MustCompile(fmt.Sprintf(`%s:\s*(?P<value>\d+\.\d+(e(\+|-)\d+)?)`, regexp.QuoteMeta(field)))
	vi := re.SubexpIndex("value")
	matches := re.FindStringSubmatch(output)
	if len(matches) <= vi {
		return 0, fmt.Errorf("failed to find summary field pattern in output")
	}
	return strconv.ParseFloat(matches[vi], 64)
}

func run(cfg *config) (err error) {
	// Launch the server.
	instances, err := launchEtcdCluster(cfg)
	if err != nil {
		return fmt.Errorf("starting cluster: %v\n", err)
	}

	// Clean up the cluster after we're done.
	defer func() {
		for _, inst := range instances {
			if r := inst.shutdown(); r != nil {
				if err == nil {
					err = r
				} else {
					fmt.Fprintf(os.Stderr, "failed to shutdown %s: %v", inst.name, r)
				}
			}
			if inst.output.Len() != 0 {
				fmt.Fprintf(os.Stderr, "=== Instance %q stdout+stderr ===\n", inst.name)
				fmt.Fprintln(os.Stderr, inst.output.String())
			}
		}
	}()

	// TODO(mknyszek): Consider collecting summed memory metrics for all instances.
	// TODO(mknyszek): Consider running all instances under perf.
	opts := []driver.RunOption{
		driver.DoPeakRSS(true),
		driver.DoPeakVM(true),
		driver.DoDefaultAvgRSS(),
		driver.DoCoreDump(true),
		driver.BenchmarkPID(instances[0].cmd.Process.Pid),
		driver.DoPerf(true),
		driver.WithGOMAXPROCS(cfg.gomaxprocs),
	}
	return driver.RunBenchmark(cfg.bench.reportName, func(d *driver.B) error {
		// Set up diagnostics.
		var stopAll par.Funcs
		diag := driver.NewDiagnostics(cfg.bench.reportName)
		for _, typ := range diagnostics.Types() {
			if typ.HTTPEndpoint() == "" {
				continue
			}
			for i, inst := range instances {
				name := ""
				if !typ.CanMerge() {
					// Create a separate file for each instance.
					name = fmt.Sprintf("inst%d", i)
				}
				stop := server.FetchDiagnostic(inst.host(clientPort), diag, typ, name)
				stopAll.Add(stop)
			}
		}
		defer diag.Commit(d)
		defer stopAll.Run()

		// Actually run the benchmark.
		return runBenchmark(d, cfg, instances)
	}, opts...)
}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "error: unexpected args\n")
		os.Exit(1)
	}
	for i := range benchmarks {
		if benchmarks[i].name == cliCfg.benchName {
			cliCfg.bench = &benchmarks[i]
			break
		}
	}
	if cliCfg.bench == nil {
		fmt.Fprintf(os.Stderr, "error: unknown benchmark %q\n", cliCfg.benchName)
		os.Exit(1)
	}
	if err := run(&cliCfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
