// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

// A illuminant conversion functions
// A_B functions
func A_B_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8905163, -0.0829136, 0.2680945},
		{-0.0971524, 1.0754262, 0.0879463},
		{0.0538970, -0.0908558, 2.4838553}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_B_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9574884, -0.1643613, 0.2902356},
		{-0.0180539, 1.0185379, 0.0036373},
		{0.0000000, 0.0000000, 2.3949136}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_B_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9018844, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 2.3949136}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_C functions
func A_C_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8530161, -0.1130268, 0.4404346},
		{-0.1238786, 1.0853435, 0.1425803},
		{0.0911907, -0.1553658, 3.4776250}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_C_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9418277, -0.2249131, 0.4806950},
		{-0.0247051, 1.0253682, 0.0049749},
		{0.0000000, 0.0000000, 3.3225235}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_C_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8927993, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 3.3225235}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_D50 functions
func A_D50_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8779529, -0.0915288, 0.2566181},
		{-0.1117372, 1.0924325, 0.0851788},
		{0.0502012, -0.0837636, 2.3994031}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D50_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9531871, -0.1809896, 0.2757725},
		{-0.0198804, 1.0204129, 0.0040065},
		{0.0000000, 0.0000000, 2.3189827}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D50_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8777606, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 2.3189827}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_D55 functions
func A_D55_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8644459, -0.1021330, 0.3073182},
		{-0.1222890, 1.0982532, 0.1013945},
		{0.0609732, -0.1022820, 2.6887535}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D55_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9477017, -0.2021983, 0.3315105},
		{-0.0222101, 1.0228052, 0.0044754},
		{0.0000000, 0.0000000, 2.5895462}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D55_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8710241, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 2.5895462}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_D65 functions
func A_D65_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8446965, -0.1179225, 0.3948108},
		{-0.1366303, 1.1041226, 0.1291718},
		{0.0798489, -0.1348999, 3.1924009}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D65_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9394987, -0.2339150, 0.4281177},
		{-0.0256939, 1.0263828, 0.0051761},
		{0.0000000, 0.0000000, 3.0598005}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D65_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8652435, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 3.0598005}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_D75 functions
func A_D75_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8310208, -0.1290882, 0.4662974},
		{-0.1456744, 1.1060318, 0.1517254},
		{0.0954696, -0.1620076, 3.6068968}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D75_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9336695, -0.2564537, 0.5073423},
		{-0.0281696, 1.0289253, 0.0056737},
		{0.0000000, 0.0000000, 3.4463397}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_D75_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8645608, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 3.4463397}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_E functions
func A_E_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8815963, -0.0908179, 0.3439213},
		{-0.1006757, 1.0708986, 0.1115462},
		{0.0709158, -0.1206464, 2.9302950}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_E_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9532875, -0.1806053, 0.3749305},
		{-0.0198382, 1.0203706, 0.0039951},
		{0.0000000, 0.0000000, 2.8101728}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_E_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9103323, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 2.8101728}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_F2 functions
func A_F2_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9083396, -0.0683719, 0.1754134},
		{-0.0853260, 1.0728419, 0.0587007},
		{0.0336526, -0.0557284, 1.9465816}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F2_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9650788, -0.1350128, 0.1875333},
		{-0.0148302, 1.0152272, 0.0029893},
		{0.0000000, 0.0000000, 1.8938598}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F2_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9029222, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.8938598}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_F7 functions
func A_F7_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8447932, -0.1178395, 0.3941104},
		{-0.1365823, 1.1041477, 0.1289531},
		{0.0796929, -0.1346275, 3.1882950}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F7_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9395426, -0.2337454, 0.4273371},
		{-0.0256753, 1.0263637, 0.0051723},
		{0.0000000, 0.0000000, 3.0559786}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F7_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.8651889, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 3.0559786}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// A_F11 functions
func A_F11_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9214338, -0.0587653, 0.1579041},
		{-0.0725224, 1.0609434, 0.0526133},
		{0.0306110, -0.0508982, 1.8568836}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F11_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9699644, -0.1161244, 0.1692805},
		{-0.0127554, 1.0130970, 0.0025709},
		{0.0000000, 0.0000000, 1.8083462}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func A_F11_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9190897, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.8083462}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}
