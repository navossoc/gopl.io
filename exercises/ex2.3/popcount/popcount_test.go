package popcount_test

import (
	"testing"

	"github.com/navossoc/gopl.io/exercises/ex2.3/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

/*
$ go test -bench=. github.com/navossoc/gopl.io/exercises/ex2.3/popcount
goos: windows
goarch: amd64
pkg: github.com/navossoc/gopl.io/exercises/ex2.3/popcount
BenchmarkPopCount-8             2000000000               0.33 ns/op
BenchmarkPopCountLoop-8         100000000               22.7 ns/op
PASS
ok      github.com/navossoc/gopl.io/exercises/ex2.3/popcount    3.013s
*/
