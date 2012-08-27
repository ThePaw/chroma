// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Delta E 2000
// Compute the CIEDE2000 color-difference according to [18]. 

package delta

import "math"

// DeltaE00Sharma computes the CIEDE2000 color-difference using the algorithm of Sharma. See [18]. 
func DeltaE00Sharma(l0, a0, b0, l1, a1, b1, kl, kc, kh float64) (delta float64) {
	var hp, hp0, hp1 float64
	rad2deg := 180 / math.Pi
	deg2rad := math.Pi / 180

	cab0 := math.Sqrt(a0*a0 + b0*b0) //(2)
	cab1 := math.Sqrt(a1*a1 + b1*b1)

	cMean := (cab0 + cab1) / 2 //(3)
	// 6103515625 == 25^7  
	g := 0.5 * (1 - math.Sqrt(math.Pow(cMean, 7)/(math.Pow(cMean, 7)+6103515625))) //(4)

	ap0 := (1 + g) * a0 //(5)
	ap1 := (1 + g) * a1

	cp0 := math.Sqrt(ap0*ap0 + b0*b0) //(6)
	cp1 := math.Sqrt(ap1*ap1 + b1*b1)
	cpProd := (cp0 * cp1)

	if ap0 == 0 && b0 == 0 { //(7)
		hp0 = 0
	} else {
		hp0 = math.Atan2(b0, ap0)
	}
	for hp0 >= 2*math.Pi || hp0 < 0 {
		if hp0 >= 2*math.Pi {
			hp0 -= 2 * math.Pi
		} else {
			hp0 += 2 * math.Pi
		}
	}

	if math.Abs(ap1) == 0 && math.Abs(b1) == 0 {
		hp1 = 0
	} else {
		hp1 = math.Atan2(b1, ap1)
	}
	for hp1 >= 2*math.Pi || hp1 < 0 {
		if hp1 >= 2*math.Pi {
			hp1 -= 2 * math.Pi
		} else {
			hp1 += 2 * math.Pi
		}
	}
	dL := (l1 - l0)   //(8)
	dC := (cp1 - cp0) //(9)

	// Compute hue difference
	dhp := hp1 - hp0
	switch { //(10) 
	case cpProd == 0:
		dhp = 0
	case cpProd != 0 && math.Abs(dhp) <= math.Pi:
		// do nothing
	case cpProd != 0 && dhp > math.Pi:
		dhp -= 2 * math.Pi
	case cpProd != 0 && dhp < -math.Pi:
		dhp += 2 * math.Pi
	default:
		// do nothing
	}
	for dhp > math.Pi || dhp <= -math.Pi {
		if dhp > math.Pi {
			dhp -= 2 * math.Pi
		} else {
			dhp += 2 * math.Pi
		}
	}

	// Note that the defining equations actually need
	// signed Hue and chroma differences which is different
	// from prior color difference formulae
	dH := 2 * math.Sqrt(cpProd) * math.Sin(dhp/2) //(11)

	// weighting functions
	lp := (l1 + l0) / 2   //(12)
	cp := (cp0 + cp1) / 2 //(13)

	// Average Hue Computation
	// Note that average hue is computed in radians and converted to degrees only where needed
	hp = (hp0 + hp1) / 2
	switch { //(14)
	case cpProd != 0 && math.Abs(hp0-hp1) <= math.Pi:
		hp = (hp0 + hp1) / 2
	case cpProd != 0 && math.Abs(hp0-hp1) > math.Pi && (hp0+hp1 < 2*math.Pi):
		hp = (hp0 + hp1 + 2*math.Pi) / 2
	case cpProd != 0 && math.Abs(hp0-hp1) > math.Pi && (hp0+hp1 >= 2*math.Pi):
		hp = (hp0 + hp1 - 2*math.Pi) / 2
	case cpProd == 0:
		hp = hp0 + hp1
	default:
		hp = (hp0 + hp1) / 2
	}
	for hp >= 2*math.Pi || hp < 0 {
		if hp >= 2*math.Pi {
			hp -= 2 * math.Pi
		} else {
			hp += 2 * math.Pi

		}
	}

	t := 1 - 0.17*math.Cos(hp-math.Pi/6) + 0.24*math.Cos(2*hp) + 0.32*math.Cos(3*hp+math.Pi/30) - 0.20*math.Cos(4*hp-63*math.Pi/180) //(15)

	bb := ((rad2deg*hp - 275) / 25.0) //(16)
	dThetaRad := deg2rad * (30 * math.Exp(-1*bb*bb))

	cp7 := cp * cp * cp * cp * cp * cp * cp //(17)
	rc := 2 * math.Sqrt((cp7)/(cp7+6103515625))

	lpm502 := (lp - 50) * (lp - 50) //(18)
	sl := 1 + 0.015*lpm502/math.Sqrt(20+lpm502)
	sc := 1 + 0.045*cp   //(19)
	sh := 1 + 0.015*cp*t //(20)

	rt := -math.Sin(2*dThetaRad) * rc //(21)

	// The CIE 00 color difference
	klSl := kl * sl //(22)
	kcSc := kc * sc
	khSh := kh * sh

	dL2 := (dL / klSl) * (dL / klSl)
	dC2 := (dC / kcSc) * (dC / kcSc)
	dH2 := (dH / khSh) * (dH / khSh)
	delta = math.Sqrt(dL2 + dC2 + dH2 + rt*(dC/kcSc)*(dH/khSh))

	return
}
