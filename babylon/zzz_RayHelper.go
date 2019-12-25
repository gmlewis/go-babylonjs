// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RayHelper represents a babylon.js RayHelper.
// As raycast might be hard to debug, the RayHelper can help rendering the different rays
// in order to better appreciate the issue one might have.
//
// See: http://doc.babylonjs.com/babylon101/raycasts#debugging
type RayHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RayHelper) JSObject() js.Value { return r.p }

// RayHelper returns a RayHelper JavaScript class.
func (ba *Babylon) RayHelper() *RayHelper {
	p := ba.ctx.Get("RayHelper")
	return RayHelperFromJSObject(p, ba.ctx)
}

// RayHelperFromJSObject returns a wrapped RayHelper JavaScript class.
func RayHelperFromJSObject(p js.Value, ctx js.Value) *RayHelper {
	return &RayHelper{p: p, ctx: ctx}
}

// RayHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func RayHelperArrayToJSArray(array []*RayHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRayHelper returns a new RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#constructor
func (ba *Babylon) NewRayHelper(ray *Ray) *RayHelper {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ray.JSObject())

	p := ba.ctx.Get("RayHelper").New(args...)
	return RayHelperFromJSObject(p, ba.ctx)
}

// RayHelperAttachToMeshOpts contains optional parameters for RayHelper.AttachToMesh.
type RayHelperAttachToMeshOpts struct {
	MeshSpaceDirection *Vector3
	MeshSpaceOrigin    *Vector3
	Length             *float64
}

// AttachToMesh calls the AttachToMesh method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#attachtomesh
func (r *RayHelper) AttachToMesh(mesh *AbstractMesh, opts *RayHelperAttachToMeshOpts) {
	if opts == nil {
		opts = &RayHelperAttachToMeshOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.MeshSpaceDirection == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MeshSpaceDirection.JSObject())
	}
	if opts.MeshSpaceOrigin == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MeshSpaceOrigin.JSObject())
	}
	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	r.p.Call("attachToMesh", args...)
}

// CreateAndShow calls the CreateAndShow method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#createandshow
func (r *RayHelper) CreateAndShow(ray *Ray, scene *Scene, color *Color3) *RayHelper {

	args := make([]interface{}, 0, 3+0)

	if ray == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ray.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if color == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color.JSObject())
	}

	retVal := r.p.Call("CreateAndShow", args...)
	return RayHelperFromJSObject(retVal, r.ctx)
}

// DetachFromMesh calls the DetachFromMesh method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#detachfrommesh
func (r *RayHelper) DetachFromMesh() {

	r.p.Call("detachFromMesh")
}

// Dispose calls the Dispose method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#dispose
func (r *RayHelper) Dispose() {

	r.p.Call("dispose")
}

// Hide calls the Hide method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#hide
func (r *RayHelper) Hide() {

	r.p.Call("hide")
}

// RayHelperShowOpts contains optional parameters for RayHelper.Show.
type RayHelperShowOpts struct {
	Color *Color3
}

// Show calls the Show method on the RayHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#show
func (r *RayHelper) Show(scene *Scene, opts *RayHelperShowOpts) {
	if opts == nil {
		opts = &RayHelperShowOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if opts.Color == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Color.JSObject())
	}

	r.p.Call("show", args...)
}

// Ray returns the Ray property of class RayHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#ray
func (r *RayHelper) Ray() *Ray {
	retVal := r.p.Get("ray")
	return RayFromJSObject(retVal, r.ctx)
}

// SetRay sets the Ray property of class RayHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.rayhelper#ray
func (r *RayHelper) SetRay(ray *Ray) *RayHelper {
	r.p.Set("ray", ray.JSObject())
	return r
}
