// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

// Rgb2Ycbcr converts sRGB to Y'CbCr Luma + Chroma. 
func Rgb2Ycbcr(r, g, b float64) (y, cb, cr float64) {
	y = 65.481*r + 128.553*g + 24.966*b + 16
	cb = -37.797*r - 74.203*g + 112.0*b + 128
	cr = 112.0*r - 93.786*g - 18.214*b + 128
	return
}

// Ycbcr2Rgb converts  Y'CbCr to sRGB. 
func Ycbcr2Rgb(y, cb, cr float64) (r, g, b float64) {
	y -= 16.0
	cb -= 128.0
	cr -= 128.0
	r = 0.00456621004566210107*y + 1.1808799897946415e-09*cr + 0.00625892896994393634*cb
	g = 0.00456621004566210107*y - 0.00153632368604490212*cr - 0.00318811094965570701*cb
	b = 0.00456621004566210107*y + 0.00791071623355474145*cr + 1.1977497040190077e-08*cb
	return
}
