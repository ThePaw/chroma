// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import "math"

// Clip values to fit in 0..maxVal. 
func clip(x, maxVal float64) float64 {
	if x < 0 {
		x = 0
	} else if x > maxVal {
		x = maxVal
	}
	return x
}

// Round to nearest integer
func round(x float64) float64 {
	var i float64
	f := math.Floor(x)
	c := math.Ceil(x)
	if x-f < c-x {
		i = f
	} else {
		i = c
	}
	return i
}

// Minimum of a, b, and c
func min3(a, b, c float64) (y float64) {
	if a <= b {
		y = math.Min(a, c)
	} else {
		y = math.Min(b, c)
	}
	return
}

// Maximum of a, b, and c
func max3(a, b, c float64) (y float64) {
	if a >= b {
		y = math.Max(a, c)
	} else {
		y = math.Max(b, c)
	}
	return
}
