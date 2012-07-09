package chroma

import (
	"fmt"
	"testing"
)

func check(x, y float64) bool {
	const acc float64 = 1e-6 // accuracy
	var z float64
	if x/y > 1.00 {
		z = y / x
	} else {
		z = x / y
	}
	if 1-z > acc {
		return false
	}
	return true
}

// test against known values
func TestSRGB2XYZ(t *testing.T) {
	fmt.Println("test of sRGB to XYZ")
	r := uint8(50)
	g := uint8(55)
	b := uint8(100)
	x, y, z := (SRGBToXYZ(r, g, b))
	x2 := (4.982)
	y2 := (4.331)
	z2 := (12.630)
		fmt.Println(x, x2)
		fmt.Println(y, y2)
		fmt.Println(z, z2)
}

