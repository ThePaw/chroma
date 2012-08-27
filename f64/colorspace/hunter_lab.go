// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

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

// XYZToHlab converts an XYZ triple to an Hunter Lab triple.
func XYZToHlab(x, y, z float64) (float64, float64, float64) {
	hL := 10 * math.Sqrt(y)
	ha := 17.5 * (((1.02 * x) - y) / math.Sqrt(y))
	hb := 7 * ((y - (0.847 * z)) / math.Sqrt(y))
	return hL, ha, hb
}
