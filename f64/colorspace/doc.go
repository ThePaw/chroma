// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Colorspace conversion functions.
package colorspace

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
