// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

// Rgb2Yiq converts from sRGB to NTSC YIQ
func Rgb2Yiq(r, g, b float64) (y, i, q float64) {
	y = 0.299*r + 0.587*g + 0.114*b
	i = 0.595716*r - 0.274453*g - 0.321263*b
	q = 0.211456*r - 0.522591*g + 0.311135*b
	return
}

//Yiq2Rgb converts from NTSC YIQ to sRGB */
func Yiq2Rgb(y, i, q float64) (r, g, b float64) {
	r = y + 0.9562957197589482261*i + 0.6210244164652610754*q
	g = y - 0.2721220993185104464*i - 0.6473805968256950427*q
	b = y - 1.1069890167364901945*i + 1.7046149983646481374*q
	return
}
