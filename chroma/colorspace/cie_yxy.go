// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

// YxyToXYZ converts an Yxy triple to an XYZ triple.
func YxyToXYZ(yY, x, y float64) (float64, float64, float64) {
	//yY from 0 to 100
	//x from 0 to 1
	//y from 0 to 1

	xx := x * (yY / y)
	yy := yY
	zz := (1 - x - y) * (yY / y)

	return xx, yy, zz
}

// XYZToYxy converts an XYZ triple to an Yxy triple.
func XYZToYxy(xx, yy, zz float64) (float64, float64, float64) {
	//X from 0 to 95.047       Observer. = 2°, Illuminant = D65
	//Y from 0 to 100.000
	//Z from 0 to 108.883

	yY := yy
	x := xx / (xx + yy + zz)
	y := yy / (xx + yy + zz)
	return yY, x, y
}
