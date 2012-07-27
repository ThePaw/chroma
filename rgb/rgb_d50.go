// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Convert different RGB colorspaces with D50 illuminator to CIE XYZ and back. 

package rgb

// Converts from Adobe RGB (1998) with D50 illuminator to CIE XYZ. 
func Adobe2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.6097559, 0.2052401, 0.1492240},
		{0.3111242, 0.6256560, 0.0632197},
		{0.0194811, 0.0608902, 0.7448387}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to Adobe RGB(1998) with D50 illuminator. 
func Xyz2Adobe_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{1.9624274, -0.6105343, -0.3413404},
		{-0.9787684, 1.9161415, 0.0334540},
		{0.0286869, -0.1406752, 1.3487655}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from Apple RGB with D50 illuminator to CIE XYZ. 
func Apple2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4755678, 0.3396722, 0.1489800},
		{0.2551812, 0.6725693, 0.0722496},
		{0.0184697, 0.1133771, 0.6933632}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to Apple RGB with D50 illuminator. 
func Xyz2Apple_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{2.8510695, -1.3605261, -0.4708281},
		{-1.0927680, 2.0348871, 0.0227598},
		{0.1027403, -0.2964984, 1.4510659}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from Bruce RGB with D50 illuminator to CIE XYZ. 
func Bruce2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4941816, 0.3204834, 0.1495550},
		{0.2521531, 0.6844869, 0.0633600},
		{0.0157886, 0.0629304, 0.7464909}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to Bruce RGB with D50 illuminator. 
func Xyz2Bruce_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{2.6502856, -1.2014485, -0.4289936},
		{-0.9787684, 1.9161415, 0.0334540},
		{0.0264570, -0.1361227, 1.3458542}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from CIE RGB with D50 illuminator to CIE XYZ. 
func Cie2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4868870, 0.3062984, 0.1710347},
		{0.1746583, 0.8247541, 0.0005877},
		{-0.0012563, 0.0169832, 0.8094831}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to CIE RGB with D50 illuminator. 
func Xyz2Cie(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{2.3638081, -0.8676030, -0.4988161},
		{-0.5005940, 1.3962369, 0.1047562},
		{0.0141712, -0.0306400, 1.2323842}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from  NTSC RGB with D50 illuminator to CIE XYZ. 
func Ntsc2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.6343706, 0.1852204, 0.1446290},
		{0.3109496, 0.5915984, 0.0974520},
		{-0.0011817, 0.0555518, 0.7708399}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to NTSC RGB with D50 illuminator. 
func Xyz2Ntsc_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{1.8464881, -0.5521299, -0.2766458},
		{-0.9826630, 2.0044755, -0.0690396},
		{0.0736477, -0.1453020, 1.3018376}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from  PAL/SECAM RGB with D50 illuminator to CIE XYZ. 
func Pal2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4552773, 0.3675500, 0.1413926},
		{0.2323025, 0.7077956, 0.0599019},
		{0.0145457, 0.1049154, 0.7057489}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to PAL/SECAM RGB with D50 illuminator. 
func Xyz2Pal_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{2.9603944, -1.4678519, -0.4685105},
		{-0.9787684, 1.9161415, 0.0334540},
		{0.0844874, -0.2545973, 1.4216174}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from  SMPTE-C RGB with D50 illuminator to CIE XYZ. 
func SmpteC2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4163290, 0.3931464, 0.1547446},
		{0.2216999, 0.7032549, 0.0750452},
		{0.0136576, 0.0913604, 0.7201920}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to SMPTE-C RGB with D50 illuminator. 
func Xyz2SmpteC_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{3.3921940, -1.8264027, -0.5385522},
		{-1.0770996, 2.0213975, 0.0207989},
		{0.0723073, -0.2217902, 1.3960932}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}

// Converts from  sRGB with D50 illuminator to CIE XYZ. 
func Srgb2Xyz_D50(r, g, b float64) (x, y, z float64) {
	m := [3][3]float64{
		{0.4360747, 0.3850649, 0.1430804},
		{0.2225045, 0.7168786, 0.0606169},
		{0.0139322, 0.0971045, 0.7141733}}
	x = m[0][0]*r + m[0][1]*g + m[0][2]*b
	y = m[1][0]*r + m[1][1]*g + m[1][2]*b
	z = m[2][0]*r + m[2][1]*g + m[2][2]*b
	return
}

// Converts from CIE XYZ to sRGB with D50 illuminator. 
func Xyz2Srgb_D50(x, y, z float64) (r, g, b float64) {
	m := [3][3]float64{
		{3.1338561, -1.6168667, -0.4906146},
		{-0.9787684, 1.9161415, 0.0334540},
		{0.0719453, -0.2289914, 1.4052427}}
	r = m[0][0]*x + m[0][1]*y + m[0][2]*z
	g = m[1][0]*x + m[1][1]*y + m[1][2]*z
	b = m[2][0]*x + m[2][1]*y + m[2][2]*z
	return
}
