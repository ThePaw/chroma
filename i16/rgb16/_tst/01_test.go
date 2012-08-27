// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

import (
	"fmt"
	"math"
	"testing"
)

func check_diff(x, y uint16) bool {
	const acc uint16 = 3 // accuracy
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
	var (
		i, j, k, l, step	int
		r, g, b              uint16
		rw, gw, bw, r1w, g1w, b1w uint16
	)
	type fn1 func(r, g, b uint16) (x, y, z int)
	type fn2 func(x, y, z int) (r, g, b uint16)
	f1 := []fn1{AdobeToXYZ, AppleToXYZ, BestToXYZ, BetaToXYZ, BruceToXYZ, CIEToXYZ, ColorMatchToXYZ, DonToXYZ, EciToXYZ, EktaSpaceToXYZ, NTSCToXYZ, PALToXYZ, ProPhotoToXYZ, SMPTE_CToXYZ, SRGBToXYZ, WGamutToXYZ, AdobeToXYZ_D50, AppleToXYZ_D50, BruceToXYZ_D50, CieToXYZ_D50, NTSCToXYZ_D50, PALToXYZ_D50, SMPTE_CToXYZ_D50, SRGBToXYZ_D50}

	f2 := []fn2{XYZToAdobe, XYZToApple, XYZToBest, XYZToBeta, XYZToBruce, XYZToCIE, XYZToColorMatch, XYZToDon, XYZToEci, XYZToEktaSpace, XYZToNTSC, XYZToPAL, XYZToProPhoto, XYZToSMPTE_C, XYZToSRGB, XYZToWGamut, XYZToAdobe_D50, XYZToApple_D50, XYZToBruce_D50, XYZToCie_D50, XYZToNTSC_D50, XYZToPAL_D50, XYZToSMPTE_C_D50, XYZToSRGB_D50}

tests := 0
	nFunc := len(f1)
	maxSqErr := 0
	for i = 0; i < nFunc; i++ {
		fmt.Println("Testing fn #", i)
		step = 512
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
					tests++
					r1, g1, b1 := f2[i](x, y, z)
					if !(check_diff(r, r1) && check_diff(g, g1) && check_diff(b, b1)) {
						//fmt.Printf(" fn # %d: x, y, z = %d, %d, %d  \n", i, x, y, z)
						//fmt.Printf(" fn # %d: r, g, b = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
						t.Errorf(" fn # %d: r, g, b = %d, %d, %d   r1, g1, b1 = %d, %d, %d\n", i, r, g, b, r1, g1, b1)
					}
					dr := (r - r1) * (r - r1)
					dg := (g - g1) * (g - g1)
					db := (b - b1) * (b - b1)
					d := dr + dg + db
					if int(d) > maxSqErr {
						rw, gw, bw = r, g, b
						r1w, g1w, b1w = r1, g1, b1
						maxSqErr = int(d)
					}
				}
			}
		}
	}
	fmt.Println("maxErr = ", math.Pow(float64(maxSqErr), 0.5))
	fmt.Println("RGB 0 = ", rw, gw, bw)
	fmt.Println("RGB 1 = ", r1w, g1w, b1w)
	fmt.Println("Number of tests = ", tests)
}
