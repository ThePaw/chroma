// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import "math"

// Convert RGB 0..1 to RGB 0..255. 
func RGBToRGB8(r, g, b float64) (r8, g8, b8 uint8) {
	r8 = uint8(clip(round(r), 255))
	g8 = uint8(clip(round(g), 255))
	b8 = uint8(clip(round(b), 255))
	return
}

// Gamma Correction
// CRT monitors have a nonlinear relationship between the input voltages and the rendered intensities. To reproduce an image accurately, the image is gamma-corrected in such a way that the monitor displays the desired intensities.
// Here, in chroma/colorspace, the “RGB” space is sRGB. The sRGB space was designed in 1996 for direct display on typical CRT monitors and standardized in 1999 by International Electrotechnical Commission (IEC) as IEC 61966-2-1. The “s” in sRGB is to mean “standard.” From linear RGB values, the sRGB values are gamma-compensated by the formula
// R' = 1.055 R1/2.4 − 0.055  	if R ≤ 0.0031308,
// R' = 12.92 R, 	if R > 0.0031308,
// and similarly for G' and B' [4].
// A standard notation is to denote R',G',B' quantities and derived quantities with a prime ' to signify gamma-correction. Since gamma-correction is already applied by digital cameras as standard practice, most digital image data should be interpreted as R'G'B', and not RGB.

// sRGB gamma correction, transforms RGB to R'G'B'
func gammaCorr(x float64) (y float64) {
	if x <= 0.0031306684425005883 {
		y = 12.92 * x
	} else {
		y = 1.055*math.Pow(x, 0.416666666666666667) - 0.055
	}
	return
}

// Inverse sRGB gamma correction, transforms R'G'B' to RGB 
func gammaCorrInv(x float64) (y float64) {
	if x <= 0.0404482362771076 {
		y = x / 12.92
	} else {
		y = math.Pow((x+0.055)/1.055, 2.4)
	}
	return
}
