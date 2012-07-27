// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package white

import (
	"fmt"
	"testing"
)

func check_diff(x, y float64) bool {
	const acc float64 = 1e-6// accuracy
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


func TestRoundTrip(t *testing.T) {
	var xd, yd, zd float64
	x, y, z := 0.500000, 0.026772, -0.797751
	fmt.Println("Roundtrip tests:")

// A_B
	xd, yd, zd = A_B_Bradford(x, y, z)
	xd, yd, zd = B_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_B_vonKries(x, y, z)
	xd, yd, zd = B_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_B_Xyz(x, y, z)
	xd, yd, zd = B_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_C
	xd, yd, zd = A_C_Bradford(x, y, z)
	xd, yd, zd = C_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_C_vonKries(x, y, z)
	xd, yd, zd = C_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_C_Xyz(x, y, z)
	xd, yd, zd = C_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_D50
	xd, yd, zd = A_D50_Bradford(x, y, z)
	xd, yd, zd = D50_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D50_vonKries(x, y, z)
	xd, yd, zd = D50_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D50_Xyz(x, y, z)
	xd, yd, zd = D50_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_D55
	xd, yd, zd = A_D55_Bradford(x, y, z)
	xd, yd, zd = D55_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D55_vonKries(x, y, z)
	xd, yd, zd = D55_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D55_Xyz(x, y, z)
	xd, yd, zd = D55_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_D65
	xd, yd, zd = A_D65_Bradford(x, y, z)
	xd, yd, zd = D65_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D65_vonKries(x, y, z)
	xd, yd, zd = D65_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D65_Xyz(x, y, z)
	xd, yd, zd = D65_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_D75
	xd, yd, zd = A_D75_Bradford(x, y, z)
	xd, yd, zd = D75_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D75_vonKries(x, y, z)
	xd, yd, zd = D75_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_D75_Xyz(x, y, z)
	xd, yd, zd = D75_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_E
	xd, yd, zd = A_E_Bradford(x, y, z)
	xd, yd, zd = E_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_E_vonKries(x, y, z)
	xd, yd, zd = E_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_E_Xyz(x, y, z)
	xd, yd, zd = E_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_F2
	xd, yd, zd = A_F2_Bradford(x, y, z)
	xd, yd, zd = F2_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F2_vonKries(x, y, z)
	xd, yd, zd = F2_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F2_Xyz(x, y, z)
	xd, yd, zd = F2_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_F7
	xd, yd, zd = A_F7_Bradford(x, y, z)
	xd, yd, zd = F7_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F7_vonKries(x, y, z)
	xd, yd, zd = F7_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F7_Xyz(x, y, z)
	xd, yd, zd = F7_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// A_F11
	xd, yd, zd = A_F11_Bradford(x, y, z)
	xd, yd, zd = F11_A_Bradford(xd, yd, zd)
	fmt.Println("test of A_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F11_vonKries(x, y, z)
	xd, yd, zd = F11_A_vonKries(xd, yd, zd)
	fmt.Println("test of A_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = A_F11_Xyz(x, y, z)
	xd, yd, zd = F11_A_Xyz(xd, yd, zd)
	fmt.Println("test of A_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

//B
// B_A
	xd, yd, zd = B_A_Bradford(x, y, z)
	xd, yd, zd = A_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_A_vonKries(x, y, z)
	xd, yd, zd = A_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_A_Xyz(x, y, z)
	xd, yd, zd = A_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_C
	xd, yd, zd = B_C_Bradford(x, y, z)
	xd, yd, zd = C_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_C_vonKries(x, y, z)
	xd, yd, zd = C_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_C_Xyz(x, y, z)
	xd, yd, zd = C_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_D50
	xd, yd, zd = B_D50_Bradford(x, y, z)
	xd, yd, zd = D50_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D50_vonKries(x, y, z)
	xd, yd, zd = D50_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D50_Xyz(x, y, z)
	xd, yd, zd = D50_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_D55
	xd, yd, zd = B_D55_Bradford(x, y, z)
	xd, yd, zd = D55_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D55_vonKries(x, y, z)
	xd, yd, zd = D55_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D55_Xyz(x, y, z)
	xd, yd, zd = D55_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_D65
	xd, yd, zd = B_D65_Bradford(x, y, z)
	xd, yd, zd = D65_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D65_vonKries(x, y, z)
	xd, yd, zd = D65_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D65_Xyz(x, y, z)
	xd, yd, zd = D65_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_D75
	xd, yd, zd = B_D75_Bradford(x, y, z)
	xd, yd, zd = D75_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D75_vonKries(x, y, z)
	xd, yd, zd = D75_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_D75_Xyz(x, y, z)
	xd, yd, zd = D75_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_E
	xd, yd, zd = B_E_Bradford(x, y, z)
	xd, yd, zd = E_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_E_vonKries(x, y, z)
	xd, yd, zd = E_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_E_Xyz(x, y, z)
	xd, yd, zd = E_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_F2
	xd, yd, zd = B_F2_Bradford(x, y, z)
	xd, yd, zd = F2_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F2_vonKries(x, y, z)
	xd, yd, zd = F2_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F2_Xyz(x, y, z)
	xd, yd, zd = F2_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_F7
	xd, yd, zd = B_F7_Bradford(x, y, z)
	xd, yd, zd = F7_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F7_vonKries(x, y, z)
	xd, yd, zd = F7_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F7_Xyz(x, y, z)
	xd, yd, zd = F7_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// B_F11
	xd, yd, zd = B_F11_Bradford(x, y, z)
	xd, yd, zd = F11_B_Bradford(xd, yd, zd)
	fmt.Println("test of B_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F11_vonKries(x, y, z)
	xd, yd, zd = F11_B_vonKries(xd, yd, zd)
	fmt.Println("test of B_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = B_F11_Xyz(x, y, z)
	xd, yd, zd = F11_B_Xyz(xd, yd, zd)
	fmt.Println("test of B_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C
// C_B
	xd, yd, zd = C_B_Bradford(x, y, z)
	xd, yd, zd = B_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_B_vonKries(x, y, z)
	xd, yd, zd = B_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_B_Xyz(x, y, z)
	xd, yd, zd = B_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_A
	xd, yd, zd = C_A_Bradford(x, y, z)
	xd, yd, zd = A_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_A_vonKries(x, y, z)
	xd, yd, zd = A_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_A_Xyz(x, y, z)
	xd, yd, zd = A_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_D50
	xd, yd, zd = C_D50_Bradford(x, y, z)
	xd, yd, zd = D50_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D50_vonKries(x, y, z)
	xd, yd, zd = D50_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D50_Xyz(x, y, z)
	xd, yd, zd = D50_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_D55
	xd, yd, zd = C_D55_Bradford(x, y, z)
	xd, yd, zd = D55_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D55_vonKries(x, y, z)
	xd, yd, zd = D55_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D55_Xyz(x, y, z)
	xd, yd, zd = D55_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_D65
	xd, yd, zd = C_D65_Bradford(x, y, z)
	xd, yd, zd = D65_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D65_vonKries(x, y, z)
	xd, yd, zd = D65_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D65_Xyz(x, y, z)
	xd, yd, zd = D65_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_D75
	xd, yd, zd = C_D75_Bradford(x, y, z)
	xd, yd, zd = D75_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D75_vonKries(x, y, z)
	xd, yd, zd = D75_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_D75_Xyz(x, y, z)
	xd, yd, zd = D75_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_E
	xd, yd, zd = C_E_Bradford(x, y, z)
	xd, yd, zd = E_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_E_vonKries(x, y, z)
	xd, yd, zd = E_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_E_Xyz(x, y, z)
	xd, yd, zd = E_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_F2
	xd, yd, zd = C_F2_Bradford(x, y, z)
	xd, yd, zd = F2_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F2_vonKries(x, y, z)
	xd, yd, zd = F2_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F2_Xyz(x, y, z)
	xd, yd, zd = F2_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_F7
	xd, yd, zd = C_F7_Bradford(x, y, z)
	xd, yd, zd = F7_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F7_vonKries(x, y, z)
	xd, yd, zd = F7_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F7_Xyz(x, y, z)
	xd, yd, zd = F7_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// C_F11
	xd, yd, zd = C_F11_Bradford(x, y, z)
	xd, yd, zd = F11_C_Bradford(xd, yd, zd)
	fmt.Println("test of C_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F11_vonKries(x, y, z)
	xd, yd, zd = F11_C_vonKries(xd, yd, zd)
	fmt.Println("test of C_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = C_F11_Xyz(x, y, z)
	xd, yd, zd = F11_C_Xyz(xd, yd, zd)
	fmt.Println("test of C_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50
// D50_B
	xd, yd, zd = D50_B_Bradford(x, y, z)
	xd, yd, zd = B_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_B_vonKries(x, y, z)
	xd, yd, zd = B_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_B_Xyz(x, y, z)
	xd, yd, zd = B_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_C
	xd, yd, zd = D50_C_Bradford(x, y, z)
	xd, yd, zd = C_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_C_vonKries(x, y, z)
	xd, yd, zd = C_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_C_Xyz(x, y, z)
	xd, yd, zd = C_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_A
	xd, yd, zd = D50_A_Bradford(x, y, z)
	xd, yd, zd = A_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_A_vonKries(x, y, z)
	xd, yd, zd = A_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_A_Xyz(x, y, z)
	xd, yd, zd = A_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_D55
	xd, yd, zd = D50_D55_Bradford(x, y, z)
	xd, yd, zd = D55_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D55_vonKries(x, y, z)
	xd, yd, zd = D55_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D55_Xyz(x, y, z)
	xd, yd, zd = D55_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_D65
	xd, yd, zd = D50_D65_Bradford(x, y, z)
	xd, yd, zd = D65_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D65_vonKries(x, y, z)
	xd, yd, zd = D65_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D65_Xyz(x, y, z)
	xd, yd, zd = D65_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_D75
	xd, yd, zd = D50_D75_Bradford(x, y, z)
	xd, yd, zd = D75_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D75_vonKries(x, y, z)
	xd, yd, zd = D75_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_D75_Xyz(x, y, z)
	xd, yd, zd = D75_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_E
	xd, yd, zd = D50_E_Bradford(x, y, z)
	xd, yd, zd = E_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_E_vonKries(x, y, z)
	xd, yd, zd = E_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_E_Xyz(x, y, z)
	xd, yd, zd = E_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_F2
	xd, yd, zd = D50_F2_Bradford(x, y, z)
	xd, yd, zd = F2_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F2_vonKries(x, y, z)
	xd, yd, zd = F2_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F2_Xyz(x, y, z)
	xd, yd, zd = F2_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_F7
	xd, yd, zd = D50_F7_Bradford(x, y, z)
	xd, yd, zd = F7_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F7_vonKries(x, y, z)
	xd, yd, zd = F7_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F7_Xyz(x, y, z)
	xd, yd, zd = F7_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D50_F11
	xd, yd, zd = D50_F11_Bradford(x, y, z)
	xd, yd, zd = F11_D50_Bradford(xd, yd, zd)
	fmt.Println("test of D50_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F11_vonKries(x, y, z)
	xd, yd, zd = F11_D50_vonKries(xd, yd, zd)
	fmt.Println("test of D50_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D50_F11_Xyz(x, y, z)
	xd, yd, zd = F11_D50_Xyz(xd, yd, zd)
	fmt.Println("test of D50_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55
// D55_B
	xd, yd, zd = D55_B_Bradford(x, y, z)
	xd, yd, zd = B_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_B_vonKries(x, y, z)
	xd, yd, zd = B_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_B_Xyz(x, y, z)
	xd, yd, zd = B_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_C
	xd, yd, zd = D55_C_Bradford(x, y, z)
	xd, yd, zd = C_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_C_vonKries(x, y, z)
	xd, yd, zd = C_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_C_Xyz(x, y, z)
	xd, yd, zd = C_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_D50
	xd, yd, zd = D55_D50_Bradford(x, y, z)
	xd, yd, zd = D50_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D50_vonKries(x, y, z)
	xd, yd, zd = D50_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D50_Xyz(x, y, z)
	xd, yd, zd = D50_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_A
	xd, yd, zd = D55_A_Bradford(x, y, z)
	xd, yd, zd = A_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_A_vonKries(x, y, z)
	xd, yd, zd = A_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_A_Xyz(x, y, z)
	xd, yd, zd = A_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_D65
	xd, yd, zd = D55_D65_Bradford(x, y, z)
	xd, yd, zd = D65_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D65_vonKries(x, y, z)
	xd, yd, zd = D65_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D65_Xyz(x, y, z)
	xd, yd, zd = D65_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_D75
	xd, yd, zd = D55_D75_Bradford(x, y, z)
	xd, yd, zd = D75_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D75_vonKries(x, y, z)
	xd, yd, zd = D75_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_D75_Xyz(x, y, z)
	xd, yd, zd = D75_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_E
	xd, yd, zd = D55_E_Bradford(x, y, z)
	xd, yd, zd = E_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_E_vonKries(x, y, z)
	xd, yd, zd = E_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_E_Xyz(x, y, z)
	xd, yd, zd = E_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_F2
	xd, yd, zd = D55_F2_Bradford(x, y, z)
	xd, yd, zd = F2_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F2_vonKries(x, y, z)
	xd, yd, zd = F2_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F2_Xyz(x, y, z)
	xd, yd, zd = F2_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_F7
	xd, yd, zd = D55_F7_Bradford(x, y, z)
	xd, yd, zd = F7_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F7_vonKries(x, y, z)
	xd, yd, zd = F7_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F7_Xyz(x, y, z)
	xd, yd, zd = F7_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D55_F11
	xd, yd, zd = D55_F11_Bradford(x, y, z)
	xd, yd, zd = F11_D55_Bradford(xd, yd, zd)
	fmt.Println("test of D55_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F11_vonKries(x, y, z)
	xd, yd, zd = F11_D55_vonKries(xd, yd, zd)
	fmt.Println("test of D55_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D55_F11_Xyz(x, y, z)
	xd, yd, zd = F11_D55_Xyz(xd, yd, zd)
	fmt.Println("test of D55_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	
// D65
// D65_B
	xd, yd, zd = D65_B_Bradford(x, y, z)
	xd, yd, zd = B_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_B_vonKries(x, y, z)
	xd, yd, zd = B_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_B_Xyz(x, y, z)
	xd, yd, zd = B_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_C
	xd, yd, zd = D65_C_Bradford(x, y, z)
	xd, yd, zd = C_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_C_vonKries(x, y, z)
	xd, yd, zd = C_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_C_Xyz(x, y, z)
	xd, yd, zd = C_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_D50
	xd, yd, zd = D65_D50_Bradford(x, y, z)
	xd, yd, zd = D50_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D50_vonKries(x, y, z)
	xd, yd, zd = D50_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D50_Xyz(x, y, z)
	xd, yd, zd = D50_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_D55
	xd, yd, zd = D65_D55_Bradford(x, y, z)
	xd, yd, zd = D55_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D55_vonKries(x, y, z)
	xd, yd, zd = D55_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D55_Xyz(x, y, z)
	xd, yd, zd = D55_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_A
	xd, yd, zd = D65_A_Bradford(x, y, z)
	xd, yd, zd = A_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_A_vonKries(x, y, z)
	xd, yd, zd = A_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_A_Xyz(x, y, z)
	xd, yd, zd = A_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_D75
	xd, yd, zd = D65_D75_Bradford(x, y, z)
	xd, yd, zd = D75_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D75_vonKries(x, y, z)
	xd, yd, zd = D75_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_D75_Xyz(x, y, z)
	xd, yd, zd = D75_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_E
	xd, yd, zd = D65_E_Bradford(x, y, z)
	xd, yd, zd = E_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_E_vonKries(x, y, z)
	xd, yd, zd = E_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_E_Xyz(x, y, z)
	xd, yd, zd = E_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_F2
	xd, yd, zd = D65_F2_Bradford(x, y, z)
	xd, yd, zd = F2_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F2_vonKries(x, y, z)
	xd, yd, zd = F2_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F2_Xyz(x, y, z)
	xd, yd, zd = F2_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_F7
	xd, yd, zd = D65_F7_Bradford(x, y, z)
	xd, yd, zd = F7_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F7_vonKries(x, y, z)
	xd, yd, zd = F7_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F7_Xyz(x, y, z)
	xd, yd, zd = F7_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D65_F11
	xd, yd, zd = D65_F11_Bradford(x, y, z)
	xd, yd, zd = F11_D65_Bradford(xd, yd, zd)
	fmt.Println("test of D65_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F11_vonKries(x, y, z)
	xd, yd, zd = F11_D65_vonKries(xd, yd, zd)
	fmt.Println("test of D65_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D65_F11_Xyz(x, y, z)
	xd, yd, zd = F11_D65_Xyz(xd, yd, zd)
	fmt.Println("test of D65_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75
// D75_B
	xd, yd, zd = D75_B_Bradford(x, y, z)
	xd, yd, zd = B_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_B_vonKries(x, y, z)
	xd, yd, zd = B_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_B_Xyz(x, y, z)
	xd, yd, zd = B_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_C
	xd, yd, zd = D75_C_Bradford(x, y, z)
	xd, yd, zd = C_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_C_vonKries(x, y, z)
	xd, yd, zd = C_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_C_Xyz(x, y, z)
	xd, yd, zd = C_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_D50
	xd, yd, zd = D75_D50_Bradford(x, y, z)
	xd, yd, zd = D50_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D50_vonKries(x, y, z)
	xd, yd, zd = D50_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D50_Xyz(x, y, z)
	xd, yd, zd = D50_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_D55
	xd, yd, zd = D75_D55_Bradford(x, y, z)
	xd, yd, zd = D55_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D55_vonKries(x, y, z)
	xd, yd, zd = D55_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D55_Xyz(x, y, z)
	xd, yd, zd = D55_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_D65
	xd, yd, zd = D75_D65_Bradford(x, y, z)
	xd, yd, zd = D65_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D65_vonKries(x, y, z)
	xd, yd, zd = D65_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_D65_Xyz(x, y, z)
	xd, yd, zd = D65_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_A
	xd, yd, zd = D75_A_Bradford(x, y, z)
	xd, yd, zd = A_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_A_vonKries(x, y, z)
	xd, yd, zd = A_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_A_Xyz(x, y, z)
	xd, yd, zd = A_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_E
	xd, yd, zd = D75_E_Bradford(x, y, z)
	xd, yd, zd = E_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_E_vonKries(x, y, z)
	xd, yd, zd = E_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_E_Xyz(x, y, z)
	xd, yd, zd = E_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_F2
	xd, yd, zd = D75_F2_Bradford(x, y, z)
	xd, yd, zd = F2_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F2_vonKries(x, y, z)
	xd, yd, zd = F2_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F2_Xyz(x, y, z)
	xd, yd, zd = F2_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_F7
	xd, yd, zd = D75_F7_Bradford(x, y, z)
	xd, yd, zd = F7_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F7_vonKries(x, y, z)
	xd, yd, zd = F7_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F7_Xyz(x, y, z)
	xd, yd, zd = F7_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// D75_F11
	xd, yd, zd = D75_F11_Bradford(x, y, z)
	xd, yd, zd = F11_D75_Bradford(xd, yd, zd)
	fmt.Println("test of D75_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F11_vonKries(x, y, z)
	xd, yd, zd = F11_D75_vonKries(xd, yd, zd)
	fmt.Println("test of D75_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = D75_F11_Xyz(x, y, z)
	xd, yd, zd = F11_D75_Xyz(xd, yd, zd)
	fmt.Println("test of D75_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E
// E_B
	xd, yd, zd = E_B_Bradford(x, y, z)
	xd, yd, zd = B_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_B_vonKries(x, y, z)
	xd, yd, zd = B_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_B_Xyz(x, y, z)
	xd, yd, zd = B_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_C
	xd, yd, zd = E_C_Bradford(x, y, z)
	xd, yd, zd = C_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_C_vonKries(x, y, z)
	xd, yd, zd = C_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_C_Xyz(x, y, z)
	xd, yd, zd = C_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_D50
	xd, yd, zd = E_D50_Bradford(x, y, z)
	xd, yd, zd = D50_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D50_vonKries(x, y, z)
	xd, yd, zd = D50_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D50_Xyz(x, y, z)
	xd, yd, zd = D50_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_D55
	xd, yd, zd = E_D55_Bradford(x, y, z)
	xd, yd, zd = D55_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D55_vonKries(x, y, z)
	xd, yd, zd = D55_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D55_Xyz(x, y, z)
	xd, yd, zd = D55_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_D65
	xd, yd, zd = E_D65_Bradford(x, y, z)
	xd, yd, zd = D65_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D65_vonKries(x, y, z)
	xd, yd, zd = D65_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D65_Xyz(x, y, z)
	xd, yd, zd = D65_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_D75
	xd, yd, zd = E_D75_Bradford(x, y, z)
	xd, yd, zd = D75_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D75_vonKries(x, y, z)
	xd, yd, zd = D75_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_D75_Xyz(x, y, z)
	xd, yd, zd = D75_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_A
	xd, yd, zd = E_A_Bradford(x, y, z)
	xd, yd, zd = A_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_A_vonKries(x, y, z)
	xd, yd, zd = A_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_A_Xyz(x, y, z)
	xd, yd, zd = A_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_F2
	xd, yd, zd = E_F2_Bradford(x, y, z)
	xd, yd, zd = F2_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F2_vonKries(x, y, z)
	xd, yd, zd = F2_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F2_Xyz(x, y, z)
	xd, yd, zd = F2_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_F7
	xd, yd, zd = E_F7_Bradford(x, y, z)
	xd, yd, zd = F7_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F7_vonKries(x, y, z)
	xd, yd, zd = F7_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F7_Xyz(x, y, z)
	xd, yd, zd = F7_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// E_F11
	xd, yd, zd = E_F11_Bradford(x, y, z)
	xd, yd, zd = F11_E_Bradford(xd, yd, zd)
	fmt.Println("test of E_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F11_vonKries(x, y, z)
	xd, yd, zd = F11_E_vonKries(xd, yd, zd)
	fmt.Println("test of E_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = E_F11_Xyz(x, y, z)
	xd, yd, zd = F11_E_Xyz(xd, yd, zd)
	fmt.Println("test of E_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2
// F2_B
	xd, yd, zd = F2_B_Bradford(x, y, z)
	xd, yd, zd = B_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_B_vonKries(x, y, z)
	xd, yd, zd = B_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_B_Xyz(x, y, z)
	xd, yd, zd = B_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_C
	xd, yd, zd = F2_C_Bradford(x, y, z)
	xd, yd, zd = C_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_C_vonKries(x, y, z)
	xd, yd, zd = C_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_C_Xyz(x, y, z)
	xd, yd, zd = C_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_D50
	xd, yd, zd = F2_D50_Bradford(x, y, z)
	xd, yd, zd = D50_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D50_vonKries(x, y, z)
	xd, yd, zd = D50_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D50_Xyz(x, y, z)
	xd, yd, zd = D50_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_D55
	xd, yd, zd = F2_D55_Bradford(x, y, z)
	xd, yd, zd = D55_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D55_vonKries(x, y, z)
	xd, yd, zd = D55_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D55_Xyz(x, y, z)
	xd, yd, zd = D55_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_D65
	xd, yd, zd = F2_D65_Bradford(x, y, z)
	xd, yd, zd = D65_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D65_vonKries(x, y, z)
	xd, yd, zd = D65_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D65_Xyz(x, y, z)
	xd, yd, zd = D65_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_D75
	xd, yd, zd = F2_D75_Bradford(x, y, z)
	xd, yd, zd = D75_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D75_vonKries(x, y, z)
	xd, yd, zd = D75_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_D75_Xyz(x, y, z)
	xd, yd, zd = D75_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_E
	xd, yd, zd = F2_E_Bradford(x, y, z)
	xd, yd, zd = E_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_E_vonKries(x, y, z)
	xd, yd, zd = E_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_E_Xyz(x, y, z)
	xd, yd, zd = E_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_A
	xd, yd, zd = F2_A_Bradford(x, y, z)
	xd, yd, zd = A_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_A_vonKries(x, y, z)
	xd, yd, zd = A_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_A_Xyz(x, y, z)
	xd, yd, zd = A_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_F7
	xd, yd, zd = F2_F7_Bradford(x, y, z)
	xd, yd, zd = F7_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_F7_vonKries(x, y, z)
	xd, yd, zd = F7_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_F7_Xyz(x, y, z)
	xd, yd, zd = F7_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F2_F11
	xd, yd, zd = F2_F11_Bradford(x, y, z)
	xd, yd, zd = F11_F2_Bradford(xd, yd, zd)
	fmt.Println("test of F2_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_F11_vonKries(x, y, z)
	xd, yd, zd = F11_F2_vonKries(xd, yd, zd)
	fmt.Println("test of F2_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F2_F11_Xyz(x, y, z)
	xd, yd, zd = F11_F2_Xyz(xd, yd, zd)
	fmt.Println("test of F2_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7
// F7_B
	xd, yd, zd = F7_B_Bradford(x, y, z)
	xd, yd, zd = B_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_B_vonKries(x, y, z)
	xd, yd, zd = B_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_B_Xyz(x, y, z)
	xd, yd, zd = B_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_C
	xd, yd, zd = F7_C_Bradford(x, y, z)
	xd, yd, zd = C_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_C_vonKries(x, y, z)
	xd, yd, zd = C_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_C_Xyz(x, y, z)
	xd, yd, zd = C_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_D50
	xd, yd, zd = F7_D50_Bradford(x, y, z)
	xd, yd, zd = D50_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D50_vonKries(x, y, z)
	xd, yd, zd = D50_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D50_Xyz(x, y, z)
	xd, yd, zd = D50_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_D55
	xd, yd, zd = F7_D55_Bradford(x, y, z)
	xd, yd, zd = D55_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D55_vonKries(x, y, z)
	xd, yd, zd = D55_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D55_Xyz(x, y, z)
	xd, yd, zd = D55_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_D65
	xd, yd, zd = F7_D65_Bradford(x, y, z)
	xd, yd, zd = D65_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D65_vonKries(x, y, z)
	xd, yd, zd = D65_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D65_Xyz(x, y, z)
	xd, yd, zd = D65_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_D75
	xd, yd, zd = F7_D75_Bradford(x, y, z)
	xd, yd, zd = D75_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D75_vonKries(x, y, z)
	xd, yd, zd = D75_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_D75_Xyz(x, y, z)
	xd, yd, zd = D75_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_E
	xd, yd, zd = F7_E_Bradford(x, y, z)
	xd, yd, zd = E_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_E_vonKries(x, y, z)
	xd, yd, zd = E_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_E_Xyz(x, y, z)
	xd, yd, zd = E_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_F2
	xd, yd, zd = F7_F2_Bradford(x, y, z)
	xd, yd, zd = F2_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_F2_vonKries(x, y, z)
	xd, yd, zd = F2_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_F2_Xyz(x, y, z)
	xd, yd, zd = F2_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_A
	xd, yd, zd = F7_A_Bradford(x, y, z)
	xd, yd, zd = A_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_A_vonKries(x, y, z)
	xd, yd, zd = A_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_A_Xyz(x, y, z)
	xd, yd, zd = A_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F7_F11
	xd, yd, zd = F7_F11_Bradford(x, y, z)
	xd, yd, zd = F11_F7_Bradford(xd, yd, zd)
	fmt.Println("test of F7_F11_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_F11_vonKries(x, y, z)
	xd, yd, zd = F11_F7_vonKries(xd, yd, zd)
	fmt.Println("test of F7_F11_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F7_F11_Xyz(x, y, z)
	xd, yd, zd = F11_F7_Xyz(xd, yd, zd)
	fmt.Println("test of F7_F11_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11
// F11_B
	xd, yd, zd = F11_B_Bradford(x, y, z)
	xd, yd, zd = B_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_B_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_B_vonKries(x, y, z)
	xd, yd, zd = B_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_B_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_B_Xyz(x, y, z)
	xd, yd, zd = B_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_B_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_C
	xd, yd, zd = F11_C_Bradford(x, y, z)
	xd, yd, zd = C_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_C_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_C_vonKries(x, y, z)
	xd, yd, zd = C_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_C_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_C_Xyz(x, y, z)
	xd, yd, zd = C_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_C_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_D50
	xd, yd, zd = F11_D50_Bradford(x, y, z)
	xd, yd, zd = D50_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_D50_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D50_vonKries(x, y, z)
	xd, yd, zd = D50_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_D50_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D50_Xyz(x, y, z)
	xd, yd, zd = D50_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_D50_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_D55
	xd, yd, zd = F11_D55_Bradford(x, y, z)
	xd, yd, zd = D55_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_D55_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D55_vonKries(x, y, z)
	xd, yd, zd = D55_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_D55_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D55_Xyz(x, y, z)
	xd, yd, zd = D55_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_D55_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_D65
	xd, yd, zd = F11_D65_Bradford(x, y, z)
	xd, yd, zd = D65_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_D65_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D65_vonKries(x, y, z)
	xd, yd, zd = D65_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_D65_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D65_Xyz(x, y, z)
	xd, yd, zd = D65_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_D65_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_D75
	xd, yd, zd = F11_D75_Bradford(x, y, z)
	xd, yd, zd = D75_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_D75_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D75_vonKries(x, y, z)
	xd, yd, zd = D75_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_D75_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_D75_Xyz(x, y, z)
	xd, yd, zd = D75_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_D75_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_E
	xd, yd, zd = F11_E_Bradford(x, y, z)
	xd, yd, zd = E_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_E_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_E_vonKries(x, y, z)
	xd, yd, zd = E_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_E_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_E_Xyz(x, y, z)
	xd, yd, zd = E_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_E_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_F2
	xd, yd, zd = F11_F2_Bradford(x, y, z)
	xd, yd, zd = F2_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_F2_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_F2_vonKries(x, y, z)
	xd, yd, zd = F2_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_F2_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_F2_Xyz(x, y, z)
	xd, yd, zd = F2_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_F2_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_F7
	xd, yd, zd = F11_F7_Bradford(x, y, z)
	xd, yd, zd = F7_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_F7_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_F7_vonKries(x, y, z)
	xd, yd, zd = F7_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_F7_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_F7_Xyz(x, y, z)
	xd, yd, zd = F7_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_F7_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

// F11_A
	xd, yd, zd = F11_A_Bradford(x, y, z)
	xd, yd, zd = A_F11_Bradford(xd, yd, zd)
	fmt.Println("test of F11_A_Bradford")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_A_vonKries(x, y, z)
	xd, yd, zd = A_F11_vonKries(xd, yd, zd)
	fmt.Println("test of F11_A_vonKries")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	xd, yd, zd = F11_A_Xyz(x, y, z)
	xd, yd, zd = A_F11_Xyz(xd, yd, zd)
	fmt.Println("test of F11_A_Xyz")
	if !(check_diff(xd, x) && check_diff(yd, y) && check_diff(zd, z)) {
		t.Errorf("x, y, z : %f, %f, %f ", x, y, z)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

return
}
