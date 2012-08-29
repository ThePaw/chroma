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

// Test that a subset of RGB space can be converted to XYZ and back to within tolerance. 

func TestRgb2RgbRoundtrip(t *testing.T) {
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
