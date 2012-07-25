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

/*
// fails
func TestChromAd(t *testing.T) {
	xws, yws, zws := 1.09850, 1.00000, 0.35585
	xwd, ywd, zwd := 0.95047, 1.00000, 1.08883
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.104230, -0.141803, -2.510429 // Pre-calculated values

	xd, yd, zd := ChromAd(xs, ys, zs, xws, yws, zws, xwd, ywd, zwd, 0)

	fmt.Println("test of Chromatic Adaptation from A to D65, _Bradford")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}

	return
}
*/

func TestD50_D65_Bradford(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.426782, -0.003866, -1.055336 // Pre-calculated values

	xd, yd, zd := D50_D65_Bradford(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D50 to D65, _Bradford")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}

func TestD50_D65_vonKries(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.436835, 0.022936, -1.052595 // Pre-calculated values

	xd, yd, zd := D50_D65_vonKries(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D50 to D65, _vonKries")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}

func TestD50_D65_Xyz(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.492870, 0.026772, -1.052599 // Pre-calculated values

	xd, yd, zd := D50_D65_Xyz(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D50 to D65, _Xyz")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}

func TestD65_D50_Bradford(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.564507, 0.054889, -0.604228

	xd, yd, zd := D65_D50_Bradford(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D65 to D50, _Bradford")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}

func TestD65_D50_vonKries(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.551109, 0.030663, -0.604603 // Pre-calculated values

	xd, yd, zd := D65_D50_vonKries(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D65 to D50, _vonKries")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}

func TestD65_D50_Xyz(t *testing.T) {
	xs, ys, zs := 0.500000, 0.026772, -0.797751
	xd0, yd0, zd0 := 0.507233, 0.026772, -0.604605 // Pre-calculated values

	xd, yd, zd := D65_D50_Xyz(xs, ys, zs)

	fmt.Println("test of Chromatic Adaptation from D65 to D50, _Xyz")
	if !(check_diff(xd, xd0) && check_diff(yd, yd0) && check_diff(zd, zd0)) {
		t.Errorf("xd0, yd0, zd0 : %f, %f, %f ", xd0, yd0, zd0)
		t.Errorf("xd, yd, zd : %f, %f, %f \n\n", xd, yd, zd)

	}
	return
}
