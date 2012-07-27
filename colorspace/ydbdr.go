// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// Rgb2Ydbdr converts sRGB to SECAM Y'DbDr Luma + Chroma */
func Rgb2Ydbdr(r, g, b float64) (y, db, dr float64) {
	y = 0.299*r + 0.587*g + 0.114*b
	db = -0.450*r - 0.883*g + 1.333*b
	dr = -1.333*r + 1.116*g + 0.217*b
	return
}

// Ydbdr2Rgb converts SECAM Y'DbDr to sRGB */
func Ydbdr2Rgb(y, db, dr float64) (r, g, b float64) {
	r = y + 9.2303716147657e-05*db - 0.52591263066186533*dr
	g = y - 0.12913289889050927*db + 0.26789932820759876*dr
	b = y + 0.66467905997895482*db - 7.9202543533108e-05*dr
	return
}
