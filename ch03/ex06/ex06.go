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
	)

	// 画像を4倍のサイズに
	img4 := image.NewRGBA(image.Rect(0, 0, width*2, height*2))
	for py := 0; py < (height * 2); py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < (width * 2); px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img4.Set(px, py, mandelbrot(z))
		}
	}

	// サブピクセルの色を決める
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py = py + 2 {
		for px := 0; px < width*2; px = px + 2 {
			var c = []color.RGBA{
				img4.RGBAAt(px, py),
				img4.RGBAAt(px+1, py),
				img4.RGBAAt(px, py+1),
				img4.RGBAAt(px+1, py+1),
			}
			var pr, pg, pb, pa int
			for n := 0; n < 4; n++ {
				pr += int(c[n].R)
				pg += int(c[n].G)
				pb += int(c[n].B)
				pa += int(c[n].A)
			}
			img.SetRGBA(px/2, py/2, color.RGBA{uint8(pr / 4), uint8(pg / 4), uint8(pb / 4), uint8(pa / 4)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast * (n + 8), contrast * (n + 4), contrast * n, 255}
		}
	}
	return color.RGBA{255, 255, 255, 255}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
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
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
