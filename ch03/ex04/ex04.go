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
)

const (
	widthDefault, heightDefault = 600, 320                   // canvas size in pixels
	cells                       = 100                        // number of grid cells
	xyrange                     = 30.0                       // axis ranges (-xyrange..+xyrange)
	xyscale                     = widthDefault / 2 / xyrange // pixels per x or y unit
	zscale                      = heightDefault * 0.4        // pixels per z unit
	angle                       = math.Pi / 6                // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//!+http
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w, r.FormValue("width"), r.FormValue("height"), r.FormValue("highColor"), r.FormValue("lowColor"))
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svg(out io.Writer, widthStr string, heightStr string, highColorStr string, lowColorStr string) {
	width, height := widthDefault, heightDefault // canvas size in pixels
	var err error
	if widthStr != "" {
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if heightStr != "" {
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if highColorStr == "" {
		highColorStr = "#ff0000"
	}
	if lowColorStr == "" {
		lowColorStr = "#0000ff"
	}
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	io.WriteString(out, s)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, width, height)
			bx, by, bz := corner(i, j, width, height)
			cx, cy, cz := corner(i, j+1, width, height)
			dx, dy, dz := corner(i+1, j+1, width, height)
			fillColor := "style=\"fill:" + lowColorStr + ";\""
			if az+bz+cz+dz > 0 {
				fillColor = "style=\"fill:" + highColorStr + ";\""
			}
			s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' %s/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
			io.WriteString(out, s)
		}
	}
	s = fmt.Sprintf("</svg>")
	io.WriteString(out, s)
}

func corner(i, j, width, height int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
