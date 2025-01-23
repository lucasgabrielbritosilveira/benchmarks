package stats

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func PrintGCStats(memStats *runtime.MemStats) {
	fmt.Println("")
	fmt.Print(memStats.Mallocs, ",", memStats.Frees, ",",
		memStats.HeapAlloc, ",", memStats.HeapSys, ",", memStats.HeapReleased, ",", memStats.NextGC, ",",
		memStats.LastGC, ",", memStats.PauseTotalNs, ",", memStats.NumGC, ",", memStats.NumForcedGC, memStats.GCCPUFraction, '\n')

}

func GenerateExecutionData(memStats *runtime.MemStats, totalTime time.Duration) []string {
	return []string{
		os.Getenv("GOGC"),
		strconv.FormatUint(memStats.Mallocs, 10), strconv.FormatUint(memStats.Frees, 10),
		strconv.FormatUint(memStats.HeapAlloc, 10), strconv.FormatUint(memStats.HeapSys, 10), strconv.FormatUint(memStats.HeapReleased, 10), strconv.FormatUint(memStats.NextGC, 10),
		strconv.FormatUint(memStats.LastGC, 10), strconv.FormatUint(memStats.PauseTotalNs, 10), strconv.FormatUint(uint64(memStats.NumGC), 10), strconv.FormatUint(uint64(memStats.NumForcedGC), 10),
		strconv.FormatFloat(memStats.GCCPUFraction, 'f', 2, 64), strconv.FormatInt(totalTime.Nanoseconds(), 10),
	}
}
