package popcount_test

import (
	"testing"

	"github.com/navossoc/gopl.io/exercises/ex2.4/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(0x1234567890ABCDEF)
	}
}

/*
$ go test -bench=. github.com/navossoc/gopl.io/exercises/ex2.4/popcount
goos: windows
goarch: amd64
pkg: github.com/navossoc/gopl.io/exercises/ex2.4/popcount
BenchmarkPopCount-8             2000000000               0.33 ns/op
BenchmarkPopCountShift-8        20000000                62.1 ns/op
PASS
ok      github.com/navossoc/gopl.io/exercises/ex2.4/popcount    2.027s
*/
