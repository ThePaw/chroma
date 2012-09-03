// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

// Conversion of natively non-D50 RGB colorspaces with D50 illuminator to CIE XYZ and back. 
// Bradford adaptation was used to calculate D50 matrices from colorspaces' native illuminators. 
// RGB values must be linear and in the nominal range [0, 255]. 
// XYZ values are usually in [0, 255] but may be greater
// To get quick and dirty XYZ approximations, divide by 255, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ_D50 converts from Adobe RGB 1998 with D50 illuminator to CIE XYZ.
func AdobeToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((93042786297*rr + 31317631799*gg + 22770122834*bb) / 1e7)
	y = int((47474509803*rr + 95468986037*gg + 9646707865*bb) / 1e7)
	z = int((2972625314*rr + 9291248950*gg + 113655100327*bb) / 1e7)
	return
}

// XYZToAdobe_D50 converts from CIE XYZ to Adobe RGB 1998 with D50 illuminator.
func XYZToAdobe_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((19624573*xx + -6105343*yy + -3413404*zz) / 152587890625)
	gg := int((-9787684*xx + 19161707*yy + 334545*zz) / 152587890625)
	bb := int((286873*xx + -1406752*yy + 13487860*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// AppleToXYZ_D50 converts from Apple RGB with D50 illuminator to CIE XYZ. 
func AppleToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((72566994735*rr + 51830655374*gg + 22732890821*bb) / 1e7)
	y = int((38938155184*rr + 102627496757*gg + 11024582283*bb) / 1e7)
	z = int((2818295566*rr + 17300236514*gg + 105800442511*bb) / 1e7)
	return
}

// XYZToApple_D50 converts from CIE XYZ to Apple RGB with D50 illuminator. 
func XYZToApple_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((28511130*xx + -13605261*yy + -4708281*zz) / 152587890625)
	gg := int((-10927680*xx + 20349181*yy + 227601*zz) / 152587890625)
	bb := int((1027418*xx + -2964984*yy + 14510880*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// BruceToXYZ_D50 converts from Bruce RGB with D50 illuminator to CIE XYZ. 
func BruceToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((75407278552*rr + 48902632180*gg + 22820630197*bb) / 1e7)
	y = int((38476096741*rr + 104446005950*gg + 9668116273*bb) / 1e7)
	z = int((2409185930*rr + 9602563515*gg + 113907209887*bb) / 1e7)
	return
}

// XYZToBruce_D50 converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func XYZToBruce_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((26503260*xx + -12014485*yy + -4289936*zz) / 152587890625)
	gg := int((-9787684*xx + 19161707*yy + 334545*zz) / 152587890625)
	bb := int((264574*xx + -1361228*yy + 13458747*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// CIEToXYZ_D50 converts from CIE RGB with D50 illuminator to CIE XYZ. 
func CIEToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((74294193941*rr + 46738139924*gg + 26098222323*bb) / 1e7)
	y = int((26651148241*rr + 125849408712*gg + 89677271*bb) / 1e7)
	z = int((-191696167*rr + 2591470206*gg + 123519203478*bb) / 1e7)
	return
}

// XYZToCIE_D50 converts from CIE XYZ to CIE RGB with D50 illuminator. 
func XYZToCIE_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((23638441*xx + -8676030*yy + -4988161*zz) / 152587890625)
	gg := int((-5005940*xx + 13962581*yy + 1047576*zz) / 152587890625)
	bb := int((141714*xx + -306400*yy + 12324030*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// NTSCToXYZ_D50 converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func NTSCToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((96798748759*rr + 28262821392*gg + 22068970778*bb) / 1e7)
	y = int((47447867551*rr + 90272129396*gg + 14870222018*bb) / 1e7)
	z = int((-180313111*rr + 8476661325*gg + 117622629128*bb) / 1e7)
	return
}

// XYZToNTSC_D50 converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((18465162*xx + -5521299*yy + -2766458*zz) / 152587890625)
	gg := int((-9826630*xx + 20045060*yy + -690397*zz) / 152587890625)
	bb := int((736488*xx + -1453020*yy + 13018574*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// PALToXYZ_D50 converts from  PAL/SECAM RGB with D50 illuminator to CIE XYZ. 
func PALToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((69470862897*rr + 56084534981*gg + 21575127793*bb) / 1e7)
	y = int((35447089340*rr + 108002685587*gg + 9140444037*bb) / 1e7)
	z = int((2219531547*rr + 16009063858*gg + 107690379185*bb) / 1e7)
	return
}

// XYZToPAL_D50 converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func XYZToPAL_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((29604395*xx + -14678520*yy + -4685105*zz) / 152587890625)
	gg := int((-9787684*xx + 19161707*yy + 334545*zz) / 152587890625)
	bb := int((844886*xx + -2545974*yy + 14216390*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// SMPTE_CToXYZ_D50 converts from  SMPTE-C RGB with D50 illuminator to CIE XYZ. 
func SMPTE_CToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((63527733272*rr + 59990295261*gg + 23612512397*bb) / 1e7)
	y = int((33829236285*rr + 107309819179*gg + 11451163499*bb) / 1e7)
	z = int((2084016174*rr + 13940703440*gg + 109894254978*bb) / 1e7)
	return
}

// XYZToSMPTE_C_D50 converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func XYZToSMPTE_C_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((33922457*xx + -18264027*yy + -5385522*zz) / 152587890625)
	gg := int((-10770996*xx + 20214283*yy + 207992*zz) / 152587890625)
	bb := int((723084*xx + -2217902*yy + 13961145*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}

// SRGBToXYZ_D50 converts from  sRGB with D50 illuminator to CIE XYZ. 
func SRGBToXYZ_D50(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((66540733958*rr + 58757137406*gg + 21832669565*bb) / 1e7)
	y = int((33952010375*rr + 109388662546*gg + 9249546043*bb) / 1e7)
	z = int((2125917447*rr + 14817196917*gg + 108975860226*bb) / 1e7)
	return
}

// XYZToSRGB_D50 converts from CIE XYZ to sRGB with D50 illuminator. 
func XYZToSRGB_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((31339039*xx + -16168668*yy + -4906146*zz) / 152587890625)
	gg := int((-9787684*xx + 19161707*yy + 334545*zz) / 152587890625)
	bb := int((719463*xx + -2289914*yy + 14052641*zz) / 152587890625)
	if rr < 0 {
		rr = 0
	} else if rr > 65535 {
		rr = 65535
	}
	if gg < 0 {
		gg = 0
	} else if gg > 65535 {
		gg = 65535
	}
	if bb < 0 {
		bb = 0
	} else if bb > 65535 {
		bb = 65535
	}
	r, g, b = uint16(rr), uint16(gg), uint16(bb)
	return
}
