package main

import (
	"io/ioutil"
	"testing"
)

var (
	args0 = []string{}
	args1 = []string{"apples"}
	args2 = []string{"apples", "bananas"}
	args3 = []string{"apples", "bananas", "coconuts"}
	args9 = []string{"apples", "bananas", "coconuts", "apples", "bananas", "coconuts", "apples", "bananas", "coconuts"}
)

func BenchmarkEcho1Args0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ioutil.Discard, args0)
	}
}

func BenchmarkEcho1Args1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ioutil.Discard, args1)
	}
}

func BenchmarkEcho1Args2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ioutil.Discard, args2)
	}
}

func BenchmarkEcho1Args3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ioutil.Discard, args3)
	}
}

func BenchmarkEcho1Args9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ioutil.Discard, args9)
	}
}

func BenchmarkEcho2Args0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(ioutil.Discard, args0)
	}
}

func BenchmarkEcho2Args1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(ioutil.Discard, args1)
	}
}

func BenchmarkEcho2Args2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(ioutil.Discard, args2)
	}
}

func BenchmarkEcho2Args3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(ioutil.Discard, args3)
	}
}

func BenchmarkEcho2Args9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(ioutil.Discard, args9)
	}
}

func BenchmarkEcho3Args0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ioutil.Discard, args0)
	}
}

func BenchmarkEcho3Args1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ioutil.Discard, args1)
	}
}

func BenchmarkEcho3Args2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ioutil.Discard, args2)
	}
}

func BenchmarkEcho3Args3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ioutil.Discard, args3)
	}
}

func BenchmarkEcho3Args9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ioutil.Discard, args9)
	}
}

/*
$ go test -bench=. github.com/navossoc/gopl.io/exercises/ex1.3
goos: windows
goarch: amd64
pkg: github.com/navossoc/gopl.io/exercises/ex1.3
BenchmarkEcho1Args0-8           20000000                80.6 ns/op
BenchmarkEcho1Args1-8           20000000                79.5 ns/op
BenchmarkEcho1Args2-8           10000000               130 ns/op
BenchmarkEcho1Args3-8           10000000               209 ns/op
BenchmarkEcho1Args9-8            2000000               640 ns/op
BenchmarkEcho2Args0-8           20000000                80.2 ns/op
BenchmarkEcho2Args1-8           10000000               131 ns/op
BenchmarkEcho2Args2-8           10000000               204 ns/op
BenchmarkEcho2Args3-8            5000000               278 ns/op
BenchmarkEcho2Args9-8            2000000               740 ns/op
BenchmarkEcho3Args0-8           20000000                84.8 ns/op
BenchmarkEcho3Args1-8           10000000               121 ns/op
BenchmarkEcho3Args2-8           10000000               191 ns/op
BenchmarkEcho3Args3-8           10000000               211 ns/op
BenchmarkEcho3Args9-8            5000000               307 ns/op
PASS
ok      github.com/navossoc/gopl.io/exercises/ex1.3     27.873s
*/
