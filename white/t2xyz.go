// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package white

// T2Xy converts from the correlated color temperature of a CIE D-illuminant to the chromaticity of that D-illuminant. 
// Ref: [25]
// http://www.brucelindbloom.com/Eqn_T_to_xy.html
func T2Xy(t float64) (x, y float64) {

	switch {
	case t < 4000:
		panic("Too low a temperature, must be >= 4000K")
	case t >= 4000 && t < 7000:
		x = -4.6070e+9/(t*t*t) + 2.9678e+6/(t*t) + 0.09911e+3/t + 0.244063
	case t >= 7000 && t < 25000:
		x = -2.0064e+9/(t*t*t) + 1.9018e+6/(t*t) + 0.24748e+3/t + 0.237040
	case t > 25000:
		panic("Too high a temperature, must be <= 25000K")
	}
	y = -3.0*x*x + 2.870*x - 0.275
	return
}
