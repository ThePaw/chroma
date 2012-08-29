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
r, g, b = 6425, 13107, 46003

//rgb 

x0, y0, z0 := 225752000, 207464000, 712507000

x, y, z := AdobeToXYZ(r, g, b)
	if !(check_diff(x, x0) && check_diff(y, y0) && check_diff(z, z0)) {
						fmt.Println("x, y, z = ",  x, y, z ,  "x0, y0, z0 = ",x0, y0, z0 )


r1, g1, b1 := XYZToAdobe(x, y, z)
						fmt.Println("r,g,b = ",  r,g,b ,  "r1, g1, b1 = ",r1, g1, b1)


r, g, b = 65535, 65535, 65535
x, y, z = AdobeToXYZ(r, g, b)
						fmt.Println("x, y, z = ",  x, y, z  )

// 33, 68, 27
r, g, b = 8481, 17476, 6939
x0, y0, z0 = 144042000, 213747000, 127289000
x, y, z = AdobeToXYZ(r, g, b)
fmt.Println("coeff = ",  float64(x)/ float64(x0) )
fmt.Println("coeff = ",  float64(y)/ float64(y0) )
fmt.Println("coeff = ",  float64(z)/ float64(z0) )

// 155, 16, 99
//r, g, b = 155*257, 16*257, 99*257
r, g, b = 39835, 4112, 25443
x0, y0, z0 = 435265000, 249346000, 405651000
x, y, z = AdobeToXYZ(r, g, b)
fmt.Println("coeff = ",  float64(x)/ float64(x0) )
fmt.Println("coeff = ",  float64(y)/ float64(y0) )
fmt.Println("coeff = ",  float64(z)/ float64(z0) )
}
}

