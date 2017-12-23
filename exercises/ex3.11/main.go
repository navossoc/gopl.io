// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {

	var sign string
	if strings.ContainsAny(s, "+-") {
		sign = s[:1]
	}

	var integer, fractional string
	if dot := strings.IndexByte(s, '.'); dot == -1 {
		integer = s[len(sign):]
	} else {
		integer = s[len(sign):dot]
		fractional = s[dot:]
	}

	n := len(integer)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	buf.WriteString(sign)

	i := n % 3
	if i > 0 {
		buf.WriteString(integer[:i])
		buf.WriteByte(',')
	}

	for ; i < n; i += 3 {
		buf.WriteString(integer[i : i+3])
		buf.WriteByte(',')
	}

	buf.Truncate(buf.Len() - 1)

	buf.WriteString(fractional)

	return buf.String()
}

//!-
