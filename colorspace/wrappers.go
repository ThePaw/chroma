// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// Wrapper functions for multi-stage transformations

func Rgb2Lab(r, g, b float64) (l, a, bb float64) {
	x, y, z := Rgb2Xyz(r, g, b)
	l, a, bb = Xyz2Lab(x, y, z)
	return
}

func Lab2Rgb(l, a, bb float64) (r, g, b float64) {
	x, y, z := Lab2Xyz(l, a, bb)
	r, g, b = Xyz2Rgb(x, y, z)
	return
}

func Rgb2Luv(r, g, b float64) (l, u, v float64) {
	x, y, z := Rgb2Xyz(r, g, b)
	l, u, v = Xyz2Luv(x, y, z)
	return
}

func Luv2Rgb(l, u, v float64) (r, g, b float64) {
	x, y, z := Luv2Xyz(l, u, v)
	r, g, b = Xyz2Rgb(x, y, z)
	return
}

func Rgb2Lch(r, g, b float64) (l, c, h float64) {
	x, y, z := Rgb2Xyz(r, g, b)
	l, c, h = Xyz2Lch(x, y, z)
	return
}

func Lch2Rgb(l, c, h float64) (r, g, b float64) {
	x, y, z := Lch2Xyz(l, c, h)
	r, g, b = Xyz2Rgb(x, y, z)
	return
}

func Rgb2Lms(r, g, b float64) (l, m, s float64) {
	x, y, z := Rgb2Xyz(r, g, b)
	l, m, s = Xyz2Lms(x, y, z)
	return
}

func Lms2Rgb(l, m, s float64) (r, g, b float64) {
	x, y, z := Lms2Xyz(l, m, s)
	r, g, b = Xyz2Rgb(x, y, z)
	return
}
