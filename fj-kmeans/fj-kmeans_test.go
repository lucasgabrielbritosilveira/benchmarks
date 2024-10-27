package fjkmeans

import (
	"testing"
)

func BenchmarkFJKMeans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result uint8 = 1
		result++
	}
}
