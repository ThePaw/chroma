// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Color delta / comparison math. 
// Riemersma, T. 2012: CompuPhase: Colour metric. http://www.compuphase.com/cmetric.htm

package delta

import (
	. "image/color"
	"math"
)

// Riemersma Colour Distance function
// The algorithm (used by the products EGI, AniSprite and PaletteMaker) is a combination both weighted Euclidean distance functions, 
// where the weight factors depend on how big the "red" component of the colour is. 
// First one calculates the mean level of "red" and then weights the ΔR′ and ΔB′ signals as a function of the mean red level. 
func ColourDistance(e1, e2 NRGBA) float64 {
	rMean := uint32(e1.R) + uint32(e2.R)/2
	r := uint32(e1.R) - uint32(e2.R)
	g := uint32(e1.G) - uint32(e2.G)
	b := uint32(e1.B) - uint32(e2.B)
	d := (((512 + rMean) * r * r) >> 8) + 4*g*g + (((767 - rMean) * b * b) >> 8)
	return math.Sqrt(float64(d))
}
