// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Converts from Munsell  HVC with D65 illuminator to CIE XYZ according to 1967 renotation data [34]. 
// XYZ values will be in the nominal range [0,1]. 
// Ref.: [32][33] 

package munsell

// MunsellToXYZ converts from Munsell  HVC with D65 illuminator to CIE XYZ. 
func MunsellToXYZ(h string, v, c float64) (x, y, z float64) {
	var (
		err    bool
		i, idx int
		hue    = [40]string{"2.5GY", "5GY", "7.5GY", "10GY", "2.5G", "5G", "7.5G", "10G", "2.5BG", "5BG", "7.5BG", "10BG", "2.5B", "5B", "7.5B", "10B", "2.5PB", "5PB", "7.5PB", "10PB", "2.5P", "5P", "7.5P", "10P", "2.5RP", "5RP", "7.5RP", "10RP", "2.5R", "5R", "7.5R", "10R", "2.5YR", "5YR", "7.5YR", "10YR", "2.5Y", "5Y", "7.5Y", "10Y"}
	)

	// parse hue given as string into an index for LUT
	err = true
	for idx = 0; idx < len(hue); idx++ {
		if hue[idx] == h {
			err = false
			break
		}
	}
	if err == true {
		panic("Illegal hue")
	}

	// look up the color number (index)
	err = true
	for i = 0; i < len(hvc); i++ {
		if float64(idx) == hvc[i][0] && v == hvc[i][1] && c == hvc[i][2] {
			err = false
			break
		}
	}
	if err == true {
		panic("Color not found")
	}

	x, y, z = xyz[i][0], xyz[i][1], xyz[i][2]
	return
}
