// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb8

// Conversion of different RGB colorspaces with their native illuminators (reference whites) to CIE XYZ scaled to 256 and back. 
// RGB values must be linear and in the nominal range [0, 255]. 
// XYZ values are usually in [0, 1e9] but may slightly outside. 
// To get quick and dirty XYZ approximations, divide by 1e9, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ converts from Adobe RGB (1998) with D65 illuminator to CIE XYZ. 
func AdobeToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2261689*rr + 727662*gg + 737981*bb
	y = 1166183*rr + 2460192*gg + 295192*bb
	z = 106016*rr + 277204*gg + 3886699*bb
	return
}

// XYZToAdobe converts from CIE XYZ to Adobe RGB (1998) with D65 illuminator. 
func XYZToAdobe(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (5205*x - 1440*y - 878*z) / 1e6
	gg := (-2471*x + 4783*y + 105*z) / 1e6
	bb := (34*x - 301*y + 2589*z) / 1e6
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

// AppleToXYZ converts from Apple RGB with D65 illuminator to CIE XYZ. 
func AppleToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1763642*rr + 1240190*gg + 723500*bb
	y = 959421*rr + 2635405*gg + 326741*bb
	z = 98763*rr + 553656*gg + 3617501*bb
	return
}

// XYZToApple converts from CIE XYZ to Apple RGB with D65 illuminator. 
func XYZToApple(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7526*x - 3287*y - 1208*z) / 1e6
	gg := (-2767*x + 5076*y + 94*z) / 1e6
	bb := (218*x - 687*y + 2782*z) / 1e6
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

// BestToXYZ converts from Best RGB with D50 illuminator to CIE XYZ. 
func BestToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2481057*rr + 802179*gg + 498018*bb
	y = 895909*rr + 2891577*gg + 134081*bb
	z = 37310*gg + 3198807*bb
	return
}

// XYZToBest converts from CIE XYZ to Best RGB with D50 illuminator. 
func XYZToBest(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4475*x - 1233*y - 645*z) / 1e6
	gg := (-1387*x + 3842*y + 54*z) / 1e6
	bb := (16*x - 44*y + 3125*z) / 1e6
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

// BetaToXYZ converts from Beta RGB with D50 illuminator to CIE XYZ. 
func BetaToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2632367*rr + 684640*gg + 464246*bb
	y = 1189304*rr + 2603082*gg + 129181*bb
	z = 159611*gg + 3076505*bb
	return
}

// XYZToBeta converts from CIE XYZ to Beta RGB with D50 illuminator. 
func XYZToBeta(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4292*x - 1092*y - 601*z) / 1e6
	gg := (-1966*x + 4351*y + 113*z) / 1e6
	bb := (102*x - 225*y + 3244*z) / 1e6
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

// BruceToXYZ converts from Bruce RGB with D65 illuminator to CIE XYZ. 
func BruceToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1833004*rr + 1154710*gg + 739618*bb
	y = 945143*rr + 2680578*gg + 295847*bb
	z = 85921*rr + 288677*gg + 3895321*bb
	return
}

// XYZToBruce converts from CIE XYZ to Bruce RGB with D65 illuminator. 
func XYZToBruce(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7000*x - 2896*y - 1109*z) / 1e6
	gg := (-2471*x + 4783*y + 105*z) / 1e6
	bb := (28*x - 290*y + 2583*z) / 1e6
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

// CIEToXYZ converts from CIE RGB with E illuminator to CIE XYZ. 
func CIEToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1916541*rr + 1218354*gg + 786673*bb
	y = 690997*rr + 3188175*gg + 42395*bb
	z = 40018*gg + 3881549*bb
	return
}

// XYZToCIE converts from CIE XYZ to CIE RGB with E illuminator. 
func XYZToCIE(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (6045*x - 2295*y - 1200*z) / 1e6
	gg := (-1310*x + 3634*y + 225*z) / 1e6
	bb := (13*x - 37*y + 2573*z) / 1e6
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

// ColorMatchToXYZ converts from  ColorMatch RGB with D50 illuminator to CIE XYZ. 
func ColorMatchToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1997427*rr + 1258459*gg + 525369*bb
	y = 1077976*rr + 2580907*gg + 262684*bb
	z = 95115*rr + 426596*gg + 2714405*bb
	return
}

// XYZToColorMatch converts from CIE XYZ to ColorMatch RGB with D50 illuminator. 
func XYZToColorMatch(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (6737*x - 3119*y - 1002*z) / 1e6
	gg := (-2835*x + 5250*y + 40*z) / 1e6
	bb := (209*x - 715*y + 3712*z) / 1e6
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

//DonToXYZ converts from  Don RGB-4 with D50 illuminator to CIE XYZ. 
func DonToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2532435*rr + 758239*gg + 490579*bb
	y = 1091567*rr + 2697922*gg + 132079*bb
	z = 14554*rr + 70533*gg + 3151029*bb
	return
}

// XYZToDon converts from CIE XYZ to Don RGB-4 with D50 illuminator. 
func XYZToDon(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4488*x - 1244*y - 646*z) / 1e6
	gg := (-1817*x + 4214*y + 106*z) / 1e6
	bb := (19*x - 88*y + 3174*z) / 1e6
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

// ECIToXYZ converts from  ECI RGB with D50 illuminator to CIE XYZ. 
func ECIToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2549820*rr + 698342*gg + 533091*bb
	y = 1255881*rr + 2361063*gg + 304623*bb
	z = 266035*gg + 2970082*bb
	return
}

// XYZToECI converts from CIE XYZ to ECI RGB with D50 illuminator. 
func XYZToECI(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4546*x - 1267*y - 685*z) / 1e6
	gg := (-2446*x + 4966*y - 70*z) / 1e6
	bb := (219*x - 444*y + 3373*z) / 1e6
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

// EktaSpaceToXYZ converts from  Ekta Space PS5 with D50 illuminator to CIE XYZ. 
func EktaSpaceToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2328985*rr + 1070510*gg + 381758*bb
	y = 1022072*rr + 2882143*gg + 17352*bb
	z = 164693*gg + 3071423*bb
	return
}

// XYZToEktaSpace converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. 
func XYZToEktaSpace(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (5111*x - 1862*y - 624*z) / 1e6
	gg := (-1813*x + 4131*y + 202*z) / 1e6
	bb := (97*x - 221*y + 3244*z) / 1e6
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

// NTSCToXYZ converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func NTSCToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2379964*rr + 680396*gg + 785678*bb
	y = 1172221*rr + 2300388*gg + 448958*bb
	z = 259198*gg + 4377350*bb
	return
}

// XYZToNTSC converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4870*x - 1357*y - 734*z) / 1e6
	gg := (-2510*x + 5097*y - 72*z) / 1e6
	bb := (148*x - 301*y + 2288*z) / 1e6
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

// PALToXYZ converts from  PAL/SECAM RGB with D65 illuminator to CIE XYZ. 
func PALToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1688701*rr + 1339380*gg + 699251*bb
	y = 870736*rr + 2771130*gg + 279700*bb
	z = 79158*rr + 508040*gg + 3682723*bb
	return
}

// XYZToPAL converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. 
func XYZToPAL(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7810*x - 3552*y - 1213*z) / 1e6
	gg := (-2471*x + 4783*y + 105*z) / 1e6
	bb := (173*x - 583*y + 2726*z) / 1e6
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

// ProPhotoToXYZ converts from  ProPhoto RGB with D50 illuminator to CIE XYZ. 
func ProPhotoToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 3128136*rr + 530163*gg + 122954*bb
	y = 1129569*rr + 2791663*gg + 336*bb
	z = 3236117 * bb
	return
}

// XYZToProPhoto converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. 
func XYZToProPhoto(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (3432*x - 651*y - 130*z) / 1e6
	gg := (-1388*x + 3845*y + 52*z) / 1e6
	bb := (0*x + 0*y + 3090*z) / 1e6
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

// SMPTE_CToXYZ converts from  SMPTE-C RGB with D65 illuminator to CIE XYZ. 
func SMPTE_CToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1543486*rr + 1432351*gg + 751495*bb
	y = 832992*rr + 2749190*gg + 339385*bb
	z = 73499*rr + 438946*gg + 3757475*bb
	return
}

// XYZToSMPTE_C converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. 
func XYZToSMPTE_C(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (8938*x - 4435*y - 1387*z) / 1e6
	gg := (-2726*x + 5043*y + 89*z) / 1e6
	bb := (143*x - 502*y + 2678*z) / 1e6
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

// SRGBToXYZ converts from  sRGB with D65 illuminator to CIE XYZ. 
func SRGBToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1617476*rr + 1402259*gg + 707598*bb
	y = 834011*rr + 2804518*gg + 283039*bb
	z = 75819*rr + 467419*gg + 3726682*bb
	return
}

// XYZToSRGB converts from CIE XYZ to sRGB with D65 illuminator. 
func XYZToSRGB(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (8263*x - 3919*y - 1271*z) / 1e6
	gg := (-2471*x + 4783*y + 105*z) / 1e6
	bb := (141*x - 520*y + 2695*z) / 1e6
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

// WGamutToXYZ converts from  Wide Gamut RGB with D50 illuminator to CIE XYZ. 
func WGamutToXYZ(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2808253*rr + 395802*gg + 577199*bb
	y = 1012499*rr + 2842893*gg + 66175*bb
	z = 203063*gg + 3033053*bb
	return
}

// XYZToWGamut converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. 
func XYZToWGamut(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (3730*x - 469*y - 699*z) / 1e6
	gg := (-1330*x + 3690*y + 172*z) / 1e6
	bb := (89*x - 247*y + 3285*z) / 1e6
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
