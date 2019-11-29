// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphericLight represents a babylon.js HemisphericLight.
// The HemisphericLight simulates the ambient environment light,
// so the passed direction is the light reflection direction, not the incoming direction.
type HemisphericLight struct{ *Light }

// JSObject returns the underlying js.Value.
func (h *HemisphericLight) JSObject() js.Value { return h.p }

// HemisphericLight returns a HemisphericLight JavaScript class.
func (b *Babylon) HemisphericLight() *HemisphericLight {
	p := b.ctx.Get("HemisphericLight")
	return HemisphericLightFromJSObject(p)
}

// HemisphericLightFromJSObject returns a wrapped HemisphericLight JavaScript class.
func HemisphericLightFromJSObject(p js.Value) *HemisphericLight {
	return &HemisphericLight{LightFromJSObject(p)}
}

// NewHemisphericLight returns a new HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight
func (b *Babylon) NewHemisphericLight(todo parameters) *HemisphericLight {
	p := b.ctx.Get("HemisphericLight").New(todo)
	return HemisphericLightFromJSObject(p)
}

// TODO: methods
