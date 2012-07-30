// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.


package rgb

import (
	"fmt"
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
var r, g, b float64
r, g, b = 0.1, 0.2, 0.7

x, y, z := Adobe2Xyz(r, g, b)
	fmt.Printf(" x, y, z = %f, %f, %f \n", x, y, z )
}

