// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMaterialOcclusionTextureInfo represents a babylon.js IMaterialOcclusionTextureInfo.
// Loader interface with additional members.
type IMaterialOcclusionTextureInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMaterialOcclusionTextureInfo) JSObject() js.Value { return i.p }

// IMaterialOcclusionTextureInfo returns a IMaterialOcclusionTextureInfo JavaScript class.
func (ba *Babylon) IMaterialOcclusionTextureInfo() *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo")
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// IMaterialOcclusionTextureInfoFromJSObject returns a wrapped IMaterialOcclusionTextureInfo JavaScript class.
func IMaterialOcclusionTextureInfoFromJSObject(p js.Value, ctx js.Value) *IMaterialOcclusionTextureInfo {
	return &IMaterialOcclusionTextureInfo{p: p, ctx: ctx}
}

// IMaterialOcclusionTextureInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMaterialOcclusionTextureInfoArrayToJSArray(array []*IMaterialOcclusionTextureInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
