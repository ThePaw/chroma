// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb

import (
	"testing"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-3// accuracy
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

func TestRgb2Xyz(t *testing.T) {
r, g, b := 0.1, 0.2, 0.7
x0, y0, z0 := 0.206480, 0.200497, 0.535513
x, y, z := Adobe2Xyz_D50(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Adobe2Xyz_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

x0, y0, z0 = 0.226513, 0.207899, 0.710617
x, y, z = Adobe2Xyz(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Adobe2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}


x0, y0, z0 = 0.219777, 0.210607, 0.509877
x, y, z = Apple2Xyz_D50(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Apple2Xyz_D50: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

x0, y0, z0 = 0.237367, 0.217194, 0.676479
x, y, z = Apple2Xyz(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Apple2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

x0, y0, z0 = 0.193074, 0.194250, 0.572890
x, y, z = Best2Xyz(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Best2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}


x0, y0, z0 = 0.184910, 0.186143, 0.557297
x, y, z = Beta2Xyz(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Beta2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

x0, y0, z0 = 0.184910, 0.186143, 0.557297
x, y, z = EktaSpace2Xyz(r, g, b)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" EktaSpace2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

}


