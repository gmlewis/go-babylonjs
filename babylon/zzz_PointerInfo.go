// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerInfo represents a babylon.js PointerInfo.
// This type contains all the data related to a pointer event in Babylon.js.
// The event member is an instance of PointerEvent for all types except PointerWheel and is of type MouseWheelEvent when type equals PointerWheel. The different event types can be found in the PointerEventTypes class.
type PointerInfo struct{ *PointerInfoBase }

// JSObject returns the underlying js.Value.
func (p *PointerInfo) JSObject() js.Value { return p.p }

// PointerInfo returns a PointerInfo JavaScript class.
func (b *Babylon) PointerInfo() *PointerInfo {
	p := b.ctx.Get("PointerInfo")
	return PointerInfoFromJSObject(p)
}

// PointerInfoFromJSObject returns a wrapped PointerInfo JavaScript class.
func PointerInfoFromJSObject(p js.Value) *PointerInfo {
	return &PointerInfo{PointerInfoBaseFromJSObject(p)}
}

// NewPointerInfo returns a new PointerInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfo
func (b *Babylon) NewPointerInfo(jsType float64, event PointerEvent, pickInfo *PickingInfo) *PointerInfo {
	p := b.ctx.Get("PointerInfo").New(jsType, event.JSObject(), pickInfo.JSObject())
	return PointerInfoFromJSObject(p)
}

// TODO: methods
