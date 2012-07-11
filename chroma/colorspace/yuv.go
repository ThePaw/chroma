package colorspace

// Convert sRGB to NTSC/PAL Y'UV Luma + Chroma
// http://en.wikipedia.org/wiki/YUV
func Rgb2Yuv(r, g, b float64) (y, u, v float64) {
	y = 0.299*r + 0.587*g + 0.114*b
	u = -0.147*r - 0.289*g + 0.436*b
	v = 0.615*r - 0.515*g - 0.100*b
	return
}

// Convert NTSC/PAL Y'UV to sRGB
func Yuv2Rgb(y, u, v float64) (r, g, b float64) {
	r = y - 3.945707070708279e-05*u + 1.1398279671717170825*v
	g = y - 0.3946101641414141437*u - 0.5805003156565656797*v
	b = y + 2.0319996843434342537*u - 4.813762626262513e-04*v
	return
}
