// Copyright 131858432 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert different RGB colorspaces with their native illuminators to CIE XYZ and back. 
// RGB values must be linear and in the nominal range [0, 255]. 
// Ref.: [24] 

package rgb8

// Converts from Adobe RGB with D65 illuminator to CIE XYZ. 
func Adobe2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{37796, 12160, 12332},
		{19488, 41113, 4933},
		{1771, 4632, 64953}}
	r1 := int(r)
	g1 := int(g)
	b1 := int(b)
	/*
		x1 := (37796*r1 + 12160*g1 + 12332*b1 + 1<<15) >> 16
		y1 := (19488*r1 + 41113*g1 + 4933*b1 + 1<<15) >> 16
		z1 := (1771*r1 + 4632*g1 + 64953*b1 + 1<<15) >> 16
	*/
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

// Converts from CIE XYZ to Adobe RGB0 with D65 illuminator. 
func Xyz2Adobe(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{133783, -37024, -22589},
		{-63521, 122946, 2723},
		{881, -7758, 66545}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from Apple RGB with D65 illuminator to CIE XYZ. 
func Apple2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{29473, 20725, 12090},
		{16033, 44042, 5460},
		{1650, 9252, 60454}}
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

// Converts from CIE XYZ to Apple RGB with D65 illuminator. 
func Xyz2Apple(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{193431, -84502, -31053},
		{-71113, 130472, 2438},
		{5602, -17661, 71519}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from Best RGB with D50 illuminator to CIE XYZ. 
func Best2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{41462, 13405, 8322},
		{14972, 48323, 2240},
		{0, 623, 53457}}
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

// Converts from CIE XYZ to Best RGB with D50 illuminator. 
func Xyz2Best(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{115032, -31698, -16580},
		{-35660, 98754, 1412},
		{415, -1151, 80327}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from Beta RGB with D50 illuminator to CIE XYZ. 
func Beta2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{43991, 11441, 7758},
		{19875, 43501, 2158},
		{0, 2667, 51413}}
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

// Converts from CIE XYZ to Beta RGB with D50 illuminator. 
func Xyz2Beta(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{110311, -28064, -15467},
		{-50529, 111840, 2928},
		{2621, -5802, 83385}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from Bruce RGB with D65 illuminator to CIE XYZ. 
func Bruce2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{30632, 19297, 12360},
		{15794, 44796, 4944},
		{1435, 4824, 65097}}
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

// Converts from CIE XYZ to Bruce RGB with D65 illuminator. 
func Xyz2Bruce(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{179926, -74436, -28509},
		{-63521, 122946, 2723},
		{738, -7469, 66404}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from CIE RGB with E illuminator to CIE XYZ. 
func CieE2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{32028, 20360, 13146},
		{11547, 53279, 708},
		{0, 668, 64867}}
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

// Converts from CIE XYZ to CIE RGB with E illuminator. 
func Xyz2CieE(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{155364, -58985, -30843},
		{-33677, 93408, 5805},
		{347, -963, 66151}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  ColorMatch RGB with D50 illuminator to CIE XYZ. 
func ColorMatch2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{33380, 21030, 8779},
		{18014, 43131, 4389},
		{1589, 7129, 45362}}
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

// Converts from CIE XYZ to ColorMatch RGB with D50 illuminator. 
func Xyz2ColorMatch(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{173164, -80178, -25756},
		{-72874, 134939, 1046},
		{5385, -18397, 95419}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  Don RGB-4 with D50 illuminator to CIE XYZ. 
func Don2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{42321, 12671, 8198},
		{18241, 45086, 2207},
		{243, 1178, 52658}}
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

// Converts from CIE XYZ to Don RGB-4 with D50 illuminator. 
func Xyz2Don(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{115368, -31989, -16620},
		{-46702, 108314, 2730},
		{512, -2276, 81577}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  ECI RGB with D50 illuminator to CIE XYZ. 
func Eci2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{42611, 11670, 8908},
		{20987, 39457, 5090},
		{0, 4445, 49635}}
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

// Converts from CIE XYZ to ECI RGB with D50 illuminator. 
func Xyz2Eci(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{116835, -32570, -17629},
		{-62872, 127650, -1807},
		{5631, -11433, 86692}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  Ekta Space PS5 with D50 illuminator to CIE XYZ. 
func EktaSpace2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{38921, 17890, 6379},
		{17080, 48165, 289},
		{0, 2752, 51328}}
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

// Converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. 
func Xyz2EktaSpace(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{131359, -47873, -16056},
		{-46597, 106182, 5191},
		{2498, -5693, 83397}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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
func Ntsc2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{39773, 11370, 13130},
		{19589, 38443, 7502},
		{0, 4331, 73152}}
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
func Xyz2Ntsc(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{125173, -34894, -18888},
		{-64531, 131017, -1855},
		{3821, -7758, 58822}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  PAL/SECAM RGB with D65 illuminator to CIE XYZ. 
func Pal2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{28221, 22383, 11685},
		{14551, 46310, 4674},
		{1322, 8490, 61544}}
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

// Converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. 
func Xyz2Pal(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{200730, -91303, -31178},
		{-63521, 122946, 2723},
		{4448, -14998, 70080}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  ProPhoto RGB with D50 illuminator to CIE XYZ. 
func ProPhoto2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{52276, 8859, 2054},
		{18877, 46653, 5},
		{0, 0, 54080}}
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

// Converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. 
func Xyz2ProPhoto(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{88207, -16751, -3349},
		{-35690, 98839, 1345},
		{0, 0, 79417}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  SMPTE-C RGB with D65 illuminator to CIE XYZ. 
func SmpteC2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{25794, 23937, 12558},
		{13920, 45943, 5671},
		{1228, 7335, 62793}}
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

// Converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. 
func Xyz2SmpteC(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{229729, -113999, -35649},
		{-70062, 129618, 2305},
		{3690, -12912, 68826}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  sRGB with D65 illuminator to CIE XYZ. 
func Srgb2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{27030, 23434, 11825},
		{13937, 46868, 4730},
		{1267, 7811, 62279}}
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

// Converts from CIE XYZ to sRGB with D65 illuminator. 
func Xyz2Srgb(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{212366, -100737, -32671},
		{-63521, 122946, 2723},
		{3646, -13371, 69286}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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

// Converts from  Wide Gamut RGB with D50 illuminator to CIE XYZ. 
func WGamut2Xyz(r, g, b uint8) (x, y, z uint8) {
	m := [3][3]int{
		{46930, 6614, 9645},
		{16920, 47509, 1105},
		{0, 3393, 50687}}
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

// Converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. 
func Xyz2WGamut(x, y, z uint8) (r, g, b uint8) {
	m := [3][3]int{
		{95866, -12062, -17980},
		{-34196, 94846, 4438},
		{2289, -6349, 84437}}
	x1 := int(x)
	y1 := int(y)
	z1 := int(z)
	r1 := (m[0][0]*x1 + m[0][1]*y1 + m[0][2]*z1 + 1<<15) >> 16
	g1 := (m[1][0]*x1 + m[1][1]*y1 + m[1][2]*z1 + 1<<15) >> 16
	b1 := (m[2][0]*x1 + m[2][1]*y1 + m[2][2]*z1 + 1<<15) >> 16
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
