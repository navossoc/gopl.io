package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type wordCount struct {
	word  string
	count int
}

var words = make(map[string]int)

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "need a file to read")
		os.Exit(1)
	}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		func() {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "open file:", err)
				return
			}
			defer f.Close()

			// count all words and their frequency
			wordfreq(f)
		}()
	}

	// copy results to a slice (ordered)
	wc := make([]wordCount, len(words))

	var i int
	for w, c := range words {
		wc[i] = wordCount{w, c}
		i++
	}

	// sort by count and word
	sort.Slice(wc, func(i int, j int) bool {
		if wc[i].count == wc[j].count {
			return wc[i].word < wc[j].word
		}

		return wc[i].count < wc[j].count
	})

	// print results
	for i, wc := range wc {
		fmt.Printf("%5d: %v %v\n", i, wc.word, wc.count)
	}
}

func wordfreq(r io.Reader) {
	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read file:", err)
	}
}
