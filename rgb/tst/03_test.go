// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb

import (
	"fmt"
	"testing"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-3 // accuracy
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

// Test whether XYZ is always within [0.0, 1.0]. . 

func TestRgb2XyzWithin01(t *testing.T) {
	const eps float64 = 0.174 // tolerance
lo:=0.0-eps
hi:=1.0+eps
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
	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step := 1.0 / 256
		for r0 := 0.0; r0 <= 1.0; r0 += step {
			for g0 := 0.0; g0 < 1.0; g0 += step {
				for b0 := 0.0; b0 < 1.0; b0 += step {
					x, y, z := f[i][0](r0, g0, b0)
					if !(x >= lo && x <= hi && y >= lo && y <= hi && z >= lo && z <= hi) {
						t.Errorf(" fn # %d: RGB = %f, %f, %f   XYZ = %f, %f, %f\n", i, r0, g0, b0, x, y, z)
					}
				}
			}
		}
	}
}

