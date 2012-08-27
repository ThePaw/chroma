// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Color delta / comparison math. 
// Euclidean color distance in sRGB

package delta

import "math"

//Euclidean Delta in sRGB
func EuclidRgb(r1, g1, b1, r2, g2, b2 float64) (delta float64) {
	rd := r2 - r1
	gd := g2 - g1
	bd := b2 - b1
	delta = math.Sqrt(rd*rd + gd*gd + bd*bd)
	return
}
