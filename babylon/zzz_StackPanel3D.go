// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StackPanel3D represents a babylon.js StackPanel3D.
// Class used to create a stack panel in 3D on XY plane
type StackPanel3D struct {
	*Container3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StackPanel3D) JSObject() js.Value { return s.p }

// StackPanel3D returns a StackPanel3D JavaScript class.
func (ba *Babylon) StackPanel3D() *StackPanel3D {
	p := ba.ctx.Get("StackPanel3D")
	return StackPanel3DFromJSObject(p, ba.ctx)
}

// StackPanel3DFromJSObject returns a wrapped StackPanel3D JavaScript class.
func StackPanel3DFromJSObject(p js.Value, ctx js.Value) *StackPanel3D {
	return &StackPanel3D{Container3D: Container3DFromJSObject(p, ctx), ctx: ctx}
}

// NewStackPanel3DOpts contains optional parameters for NewStackPanel3D.
type NewStackPanel3DOpts struct {
	IsVertical *JSBool
}

// NewStackPanel3D returns a new StackPanel3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel3d
func (ba *Babylon) NewStackPanel3D(opts *NewStackPanel3DOpts) *StackPanel3D {
	if opts == nil {
		opts = &NewStackPanel3DOpts{}
	}

	p := ba.ctx.Get("StackPanel3D").New(opts.IsVertical.JSObject())
	return StackPanel3DFromJSObject(p, ba.ctx)
}

// TODO: methods
