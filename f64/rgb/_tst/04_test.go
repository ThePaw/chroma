// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.
// Testing values from http://www.brucelindbloom.com/ColorCalculator.html
// Thanks, Bruce!

package rgb

import (
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

func TestRgbToXYZ(t *testing.T) {

// first, native illuminators

	r, g, b := 0.1, 0.2, 0.7
	x0, y0, z0 := 0.206480, 0.200497, 0.535513
	x, y, z := AdobeToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AdobeToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.226513, 0.207899, 0.710617
	x, y, z = AdobeToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AdobeToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.219777, 0.210607, 0.509877
	x, y, z = AppleToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AppleToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.237367, 0.217194, 0.676479
	x, y, z = AppleToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AppleToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.193074, 0.194250, 0.572890
	x, y, z = BestToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" BestToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.184910, 0.186143, 0.557297
	x, y, z = BetaToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" BetaToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.237654, 0.213619, 0.712229
	x, y, z = BruceToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" BruceToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.251429, 0.187785, 0.694898
	x, y, z = CIEToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" CIEToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.208894, 0.206004, 0.508703
	x, y, z = ColorMatchToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" ColorMatchToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.190816, 0.189005, 0.566427
	x, y, z = DonToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" DonToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.195793, 0.206815, 0.543728
	x, y, z = EciToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" EciToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.182129, 0.176150, 0.556649
	x, y, z = EktaSpaceToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" EktaSpaceToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.235633, 0.227351, 0.794576
	x, y, z = NTSCToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" NTSCToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.236187, 0.213458, 0.685295
	x, y, z = PALToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" PALToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.128753, 0.171239, 0.577647
	x, y, z = ProPhotoToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" ProPhotoToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.246551, 0.222030, 0.694970
	x, y, z = SMPTE_CToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" SMPTE_CToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.239067, 0.214820, 0.690985
	x, y, z = SRGBToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" SRGBToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.194826, 0.182619, 0.551756
	x, y, z = WGamutToXYZ(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" WGamutToXYZ: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

// now with D50

	x0, y0, z0 = 0.206480, 0.200497, 0.535513
	x, y, z = AdobeToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AdobeToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.219777, 0.210607, 0.509877
	x, y, z = AppleToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" AppleToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.218203, 0.206465, 0.536709
	x, y, z = BruceToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" BruceToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.229673, 0.182828, 0.569909
	x, y, z = CieToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" CieToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.201721, 0.217631, 0.550580
	x, y, z = NTSCToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" NTSCToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.218013, 0.206721, 0.516462
	x, y, z = PALToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" PALToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.228583, 0.215353, 0.523772
	x, y, z = SMPTE_CToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" SMPTE_CToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

	x0, y0, z0 = 0.220777, 0.208058, 0.520735
	x, y, z = SRGBToXYZ_D50(r, g, b)
	if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
		t.Errorf(" SRGBToXYZ_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
	}

}
