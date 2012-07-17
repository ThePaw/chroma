// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

import (
	"fmt"
	"testing"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-8 // accuracy
	var d float64
	if x >= y {
		d = x - y
	} else {
		d = y - x
	}
	if d > acc {
		return false
	}
	return true
}

// Test that a subset of RGB space can be converted to XYZ and back to within tolerance. 

func TestRgb2AnyRoundtrip(t *testing.T) {
	type fn func(r, g, b float64) (x, y, z float64)
	nFunc := 14
	f := [][]fn{{Rgb2Xyz, Xyz2Rgb},
		{Rgb2Lms, Lms2Rgb},
		{Rgb2Lch, Lch2Rgb},
		{Rgb2Hsi, Hsi2Rgb},
		{Rgb2Hsl, Hsl2Rgb},
		{Rgb2Hsv, Hsv2Rgb},
		{Rgb2Jycbcr, Jycbcr2Rgb},
		{Rgb2Lab, Lab2Rgb},
		{Rgb2Luv, Luv2Rgb},
		{Rgb2Ydbdr, Ydbdr2Rgb},
		{Rgb2Yiq, Yiq2Rgb},
		{Rgb2Ypbpr, Ypbpr2Rgb},
		{Rgb2Yuv, Yuv2Rgb},
		{Rgb2Ycbcr, Ycbcr2Rgb}}

	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		for r0 := 0.0; r0 < 1.0; r0 += 0.01 {
			for g0 := 0.0; g0 < 1.0; g0 += 0.01 {
				for b0 := 0.0; b0 < 1.0; b0 += 0.01 {
					x, y, z := f[i][0](r0, g0, b0)
					r1, g1, b1 := f[i][1](x, y, z)
					if !(check_diff(r0, r1) && check_diff(g0, g1) && check_diff(b0, b1)) {
						t.Fatalf("r0, g0, b0 = %f, %f, %f   r1, g1, b1 = %f, %f, %f\n", r0, g0, b0, r1, g1, b1)
					}
				}
			}
		}
	}
}
