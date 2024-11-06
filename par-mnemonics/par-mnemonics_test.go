package parmnemonics_test

import (
	parmnemonics "benchmarks/par-mnemonics"
	stats "benchmarks/utils"
	"runtime"
	"testing"
)

func BenchmarkParMnemonics(b *testing.B) {
	var memStatsBeforeGCExecution runtime.MemStats
	var memStatsAfterGCExecution runtime.MemStats
	for i := 0; i < b.N; i++ {
		runtime.ReadMemStats(&memStatsBeforeGCExecution)
		parmnemonics.Run("5225")
		runtime.ReadMemStats(&memStatsAfterGCExecution)
		stats.PrintGCStats(memStatsBeforeGCExecution)
		stats.PrintGCStats(memStatsAfterGCExecution)
	}
}
