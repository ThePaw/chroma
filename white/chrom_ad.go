// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

// The chromatic adaptation algorithms on this web site may all be implemented as a linear transformation of a 
// source color (XS, YS, ZS) into a destination color (XD, YD, ZD) by a linear transformation [M] 
// which is dependent on the source reference white (XWS, YWS, ZWS) and the 
// destination reference white (XWD, YWD, ZWD). The idea behind all of these algorithms is to follow three steps:
// Transform from XYZ into a cone response domain, (ρ, γ, β).
// Scale the vector components by factors dependent upon both the source and destination reference whites.
// Transform from (ρ, γ, β) back to XYZ using the inverse transform of step 1.
// Thanks to Bruce Lindbloom
// http://www.brucelindbloom.com

// Bradford
func bradford() (ma, maInv [3][3]float64) {
	ma = [3][3]float64{{0.8951000, 0.2664000, -0.1614000},
		{-0.7502000, 1.7135000, 0.0367000},
		{0.0389000, -0.0685000, 1.0296000}}

	maInv = [3][3]float64{{0.9869929, -0.1470543, 0.1599627},
		{0.4323053, 0.5183603, 0.0492912},
		{-0.0085287, 0.0400428, 0.9684867}}
	return
}

// Von Kries
func vonKries() (ma, maInv [3][3]float64) {
	ma = [3][3]float64{{0.4002400, 0.7076000, -0.0808100},
		{-0.2263000, 1.1653200, 0.0457000},
		{0.0000000, 0.0000000, 0.9182200}}

	maInv = [3][3]float64{{1.8599364, -1.1293816, 0.2198974},
		{0.3611914, 0.6388125, -0.0000064},
		{0.0000000, 0.0000000, 1.0890636}}
	return
}

// XYZ Scaling
func xyzScaling() (ma, maInv [3][3]float64) {
	ma = [3][3]float64{{1.0000000, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0000000}}

	maInv = [3][3]float64{{1.0000000, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.0000000}}
	return
}

func ChromAd(xs, ys, zs, xws, yws, zws, xwd, ywd, zwd float64, method uint8) (xd, yd, zd float64) {

	var ma, maInv, m [3][3]float64

	switch method {
	case 0:
		{ // Bradford
			ma, maInv = bradford()
		}

	case 1:
		{
			ma, maInv = vonKries()
		}

	case 2:
		{
			ma, maInv = xyzScaling()
		}
	}

	// ρ, γ, β
	ρs := ma[0][0]*xws + ma[0][1]*yws + ma[0][2]*zws
	γs := ma[1][0]*xws + ma[1][1]*yws + ma[1][2]*zws
	βs := ma[2][0]*xws + ma[2][1]*yws + ma[2][2]*zws

	ρd := maInv[0][0]*xwd + maInv[0][1]*ywd + maInv[0][2]*zwd
	γd := maInv[1][0]*xwd + maInv[1][1]*ywd + maInv[1][2]*zwd
	βd := maInv[2][0]*xwd + maInv[2][1]*ywd + maInv[2][2]*zwd

	m[0][0] = maInv[0][0] * (ρd / ρs)
	m[0][1] = maInv[0][1] * (γd / γs)
	m[0][2] = maInv[0][2] * (βd / βs)
	m[1][0] = maInv[1][0] * (ρd / ρs)
	m[1][1] = maInv[1][1] * (γd / γs)
	m[1][2] = maInv[1][2] * (βd / βs)
	m[2][0] = maInv[2][0] * (ρd / ρs)
	m[2][1] = maInv[2][1] * (γd / γs)
	m[2][2] = maInv[2][2] * (βd / βs)

	m[0][0] = m[0][0]*ma[0][0] + m[0][1]*ma[1][0] + m[0][2]*ma[2][0]
	m[0][1] = m[0][0]*ma[0][1] + m[0][1]*ma[1][1] + m[0][2]*ma[2][1]
	m[0][2] = m[0][0]*ma[0][2] + m[0][1]*ma[1][2] + m[0][2]*ma[2][2]

	m[1][0] = m[1][0]*ma[0][0] + m[1][1]*ma[1][0] + m[1][2]*ma[2][0]
	m[1][1] = m[1][0]*ma[0][1] + m[1][1]*ma[1][1] + m[1][2]*ma[2][1]
	m[1][2] = m[1][0]*ma[0][2] + m[1][1]*ma[1][2] + m[1][2]*ma[2][2]

	m[2][0] = m[2][0]*ma[0][0] + m[2][1]*ma[1][0] + m[2][2]*ma[2][0]
	m[2][1] = m[2][0]*ma[0][1] + m[2][1]*ma[1][1] + m[2][2]*ma[2][1]
	m[2][2] = m[2][0]*ma[0][2] + m[2][1]*ma[1][2] + m[2][2]*ma[2][2]

	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs

	return
}
