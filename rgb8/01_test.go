// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb8

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y uint8) bool {
	const acc uint8 = 1// accuracy
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
	var r0, g0, b0, step uint8
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
		step = 64
		for r0 = 0; r0 <= 255; r0 += step {
			for g0 = 0; g0 <= 255; g0 += step {
				for b0 = 0; b0 <= 255 ; b0 += step {
					x, y, z := f[i][0](r0, g0, b0)
					r1, g1, b1 := f[i][1](x, y, z)
					if !(check_diff(r0, r1) && check_diff(g0, g1) && check_diff(b0, b1)) {
						t.Errorf(" fn # %d: r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r0, g0, b0, r1, g1, b1)
					}

//						fmt.Printf(" fn # %d: r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r0, g0, b0, r1, g1, b1)
						fmt.Printf(" fn # %d: r0, g0, b0 = %d, %d, %d   x, y, z = %d, %d, %d\n", i, r0, g0, b0, x, y, z)
				dr := (r0 - r1) * (r0 - r1)
					dg := (g0 - g1) * (g0 - g1)
					db := (b0 - b1) * (b0 - b1)
					d := dr + dg + db
					if int(d) > maxSqErr {
						maxSqErr = int(d)
					}
				}
			}
		}
	}
	fmt.Println("maxErr = ", math.Pow(float64(maxSqErr), 0.5))
}

func TestRgb2Rgb(t *testing.T) {
var x, y, z uint8
x, y, z = 29, 12, 143

r1, g1, b1 := Xyz2Adobe_D50(x, y, z)
						fmt.Printf(" r, g, b = %d, %d, %d \n", r1, g1, b1 )
}