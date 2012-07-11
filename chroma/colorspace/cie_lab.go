// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import "math"

func lab(x float64) (y float64) {
	if x >= 8.85645167903563082e-3 {
		y = math.Pow(x, 0.333333333333333)
	} else {
		y = (841.0/108.0)*x + (4.0 / 29.0)
		// y = ( 7.787 * x ) + ( 16.0 / 116.0 )	
	}
	return
}

func labInv(x float64) (y float64) {
	if x >= 0.206896551724137931 {
		y = x * x * x
	} else {
		y = (108.0 / 841.0) * (x - 4.0/29.0)
		//		y = (x - 16.0/116.0) / 7.787
	}
	return
}

// LabToXYZ converts a CIE L*ab triple to a CIE XYZ triple.
func LabToXYZ(l, a, b float64) (x, y, z float64) {
	l = (l + 16)/116
	a = l + a/500
	b = l - b/200
	x = wX * labInv(a)
	y = wY * labInv(l)
	z = wZ * labInv(b)
	return
}

// XYZToLab converts an XYZ triple to a CIE L*ab triple.
func XYZToLab(x, y, z float64) (l, a, b float64) {
	x /= wX
	y /= wY
	z /= wZ
	x = lab(x)
	y = lab(y)
	z = lab(z)
	l = 116*y - 16
	a = 500 * (x - y)
	b = 200 * (y - z)
	return
}
