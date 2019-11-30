// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiMaterial represents a babylon.js MultiMaterial.
// A multi-material is used to apply different materials to different parts of the same object without the need of
// separate meshes. This can be use to improve performances.
//
// See: http://doc.babylonjs.com/how_to/multi_materials
type MultiMaterial struct {
	*Material
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiMaterial) JSObject() js.Value { return m.p }

// MultiMaterial returns a MultiMaterial JavaScript class.
func (ba *Babylon) MultiMaterial() *MultiMaterial {
	p := ba.ctx.Get("MultiMaterial")
	return MultiMaterialFromJSObject(p, ba.ctx)
}

// MultiMaterialFromJSObject returns a wrapped MultiMaterial JavaScript class.
func MultiMaterialFromJSObject(p js.Value, ctx js.Value) *MultiMaterial {
	return &MultiMaterial{Material: MaterialFromJSObject(p, ctx), ctx: ctx}
}

// NewMultiMaterial returns a new MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial
func (ba *Babylon) NewMultiMaterial(name string, scene *Scene) *MultiMaterial {
	p := ba.ctx.Get("MultiMaterial").New(name, scene.JSObject())
	return MultiMaterialFromJSObject(p, ba.ctx)
}

// TODO: methods
