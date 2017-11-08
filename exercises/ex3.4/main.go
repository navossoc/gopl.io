// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// colors used to blend
var (
	width, height int     = 600, 320 // canvas size in pixels
	cells                 = 100      // number of grid cells
	xyscale       float64            // pixels per x or y unit
	zscale        float64            // pixels per z unit
	peak, _       = colorful.Hex("#ff0000")
	valley, _     = colorful.Hex("#0000ff")
)

type polygon struct {
	ax, ay float64
	bx, by float64
	cx, cy float64
	dx, dy float64
	zavg   float64
}

// http://localhost:8000/?width=1200&height=640&cells=200&peak=ff00ff&valley=00ffff
func main() {
	http.HandleFunc("/", svgHandler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
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
			xyscale = float64(width) / 2 / xyrange
		case "height":
			height, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, "height parameter should be an integer value. %s", err)
				return
			}
			zscale = float64(height) * 0.4
		case "cells":
			cells, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, "cells parameter should be an integer value. %s", err)
				return
			}
		case "peak":
			peak, err = colorful.Hex("#" + v[0])
			if err != nil {
				fmt.Fprintf(w, "peak color parameter should be a RGB color in hex format. %s", err)
				return
			}

		case "valley":
			valley, err = colorful.Hex("#" + v[0])
			if err != nil {
				fmt.Fprintf(w, "valley color parameter should be a RGB color in hex format. %s", err)
				return
			}
		}
	}

	// Set header
	w.Header().Set("Content-Type", "image/svg+xml")

	// Render image
	svg(w)
}

func svg(w io.Writer) {
	var zmax, zmin float64
	polygons := make([]polygon, 0, cells*cells)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, bz, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, cz, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dz, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			// calculates an average height
			zavg := (az + bz + cz + dz) / 4.0
			zmax = math.Max(zavg, zmax)
			zmin = math.Min(zavg, zmin)

			p := polygon{ax, ay, bx, by, cx, cy, dx, dy, zavg}
			polygons = append(polygons, p)
		}
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for _, p := range polygons {
		c := colorHeight(p.zavg, zmin, zmax)
		fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
			p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, c.Hex())
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r

	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}

	return r, true
}

func colorHeight(zavg, zmin, zmax float64) colorful.Color {
	z := (zavg - zmin) / (zmax - zmin)
	return valley.BlendLuv(peak, z)
}

//!-
