// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb8

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y uint8) bool {
	const acc uint8 = 20 // accuracy
	var d uint8
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
	var (
		r0, g0, b0, step             int
		r0w, g0w, b0w, r1w, g1w, b1w uint8
	)
	type fn func(r, g, b uint8) (x, y, z uint8)
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
	maxSqErr := 0
	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step = 1
		for r0 = 0; r0 <= 255; r0 += step {
			for g0 = 0; g0 <= 255; g0 += step {
				for b0 = 0; b0 <= 255; b0 += step {
					r, g, b := uint8(r0), uint8(g0), uint8(b0)
					x, y, z := f[i][0](r, g, b)
					r1, g1, b1 := f[i][1](x, y, z)
					if !(check_diff(r, r1) && check_diff(g, g1) && check_diff(b, b1)) {
						fmt.Printf(" fn # %d: x, y, z = %d, %d, %d  \n", i, x, y, z)
						t.Errorf(" fn # %d: r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r0, g0, b0, r1, g1, b1)
					}
					dr := (r - r1) * (r - r1)
					dg := (g - g1) * (g - g1)
					db := (b - b1) * (b - b1)
					d := dr + dg + db
					if int(d) > maxSqErr {
						r0w, g0w, b0w = r, g, b
						r1w, g1w, b1w = r1, g1, b1
						maxSqErr = int(d)
					}
				}
			}
		}
	}
	fmt.Println("maxErr = ", math.Pow(float64(maxSqErr), 0.5))
	fmt.Println("RGB 0 = ", r0w, g0w, b0w)
	fmt.Println("RGB 1 = ", r1w, g1w, b1w)
}
