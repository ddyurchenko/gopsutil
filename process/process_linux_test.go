package process

import (
	"testing"
)

func BenchmarkAllProcesses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllProcesses()
	}
}
