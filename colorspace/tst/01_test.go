// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

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
	fmt.Println("test of sRGB to XYZ")
	r, g, b := 0.545877, 0.966567, 0.463759
	x, y, z := (Rgb2Xyz(r, g, b))
	x2, y2, z2 := 0.470645, 0.730177, 0.288323
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
	fmt.Println("test of XYZ to sRGB")
	x, y, z := 0.470645, 0.730177, 0.288323
	r, g, b := Xyz2Rgb(x, y, z)
	r2, g2, b2 := 0.545877, 0.966567, 0.463759
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
	x, y, z := 0.470645, 0.730177, 0.288323
	l, a, b := (Xyz2Lab(x, y, z))
	l2, a2, b2 := 88.456154, -54.671483, 51.662818
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
		fmt.Println("b: ", b, b2)
	}
}

func TestLab2Xyz(t *testing.T) {
	fmt.Println("test of Lab to XYZ")
	l, a, b := 88.456154, -54.671483, 51.662818
	x, y, z := (Lab2Xyz(l, a, b))
	x2, y2, z2 := 0.470645, 0.730177, 0.288323
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
	x, y, z := 0.470645, 0.730177, 0.288323
	l, u, v := (Xyz2Luv(x, y, z))
	l2, u2, v2 := 88.456154, -51.330414, 76.405526
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
	l, u, v := Xyz2Luv(x, y, z)
	x2, y2, z2 := Luv2Xyz(l, u, v)
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
	l, a, b := Xyz2Lab(x, y, z)
	x2, y2, z2 := Lab2Xyz(l, a, b)
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
// Tests from sRGB, from the table below
/*
sRGB         0.545877, 0.966567, 0.463759
Y'CbCr       187.577791, 87.586330, 90.040886
JPEG-Y'CbCr  0.783460, 0.319581, 0.330539
Y'PbPr       0.783460, -0.180419, -0.169461
Y'DbDr       0.783460, -0.480932, 0.451670
Y'UV         0.783460, -0.157383, -0.208443
Y'IQ         0.783460, -0.089078, -0.245399
HSV          110.200859, 0.520200, 0.966567
HSL          110.200859, 0.882623, 0.715163
HSI          111.244375, 0.295985, 0.658734
XYZ          0.470645, 0.730177, 0.288323
L*a*b*        88.456154, -54.671483,  51.662818
L*u*v*        88.456154, -51.330414, 76.405526
L*C*H*        88.456154, 75.219797, 136.620717
CAT02 LMS     0.611749, 0.910088, 0.294880
*/
func TestRGBtoYCbCr(t *testing.T) {
	fmt.Println("test of sRGB to Y'CbCr and back")
	r, g, b := 0.540000, 0.960000, 0.460000
	d, e, f := 186.254980, 87.874740, 90.067000

	fn1 := Rgb2Ycbcr
	fn2 := Ycbcr2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d2, e2, f2)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoJPEGYCbCr(t *testing.T) {
	fmt.Println("test of sRGB to JPEG-Y'CbCr and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.783460, 0.319581, 0.330539
	d2, e2, f2 := Rgb2Jycbcr(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := Jycbcr2Rgb(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoYPbPr(t *testing.T) {
	fmt.Println("test of sRGB to Y'PbPr and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.783460, -0.180419, -0.169461
	fn1 := Rgb2Ypbpr
	fn2 := Ypbpr2Rgb
	d2, e2, f2 := fn1(r, g, b)

	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoYDbDr(t *testing.T) {
	fmt.Println("test of sRGB to Y'DbDr and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.783460, -0.480932, 0.451670
	fn1 := Rgb2Ydbdr
	fn2 := Ydbdr2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoYUV(t *testing.T) {
	fmt.Println("test of sRGB to Y'UV and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.783460, -0.157383, -0.208443
	fn1 := Rgb2Yuv
	fn2 := Yuv2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoYIQ(t *testing.T) {
	fmt.Println("test of sRGB to Y'IQ and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.783460, -0.089078, -0.245399
	fn1 := Rgb2Yiq
	fn2 := Yiq2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoHSV(t *testing.T) {
	fmt.Println("test of sRGB to HSV and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 110.200859, 0.520200, 0.966567
	fn1 := Rgb2Hsv
	fn2 := Hsv2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoHSL(t *testing.T) {
	fmt.Println("test of sRGB to HSL and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 110.200859, 0.882623, 0.715163
	fn1 := Rgb2Hsl
	fn2 := Hsl2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoHSI(t *testing.T) {
	fmt.Println("test of sRGB to HSI and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 111.244375, 0.295985, 0.658734
	fn1 := Rgb2Hsi
	fn2 := Hsi2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoXYZ(t *testing.T) {
	fmt.Println("test of sRGB to XYZ and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.470645, 0.730177, 0.288323
	fn1 := Rgb2Xyz
	fn2 := Xyz2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoLab(t *testing.T) {
	fmt.Println("test of sRGB to L*a*b* and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 88.456154, -54.671483, 51.662818
	fn1 := Rgb2Lab
	fn2 := Lab2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoLuv(t *testing.T) {
	fmt.Println("test of sRGB to L*u*v* and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 88.456154, -51.330414, 76.405526
	fn1 := Rgb2Luv
	fn2 := Luv2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoLch(t *testing.T) {
	fmt.Println("test of sRGB to L*C*H* and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 88.456154, 75.219797, 136.620717
	fn1 := Rgb2Lch
	fn2 := Lch2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}

func TestRGBtoLms(t *testing.T) {
	fmt.Println("test of sRGB to CAT02 LMS and back")
	r, g, b := 0.545877, 0.966567, 0.463759
	d, e, f := 0.611749, 0.910088, 0.294880
	fn1 := Rgb2Lms
	fn2 := Lms2Rgb
	//forward
	d2, e2, f2 := fn1(r, g, b)
	if !check(d, d2) {
		t.Error()
		fmt.Println("d: ", d, d2)
	}
	if !check(e, e2) {
		t.Error()
		fmt.Println("e: ", e, e2)
	}
	if !check(f, f2) {
		t.Error()
		fmt.Println("f: ", f, f2)
	}
	//back
	r2, g2, b2 := fn2(d, e, f)
	if !check(r, r2) {
		t.Error()
		fmt.Println("r: ", r, r2)
	}
	if !check(g, g2) {
		t.Error()
		fmt.Println("g: ", g, g2)
	}
	if !check(b, b2) {
		t.Error()
		fmt.Println("b: ", b, b2)
	}
}
