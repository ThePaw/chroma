// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Colorspace conversion functions.
// All conversions assume 2 degree observer and D65 illuminant (illuminant can be changed in 'const.go'). 
// Supported colorspaces:
// 'RGB' 	sRGB IEC 61966-2-1
// 'YPbPr' 	Luma (ITU-R BT.601) + Chroma
// 'YCbCr' 	Luma + Chroma (digitized version of Y'PbPr)
// 'JPEG-YCbCr' 	Luma + Chroma space used in JFIF JPEG
// 'YUV' 	NTSC PAL Y'UV Luma + Chroma
// 'YIQ' 	NTSC Y'IQ Luma + Chroma
// 'YDbDr' 	SECAM Luma + Chroma
// 'HSV' or 'HSB'   	Hue Saturation Value/Brightness
// 'HSL' or 'HLS' 	Hue Saturation Luminance
// 'HSI' 	Hue Saturation Intensity
// 'XYZ' 	CIE XYZ
// 'Lab' 	CIE L*a*b* (CIELAB)
// 'Luv' 	CIE L*u*v* (CIELUV)
// 'LCH' 	CIE L*C*H* (CIELCH)
// 'CAT02 LMS' 	CIE CAT02 LMS
package colorspace

// Auxilliary functions. 

import "math"

// Clip values to fit in 0..maxVal. 
func clip(x, maxVal float64) float64 {
	if x < 0 {
		x = 0
	} else if x > maxVal {
		x = maxVal
	}
	return x
}

// Round to nearest integer
func round(x float64) float64 {
	var i float64
	f := math.Floor(x)
	c := math.Ceil(x)
	if x-f < c-x {
		i = f
	} else {
		i = c
	}
	return i
}

// Minimum of a, b, and c
func min3(a, b, c float64) (y float64) {
	if a <= b {
		y = math.Min(a, c)
	} else {
		y = math.Min(b, c)
	}
	return
}

// Maximum of a, b, and c
func max3(a, b, c float64) (y float64) {
	if a >= b {
		y = math.Max(a, c)
	} else {
		y = math.Max(b, c)
	}
	return
}
