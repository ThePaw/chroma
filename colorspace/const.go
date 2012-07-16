// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colorspace

// All transformations assume a two degree observer angle and a D65 illuminant.
// The white point can be changed by modifying the wX, wY, wZ below.
const (
	// CIE XYZ color of the D65 white point
	wX = 0.950456
	wY = 1.0
	wZ = 1.088754

	// u'v' coordinates of the white point for CIE Lu*v*
	wU = (4 * wX) / (wX + 15*wY + 3*wZ)
	wV = (9 * wY) / (wX + 15*wY + 3*wZ)
)
