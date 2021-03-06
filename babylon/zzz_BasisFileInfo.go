// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BasisFileInfo represents a babylon.js BasisFileInfo.
// Info about the .basis files
type BasisFileInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BasisFileInfo) JSObject() js.Value { return b.p }

// BasisFileInfo returns a BasisFileInfo JavaScript class.
func (ba *Babylon) BasisFileInfo() *BasisFileInfo {
	p := ba.ctx.Get("BasisFileInfo")
	return BasisFileInfoFromJSObject(p, ba.ctx)
}

// BasisFileInfoFromJSObject returns a wrapped BasisFileInfo JavaScript class.
func BasisFileInfoFromJSObject(p js.Value, ctx js.Value) *BasisFileInfo {
	return &BasisFileInfo{p: p, ctx: ctx}
}

// BasisFileInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func BasisFileInfoArrayToJSArray(array []*BasisFileInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// HasAlpha returns the HasAlpha property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#hasalpha
func (b *BasisFileInfo) HasAlpha() bool {
	retVal := b.p.Get("hasAlpha")
	return retVal.Bool()
}

// SetHasAlpha sets the HasAlpha property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#hasalpha
func (b *BasisFileInfo) SetHasAlpha(hasAlpha bool) *BasisFileInfo {
	b.p.Set("hasAlpha", hasAlpha)
	return b
}

// Images returns the Images property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#images
func (b *BasisFileInfo) Images() js.Value {
	retVal := b.p.Get("images")
	return retVal
}

// SetImages sets the Images property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#images
func (b *BasisFileInfo) SetImages(images []js.Value) *BasisFileInfo {
	b.p.Set("images", images)
	return b
}
