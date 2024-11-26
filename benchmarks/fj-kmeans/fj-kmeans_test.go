package fjkmeans_test

import (
	fjkmeans "benchmarks/fj-kmeans"
	"testing"
)

func BenchmarkFJKMeans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fjkmeans.Run()
	}
}
