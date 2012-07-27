// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Attention: This test took 76.464s on Intel® Core™2 CPU 6600 @ 2.40GHz for 8-bit depth all colors
// For every additional bit of depth the time multiplies by 8
// so that to test for all 12-bit colors would take 3.62 days
// 14-bit:  231.99 days
// 16-bit:  40.65 years
// so be careful while adjusting 'step' below.
// Also, you then should have to run 'go test' with '-test.timeout t' flag.


package rgb

import (
	"fmt"
	"math"
	"testing"
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

func TestRgb2RgbRoundtrip(t *testing.T) {
	type fn func(r, g, b float64) (x, y, z float64)
	f := [][]fn{
		{Adobe2Xyz_D50, Xyz2Adobe_D50},
		{Apple2Xyz_D50, Xyz2Apple_D50},
		{Bruce2Xyz_D50, Xyz2Bruce_D50},
		{Cie2Xyz_D50, Xyz2Cie},
		{Ntsc2Xyz_D50, Xyz2Ntsc_D50},
		{Pal2Xyz_D50, Xyz2Pal_D50},
		{SmpteC2Xyz_D50, Xyz2SmpteC_D50},
		{Srgb2Xyz_D50, Xyz2Srgb_D50},
		{Adobe2Xyz, Xyz2Adobe},
		{Apple2Xyz, Xyz2Apple},
		{Best2Xyz, Xyz2Best},
		{Beta2Xyz, Xyz2Beta},
		{Bruce2Xyz, Xyz2Bruce},
		{CieE2Xyz, Xyz2CieE},
		{ColorMatch2Xyz, Xyz2ColorMatch},
		{Don2Xyz, Xyz2Don},
		{Eci2Xyz, Xyz2Eci},
		{EktaSpace2Xyz, Xyz2EktaSpace},
		{Ntsc2Xyz, Xyz2Ntsc},
		{Pal2Xyz, Xyz2Pal},
		{ProPhoto2Xyz, Xyz2ProPhoto},
		{SmpteC2Xyz, Xyz2SmpteC},
		{Srgb2Xyz, Xyz2Srgb},
		{WGamut2Xyz, Xyz2WGamut}}
	nFunc := len(f)
	maxSqErr := 0.0
	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step := 1.0 / 255
		for r0 := 0.0; r0 <= 1.0; r0 += step {
			for g0 := 0.0; g0 < 1.0; g0 += step {
				for b0 := 0.0; b0 < 1.0; b0 += step {
					x, y, z := f[i][0](r0, g0, b0)
					r1, g1, b1 := f[i][1](x, y, z)
					if !(check_diff(r0, r1) && check_diff(g0, g1) && check_diff(b0, b1)) {
						t.Errorf(" fn # %d: r0, g0, b0 = %f, %f, %f   r1, g1, b1 = %f, %f, %f\n", i, r0, g0, b0, r1, g1, b1)
					}
					dr := (r0 - r1) * (r0 - r1)
					dg := (g0 - g1) * (g0 - g1)
					db := (b0 - b1) * (b0 - b1)
					d := dr + dg + db
					if d > maxSqErr {
						maxSqErr = d
					}
				}
			}
		}
	}
	fmt.Println("maxErr = ", math.Pow(maxSqErr, 0.5))
}
