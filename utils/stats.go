package stats

import (
	"fmt"
	"runtime"
)

func PrintGCStats(memStats runtime.MemStats) {
	runtime.ReadMemStats(&memStats)
	fmt.Println("")
	fmt.Print(memStats.Mallocs, ",", memStats.Frees, ",",
		memStats.HeapAlloc, ",", memStats.HeapSys, ",", memStats.HeapReleased, ",", memStats.NextGC, ",",
		memStats.LastGC, ",", memStats.PauseTotalNs, ",", memStats.NumGC, ",", memStats.NumForcedGC, '\n')

}
