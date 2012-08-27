// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb16

// Conversion of natively non-D50 RGB colorspaces with D50 illuminator to CIE XYZ and back. 
// Bradford adaptation was used to calculate D50 matrices from colorspaces' native illuminators. 
// RGB values must be linear and in the nominal range [0, 255]. 
// XYZ values are usually in [0, 255] but may be greater
// To get quick and dirty XYZ approximations, divide by 255, otherwise use the float64 version of these functions. 
// Ref.: [24] 

// AdobeToXYZ_D50 converts from Adobe RGB 1998 with D50 illuminator to CIE XYZ. 
func AdobeToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 93041366577 *rr +  31317153930 *gg+  22769775390 *bb)/1e7)
	 y  = int(( 47473785400 *rr +  95467529296 *gg+  9646560668 *bb)/1e7)
	 z  = int(( 2972579956 *rr +  9291107177 *gg+  113653366088 *bb)/1e7)
	return
}


// XYZToAdobe_D50 converts from CIE XYZ to Adobe RGB 1998 with D50 illuminator. 
func XYZToAdobe_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 19624274 *xx +  -6105343 *yy+  -3413404 *zz)/152587890625)
	 gg  := int(( -9787684 *xx +  19161415 *yy+  334540 *zz)/152587890625)
	 bb  := int(( 286869 *xx +  -1406752 *yy+  13487655 *zz)/152587890625)
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
func AppleToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 72565887451 *rr +  51829864501 *gg+  22732543945 *bb)/1e7)
	 y  = int(( 38937561035 *rr +  102625930786 *gg+  11024414062 *bb)/1e7)
	 z  = int(( 2818252563 *rr +  17299972534 *gg+  105798828125 *bb)/1e7)
	return
}



// XYZToApple_D50 converts from CIE XYZ to Apple RGB with D50 illuminator. 
func XYZToApple_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 28510695 *xx +  -13605261 *yy+  -4708281 *zz)/152587890625)
	 gg  := int(( -10927680 *xx +  20348871 *yy+  227598 *zz)/152587890625)
	 bb  := int(( 1027403 *xx +  -2964984 *yy+  14510659 *zz)/152587890625)
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
func BruceToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 75406127929 *rr +  48901885986 *gg+  22820281982 *bb)/1e7)
	 y  = int(( 38475509643 *rr +  104444412231 *gg+  9667968750 *bb)/1e7)
	 z  = int(( 2409149169 *rr +  9602416992 *gg+  113905471801 *bb)/1e7)
	return
}

// XYZToBruce_D50 converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func XYZToBruce_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 26502856 *xx +  -12014485 *yy+  -4289936 *zz)/152587890625)
	 gg  := int(( -9787684 *xx +  19161415 *yy+  334540 *zz)/152587890625)
	 bb  := int(( 264570 *xx +  -1361228 *yy+  13458542 *zz)/152587890625)
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

// CieToXYZ_D50 converts from CIE RGB with D50 illuminator to CIE XYZ. 
func CieToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 74293060302 *rr +  46737426757 *gg+  26097824096 *bb)/1e7)
	 y  = int(( 26650741577 *rr +  125847488403 *gg+  89675903 *bb)/1e7)
	 z  = int(( -191696167 *rr +  2591430664 *gg+  123517318725 *bb)/1e7)
	return
}

// XYZToCie_D50 converts from CIE XYZ to CIE RGB with D50 illuminator. 
func XYZToCie_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 23638081 *xx +  -8676030 *yy+  -4988161 *zz)/152587890625)
	 gg  := int(( -5005940 *xx +  13962368 *yy+  1047561 *zz)/152587890625)
	 bb  := int(( 141712 *xx +  -306400 *yy+  12323842 *zz)/152587890625)
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
func NTSCToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 96797271728 *rr +  28262390136 *gg+  22068634033 *bb)/1e7)
	 y  = int(( 47447143554 *rr +  90270751953 *gg+  14869995117 *bb)/1e7)
	 z  = int(( -180313111 *rr +  8476531982 *gg+  117620834350 *bb)/1e7)
	return
}



// XYZToNTSC_D50 converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func XYZToNTSC_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 18464881 *xx +  -5521299 *yy+  -2766458 *zz)/152587890625)
	 gg  := int(( -9826630 *xx +  20044755 *yy+  -690397 *zz)/152587890625)
	 bb  := int(( 736477 *xx +  -1453020 *yy+  13018376 *zz)/152587890625)
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
func PALToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 69469802856 *rr +  56083679199 *gg+  21574798583 *bb)/1e7)
	 y  = int(( 35446548461 *rr +  108001037597 *gg+  9140304565 *bb)/1e7)
	 z  = int(( 2219497680 *rr +  16008819580 *gg+  107688735961 *bb)/1e7)
	return
}



// XYZToPAL_D50 converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func XYZToPAL_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 29603944 *xx +  -14678520 *yy+  -4685105 *zz)/152587890625)
	 gg  := int(( -9787684 *xx +  19161415 *yy+  334540 *zz)/152587890625)
	 bb  := int(( 844874 *xx +  -2545974 *yy+  14216174 *zz)/152587890625)
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
func SMPTE_CToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 63526763916 *rr +  59989379882 *gg+  23612152099 *bb)/1e7)
	 y  = int(( 33828720092 *rr +  107308181762 *gg+  11450988769 *bb)/1e7)
	 z  = int(( 2083984375 *rr +  13940490722 *gg+  109892578125 *bb)/1e7)
	return
}


// XYZToSMPTE_C_D50 converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func XYZToSMPTE_C_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 33921940 *xx +  -18264027 *yy+  -5385522 *zz)/152587890625)
	 gg  := int(( -10770996 *xx +  20213975 *yy+  207989 *zz)/152587890625)
	 bb  := int(( 723073 *xx +  -2217902 *yy+  13960932 *zz)/152587890625)
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
func SRGBToXYZ_D50(r,g,b uint16) (x, y, z int) {
rr:=int64(r)
gg:=int64(g)
bb:=int64(b)
	 x  = int(( 66539718627 *rr +  58756240844 *gg+  21832336425 *bb)/1e7)
	 y  = int(( 33951492309 *rr +  109386993408 *gg+  9249404907 *bb)/1e7)
	 z  = int(( 2125885009 *rr +  14816970825 *gg+  108974197387 *bb)/1e7)
	return
}




// XYZToSRGB_D50 converts from CIE XYZ to sRGB with D50 illuminator. 
func XYZToSRGB_D50(x, y, z int) (r, g, b uint16) {
	xx, yy, zz := int64(x), int64(y), int64(z)
	 rr  := int(( 31338561 *xx +  -16168668 *yy+  -4906146 *zz)/152587890625)
	 gg  := int(( -9787684 *xx +  19161415 *yy+  334540 *zz)/152587890625)
	 bb  := int(( 719453 *xx +  -2289914 *yy+  14052427 *zz)/152587890625)
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


