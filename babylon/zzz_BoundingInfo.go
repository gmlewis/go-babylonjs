// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingInfo represents a babylon.js BoundingInfo.
// Info for a bounding data of a mesh
type BoundingInfo struct{}

// JSObject returns the underlying js.Value.
func (b *BoundingInfo) JSObject() js.Value { return b.p }

// BoundingInfo returns a BoundingInfo JavaScript class.
func (b *Babylon) BoundingInfo() *BoundingInfo {
	p := b.ctx.Get("BoundingInfo")
	return BoundingInfoFromJSObject(p)
}

// BoundingInfoFromJSObject returns a wrapped BoundingInfo JavaScript class.
func BoundingInfoFromJSObject(p js.Value) *BoundingInfo {
	return &BoundingInfo{p: p}
}

// NewBoundingInfo returns a new BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo
func (b *Babylon) NewBoundingInfo(todo parameters) *BoundingInfo {
	p := b.ctx.Get("BoundingInfo").New(todo)
	return BoundingInfoFromJSObject(p)
}

// TODO: methods
