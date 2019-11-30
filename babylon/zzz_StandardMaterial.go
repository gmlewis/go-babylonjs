// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StandardMaterial represents a babylon.js StandardMaterial.
// This is the default material used in Babylon. It is the best trade off between quality
// and performances.
//
// See: http://doc.babylonjs.com/babylon101/materials
type StandardMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StandardMaterial) JSObject() js.Value { return s.p }

// StandardMaterial returns a StandardMaterial JavaScript class.
func (ba *Babylon) StandardMaterial() *StandardMaterial {
	p := ba.ctx.Get("StandardMaterial")
	return StandardMaterialFromJSObject(p, ba.ctx)
}

// StandardMaterialFromJSObject returns a wrapped StandardMaterial JavaScript class.
func StandardMaterialFromJSObject(p js.Value, ctx js.Value) *StandardMaterial {
	return &StandardMaterial{p: p, ctx: ctx}
}

// NewStandardMaterial returns a new StandardMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.standardmaterial
func (ba *Babylon) NewStandardMaterial(name string, scene *Scene) *StandardMaterial {
	p := ba.ctx.Get("StandardMaterial").New(name, scene.JSObject())
	return StandardMaterialFromJSObject(p, ba.ctx)
}

// TODO: methods
