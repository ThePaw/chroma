// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package colorspace

import "math"

// Convert an sRGB color to Hue-Saturation-Lightness (HSL)
// The sRGB values are assumed to be between 0 and 1.  
// The outputs are:
// H = hexagonal hue angle                (0 <= H < 360),
// S = { C/(2L)     if L <= 1/2           (0 <= S <= 1),
// S = { C/(2 - 2L) if L >  1/2
// L = (max(R',G',B') + min(R',G',B'))/2  (0 <= L <= 1),
// where C = max(R',G',B') - min(R',G',B').  
func Rgb2Hsl(r, g, b float64) (h, s, l float64) {
	max := max3(r, g, b)
	min := min3(r, g, b)
	c := max - min

	l = (max + min) / 2

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
		//		s = (l <= 0.5) ? (c/(2*(l))) : (c/(2 - 2*(l)))
		if l <= 0.5 {
			s = c / (2 * l)
		} else {
			s = c / (2 - 2*l)
		}
	} else {
		h = 0
		s = 0
	}
	return
}

// Convert a Hue-Saturation-Lightness (HSL) color to sRGB
//    0 <= H < 360,
//    0 <= S <= 1,
//    0 <= L <= 1. 
// The output sRGB values are scaled between 0 and 1.
func Hsl2Rgb(h, s, l float64) (r, g, b float64) {
	var c float64
	if l <= 0.5 {
		c = 2 * l * s
	} else {
		c = (2 - 2*l) * s
	}
	min := l - 0.5*c
	h -= 360 * math.Floor(h/360)
	h /= 60
	x := c * (1 - math.Abs(h-2*math.Floor(h/2)-1))

	switch int(math.Floor(h)) {
	case 0:
		r = min + c
		g = min + x
		b = min
		break
	case 1:
		r = min + x
		g = min + c
		b = min
		break
	case 2:
		r = min
		g = min + c
		b = min + x
		break
	case 3:
		r = min
		g = min + x
		b = min + c
		break
	case 4:
		r = min + x
		g = min
		b = min + c
		break
	case 5:
		r = min + c
		g = min
		b = min + x
		break
	default:
		r = 0
		g = 0
		b = 0
	}
	return
}
