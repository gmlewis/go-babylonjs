// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerInfoPre represents a babylon.js PointerInfoPre.
// This class is used to store pointer related info for the onPrePointerObservable event.
// Set the skipOnPointerObservable property to true if you want the engine to stop any process after this event is triggered, even not calling onPointerObservable
type PointerInfoPre struct{ *PointerInfoBase }

// JSObject returns the underlying js.Value.
func (p *PointerInfoPre) JSObject() js.Value { return p.p }

// PointerInfoPre returns a PointerInfoPre JavaScript class.
func (b *Babylon) PointerInfoPre() *PointerInfoPre {
	p := b.ctx.Get("PointerInfoPre")
	return PointerInfoPreFromJSObject(p)
}

// PointerInfoPreFromJSObject returns a wrapped PointerInfoPre JavaScript class.
func PointerInfoPreFromJSObject(p js.Value) *PointerInfoPre {
	return &PointerInfoPre{PointerInfoBaseFromJSObject(p)}
}

// NewPointerInfoPre returns a new PointerInfoPre object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre
func (b *Babylon) NewPointerInfoPre(jsType float64, event PointerEvent, localX float64, localY float64) *PointerInfoPre {
	p := b.ctx.Get("PointerInfoPre").New(jsType, event.JSObject(), localX, localY)
	return PointerInfoPreFromJSObject(p)
}

// TODO: methods
