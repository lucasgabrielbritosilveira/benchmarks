package stats

import (
	"fmt"
	"runtime"
)

func PrintGCStats(memStats runtime.MemStats) {
	runtime.ReadMemStats(&memStats)
	fmt.Println(" ")
	fmt.Println("MALLOC:", memStats.Mallocs)
	fmt.Println("FREES:", memStats.Frees)
	fmt.Println("HEAP ALLOC:", memStats.HeapAlloc)
	fmt.Println("HEAP SYS:", memStats.HeapSys)
	fmt.Println("HEAP RELEASED:", memStats.HeapReleased)
	fmt.Println("NEXT GC:", memStats.NextGC)
	fmt.Println("LAST GC:", memStats.LastGC)
	fmt.Println("PAUSE TOTAL NS:", memStats.PauseTotalNs)
	fmt.Println("NUM GC:", memStats.NumGC)
	fmt.Println("NUM FORCED GC:", memStats.NumForcedGC)
	fmt.Println("-------------------------------------")
}
