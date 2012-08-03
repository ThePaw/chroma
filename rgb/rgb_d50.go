// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert natively non-D50 RGB colorspaces with D50 illuminator to CIE XYZ and back. 
// Bradford adaptation was used to calculate D50 matrices from colorspaces' native illuminators. 
// RGB values must be linear and in the nominal range [0.0, 1.0]. 
// Ref.: [24][30][31]

package rgb

// AdobeToXYZ_D50 converts from Adobe RGB (1998) with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func AdobeToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.6097559*r + 0.2052401*g + 0.1492240*b)
	y = (0.3111242*r + 0.6256560*g + 0.0632197*b)
	z = (0.0194811*r + 0.0608902*g + 0.7448387*b)
	return
}

// XYZToAdobe_D50 converts from CIE XYZ to Adobe RGB(1998) with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToAdobe_D50(x, y, z float64) (r, g, b float64) {
	r = (1.9624274*x + -0.6105343*y + -0.3413404*z)
	g = (-0.9787684*x + 1.9161415*y + 0.0334540*z)
	b = (0.0286869*x + -0.1406752*y + 1.3487655*z)
	return
}

// AppleToXYZ_D50 converts from Apple RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func AppleToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4755678*r + 0.3396722*g + 0.1489800*b)
	y = (0.2551812*r + 0.6725693*g + 0.0722496*b)
	z = (0.0184697*r + 0.1133771*g + 0.6933632*b)
	return
}

// XYZToApple_D50 converts from CIE XYZ to Apple RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToApple_D50(x, y, z float64) (r, g, b float64) {
	return
}

// BruceToXYZ_D50 converts from Bruce RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func BruceToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4941816*r + 0.3204834*g + 0.1495550*b)
	y = (0.2521531*r + 0.6844869*g + 0.0633600*b)
	z = (0.0157886*r + 0.0629304*g + 0.7464909*b)
	return
}

// XYZToBruce_D50 converts from CIE XYZ to Bruce RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToBruce_D50(x, y, z float64) (r, g, b float64) {
	r = (2.6502856*x + -1.2014485*y + -0.4289936*z)
	g = (-0.9787684*x + 1.9161415*y + 0.0334540*z)
	b = (0.0264570*x + -0.1361227*y + 1.3458542*z)
	return
}

// CieToXYZ_D50 converts from CIE RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func CieToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4868870*r + 0.3062984*g + 0.1710347*b)
	y = (0.1746583*r + 0.8247541*g + 0.0005877*b)
	z = (-0.0012563*r + 0.0169832*g + 0.8094831*b)
	return
}

// XYZToCie converts from CIE XYZ to CIE RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToCie(x, y, z float64) (r, g, b float64) {
	r = (2.3638081*x + -0.8676030*y + -0.4988161*z)
	g = (-0.5005940*x + 1.3962369*y + 0.1047562*z)
	b = (0.0141712*x + -0.0306400*y + 1.2323842*z)
	return
}

// NTSCToXYZ_D50 converts from  NTSC RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func NTSCToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.6343706*r + 0.1852204*g + 0.1446290*b)
	y = (0.3109496*r + 0.5915984*g + 0.0974520*b)
	z = (-0.0011817*r + 0.0555518*g + 0.7708399*b)
	return
}

// XYZToNTSC_D50 converts from CIE XYZ to NTSC RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToNTSC_D50(x, y, z float64) (r, g, b float64) {
	r = (1.8464881*x + -0.5521299*y + -0.2766458*z)
	g = (-0.9826630*x + 2.0044755*y + -0.0690396*z)
	b = (0.0736477*x + -0.1453020*y + 1.3018376*z)
	return
}

// PALToXYZ_D50 converts from  PAL/SECAM RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func PALToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4552773*r + 0.3675500*g + 0.1413926*b)
	y = (0.2323025*r + 0.7077956*g + 0.0599019*b)
	z = (0.0145457*r + 0.1049154*g + 0.7057489*b)
	return
}

// XYZToPAL_D50 converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToPAL_D50(x, y, z float64) (r, g, b float64) {
	r = (2.9603944*x + -1.4678519*y + -0.4685105*z)
	g = (-0.9787684*x + 1.9161415*y + 0.0334540*z)
	b = (0.0844874*x + -0.2545973*y + 1.4216174*z)
	return
}

// SMPTE_CToXYZ_D50 converts from  SMPTE-C RGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func SMPTE_CToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4163290*r + 0.3931464*g + 0.1547446*b)
	y = (0.2216999*r + 0.7032549*g + 0.0750452*b)
	z = (0.0136576*r + 0.0913604*g + 0.7201920*b)
	return
}

// XYZToSMPTE_C_D50 converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToSMPTE_C_D50(x, y, z float64) (r, g, b float64) {
	r = (3.3921940*x + -1.8264027*y + -0.5385522*z)
	g = (-1.0770996*x + 2.0213975*y + 0.0207989*z)
	b = (0.0723073*x + -0.2217902*y + 1.3960932*z)
	return
}

// SRGBToXYZ_D50 converts from  sRGB with D50 illuminator to CIE XYZ. RGB values must be linear  and in the nominal range [0.0, 1.0]. 
func SRGBToXYZ_D50(r, g, b float64) (x, y, z float64) {
	x = (0.4360747*r + 0.3850649*g + 0.1430804*b)
	y = (0.2225045*r + 0.7168786*g + 0.0606169*b)
	z = (0.0139322*r + 0.0971045*g + 0.7141733*b)
	return
}

// XYZToSRGB_D50 converts from CIE XYZ to sRGB with D50 illuminator. Returned RGB values are linear and in the nominal range [0.0, 1.0]. 
func XYZToSRGB_D50(x, y, z float64) (r, g, b float64) {
	r = (3.1338561*x + -1.6168667*y + -0.4906146*z)
	g = (-0.9787684*x + 1.9161415*y + 0.0334540*z)
	b = (0.0719453*x + -0.2289914*y + 1.4052427*z)
	return
}
