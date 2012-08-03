// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Attention: This test took 195196.846s on Intel® Core™2 CPU 6600 @ 2.40GHz for 12-bit depth (step = 1.0/2048)

package colorspace

import (
	"fmt"
	"testing"
	"math"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-3// accuracy
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

func TestRGBToAnyRoundtrip(t *testing.T) {
	type fn func(r, g, b float64) (x, y, z float64)
	f := [][]fn{{RGBToXYZ, XYZToRGB},
		{RGBToLMS, LMSToRGB},
		{RGBToLch, LchToRGB},
		{RGBToHSI, HSIToRGB},
		{RGBToHSL, HSLToRGB},
		{RGBToHSV, HSVToRGB},
		{RGBToJPEG_YCbCr, JPEG_YCbCrToRGB},
		{RGBToLab, LabToRGB},
		{RGBToLuv, LuvToRGB},
		{RGBToYDbDr, YDbDrToRGB},
		{RGBToYIQ, YIQToRGB},
		{RGBToYPbPr, YPbPrToRGB},
		{RGBToYUV, YUVToRGB},
		{RGBToYCbCr, YCbCrToRGB}}
	nFunc := len(f)
	maxSqErr := 0.0
	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step := 1.0/32
		for r0 := 0.0; r0 <= 1.0; r0 += step {
			for g0 := 0.0; g0 < 1.0; g0 += step {
				for b0 := 0.0; b0 < 1.0; b0 += step {
					x, y, z := f[i][0](r0, g0, b0)
					r1, g1, b1 := f[i][1](x, y, z)
					if !(check_diff(r0, r1) && check_diff(g0, g1) && check_diff(b0, b1)) {
						t.Errorf("r0, g0, b0 = %f, %f, %f   r1, g1, b1 = %f, %f, %f\n", r0, g0, b0, r1, g1, b1)
					}
					dr := (r0-r1)*(r0-r1)
					dg := (g0-g1)*(g0-g1)
					db := (b0-b1)*(b0-b1)
					d:= dr+dg+db
					if d > maxSqErr {
						maxSqErr = d
					}
				}
			}
		}
	}
	fmt.Println("maxErr = ", math.Pow(maxSqErr, 0.5))
}
