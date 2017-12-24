// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	if n == 0 {
		return
	}

	total := len(s)

	n = n % total
	if n < 0 {
		n += total
	}

	for j := 0; j < n; j++ {
		for i := 0; i < total-1; i++ {
			s[i], s[i+1] = s[i+1], s[i]
		}
	}
}
