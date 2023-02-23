package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/muesli/gamut"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	palette, _ := gamut.Generate(50, gamut.PastelGenerator{})
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			iterations := mandelbrot(z)
			if iterations == -1 {
				img.Set(px, py, color.Black)
			} else {
				img.Set(px, py, palette[iterations%len(palette)])
			}
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) int {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return int(n)
		}

	}
	return -1
}
