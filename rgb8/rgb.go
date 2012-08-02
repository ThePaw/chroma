// Copyright 131858432 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert different RGB colorspaces with their native illuminators (reference whites) to CIE XYZ scaled to 256 and back. 
// RGB values must be linear and in the nominal range [0, 255]. 
// XYZ values are usually in [0, 255] but may be greater
// To get quick and dirty XYZ approximations, divide by 255, otherwise use the float64 version of these functions. 
// Ref.: [24] 

package rgb8
//import "fmt"
// Converts from Adobe RGB 0 with D65 illuminator to CIE XYZ. 
func Adobe2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (37796*rr + 12160*gg + 12332*bb + 1<<15) >> 16 
	y = (19488*rr + 41113*gg + 4933*bb + 1<<15) >> 16
	z = (1771*rr + 4632*gg + 64953*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Adobe RGB0 with D65 illuminator. 
func Xyz2Adobe(x, y, z int) (r, g, b uint8) {
	rr := (133783*x  -37024*y  -22589*z + 1<<15) >> 16
	gg := (-63521*x + 122946*y + 2723*z + 1<<15) >> 16
	bb := (881*x  -7758*y + 66545*z + 1<<15) >> 16
	if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from Apple RGB with D65 illuminator to CIE XYZ. 
func Apple2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (29473*rr + 20725*gg + 12090*bb + 1<<15) >> 16
	y = (16033*rr + 44042*gg + 5460*bb + 1<<15) >> 16
	z = (1650*rr + 9252*gg + 60454*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Apple RGB with D65 illuminator. 
func Xyz2Apple(x, y, z int) (r, g, b uint8) {
	rr := (193431*x  -84502*y  -31053*z + 1<<15) >> 16
	gg := (-71113*x + 130472*y + 2438*z + 1<<15) >> 16
	bb := (5602*x  -17661*y + 71519*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from Best RGB with D50 illuminator to CIE XYZ. 
func Best2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (41462*rr + 13405*gg + 8322*bb + 1<<15) >> 16
	y = (14972*rr + 48323*gg + 2240*bb + 1<<15) >> 16
	z = (0*rr + 623*gg + 53457*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Best RGB with D50 illuminator. 
func Xyz2Best(x, y, z int) (r, g, b uint8) {
	rr := (115032*x  -31698*y  -16580*z + 1<<15) >> 16
	gg := (-35660*x + 98754*y + 1412*z + 1<<15) >> 16
	bb := (415*x  -1151*y + 80327*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from Beta RGB with D50 illuminator to CIE XYZ. 
func Beta2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (43991*rr + 11441*gg + 7758*bb + 1<<15) >> 16
	y = (19875*rr + 43501*gg + 2158*bb + 1<<15) >> 16
	z = (0*rr + 2667*gg + 51413*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Beta RGB with D50 illuminator. 
func Xyz2Beta(x, y, z int) (r, g, b uint8) {
	rr := (110311*x  -28064*y  -15467*z + 1<<15) >> 16
	gg := (-50529*x + 111840*y + 2928*z + 1<<15) >> 16
	bb := (2621*x  -5802*y + 83385*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from Bruce RGB with D65 illuminator to CIE XYZ. 
func Bruce2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (30632*rr + 19297*gg + 12360*bb + 1<<15) >> 16
	y = (15794*rr + 44796*gg + 4944*bb + 1<<15) >> 16
	z = (1435*rr + 4824*gg + 65097*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Bruce RGB with D65 illuminator. 
func Xyz2Bruce(x, y, z int) (r, g, b uint8) {
	rr := (179926*x  -74436*y  -28509*z + 1<<15) >> 16
	gg := (-63521*x + 122946*y + 2723*z + 1<<15) >> 16
	bb := (738*x  -7469*y + 66404*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from CIE RGB with E illuminator to CIE XYZ. 
func CieE2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (32028*rr + 20360*gg + 13146*bb + 1<<15) >> 16
	y = (11547*rr + 53279*gg + 708*bb + 1<<15) >> 16
	z = (0*rr + 668*gg + 64867*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to CIE RGB with E illuminator. 
func Xyz2CieE(x, y, z int) (r, g, b uint8) {
	rr := (155364*x  -58985*y  -30843*z + 1<<15) >> 16
	gg := (-33677*x + 93408*y + 5805*z + 1<<15) >> 16
	bb := (347*x  -963*y + 66151*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  ColorMatch RGB with D50 illuminator to CIE XYZ. 
func ColorMatch2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (33380*rr + 21030*gg + 8779*bb + 1<<15) >> 16
	y = (18014*rr + 43131*gg + 4389*bb + 1<<15) >> 16
	z = (1589*rr + 7129*gg + 45362*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to ColorMatch RGB with D50 illuminator. 
func Xyz2ColorMatch(x, y, z int) (r, g, b uint8) {
	rr := (173164*x  -80178*y  -25756*z + 1<<15) >> 16
	gg := (-72874*x + 134939*y + 1046*z + 1<<15) >> 16
	bb := (5385*x  -18397*y + 95419*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  Don RGB-4 with D50 illuminator to CIE XYZ. 
func Don2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (42321*rr + 12671*gg + 8198*bb + 1<<15) >> 16
	y = (18241*rr + 45086*gg + 2207*bb + 1<<15) >> 16
	z = (243*rr + 1178*gg + 52658*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Don RGB-4 with D50 illuminator. 
func Xyz2Don(x, y, z int) (r, g, b uint8) {
	rr := (115368*x  -31989*y  -16620*z + 1<<15) >> 16
	gg := (-46702*x + 108314*y + 2730*z + 1<<15) >> 16
	bb := (512*x  -2276*y + 81577*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  ECI RGB with D50 illuminator to CIE XYZ. 
func Eci2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (42611*rr + 11670*gg + 8908*bb + 1<<15) >> 16
	y = (20987*rr + 39457*gg + 5090*bb + 1<<15) >> 16
	z = (0*rr + 4445*gg + 49635*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to ECI RGB with D50 illuminator. 
func Xyz2Eci(x, y, z int) (r, g, b uint8) {
	rr := (116835*x  -32570*y  -17629*z + 1<<15) >> 16
	gg := (-62872*x + 127650*y  -1807*z + 1<<15) >> 16
	bb := (5631*x  -11433*y + 86692*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  Ekta Space PS5 with D50 illuminator to CIE XYZ. 
func EktaSpace2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (38921*rr + 17890*gg + 6379*bb + 1<<15) >> 16
	y = (17080*rr + 48165*gg + 289*bb + 1<<15) >> 16
	z = (0*rr + 2752*gg + 51328*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. 
func Xyz2EktaSpace(x, y, z int) (r, g, b uint8) {
	rr := (131359*x  -47873*y  -16056*z + 1<<15) >> 16
	gg := (-46597*x + 106182*y + 5191*z + 1<<15) >> 16
	bb := (2498*x  -5693*y + 83397*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func Ntsc2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (39773*rr + 11370*gg + 13130*bb + 1<<15) >> 16
	y = (19589*rr + 38443*gg + 7502*bb + 1<<15) >> 16
	z = (0*rr + 4331*gg + 73152*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func Xyz2Ntsc(x, y, z int) (r, g, b uint8) {
	rr := (125173*x  -34894*y  -18888*z + 1<<15) >> 16
	gg := (-64531*x + 131017*y  -1855*z + 1<<15) >> 16
	bb := (3821*x  -7758*y + 58822*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  PAL/SECAM RGB with D65 illuminator to CIE XYZ. 
func Pal2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (28221*rr + 22383*gg + 11685*bb + 1<<15) >> 16
	y = (14551*rr + 46310*gg + 4674*bb + 1<<15) >> 16
	z = (1322*rr + 8490*gg + 61544*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. 
func Xyz2Pal(x, y, z int) (r, g, b uint8) {
	rr := (200730*x +  -91303*y +  -31178*z +1<<15) >> 16
	gg := (-63521*x + 122946*y + 2723*z + 1<<15) >> 16
	bb := (4448*x  -14998*y + 70080*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  ProPhoto RGB with D50 illuminator to CIE XYZ. 
func ProPhoto2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (52276*rr + 8859*gg + 2054*bb + 1<<15) >> 16
	y = (18877*rr + 46653*gg + 5*bb + 1<<15) >> 16
	z = (0*rr + 0*gg + 54080*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. 
func Xyz2ProPhoto(x, y, z int) (r, g, b uint8) {
	rr := (88207*x  -16751*y  -3349*z + 1<<15) >> 16
	gg := (-35690*x + 98839*y + 1345*z + 1<<15) >> 16
	bb := (0*x + 0*y + 79417*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  SMPTE-C RGB with D65 illuminator to CIE XYZ. 
func SmpteC2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (25794*rr + 23937*gg + 12558*bb + 1<<15) >> 16
	y = (13920*rr + 45943*gg + 5671*bb + 1<<15) >> 16
	z = (1228*rr + 7335*gg + 62793*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. 
func Xyz2SmpteC(x, y, z int) (r, g, b uint8) {
	rr := (229729*x  -113999*y  -35649*z + 1<<15) >> 16
	gg := (-70062*x + 129618*y + 2305*z + 1<<15) >> 16
	bb := (3690*x  -12912*y + 68826*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  sRGB with D65 illuminator to CIE XYZ. 
func Srgb2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (27030*rr + 23434*gg + 11825*bb + 1<<15) >> 16
	y = (13937*rr + 46868*gg + 4730*bb + 1<<15) >> 16
	z = (1267*rr + 7811*gg + 62279*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to sRGB with D65 illuminator. 
func Xyz2Srgb(x, y, z int) (r, g, b uint8) {
	rr := (212366*x  -100737*y  -32671*z + 1<<15) >> 16
	gg := (-63521*x + 122946*y + 2723*z + 1<<15) >> 16
	bb := (3646*x  -13371*y + 69286*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}

// Converts from  Wide Gamut RGB with D50 illuminator to CIE XYZ. 
func WGamut2Xyz(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = (46930*rr + 6614*gg + 9645*bb + 1<<15) >> 16
	y = (16920*rr + 47509*gg + 1105*bb + 1<<15) >> 16
	z = (0*rr + 3393*gg + 50687*bb + 1<<15) >> 16
	return
}

// Converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. 
func Xyz2WGamut(x, y, z int) (r, g, b uint8) {
	rr := (95866*x  -12062*y  -17980*z + 1<<15) >> 16
	gg := (-34196*x + 94846*y + 4438*z + 1<<15) >> 16
	bb := (2289*x  -6349*y + 84437*z + 1<<15) >> 16
		if rr < 0 {
		rr = 0
	} else if rr > 255 {
		rr = 255
	}
	if gg < 0 {
		gg = 0
	} else if gg > 255 {
		gg = 255
	}
	if bb < 0 {
		bb = 0
	} else if bb > 255 {
		bb = 255
	}
	r, g, b = uint8(rr), uint8(gg), uint8(bb)
	return
}
