// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

// Auxilliary functions

// Clip 3 channels to 0..255
func clip3(x, y, z *int) {
	if *x < 0 {
		*x = 0
	} else if *x > 255 {
		*x = 255
	}
	if *y < 0 {
		*y = 0
	} else if *y > 255 {
		*y = 255
	}
	if *z < 0 {
		*z = 0
	} else if *z > 255 {
		*z = 255
	}
	return
}

// Clip 4 channels to 0..255
func clip4(x, y, z, a *int) {
	if *x < 0 {
		*x = 0
	} else if *x > 255 {
		*x = 255
	}
	if *y < 0 {
		*y = 0
	} else if *y > 255 {
		*y = 255
	}
	if *z < 0 {
		*z = 0
	} else if *z > 255 {
		*z = 255
	}
	if *a < 0 {
		*a = 0
	} else if *a > 255 {
		*a = 255
	}
	return
}


