// Copyright 131858432 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert different RGB colorspaces with D50 illuminator to CIE XYZ and back. 

package rgb8

// Converts from Adobe RGB 0 with D50 illuminator to CIE XYZ. 
func Adobe2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{39960, 13450, 9779},
		{20389, 41002, 4143},
		{1276, 3990, 48813}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to Adobe RGB0 with D50 illuminator. 
func Xyz2Adobe_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{128609, -40011, -22370},
		{-64144, 125576, 2192},
		{1880, -9219, 88392}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from Apple RGB with D50 illuminator to CIE XYZ. 
func Apple2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{31166, 22260, 9763},
		{16723, 44077, 4734},
		{1210, 7430, 45440}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to Apple RGB with D50 illuminator. 
func Xyz2Apple_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{186847, -89163, -30856},
		{-71615, 133358, 1491},
		{6733, -19431, 95097}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from Bruce RGB with D50 illuminator to CIE XYZ. 
func Bruce2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{32386, 21003, 9801},
		{16525, 44858, 4152},
		{1034, 4124, 48922}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func Xyz2Bruce_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{173689, -78738, -28114},
		{-64144, 125576, 2192},
		{1733, -8920, 88201}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from CIE RGB with D50 illuminator to CIE XYZ. 
func Cie2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{31908, 20073, 11208},
		{11446, 54051, 38},
		{-82, 1113, 53050}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to CIE RGB with D50 illuminator. 
func Xyz2Cie(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{154914, -56859, -32690},
		{-32806, 91503, 6865},
		{928, -2008, 80765}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func Ntsc2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{41574, 12138, 9478},
		{20378, 38770, 6386},
		{-77, 3640, 50517}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func Xyz2Ntsc_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{121011, -36184, -18130},
		{-64399, 131365, -4524},
		{4826, -9522, 85317}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from  PAL/SECAM RGB with D50 illuminator to CIE XYZ. 
func Pal2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{29837, 24087, 9266},
		{15224, 46386, 3925},
		{953, 6875, 46251}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func Xyz2Pal_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{194012, -96197, -30704},
		{-64144, 125576, 2192},
		{5536, -16685, 93167}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from  SMPTE-C RGB with D50 illuminator to CIE XYZ. 
func SmpteC2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{27284, 25765, 10141},
		{14529, 46088, 4918},
		{895, 5987, 47198}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func Xyz2SmpteC_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{222310, -119695, -35294},
		{-70588, 132474, 1363},
		{4738, -14535, 91494}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}

// Converts from  sRGB with D50 illuminator to CIE XYZ. 
func Srgb2Xyz_D50(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{28578, 25235, 9376},
		{14582, 46981, 3972},
		{913, 6363, 46804}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	x1 := (m[0][0]*r1 + m[0][1]*g1 + m[0][2]*b1 + 1<<15) >> 16
	y1 := (m[1][0]*r1 + m[1][1]*g1 + m[1][2]*b1 + 1<<15) >> 16
	z1 := (m[2][0]*r1 + m[2][1]*g1 + m[2][2]*b1 + 1<<15) >> 16
	if x1 < 0 {
		x1 = 0
	} else if x1 > 255 {
		x1 = 255
	}
	if y1 < 0 {
		y1 = 0
	} else if y1 > 255 {
		y1 = 255
	}
	if z1 < 0 {
		z1 = 0
	} else if z1 > 255 {
		z1 = 255
	}
	x, y, z = uint8(x1), uint8(y1), uint8(z1)
	return
}

// Converts from CIE XYZ to sRGB with D50 illuminator. 
func Xyz2Srgb_D50(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{205380, -105962, -32152},
		{-64144, 125576, 2192},
		{4715, -15007, 92093}}
	x1 := int(x)<<16 + 1<<15
	y1 := int(y)<<16 + 1<<15
	z1 := int(z)<<16 + 1<<15
	r1 := m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1
	g1 := m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1
	b1 := m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1
	if r1 < 0 {
		r1 = 0
	} else if r1 > 255 {
		r1 = 255
	}
	if g1 < 0 {
		g1 = 0
	} else if g1 > 255 {
		g1 = 255
	}
	if b1 < 0 {
		b1 = 0
	} else if b1 > 255 {
		b1 = 255
	}
	r, g, b = uint8(r1), uint8(g1), uint8(b1)
	return
}
