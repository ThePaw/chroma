// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// RGBToJPEG_YCbCr converts sRGB to JPEG Y'CbCr Luma + Chroma. 
func RGBToJPEG_YCbCr(r, g, b float64) (y, cb, cr float64) {
	y, cb, cr = RGBToYPbPr(r, g, b)
	cb += 0.5
	cr += 0.5
	return
}

// JPEG_YCbCrToRGB converts JPEG Y'CbCr to sRGB. 
func JPEG_YCbCrToRGB(y, cb, cr float64) (r, g, b float64) {
	cb -= 0.5
	cr -= 0.5
	r, g, b = YPbPrToRGB(y, cb, cr)
	return
}
