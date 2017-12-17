// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var (
		rxdiff = new(big.Rat).SetInt64(xmax - xmin)
		rydiff = new(big.Rat).SetInt64(ymax - ymin)
		rymin  = new(big.Rat).SetInt64(ymin)
		rxmin  = new(big.Rat).SetInt64(xmin)
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewRat(int64(py), height)
		y.Mul(y, rydiff).Add(y, rymin)

		for px := 0; px < width; px++ {
			x := big.NewRat(int64(px), width)
			x.Mul(x, rxdiff).Add(x, rxmin)

			z := NewComplexBigRat(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z *ComplexBigRat) color.Color {
	const iterations = 200
	const contrast = 15

	root := new(big.Rat).SetInt64(4)

	v := &ComplexBigRat{new(big.Rat).SetInt64(0), new(big.Rat).SetInt64(0)}
	for n := uint8(0); n < iterations; n++ {
		v.Mul(v, v).Add(v, z)
		if SquareAbs(v).Cmp(root) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

///

type ComplexBigRat struct {
	real *big.Rat
	imag *big.Rat
}

func NewComplexBigRat(real, imag *big.Rat) *ComplexBigRat {
	return &ComplexBigRat{real, imag}
}

func (z *ComplexBigRat) Add(x, y *ComplexBigRat) *ComplexBigRat {
	z.real.Add(x.real, y.real)
	z.imag.Add(x.imag, y.imag)
	return z
}

func (z *ComplexBigRat) Mul(x, y *ComplexBigRat) *ComplexBigRat {
	var real, imag, temp big.Rat
	real.Mul(x.real, y.real).Sub(&real, temp.Mul(x.imag, y.imag))
	imag.Mul(x.real, y.imag).Add(&imag, temp.Mul(x.imag, y.real))
	z.real, z.imag = &real, &imag
	return z
}

func SquareAbs(x *ComplexBigRat) *big.Rat {
	var temp, temp2 big.Rat
	return temp.Mul(x.real, x.real).Add(&temp, temp2.Mul(x.imag, x.imag))
}

func (z *ComplexBigRat) String() string {
	return fmt.Sprintf("(%v %+vi)", z.real.FloatString(10), z.imag.FloatString(10))
}

//!-
