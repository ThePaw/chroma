// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

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
var r, g, b uint16
r, g, b = 6554, 13107, 45875

//rgb 

x0, y0, z0 := 226513000, 207899000, 710617000

x, y, z := AdobeToXYZ(r, g, b)
	if !(check_diff(x, x0) && check_diff(y, y0) && check_diff(z, z0)) {
						fmt.Println("x, y, z = ",  x, y, z ,  "x0, y0, z0 = ",x0, y0, z0 )


r1, g1, b1 := XYZToAdobe(x, y, z)
						fmt.Println("r,g,b = ",  r,g,b ,  "r1, g1, b1 = ",r1, g1, b1)

}
}

