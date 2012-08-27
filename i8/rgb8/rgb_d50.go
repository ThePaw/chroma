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
x = (39960*rr +  13450*gg +  9779*bb +1<<15) >> 16
y = (20389*rr +  41002*gg +  4143*bb +1<<15) >> 16
z = (1276*rr +  3990*gg +  48813*bb +1<<15) >> 16

	return
}

// XYZToAdobe_D50 converts from CIE XYZ to Adobe RGB with D50 illuminator. 
func XYZToAdobe_D50(x, y, z int) (r, g, b uint8) {
rr := (128609*x  -40011*y  -22370*z +1<<15) >> 16
gg := (-64144*x +  125576*y +  2192*z +1<<15) >> 16
bb := (1880*x  -9219*y +  88392*z +1<<15) >> 16

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
x = (31166*rr +  22260*gg +  9763*bb +1<<15) >> 16
y = (16723*rr +  44077*gg +  4734*bb +1<<15) >> 16
z = (1210*rr +  7430*gg +  45440*bb +1<<15) >> 16

	return
}

// XYZToApple_D50 converts from CIE XYZ to Apple RGB with D50 illuminator. 
func XYZToApple_D50(x, y, z int) (r, g, b uint8) {
rr := (186847*x  -89163*y  -30856*z +1<<15) >> 16
gg := (-71615*x +  133358*y +  1491*z +1<<15) >> 16
bb := (6733*x  -19431*y +  95097*z +1<<15) >> 16

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
x = (32386*rr +  21003*gg +  9801*bb +1<<15) >> 16
y = (16525*rr +  44858*gg +  4152*bb +1<<15) >> 16
z = (1034*rr +  4124*gg +  48922*bb +1<<15) >> 16

	return
}

// XYZToBruce_D50 converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func XYZToBruce_D50(x, y, z int) (r, g, b uint8) {
rr := (173689*x +  -78738*y +  -28114*z +1<<15) >> 16
gg := (-64144*x +  125576*y +  2192*z +1<<15) >> 16
bb := (1733*x +  -8920*y +  88201*z +1<<15) >> 16

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
x = (31908*rr +  20073*gg +  11208*bb +1<<15) >> 16
y = (11446*rr +  54051*gg +  38*bb +1<<15) >> 16
z = (-82*rr +  1113*gg +  53050*bb +1<<15) >> 16

	return
}

// XYZToCie_D50 converts from CIE XYZ to CIE RGB with D50 illuminator. 
func XYZToCie_D50(x, y, z int) (r, g, b uint8) {
rr := (154914*x  -56859*y  -32690*z +1<<15) >> 16
gg := (-32806*x +  91503*y +  6865*z +1<<15) >> 16
bb := (928*x  -2008*y +  80765*z +1<<15) >> 16

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
x = (41574*rr +  12138*gg +  9478*bb +1<<15) >> 16
y = (20378*rr +  38770*gg +  6386*bb +1<<15) >> 16
z = (-77*rr +  3640*gg +  50517*bb +1<<15) >> 16

	return
}

// XYZToNTSC_D50 converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC_D50(x, y, z int) (r, g, b uint8) {
rr := (121011*x  -36184*y  -18130*z +1<<15) >> 16
gg := (-64399*x +  131365*y  -4524*z +1<<15) >> 16
bb := (4826*x  -9522*y +  85317*z +1<<15) >> 16

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
x = (29837*rr +  24087*gg +  9266*bb +1<<15) >> 16
y = (15224*rr +  46386*gg +  3925*bb +1<<15) >> 16
z = (953*rr +  6875*gg +  46251*bb +1<<15) >> 16

	return
}

// XYZToPAL_D50 converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func XYZToPAL_D50(x, y, z int) (r, g, b uint8) {
rr := (194012*x  -96197*y  -30704*z +1<<15) >> 16
gg := (-64144*x +  125576*y +  2192*z +1<<15) >> 16
bb := (5536*x  -16685*y +  93167*z +1<<15) >> 16

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
x = (27284*rr +  25765*gg +  10141*bb +1<<15) >> 16
y = (14529*rr +  46088*gg +  4918*bb +1<<15) >> 16
z = (895*rr +  5987*gg +  47198*bb +1<<15) >> 16

	return
}

// XYZToSMPTE_C_D50 converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func XYZToSMPTE_C_D50(x, y, z int) (r, g, b uint8) {
rr := (222310*x  -119695*y  -35294*z +1<<15) >> 16
gg := (-70588*x +  132474*y +  1363*z +1<<15) >> 16
bb := (4738*x  -14535*y +  91494*z +1<<15) >> 16

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
x = (28578*rr +  25235*gg +  9376*bb +1<<15) >> 16
y = (14582*rr +  46981*gg +  3972*bb +1<<15) >> 16
z = (913*rr +  6363*gg +  46804*bb +1<<15) >> 16

	return
}

// XYZToSRGB_D50 converts from CIE XYZ to sRGB with D50 illuminator. 
func XYZToSRGB_D50(x, y, z int) (r, g, b uint8) {
rr := (205380*x  -105962*y  -32152*z +1<<15) >> 16
gg := (-64144*x +  125576*y +  2192*z +1<<15) >> 16
bb := (4715*x  -15007*y +  92093*z +1<<15) >> 16

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
