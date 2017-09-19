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
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type params struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

var defaultValue = params{
	cycles:  5,
	res:     0.001,
	size:    100,
	nframes: 64,
	delay:   8,
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
			params, err := parseOptions(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			lissajous(w, params)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, defaultValue)
}

func lissajous(out io.Writer, p params) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5),
				p.size+int(y*float64(p.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main

func parseOptions(r *http.Request) (p params, err error) {
	p = defaultValue

	err = r.ParseForm()
	if err != nil {
		return
	}

	if v := r.FormValue("cycles"); v != "" {
		p.cycles, err = strconv.Atoi(v)
		if err != nil {
			return p, fmt.Errorf("cycles parameter should be an integer value. %s", err)
		}
	}

	if v := r.FormValue("res"); v != "" {
		p.res, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return p, fmt.Errorf("res parameter should be a float value. %s", err)
		}
	}

	if v := r.FormValue("size"); v != "" {
		p.size, err = strconv.Atoi(v)
		if err != nil {
			return p, fmt.Errorf("size parameter should be an integer value. %s", err)
		}
	}

	if v := r.FormValue("nframes"); v != "" {
		p.nframes, err = strconv.Atoi(v)
		if err != nil {
			return p, fmt.Errorf("cynframescles parameter should be an integer value. %s", err)
		}
	}

	if v := r.FormValue("delay"); v != "" {
		p.delay, err = strconv.Atoi(v)
		if err != nil {
			return p, fmt.Errorf("delay parameter should be an integer value. %s", err)
		}
	}

	return p, nil
}
