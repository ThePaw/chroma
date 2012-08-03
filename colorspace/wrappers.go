// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// Wrapper functions for multi-stage transformations

func RGBToLab(r, g, b float64) (l, a, bb float64) {
	x, y, z := RGBToXYZ(r, g, b)
	l, a, bb = XYZToLab(x, y, z)
	return
}

func LabToRGB(l, a, bb float64) (r, g, b float64) {
	x, y, z := LabToXYZ(l, a, bb)
	r, g, b = XYZToRGB(x, y, z)
	return
}

func RGBToLuv(r, g, b float64) (l, u, v float64) {
	x, y, z := RGBToXYZ(r, g, b)
	l, u, v = XYZToLuv(x, y, z)
	return
}

func LuvToRGB(l, u, v float64) (r, g, b float64) {
	x, y, z := LuvToXYZ(l, u, v)
	r, g, b = XYZToRGB(x, y, z)
	return
}

func RGBToLch(r, g, b float64) (l, c, h float64) {
	x, y, z := RGBToXYZ(r, g, b)
	l, c, h = XYZToLch(x, y, z)
	return
}

func LchToRGB(l, c, h float64) (r, g, b float64) {
	x, y, z := LchToXYZ(l, c, h)
	r, g, b = XYZToRGB(x, y, z)
	return
}

func RGBToLMS(r, g, b float64) (l, m, s float64) {
	x, y, z := RGBToXYZ(r, g, b)
	l, m, s = XYZToLMS(x, y, z)
	return
}

func LMSToRGB(l, m, s float64) (r, g, b float64) {
	x, y, z := LMSToXYZ(l, m, s)
	r, g, b = XYZToRGB(x, y, z)
	return
}
