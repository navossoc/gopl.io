package popcount_test

import (
	"testing"

	"github.com/navossoc/gopl.io/exercises/ex2.5/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClear(0x1234567890ABCDEF)
	}
}

/*
$ go test -bench=. github.com/navossoc/gopl.io/exercises/ex2.5/popcount
goos: windows
goarch: amd64
pkg: github.com/navossoc/gopl.io/exercises/ex2.5/popcount
BenchmarkPopCount-8             2000000000               0.33 ns/op
BenchmarkPopCountClear-8        50000000                24.2 ns/op
PASS
ok      github.com/navossoc/gopl.io/exercises/ex2.5/popcount    1.955s
*/
