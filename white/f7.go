// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// F7 illuminant conversion functions

package white

// F7_A functions
func F7_A_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.2162616, 0.1109265, -0.1548306},
		{0.1532455, 0.9152079, -0.0559592},
		{-0.0239302, 0.0358725, 0.3151544}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_A_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0710133, 0.2439140, -0.1501795},
		{0.0267922, 0.9804152, -0.0054059},
		{0.0000000, 0.0000000, 0.3272274}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_A_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1558170, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.3272274}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_B functions
func F7_B_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0639791, 0.0325159, -0.0487483},
		{0.0445370, 0.9766166, -0.0174212},
		{-0.0078094, 0.0119287, 0.7795373}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_B_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0210792, 0.0724024, -0.0479336},
		{0.0079529, 0.9941864, -0.0016046},
		{0.0000000, 0.0000000, 0.7836814}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_B_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0424133, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.7836814}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_C functions
func F7_C_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0096302, 0.0069785, 0.0130568},
		{0.0122432, 0.9846882, 0.0033800},
		{0.0038825, -0.0073255, 1.0905639}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_C_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0026841, 0.0092168, 0.0170692},
		{0.0010124, 0.9992607, -0.0002049},
		{0.0000000, 0.0000000, 1.0872208}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_C_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0319125, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0872208}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_D50 functions
func F7_D50_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0476532, 0.0228258, -0.0499378},
		{0.0294704, 0.9904638, -0.0169869},
		{-0.0091967, 0.0149801, 0.7530971}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D50_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0160269, 0.0550506, -0.0519304},
		{0.0060469, 0.9955792, -0.0012196},
		{0.0000000, 0.0000000, 0.7588347}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D50_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0145306, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.7588347}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_D55 functions
func F7_D55_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0283868, 0.0134412, -0.0312747},
		{0.0171406, 0.9952022, -0.0105684},
		{-0.0058572, 0.0096065, 0.8436556}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D55_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0095838, 0.0329194, -0.0327530},
		{0.0036160, 0.9973564, -0.0007292},
		{0.0000000, 0.0000000, 0.8473705}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D55_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0067445, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.8473705}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_D65 functions
func F7_D65_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9998530, -0.0000616, 0.0002403},
		{-0.0000675, 0.9999796, 0.0000778},
		{0.0000498, -0.0000848, 1.0012851}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D65_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9999485, -0.0001770, 0.0002629},
		{-0.0000194, 1.0000142, 0.0000039},
		{0.0000000, 0.0000000, 1.0012506}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D65_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0000631, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0012506}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_D75 functions
func F7_D75_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9797979, -0.0092332, 0.0255119},
		{-0.0113145, 1.0015327, 0.0084791},
		{0.0049755, -0.0082922, 1.1310137}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D75_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9931015, -0.0236961, 0.0271847},
		{-0.0026028, 1.0019031, 0.0005248},
		{0.0000000, 0.0000000, 1.1277369}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_D75_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9992740, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.1277369}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_E functions
func F7_E_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0501042, 0.0270124, -0.0230277},
		{0.0389931, 0.9729287, -0.0091847},
		{-0.0023588, 0.0025668, 0.9192667}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_E_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0161448, 0.0554519, -0.0195004},
		{0.0060910, 0.9955480, -0.0012294},
		{0.0000000, 0.0000000, 0.9195656}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_E_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0521775, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.9195656}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_F2 functions
func F7_F2_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0901032, 0.0444769, -0.0815305},
		{0.0592248, 0.9745143, -0.0283245},
		{-0.0141918, 0.0225586, 0.6113819}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_F2_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0299949, 0.1030276, -0.0828392},
		{0.0113169, 0.9917269, -0.0022829},
		{0.0000000, 0.0000000, 0.6197228}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_F2_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0436128, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.6197228}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F7_F11 functions
func F7_F11_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1079204, 0.0540934, -0.0896135},
		{0.0731195, 0.9648265, -0.0315596},
		{-0.0150044, 0.0234241, 0.5833138}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_F11_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0357335, 0.1227378, -0.0896478},
		{0.0134819, 0.9901444, -0.0027199},
		{0.0000000, 0.0000000, 0.5917405}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F7_F11_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0622994, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.5917405}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}
