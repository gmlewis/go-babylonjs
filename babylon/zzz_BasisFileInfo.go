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

/*

// HasAlpha returns the HasAlpha property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#hasalpha
func (b *BasisFileInfo) HasAlpha(hasAlpha bool) *BasisFileInfo {
	p := ba.ctx.Get("BasisFileInfo").New(hasAlpha)
	return BasisFileInfoFromJSObject(p, ba.ctx)
}

// SetHasAlpha sets the HasAlpha property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#hasalpha
func (b *BasisFileInfo) SetHasAlpha(hasAlpha bool) *BasisFileInfo {
	p := ba.ctx.Get("BasisFileInfo").New(hasAlpha)
	return BasisFileInfoFromJSObject(p, ba.ctx)
}

// Images returns the Images property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#images
func (b *BasisFileInfo) Images(images []object) *BasisFileInfo {
	p := ba.ctx.Get("BasisFileInfo").New(images.JSObject())
	return BasisFileInfoFromJSObject(p, ba.ctx)
}

// SetImages sets the Images property of class BasisFileInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.basisfileinfo#images
func (b *BasisFileInfo) SetImages(images []object) *BasisFileInfo {
	p := ba.ctx.Get("BasisFileInfo").New(images.JSObject())
	return BasisFileInfoFromJSObject(p, ba.ctx)
}

*/
