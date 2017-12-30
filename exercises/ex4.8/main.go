// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var counts [invalid + 1]int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		switch {
		case unicode.IsControl(r):
			counts[control]++
		case unicode.IsDigit(r):
			counts[digit]++
		case unicode.IsGraphic(r):
			counts[graphic]++
		case unicode.IsLetter(r):
			counts[letter]++
		case unicode.IsLower(r):
			counts[lower]++
		case unicode.IsMark(r):
			counts[mark]++
		case unicode.IsNumber(r):
			counts[number]++
		case unicode.IsPrint(r):
			counts[print]++
		case unicode.IsPunct(r):
			counts[punct]++
		case unicode.IsSpace(r):
			counts[space]++
		case unicode.IsSymbol(r):
			counts[symbol]++
		case unicode.IsTitle(r):
			counts[title]++
		case unicode.IsUpper(r):
			counts[upper]++
		case r == unicode.ReplacementChar && n == 1:
			counts[invalid]++
		}
	}

	fmt.Printf("category\tcount\n")
	for i, n := range counts {
		fmt.Printf("%v\t\t%d\n", UnicodeCategory(i), n)
	}
}

//!-
