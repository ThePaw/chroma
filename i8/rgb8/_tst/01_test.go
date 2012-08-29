// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.
/*
These functions have problems with accuracy == 1

and these with 2:

At accuracy == XXXX all passed
*/

package rgb8

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y, acc uint8) bool {
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
	const acc uint8 = 1 // accuracy
	var (
		r0, g0, b0, nTests, step             int
		r0w, g0w, b0w, r1w, g1w, b1w uint8
	)
	type fn1 func(r, g, b uint8) (x, y, z int)
	type fn2 func(x, y, z int) (r, g, b uint8)
	f1 := []fn1{AdobeToXYZ, AppleToXYZ, BestToXYZ, BetaToXYZ, BruceToXYZ, CIEToXYZ, ColorMatchToXYZ, DonToXYZ, ECIToXYZ, EktaSpaceToXYZ, NTSCToXYZ, PALToXYZ, ProPhotoToXYZ, SMPTE_CToXYZ, SRGBToXYZ, WGamutToXYZ, AdobeToXYZ_D50, AppleToXYZ_D50, BruceToXYZ_D50, CieToXYZ_D50, NTSCToXYZ_D50, PALToXYZ_D50, SMPTE_CToXYZ_D50, SRGBToXYZ_D50}

	f2 := []fn2{XYZToAdobe, XYZToApple, XYZToBest, XYZToBeta, XYZToBruce, XYZToCIE, XYZToColorMatch, XYZToDon, XYZToECI, XYZToEktaSpace, XYZToNTSC, XYZToPAL, XYZToProPhoto, XYZToSMPTE_C, XYZToSRGB, XYZToWGamut, XYZToAdobe_D50, XYZToApple_D50, XYZToBruce_D50, XYZToCie_D50, XYZToNTSC_D50, XYZToPAL_D50, XYZToSMPTE_C_D50, XYZToSRGB_D50}

	nFunc := len(f1)
	maxSqErr := 0
	nTests =0

	for i := 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step = 1
		for r0 = 0; r0 <= 255; r0 += step {
			for g0 = 0; g0 <= 255; g0 += step {
				for b0 = 0; b0 <= 255; b0 += step {
					nTests++
					r, g, b := uint8(r0), uint8(g0), uint8(b0)
					x, y, z := f1[i](r, g, b)
					r1, g1, b1 := f2[i](x, y, z)
					if !(check_diff(r, r1, acc) && check_diff(g, g1, acc) && check_diff(b, b1, acc)) {
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

	fmt.Println("Number of tests = ", nTests)
	fmt.Println("Accuracy = ", acc)
	fmt.Println("Step = ", step)
	fmt.Println("maxErr = ", math.Pow(float64(maxSqErr), 0.5))
	fmt.Println("Worst case: ")
	fmt.Println("RGB 0 = ", r0w, g0w, b0w)
	fmt.Println("RGB 1 = ", r1w, g1w, b1w)
}
