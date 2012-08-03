// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

import "math"

// Convert an sRGB color to the cylindrical Hue-Saturation-Intensity (HSI).
// The sRGB values are assumed to be between 0 and 1.  
// The outputs are:
//   H = polar hue angle         (0 <= H < 360),
//   S = 1 - min(R',G',B')/I     (0 <= S <= 1),
//   I = (R'+G'+B')/3            (0 <= I <= 1).
func RGBToHSI(r, g, b float64) (h, s, i float64) {
	alpha := 0.5 * (2*r - g - b)
	beta := 0.866025403784439 * (g - b)

	i = (r + g + b) / 3

	if i > 0 {
		s = 1 - min3(r, g, b)/i
		h = math.Atan2(beta, alpha) * (180 / math.Pi)

		if h < 0 {
			h += 360
		}
	} else {
		h = 0
		s = 0
	}
	return
}

// Convert a Hue-Saturation-Intesity (HSI) color to sRGB. 
//    0 <= H < 360,
//    0 <= S <= 1,
//    0 <= I <= 1. 
// The output sRGB values are scaled between 0 and 1.
func HSIToRGB(h, s, i float64) (r, g, b float64) {
	h -= 360 * math.Floor(h/360)

	if h < 120 {
		b = i * (1 - s)
		r = i * (1 + s*math.Cos(h*(math.Pi/180))/math.Cos((60-h)*(math.Pi/180)))
		g = 3*i - r - b
	} else if h < 240 {
		h -= 120
		r = i * (1 - s)
		g = i * (1 + s*math.Cos(h*(math.Pi/180))/math.Cos((60-h)*(math.Pi/180)))
		b = 3*i - r - g
	} else {
		h -= 240
		g = i * (1 - s)
		b = i * (1 + s*math.Cos(h*(math.Pi/180))/math.Cos((60-h)*(math.Pi/180)))
		r = 3*i - g - b
	}
	return
}
