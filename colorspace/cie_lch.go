// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

import "math"

// Lab2Lch converts a CIE L*ab triple to a CIE L*CH° triple.
func Lab2Lch(l, a, b float64) (float64, float64, float64) {
	h := math.Atan2(b, a) //Quadrant by signs

	if h > 0 {
		h = (h / math.Pi) * 180.0
	} else {
		h = 360.0 - (math.Abs(h)/math.Pi)*180.0
	}
	c := math.Sqrt(a*a + b*b)
	return l, c, h
}

// Lab2Lch converts a CIE L*CH° triple to a CIE L*ab triple.
func Lch2Lab(l, c, h float64) (float64, float64, float64) {
	//CIE-H° from 0 to 360°
	deg2rad := math.Pi / 180.0
	a := math.Cos(deg2rad*h) * c
	b := math.Sin(deg2rad*h) * c
	return l, a, b
}

// Xyz2Lch converts CIE XYZ to CIE L*C*H*. 
// CIE L*C*H* is related to CIE L*a*b* by
//    a* = C* math.Cos(H* pi/180),
//    b* = C* math.Sin(H* pi/180).
func Xyz2Lch(x, y, z float64) (l, c, h float64) {
	l, a, b := Xyz2Lab(x, y, z)
	c = math.Sqrt(a*a + b*b)
	h = math.Atan2(b, a) * 180.0 / math.Pi
	if h < 0 {
		h += 360
	}
	return
}

// Lch2Xyz converts CIE L*C*H* to CIE XYZ. 
func Lch2Xyz(l, c, h float64) (x, y, z float64) {
	a := c * math.Cos(h*(math.Pi/180.0))
	b := c * math.Sin(h*(math.Pi/180.0))
	x, y, z = Lab2Xyz(l, a, b)
	return
}
