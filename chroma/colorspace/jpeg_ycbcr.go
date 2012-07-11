// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

// Convert sRGB to JPEG Y'CbCr Luma + Chroma. 
func RGB2JpegYCbCr(r, g, b float64) (y, cb, cr float64) {
	y, cb, cr = RGB2YPbPr(r, g, b)
	cb += 0.5
	cr += 0.5
	return
}

// Convert JPEG Y'CbCr to sRGB. 
func JpegYCbCr2RGB(y, cb, cr float64) (r, g, b float64) {
	cb -= 0.5
	cr -= 0.5
	r, g, b = YPbPr2RGB(y, cb, cr)
	return
}
