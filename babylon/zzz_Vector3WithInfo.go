// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Vector3WithInfo represents a babylon.js Vector3WithInfo.
// Class used to transport BABYLON.Vector3 information for pointer events
type Vector3WithInfo struct {
	*Vector3
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *Vector3WithInfo) JSObject() js.Value { return v.p }

// Vector3WithInfo returns a Vector3WithInfo JavaScript class.
func (ba *Babylon) Vector3WithInfo() *Vector3WithInfo {
	p := ba.ctx.Get("Vector3WithInfo")
	return Vector3WithInfoFromJSObject(p, ba.ctx)
}

// Vector3WithInfoFromJSObject returns a wrapped Vector3WithInfo JavaScript class.
func Vector3WithInfoFromJSObject(p js.Value, ctx js.Value) *Vector3WithInfo {
	return &Vector3WithInfo{Vector3: Vector3FromJSObject(p, ctx), ctx: ctx}
}

// NewVector3WithInfoOpts contains optional parameters for NewVector3WithInfo.
type NewVector3WithInfoOpts struct {
	ButtonIndex *JSFloat64
}

// NewVector3WithInfo returns a new Vector3WithInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector3withinfo
func (ba *Babylon) NewVector3WithInfo(source *Vector3, opts *NewVector3WithInfoOpts) *Vector3WithInfo {
	if opts == nil {
		opts = &NewVector3WithInfoOpts{}
	}

	p := ba.ctx.Get("Vector3WithInfo").New(source.JSObject(), opts.ButtonIndex.JSObject())
	return Vector3WithInfoFromJSObject(p, ba.ctx)
}

// TODO: methods
