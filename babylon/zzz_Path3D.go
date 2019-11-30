// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Path3D represents a babylon.js Path3D.
// Represents a 3D path made up of multiple 3D points
type Path3D struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *Path3D) JSObject() js.Value { return p.p }

// Path3D returns a Path3D JavaScript class.
func (ba *Babylon) Path3D() *Path3D {
	p := ba.ctx.Get("Path3D")
	return Path3DFromJSObject(p, ba.ctx)
}

// Path3DFromJSObject returns a wrapped Path3D JavaScript class.
func Path3DFromJSObject(p js.Value, ctx js.Value) *Path3D {
	return &Path3D{p: p, ctx: ctx}
}

// NewPath3DOpts contains optional parameters for NewPath3D.
type NewPath3DOpts struct {
	FirstNormal *Vector3

	Raw *JSBool

	AlignTangentsWithPath *JSBool
}

// NewPath3D returns a new Path3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.path3d
func (ba *Babylon) NewPath3D(path *Vector3, opts *NewPath3DOpts) *Path3D {
	if opts == nil {
		opts = &NewPath3DOpts{}
	}

	p := ba.ctx.Get("Path3D").New(path.JSObject(), opts.FirstNormal.JSObject(), opts.Raw.JSObject(), opts.AlignTangentsWithPath.JSObject())
	return Path3DFromJSObject(p, ba.ctx)
}

// TODO: methods
