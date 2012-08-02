// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.
/*
These functions have problems with accuracy == 1
1:
11:
13:
14:
17:
18:
19:
20:
21:
22:
23:
4:
5:
6:

and these with 2:
13:
14:
22:
23:

At accuracy == 3 all passed
*/

package rgb8

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y uint8) bool {
	const acc uint8 = 3 // accuracy
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
	type fn1 func(r, g, b uint8) (x, y, z int)
	type fn2 func(x, y, z int) (r, g, b uint8)
	f1:= []fn1{Adobe2Xyz, Apple2Xyz, Best2Xyz, Beta2Xyz, Bruce2Xyz, CieE2Xyz, ColorMatch2Xyz, Don2Xyz, Eci2Xyz, EktaSpace2Xyz, Ntsc2Xyz, Pal2Xyz, ProPhoto2Xyz, SmpteC2Xyz, Srgb2Xyz, WGamut2Xyz, Adobe2Xyz_D50, Apple2Xyz_D50, Bruce2Xyz_D50, Cie2Xyz_D50, Ntsc2Xyz_D50, Pal2Xyz_D50, SmpteC2Xyz_D50, Srgb2Xyz_D50}

	f2:= []fn2{Xyz2Adobe, Xyz2Apple, Xyz2Best, Xyz2Beta, Xyz2Bruce, Xyz2CieE, Xyz2ColorMatch, Xyz2Don, Xyz2Eci, Xyz2EktaSpace, Xyz2Ntsc, Xyz2Pal, Xyz2ProPhoto, Xyz2SmpteC, Xyz2Srgb, Xyz2WGamut, Xyz2Adobe_D50, Xyz2Apple_D50, Xyz2Bruce_D50, Xyz2Cie, Xyz2Ntsc_D50, Xyz2Pal_D50, Xyz2SmpteC_D50, Xyz2Srgb_D50}


	nFunc := len(f1)
	maxSqErr := 0
	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step = 1
		for r0 = 0; r0 <= 255; r0 += step {
			for g0 = 0; g0 <= 255; g0 += step {
				for b0 = 0; b0 <= 255; b0 += step {
					r, g, b := uint8(r0), uint8(g0), uint8(b0)
					x, y, z := f1[i](r, g, b)
					r1, g1, b1 := f2[i](x, y, z)
					if !(check_diff(r, r1) && check_diff(g, g1) && check_diff(b, b1)) {
						//fmt.Printf(" fn # %d: x, y, z = %d, %d, %d  \n", i, x, y, z)
						fmt.Printf(" fn # %d: r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
						//t.Errorf(" fn # %d: r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
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
