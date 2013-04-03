// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Conversion of different RGB colorspaces to CIE-XYZ and back. 16 bits-per-channel.
package rgb16

// Intended for conversion between RGB-colorspaces.
// Supported RGB colorspaces and their reference whites:
// Adobe RGB (1998) D65
// Apple RGB D65
// Best RGB D50
// Beta RGB D50
// Bruce RGB D65
// CIE RGB E
// ColorMatch RGB D50
// Don RGB 4 D50
// ECI RGB D50
// Ekta Space PS5 D50
// NTSC RGB C
// PAL/SECAM RGB D65
// ProPhoto RGB D50
// SMPTE-C RGB D65
// sRGB D65
// Wide Gamut RGB D50
// For colorspaces whose native reference white is not D50, Bradford adaptation was used:
// Adobe RGB (1998) D50
// Apple RGB D50
// Bruce RGB D50
// CIE RGB D50
// NTSC RGB D50
// PAL/SECAM RGB D50
// SMPTE-C RGB D50
// sRGB D50
