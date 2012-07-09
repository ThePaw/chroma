// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chroma

import (
	"math"
)

// Hue-Saturation-Value (HSV) color to sRGB
// The input values are assumed to be scaled as 
//    0 <= H < 360,
//    0 <= S <= 1,
//    0 <= V <= 1. 
// The output sRGB values are scaled between 0 and 1.  This is the inverse
// transformation of Rgb2Hsv.
//
// Wikipedia: http://en.wikipedia.org/wiki/HSL_and_HSV
//
func HSVTosRGB(h, s, v float64) (r, g, b float64) {
	c := s * v
	min := v - c

	h -= 360 * math.Floor(h/360)
	h /= 60
	x := c * (1 - math.Abs(h-2*math.Floor(h/2)-1))

	switch int(h) {
	case 0:
		{
			r = min + c
			g = min + x
			b = min
		}
	case 1:
		{
			r = min + x
			g = min + c
			b = min
		}
	case 2:
		{
			r = min
			g = min + c
			b = min + x
		}
	case 3:
		{
			r = min
			g = min + x
			b = min + c
		}
	case 4:
		{
			r = min + x
			g = min
			b = min + c
		}
	case 5:
		{
			r = min + c
			g = min
			b = min + x
		}
	default:
		r = 0
		g = 0
		b = 0
	}
	return
}

// Converts an sRGB color to  hexcone Hue-Saturation-Value (HSV) color space. 
// The input sRGB values are scaled in [0,1]
// H = hexagonal hue angle   (0 <= H < 360),
// S = C/V                   (0 <= S <= 1),
// V = max(R',G',B')         (0 <= V <= 1),
// where C = max(R',G',B') - min(R',G',B'). 
// Wikipedia: http://en.wikipedia.org/wiki/HSL_and_HSV

func SRGBToHSV(r, g, b float64) (h, s, v float64) {
	max := max3(r, g, b)
	min := min3(r, g, b)
	c := max - min

	v = max

	if c > 0 {
		if max == r {
			h = (g - b) / c

			if g < b {
				h += 6
			}
		} else if max == g {
			h = 2 + (b-r)/c
		} else {
			h = 4 + (r-g)/c
		}
		h *= 60
		s = c / max
	} else {
		h = 0
		s = 0
	}
	return
}
