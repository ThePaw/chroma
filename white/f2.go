// Copyright 2012 The Chroma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package white

// F2 illuminant conversion functions
// F2_A functions
func F2_A_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1108436, 0.0654914, -0.1020770},
		{0.0892593, 0.9359084, -0.0362665},
		{-0.0166489, 0.0256618, 0.5144475}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_A_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0383067, 0.1380821, -0.1030329},
		{0.0151673, 0.9870182, -0.0030598},
		{0.0000000, 0.0000000, 0.5280222}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_A_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.1075152, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.5280222}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_B functions
func F2_B_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9773601, -0.0123986, 0.0500263},
		{-0.0133935, 1.0023946, 0.0161588},
		{0.0104080, -0.0177628, 1.2756066}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_B_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9916737, -0.0300157, 0.0551010},
		{-0.0032970, 1.0028226, 0.0006642},
		{0.0000000, 0.0000000, 1.2645675}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_B_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9988506, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.2645675}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_C functions
func F2_C_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9301461, -0.0386152, 0.1436063},
		{-0.0431066, 1.0113279, 0.0466336},
		{0.0295321, -0.0501939, 1.7853816}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_C_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9744947, -0.0919438, 0.1574666},
		{-0.0100994, 1.0086458, 0.0020348},
		{0.0000000, 0.0000000, 1.7543662}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_C_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9887887, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.7543662}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_D50 functions
func F2_D50_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9628262, -0.0215790, 0.0457172},
		{-0.0280310, 1.0172847, 0.0156071},
		{0.0083415, -0.0135344, 1.2322804}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D50_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9869554, -0.0470220, 0.0479582},
		{-0.0051650, 1.0044210, 0.0010416},
		{0.0000000, 0.0000000, 1.2244744}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D50_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9721332, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.2244744}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_D55 functions
func F2_D55_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9460314, -0.0310871, 0.0735630},
		{-0.0395027, 1.0224574, 0.0248152},
		{0.0138373, -0.0227352, 1.3807080}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D55_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9809382, -0.0687128, 0.0780192},
		{-0.0075476, 1.0064606, 0.0015219},
		{0.0000000, 0.0000000, 1.3673379}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D55_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9646724, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.3673379}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_D65 functions
func F2_D65_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9212269, -0.0449128, 0.1211620},
		{-0.0553723, 1.0277243, 0.0403563},
		{0.0235086, -0.0391019, 1.6390644}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D65_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9719399, -0.1011504, 0.1299721},
		{-0.0111107, 1.0095107, 0.0022399},
		{0.0000000, 0.0000000, 1.6156426}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D65_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9582703, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.6156426}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_D75 functions
func F2_D75_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9038485, -0.0544241, 0.1597390},
		{-0.0656239, 1.0294976, 0.0528128},
		{0.0315403, -0.0528125, 1.8516893}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D75_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9655456, -0.1242015, 0.1724740},
		{-0.0136427, 1.0116783, 0.0027499},
		{0.0000000, 0.0000000, 1.8197439}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_D75_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9575142, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.8197439}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_E functions
func F2_E_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9654834, -0.0184347, 0.0902324},
		{-0.0181044, 0.9985320, 0.0288236},
		{0.0192213, -0.0330731, 1.5046195}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_E_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9870655, -0.0466289, 0.1003043},
		{-0.0051218, 1.0043851, 0.0010314},
		{0.0000000, 0.0000000, 1.4838336}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_E_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0082068, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.4838336}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_F7 functions
func F2_F7_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9213533, -0.0448467, 0.1207888},
		{-0.0553131, 1.0277453, 0.0402379},
		{0.0234280, -0.0389625, 1.6369581}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_F7_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9719880, -0.1009770, 0.1295549},
		{-0.0110916, 1.0094944, 0.0022360},
		{0.0000000, 0.0000000, 1.6136246}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_F7_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{0.9582098, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 1.6136246}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

// F2_F11 functions
func F2_F11_Bradford(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0156947, 0.0093992, -0.0106926},
		{0.0132620, 0.9895464, -0.0040071},
		{-0.0014541, 0.0020196, 0.9539904}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_F11_vonKries(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0053592, 0.0193178, -0.0101990},
		{0.0021219, 0.9981839, -0.0004282},
		{0.0000000, 0.0000000, 0.9548469}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}

func F2_F11_Xyz(xs, ys, zs float64) (xd, yd, zd float64) {
	m := [3][3]float64{
		{1.0179058, 0.0000000, 0.0000000},
		{0.0000000, 1.0000000, 0.0000000},
		{0.0000000, 0.0000000, 0.9548469}}
	xd = m[0][0]*xs + m[0][1]*ys + m[0][2]*zs
	yd = m[1][0]*xs + m[1][1]*ys + m[1][2]*zs
	zd = m[2][0]*xs + m[2][1]*ys + m[2][2]*zs
	return
}
