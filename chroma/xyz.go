// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import (
	. "image/color"
)

// SRGBToXYZ converts an sRGB triple to an XYZ triple.
func SRGBToXYZ(r, g, b uint8) (uint8, uint8, uint8) {
	ri := int(r)
	gi := int(g)
	bi := int(b)

	c := 65536
	m := sRGB_xyz

	r1 := c * int(m[0][0]) * ri
	g1 := c * int(m[0][1]) * gi
	b1 := c * int(m[0][2]) * bi
	r2 := c * int(m[1][0]) * ri
	g2 := c * int(m[1][1]) * gi
	b2 := c * int(m[1][2]) * bi
	r3 := c * int(m[2][0]) * ri
	g3 := c * int(m[2][1]) * gi
	b3 := c * int(m[2][2]) * bi

	x := (r1 + g1 + b1 + 1<<15) >> 16
	y := (r2 + g2 + b2 + 1<<15) >> 16
	z := (r3 + g3 + b3 + 1<<15) >> 16

	clip3(&x, &y, &z)
	return uint8(x), uint8(y), uint8(z)
}

// XYZToSRGB converts an XYZ triple to an sRGB triple.
func XYZToSRGB(x, y, z uint8) (uint8, uint8, uint8) {
	xi := int(x)
	yi := int(y)
	zi := int(z)

	c := 65536
	m := xyz_sRGB

	/*
		r := (xyz_sRGB[0][0]*xf + xyz_sRGB[0][1]*yf + xyz_sRGB[0][2]*zf)
		g := (xyz_sRGB[1][0]*xf + xyz_sRGB[1][1]*yf + xyz_sRGB[1][2]*zf)
		b := (xyz_sRGB[2][0]*xf + xyz_sRGB[2][1]*yf + xyz_sRGB[2][2]*zf)
	*/
	x1 := c * int(m[0][0]) * xi
	y1 := c * int(m[0][1]) * yi
	z1 := c * int(m[0][2]) * zi
	x2 := c * int(m[1][0]) * xi
	y2 := c * int(m[1][1]) * yi
	z2 := c * int(m[1][2]) * zi
	x3 := c * int(m[2][0]) * xi
	y3 := c * int(m[2][1]) * yi
	z3 := c * int(m[2][2]) * zi
	r := (x1 + y1 + z1 + 1<<15) >> 16
	g := (x2 + y2 + z2 + 1<<15) >> 16
	b := (x3 + y3 + z3 + 1<<15) >> 16

	clip3(&r, &g, &b)
	return uint8(r), uint8(g), uint8(b)
}

// XYZ represents a fully opaque 24-bit XYZ color, having 8 bits each .
type XYZ struct {
	X, Y, Z uint8
}

func (c XYZ) RGBA() (uint32, uint32, uint32, uint32) {
	r, g, b := XYZToSRGB(c.X, c.Y, c.Z)
	return uint32(r) * 0x101, uint32(g) * 0x101, uint32(b) * 0x101, 0xffff
}

// XYZModel is the Model for XYZ colors.
var XYZModel Model = ModelFunc(xYZModel)

func xYZModel(c Color) Color {
	if _, ok := c.(XYZ); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	x, y, z := SRGBToXYZ(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	return XYZ{x, y, z}
}
