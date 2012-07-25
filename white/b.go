// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

/*

 s/ /, /g 	

remove  func B_B...()  first!

m := [3][3]float64{
		{xx},
		{xx},
		{xx}}
*/

// B illuminant conversion functions
// B_A functions
func B_A_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1389718, 0.0771960, -0.1256680},
		{0.1046012, 0.9341802, -0.0443668},
		{-0.0208883, 0.0324959, 0.4037039}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_A_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0475866, 0.1690489, -0.1272120},
		{0.0185688, 0.9847959, -0.0037460},
		{0.0000000, 0.0000000, 0.4175516}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_A_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1087896, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.4175516}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_C functions
func B_C_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9505386, -0.0254257, 0.0756230},
		{-0.0305443, 1.0089768, 0.0249746},
		{0.0149704, -0.0250917, 1.3993642}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_C_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9824697, -0.0622785, 0.0817457},
		{-0.0068409, 1.0056021, 0.0013790},
		{0.0000000, 0.0000000, 1.3873250}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_C_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9899265, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.3873250}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_D50 functions
func B_D50_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9850292, -0.0093910, -0.0026720},
		{-0.0147751, 1.0146711, -0.0000389},
		{-0.0017035, 0.0035957, 0.9660561}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D50_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9951852, -0.0171026, -0.0054296},
		{-0.0018786, 1.0015377, 0.0003795},
		{0.0000000, 0.0000000, 0.9682949}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D50_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9732518, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.9682949}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_D55 functions
func B_D55_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9674769, -0.0186924, 0.0199637},
		{-0.0265230, 1.0198211, 0.0075752},
		{0.0025844, -0.0034696, 1.0823359}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D55_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9890450, -0.0389162, 0.0186211},
		{-0.0042747, 1.0034998, 0.0008627},
		{0.0000000, 0.0000000, 1.0812691}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D55_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9657825, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0812691}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_D65 functions
func B_D65_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9415037, -0.0321240, 0.0584672},
		{-0.0428238, 1.0250998, 0.0203309},
		{0.0101511, -0.0161170, 1.2847354}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D65_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9798627, -0.0715373, 0.0601219},
		{-0.0078579, 1.0064341, 0.0015850},
		{0.0000000, 0.0000000, 1.2776246}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D65_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9593730, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.2776246}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_D75 functions
func B_D75_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9232662, -0.0412875, 0.0895406},
		{-0.0533960, 1.0269180, 0.0304876},
		{0.0164490, -0.0267652, 1.4513087}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D75_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9733376, -0.0947188, 0.0940283},
		{-0.0104042, 1.0085195, 0.0020982},
		{0.0000000, 0.0000000, 1.4390247}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_D75_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9586160, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.4390247}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_E functions
func B_E_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9874297, -0.0056086, 0.0320832},
		{-0.0049796, 0.9962653, 0.0101710},
		{0.0069424, -0.0120086, 1.1794125}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_E_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9952976, -0.0167072, 0.0359597},
		{-0.0018352, 1.0015032, 0.0003695},
		{0.0000000, 0.0000000, 1.1733922}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_E_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0093669, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.1733922}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_F2 functions
func B_F2_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0237573, 0.0119487, -0.0403007},
		{0.0138105, 0.9975484, -0.0131781},
		{-0.0081608, 0.0137933, 0.7840861}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F2_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0084966, 0.0301855, -0.0439590},
		{0.0033157, 0.9972846, -0.0006683},
		{0.0000000, 0.0000000, 0.7907842}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F2_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0011507, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.7907842}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_F7 functions
func B_F7_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9416371, -0.0320617, 0.0581686},
		{-0.0427619, 1.0251198, 0.0202354},
		{0.0100877, -0.0160079, 1.2830854}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F7_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9799118, -0.0713629, 0.0597898},
		{-0.0078387, 1.0064184, 0.0015812},
		{0.0000000, 0.0000000, 1.2760288}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F7_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9593124, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.2760288}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// B_F11 functions
func B_F11_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0400419, 0.0213649, -0.0494410},
		{0.0272759, 0.9872237, -0.0167167},
		{-0.0092461, 0.0151560, 0.7480426}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F11_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0139654, 0.0496126, -0.0522728},
		{0.0054496, 0.9955375, -0.0010989},
		{0.0000000, 0.0000000, 0.7550779}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func B_F11_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0190770, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.7550779}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}
