// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

// palette from https://androidarts.com/palette/16pal.htm
var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x9D, 0x9D, 0x9D, 0xFF},
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	color.RGBA{0xBE, 0x26, 0x33, 0xFF},
	color.RGBA{0xE0, 0x6F, 0x8B, 0xFF},
	color.RGBA{0x49, 0x3C, 0x2B, 0xFF},
	color.RGBA{0xA4, 0x64, 0x22, 0xFF},
	color.RGBA{0xEB, 0x89, 0x31, 0xFF},
	color.RGBA{0xF7, 0xE2, 0x6B, 0xFF},
	color.RGBA{0x2F, 0x48, 0x4E, 0xFF},
	color.RGBA{0x44, 0x89, 0x1A, 0xFF},
	color.RGBA{0xA3, 0xCE, 0x27, 0xFF},
	color.RGBA{0x1B, 0x26, 0x32, 0xFF},
	color.RGBA{0x00, 0x57, 0x84, 0xFF},
	color.RGBA{0x31, 0xA2, 0xF2, 0xFF},
	color.RGBA{0xB2, 0xDC, 0xEF, 0xFF},
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// cool color effect
	index := nextIndex(len(palette))

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// one color per frame
		colorIndex := uint8(index())

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func nextIndex(max int) func() int {
	i := 0
	dir := 1

	return func() int {
		i += dir

		n := i % (max + 1)
		if n <= 0 {
			n = 1
			dir = -dir
		} else if n >= max {
			n = max - 1
			dir = -dir
		}

		return n
	}
}

//!-main
