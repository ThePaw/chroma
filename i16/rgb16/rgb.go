// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

// Conversion of different RGB colorspaces with their native illuminators (reference whites) to CIE XYZ scaled to [0, 1e9] and back. 
// RGB values MUST BE LINEAR and in the nominal range [0, 2^16]. 
// XYZ values are usually in [0, 1e9] but may be slightly outside this interval.
// To get quick and dirty XYZ [0.0, 1.0] approximations, divide by 1e9, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ converts from Adobe RGB 1998 with D65 illuminator to CIE XYZ.
func AdobeToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((88003494315*rr + 28313725490*gg + 28715220873*bb) / 1e7)
	y = int((45376806286*rr + 95727336537*gg + 11486091400*bb) / 1e7)
	z = int((4125169756*rr + 10786175325*gg + 151233463034*bb) / 1e7)
	return
}

// XYZToAdobe converts from CIE XYZ to Adobe RGB 1998 with D65 illuminator.
func XYZToAdobe(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((20414001*xx - 5649464*yy - 3446944*zz) / 152587890625)
	gg := int((-9692660*xx + 18760394*yy + 415566*zz) / 152587890625)
	bb := int((134476*xx - 1183897*yy + 10154250*zz) / 152587890625)

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

// AppleToXYZ converts from Apple RGB with D65 illuminator to CIE XYZ. 
func AppleToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((68624216067*rr + 48256443121*gg + 28151766230*bb) / 1e7)
	y = int((37331578545*rr + 102544945448*gg + 12713694971*bb) / 1e7)
	z = int((3842954145*rr + 21543053329*gg + 140758800640*bb) / 1e7)
	return
}

// XYZToApple converts from CIE XYZ to Apple RGB with D65 illuminator. 
func XYZToApple(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((29515823*xx - 12894116*yy - 4738445*zz) / 152587890625)
	gg := int((-10851093*xx + 19908869*yy + 372031*zz) / 152587890625)
	bb := int((854947*xx - 2694965*yy + 10913141*zz) / 152587890625)
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

// BestToXYZ converts from Best RGB with D50 illuminator to CIE XYZ. 
func BestToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((96539192797*rr + 31213214312*gg + 19378133820*bb) / 1e7)
	y = int((34860288394*rr + 112512748912*gg + 5217181657*bb) / 1e7)
	z = int((0*rr + 1451773860*gg + 124467200731*bb) / 1e7)
	return
}

// XYZToBest converts from CIE XYZ to Best RGB with D50 illuminator. 
func XYZToBest(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((17552866*xx - 4836786*yy - 2530000*zz) / 152587890625)
	gg := int((-5441336*xx + 15069018*yy + 215531*zz) / 152587890625)
	bb := int((63467*xx - 175761*yy + 12257146*zz) / 152587890625)

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

// BetaToXYZ converts from Beta RGB with D50 illuminator to CIE XYZ. 
func BetaToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((102426749065*rr + 26639719233*gg + 18064072632*bb) / 1e7)
	y = int((46276432440*rr + 101287266346*gg + 5026520179*bb) / 1e7)
	z = int((0*rr + 6210574501*gg + 119708400091*bb) / 1e7)
	return
}

// XYZToBeta converts from CIE XYZ to Beta RGB with D50 illuminator. 
func XYZToBeta(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((16832526*xx - 4282363*yy - 2360185*zz) / 152587890625)
	gg := int((-7710229*xx + 17065831*yy + 446906*zz) / 152587890625)
	bb := int((400018*xx - 885376*yy + 12723834*zz) / 152587890625)

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

// BruceToXYZ converts from Bruce RGB with D65 illuminator to CIE XYZ. 
func BruceToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((71323140305*rr + 44930373083*gg + 28778912031*bb) / 1e7)
	y = int((36775997558*rr + 104302662698*gg + 11511558708*bb) / 1e7)
	z = int((3343266956*rr + 11232593270*gg + 151568947889*bb) / 1e7)
	return
}

// XYZToBruce converts from CIE XYZ to Bruce RGB with D65 illuminator. 
func XYZToBruce(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((27455087*xx - 11358136*yy - 4350269*zz) / 152587890625)
	gg := int((-9692660*xx + 18760394*yy + 415566*zz) / 152587890625)
	bb := int((112724*xx - 1139754*yy + 10132695*zz) / 152587890625)

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

// CIEToXYZ converts from CIE RGB with E illuminator to CIE XYZ. 
func CIEToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((74573586632*rr + 47406775004*gg + 30609857327*bb) / 1e7)
	y = int((26887067978*rr + 124053513389*gg + 1649637597*bb) / 1e7)
	z = int((0*rr + 1557152666*gg + 151033066299*bb) / 1e7)
	return
}

// XYZToCIE converts from CIE XYZ to CIE RGB with E illuminator. 
func XYZToCIE(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((23707104*xx - 9000405*yy - 4706338*zz) / 152587890625)
	gg := int((-5138850*xx + 14253252*yy + 885827*zz) / 152587890625)
	bb := int((52982*xx - 146949*yy + 10094122*zz) / 152587890625)

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

// ColorMatchToXYZ converts from  ColorMatch RGB with D50 illuminator to CIE XYZ. 
func ColorMatchToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((77720897229*rr + 48967284656*gg + 20442374302*bb) / 1e7)
	y = int((41944609749*rr + 100424429693*gg + 10221179521*bb) / 1e7)
	z = int((3700999465*rr + 16599084457*gg + 105618905927*bb) / 1e7)
	return
}

// XYZToColorMatch converts from CIE XYZ to ColorMatch RGB with D50 illuminator. 
func XYZToColorMatch(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((26423277*xx - 12234270*yy - 3930143*zz) / 152587890625)
	gg := int((-11119763*xx + 20590497*yy + 159616*zz) / 152587890625)
	bb := int((821711*xx - 2807254*yy + 14560099*zz) / 152587890625)

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

//DonToXYZ converts from  Don RGB-4 with D50 illuminator to CIE XYZ. 
func DonToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((98538353550*rr + 29503486686*gg + 19088700693*bb) / 1e7)
	y = int((42473426413*rr + 104977523459*gg + 5139269092*bb) / 1e7)
	z = int((566308079*rr + 2744502936*gg + 122608148316*bb) / 1e7)
	return
}

// XYZToDon converts from CIE XYZ to Don RGB-4 with D50 illuminator. 
func XYZToDon(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((17604170*xx - 4881198*yy - 2536126*zz) / 152587890625)
	gg := int((-7126288*xx + 16527684*yy + 416721*zz) / 152587890625)
	bb := int((78208*xx - 347411*yy + 12447932*zz) / 152587890625)

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

// ECIToXYZ converts from  ECI RGB with D50 illuminator to CIE XYZ. 
func ECIToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((99214816509*rr + 27172869458*gg + 20742870221*bb) / 1e7)
	y = int((48867002364*rr + 91870160982*gg + 11853070877*bb) / 1e7)
	z = int((0*rr + 10351567864*gg + 115567406728*bb) / 1e7)
	return
}

// XYZToECI converts from CIE XYZ to ECI RGB with D50 illuminator. 
func XYZToECI(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((17827890*xx - 4969847*yy - 2690101*zz) / 152587890625)
	gg := int((-9593623*xx + 19478259*yy - 275807*zz) / 152587890625)
	bb := int((859330*xx - 1744674*yy + 13228474*zz) / 152587890625)

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

// EktaSpaceToXYZ converts from  Ekta Space PS5 with D50 illuminator to CIE XYZ. 
func EktaSpaceToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((90622018768*rr + 41654093232*gg + 14854428930*bb) / 1e7)
	y = int((39769375142*rr + 112145647363*gg + 675196459*bb) / 1e7)
	z = int((0*rr + 6408316166*gg + 119510658425*bb) / 1e7)
	return
}

// XYZToEktaSpace converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. 
func XYZToEktaSpace(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((20044124*xx - 7304844*yy - 2450052*zz) / 152587890625)
	gg := int((-7110285*xx + 16202372*yy + 792238*zz) / 152587890625)
	bb := int((381268*xx - 868780*yy + 12725632*zz) / 152587890625)

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

// NTSCToXYZ converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func NTSCToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((92605615319*rr + 26474570839*gg + 30571145188*bb) / 1e7)
	y = int((45611718928*rr + 89509269855*gg + 17469214923*bb) / 1e7)
	z = int((0*rr + 10085557335*gg + 170324910352*bb) / 1e7)
	return
}

// XYZToNTSC converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((19100252*xx - 5324542*yy - 2882091*zz) / 152587890625)
	gg := int((-9846663*xx + 19992015*yy - 283082*zz) / 152587890625)
	bb := int((583064*xx - 1183781*yy + 8975671*zz) / 152587890625)

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

// PALToXYZ converts from  PAL/SECAM RGB with D65 illuminator to CIE XYZ. 
func PALToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((65708247501*rr + 52115953307*gg + 27208224612*bb) / 1e7)
	y = int((33880811779*rr + 107826108185*gg + 10883283740*bb) / 1e7)
	z = int((3080079346*rr + 19768123902*gg + 143296620125*bb) / 1e7)
	return
}

// XYZToPAL converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. 
func XYZToPAL(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((30629438*xx - 13931791*yy - 4757517*zz) / 152587890625)
	gg := int((-9692660*xx + 18760394*yy + 415566*zz) / 152587890625)
	bb := int((678784*xx - 2288548*yy + 10693653*zz) / 152587890625)

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

// ProPhotoToXYZ converts from  ProPhoto RGB with D50 illuminator to CIE XYZ. 
func ProPhotoToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((121717387654*rr + 20628931105*gg + 4784222170*bb) / 1e7)
	y = int((43952117189*rr + 108625024795*gg + 13076981*bb) / 1e7)
	z = int((0*rr + 0*gg + 125918974593*bb) / 1e7)
	return
}

// XYZToProPhoto converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. 
func XYZToProPhoto(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((13459638*xx - 2556075*yy - 511118*zz) / 152587890625)
	gg := int((-5445989*xx + 15081903*yy + 205354*zz) / 152587890625)
	bb := int((0*xx + 0*yy + 12118312*zz) / 152587890625)

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

// SMPTE_CToXYZ converts from  SMPTE-C RGB with D65 illuminator to CIE XYZ. 
func SMPTE_CToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((60057846951*rr + 55733531700*gg + 29241062027*bb) / 1e7)
	y = int((32412176698*rr + 106972411687*gg + 13205645837*bb) / 1e7)
	z = int((2859891660*rr + 17079621575*gg + 146205279621*bb) / 1e7)
	return
}

// XYZToSMPTE_C converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. 
func XYZToSMPTE_C(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((35054494*xx - 17394894*yy - 5439640*zz) / 152587890625)
	gg := int((-10690722*xx + 19778546*yy + 351727*zz) / 152587890625)
	bb := int((563208*xx - 1970226*yy + 10502186*zz) / 152587890625)

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

// SRGBToXYZ converts from  sRGB with D65 illuminator to CIE XYZ. 
func SRGBToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((62936812389*rr + 54562615395*gg + 27532997634*bb) / 1e7)
	y = int((32451804379*rr + 109125230791*gg + 11013199053*bb) / 1e7)
	z = int((2950164033*rr + 18187533378*gg + 145007110703*bb) / 1e7)
	return
}

// XYZToSRGB converts from CIE XYZ to sRGB with D65 illuminator. 
func XYZToSRGB(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((32405036*xx - 15371385*yy - 4985314*zz) / 152587890625)
	gg := int((-9692660*xx + 18760394*yy + 415566*zz) / 152587890625)
	bb := int((556442*xx - 2040259*yy + 10572413*zz) / 152587890625)

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

// WGamutToXYZ converts from  Wide Gamut RGB with D50 illuminator to CIE XYZ. 
func WGamutToXYZ(r, g, b uint16) (x, y, z int) {
	rr := int64(r)
	gg := int64(g)
	bb := int64(b)
	x = int((109270557716*rr + 15400869763*gg + 22459113449*bb) / 1e7)
	y = int((39396871899*rr + 110618417639*gg + 2574929426*bb) / 1e7)
	z = int((0*rr + 7901319904*gg + 118017654687*bb) / 1e7)
	return
}

// XYZToWGamut converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. 
func XYZToWGamut(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	rr := int((14628290*xx - 1840624*yy - 2743606*zz) / 152587890625)
	gg := int((-5217933*xx + 14472601*yy + 677237*zz) / 152587890625)
	bb := int((349347*xx - 968931*yy + 12884295*zz) / 152587890625)

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
