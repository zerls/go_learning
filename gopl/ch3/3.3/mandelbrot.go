package main

import (
	"fmt"
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

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelBrot2(z))
		}
	}
	file, _ := os.Create("mandelBrot.png")
	png.Encode(file, img) // NOTE: ignoring errors

	vs := 3 + 4i
	fmt.Println(cmplx.Abs(vs))
}

func mandelBrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}

func mandelBrot2(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 1.5 {
			//return  color.Gray{255-contrast*n}
			//return  color.RGBA{255-uint8(real(v)),255-uint8(imag(v)),255-uint8(cmplx.Abs(v)),255-contrast*n}
			return color.YCbCr{255 - uint8(real(v)), 255 - uint8(imag(v)), 255 - uint8(cmplx.Abs(v))}
		}
	}

	return color.Black
}
