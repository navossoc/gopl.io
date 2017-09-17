// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Code based on Echo1, Echo2, Echo3
package main

import (
	"fmt"
	"io"
	"strings"
)

func echo1(w io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo2(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo3(w io.Writer, args []string) {
	fmt.Fprintln(w, strings.Join(args, " "))
}
