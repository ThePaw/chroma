// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert different RGB colorspaces with their native illuminator to CIE XYZ and back. 
// RGB values must be linear and in the nominal range [0.0, 1.0]. 
// Ref.: [24][30][31]. 
package rgb

// AdobeToXYZ converts from Adobe RGB (1998) with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func AdobeToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.5767309*r + 0.1855540*g + 0.1881852*b
	y = 0.2973769*r + 0.6273491*g + 0.0752741*b
	z = 0.0270343*r + 0.0706872*g + 0.9911085*b
	return
}

//  XYZToAdobe converts from CIE XYZ to Adobe RGB (1998) with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToAdobe(x, y, z float64) (r, g, b float64) {
	r = 2.0413690*x -0.5649464*y -0.3446944*z
	g = -0.9692660*x + 1.8760108*y + 0.0415560*z
	b = 0.0134474*x -0.1183897*y + 1.0154096*z
	return
}

// AppleToXYZ converts from Apple RGB with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func AppleToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4497288*r + 0.3162486*g + 0.1844926*b
	y = 0.2446525*r + 0.6720283*g + 0.0833192*b
	z = 0.0251848*r + 0.1411824*g + 0.9224628*b
	return
}

// XYZToApple converts from CIE XYZ to Apple RGB with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToApple(x, y, z float64) (r, g, b float64) {
	r = 2.9515373*x -1.2894116*y -0.4738445*z
	g = -1.0851093*x + 1.9908566*y + 0.0372026*z
	b = 0.0854934*x -0.2694964*y + 1.0912975*z
	return
}

// BestToXYZ converts from Best RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func BestToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.6326696*r + 0.2045558*g + 0.1269946*b
	y = 0.2284569*r + 0.7373523*g + 0.0341908*b
	z = 0.0000000*r + 0.0095142*g + 0.8156958*b
	return
}

// XYZToBest converts from CIE XYZ to Best RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToBest(x, y, z float64) (r, g, b float64) {
	r = 1.7552599*x -0.4836786*y -0.2530000*z
	g = -0.5441336*x + 1.5068789*y + 0.0215528*z
	b = 0.0063467*x -0.0175761*y + 1.2256959*z
	return
}

// BetaToXYZ converts from Beta RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func BetaToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.6712537*r + 0.1745834*g + 0.1183829*b
	y = 0.3032726*r + 0.6637861*g + 0.0329413*b
	z = 0.0000000*r + 0.0407010*g + 0.7845090*b
	return
}

// XYZToBeta converts from CIE XYZ to Beta RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToBeta(x, y, z float64) (r, g, b float64) {
	r = 1.6832270*x -0.4282363*y -0.2360185*z
	g = -0.7710229*x + 1.7065571*y + 0.0446900*z
	b = 0.0400013*x -0.0885376*y + 1.2723640*z
	return
}

// BruceToXYZ converts from Bruce RGB with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func BruceToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4674162*r + 0.2944512*g + 0.1886026*b
	y = 0.2410115*r + 0.6835475*g + 0.0754410*b
	z = 0.0219101*r + 0.0736128*g + 0.9933071*b
	return
}

// XYZToBruce converts from CIE XYZ to Bruce RGB with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToBruce(x, y, z float64) (r, g, b float64) {
	r = 2.7454669*x -1.1358136*y -0.4350269*z
	g = -0.9692660*x + 1.8760108*y + 0.0415560*z
	b = 0.0112723*x -0.1139754*y + 1.0132541*z
	return
}

// CIEToXYZ converts from CIE RGB with E illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func CIEToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4887180*r + 0.3106803*g + 0.2006017*b
	y = 0.1762044*r + 0.8129847*g + 0.0108109*b
	z = 0.0000000*r + 0.0102048*g + 0.9897952*b
	return
}

// XYZToCIE converts from CIE XYZ to CIE RGB with E illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToCIE(x, y, z float64) (r, g, b float64) {
	r = 2.3706743*x -0.9000405*y -0.4706338*z
	g = -0.5138850*x + 1.4253036*y + 0.0885814*z
	b = 0.0052982*x -0.0146949*y + 1.0093968*z
	return
}

// ColorMatchToXYZ converts from  ColorMatch RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func ColorMatchToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.5093439*r + 0.3209071*g + 0.1339691*b
	y = 0.2748840*r + 0.6581315*g + 0.0669845*b
	z = 0.0242545*r + 0.1087821*g + 0.6921735*b
	return
}

// XYZToColorMatch converts from CIE XYZ to ColorMatch RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToColorMatch(x, y, z float64) (r, g, b float64) {
	r = 2.6422874*x -1.2234270*y -0.3930143*z
	g = -1.1119763*x + 2.0590183*y + 0.0159614*z
	b = 0.0821699*x -0.2807254*y + 1.4559877*z
	return
}

// DonToXYZ converts from  Don RGB-4 with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func DonToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.6457711*r + 0.1933511*g + 0.1250978*b
	y = 0.2783496*r + 0.6879702*g + 0.0336802*b
	z = 0.0037113*r + 0.0179861*g + 0.8035125*b
	return
}

// XYZToDon converts from CIE XYZ to Don RGB-4 with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToDon(x, y, z float64) (r, g, b float64) {
	r = 1.7603902*x -0.4881198*y -0.2536126*z
	g = -0.7126288*x + 1.6527432*y + 0.0416715*z
	b = 0.0078207*x -0.0347411*y + 1.2447743*z
	return
}

// EciToXYZ converts from  ECI RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func EciToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.6502043*r + 0.1780774*g + 0.1359384*b
	y = 0.3202499*r + 0.6020711*g + 0.0776791*b
	z = 0.0000000*r + 0.0678390*g + 0.7573710*b
	return
}

// XYZToEci converts from CIE XYZ to ECI RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToEci(x, y, z float64) (r, g, b float64) {
	r = 1.7827618*x -0.4969847*y -0.2690101*z
	g = -0.9593623*x + 1.9477962*y -0.0275807*z
	b = 0.0859317*x -0.1744674*y + 1.3228273*z
	return
}

// EktaSpaceToXYZ converts from  Ekta Space PS5 with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func EktaSpaceToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.5938914*r + 0.2729801*g + 0.0973485*b
	y = 0.2606286*r + 0.7349465*g + 0.0044249*b
	z = 0.0000000*r + 0.0419969*g + 0.7832131*b
	return
}

// XYZToEktaSpace converts from CIE XYZ to Ekta Space PS5 with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToEktaSpace(x, y, z float64) (r, g, b float64) {
	r = 2.0043819*x -0.7304844*y -0.2450052*z
	g = -0.7110285*x + 1.6202126*y + 0.0792227*z
	b = 0.0381263*x -0.0868780*y + 1.2725438*z
	return
}

// NTSCToXYZ converts from  NTSC RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func NTSCToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.6068909*r + 0.1735011*g + 0.2003480*b
	y = 0.2989164*r + 0.5865990*g + 0.1144845*b
	z = 0.0000000*r + 0.0660957*g + 1.1162243*b
	return
}

// XYZToNTSC converts from CIE XYZ to NTSC RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToNTSC(x, y, z float64) (r, g, b float64) {
	r = 1.9099961*x -0.5324542*y -0.2882091*z
	g = -0.9846663*x + 1.9991710*y -0.0283082*z
	b = 0.0583056*x -0.1183781*y + 0.8975535*z
	return
}

// PALToXYZ converts from  PAL/SECAM RGB with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func PALToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4306190*r + 0.3415419*g + 0.1783091*b
	y = 0.2220379*r + 0.7066384*g + 0.0713236*b
	z = 0.0201853*r + 0.1295504*g + 0.9390944*b
	return
}

// XYZToPAL converts from CIE XYZ to PAL/SECAM RGB with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToPAL(x, y, z float64) (r, g, b float64) {
	r = 3.0628971*x -1.3931791*y -0.4757517*z
	g = -0.9692660*x + 1.8760108*y + 0.0415560*z
	b = 0.0678775*x -0.2288548*y + 1.0693490*z
	return
}

// ProPhotoToXYZ converts from  ProPhoto RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func ProPhotoToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.7976749*r + 0.1351917*g + 0.0313534*b
	y = 0.2880402*r + 0.7118741*g + 0.0000857*b
	z = 0.0000000*r + 0.0000000*g + 0.8252100*b
	return
}

// XYZToProPhoto converts from CIE XYZ to  ProPhoto RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToProPhoto(x, y, z float64) (r, g, b float64) {
	r = 1.3459433*x -0.2556075*y -0.0511118*z
	g = -0.5445989*x + 1.5081673*y + 0.0205351*z
	b = 0.0000000*x + 0.0000000*y + 1.2118128*z
	return
}

// SMPTE_CToXYZ converts from  SMPTE-C RGB with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func SMPTE_CToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.3935891*r + 0.3652497*g + 0.1916313*b
	y = 0.2124132*r + 0.7010437*g + 0.0865432*b
	z = 0.0187423*r + 0.1119313*g + 0.9581563*b
	return
}

// XYZToSMPTE_C converts from CIE XYZ to SMPTE-C RGB with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToSMPTE_C(x, y, z float64) (r, g, b float64) {
	r = 3.5053960*x -1.7394894*y -0.5439640*z
	g = -1.0690722*x + 1.9778245*y + 0.0351722*z
	b = 0.0563200*x -0.1970226*y + 1.0502026*z
	return
}

// SRGBToXYZ converts from  sRGB with D65 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func SRGBToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4124564*r + 0.3575761*g + 0.1804375*b
	y = 0.2126729*r + 0.7151522*g + 0.0721750*b
	z = 0.0193339*r + 0.1191920*g + 0.9503041*b
	return
}

// XYZToSRGB converts from CIE XYZ to sRGB with D65 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToSRGB(x, y, z float64) (r, g, b float64) {
	r = 3.2404542*x -1.5371385*y -0.4985314*z
	g = -0.9692660*x + 1.8760108*y + 0.0415560*z
	b = 0.0556434*x -0.2040259*y + 1.0572252*z
	return
}

// WGamutToXYZ converts from  Wide Gamut RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func WGamutToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.7161046*r + 0.1009296*g + 0.1471858*b
	y = 0.2581874*r + 0.7249378*g + 0.0168748*b
	z = 0.0000000*r + 0.0517813*g + 0.7734287*b
	return
}

// XYZToWGamut converts from CIE XYZ to Wide Gamut RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToWGamut(x, y, z float64) (r, g, b float64) {
	r = 1.4628067*x -0.1840623*y -0.2743606*z
	g = -0.5217933*x + 1.4472381*y + 0.0677227*z
	b = 0.0349342*x -0.0968930*y + 1.2884099*z
	return
}
