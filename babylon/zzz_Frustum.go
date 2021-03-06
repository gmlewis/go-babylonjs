// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Frustum represents a babylon.js Frustum.
// Represents a camera frustum
type Frustum struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *Frustum) JSObject() js.Value { return f.p }

// Frustum returns a Frustum JavaScript class.
func (ba *Babylon) Frustum() *Frustum {
	p := ba.ctx.Get("Frustum")
	return FrustumFromJSObject(p, ba.ctx)
}

// FrustumFromJSObject returns a wrapped Frustum JavaScript class.
func FrustumFromJSObject(p js.Value, ctx js.Value) *Frustum {
	return &Frustum{p: p, ctx: ctx}
}

// FrustumArrayToJSArray returns a JavaScript Array for the wrapped array.
func FrustumArrayToJSArray(array []*Frustum) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// GetBottomPlaneToRef calls the GetBottomPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getbottomplanetoref
func (f *Frustum) GetBottomPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetBottomPlaneToRef", args...)
}

// GetFarPlaneToRef calls the GetFarPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getfarplanetoref
func (f *Frustum) GetFarPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetFarPlaneToRef", args...)
}

// GetLeftPlaneToRef calls the GetLeftPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getleftplanetoref
func (f *Frustum) GetLeftPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetLeftPlaneToRef", args...)
}

// GetNearPlaneToRef calls the GetNearPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getnearplanetoref
func (f *Frustum) GetNearPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetNearPlaneToRef", args...)
}

// GetPlanes calls the GetPlanes method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getplanes
func (f *Frustum) GetPlanes(transform *Matrix) []*Plane {

	args := make([]interface{}, 0, 1+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	retVal := f.p.Call("GetPlanes", args...)
	result := []*Plane{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, PlaneFromJSObject(retVal.Index(ri), f.ctx))
	}
	return result
}

// GetPlanesToRef calls the GetPlanesToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getplanestoref
func (f *Frustum) GetPlanesToRef(transform *Matrix, frustumPlanes []*Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	f.p.Call("GetPlanesToRef", args...)
}

// GetRightPlaneToRef calls the GetRightPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#getrightplanetoref
func (f *Frustum) GetRightPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetRightPlaneToRef", args...)
}

// GetTopPlaneToRef calls the GetTopPlaneToRef method on the Frustum object.
//
// https://doc.babylonjs.com/api/classes/babylon.frustum#gettopplanetoref
func (f *Frustum) GetTopPlaneToRef(transform *Matrix, frustumPlane *Plane) {

	args := make([]interface{}, 0, 2+0)

	if transform == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transform.JSObject())
	}

	if frustumPlane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, frustumPlane.JSObject())
	}

	f.p.Call("GetTopPlaneToRef", args...)
}
