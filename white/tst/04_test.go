// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package white

import (
	"fmt"
	"testing"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-5 // accuracy
	var d float64
	if x >= y {
		d = x - y
	} else {
		d = y - x
	}
	if d > acc {
		return false
	}
	return true
}

func TestSD(t *testing.T) {
	sd0 := 0.0 // waiting for known values
	sd := SD(300, 0.5, 0.4)
	//	sd := SD(830)

	fmt.Println("test of SD")
	if !(check_diff(sd, sd0)) {
		t.Errorf("sd : %f ", sd)

	}
	s0, s1, s2 := s(300)
	fmt.Println("test of s")
	fmt.Println("s 300 :", s0, s1, s2)
	s0, s1, s2 = s(830)
	fmt.Println("s 830 :", s0, s1, s2)

	return
}
