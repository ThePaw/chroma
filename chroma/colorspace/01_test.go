package chroma

import (
	"fmt"
	"testing"
)

func check(x, y float64) bool {
	const acc float64 = 1e-2 // accuracy
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
func TestRGB2XYZ(t *testing.T) {
	fmt.Println("test of RGB to XYZ")
	r := uint8(50)
	g := uint8(55)
	b := uint8(100)
	x, y, z := (RGBToXYZ(r, g, b))
	x2 := (4.982)
	y2 := (4.331)
	z2 := (12.630)
	if !check(x, x2) {
		t.Error()
		fmt.Println(x, x2)
	}
	if !check(y, y2) {
		t.Error()
		fmt.Println(y, y2)
	}
	if !check(z, z2) {
		t.Error()
		fmt.Println(z, z2)
	}
}

func TestXYZ2RGB(t *testing.T) {
	fmt.Println("test of XYZ to RGB")
	x, y, z := 56.060,   50.490,    8.160
r, g, b := XYZToRGB(x, y, z)
r2, g2, b2 := 254,   171,    32
v := float64(r)
v2 := float64(r2)
	if !check(v, v2) {
		t.Error()
		fmt.Println("R: ", v, v2)
	}
v = float64(g)
v2 = float64(g2)
	if !check(v, v2) {
		t.Error()
		fmt.Println("G: ", v, v2)
	}
v = float64(b)
v2 = float64(b2)
	if !check(v, v2) {
		t.Error()
		fmt.Println("B: ", v, v2)
	}
}

func TestXYZtoLab(t *testing.T) {
	fmt.Println("test of XYZ to Lab")
	x, y, z := 56.060, 50.490, 8.160
	l, a, b := (XYZToLab(x, y, z))
	l2, a2, b2 := 76.369, 21.174, 74.935
	if !check(l, l2) {
		t.Error()
		fmt.Println(l, l2)
	}
	if !check(a, a2) {
		t.Error()
		fmt.Println(a, a2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println(b, b2)
	}
}

func TestLabToXYZ(t *testing.T) {
	fmt.Println("test of Lab to XYZ")
	l, a, b := 76.369, 21.174, 74.935
	x, y, z := (LabToXYZ(l, a, b))
	x2, y2, z2 := 56.060, 50.490, 8.160
	if !check(x, x2) {
		t.Error()
		fmt.Println(x, x2)
	}
	if !check(y, y2) {
		t.Error()
		fmt.Println(y, y2)
	}
	if !check(z, z2) {
		t.Error()
		fmt.Println(z, z2)
	}
}

func TestXYZtoLuv(t *testing.T) {
	fmt.Println("test of XYZ to Luv")
	x, y, z := 56.060, 50.490, 8.160
	l, u, v := (XYZToLuv(x, y, z))
	l2, u2, v2 := 76.369, 69.282, 73.457
	if !check(l, l2) {
		t.Error()
		fmt.Println(l, l2)
	}
	if !check(u, u2) {
		t.Error()
		fmt.Println(u, u2)
	}
	if !check(v, v2) {
		t.Error()
		fmt.Println(v, v2)
	}
}

func TestXYZtoLuvAndBack(t *testing.T) {
	fmt.Println("test of XYZ to Luv and back")
	x, y, z := 56.060, 50.490, 8.160
	l, u, v := XYZToLuv(x, y, z)
	x2, y2, z2 := LuvToXYZ(l, u, v)
	if !check(x, x2) {
		t.Error()
		fmt.Println(x, x2)
	}
	if !check(y, y2) {
		t.Error()
		fmt.Println(y, y2)
	}
	if !check(z, z2) {
		t.Error()
		fmt.Println(z, z2)
	}
}

func TestXYZtoLabAndBack(t *testing.T) {
	fmt.Println("test of XYZ to Lab and back")
	x, y, z := 56.060, 50.490, 8.160
	l, a, b := XYZToLab(x, y, z)
	x2, y2, z2 := LabToXYZ(l, a, b)
	if !check(x, x2) {
		t.Error()
		fmt.Println(x, x2)
	}
	if !check(y, y2) {
		t.Error()
		fmt.Println(y, y2)
	}
	if !check(z, z2) {
		t.Error()
		fmt.Println(z, z2)
	}
}

/*
RGB 0-255  =   254.98   171.01    32.04
RGB 0-FF   =       FE       AB       20
RGB 0-0.1  =  0.99994  0.67064  0.12566
CMY 0-0.1  =  0.00006  0.32936  0.87434
CMYK %     =    0.000   32.932   87.433    0.006
XYZ        =   56.060   50.490    8.160
Yxy        =   50.490  0.48871  0.44015
CIE-L*ab   =   76.369   21.174   74.935
CIE-L*CH   =   76.369   77.869   74.222
CIE-L*uv   =   76.369   69.282   73.457
HunterLab  =   71.056   16.479   42.931

Illuminant =  D65
Observer   =  2° (1931)

*/
