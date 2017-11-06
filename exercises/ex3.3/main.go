// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// colors used to blend
var (
	peak, _   = colorful.Hex("#ff0000")
	valley, _ = colorful.Hex("#0000ff")
)

type polygon struct {
	ax, ay float64
	bx, by float64
	cx, cy float64
	dx, dy float64
	zavg   float64
}

func main() {
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

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for _, p := range polygons {
		c := colorHeight(p.zavg, zmin, zmax)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
			p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, c.Hex())
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
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
