// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.
// With step = 32 it will run about 3 days. Be careful when adjusting step: step /=2  --> time *=8
// You will have to run 'go test -test.timeout 120h'.
package rgb16

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y, acc uint16) bool {
	var d uint16
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
	const acc uint16 = 4 // accuracy
	var (
		i, j, k, l, step             int
		r, g, b                      uint16
		r0w, g0w, b0w, r1w, g1w, b1w uint16
	)
	type fn1 func(r, g, b uint16) (x, y, z int)
	type fn2 func(x, y, z int) (r, g, b uint16)
	f1 := []fn1{AdobeToXYZ, AppleToXYZ, BestToXYZ, BetaToXYZ, BruceToXYZ, CIEToXYZ, ColorMatchToXYZ, DonToXYZ, ECIToXYZ, EktaSpaceToXYZ, NTSCToXYZ, PALToXYZ, ProPhotoToXYZ, SMPTE_CToXYZ, SRGBToXYZ, WGamutToXYZ, AdobeToXYZ_D50, AppleToXYZ_D50, BruceToXYZ_D50, CIEToXYZ_D50, NTSCToXYZ_D50, PALToXYZ_D50, SMPTE_CToXYZ_D50, SRGBToXYZ_D50}

	f2 := []fn2{XYZToAdobe, XYZToApple, XYZToBest, XYZToBeta, XYZToBruce, XYZToCIE, XYZToColorMatch, XYZToDon, XYZToECI, XYZToEktaSpace, XYZToNTSC, XYZToPAL, XYZToProPhoto, XYZToSMPTE_C, XYZToSRGB, XYZToWGamut, XYZToAdobe_D50, XYZToApple_D50, XYZToBruce_D50, XYZToCIE_D50, XYZToNTSC_D50, XYZToPAL_D50, XYZToSMPTE_C_D50, XYZToSRGB_D50}

	nTests := 0
	fun := 0
	step = 32

	nFunc := len(f1)
	maxSqErr := 0
	fmt.Println("### TestRgb2RgbRoundtrip ###")
	for i = 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		err := false
		for j = 0; j <= 65536; j += step {

			for k = 0; k <= 65536; k += step {
				for l = 0; l <= 65536; l += step {
					if j > 65535 {
						j = 65535
					}
					if k > 65535 {
						k = 65535
					}
					if l > 65535 {
						l = 65535
					}

					r, g, b = uint16(j), uint16(k), uint16(l)
					x, y, z := f1[i](r, g, b)
					//		fmt.Println("rgb: ", r,g,b)
					nTests++
					r1, g1, b1 := f2[i](x, y, z)
					if !(check_diff(r, r1, acc) && check_diff(g, g1, acc) && check_diff(b, b1, acc)) {
						// fmt.Printf(" fn # %d: x, y, z = %d, %d, %d  \n", i, x, y, z)
						// fmt.Printf(" fn # %d: r, g, b = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
						err = true
						// t.Errorf(" fn # %d: r, g, b = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
					}
					dr := (r - r1) * (r - r1)
					dg := (g - g1) * (g - g1)
					db := (b - b1) * (b - b1)
					d := dr + dg + db
					if int(d) > maxSqErr {
						r0w, g0w, b0w = r, g, b
						r1w, g1w, b1w = r1, g1, b1
						fun = i
						maxSqErr = int(d)
					}
				}
			}
		}
		if err {
			t.Errorf(" fn # %d failed \n", i)
		}

	}
	fmt.Println("Number of tests = ", nTests)
	fmt.Println("Accuracy = ", acc)
	fmt.Println("Step = ", step)
	fmt.Println("maxErr = ", math.Pow(float64(maxSqErr), 0.5))
	fmt.Println("Worst case (appeared in function #", fun, "): ")
	fmt.Println("RGB 0 = ", r0w, g0w, b0w)
	fmt.Println("RGB 1 = ", r1w, g1w, b1w)
}
