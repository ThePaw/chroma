// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

// XYZ to CIE-CAT02 LMS */
func XYZToLMS(x, y, z float64) (l, m, s float64) {
	l = 0.7328*x + 0.4296*y - 0.1624*z
	m = -0.7036*x + 1.6975*y + 0.0061*z
	s = 0.0030*x + 0.0136*y + 0.9834*z
	return
}

// CIE-CAT02 LMS to XYZ */
func LMSToXYZ(l, m, s float64) (x, y, z float64) {
	x = 1.096123820835514*l - 0.278869000218287*m + 0.182745179382773*s
	y = 0.454369041975359*l + 0.473533154307412*m + 0.072097803717229*s
	z = -0.009627608738429*l - 0.005698031216113*m + 1.015325639954543*s
	return
}
