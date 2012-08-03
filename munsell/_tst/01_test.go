// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package munsell

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

func TestHvc2Xyz(t *testing.T) {
h := "5B"
v, c := 0.2, 2.0
x0, y0, z0 := 0.00145008306932927,0.00235643659115854,0.00962998199615853
x, y, z := Munsell2Xyz(h , v, c)
					if !(check_diff(x0, x) && check_diff(y0, y) && check_diff(z0, z)) {
						t.Errorf(" Munsell2Xyz: \tXYZ 0 = %f, %f, %f   \n\t\t\tXYZ 1 = %f, %f, %f\n", x0, y0, z0, x, y, z)
					}

}
