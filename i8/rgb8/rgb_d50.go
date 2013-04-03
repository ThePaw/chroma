// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb8

// Conversion of natively non-D50 RGB colorspaces with D50 illuminator to CIE XYZ and back. 
// Bradford adaptation was used to calculate D50 matrices from colorspaces' native illuminators. 
// RGB values must be linear and in the nominal range [0, 255]. 
// XYZ values are usually in [0, 255] but may be greater
// To get quick and dirty XYZ approximations, divide by 255, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ_D50 converts from Adobe RGB 0 with D50 illuminator to CIE XYZ. 
func AdobeToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2391199*rr + 804863*gg + 585192*bb
	y = 1220094*rr + 2453552*gg + 247920*bb
	z = 76396*rr + 238785*gg + 2920936*bb
	return
}

// XYZToAdobe_D50 converts from CIE XYZ to Adobe RGB with D50 illuminator. 
func XYZToAdobe_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (5004*x - 1556*y - 870*z) / 1e6
	gg := (-2495*x + 4886*y + 85*z) / 1e6
	bb := (73*x - 358*y + 3439*z) / 1e6

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

// AppleToXYZ_D50 converts from Apple RGB with D50 illuminator to CIE XYZ. 
func AppleToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1864971*rr + 1332047*gg + 584235*bb
	y = 1000710*rr + 2637526*gg + 283331*bb
	z = 72430*rr + 444616*gg + 2719071*bb
	return
}

// XYZToApple_D50 converts from CIE XYZ to Apple RGB with D50 illuminator. 
func XYZToApple_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7270*x - 3469*y - 1200*z) / 1e6
	gg := (-2786*x + 5188*y + 58*z) / 1e6
	bb := (261*x - 756*y + 3700*z) / 1e6

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

// BruceToXYZ_D50 converts from Bruce RGB with D50 illuminator to CIE XYZ. 
func BruceToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1937967*rr + 1256797*gg + 586490*bb
	y = 988835*rr + 2684262*gg + 248470*bb
	z = 61916*rr + 246785*gg + 2927415*bb

	return
}

// XYZToBruce_D50 converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func XYZToBruce_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (6758*x - 3063*y - 1093*z) / 1e6
	gg := (-2495*x + 4886*y + 85*z) / 1e6
	bb := (67*x - 347*y + 3431*z) / 1e6

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

// CieToXYZ_D50 converts from CIE RGB with D50 illuminator to CIE XYZ. 
func CieToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1909360*rr + 1201170*gg + 670724*bb
	y = 684934*rr + 3234329*gg + 2304*bb
	z = -4926*rr + 66600*gg + 3174443*bb

	return
}

// XYZToCie_D50 converts from CIE XYZ to CIE RGB with D50 illuminator. 
func XYZToCie_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (6027*x - 2212*y - 1271*z) / 1e6
	gg := (-1276*x + 3560*y + 267*z) / 1e6
	bb := (36*x - 78*y + 3142*z) / 1e6

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

// NTSCToXYZ_D50 converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func NTSCToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 2487727*rr + 726354*gg + 567172*bb
	y = 1219410*rr + 2319993*gg + 382164*bb
	z = -4634*rr + 217850*gg + 3022901*bb
	return
}

// XYZToNTSC_D50 converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (4708*x - 1407*y - 705*z) / 1e6
	gg := (-2505*x + 5111*y - 176*z) / 1e6
	bb := (187*x - 370*y + 3319*z) / 1e6

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

// PALToXYZ_D50 converts from  PAL/SECAM RGB with D50 illuminator to CIE XYZ. 
func PALToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1785401*rr + 1441372*gg + 554480*bb
	y = 910990*rr + 2775669*gg + 234909*bb
	z = 57041*rr + 411432*gg + 2767642*bb
	return
}

// XYZToPAL_D50 converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func XYZToPAL_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7549*x - 3743*y - 1194*z) / 1e6
	gg := (-2495*x + 4886*y + 85*z) / 1e6
	bb := (215*x - 649*y + 3625*z) / 1e6

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

// SMPTE_CToXYZ_D50 converts from  SMPTE-C RGB with D50 illuminator to CIE XYZ. 
func SMPTE_CToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1632662*rr + 1541750*gg + 606841*bb
	y = 869411*rr + 2757862*gg + 294294*bb
	z = 53559*rr + 358276*gg + 2824282*bb
	return
}

// XYZToSMPTE_C_D50 converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func XYZToSMPTE_C_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (8650*x - 4657*y - 1373*z) / 1e6
	gg := (-2746*x + 5154*y + 53*z) / 1e6
	bb := (184*x - 565*y + 3560*z) / 1e6

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

// SRGBToXYZ_D50 converts from  sRGB with D50 illuminator to CIE XYZ. 
func SRGBToXYZ_D50(r, g, b uint8) (x, y, z int) {
	rr, gg, bb := int(r), int(g), int(b)
	x = 1710096*rr + 1510058*gg + 561099*bb
	y = 872566*rr + 2811288*gg + 237713*bb
	z = 54636*rr + 380801*gg + 2800679*bb
	return
}

// XYZToSRGB_D50 converts from CIE XYZ to sRGB with D50 illuminator. 
func XYZToSRGB_D50(x, y, z int) (r, g, b uint8) {
	x /= 1e4
	y /= 1e4
	z /= 1e4
	rr := (7991*x - 4123*y - 1251*z) / 1e6
	gg := (-2495*x + 4886*y + 85*z) / 1e6
	bb := (183*x - 583*y + 3583*z) / 1e6
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
