// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma
import (
	"math"
	. "image/color"
)


// RGBToXYZ converts an RGB triple to an XYZ triple.
func RGBToXYZ(r, g, b uint8) (float64, float64, float64) {
	aa := 0.04045
	c1 := 0.055
	c2 := 1.055
	c3 := 2.4
	c4 := 12.92
	rf := float64(r)/255.0
	gf := float64(g)/255.0
	bf := float64(b)/255.0
	if rf > aa {
		rf = math.Pow((rf+c1)/c2, c3)
	} else  {
		rf /= c4
	}

	if gf > aa {
		gf = math.Pow((gf+c1)/c2, c3)
	} else  {
		gf /= c4
	}

	if bf > aa {
		bf = math.Pow((bf+c1)/c2, c3)
	} else  {
		bf /= c4
	}

	rf *=100.0
	gf *=100.0
	bf *=100.0

	//Observer. = 2°, Illuminant = D65
	x := rf * 0.4124 + gf * 0.3576 + bf * 0.1805
	y := rf * 0.2126 + gf * 0.7152 + bf * 0.0722
	z := rf * 0.0193 + gf * 0.1192 + bf * 0.9505
	return x, y, z
}



// XYZToRGB converts an XYZ triple to an RGB triple.
func XYZToRGB(x, y, z float64) (uint8, uint8, uint8) {
//X from 0 to  95.047      (Observer = 2°, Illuminant = D65)
//Y from 0 to 100.000
//Z from 0 to 108.883
x /= 100.0
y /= 100.0
z /= 100.0

r := x *  3.2406 + y * -1.5372 + z * -0.4986
g := x * -0.9689 + y *  1.8758 + z *  0.0415
b := x *  0.0557 + y * -0.2040 + z *  1.0570

if r > 0.0031308 {
	r = 1.055 * ( math.Pow (r, 1 / 2.4 ) ) - 0.055
} else {
	r = 12.92 * r 
}

if g > 0.0031308 {
	g = 1.055 * ( math.Pow (g, 1 / 2.4 ) ) - 0.055
} else {
	g = 12.92 * g 
}

if b > 0.0031308 {
	b = 1.055 * ( math.Pow (b, 1 / 2.4 ) ) - 0.055
} else {
	b = 12.92 * b 
}


r *= 255.0
g *= 255.0
b *= 255.0

// clip
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}
	return uint8(r), uint8(g), uint8(b)
}

// XYZ represents a fully opaque 24-bit XYZ color, having 8 bits each .
type XYZ struct {
	X, Y, Z float64
}

func (c XYZ) RGBA() (uint32, uint32, uint32, uint32) {
	r, g, b := XYZToRGB(c.X, c.Y, c.Z)
	return uint32(r) * 0x101, uint32(g) * 0x101, uint32(b) * 0x101, 0xffff
}

// XYZModel is the Model for XYZ colors.
var XYZModel Model = ModelFunc(xYZModel)

func xYZModel(c Color) Color {
	if _, ok := c.(XYZ); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	x, y, z := RGBToXYZ(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	return XYZ{x, y, z}
}

