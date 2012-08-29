// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb8

import (
	"fmt"
	"testing"
)

func check_diff(x, y int) bool {
	const acc int = 3 // accuracy
	var d int
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

func TestRgb(t *testing.T) {
var r, g, b uint8
r, g, b = 33, 68, 27
x0, y0, z0 := 144042000, 213747000, 127289000
x, y, z := AdobeToXYZ(r, g, b)
fmt.Println("coeff = ",  float64(x)/ float64(x0) )
fmt.Println("coeff = ",  float64(y)/ float64(y0) )
fmt.Println("coeff = ",  float64(z)/ float64(z0) )
r1, g1, b1 := XYZToAdobe(x, y, z)
fmt.Println("rgb ",  r, g, b )
fmt.Println("rgb ",  r1, g1, b1 )

}