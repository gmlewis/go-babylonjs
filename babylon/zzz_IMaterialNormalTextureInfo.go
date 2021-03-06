// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMaterialNormalTextureInfo represents a babylon.js IMaterialNormalTextureInfo.
// Loader interface with additional members.
type IMaterialNormalTextureInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMaterialNormalTextureInfo) JSObject() js.Value { return i.p }

// IMaterialNormalTextureInfo returns a IMaterialNormalTextureInfo JavaScript class.
func (ba *Babylon) IMaterialNormalTextureInfo() *IMaterialNormalTextureInfo {
	p := ba.ctx.Get("IMaterialNormalTextureInfo")
	return IMaterialNormalTextureInfoFromJSObject(p, ba.ctx)
}

// IMaterialNormalTextureInfoFromJSObject returns a wrapped IMaterialNormalTextureInfo JavaScript class.
func IMaterialNormalTextureInfoFromJSObject(p js.Value, ctx js.Value) *IMaterialNormalTextureInfo {
	return &IMaterialNormalTextureInfo{p: p, ctx: ctx}
}

// IMaterialNormalTextureInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMaterialNormalTextureInfoArrayToJSArray(array []*IMaterialNormalTextureInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
