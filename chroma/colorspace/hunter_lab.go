// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import (
	"math"
)

// HunterLabToXYZ converts an Hunter Lab triple to an XYZ triple.
func HunterLabToXYZ(hL, ha, hb float64) (float64, float64, float64) {
	y := hL / 10
	x := ha / 17.5 * hL / 10
	z := hb / 7 * hL / 10

	y = math.Pow(y, 2)
	x = (x + y) / 1.02
	z = -(z - y) / 0.847
	return x, y, z
}

// XYZToHunterLab converts an XYZ triple to an Hunter Lab triple.
func XYZToHunterLab(x, y, z float64) (float64, float64, float64) {
	hL := 10 * math.Sqrt(y)
	ha := 17.5 * (((1.02 * x) - y) / math.Sqrt(y))
	hb := 7 * ((y - (0.847 * z)) / math.Sqrt(y))
	return hL, ha, hb
}
