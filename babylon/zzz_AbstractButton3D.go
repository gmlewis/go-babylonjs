// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AbstractButton3D represents a babylon.js AbstractButton3D.
// Class used as a root to all buttons
type AbstractButton3D struct {
	*Control3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AbstractButton3D) JSObject() js.Value { return a.p }

// AbstractButton3D returns a AbstractButton3D JavaScript class.
func (ba *Babylon) AbstractButton3D() *AbstractButton3D {
	p := ba.ctx.Get("AbstractButton3D")
	return AbstractButton3DFromJSObject(p, ba.ctx)
}

// AbstractButton3DFromJSObject returns a wrapped AbstractButton3D JavaScript class.
func AbstractButton3DFromJSObject(p js.Value, ctx js.Value) *AbstractButton3D {
	return &AbstractButton3D{Control3D: Control3DFromJSObject(p, ctx), ctx: ctx}
}

// AbstractButton3DArrayToJSArray returns a JavaScript Array for the wrapped array.
func AbstractButton3DArrayToJSArray(array []*AbstractButton3D) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAbstractButton3DOpts contains optional parameters for NewAbstractButton3D.
type NewAbstractButton3DOpts struct {
	Name *string
}

// NewAbstractButton3D returns a new AbstractButton3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractbutton3d
func (gui *GUI) NewAbstractButton3D(opts *NewAbstractButton3DOpts) *AbstractButton3D {
	if opts == nil {
		opts = &NewAbstractButton3DOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("AbstractButton3D").New(args...)
	return AbstractButton3DFromJSObject(p, gui.ctx)
}
