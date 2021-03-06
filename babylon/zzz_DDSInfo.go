// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DDSInfo represents a babylon.js DDSInfo.
// Direct draw surface info
//
// See: https://docs.microsoft.com/en-us/windows/desktop/direct3ddds/dx-graphics-dds-pguide
type DDSInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DDSInfo) JSObject() js.Value { return d.p }

// DDSInfo returns a DDSInfo JavaScript class.
func (ba *Babylon) DDSInfo() *DDSInfo {
	p := ba.ctx.Get("DDSInfo")
	return DDSInfoFromJSObject(p, ba.ctx)
}

// DDSInfoFromJSObject returns a wrapped DDSInfo JavaScript class.
func DDSInfoFromJSObject(p js.Value, ctx js.Value) *DDSInfo {
	return &DDSInfo{p: p, ctx: ctx}
}

// DDSInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func DDSInfoArrayToJSArray(array []*DDSInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// DxgiFormat returns the DxgiFormat property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#dxgiformat
func (d *DDSInfo) DxgiFormat() float64 {
	retVal := d.p.Get("dxgiFormat")
	return retVal.Float()
}

// SetDxgiFormat sets the DxgiFormat property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#dxgiformat
func (d *DDSInfo) SetDxgiFormat(dxgiFormat float64) *DDSInfo {
	d.p.Set("dxgiFormat", dxgiFormat)
	return d
}

// Height returns the Height property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#height
func (d *DDSInfo) Height() float64 {
	retVal := d.p.Get("height")
	return retVal.Float()
}

// SetHeight sets the Height property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#height
func (d *DDSInfo) SetHeight(height float64) *DDSInfo {
	d.p.Set("height", height)
	return d
}

// IsCompressed returns the IsCompressed property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#iscompressed
func (d *DDSInfo) IsCompressed() bool {
	retVal := d.p.Get("isCompressed")
	return retVal.Bool()
}

// SetIsCompressed sets the IsCompressed property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#iscompressed
func (d *DDSInfo) SetIsCompressed(isCompressed bool) *DDSInfo {
	d.p.Set("isCompressed", isCompressed)
	return d
}

// IsCube returns the IsCube property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#iscube
func (d *DDSInfo) IsCube() bool {
	retVal := d.p.Get("isCube")
	return retVal.Bool()
}

// SetIsCube sets the IsCube property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#iscube
func (d *DDSInfo) SetIsCube(isCube bool) *DDSInfo {
	d.p.Set("isCube", isCube)
	return d
}

// IsFourCC returns the IsFourCC property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isfourcc
func (d *DDSInfo) IsFourCC() bool {
	retVal := d.p.Get("isFourCC")
	return retVal.Bool()
}

// SetIsFourCC sets the IsFourCC property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isfourcc
func (d *DDSInfo) SetIsFourCC(isFourCC bool) *DDSInfo {
	d.p.Set("isFourCC", isFourCC)
	return d
}

// IsLuminance returns the IsLuminance property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isluminance
func (d *DDSInfo) IsLuminance() bool {
	retVal := d.p.Get("isLuminance")
	return retVal.Bool()
}

// SetIsLuminance sets the IsLuminance property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isluminance
func (d *DDSInfo) SetIsLuminance(isLuminance bool) *DDSInfo {
	d.p.Set("isLuminance", isLuminance)
	return d
}

// IsRGB returns the IsRGB property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isrgb
func (d *DDSInfo) IsRGB() bool {
	retVal := d.p.Get("isRGB")
	return retVal.Bool()
}

// SetIsRGB sets the IsRGB property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#isrgb
func (d *DDSInfo) SetIsRGB(isRGB bool) *DDSInfo {
	d.p.Set("isRGB", isRGB)
	return d
}

// MipmapCount returns the MipmapCount property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#mipmapcount
func (d *DDSInfo) MipmapCount() float64 {
	retVal := d.p.Get("mipmapCount")
	return retVal.Float()
}

// SetMipmapCount sets the MipmapCount property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#mipmapcount
func (d *DDSInfo) SetMipmapCount(mipmapCount float64) *DDSInfo {
	d.p.Set("mipmapCount", mipmapCount)
	return d
}

// SphericalPolynomial returns the SphericalPolynomial property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#sphericalpolynomial
func (d *DDSInfo) SphericalPolynomial() *SphericalPolynomial {
	retVal := d.p.Get("sphericalPolynomial")
	return SphericalPolynomialFromJSObject(retVal, d.ctx)
}

// SetSphericalPolynomial sets the SphericalPolynomial property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#sphericalpolynomial
func (d *DDSInfo) SetSphericalPolynomial(sphericalPolynomial *SphericalPolynomial) *DDSInfo {
	d.p.Set("sphericalPolynomial", sphericalPolynomial.JSObject())
	return d
}

// TextureType returns the TextureType property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#texturetype
func (d *DDSInfo) TextureType() float64 {
	retVal := d.p.Get("textureType")
	return retVal.Float()
}

// SetTextureType sets the TextureType property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#texturetype
func (d *DDSInfo) SetTextureType(textureType float64) *DDSInfo {
	d.p.Set("textureType", textureType)
	return d
}

// Width returns the Width property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#width
func (d *DDSInfo) Width() float64 {
	retVal := d.p.Get("width")
	return retVal.Float()
}

// SetWidth sets the Width property of class DDSInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.ddsinfo#width
func (d *DDSInfo) SetWidth(width float64) *DDSInfo {
	d.p.Set("width", width)
	return d
}
