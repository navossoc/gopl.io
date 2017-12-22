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
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/fractal", fractalHandler)

	http.ListenAndServe("localhost:8000", nil)
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
	var (
		xmin, ymin, xmax, ymax = -2.0, -2.0, +2.0, +2.0
		width, height          = 512, 512
	)

	var err error
	if err = r.ParseForm(); err != nil {
		fmt.Fprintf(w, "failed to parse form values. %s", err)
		return
	}

	for k, v := range r.Form {
		switch k {
		case "width":
			width, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, "width parameter should be an integer value. %s", err)
				return
			}
		case "height":
			height, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, "height parameter should be an integer value. %s", err)
				return
			}
		case "xmin":
			xmin, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Fprintf(w, "xmin parameter should be an integer value. %s", err)
				return
			}
		case "ymin":
			ymin, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Fprintf(w, "ymin parameter should be an integer value. %s", err)
				return
			}
		case "xmax":
			xmax, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Fprintf(w, "xmax parameter should be an integer value. %s", err)
				return
			}
		case "ymax":
			ymax, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Fprintf(w, "ymax parameter should be an integer value. %s", err)
				return
			}
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-
