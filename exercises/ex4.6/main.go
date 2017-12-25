package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	sentence := "  Hello, \t\t 世界  ! "
	sentence = string(removeAdjacentSpaces([]byte(sentence)))
	fmt.Printf("%q\n", sentence)
}

func removeAdjacentSpaces(s []byte) []byte {
	if utf8.RuneCount(s) <= 1 {
		return s
	}

	var p int
	for b := s; len(b) > 0; {
		r, size := utf8.DecodeRune(b)
		b = b[size:]

		if unicode.IsSpace(r) {
			rr, _ := utf8.DecodeRune(b)
			if unicode.IsSpace(rr) {
				r = ' '
				continue
			}
		}

		copy(s[p:], string(r))
		p += size
	}

	return s[:p]
}
