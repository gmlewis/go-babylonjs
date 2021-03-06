// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BasisTranscodeConfiguration represents a babylon.js BasisTranscodeConfiguration.
// Configuration options for the Basis transcoder
type BasisTranscodeConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BasisTranscodeConfiguration) JSObject() js.Value { return b.p }

// BasisTranscodeConfiguration returns a BasisTranscodeConfiguration JavaScript class.
func (ba *Babylon) BasisTranscodeConfiguration() *BasisTranscodeConfiguration {
	p := ba.ctx.Get("BasisTranscodeConfiguration")
	return BasisTranscodeConfigurationFromJSObject(p, ba.ctx)
}

// BasisTranscodeConfigurationFromJSObject returns a wrapped BasisTranscodeConfiguration JavaScript class.
func BasisTranscodeConfigurationFromJSObject(p js.Value, ctx js.Value) *BasisTranscodeConfiguration {
	return &BasisTranscodeConfiguration{p: p, ctx: ctx}
}

// BasisTranscodeConfigurationArrayToJSArray returns a JavaScript Array for the wrapped array.
func BasisTranscodeConfigurationArrayToJSArray(array []*BasisTranscodeConfiguration) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// LoadMipmapLevels returns the LoadMipmapLevels property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#loadmipmaplevels
func (b *BasisTranscodeConfiguration) LoadMipmapLevels() bool {
	retVal := b.p.Get("loadMipmapLevels")
	return retVal.Bool()
}

// SetLoadMipmapLevels sets the LoadMipmapLevels property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#loadmipmaplevels
func (b *BasisTranscodeConfiguration) SetLoadMipmapLevels(loadMipmapLevels bool) *BasisTranscodeConfiguration {
	b.p.Set("loadMipmapLevels", loadMipmapLevels)
	return b
}

// LoadSingleImage returns the LoadSingleImage property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#loadsingleimage
func (b *BasisTranscodeConfiguration) LoadSingleImage() float64 {
	retVal := b.p.Get("loadSingleImage")
	return retVal.Float()
}

// SetLoadSingleImage sets the LoadSingleImage property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#loadsingleimage
func (b *BasisTranscodeConfiguration) SetLoadSingleImage(loadSingleImage float64) *BasisTranscodeConfiguration {
	b.p.Set("loadSingleImage", loadSingleImage)
	return b
}

// SupportedCompressionFormats returns the SupportedCompressionFormats property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#supportedcompressionformats
func (b *BasisTranscodeConfiguration) SupportedCompressionFormats() js.Value {
	retVal := b.p.Get("supportedCompressionFormats")
	return retVal
}

// SetSupportedCompressionFormats sets the SupportedCompressionFormats property of class BasisTranscodeConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.basistranscodeconfiguration#supportedcompressionformats
func (b *BasisTranscodeConfiguration) SetSupportedCompressionFormats(supportedCompressionFormats js.Value) *BasisTranscodeConfiguration {
	b.p.Set("supportedCompressionFormats", supportedCompressionFormats)
	return b
}
