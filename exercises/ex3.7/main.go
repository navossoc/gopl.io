// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// newton roots
const (
	r1 = -1
	r2 = 1
	r3 = -1.0i
	r4 = 1.0i
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			y := 255 - (i * contrast)
			switch {
			case cmplx.Abs(z-r1) < 1e-6:
				return color.YCbCr{y, 64, 128}
			case cmplx.Abs(z-r2) < 1e-6:
				return color.YCbCr{y, 255, 80}
			case cmplx.Abs(z-r3) < 1e-6:
				return color.YCbCr{y, 72, 255}
			case cmplx.Abs(z-r4) < 1e-6:
				return color.YCbCr{y, 64, 160}
			}
		}
	}
	return color.Black
}

//!-
