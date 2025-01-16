package xalan_test

import (
	stats "benchmarks/utils"
	"benchmarks/xalan"
	"encoding/csv"
	"os"
	"runtime"
	"testing"
	"time"
)

func BenchmarkParMnemonics(b *testing.B) {
	file, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var memStatsAfterGCExecution runtime.MemStats
	executions_data := [][]string{}

	for i := 0; i < b.N; i++ {
		runtime.GC()
		start := time.Now()
		xalan.Run()
		runtime.ReadMemStats(&memStatsAfterGCExecution)
		executions_data = append(executions_data, stats.GenerateExecutionData(&memStatsAfterGCExecution, time.Since(start)))
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, data := range executions_data {
		err := writer.Write(data)
		if err != nil {
			panic(err)
		}
	}
}
