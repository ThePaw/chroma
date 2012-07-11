// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

/*
// Luv2Xyz converts a CIE-L*uv triple to a CIE-XYZ triple.
func Luv2Xyz(l, u, v float64) (x, y, z float64) {
	// Observer= 2°, Illuminant= D65
	ref_x := 95.047
	ref_y := 100.000
	ref_z := 108.883

	y = (l + 16) / 116.0
	if y*y*y > 0.008856 {
		y = y * y * y
	} else {
		y = (y - 16/116) / 7.787
	}
	ref_u := (4 * ref_x) / (ref_x + (15 * ref_y) + (3 * ref_z))
	ref_v := (9 * ref_y) / (ref_x + (15 * ref_y) + (3 * ref_z))

	u = u/(13*l) + ref_u
	v = v/(13*l) + ref_v

	y = y * 100
	x = -(9 * y * u) / ((u-4)*v - u*v)
	z = (9*y - (15 * v * y) - (v * x)) / (3 * v)
	return
}

// Xyz2Luv converts a CIE-XYZ triple to a CIE-L*uv triple.
func Xyz2Luv(x, y, z float64) (l, u, v float64) {
	// Observer= 2°, Illuminant= D65
	ref_x := 95.047
	ref_y := 100.000
	ref_z := 108.883

	u = (4 * x) / (x + (15 * y) + (3 * z))
	v = (9 * y) / (x + (15 * y) + (3 * z))

	y = y / 100.0
	if y > 0.008856 {
		y = math.Pow(y, 1.0/3.0)
	} else {
		y = (7.787 * y) + (16 / 116.0)
	}

	ref_u := (4 * ref_x) / (ref_x + (15 * ref_y) + (3 * ref_z))
	ref_v := (9 * ref_y) / (ref_x + (15 * ref_y) + (3 * ref_z))

	l = (116 * y) - 16
	u = 13 * l * (u - ref_u)
	v = 13 * l * (v - ref_v)
	return
}
*/

// Luv2Xyz converts a CIE L*uv triple to a CIE XYZ triple.
func Luv2Xyz(l, u, v float64) (x, y, z float64) {
	y = (l + 16) / 116.0
	y = wY * labInv(y)

	if l != 0.0 {
		u /= l
		v /= l
	}

	u = u/13.0 + wU
	v = v/13.0 + wV
	x = y * ((9.0 * u) / (4.0 * v))
	z = y * ((3.0-0.75*u)/v - 5.0)
	return
}

// Xyz2Luv converts a CIE XYZ triple to a CIE L*uv triple.
func Xyz2Luv(x, y, z float64) (l, u, v float64) {
	var u1, v1 float64
	den := x + 15*y + 3*z
	if den > 0 {
		u1 = (4 * x) / den
		v1 = (9 * y) / den
	} else {
		u1 = 0
		v1 = 0
	}
	y /= wY
	y = lab(y)
	l = 116*y - 16
	u = 13 * (l) * (u1 - wU)
	v = 13 * (l) * (v1 - wV)
	return
}
