// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// Rgb2Xyz converts an RGB triple to an XYZ triple.
func Rgb2Xyz(r, g, b float64) (x, y, z float64) {
	r = gammaCorrInv(r)
	g = gammaCorrInv(g)
	b = gammaCorrInv(b)
	x = 0.4123955889674142161*r + 0.3575834307637148171*g + 0.1804926473817015735*b
	y = 0.2125862307855955516*r + 0.7151703037034108499*g + 0.07220049864333622685*b
	z = 0.01929721549174694484*r + 0.1191838645808485318*g + 0.9504971251315797660*b
	return
}

// Xyz2Rgb converts an XYZ triple to an RGB triple.
func Xyz2Rgb(x, y, z float64) (r, g, b float64) {
	// convert to linear RGB
	r1 := 3.2406*x - 1.5372*y - 0.4986*z
	g1 := -0.9689*x + 1.8758*y + 0.0415*z
	b1 := 0.0557*x - 0.2040*y + 1.0570*z

	min := min3(r1, g1, b1)

	// force nonnegative values so that gamma correction is well-defined
	if min < 0 {
		r1 -= min
		g1 -= min
		b1 -= min
	}

	// convert from RGB to R'G'B'
	r = gammaCorr(r1)
	g = gammaCorr(g1)
	b = gammaCorr(b1)
	return
}

// XYZ represents a fully opaque 3x64-bit XYZ color.
type XYZ struct {
	X, Y, Z float64
}

/*
func (c XYZ) RGBA() (uint32, uint32, uint32, uint32) {
	r, g, b := xyz2Rgb(c.X, c.Y, c.Z)
	return uint32(r) * 0x101, uint32(g) * 0x101, uint32(b) * 0x101, 0xffff
}

// XYZModel is the Model for XYZ colors.
var XYZModel Model = ModelFunc(xYZModel)

func xYZModel(c Color) Color {
	if _, ok := c.(XYZ) ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	x, y, z := Rgb2Xyz(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	return XYZ{x, y, z}
}
*/
