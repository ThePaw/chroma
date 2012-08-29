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

func TestRgbToXYZWithin01(t *testing.T) {
	const eps float64 = 0.174 // tolerance
lo:=0.0-eps
hi:=1.0+eps
	type fn func(r, g, b float64) (x, y, z float64)
	f := [][]fn{
		{AdobeToXYZ_D50, XYZToAdobe_D50},
		{AppleToXYZ_D50, XYZToApple_D50},
		{BruceToXYZ_D50, XYZToBruce_D50},
		{CieToXYZ_D50, XYZToCie},
		{NTSCToXYZ_D50, XYZToNTSC_D50},
		{PALToXYZ_D50, XYZToPAL_D50},
		{SMPTE_CToXYZ_D50, XYZToSMPTE_C_D50},
		{SRGBToXYZ_D50, XYZToSRGB_D50},
		{AdobeToXYZ, XYZToAdobe},
		{AppleToXYZ, XYZToApple},
		{BestToXYZ, XYZToBest},
		{BetaToXYZ, XYZToBeta},
		{BruceToXYZ, XYZToBruce},
		{CIEToXYZ, XYZToCIE},
		{ColorMatchToXYZ, XYZToColorMatch},
		{DonToXYZ, XYZToDon},
		{ECIToXYZ, XYZToECI},
		{EktaSpaceToXYZ, XYZToEktaSpace},
		{NTSCToXYZ, XYZToNTSC},
		{PALToXYZ, XYZToPAL},
		{ProPhotoToXYZ, XYZToProPhoto},
		{SMPTE_CToXYZ, XYZToSMPTE_C},
		{SRGBToXYZ, XYZToSRGB},
		{WGamutToXYZ, XYZToWGamut}}
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

