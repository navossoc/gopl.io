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
		bxdiff  = big.NewFloat(xmax - xmin)
		bydiff  = big.NewFloat(ymax - ymin)
		bymin   = big.NewFloat(ymin)
		bxmin   = big.NewFloat(xmin)
		bheight = big.NewFloat(height)
		bwidth  = big.NewFloat(width)
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewFloat(float64(py))
		y.Quo(y, bheight).Mul(y, bydiff).Add(y, bymin)

		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px))
			x.Quo(x, bwidth).Mul(x, bxdiff).Add(x, bxmin)

			z := NewComplexBigFloat(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z *ComplexBigFloat) color.Color {
	const iterations = 200
	const contrast = 15

	root := big.NewFloat(4)

	v := &ComplexBigFloat{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		v.Mul(v, v).Add(v, z)
		if SquareAbs(v).Cmp(root) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

///

type ComplexBigFloat struct {
	real *big.Float
	imag *big.Float
}

func NewComplexBigFloat(real, imag *big.Float) *ComplexBigFloat {
	return &ComplexBigFloat{real, imag}
}

func (z *ComplexBigFloat) Add(x, y *ComplexBigFloat) *ComplexBigFloat {
	z.real.Add(x.real, y.real)
	z.imag.Add(x.imag, y.imag)
	return z
}

func (z *ComplexBigFloat) Mul(x, y *ComplexBigFloat) *ComplexBigFloat {
	var real, imag, temp big.Float
	real.Mul(x.real, y.real).Sub(&real, temp.Mul(x.imag, y.imag))
	imag.Mul(x.real, y.imag).Add(&imag, temp.Mul(x.imag, y.real))
	z.real, z.imag = &real, &imag
	return z
}

func SquareAbs(x *ComplexBigFloat) *big.Float {
	var temp, temp2 big.Float
	return temp.Mul(x.real, x.real).Add(&temp, temp2.Mul(x.imag, x.imag))
}

func (z *ComplexBigFloat) String() string {
	return fmt.Sprintf("(%v %+vi)", z.real.String(), z.imag.String())
}

//!-
