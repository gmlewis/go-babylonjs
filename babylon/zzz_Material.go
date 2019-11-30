// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Material represents a babylon.js Material.
// Base class for the main features of a material in Babylon.js
type Material struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *Material) JSObject() js.Value { return m.p }

// Material returns a Material JavaScript class.
func (ba *Babylon) Material() *Material {
	p := ba.ctx.Get("Material")
	return MaterialFromJSObject(p, ba.ctx)
}

// MaterialFromJSObject returns a wrapped Material JavaScript class.
func MaterialFromJSObject(p js.Value, ctx js.Value) *Material {
	return &Material{p: p, ctx: ctx}
}

// NewMaterialOpts contains optional parameters for NewMaterial.
type NewMaterialOpts struct {
	DoNotAdd *JSBool
}

// NewMaterial returns a new Material object.
//
// https://doc.babylonjs.com/api/classes/babylon.material
func (ba *Babylon) NewMaterial(name string, scene *Scene, opts *NewMaterialOpts) *Material {
	if opts == nil {
		opts = &NewMaterialOpts{}
	}

	p := ba.ctx.Get("Material").New(name, scene.JSObject(), opts.DoNotAdd.JSObject())
	return MaterialFromJSObject(p, ba.ctx)
}

// TODO: methods
