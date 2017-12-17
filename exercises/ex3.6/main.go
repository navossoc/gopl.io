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

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		sampling               = 2
		swidth                 = width * sampling
		sheight                = height * sampling
	)

	// scale image up
	simg := image.NewRGBA(image.Rect(0, 0, swidth, sheight))
	for py := 0; py < sheight; py++ {
		y := float64(py)/sheight*(ymax-ymin) + ymin
		for px := 0; px < swidth; px++ {
			x := float64(px)/swidth*(xmax-xmin) + xmin
			z := complex(x, y)

			simg.Set(px, py, mandelbrot(z))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			spx, spy := px*sampling, py*sampling
			r1, g1, b1, a1 := simg.At(spx+0, spy+0).RGBA()
			r2, g2, b2, a2 := simg.At(spx+0, spy+1).RGBA()
			r3, g3, b3, a3 := simg.At(spx+1, spy+0).RGBA()
			r4, g4, b4, a4 := simg.At(spx+1, spy+1).RGBA()

			avg := color.RGBA{
				uint8((r1 + r2 + r3 + r4) >> 10), // sum / 4 / 1024
				uint8((g1 + g2 + g3 + g4) >> 10), // sum / 4 / 1024
				uint8((b1 + b2 + b3 + b4) >> 10), // sum / 4 / 1024
				uint8((a1 + a2 + a3 + a4) >> 10), // sum / 4 / 1024
			}

			// Image point (px, py) represents complex value z.
			img.Set(px, py, avg)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast * n, 4 * n, 255, 255}
		}
	}
	return color.Black
}

//!-
