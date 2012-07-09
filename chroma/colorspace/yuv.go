package chroma

// Convert sRGB to NTSC/PAL Y'UV Luma + Chroma
// http://en.wikipedia.org/wiki/YUV
func RGB2YUV(r, g, b uint8) (y, u, v float64) {
	rf := float64(r)
	gf := float64(g)
	bf := float64(b)

	y = 0.299*rf + 0.587*gf + 0.114*bf
	u = -0.147*rf - 0.289*gf + 0.436*bf
	v = 0.615*rf - 0.515*gf - 0.100*bf
	return
}

// Convert NTSC/PAL Y'UV to sRGB
func YUV2RGB(y, u, v float64) (r, g, b uint8) {
	rf := y - 3.945707070708279e-05*u + 1.1398279671717170825*v
	gf := y - 0.3946101641414141437*u - 0.5805003156565656797*v
	bf := y + 2.0319996843434342537*u - 4.813762626262513e-04*v
	r = uint8(rf)
	g = uint8(gf)
	b = uint8(bf)
	return
}
