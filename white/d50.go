// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

// D50 illuminant conversion functions
// D50_A functions
func D50_A_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1573713, 0.0872411, -0.1268788},
		{0.1199410, 0.9219445, -0.0455568},
		{-0.0200278, 0.0303599, 0.4178345}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_A_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0530074, 0.1867709, -0.1255459},
		{0.0205155, 0.9836342, -0.0041391},
		{0.0000000, 0.0000000, 0.4312236}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_A_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1392628, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.4312236}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_B functions
func D50_B_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0153440, 0.0093873, 0.0028087},
		{0.0147849, 0.9856776, 0.0000806},
		{0.0017354, -0.0036522, 1.0351412}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_B_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0048705, 0.0171595, 0.0056280},
		{0.0018849, 0.9984968, -0.0003808},
		{0.0000000, 0.0000000, 1.0327432}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_B_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0274834, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0327432}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_C functions
func D50_C_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9648789, -0.0164148, 0.0809482},
		{-0.0160520, 0.9941478, 0.0258478},
		{0.0172576, -0.0297025, 1.4485796}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_C_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9871374, -0.0453262, 0.0899753},
		{-0.0049788, 1.0039731, 0.0010028},
		{0.0000000, 0.0000000, 1.4327505}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_C_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0171330, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.4327505}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_D55 functions
func D50_D55_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9820801, -0.0094156, 0.0233811},
		{-0.0118389, 1.0049382, 0.0078491},
		{0.0044511, -0.0073485, 1.1203775}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D55_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9937887, -0.0218861, 0.0248120},
		{-0.0024040, 1.0019180, 0.0004847},
		{0.0000000, 0.0000000, 1.1166733}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D55_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9923254, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.1166733}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_D65 functions
func D50_D65_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9555766, -0.0230393, 0.0631636},
		{-0.0282895, 1.0099416, 0.0210077},
		{0.0122982, -0.0204830, 1.3299098}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D65_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9845002, -0.0546158, 0.0676324},
		{-0.0059992, 1.0047864, 0.0012095},
		{0.0000000, 0.0000000, 1.3194581}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D65_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9857398, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.3194581}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_D75 functions
func D50_D75_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9369777, -0.0323563, 0.0952771},
		{-0.0389795, 1.0115975, 0.0314918},
		{0.0188243, -0.0315280, 1.5023535}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D75_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9778997, -0.0778744, 0.1026211},
		{-0.0085539, 1.0068249, 0.0017244},
		{0.0000000, 0.0000000, 1.4861429}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_D75_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9849619, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.4861429}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_E functions
func D50_E_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0025535, 0.0036238, 0.0359837},
		{0.0096914, 0.9819125, 0.0105947},
		{0.0089181, -0.0160789, 1.2208770}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_E_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0001136, 0.0003968, 0.0427451},
		{0.0000436, 0.9999663, -0.0000100},
		{0.0000000, 0.0000000, 1.2118128}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_E_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0371077, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.2118128}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_F2 functions
func D50_F2_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0395725, 0.0215350, -0.0388405},
		{0.0287482, 0.9834389, -0.0135220},
		{-0.0067213, 0.0106555, 0.8116180}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F2_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0134653, 0.0474454, -0.0397341},
		{0.0052115, 0.9958424, -0.0010512},
		{0.0000000, 0.0000000, 0.8166770}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F2_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0286657, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.8166770}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_F7 functions
func D50_F7_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9557124, -0.0229756, 0.0628550},
		{-0.0282266, 1.0099623, 0.0209090},
		{0.0122325, -0.0203701, 1.3282016}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F7_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9845499, -0.0544408, 0.0672896},
		{-0.0059799, 1.0047711, 0.0012056},
		{0.0000000, 0.0000000, 1.3178100}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F7_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9856775, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.3178100}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// D50_F11 functions
func D50_F11_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0562303, 0.0310026, -0.0482555},
		{0.0422615, 0.9734013, -0.0171480},
		{-0.0078657, 0.0121201, 0.7743049}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F11_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0189974, 0.0669372, -0.0482967},
		{0.0073526, 0.9941346, -0.0014833},
		{0.0000000, 0.0000000, 0.7798015}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func D50_F11_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0470847, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.7798015}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}
