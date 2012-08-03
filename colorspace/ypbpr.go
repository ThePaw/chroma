// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// sRGB to Y'PbPr Luma (ITU-R BT.601) + Chroma */
func RGBToYPbPr(r, g, b float64) (y, pb, pr float64) {
	y = 0.299*r + 0.587*g + 0.114*b
	pb = -0.1687367*r - 0.331264*g + 0.5*b
	pr = 0.5*r - 0.418688*g - 0.081312*b
	return
}

// Y'PbPr to sRGB */
func YPbPrToRGB(y, pb, pr float64) (r, g, b float64) {
	r = 0.99999999999914679361*y - 1.2188941887145875e-06*pb + 1.4019995886561440468*pr
	g = 0.99999975910502514331*y - 0.34413567816504303521*pb - 0.71413649331646789076*pr
	b = 1.00000124040004623180*y + 1.77200006607230409200*pb + 2.1453384174593273e-06*pr
	return
}
