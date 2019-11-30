// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphericLight represents a babylon.js HemisphericLight.
// The HemisphericLight simulates the ambient environment light,
// so the passed direction is the light reflection direction, not the incoming direction.
type HemisphericLight struct {
	*Light
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HemisphericLight) JSObject() js.Value { return h.p }

// HemisphericLight returns a HemisphericLight JavaScript class.
func (ba *Babylon) HemisphericLight() *HemisphericLight {
	p := ba.ctx.Get("HemisphericLight")
	return HemisphericLightFromJSObject(p, ba.ctx)
}

// HemisphericLightFromJSObject returns a wrapped HemisphericLight JavaScript class.
func HemisphericLightFromJSObject(p js.Value, ctx js.Value) *HemisphericLight {
	return &HemisphericLight{Light: LightFromJSObject(p, ctx), ctx: ctx}
}

// NewHemisphericLight returns a new HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight
func (ba *Babylon) NewHemisphericLight(name string, direction *Vector3, scene *Scene) *HemisphericLight {
	p := ba.ctx.Get("HemisphericLight").New(name, direction.JSObject(), scene.JSObject())
	return HemisphericLightFromJSObject(p, ba.ctx)
}

// TODO: methods
