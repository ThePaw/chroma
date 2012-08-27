// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Color delta / comparison math. 
// Functions to calculate color differences. All of them refers to the CIE-L*ab color space.
// http://www.easyrgb.com/index.php?X=DELT

package delta

import "math"

// Returns H° value
func cieLab2Hue(a, b float64) (h float64) {
	rad2deg := 180 / math.Pi
	bias := 0.
	notAssigned := 999999.0
	switch {
	case a >= 0 && b == 0:
		h = 0
	case a < 0 && b == 0:
		h = 180
	case a == 0 && b > 0:
		h = 90
	case a == 0 && b < 0:
		h = 270
	case a > 0 && b > 0:
		bias = 0
	case a < 0:
		bias = 180
	case a > 0 && b < 0:
		bias = 360
	}
	if h == notAssigned {
		h = rad2deg*math.Atan(b/a) + bias
	}
	return
}

// Delta C*
func DeltaC(a1, b1, a2, b2 float64) (delta float64) {
	delta = math.Sqrt(a2*a2+b2*b2) - math.Sqrt(a1*a1+b1*b1)
	return
}

// Delta H*
func DeltaH(a1, b1, a2, b2 float64) (delta float64) {
	xDE := math.Sqrt(a2*a2+b2*b2) - math.Sqrt(a1*a1+b1*b1)
	delta = math.Sqrt((a2-a1)*(a2-a1) + (b2-b1)*(b2-b1) - (xDE * xDE))
	return
}

// Delta E*
func DeltaE(l1, a1, b1, l2, a2, b2 float64) (delta float64) {
	delta = math.Sqrt(((l1 - l2) * (l1 - l2)) + ((a1 - a2) * (a1 - a2)) + ((b1 - b2) * (b1 - b2)))
	return
}

// Delta E 1994
func DeltaE94(l1, a1, b1, l2, a2, b2, wL, wC, wH float64) (delta float64) {
	// wL, wC, wH	weighting factors depending on the application (1 = default)
	var xDH float64
	xC1 := math.Sqrt((a1 * a1) + (b1 * b1))
	xC2 := math.Sqrt((a2 * a2) + (b2 * b2))
	xDL := l2 - l1
	xDC := xC2 - xC1
	xDE := math.Sqrt(((l1 - l2) * (l1 - l2)) + ((a1 - a2) * (a1 - a2)) + ((b1 - b2) * (b1 - b2)))
	if math.Sqrt(xDE) > math.Sqrt(math.Abs(xDL))+math.Sqrt(math.Abs(xDC)) {
		xDH = math.Sqrt((xDE * xDE) - (xDL * xDL) - (xDC * xDC))
	} else {
		xDH = 0
	}
	xSC := 1 + (0.045 * xC1)
	xSH := 1 + (0.015 * xC1)
	xDL /= wL
	xDC /= wC * xSC
	xDH /= wH * xSH
	delta = math.Sqrt(xDL*xDL + xDC*xDC + xDH*xDH)
	return
}

/* bug
// Delta E 2000
func DeltaE00(l1, a1, b1, l2, a2, b2, wL, wC, wH float64) (delta float64) {
	// l1, a1, b1          //Color #1 lab values
	// l2, a2, b2          //Color #2 lab values
	// wL, wC, wH                //Wheight factor
	var xDH, xHX float64
	const deg2rad = math.Pi / 180

	xC1 := math.Sqrt(a1*a1 + b1*b1)
	xC2 := math.Sqrt(a2*a2 + b2*b2)
	xCX := (xC1 + xC2) / 2
	xCX7 := xCX * xCX * xCX * xCX * xCX * xCX * xCX
	xGX := 0.5 * (1 - math.Sqrt((xCX7)/((xCX7)+(6103515625))))
	xNN := (1 + xGX) * a1
	xC1 = math.Sqrt(xNN*xNN + b1*b1)
	xH1 := cieLab2Hue(xNN, b1)
	xNN = (1 + xGX) * a2
	xC2 = math.Sqrt(xNN*xNN + b2*b2)
	xH2 := cieLab2Hue(xNN, b2)
	xDL := l2 - l1
	xDC := xC2 - xC1
	if xC1*xC2 == 0 {
		xDH = 0
	} else {
		//		xNN = round(xH2-xH1, 12)
		xNN = xH2 - xH1
		if math.Abs(xNN) <= 180 {
			xDH = xH2 - xH1
		} else {
			if xNN > 180 {
				xDH = xH2 - xH1 - 360
			} else {
				xDH = xH2 - xH1 + 360
			}
		}
	}
	xDH = 2 * math.Sqrt(xC1*xC2) * math.Sin(deg2rad*(xDH/2))
	xLX := (l1 + l2) / 2
	xCY := (xC1 + xC2) / 2
	if (xC1 * xC2) == 0 {
		xHX = xH1 + xH2
	} else {
		//		xNN = math.Abs(round(xH1-xH2, 12))
		xNN = math.Abs(xH1 - xH2)
		if xNN > 180 {
			if xH2+xH1 < 360 {
				xHX = xH1 + xH2 + 360
			} else {
				xHX = xH1 + xH2 - 360
			}
		} else {
			xHX = xH1 + xH2
		}
		xHX /= 2
	}
	xTX := 1 - 0.17*math.Cos(deg2rad*(xHX-30)) + 0.24*math.Cos(deg2rad*(2*xHX)) + 0.32*math.Cos(deg2rad*(3*xHX+6)) - 0.20*math.Cos(deg2rad*(4*xHX-63))
	xPH := 30 * math.Exp(-((xHX-275)/25)*((xHX-275)/25))
	xCY7 := xCY * xCY * xCY * xCY * xCY * xCY * xCY
	xRC := 2 * math.Sqrt((xCY7)/((xCY7)+(6103515625)))
	xSL := 1 + ((0.015 * ((xLX - 50) * (xLX - 50))) / math.Sqrt(20+((xLX-50)*(xLX-50))))
	xSC := 1 + 0.045*xCY
	xSH := 1 + 0.015*xCY*xTX
	xRT := -math.Sin(deg2rad*(2*xPH)) * xRC
	xDL /= wL * xSL
	xDC /= wC * xSC
	xDH /= wH * xSH
	delta = math.Sqrt(xDL*xDL + xDC*xDC + xDH*xDH + xRT*xDC*xDH)
	return
}

func DeltaCMC(l1, a1, b1, l2, a2, b2, wL, wC, wH float64) (delta float64) {

	//l1, a1, b1	Color #1 lab values
	//l2, a2, b2	Color #2 lab values
	//wL, wC	Weight factors
	const deg2rad = math.Pi / 180
	var xTT, xSL float64

	xC1 := math.Sqrt((a1 * a1) + (b1 * b1))
	xC2 := math.Sqrt((a2 * a2) + (b2 * b2))
	xC14 := xC1 * xC1 * xC1 * xC1
	xff := math.Sqrt((xC14) / ((xC14) + 1900))
	xH1 := cieLab2Hue(a1, b1)

	if xH1 < 164 || xH1 > 345 {
		xTT = 0.36 + math.Abs(0.4*math.Cos(deg2rad*(35+xH1)))
	} else {
		xTT = 0.56 + math.Abs(0.2*math.Cos(deg2rad*(168+xH1)))
	}

	if l1 < 16 {

		xSL = 0.511
	} else {
		xSL = (0.040975 * l1) / (1 + (0.01765 * l1))
	}
	xSC := ((0.0638 * xC1) / (1 + (0.0131 * xC1))) + 0.638
	xSH := ((xff * xTT) + 1 - xff) * xSC
	aa2 := (a2 - a1) * (a2 - a1)
	bb2 := (b2 - b1) * (b2 - b1)
	cc2 := (xC2 - xC1) * (xC2 - xC1)
	xDH := math.Sqrt(aa2 + bb2 - cc2)
	xSL = (l2 - l1) / wL * xSL
	xSC = (xC2 - xC1) / wC * xSC
	xSH = xDH / xSH
	delta = math.Sqrt(xSL*xSL + xSC*xSC + xSH*xSH)

	return
}
*/
