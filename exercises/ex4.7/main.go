package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var bs = []byte("Hello, 世界")
	reverseUTF8(bs)
	fmt.Printf("%q\n", bs)
}

func reverseUTF8(p []byte) {
	if utf8.RuneCount(p) <= 1 {
		return
	}

	for i := 0; i < len(p)-1; {
		r, s := utf8.DecodeLastRune(p)

		// copy to right
		copy(p[i+s:], p[i:])

		// restore selected rune
		copy(p[i:], []byte(string(r)))
		i += s
	}
}
