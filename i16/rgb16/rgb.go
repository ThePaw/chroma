// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

// Conversion of different RGB colorspaces with their native illuminators (reference whites) to CIE XYZ scaled to [0, 1e9] and back. 
// RGB values MUST BE LINEAR and in the nominal range [0, 2^16]. 
// XYZ values are usually in [0, 1e9] but may be slightly outside this interval.
// To get quick and dirty XYZ [0.0, 1.0] approximations, divide by 1e9, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ converts from Adobe RGB 1998 with D65 illuminator to CIE XYZ. 
func AdobeToXYZ(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 88002151489 *rr +  28313293457 *gg+  28714782714 *bb)/1e7)
	 y  = int(( 45376113891 *rr +  95725875854 *gg+  11485916137 *bb)/1e7)
	 z  = int(( 4125106811 *rr +  10786010742 *gg+  151231155395 *bb)/1e7)
	return
}

// XYZToAdobe converts from CIE XYZ to Adobe RGB 1998 with D65 illuminator. 
func XYZToAdobe(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 20413690 *xx -5649464 *yy-3446944 *zz)/152587890625)
	 gg  := int(( -9692660 *xx +  18760108 *yy+  415560 *zz)/152587890625)
	 bb  := int(( 134474 *xx -1183897 *yy+  10154096 *zz)/152587890625)

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
func AppleToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 68623168945 *rr +  48255706787 *gg+  28151336669 *bb)/1e7)
	 y  = int(( 37331008911 *rr +  102543380737 *gg+  12713500976 *bb)/1e7)
	 z  = int(( 3842895507 *rr +  21542724609 *gg+  140756652832 *bb)/1e7)
	return
}





// XYZToApple converts from CIE XYZ to Apple RGB with D65 illuminator. 
func XYZToApple(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 29515373 *xx -12894116 *yy-4738445 *zz)/152587890625)
	 gg  := int(( -10851093 *xx +  19908566 *yy+  372026 *zz)/152587890625)
	 bb  := int(( 854934 *xx -2694965 *yy+  10912975 *zz)/152587890625)
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
func BestToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 96537719726 *rr +  31212738037 *gg+  19377838134 *bb)/1e7)
	 y  = int(( 34859756469 *rr +  112511032104 *gg+  5217102050 *bb)/1e7)
	 z  = int(( 0 *rr +  1451751708 *gg+  124465301513 *bb)/1e7)
	return
}


// XYZToBest converts from CIE XYZ to Best RGB with D50 illuminator. 
func XYZToBest(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 17552599 *xx -4836786 *yy-2530000 *zz)/152587890625)
	 gg  := int(( -5441336 *xx +  15068789 *yy+  215528 *zz)/152587890625)
	 bb  := int(( 63467 *xx -175761 *yy+  12256959 *zz)/152587890625)

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
func BetaToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 102425186157 *rr +  26639312744 *gg+  18063796997 *bb)/1e7)
	 y  = int(( 46275726318 *rr +  101285720825 *gg+  5026443481 *bb)/1e7)
	 z  = int(( 0 *rr +  6210479736 *gg+  119706573486 *bb)/1e7)
	return
}



// XYZToBeta converts from CIE XYZ to Beta RGB with D50 illuminator. 
func XYZToBeta(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 16832270 *xx -4282363 *yy-2360185 *zz)/152587890625)
	 gg  := int(( -7710229 *xx +  17065571 *yy+  446900 *zz)/152587890625)
	 bb  := int(( 400012 *xx -885376 *yy+  12723640 *zz)/152587890625)

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
func BruceToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 71322052001 *rr +  44929687500 *gg+  28778472900 *bb)/1e7)
	 y  = int(( 36775436401 *rr +  104301071166 *gg+  11511383056 *bb)/1e7)
	 z  = int(( 3343215942 *rr +  11232421875 *gg+  151566635131 *bb)/1e7)
	return
}



// XYZToBruce converts from CIE XYZ to Bruce RGB with D65 illuminator. 
func XYZToBruce(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 27454669 *xx -11358136 *yy-4350269 *zz)/152587890625)
	 gg  := int(( -9692660 *xx +  18760108 *yy+  415560 *zz)/152587890625)
	 bb  := int(( 112723 *xx -1139754 *yy+  10132541 *zz)/152587890625)

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
func CIEToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 74572448730 *rr +  47406051635 *gg+  30609390258 *bb)/1e7)
	 y  = int(( 26886657714 *rr +  124051620483 *gg+  1649612426 *bb)/1e7)
	 z  = int(( 0 *rr +  1557128906 *gg+  151030761718 *bb)/1e7)
	return
}



// XYZToCIE converts from CIE XYZ to CIE RGB with E illuminator. 
func XYZToCIE(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 23706743 *xx -9000405 *yy-4706338 *zz)/152587890625)
	 gg  := int(( -5138850 *xx +  14253035 *yy+  885814 *zz)/152587890625)
	 bb  := int(( 52982 *xx -146949 *yy+  10093968 *zz)/152587890625)

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
func ColorMatchToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 77719711303 *rr +  48966537475 *gg+  20442062377 *bb)/1e7)
	 y  = int(( 41943969726 *rr +  100422897338 *gg+  10221023559 *bb)/1e7)
	 z  = int(( 3700942993 *rr +  16598831176 *gg+  105617294311 *bb)/1e7)
	return
}



// XYZToColorMatch converts from CIE XYZ to ColorMatch RGB with D50 illuminator. 
func XYZToColorMatch(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 26422874 *xx -12234270 *yy-3930143 *zz)/152587890625)
	 gg  := int(( -11119763 *xx +  20590183 *yy+  159614 *zz)/152587890625)
	 bb  := int(( 821699 *xx -2807254 *yy+  14559877 *zz)/152587890625)

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
func DonToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 98536849975 *rr +  29503036499 *gg+  19088409423 *bb)/1e7)
	 y  = int(( 42472778320 *rr +  104975921630 *gg+  5139190673 *bb)/1e7)
	 z  = int(( 566299438 *rr +  2744461059 *gg+  122606277465 *bb)/1e7)
	return
}



// XYZToDon converts from CIE XYZ to Don RGB-4 with D50 illuminator. 
func XYZToDon(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 17603902 *xx -4881198 *yy-2536126 *zz)/152587890625)
	 gg  := int(( -7126288 *xx +  16527432 *yy+  416715 *zz)/152587890625)
	 bb  := int(( 78207 *xx -347411 *yy+  12447743 *zz)/152587890625)

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




// EciToXYZ converts from  ECI RGB with D50 illuminator to CIE XYZ. 
func EciToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 99213302612 *rr +  27172454833 *gg+  20742553710 *bb)/1e7)
	 y  = int(( 48866256713 *rr +  91868759155 *gg+  11852890014 *bb)/1e7)
	 z  = int(( 0 *rr +  10351409912 *gg+  115565643310 *bb)/1e7)
	return
}



// XYZToEci converts from CIE XYZ to ECI RGB with D50 illuminator. 
func XYZToEci(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 17827618 *xx -4969847 *yy-2690101 *zz)/152587890625)
	 gg  := int(( -9593623 *xx +  19477962 *yy-275807 *zz)/152587890625)
	 bb  := int(( 859317 *xx -1744674 *yy+  13228273 *zz)/152587890625)

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
func EktaSpaceToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 90620635986 *rr +  41653457641 *gg+  14854202270 *bb)/1e7)
	 y  = int(( 39768768310 *rr +  112143936157 *gg+  675186157 *bb)/1e7)
	 z  = int(( 0 *rr +  6408218383 *gg+  119508834838 *bb)/1e7)
	return
}




// XYZToEktaSpace converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. 
func XYZToEktaSpace(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 20043819 *xx -7304844 *yy-2450052 *zz)/152587890625)
	 gg  := int(( -7110285 *xx +  16202125 *yy+  792226 *zz)/152587890625)
	 bb  := int(( 381263 *xx -868780 *yy+  12725438 *zz)/152587890625)

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
func NTSCToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 92604202270 *rr +  26474166870 *gg+  30570678710 *bb)/1e7)
	 y  = int(( 45611022949 *rr +  89507904052 *gg+  17468948364 *bb)/1e7)
	 z  = int(( 0 *rr +  10085403442 *gg+  170322311401 *bb)/1e7)
	return
}



// XYZToNTSC converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 19099961 *xx -5324542 *yy-2882091 *zz)/152587890625)
	 gg  := int(( -9846663 *xx +  19991710 *yy-283082 *zz)/152587890625)
	 bb  := int(( 583056 *xx -1183781 *yy+  8975535 *zz)/152587890625)

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
func PALToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 65707244873 *rr +  52115158081 *gg+  27207809448 *bb)/1e7)
	 y  = int(( 33880294799 *rr +  107824462890 *gg+  10883117675 *bb)/1e7)
	 z  = int(( 3080032348 *rr +  19767822265 *gg+  143294433593 *bb)/1e7)
	return
}



// XYZToPAL converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. 
func XYZToPAL(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 30628971 *xx -13931791 *yy-4757517 *zz)/152587890625)
	 gg  := int(( -9692660 *xx +  18760108 *yy+  415560 *zz)/152587890625)
	 bb  := int(( 678774 *xx -2288548 *yy+  10693490 *zz)/152587890625)

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
func ProPhotoToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 121715530395 *rr +  20628616333 *gg+  4784149169 *bb)/1e7)
	 y  = int(( 43951446533 *rr +  108623367309 *gg+  13076782 *bb)/1e7)
	 z  = int(( 0 *rr +  0 *gg+  125917053222 *bb)/1e7)
	return
}



// XYZToProPhoto converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. 
func XYZToProPhoto(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 13459433 *xx -2556075 *yy-511118 *zz)/152587890625)
	 gg  := int(( -5445989 *xx +  15081673 *yy+  205351 *zz)/152587890625)
	 bb  := int(( 0 *xx +  0 *yy+  12118128 *zz)/152587890625)

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
func SMPTE_CToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 60056930541 *rr +  55732681274 *gg+  29240615844 *bb)/1e7)
	 y  = int(( 32411682128 *rr +  106970779418 *gg+  13205444335 *bb)/1e7)
	 z  = int(( 2859848022 *rr +  17079360961 *gg+  146203048706 *bb)/1e7)
	return
}




// XYZToSMPTE_C converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. 
func XYZToSMPTE_C(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 35053960 *xx -17394894 *yy-5439640 *zz)/152587890625)
	 gg  := int(( -10690722 *xx +  19778245 *yy+  351722 *zz)/152587890625)
	 bb  := int(( 563200 *xx -1970226 *yy+  10502026 *zz)/152587890625)

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
func SRGBToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 62935852050 *rr +  54561782836 *gg+  27532577514 *bb)/1e7)
	 y  = int(( 32451309204 *rr +  109123565673 *gg+  11013031005 *bb)/1e7)
	 z  = int(( 2950119018 *rr +  18187255859 *gg+  145004898071 *bb)/1e7)
	return
}



// XYZToSRGB converts from CIE XYZ to sRGB with D65 illuminator. 
func XYZToSRGB(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 32404542 *xx -15371385 *yy-4985314 *zz)/152587890625)
	 gg  := int(( -9692660 *xx +  18760108 *yy+  415560 *zz)/152587890625)
	 bb  := int(( 556434 *xx -2040259 *yy+  10572252 *zz)/152587890625)

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
func WGamutToXYZ(r,g,b uint16) (x, y, z int) {
	rr:=int64(r)
	gg:=int64(g)
	bb:=int64(b)
	 x  = int(( 109268890380 *rr +  15400634765 *gg+  22458770751 *bb)/1e7)
	 y  = int(( 39396270751 *rr +  110616729736 *gg+  2574890136 *bb)/1e7)
	 z  = int(( 0 *rr +  7901199340 *gg+  118015853881 *bb)/1e7)
	return
}



// XYZToWGamut converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. 
func XYZToWGamut(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 14628067 *xx -1840624 *yy-2743606 *zz)/152587890625)
	 gg  := int(( -5217933 *xx +  14472381 *yy+  677227 *zz)/152587890625)
	 bb  := int(( 349342 *xx -968931 *yy+  12884099 *zz)/152587890625)

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

