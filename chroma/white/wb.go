// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

// White balance functions

import (
	//	. "image/color"
	"math"
)

const (
	minTemp  = 2000.0
	maxTemp  = 25000.0
	minGreen = 0.02
	maxGreen = 5.0
)

func clipTemp(temp *float64) {
	if *temp < minTemp {
		*temp = minTemp
	} else if *temp > maxTemp {
		*temp = maxTemp
	}
	return
}

func clipGreen(green *float64) {
	if *green < minGreen {
		*green = minGreen
	} else if *green > maxGreen {
		*green = maxGreen
	}
	return
}

// Colour Rendering of Spectra
// These values are often called xBar yBar zBar and are characteristics of a color / illuminant.
// Values cie_colour_match[][3] = Observer 2° x2, y2, z2
// E.g. for 380nm: x2=0.001368  y2=0.000039  z2=0.00645
func spect2xyz(spec_intens func(λ, m1, m2, temp float64) float64, _m1, _m2, _temp float64) (x, y, z float64) {
	var λ, xX, yY, zZ, xyz float64
	xX = 0.0
	yY = 0.0

	//350nm to 830nm
	// should be const, not var
	var cie_colour_match = [97][3]float64{{0.0000000, 0.000000, 0.000000}, {0.0000000, 0.000000, 0.000000}, {0.0001299, 0.000392, 0.000606},
		{0.0002321, 0.000007, 0.001086}, {0.0004149, 0.000012, 0.001946}, {0.0007416, 0.000022, 0.003846},
		{0.001368, 0.000039, 0.006450}, {0.002236, 0.000064, 0.010550}, {0.004243, 0.000120, 0.020050},
		{0.007650, 0.000217, 0.036210}, {0.014310, 0.000396, 0.067850}, {0.023190, 0.000640, 0.110200},
		{0.043510, 0.001210, 0.207400}, {0.077630, 0.002180, 0.371300}, {0.134380, 0.004000, 0.645600},
		{0.214770, 0.007300, 1.039050}, {0.283900, 0.011600, 1.385600}, {0.328500, 0.016840, 1.622960},
		{0.348280, 0.023000, 1.747100}, {0.348060, 0.029800, 1.782600}, {0.336200, 0.038000, 1.772110},
		{0.318700, 0.048000, 1.744100}, {0.290800, 0.060000, 1.669200}, {0.251100, 0.073900, 1.528100},
		{0.195360, 0.090980, 1.287640}, {0.142100, 0.112600, 1.041900}, {0.095640, 0.139020, 0.812950},
		{0.057950, 0.169300, 0.616200}, {0.032010, 0.208020, 0.465180}, {0.014700, 0.258600, 0.353300},
		{0.004900, 0.323000, 0.272000}, {0.002400, 0.407300, 0.212300}, {0.009300, 0.503000, 0.158200},
		{0.029100, 0.608200, 0.111700}, {0.063270, 0.710000, 0.078250}, {0.109600, 0.793200, 0.057250},
		{0.165500, 0.862000, 0.042160}, {0.225750, 0.914850, 0.029840}, {0.290400, 0.954000, 0.020300},
		{0.359700, 0.980300, 0.013400}, {0.433450, 0.994950, 0.008750}, {0.512050, 1.000000, 0.005750},
		{0.594500, 0.995000, 0.003900}, {0.678400, 0.978600, 0.002750}, {0.762100, 0.952000, 0.002100},
		{0.842500, 0.915400, 0.001800}, {0.916300, 0.870000, 0.001650}, {0.978600, 0.816300, 0.001400},
		{1.026300, 0.757000, 0.001100}, {1.056700, 0.694900, 0.001000}, {1.062200, 0.631000, 0.000800},
		{1.045600, 0.566800, 0.000600}, {1.002600, 0.503000, 0.000340}, {0.938400, 0.441200, 0.000240},
		{0.854450, 0.381000, 0.000190}, {0.751400, 0.321000, 0.000100}, {0.642400, 0.265000, 0.000050},
		{0.541900, 0.217000, 0.000030}, {0.447900, 0.175000, 0.000020}, {0.360800, 0.138200, 0.000010},
		{0.283500, 0.107000, 0.000000}, {0.218700, 0.081600, 0.000000}, {0.164900, 0.061000, 0.000000},
		{0.121200, 0.044580, 0.000000}, {0.087400, 0.032000, 0.000000}, {0.063600, 0.023200, 0.000000},
		{0.046770, 0.017000, 0.000000}, {0.032900, 0.011920, 0.000000}, {0.022700, 0.008210, 0.000000},
		{0.015840, 0.005723, 0.000000}, {0.011359, 0.004102, 0.000000}, {0.008111, 0.002929, 0.000000},
		{0.005790, 0.002091, 0.000000}, {0.004109, 0.001484, 0.000000}, {0.002899, 0.001047, 0.000000},
		{0.002049, 0.000740, 0.000000}, {0.001440, 0.000520, 0.000000}, {0.001000, 0.000361, 0.000000},
		{0.000690, 0.000249, 0.000000}, {0.000476, 0.000172, 0.000000}, {0.000332, 0.000120, 0.000000},
		{0.000235, 0.000085, 0.000000}, {0.000166, 0.000060, 0.000000}, {0.000117, 0.000042, 0.000000},
		{0.000083, 0.000030, 0.000000}, {0.000059, 0.000021, 0.000000}, {0.000042, 0.000015, 0.000000},
		{0.000029, 0.000011, 0.000000}, {0.000021, 0.000007, 0.000000}, {0.000015, 0.000005, 0.000000},
		{0.000010, 0.000004, 0.000000}, {0.000007, 0.000003, 0.000000}, {0.000005, 0.000002, 0.000000},
		{0.000004, 0.000001, 0.000000}, {0.000003, 0.000001, 0.000000}, {0.000002, 0.000001, 0.000000},
		{0.000001, 0.000001, 0.000000}}

	λ = 350.0
	for i := 0; λ < 830.1; i++ {
		Me := spec_intens(λ, _m1, _m2, _temp)
		xX += Me * cie_colour_match[i][0]
		yY += Me * cie_colour_match[i][1]
		zZ += Me * cie_colour_match[i][2]
		λ += 5.0
	}
	xyz = (xX + yY + zZ)
	x = xX / xyz
	y = yY / xyz
	z = zZ / xyz

	return x, y, z
}

// XYZ for specially handled illumination methods
func illum2xyz(illum uint8) (x, y, z float64) {
	const (
		Daylight = iota
		Cloudy
		Shade
		Tungsten
		Fluo_F1
		Fluo_F2
		Fluo_F3
		Fluo_F4
		Fluo_F5
		Fluo_F6
		Fluo_F7
		Fluo_F8
		Fluo_F9
		Fluo_F10
		Fluo_F11
		Fluo_F12
		HMI_Lamp
		GTI_Lamp
		JudgeIII_Lamp
		Solux_Lamp_3500K
		Solux_Lamp_4100K
		Solux_Lamp_4700K
		NG_Solux_Lamp_4700K
		LED_LSI_Lumelex_2040
		LED_CRS_SP12_WWMR16
		Flash_5500K
		Flash_6000K
		Flash_6500K
		TempGreen
	)

	fn := daylight5300_spect //default

	switch illum {
	case Daylight:
		fn = daylight5300_spect
	case Cloudy:
		fn = cloudy6200_spect
	case Shade:
		fn = shade7600_spect
	case Tungsten:
		fn = tungsten2856_spect
	case Fluo_F1:
		fn = fluoF1_spect
	case Fluo_F2:
		fn = fluoF2_spect
	case Fluo_F3:
		fn = fluoF3_spect
	case Fluo_F4:
		fn = fluoF4_spect
	case Fluo_F5:
		fn = fluoF5_spect
	case Fluo_F6:
		fn = fluoF6_spect
	case Fluo_F7:
		fn = fluoF7_spect
	case Fluo_F8:
		fn = fluoF8_spect
	case Fluo_F9:
		fn = fluoF9_spect
	case Fluo_F10:
		fn = fluoF10_spect
	case Fluo_F11:
		fn = fluoF11_spect
	case Fluo_F12:
		fn = fluoF12_spect
	case HMI_Lamp:
		fn = hmi_spect
	case GTI_Lamp:
		fn = gti_spect
	case JudgeIII_Lamp:
		fn = judgeIII_spect
	case Solux_Lamp_3500K:
		fn = solux3500_spect
	case Solux_Lamp_4100K:
		fn = solux4100_spect
	case Solux_Lamp_4700K:
		fn = solux4700_spect
	case NG_Solux_Lamp_4700K:
		fn = ng_Solux4700_spect
	case LED_LSI_Lumelex_2040:
		fn = ng_LEDLSI2040_spect
	case LED_CRS_SP12_WWMR16:
		fn = ng_CRSSP12WWMR16_spect
	case Flash_5500K:
		fn = flash5500_spect
	case Flash_6000K:
		fn = flash6000_spect
	case Flash_6500K:
		fn = flash6500_spect
	}

	x, y, z = spect2xyz(fn, 0., 0., 0.)
	return
}

// XYZ for Planck's radiation of black body
func temp2xyz(temp float64) (x, y, z float64) {
	var xD, yD, interm float64

	if temp <= 4000 {
		// if temperature is between 2000K and 4000K we use blackbody, because there will be no Daylight reference below 4000K...
		x, y, z = spect2xyz(blackBodySpect, 0., 0., temp)
	} else {
		// from 4000K up to 25000K: using the D illuminant (daylight) which is standard
		if temp <= 7000 {
			xD = -4.6070e9/(temp*temp*temp) + 2.9678e6/(temp*temp) + 0.09911e3/temp + 0.244063
		} else {
			xD = -2.0064e9/(temp*temp*temp) + 1.9018e6/(temp*temp) + 0.24748e3/temp + 0.237040
		}
		yD = -3.0*xD*xD + 2.87*xD - 0.275
		// calculate D -daylight in function of s0, s1, s2 and temp ==> xD yD
		// S(λ)=So(λ)+m1*s1(λ)+m2*s2(λ)
		interm = (0.0241 + 0.2562*xD - 0.734*yD)
		m1 := (-1.3515 - 1.7703*xD + 5.9114*yD) / interm
		m2 := (0.03 - 31.4424*xD + 30.0717*yD) / interm
		fn := daylight_spect
		x, y, z = spect2xyz(fn, m1, m2, 0.)
	}
	return
}

// Spectral data for black body at given temperature
func blackBodySpect(λ, m1, m2, temp float64) float64 {
	wlm := λ * 1e-9        // Wavelength in meters
	c1 := 3.7417715247e-16 //3.7417..= c1 = 2*Pi*h*c2  where h=Planck constant, c=velocity of light
	c2 := 1.438786e-2      //1.4387..= c2 = h*c/k  where k=Boltzmann constant
	return (c1 * math.Pow(wlm, -5.0)) / (math.Exp(c2/(wlm*temp)) - 1.0)
}

// RGB channel multipliers
func xyz2mul(x, y, z, green float64, isRaw bool) (rmul, gmul, bmul float64) {
	var xD, yD float64

	clipGreen(&green)
	xD = x
	yD = y

	xX := xD / yD
	yY := 1.0
	zZ := (1.0 - xD - yD) / yD

	//recalculate channels multipliers with new values of XYZ due to whitebalance
	if isRaw {
		rmul = sRGB_xyz[0][0]*xX + sRGB_xyz[0][1]*yY + sRGB_xyz[0][2]*zZ
		gmul = sRGB_xyz[1][0]*xX + sRGB_xyz[1][1]*yY + sRGB_xyz[1][2]*zZ
		bmul = sRGB_xyz[2][0]*xX + sRGB_xyz[2][1]*yY + sRGB_xyz[2][2]*zZ
	} else {
		rmul = sRGBd65_xyz[0][0]*xX + sRGBd65_xyz[0][1]*yY + sRGBd65_xyz[0][2]*zZ
		gmul = sRGBd65_xyz[1][0]*xX + sRGBd65_xyz[1][1]*yY + sRGBd65_xyz[1][2]*zZ
		bmul = sRGBd65_xyz[2][0]*xX + sRGBd65_xyz[2][1]*yY + sRGBd65_xyz[2][2]*zZ
	}
	gmul /= green

	max := rmul
	if gmul > max {
		max = gmul
	}
	if bmul > max {
		max = bmul
	}
	rmul /= max
	gmul /= max
	bmul /= max
	return
}

// Temperature to RGB-multipliers
func temp2mul(temp, green float64) (rmul, gmul, bmul float64) {
	x, y, z := temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	return
}

// RGB-multipliers to temperature
func mul2temp(rmul, gmul, bmul float64) (temp, green float64) {
	var tmpr, tmpg, tmpb float64

	maxT := maxTemp
	minT := minTemp
	temp = (maxT + minT) / 2.0

	for maxT-minT > 1 {
		tmpr, tmpg, tmpb = temp2mul(temp, 1.0)
		if tmpb/tmpr > bmul/rmul {
			maxT = temp
		} else {
			minT = temp
		}
		temp = (maxT + minT) / 2.0
	}
	green = (tmpg / tmpr) / (gmul / rmul)
	clipTemp(&temp)
	clipGreen(&green)
	return
}

// ++pac
func mul2RGB(r, g, b, rmul, gmul, bmul float64) (rOut, gOut, bOut float64) {
	// compute channel multipliers
	rm := 1.0 / rmul
	gm := 1.0 / gmul
	bm := 1.0 / bmul

	mul_lum := 0.299*rm + 0.587*gm + 0.114*bm

	rm /= mul_lum
	gm /= mul_lum
	bm /= mul_lum

	// compute output color
	rOut = r * rm
	gOut = g * gm
	bOut = b * bm
	return
}

// Transform input sRGB color to output sRGB color corrected for color temperature of illuminator.
func TempWhiteBalance(r, g, b, temp float64) (rOut, gOut, bOut float64) {
	green := 1.0
	x, y, z := temp2xyz(temp)
	rmul, gmul, bmul := xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	return
}

// Transform input sRGB color to output sRGB color corrected for selected types of illuminators.
// See illum2xyz() for supported illuminators.
func IllumWhiteBalance(r, g, b float64, illum uint8) (rOut, gOut, bOut float64) {
	green := 1.0
	x, y, z := illum2xyz(illum)
	rmul, gmul, bmul := xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	return
}
