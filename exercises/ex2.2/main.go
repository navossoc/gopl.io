// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument or stdin to different units.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/navossoc/gopl.io/exercises/ex2.2/unitconv"
)

func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			convert(scanner.Text())
		}
		return
	}

	for _, arg := range os.Args[1:] {
		convert(arg)
	}
}

func convert(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Temperature")
	f := unitconv.Fahrenheit(t)
	fmt.Printf("%s = %s\n", f, unitconv.FToC(f))
	c := unitconv.Celsius(t)
	fmt.Printf("%s = %s\n", c, unitconv.CToF(c))
	k := unitconv.Kelvin(t)
	fmt.Printf("%s = %s\n", k, unitconv.KToC(k))

	fmt.Println("Length")
	ft := unitconv.Feet(t)
	fmt.Printf("%s = %s\n", ft, unitconv.FToM(ft))
	m := unitconv.Meter(t)
	fmt.Printf("%s = %s\n", m, unitconv.MToF(m))

	fmt.Println("Mass")
	kg := unitconv.Kilogram(t)
	fmt.Printf("%s = %s\n", kg, unitconv.KGToP(kg))
	lb := unitconv.Pound(t)
	fmt.Printf("%s = %s\n", lb, unitconv.PToKG(lb))

	fmt.Println()
}

//!-
