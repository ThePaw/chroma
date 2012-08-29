// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

package rgb

// Conversion of different RGB colorspaces with their native illuminators (reference whites) to CIE XYZ scaled to [0, 1e9] and back. 
// RGB values MUST BE LINEAR and in the nominal range [0, 1]. 
// Ref.: [24] 

// AdobeToApple converts from Adobe RGB to Apple RGB colorspace. 
func AdobeToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// AdobeToBest converts from Adobe RGB to Best RGB colorspace. 
func AdobeToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// AdobeToBeta converts from Adobe RGB to Beta RGB colorspace. 
func AdobeToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// AdobeToBruce converts from Adobe RGB to Bruce RGB colorspace. 
func AdobeToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// AdobeToCIE converts from Adobe RGB to CIE RGB colorspace. 
func AdobeToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// AdobeToColorMatch converts from Adobe RGB to ColorMatch RGB colorspace. 
func AdobeToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// AdobeToDon converts from Adobe RGB to Don RGB colorspace. 
func AdobeToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// AdobeToECI converts from Adobe RGB to ECI RGB colorspace. 
func AdobeToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// AdobeToEktaSpace converts from Adobe RGB to EktaSpace RGB colorspace. 
func AdobeToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// AdobeToNTSC converts from Adobe RGB to NTSC RGB colorspace. 
func AdobeToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// AdobeToPAL converts from Adobe RGB to PAL RGB colorspace. 
func AdobeToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// AdobeToProPhoto converts from Adobe RGB to ProPhoto RGB colorspace. 
func AdobeToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// AdobeToSMPTE_C converts from Adobe RGB to SMPTE_C RGB colorspace. 
func AdobeToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// AdobeToSRGB converts from Adobe RGB to SRGB RGB colorspace. 
func AdobeToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// AdobeToWGamut converts from Adobe RGB to WGamut RGB colorspace. 
func AdobeToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// AdobeToAdobe_D50 converts from Adobe RGB to Adobe_D50 RGB colorspace. 
func AdobeToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// AdobeToApple_D50 converts from Adobe RGB to Apple_D50 RGB colorspace. 
func AdobeToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// AdobeToBruce_D50 converts from Adobe RGB to Bruce_D50 RGB colorspace. 
func AdobeToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// AdobeToCie_D50 converts from Adobe RGB to Cie_D50 RGB colorspace. 
func AdobeToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// AdobeToNTSC_D50 converts from Adobe RGB to NTSC_D50 RGB colorspace. 
func AdobeToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// AdobeToPAL_D50 converts from Adobe RGB to PAL_D50 RGB colorspace. 
func AdobeToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// AdobeToSMPTE_C_D50 converts from Adobe RGB to SMPTE_C_D50 RGB colorspace. 
func AdobeToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// AdobeToSRGB_D50 converts from Adobe RGB to SRGB_D50 RGB colorspace. 
func AdobeToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// AppleToAdobe converts from Apple RGB to Adobe RGB colorspace. 
func AppleToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// AppleToBest converts from Apple RGB to Best RGB colorspace. 
func AppleToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// AppleToBeta converts from Apple RGB to Beta RGB colorspace. 
func AppleToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// AppleToBruce converts from Apple RGB to Bruce RGB colorspace. 
func AppleToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// AppleToCIE converts from Apple RGB to CIE RGB colorspace. 
func AppleToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// AppleToColorMatch converts from Apple RGB to ColorMatch RGB colorspace. 
func AppleToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// AppleToDon converts from Apple RGB to Don RGB colorspace. 
func AppleToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// AppleToECI converts from Apple RGB to ECI RGB colorspace. 
func AppleToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// AppleToEktaSpace converts from Apple RGB to EktaSpace RGB colorspace. 
func AppleToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// AppleToNTSC converts from Apple RGB to NTSC RGB colorspace. 
func AppleToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// AppleToPAL converts from Apple RGB to PAL RGB colorspace. 
func AppleToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// AppleToProPhoto converts from Apple RGB to ProPhoto RGB colorspace. 
func AppleToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// AppleToSMPTE_C converts from Apple RGB to SMPTE_C RGB colorspace. 
func AppleToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// AppleToSRGB converts from Apple RGB to SRGB RGB colorspace. 
func AppleToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// AppleToWGamut converts from Apple RGB to WGamut RGB colorspace. 
func AppleToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// AppleToAdobe_D50 converts from Apple RGB to Adobe_D50 RGB colorspace. 
func AppleToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// AppleToApple_D50 converts from Apple RGB to Apple_D50 RGB colorspace. 
func AppleToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// AppleToBruce_D50 converts from Apple RGB to Bruce_D50 RGB colorspace. 
func AppleToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// AppleToCie_D50 converts from Apple RGB to Cie_D50 RGB colorspace. 
func AppleToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// AppleToNTSC_D50 converts from Apple RGB to NTSC_D50 RGB colorspace. 
func AppleToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// AppleToPAL_D50 converts from Apple RGB to PAL_D50 RGB colorspace. 
func AppleToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// AppleToSMPTE_C_D50 converts from Apple RGB to SMPTE_C_D50 RGB colorspace. 
func AppleToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// AppleToSRGB_D50 converts from Apple RGB to SRGB_D50 RGB colorspace. 
func AppleToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// BestToAdobe converts from Best RGB to Adobe RGB colorspace. 
func BestToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// BestToApple converts from Best RGB to Apple RGB colorspace. 
func BestToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// BestToBeta converts from Best RGB to Beta RGB colorspace. 
func BestToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// BestToBruce converts from Best RGB to Bruce RGB colorspace. 
func BestToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// BestToCIE converts from Best RGB to CIE RGB colorspace. 
func BestToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// BestToColorMatch converts from Best RGB to ColorMatch RGB colorspace. 
func BestToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// BestToDon converts from Best RGB to Don RGB colorspace. 
func BestToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// BestToECI converts from Best RGB to ECI RGB colorspace. 
func BestToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// BestToEktaSpace converts from Best RGB to EktaSpace RGB colorspace. 
func BestToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// BestToNTSC converts from Best RGB to NTSC RGB colorspace. 
func BestToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// BestToPAL converts from Best RGB to PAL RGB colorspace. 
func BestToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// BestToProPhoto converts from Best RGB to ProPhoto RGB colorspace. 
func BestToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// BestToSMPTE_C converts from Best RGB to SMPTE_C RGB colorspace. 
func BestToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// BestToSRGB converts from Best RGB to SRGB RGB colorspace. 
func BestToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// BestToWGamut converts from Best RGB to WGamut RGB colorspace. 
func BestToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// BestToAdobe_D50 converts from Best RGB to Adobe_D50 RGB colorspace. 
func BestToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// BestToApple_D50 converts from Best RGB to Apple_D50 RGB colorspace. 
func BestToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// BestToBruce_D50 converts from Best RGB to Bruce_D50 RGB colorspace. 
func BestToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// BestToCie_D50 converts from Best RGB to Cie_D50 RGB colorspace. 
func BestToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// BestToNTSC_D50 converts from Best RGB to NTSC_D50 RGB colorspace. 
func BestToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// BestToPAL_D50 converts from Best RGB to PAL_D50 RGB colorspace. 
func BestToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// BestToSMPTE_C_D50 converts from Best RGB to SMPTE_C_D50 RGB colorspace. 
func BestToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// BestToSRGB_D50 converts from Best RGB to SRGB_D50 RGB colorspace. 
func BestToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BestToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// BetaToAdobe converts from Beta RGB to Adobe RGB colorspace. 
func BetaToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// BetaToApple converts from Beta RGB to Apple RGB colorspace. 
func BetaToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// BetaToBest converts from Beta RGB to Best RGB colorspace. 
func BetaToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// BetaToBruce converts from Beta RGB to Bruce RGB colorspace. 
func BetaToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// BetaToCIE converts from Beta RGB to CIE RGB colorspace. 
func BetaToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// BetaToColorMatch converts from Beta RGB to ColorMatch RGB colorspace. 
func BetaToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// BetaToDon converts from Beta RGB to Don RGB colorspace. 
func BetaToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// BetaToECI converts from Beta RGB to ECI RGB colorspace. 
func BetaToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// BetaToEktaSpace converts from Beta RGB to EktaSpace RGB colorspace. 
func BetaToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// BetaToNTSC converts from Beta RGB to NTSC RGB colorspace. 
func BetaToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// BetaToPAL converts from Beta RGB to PAL RGB colorspace. 
func BetaToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// BetaToProPhoto converts from Beta RGB to ProPhoto RGB colorspace. 
func BetaToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// BetaToSMPTE_C converts from Beta RGB to SMPTE_C RGB colorspace. 
func BetaToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// BetaToSRGB converts from Beta RGB to SRGB RGB colorspace. 
func BetaToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// BetaToWGamut converts from Beta RGB to WGamut RGB colorspace. 
func BetaToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// BetaToAdobe_D50 converts from Beta RGB to Adobe_D50 RGB colorspace. 
func BetaToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// BetaToApple_D50 converts from Beta RGB to Apple_D50 RGB colorspace. 
func BetaToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// BetaToBruce_D50 converts from Beta RGB to Bruce_D50 RGB colorspace. 
func BetaToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// BetaToCie_D50 converts from Beta RGB to Cie_D50 RGB colorspace. 
func BetaToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// BetaToNTSC_D50 converts from Beta RGB to NTSC_D50 RGB colorspace. 
func BetaToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// BetaToPAL_D50 converts from Beta RGB to PAL_D50 RGB colorspace. 
func BetaToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// BetaToSMPTE_C_D50 converts from Beta RGB to SMPTE_C_D50 RGB colorspace. 
func BetaToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// BetaToSRGB_D50 converts from Beta RGB to SRGB_D50 RGB colorspace. 
func BetaToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BetaToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// BruceToAdobe converts from Bruce RGB to Adobe RGB colorspace. 
func BruceToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// BruceToApple converts from Bruce RGB to Apple RGB colorspace. 
func BruceToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// BruceToBest converts from Bruce RGB to Best RGB colorspace. 
func BruceToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// BruceToBeta converts from Bruce RGB to Beta RGB colorspace. 
func BruceToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// BruceToCIE converts from Bruce RGB to CIE RGB colorspace. 
func BruceToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// BruceToColorMatch converts from Bruce RGB to ColorMatch RGB colorspace. 
func BruceToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// BruceToDon converts from Bruce RGB to Don RGB colorspace. 
func BruceToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// BruceToECI converts from Bruce RGB to ECI RGB colorspace. 
func BruceToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// BruceToEktaSpace converts from Bruce RGB to EktaSpace RGB colorspace. 
func BruceToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// BruceToNTSC converts from Bruce RGB to NTSC RGB colorspace. 
func BruceToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// BruceToPAL converts from Bruce RGB to PAL RGB colorspace. 
func BruceToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// BruceToProPhoto converts from Bruce RGB to ProPhoto RGB colorspace. 
func BruceToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// BruceToSMPTE_C converts from Bruce RGB to SMPTE_C RGB colorspace. 
func BruceToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// BruceToSRGB converts from Bruce RGB to SRGB RGB colorspace. 
func BruceToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// BruceToWGamut converts from Bruce RGB to WGamut RGB colorspace. 
func BruceToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// BruceToAdobe_D50 converts from Bruce RGB to Adobe_D50 RGB colorspace. 
func BruceToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// BruceToApple_D50 converts from Bruce RGB to Apple_D50 RGB colorspace. 
func BruceToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// BruceToBruce_D50 converts from Bruce RGB to Bruce_D50 RGB colorspace. 
func BruceToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// BruceToCie_D50 converts from Bruce RGB to Cie_D50 RGB colorspace. 
func BruceToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// BruceToNTSC_D50 converts from Bruce RGB to NTSC_D50 RGB colorspace. 
func BruceToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// BruceToPAL_D50 converts from Bruce RGB to PAL_D50 RGB colorspace. 
func BruceToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// BruceToSMPTE_C_D50 converts from Bruce RGB to SMPTE_C_D50 RGB colorspace. 
func BruceToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// BruceToSRGB_D50 converts from Bruce RGB to SRGB_D50 RGB colorspace. 
func BruceToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// CIEToAdobe converts from CIE RGB to Adobe RGB colorspace. 
func CIEToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// CIEToApple converts from CIE RGB to Apple RGB colorspace. 
func CIEToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// CIEToBest converts from CIE RGB to Best RGB colorspace. 
func CIEToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// CIEToBeta converts from CIE RGB to Beta RGB colorspace. 
func CIEToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// CIEToBruce converts from CIE RGB to Bruce RGB colorspace. 
func CIEToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// CIEToColorMatch converts from CIE RGB to ColorMatch RGB colorspace. 
func CIEToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// CIEToDon converts from CIE RGB to Don RGB colorspace. 
func CIEToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// CIEToECI converts from CIE RGB to ECI RGB colorspace. 
func CIEToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// CIEToEktaSpace converts from CIE RGB to EktaSpace RGB colorspace. 
func CIEToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// CIEToNTSC converts from CIE RGB to NTSC RGB colorspace. 
func CIEToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// CIEToPAL converts from CIE RGB to PAL RGB colorspace. 
func CIEToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// CIEToProPhoto converts from CIE RGB to ProPhoto RGB colorspace. 
func CIEToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// CIEToSMPTE_C converts from CIE RGB to SMPTE_C RGB colorspace. 
func CIEToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// CIEToSRGB converts from CIE RGB to SRGB RGB colorspace. 
func CIEToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// CIEToWGamut converts from CIE RGB to WGamut RGB colorspace. 
func CIEToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// CIEToAdobe_D50 converts from CIE RGB to Adobe_D50 RGB colorspace. 
func CIEToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// CIEToApple_D50 converts from CIE RGB to Apple_D50 RGB colorspace. 
func CIEToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// CIEToBruce_D50 converts from CIE RGB to Bruce_D50 RGB colorspace. 
func CIEToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// CIEToCie_D50 converts from CIE RGB to Cie_D50 RGB colorspace. 
func CIEToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// CIEToNTSC_D50 converts from CIE RGB to NTSC_D50 RGB colorspace. 
func CIEToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// CIEToPAL_D50 converts from CIE RGB to PAL_D50 RGB colorspace. 
func CIEToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// CIEToSMPTE_C_D50 converts from CIE RGB to SMPTE_C_D50 RGB colorspace. 
func CIEToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// CIEToSRGB_D50 converts from CIE RGB to SRGB_D50 RGB colorspace. 
func CIEToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CIEToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// ColorMatchToAdobe converts from ColorMatch RGB to Adobe RGB colorspace. 
func ColorMatchToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// ColorMatchToApple converts from ColorMatch RGB to Apple RGB colorspace. 
func ColorMatchToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// ColorMatchToBest converts from ColorMatch RGB to Best RGB colorspace. 
func ColorMatchToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// ColorMatchToBeta converts from ColorMatch RGB to Beta RGB colorspace. 
func ColorMatchToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// ColorMatchToBruce converts from ColorMatch RGB to Bruce RGB colorspace. 
func ColorMatchToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// ColorMatchToCIE converts from ColorMatch RGB to CIE RGB colorspace. 
func ColorMatchToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// ColorMatchToDon converts from ColorMatch RGB to Don RGB colorspace. 
func ColorMatchToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// ColorMatchToECI converts from ColorMatch RGB to ECI RGB colorspace. 
func ColorMatchToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// ColorMatchToEktaSpace converts from ColorMatch RGB to EktaSpace RGB colorspace. 
func ColorMatchToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// ColorMatchToNTSC converts from ColorMatch RGB to NTSC RGB colorspace. 
func ColorMatchToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// ColorMatchToPAL converts from ColorMatch RGB to PAL RGB colorspace. 
func ColorMatchToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// ColorMatchToProPhoto converts from ColorMatch RGB to ProPhoto RGB colorspace. 
func ColorMatchToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// ColorMatchToSMPTE_C converts from ColorMatch RGB to SMPTE_C RGB colorspace. 
func ColorMatchToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// ColorMatchToSRGB converts from ColorMatch RGB to SRGB RGB colorspace. 
func ColorMatchToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// ColorMatchToWGamut converts from ColorMatch RGB to WGamut RGB colorspace. 
func ColorMatchToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// ColorMatchToAdobe_D50 converts from ColorMatch RGB to Adobe_D50 RGB colorspace. 
func ColorMatchToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// ColorMatchToApple_D50 converts from ColorMatch RGB to Apple_D50 RGB colorspace. 
func ColorMatchToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// ColorMatchToBruce_D50 converts from ColorMatch RGB to Bruce_D50 RGB colorspace. 
func ColorMatchToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// ColorMatchToCie_D50 converts from ColorMatch RGB to Cie_D50 RGB colorspace. 
func ColorMatchToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// ColorMatchToNTSC_D50 converts from ColorMatch RGB to NTSC_D50 RGB colorspace. 
func ColorMatchToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// ColorMatchToPAL_D50 converts from ColorMatch RGB to PAL_D50 RGB colorspace. 
func ColorMatchToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// ColorMatchToSMPTE_C_D50 converts from ColorMatch RGB to SMPTE_C_D50 RGB colorspace. 
func ColorMatchToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// ColorMatchToSRGB_D50 converts from ColorMatch RGB to SRGB_D50 RGB colorspace. 
func ColorMatchToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ColorMatchToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// DonToAdobe converts from Don RGB to Adobe RGB colorspace. 
func DonToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// DonToApple converts from Don RGB to Apple RGB colorspace. 
func DonToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// DonToBest converts from Don RGB to Best RGB colorspace. 
func DonToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// DonToBeta converts from Don RGB to Beta RGB colorspace. 
func DonToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// DonToBruce converts from Don RGB to Bruce RGB colorspace. 
func DonToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// DonToCIE converts from Don RGB to CIE RGB colorspace. 
func DonToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// DonToColorMatch converts from Don RGB to ColorMatch RGB colorspace. 
func DonToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// DonToECI converts from Don RGB to ECI RGB colorspace. 
func DonToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// DonToEktaSpace converts from Don RGB to EktaSpace RGB colorspace. 
func DonToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// DonToNTSC converts from Don RGB to NTSC RGB colorspace. 
func DonToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// DonToPAL converts from Don RGB to PAL RGB colorspace. 
func DonToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// DonToProPhoto converts from Don RGB to ProPhoto RGB colorspace. 
func DonToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// DonToSMPTE_C converts from Don RGB to SMPTE_C RGB colorspace. 
func DonToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// DonToSRGB converts from Don RGB to SRGB RGB colorspace. 
func DonToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// DonToWGamut converts from Don RGB to WGamut RGB colorspace. 
func DonToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// DonToAdobe_D50 converts from Don RGB to Adobe_D50 RGB colorspace. 
func DonToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// DonToApple_D50 converts from Don RGB to Apple_D50 RGB colorspace. 
func DonToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// DonToBruce_D50 converts from Don RGB to Bruce_D50 RGB colorspace. 
func DonToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// DonToCie_D50 converts from Don RGB to Cie_D50 RGB colorspace. 
func DonToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// DonToNTSC_D50 converts from Don RGB to NTSC_D50 RGB colorspace. 
func DonToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// DonToPAL_D50 converts from Don RGB to PAL_D50 RGB colorspace. 
func DonToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// DonToSMPTE_C_D50 converts from Don RGB to SMPTE_C_D50 RGB colorspace. 
func DonToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// DonToSRGB_D50 converts from Don RGB to SRGB_D50 RGB colorspace. 
func DonToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  DonToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// ECIToAdobe converts from ECI RGB to Adobe RGB colorspace. 
func ECIToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// ECIToApple converts from ECI RGB to Apple RGB colorspace. 
func ECIToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// ECIToBest converts from ECI RGB to Best RGB colorspace. 
func ECIToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// ECIToBeta converts from ECI RGB to Beta RGB colorspace. 
func ECIToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// ECIToBruce converts from ECI RGB to Bruce RGB colorspace. 
func ECIToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// ECIToCIE converts from ECI RGB to CIE RGB colorspace. 
func ECIToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// ECIToColorMatch converts from ECI RGB to ColorMatch RGB colorspace. 
func ECIToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// ECIToDon converts from ECI RGB to Don RGB colorspace. 
func ECIToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// ECIToEktaSpace converts from ECI RGB to EktaSpace RGB colorspace. 
func ECIToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// ECIToNTSC converts from ECI RGB to NTSC RGB colorspace. 
func ECIToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// ECIToPAL converts from ECI RGB to PAL RGB colorspace. 
func ECIToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// ECIToProPhoto converts from ECI RGB to ProPhoto RGB colorspace. 
func ECIToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// ECIToSMPTE_C converts from ECI RGB to SMPTE_C RGB colorspace. 
func ECIToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// ECIToSRGB converts from ECI RGB to SRGB RGB colorspace. 
func ECIToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// ECIToWGamut converts from ECI RGB to WGamut RGB colorspace. 
func ECIToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// ECIToAdobe_D50 converts from ECI RGB to Adobe_D50 RGB colorspace. 
func ECIToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// ECIToApple_D50 converts from ECI RGB to Apple_D50 RGB colorspace. 
func ECIToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// ECIToBruce_D50 converts from ECI RGB to Bruce_D50 RGB colorspace. 
func ECIToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// ECIToCie_D50 converts from ECI RGB to Cie_D50 RGB colorspace. 
func ECIToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// ECIToNTSC_D50 converts from ECI RGB to NTSC_D50 RGB colorspace. 
func ECIToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// ECIToPAL_D50 converts from ECI RGB to PAL_D50 RGB colorspace. 
func ECIToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// ECIToSMPTE_C_D50 converts from ECI RGB to SMPTE_C_D50 RGB colorspace. 
func ECIToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// ECIToSRGB_D50 converts from ECI RGB to SRGB_D50 RGB colorspace. 
func ECIToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ECIToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// EktaSpaceToAdobe converts from EktaSpace RGB to Adobe RGB colorspace. 
func EktaSpaceToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// EktaSpaceToApple converts from EktaSpace RGB to Apple RGB colorspace. 
func EktaSpaceToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// EktaSpaceToBest converts from EktaSpace RGB to Best RGB colorspace. 
func EktaSpaceToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// EktaSpaceToBeta converts from EktaSpace RGB to Beta RGB colorspace. 
func EktaSpaceToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// EktaSpaceToBruce converts from EktaSpace RGB to Bruce RGB colorspace. 
func EktaSpaceToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// EktaSpaceToCIE converts from EktaSpace RGB to CIE RGB colorspace. 
func EktaSpaceToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// EktaSpaceToColorMatch converts from EktaSpace RGB to ColorMatch RGB colorspace. 
func EktaSpaceToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// EktaSpaceToDon converts from EktaSpace RGB to Don RGB colorspace. 
func EktaSpaceToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// EktaSpaceToECI converts from EktaSpace RGB to ECI RGB colorspace. 
func EktaSpaceToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// EktaSpaceToNTSC converts from EktaSpace RGB to NTSC RGB colorspace. 
func EktaSpaceToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// EktaSpaceToPAL converts from EktaSpace RGB to PAL RGB colorspace. 
func EktaSpaceToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// EktaSpaceToProPhoto converts from EktaSpace RGB to ProPhoto RGB colorspace. 
func EktaSpaceToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// EktaSpaceToSMPTE_C converts from EktaSpace RGB to SMPTE_C RGB colorspace. 
func EktaSpaceToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// EktaSpaceToSRGB converts from EktaSpace RGB to SRGB RGB colorspace. 
func EktaSpaceToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// EktaSpaceToWGamut converts from EktaSpace RGB to WGamut RGB colorspace. 
func EktaSpaceToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// EktaSpaceToAdobe_D50 converts from EktaSpace RGB to Adobe_D50 RGB colorspace. 
func EktaSpaceToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// EktaSpaceToApple_D50 converts from EktaSpace RGB to Apple_D50 RGB colorspace. 
func EktaSpaceToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// EktaSpaceToBruce_D50 converts from EktaSpace RGB to Bruce_D50 RGB colorspace. 
func EktaSpaceToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// EktaSpaceToCie_D50 converts from EktaSpace RGB to Cie_D50 RGB colorspace. 
func EktaSpaceToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// EktaSpaceToNTSC_D50 converts from EktaSpace RGB to NTSC_D50 RGB colorspace. 
func EktaSpaceToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// EktaSpaceToPAL_D50 converts from EktaSpace RGB to PAL_D50 RGB colorspace. 
func EktaSpaceToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// EktaSpaceToSMPTE_C_D50 converts from EktaSpace RGB to SMPTE_C_D50 RGB colorspace. 
func EktaSpaceToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// EktaSpaceToSRGB_D50 converts from EktaSpace RGB to SRGB_D50 RGB colorspace. 
func EktaSpaceToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  EktaSpaceToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// NTSCToAdobe converts from NTSC RGB to Adobe RGB colorspace. 
func NTSCToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// NTSCToApple converts from NTSC RGB to Apple RGB colorspace. 
func NTSCToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// NTSCToBest converts from NTSC RGB to Best RGB colorspace. 
func NTSCToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// NTSCToBeta converts from NTSC RGB to Beta RGB colorspace. 
func NTSCToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// NTSCToBruce converts from NTSC RGB to Bruce RGB colorspace. 
func NTSCToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// NTSCToCIE converts from NTSC RGB to CIE RGB colorspace. 
func NTSCToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// NTSCToColorMatch converts from NTSC RGB to ColorMatch RGB colorspace. 
func NTSCToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// NTSCToDon converts from NTSC RGB to Don RGB colorspace. 
func NTSCToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// NTSCToECI converts from NTSC RGB to ECI RGB colorspace. 
func NTSCToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// NTSCToEktaSpace converts from NTSC RGB to EktaSpace RGB colorspace. 
func NTSCToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// NTSCToPAL converts from NTSC RGB to PAL RGB colorspace. 
func NTSCToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// NTSCToProPhoto converts from NTSC RGB to ProPhoto RGB colorspace. 
func NTSCToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// NTSCToSMPTE_C converts from NTSC RGB to SMPTE_C RGB colorspace. 
func NTSCToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// NTSCToSRGB converts from NTSC RGB to SRGB RGB colorspace. 
func NTSCToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// NTSCToWGamut converts from NTSC RGB to WGamut RGB colorspace. 
func NTSCToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// NTSCToAdobe_D50 converts from NTSC RGB to Adobe_D50 RGB colorspace. 
func NTSCToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// NTSCToApple_D50 converts from NTSC RGB to Apple_D50 RGB colorspace. 
func NTSCToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// NTSCToBruce_D50 converts from NTSC RGB to Bruce_D50 RGB colorspace. 
func NTSCToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// NTSCToCie_D50 converts from NTSC RGB to Cie_D50 RGB colorspace. 
func NTSCToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// NTSCToNTSC_D50 converts from NTSC RGB to NTSC_D50 RGB colorspace. 
func NTSCToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// NTSCToPAL_D50 converts from NTSC RGB to PAL_D50 RGB colorspace. 
func NTSCToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// NTSCToSMPTE_C_D50 converts from NTSC RGB to SMPTE_C_D50 RGB colorspace. 
func NTSCToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// NTSCToSRGB_D50 converts from NTSC RGB to SRGB_D50 RGB colorspace. 
func NTSCToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// PALToAdobe converts from PAL RGB to Adobe RGB colorspace. 
func PALToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// PALToApple converts from PAL RGB to Apple RGB colorspace. 
func PALToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// PALToBest converts from PAL RGB to Best RGB colorspace. 
func PALToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// PALToBeta converts from PAL RGB to Beta RGB colorspace. 
func PALToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// PALToBruce converts from PAL RGB to Bruce RGB colorspace. 
func PALToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// PALToCIE converts from PAL RGB to CIE RGB colorspace. 
func PALToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// PALToColorMatch converts from PAL RGB to ColorMatch RGB colorspace. 
func PALToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// PALToDon converts from PAL RGB to Don RGB colorspace. 
func PALToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// PALToECI converts from PAL RGB to ECI RGB colorspace. 
func PALToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// PALToEktaSpace converts from PAL RGB to EktaSpace RGB colorspace. 
func PALToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// PALToNTSC converts from PAL RGB to NTSC RGB colorspace. 
func PALToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// PALToProPhoto converts from PAL RGB to ProPhoto RGB colorspace. 
func PALToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// PALToSMPTE_C converts from PAL RGB to SMPTE_C RGB colorspace. 
func PALToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// PALToSRGB converts from PAL RGB to SRGB RGB colorspace. 
func PALToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// PALToWGamut converts from PAL RGB to WGamut RGB colorspace. 
func PALToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// PALToAdobe_D50 converts from PAL RGB to Adobe_D50 RGB colorspace. 
func PALToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// PALToApple_D50 converts from PAL RGB to Apple_D50 RGB colorspace. 
func PALToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// PALToBruce_D50 converts from PAL RGB to Bruce_D50 RGB colorspace. 
func PALToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// PALToCie_D50 converts from PAL RGB to Cie_D50 RGB colorspace. 
func PALToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// PALToNTSC_D50 converts from PAL RGB to NTSC_D50 RGB colorspace. 
func PALToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// PALToPAL_D50 converts from PAL RGB to PAL_D50 RGB colorspace. 
func PALToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// PALToSMPTE_C_D50 converts from PAL RGB to SMPTE_C_D50 RGB colorspace. 
func PALToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// PALToSRGB_D50 converts from PAL RGB to SRGB_D50 RGB colorspace. 
func PALToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// ProPhotoToAdobe converts from ProPhoto RGB to Adobe RGB colorspace. 
func ProPhotoToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// ProPhotoToApple converts from ProPhoto RGB to Apple RGB colorspace. 
func ProPhotoToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// ProPhotoToBest converts from ProPhoto RGB to Best RGB colorspace. 
func ProPhotoToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// ProPhotoToBeta converts from ProPhoto RGB to Beta RGB colorspace. 
func ProPhotoToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// ProPhotoToBruce converts from ProPhoto RGB to Bruce RGB colorspace. 
func ProPhotoToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// ProPhotoToCIE converts from ProPhoto RGB to CIE RGB colorspace. 
func ProPhotoToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// ProPhotoToColorMatch converts from ProPhoto RGB to ColorMatch RGB colorspace. 
func ProPhotoToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// ProPhotoToDon converts from ProPhoto RGB to Don RGB colorspace. 
func ProPhotoToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// ProPhotoToECI converts from ProPhoto RGB to ECI RGB colorspace. 
func ProPhotoToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// ProPhotoToEktaSpace converts from ProPhoto RGB to EktaSpace RGB colorspace. 
func ProPhotoToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// ProPhotoToNTSC converts from ProPhoto RGB to NTSC RGB colorspace. 
func ProPhotoToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// ProPhotoToPAL converts from ProPhoto RGB to PAL RGB colorspace. 
func ProPhotoToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// ProPhotoToSMPTE_C converts from ProPhoto RGB to SMPTE_C RGB colorspace. 
func ProPhotoToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// ProPhotoToSRGB converts from ProPhoto RGB to SRGB RGB colorspace. 
func ProPhotoToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// ProPhotoToWGamut converts from ProPhoto RGB to WGamut RGB colorspace. 
func ProPhotoToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// ProPhotoToAdobe_D50 converts from ProPhoto RGB to Adobe_D50 RGB colorspace. 
func ProPhotoToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// ProPhotoToApple_D50 converts from ProPhoto RGB to Apple_D50 RGB colorspace. 
func ProPhotoToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// ProPhotoToBruce_D50 converts from ProPhoto RGB to Bruce_D50 RGB colorspace. 
func ProPhotoToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// ProPhotoToCie_D50 converts from ProPhoto RGB to Cie_D50 RGB colorspace. 
func ProPhotoToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// ProPhotoToNTSC_D50 converts from ProPhoto RGB to NTSC_D50 RGB colorspace. 
func ProPhotoToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// ProPhotoToPAL_D50 converts from ProPhoto RGB to PAL_D50 RGB colorspace. 
func ProPhotoToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// ProPhotoToSMPTE_C_D50 converts from ProPhoto RGB to SMPTE_C_D50 RGB colorspace. 
func ProPhotoToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// ProPhotoToSRGB_D50 converts from ProPhoto RGB to SRGB_D50 RGB colorspace. 
func ProPhotoToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  ProPhotoToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// SMPTE_CToAdobe converts from SMPTE_C RGB to Adobe RGB colorspace. 
func SMPTE_CToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// SMPTE_CToApple converts from SMPTE_C RGB to Apple RGB colorspace. 
func SMPTE_CToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// SMPTE_CToBest converts from SMPTE_C RGB to Best RGB colorspace. 
func SMPTE_CToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// SMPTE_CToBeta converts from SMPTE_C RGB to Beta RGB colorspace. 
func SMPTE_CToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// SMPTE_CToBruce converts from SMPTE_C RGB to Bruce RGB colorspace. 
func SMPTE_CToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// SMPTE_CToCIE converts from SMPTE_C RGB to CIE RGB colorspace. 
func SMPTE_CToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// SMPTE_CToColorMatch converts from SMPTE_C RGB to ColorMatch RGB colorspace. 
func SMPTE_CToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// SMPTE_CToDon converts from SMPTE_C RGB to Don RGB colorspace. 
func SMPTE_CToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// SMPTE_CToECI converts from SMPTE_C RGB to ECI RGB colorspace. 
func SMPTE_CToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// SMPTE_CToEktaSpace converts from SMPTE_C RGB to EktaSpace RGB colorspace. 
func SMPTE_CToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// SMPTE_CToNTSC converts from SMPTE_C RGB to NTSC RGB colorspace. 
func SMPTE_CToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// SMPTE_CToPAL converts from SMPTE_C RGB to PAL RGB colorspace. 
func SMPTE_CToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// SMPTE_CToProPhoto converts from SMPTE_C RGB to ProPhoto RGB colorspace. 
func SMPTE_CToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// SMPTE_CToSRGB converts from SMPTE_C RGB to SRGB RGB colorspace. 
func SMPTE_CToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// SMPTE_CToWGamut converts from SMPTE_C RGB to WGamut RGB colorspace. 
func SMPTE_CToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// SMPTE_CToAdobe_D50 converts from SMPTE_C RGB to Adobe_D50 RGB colorspace. 
func SMPTE_CToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// SMPTE_CToApple_D50 converts from SMPTE_C RGB to Apple_D50 RGB colorspace. 
func SMPTE_CToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// SMPTE_CToBruce_D50 converts from SMPTE_C RGB to Bruce_D50 RGB colorspace. 
func SMPTE_CToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// SMPTE_CToCie_D50 converts from SMPTE_C RGB to Cie_D50 RGB colorspace. 
func SMPTE_CToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// SMPTE_CToNTSC_D50 converts from SMPTE_C RGB to NTSC_D50 RGB colorspace. 
func SMPTE_CToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// SMPTE_CToPAL_D50 converts from SMPTE_C RGB to PAL_D50 RGB colorspace. 
func SMPTE_CToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// SMPTE_CToSMPTE_C_D50 converts from SMPTE_C RGB to SMPTE_C_D50 RGB colorspace. 
func SMPTE_CToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// SMPTE_CToSRGB_D50 converts from SMPTE_C RGB to SRGB_D50 RGB colorspace. 
func SMPTE_CToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// SRGBToAdobe converts from SRGB RGB to Adobe RGB colorspace. 
func SRGBToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// SRGBToApple converts from SRGB RGB to Apple RGB colorspace. 
func SRGBToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// SRGBToBest converts from SRGB RGB to Best RGB colorspace. 
func SRGBToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// SRGBToBeta converts from SRGB RGB to Beta RGB colorspace. 
func SRGBToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// SRGBToBruce converts from SRGB RGB to Bruce RGB colorspace. 
func SRGBToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// SRGBToCIE converts from SRGB RGB to CIE RGB colorspace. 
func SRGBToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// SRGBToColorMatch converts from SRGB RGB to ColorMatch RGB colorspace. 
func SRGBToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// SRGBToDon converts from SRGB RGB to Don RGB colorspace. 
func SRGBToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// SRGBToECI converts from SRGB RGB to ECI RGB colorspace. 
func SRGBToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// SRGBToEktaSpace converts from SRGB RGB to EktaSpace RGB colorspace. 
func SRGBToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// SRGBToNTSC converts from SRGB RGB to NTSC RGB colorspace. 
func SRGBToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// SRGBToPAL converts from SRGB RGB to PAL RGB colorspace. 
func SRGBToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// SRGBToProPhoto converts from SRGB RGB to ProPhoto RGB colorspace. 
func SRGBToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// SRGBToSMPTE_C converts from SRGB RGB to SMPTE_C RGB colorspace. 
func SRGBToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// SRGBToWGamut converts from SRGB RGB to WGamut RGB colorspace. 
func SRGBToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// SRGBToAdobe_D50 converts from SRGB RGB to Adobe_D50 RGB colorspace. 
func SRGBToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// SRGBToApple_D50 converts from SRGB RGB to Apple_D50 RGB colorspace. 
func SRGBToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// SRGBToBruce_D50 converts from SRGB RGB to Bruce_D50 RGB colorspace. 
func SRGBToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// SRGBToCie_D50 converts from SRGB RGB to Cie_D50 RGB colorspace. 
func SRGBToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// SRGBToNTSC_D50 converts from SRGB RGB to NTSC_D50 RGB colorspace. 
func SRGBToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// SRGBToPAL_D50 converts from SRGB RGB to PAL_D50 RGB colorspace. 
func SRGBToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// SRGBToSMPTE_C_D50 converts from SRGB RGB to SMPTE_C_D50 RGB colorspace. 
func SRGBToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// SRGBToSRGB_D50 converts from SRGB RGB to SRGB_D50 RGB colorspace. 
func SRGBToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// WGamutToAdobe converts from WGamut RGB to Adobe RGB colorspace. 
func WGamutToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// WGamutToApple converts from WGamut RGB to Apple RGB colorspace. 
func WGamutToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// WGamutToBest converts from WGamut RGB to Best RGB colorspace. 
func WGamutToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// WGamutToBeta converts from WGamut RGB to Beta RGB colorspace. 
func WGamutToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// WGamutToBruce converts from WGamut RGB to Bruce RGB colorspace. 
func WGamutToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// WGamutToCIE converts from WGamut RGB to CIE RGB colorspace. 
func WGamutToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// WGamutToColorMatch converts from WGamut RGB to ColorMatch RGB colorspace. 
func WGamutToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// WGamutToDon converts from WGamut RGB to Don RGB colorspace. 
func WGamutToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// WGamutToECI converts from WGamut RGB to ECI RGB colorspace. 
func WGamutToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// WGamutToEktaSpace converts from WGamut RGB to EktaSpace RGB colorspace. 
func WGamutToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// WGamutToNTSC converts from WGamut RGB to NTSC RGB colorspace. 
func WGamutToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// WGamutToPAL converts from WGamut RGB to PAL RGB colorspace. 
func WGamutToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// WGamutToProPhoto converts from WGamut RGB to ProPhoto RGB colorspace. 
func WGamutToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// WGamutToSMPTE_C converts from WGamut RGB to SMPTE_C RGB colorspace. 
func WGamutToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// WGamutToSRGB converts from WGamut RGB to SRGB RGB colorspace. 
func WGamutToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// WGamutToAdobe_D50 converts from WGamut RGB to Adobe_D50 RGB colorspace. 
func WGamutToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// WGamutToApple_D50 converts from WGamut RGB to Apple_D50 RGB colorspace. 
func WGamutToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// WGamutToBruce_D50 converts from WGamut RGB to Bruce_D50 RGB colorspace. 
func WGamutToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// WGamutToCie_D50 converts from WGamut RGB to Cie_D50 RGB colorspace. 
func WGamutToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// WGamutToNTSC_D50 converts from WGamut RGB to NTSC_D50 RGB colorspace. 
func WGamutToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// WGamutToPAL_D50 converts from WGamut RGB to PAL_D50 RGB colorspace. 
func WGamutToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// WGamutToSMPTE_C_D50 converts from WGamut RGB to SMPTE_C_D50 RGB colorspace. 
func WGamutToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// WGamutToSRGB_D50 converts from WGamut RGB to SRGB_D50 RGB colorspace. 
func WGamutToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  WGamutToXYZ(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// Adobe_D50ToAdobe converts from Adobe_D50 RGB to Adobe RGB colorspace. 
func Adobe_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// Adobe_D50ToApple converts from Adobe_D50 RGB to Apple RGB colorspace. 
func Adobe_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// Adobe_D50ToBest converts from Adobe_D50 RGB to Best RGB colorspace. 
func Adobe_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// Adobe_D50ToBeta converts from Adobe_D50 RGB to Beta RGB colorspace. 
func Adobe_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// Adobe_D50ToBruce converts from Adobe_D50 RGB to Bruce RGB colorspace. 
func Adobe_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// Adobe_D50ToCIE converts from Adobe_D50 RGB to CIE RGB colorspace. 
func Adobe_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// Adobe_D50ToColorMatch converts from Adobe_D50 RGB to ColorMatch RGB colorspace. 
func Adobe_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// Adobe_D50ToDon converts from Adobe_D50 RGB to Don RGB colorspace. 
func Adobe_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// Adobe_D50ToECI converts from Adobe_D50 RGB to ECI RGB colorspace. 
func Adobe_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// Adobe_D50ToEktaSpace converts from Adobe_D50 RGB to EktaSpace RGB colorspace. 
func Adobe_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// Adobe_D50ToNTSC converts from Adobe_D50 RGB to NTSC RGB colorspace. 
func Adobe_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// Adobe_D50ToPAL converts from Adobe_D50 RGB to PAL RGB colorspace. 
func Adobe_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// Adobe_D50ToProPhoto converts from Adobe_D50 RGB to ProPhoto RGB colorspace. 
func Adobe_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// Adobe_D50ToSMPTE_C converts from Adobe_D50 RGB to SMPTE_C RGB colorspace. 
func Adobe_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// Adobe_D50ToSRGB converts from Adobe_D50 RGB to SRGB RGB colorspace. 
func Adobe_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// Adobe_D50ToWGamut converts from Adobe_D50 RGB to WGamut RGB colorspace. 
func Adobe_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// Adobe_D50ToApple_D50 converts from Adobe_D50 RGB to Apple_D50 RGB colorspace. 
func Adobe_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// Adobe_D50ToBruce_D50 converts from Adobe_D50 RGB to Bruce_D50 RGB colorspace. 
func Adobe_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// Adobe_D50ToCie_D50 converts from Adobe_D50 RGB to Cie_D50 RGB colorspace. 
func Adobe_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// Adobe_D50ToNTSC_D50 converts from Adobe_D50 RGB to NTSC_D50 RGB colorspace. 
func Adobe_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// Adobe_D50ToPAL_D50 converts from Adobe_D50 RGB to PAL_D50 RGB colorspace. 
func Adobe_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// Adobe_D50ToSMPTE_C_D50 converts from Adobe_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func Adobe_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// Adobe_D50ToSRGB_D50 converts from Adobe_D50 RGB to SRGB_D50 RGB colorspace. 
func Adobe_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AdobeToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// Apple_D50ToAdobe converts from Apple_D50 RGB to Adobe RGB colorspace. 
func Apple_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// Apple_D50ToApple converts from Apple_D50 RGB to Apple RGB colorspace. 
func Apple_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// Apple_D50ToBest converts from Apple_D50 RGB to Best RGB colorspace. 
func Apple_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// Apple_D50ToBeta converts from Apple_D50 RGB to Beta RGB colorspace. 
func Apple_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// Apple_D50ToBruce converts from Apple_D50 RGB to Bruce RGB colorspace. 
func Apple_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// Apple_D50ToCIE converts from Apple_D50 RGB to CIE RGB colorspace. 
func Apple_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// Apple_D50ToColorMatch converts from Apple_D50 RGB to ColorMatch RGB colorspace. 
func Apple_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// Apple_D50ToDon converts from Apple_D50 RGB to Don RGB colorspace. 
func Apple_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// Apple_D50ToECI converts from Apple_D50 RGB to ECI RGB colorspace. 
func Apple_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// Apple_D50ToEktaSpace converts from Apple_D50 RGB to EktaSpace RGB colorspace. 
func Apple_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// Apple_D50ToNTSC converts from Apple_D50 RGB to NTSC RGB colorspace. 
func Apple_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// Apple_D50ToPAL converts from Apple_D50 RGB to PAL RGB colorspace. 
func Apple_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// Apple_D50ToProPhoto converts from Apple_D50 RGB to ProPhoto RGB colorspace. 
func Apple_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// Apple_D50ToSMPTE_C converts from Apple_D50 RGB to SMPTE_C RGB colorspace. 
func Apple_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// Apple_D50ToSRGB converts from Apple_D50 RGB to SRGB RGB colorspace. 
func Apple_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// Apple_D50ToWGamut converts from Apple_D50 RGB to WGamut RGB colorspace. 
func Apple_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// Apple_D50ToAdobe_D50 converts from Apple_D50 RGB to Adobe_D50 RGB colorspace. 
func Apple_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// Apple_D50ToBruce_D50 converts from Apple_D50 RGB to Bruce_D50 RGB colorspace. 
func Apple_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// Apple_D50ToCie_D50 converts from Apple_D50 RGB to Cie_D50 RGB colorspace. 
func Apple_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// Apple_D50ToNTSC_D50 converts from Apple_D50 RGB to NTSC_D50 RGB colorspace. 
func Apple_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// Apple_D50ToPAL_D50 converts from Apple_D50 RGB to PAL_D50 RGB colorspace. 
func Apple_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// Apple_D50ToSMPTE_C_D50 converts from Apple_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func Apple_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// Apple_D50ToSRGB_D50 converts from Apple_D50 RGB to SRGB_D50 RGB colorspace. 
func Apple_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  AppleToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// Bruce_D50ToAdobe converts from Bruce_D50 RGB to Adobe RGB colorspace. 
func Bruce_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// Bruce_D50ToApple converts from Bruce_D50 RGB to Apple RGB colorspace. 
func Bruce_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// Bruce_D50ToBest converts from Bruce_D50 RGB to Best RGB colorspace. 
func Bruce_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// Bruce_D50ToBeta converts from Bruce_D50 RGB to Beta RGB colorspace. 
func Bruce_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// Bruce_D50ToBruce converts from Bruce_D50 RGB to Bruce RGB colorspace. 
func Bruce_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// Bruce_D50ToCIE converts from Bruce_D50 RGB to CIE RGB colorspace. 
func Bruce_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// Bruce_D50ToColorMatch converts from Bruce_D50 RGB to ColorMatch RGB colorspace. 
func Bruce_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// Bruce_D50ToDon converts from Bruce_D50 RGB to Don RGB colorspace. 
func Bruce_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// Bruce_D50ToECI converts from Bruce_D50 RGB to ECI RGB colorspace. 
func Bruce_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// Bruce_D50ToEktaSpace converts from Bruce_D50 RGB to EktaSpace RGB colorspace. 
func Bruce_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// Bruce_D50ToNTSC converts from Bruce_D50 RGB to NTSC RGB colorspace. 
func Bruce_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// Bruce_D50ToPAL converts from Bruce_D50 RGB to PAL RGB colorspace. 
func Bruce_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// Bruce_D50ToProPhoto converts from Bruce_D50 RGB to ProPhoto RGB colorspace. 
func Bruce_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// Bruce_D50ToSMPTE_C converts from Bruce_D50 RGB to SMPTE_C RGB colorspace. 
func Bruce_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// Bruce_D50ToSRGB converts from Bruce_D50 RGB to SRGB RGB colorspace. 
func Bruce_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// Bruce_D50ToWGamut converts from Bruce_D50 RGB to WGamut RGB colorspace. 
func Bruce_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// Bruce_D50ToAdobe_D50 converts from Bruce_D50 RGB to Adobe_D50 RGB colorspace. 
func Bruce_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// Bruce_D50ToApple_D50 converts from Bruce_D50 RGB to Apple_D50 RGB colorspace. 
func Bruce_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// Bruce_D50ToCie_D50 converts from Bruce_D50 RGB to Cie_D50 RGB colorspace. 
func Bruce_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// Bruce_D50ToNTSC_D50 converts from Bruce_D50 RGB to NTSC_D50 RGB colorspace. 
func Bruce_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// Bruce_D50ToPAL_D50 converts from Bruce_D50 RGB to PAL_D50 RGB colorspace. 
func Bruce_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// Bruce_D50ToSMPTE_C_D50 converts from Bruce_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func Bruce_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// Bruce_D50ToSRGB_D50 converts from Bruce_D50 RGB to SRGB_D50 RGB colorspace. 
func Bruce_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  BruceToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// Cie_D50ToAdobe converts from Cie_D50 RGB to Adobe RGB colorspace. 
func Cie_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// Cie_D50ToApple converts from Cie_D50 RGB to Apple RGB colorspace. 
func Cie_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// Cie_D50ToBest converts from Cie_D50 RGB to Best RGB colorspace. 
func Cie_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// Cie_D50ToBeta converts from Cie_D50 RGB to Beta RGB colorspace. 
func Cie_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// Cie_D50ToBruce converts from Cie_D50 RGB to Bruce RGB colorspace. 
func Cie_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// Cie_D50ToCIE converts from Cie_D50 RGB to CIE RGB colorspace. 
func Cie_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// Cie_D50ToColorMatch converts from Cie_D50 RGB to ColorMatch RGB colorspace. 
func Cie_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// Cie_D50ToDon converts from Cie_D50 RGB to Don RGB colorspace. 
func Cie_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// Cie_D50ToECI converts from Cie_D50 RGB to ECI RGB colorspace. 
func Cie_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// Cie_D50ToEktaSpace converts from Cie_D50 RGB to EktaSpace RGB colorspace. 
func Cie_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// Cie_D50ToNTSC converts from Cie_D50 RGB to NTSC RGB colorspace. 
func Cie_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// Cie_D50ToPAL converts from Cie_D50 RGB to PAL RGB colorspace. 
func Cie_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// Cie_D50ToProPhoto converts from Cie_D50 RGB to ProPhoto RGB colorspace. 
func Cie_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// Cie_D50ToSMPTE_C converts from Cie_D50 RGB to SMPTE_C RGB colorspace. 
func Cie_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// Cie_D50ToSRGB converts from Cie_D50 RGB to SRGB RGB colorspace. 
func Cie_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// Cie_D50ToWGamut converts from Cie_D50 RGB to WGamut RGB colorspace. 
func Cie_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// Cie_D50ToAdobe_D50 converts from Cie_D50 RGB to Adobe_D50 RGB colorspace. 
func Cie_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// Cie_D50ToApple_D50 converts from Cie_D50 RGB to Apple_D50 RGB colorspace. 
func Cie_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// Cie_D50ToBruce_D50 converts from Cie_D50 RGB to Bruce_D50 RGB colorspace. 
func Cie_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// Cie_D50ToNTSC_D50 converts from Cie_D50 RGB to NTSC_D50 RGB colorspace. 
func Cie_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// Cie_D50ToPAL_D50 converts from Cie_D50 RGB to PAL_D50 RGB colorspace. 
func Cie_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// Cie_D50ToSMPTE_C_D50 converts from Cie_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func Cie_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// Cie_D50ToSRGB_D50 converts from Cie_D50 RGB to SRGB_D50 RGB colorspace. 
func Cie_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  CieToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// NTSC_D50ToAdobe converts from NTSC_D50 RGB to Adobe RGB colorspace. 
func NTSC_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// NTSC_D50ToApple converts from NTSC_D50 RGB to Apple RGB colorspace. 
func NTSC_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// NTSC_D50ToBest converts from NTSC_D50 RGB to Best RGB colorspace. 
func NTSC_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// NTSC_D50ToBeta converts from NTSC_D50 RGB to Beta RGB colorspace. 
func NTSC_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// NTSC_D50ToBruce converts from NTSC_D50 RGB to Bruce RGB colorspace. 
func NTSC_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// NTSC_D50ToCIE converts from NTSC_D50 RGB to CIE RGB colorspace. 
func NTSC_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// NTSC_D50ToColorMatch converts from NTSC_D50 RGB to ColorMatch RGB colorspace. 
func NTSC_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// NTSC_D50ToDon converts from NTSC_D50 RGB to Don RGB colorspace. 
func NTSC_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// NTSC_D50ToECI converts from NTSC_D50 RGB to ECI RGB colorspace. 
func NTSC_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// NTSC_D50ToEktaSpace converts from NTSC_D50 RGB to EktaSpace RGB colorspace. 
func NTSC_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// NTSC_D50ToNTSC converts from NTSC_D50 RGB to NTSC RGB colorspace. 
func NTSC_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// NTSC_D50ToPAL converts from NTSC_D50 RGB to PAL RGB colorspace. 
func NTSC_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// NTSC_D50ToProPhoto converts from NTSC_D50 RGB to ProPhoto RGB colorspace. 
func NTSC_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// NTSC_D50ToSMPTE_C converts from NTSC_D50 RGB to SMPTE_C RGB colorspace. 
func NTSC_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// NTSC_D50ToSRGB converts from NTSC_D50 RGB to SRGB RGB colorspace. 
func NTSC_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// NTSC_D50ToWGamut converts from NTSC_D50 RGB to WGamut RGB colorspace. 
func NTSC_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// NTSC_D50ToAdobe_D50 converts from NTSC_D50 RGB to Adobe_D50 RGB colorspace. 
func NTSC_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// NTSC_D50ToApple_D50 converts from NTSC_D50 RGB to Apple_D50 RGB colorspace. 
func NTSC_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// NTSC_D50ToBruce_D50 converts from NTSC_D50 RGB to Bruce_D50 RGB colorspace. 
func NTSC_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// NTSC_D50ToCie_D50 converts from NTSC_D50 RGB to Cie_D50 RGB colorspace. 
func NTSC_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// NTSC_D50ToPAL_D50 converts from NTSC_D50 RGB to PAL_D50 RGB colorspace. 
func NTSC_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// NTSC_D50ToSMPTE_C_D50 converts from NTSC_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func NTSC_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// NTSC_D50ToSRGB_D50 converts from NTSC_D50 RGB to SRGB_D50 RGB colorspace. 
func NTSC_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  NTSCToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// PAL_D50ToAdobe converts from PAL_D50 RGB to Adobe RGB colorspace. 
func PAL_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// PAL_D50ToApple converts from PAL_D50 RGB to Apple RGB colorspace. 
func PAL_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// PAL_D50ToBest converts from PAL_D50 RGB to Best RGB colorspace. 
func PAL_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// PAL_D50ToBeta converts from PAL_D50 RGB to Beta RGB colorspace. 
func PAL_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// PAL_D50ToBruce converts from PAL_D50 RGB to Bruce RGB colorspace. 
func PAL_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// PAL_D50ToCIE converts from PAL_D50 RGB to CIE RGB colorspace. 
func PAL_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// PAL_D50ToColorMatch converts from PAL_D50 RGB to ColorMatch RGB colorspace. 
func PAL_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// PAL_D50ToDon converts from PAL_D50 RGB to Don RGB colorspace. 
func PAL_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// PAL_D50ToECI converts from PAL_D50 RGB to ECI RGB colorspace. 
func PAL_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// PAL_D50ToEktaSpace converts from PAL_D50 RGB to EktaSpace RGB colorspace. 
func PAL_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// PAL_D50ToNTSC converts from PAL_D50 RGB to NTSC RGB colorspace. 
func PAL_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// PAL_D50ToPAL converts from PAL_D50 RGB to PAL RGB colorspace. 
func PAL_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// PAL_D50ToProPhoto converts from PAL_D50 RGB to ProPhoto RGB colorspace. 
func PAL_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// PAL_D50ToSMPTE_C converts from PAL_D50 RGB to SMPTE_C RGB colorspace. 
func PAL_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// PAL_D50ToSRGB converts from PAL_D50 RGB to SRGB RGB colorspace. 
func PAL_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// PAL_D50ToWGamut converts from PAL_D50 RGB to WGamut RGB colorspace. 
func PAL_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// PAL_D50ToAdobe_D50 converts from PAL_D50 RGB to Adobe_D50 RGB colorspace. 
func PAL_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// PAL_D50ToApple_D50 converts from PAL_D50 RGB to Apple_D50 RGB colorspace. 
func PAL_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// PAL_D50ToBruce_D50 converts from PAL_D50 RGB to Bruce_D50 RGB colorspace. 
func PAL_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// PAL_D50ToCie_D50 converts from PAL_D50 RGB to Cie_D50 RGB colorspace. 
func PAL_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// PAL_D50ToNTSC_D50 converts from PAL_D50 RGB to NTSC_D50 RGB colorspace. 
func PAL_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// PAL_D50ToSMPTE_C_D50 converts from PAL_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func PAL_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

// PAL_D50ToSRGB_D50 converts from PAL_D50 RGB to SRGB_D50 RGB colorspace. 
func PAL_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  PALToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// SMPTE_C_D50ToAdobe converts from SMPTE_C_D50 RGB to Adobe RGB colorspace. 
func SMPTE_C_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// SMPTE_C_D50ToApple converts from SMPTE_C_D50 RGB to Apple RGB colorspace. 
func SMPTE_C_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// SMPTE_C_D50ToBest converts from SMPTE_C_D50 RGB to Best RGB colorspace. 
func SMPTE_C_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// SMPTE_C_D50ToBeta converts from SMPTE_C_D50 RGB to Beta RGB colorspace. 
func SMPTE_C_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// SMPTE_C_D50ToBruce converts from SMPTE_C_D50 RGB to Bruce RGB colorspace. 
func SMPTE_C_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// SMPTE_C_D50ToCIE converts from SMPTE_C_D50 RGB to CIE RGB colorspace. 
func SMPTE_C_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// SMPTE_C_D50ToColorMatch converts from SMPTE_C_D50 RGB to ColorMatch RGB colorspace. 
func SMPTE_C_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// SMPTE_C_D50ToDon converts from SMPTE_C_D50 RGB to Don RGB colorspace. 
func SMPTE_C_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// SMPTE_C_D50ToECI converts from SMPTE_C_D50 RGB to ECI RGB colorspace. 
func SMPTE_C_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// SMPTE_C_D50ToEktaSpace converts from SMPTE_C_D50 RGB to EktaSpace RGB colorspace. 
func SMPTE_C_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// SMPTE_C_D50ToNTSC converts from SMPTE_C_D50 RGB to NTSC RGB colorspace. 
func SMPTE_C_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// SMPTE_C_D50ToPAL converts from SMPTE_C_D50 RGB to PAL RGB colorspace. 
func SMPTE_C_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// SMPTE_C_D50ToProPhoto converts from SMPTE_C_D50 RGB to ProPhoto RGB colorspace. 
func SMPTE_C_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// SMPTE_C_D50ToSMPTE_C converts from SMPTE_C_D50 RGB to SMPTE_C RGB colorspace. 
func SMPTE_C_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// SMPTE_C_D50ToSRGB converts from SMPTE_C_D50 RGB to SRGB RGB colorspace. 
func SMPTE_C_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// SMPTE_C_D50ToWGamut converts from SMPTE_C_D50 RGB to WGamut RGB colorspace. 
func SMPTE_C_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// SMPTE_C_D50ToAdobe_D50 converts from SMPTE_C_D50 RGB to Adobe_D50 RGB colorspace. 
func SMPTE_C_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// SMPTE_C_D50ToApple_D50 converts from SMPTE_C_D50 RGB to Apple_D50 RGB colorspace. 
func SMPTE_C_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// SMPTE_C_D50ToBruce_D50 converts from SMPTE_C_D50 RGB to Bruce_D50 RGB colorspace. 
func SMPTE_C_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// SMPTE_C_D50ToCie_D50 converts from SMPTE_C_D50 RGB to Cie_D50 RGB colorspace. 
func SMPTE_C_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// SMPTE_C_D50ToNTSC_D50 converts from SMPTE_C_D50 RGB to NTSC_D50 RGB colorspace. 
func SMPTE_C_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// SMPTE_C_D50ToPAL_D50 converts from SMPTE_C_D50 RGB to PAL_D50 RGB colorspace. 
func SMPTE_C_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// SMPTE_C_D50ToSRGB_D50 converts from SMPTE_C_D50 RGB to SRGB_D50 RGB colorspace. 
func SMPTE_C_D50ToSRGB_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SMPTE_CToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB_D50(x, y, z)
	return
}

// SRGB_D50ToAdobe converts from SRGB_D50 RGB to Adobe RGB colorspace. 
func SRGB_D50ToAdobe(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe(x, y, z)
	return
}

// SRGB_D50ToApple converts from SRGB_D50 RGB to Apple RGB colorspace. 
func SRGB_D50ToApple(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple(x, y, z)
	return
}

// SRGB_D50ToBest converts from SRGB_D50 RGB to Best RGB colorspace. 
func SRGB_D50ToBest(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBest(x, y, z)
	return
}

// SRGB_D50ToBeta converts from SRGB_D50 RGB to Beta RGB colorspace. 
func SRGB_D50ToBeta(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBeta(x, y, z)
	return
}

// SRGB_D50ToBruce converts from SRGB_D50 RGB to Bruce RGB colorspace. 
func SRGB_D50ToBruce(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce(x, y, z)
	return
}

// SRGB_D50ToCIE converts from SRGB_D50 RGB to CIE RGB colorspace. 
func SRGB_D50ToCIE(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCIE(x, y, z)
	return
}

// SRGB_D50ToColorMatch converts from SRGB_D50 RGB to ColorMatch RGB colorspace. 
func SRGB_D50ToColorMatch(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToColorMatch(x, y, z)
	return
}

// SRGB_D50ToDon converts from SRGB_D50 RGB to Don RGB colorspace. 
func SRGB_D50ToDon(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToDon(x, y, z)
	return
}

// SRGB_D50ToECI converts from SRGB_D50 RGB to ECI RGB colorspace. 
func SRGB_D50ToECI(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToECI(x, y, z)
	return
}

// SRGB_D50ToEktaSpace converts from SRGB_D50 RGB to EktaSpace RGB colorspace. 
func SRGB_D50ToEktaSpace(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToEktaSpace(x, y, z)
	return
}

// SRGB_D50ToNTSC converts from SRGB_D50 RGB to NTSC RGB colorspace. 
func SRGB_D50ToNTSC(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC(x, y, z)
	return
}

// SRGB_D50ToPAL converts from SRGB_D50 RGB to PAL RGB colorspace. 
func SRGB_D50ToPAL(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL(x, y, z)
	return
}

// SRGB_D50ToProPhoto converts from SRGB_D50 RGB to ProPhoto RGB colorspace. 
func SRGB_D50ToProPhoto(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToProPhoto(x, y, z)
	return
}

// SRGB_D50ToSMPTE_C converts from SRGB_D50 RGB to SMPTE_C RGB colorspace. 
func SRGB_D50ToSMPTE_C(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C(x, y, z)
	return
}

// SRGB_D50ToSRGB converts from SRGB_D50 RGB to SRGB RGB colorspace. 
func SRGB_D50ToSRGB(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSRGB(x, y, z)
	return
}

// SRGB_D50ToWGamut converts from SRGB_D50 RGB to WGamut RGB colorspace. 
func SRGB_D50ToWGamut(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToWGamut(x, y, z)
	return
}

// SRGB_D50ToAdobe_D50 converts from SRGB_D50 RGB to Adobe_D50 RGB colorspace. 
func SRGB_D50ToAdobe_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToAdobe_D50(x, y, z)
	return
}

// SRGB_D50ToApple_D50 converts from SRGB_D50 RGB to Apple_D50 RGB colorspace. 
func SRGB_D50ToApple_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToApple_D50(x, y, z)
	return
}

// SRGB_D50ToBruce_D50 converts from SRGB_D50 RGB to Bruce_D50 RGB colorspace. 
func SRGB_D50ToBruce_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToBruce_D50(x, y, z)
	return
}

// SRGB_D50ToCie_D50 converts from SRGB_D50 RGB to Cie_D50 RGB colorspace. 
func SRGB_D50ToCie_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToCie_D50(x, y, z)
	return
}

// SRGB_D50ToNTSC_D50 converts from SRGB_D50 RGB to NTSC_D50 RGB colorspace. 
func SRGB_D50ToNTSC_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToNTSC_D50(x, y, z)
	return
}

// SRGB_D50ToPAL_D50 converts from SRGB_D50 RGB to PAL_D50 RGB colorspace. 
func SRGB_D50ToPAL_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToPAL_D50(x, y, z)
	return
}

// SRGB_D50ToSMPTE_C_D50 converts from SRGB_D50 RGB to SMPTE_C_D50 RGB colorspace. 
func SRGB_D50ToSMPTE_C_D50(r, g, b float64) (r1, g1, b1 float64) {
	x, y, z :=  SRGBToXYZ_D50(r, g, b)
	r1, g1, b1 = XYZToSMPTE_C_D50(x, y, z)
	return
}

