// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

// Rgb2Jycbcr converts sRGB to JPEG Y'CbCr Luma + Chroma. 
func Rgb2Jycbcr(r, g, b float64) (y, cb, cr float64) {
	y, cb, cr = Rgb2Ypbpr(r, g, b)
	cb += 0.5
	cr += 0.5
	return
}

// Jycbcr2Rgb converts JPEG Y'CbCr to sRGB. 
func Jycbcr2Rgb(y, cb, cr float64) (r, g, b float64) {
	cb -= 0.5
	cr -= 0.5
	r, g, b = Ypbpr2Rgb(y, cb, cr)
	return
}
