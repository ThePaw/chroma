// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import "math"

// CIE_LabToCIE_LCH converts a CIE L*ab triple to a CIE L*CH° triple.
func CIE_LabToCIE_LCH(l, a, b float64) (float64, float64, float64) {
	h := math.Atan2(b, a) //Quadrant by signs

	if h > 0 {
		h = (h / math.Pi) * 180.0
	} else {
		h = 360.0 - (math.Abs(h)/math.Pi)*180.0
	}
	c := math.Sqrt(a*a + b*b)
	return l, c, h
}

// CIE_LabToCIE_LCH converts a CIE L*CH° triple to a CIE L*ab triple.
func CIE_LCHToCIE_Lab(l, c, h float64) (float64, float64, float64) {
	//CIE-H° from 0 to 360°
	deg2rad := math.Pi / 180.0
	a := math.Cos(deg2rad*h) * c
	b := math.Sin(deg2rad*h) * c
	return l, a, b
}
