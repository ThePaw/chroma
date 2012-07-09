package chroma

import "math"

// All transformations assume a two degree observer angle and a D65 illuminant.
// The white point can be changed by modifying the wX, wY, wZ below.
const (
	// CIE XYZ color of the D65 white point
	wX = 0.950456e+2
	wY = 1.0e+2
	wZ = 1.088754e+2

	// u'v' coordinates of the white point for CIE Lu*v*
	wU = (4 * wX) / (wX + 15*wY + 3*wZ)
	wV = (9 * wY) / (wX + 15*wY + 3*wZ)
)

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

// sRGB gamma correction, transforms RGB to R'G'B'
// http://en.wikipedia.org/wiki/SRGB
func gammaCorr(x float64) (y float64) {
	if x <= 0.0031306684425005883 {
		y = 12.92 * x
	} else {
		y = 1.055*math.Pow(x, 0.416666666666666667) - 0.055
	}
	return
}

// Inverse sRGB gamma correction, transforms R'G'B' to RGB 
func gammaCorrInv(x float64) (y float64) {
	if x <= 0.0404482362771076 {
		y = x / 12.92
	} else {
		y = math.Pow((x+0.055)/1.055, 2.4)
	}
	return
}
