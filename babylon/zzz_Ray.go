// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Ray represents a babylon.js Ray.
// Class representing a ray with position and direction
type Ray struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *Ray) JSObject() js.Value { return r.p }

// Ray returns a Ray JavaScript class.
func (ba *Babylon) Ray() *Ray {
	p := ba.ctx.Get("Ray")
	return RayFromJSObject(p, ba.ctx)
}

// RayFromJSObject returns a wrapped Ray JavaScript class.
func RayFromJSObject(p js.Value, ctx js.Value) *Ray {
	return &Ray{p: p, ctx: ctx}
}

// RayArrayToJSArray returns a JavaScript Array for the wrapped array.
func RayArrayToJSArray(array []*Ray) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRayOpts contains optional parameters for NewRay.
type NewRayOpts struct {
	Length *float64
}

// NewRay returns a new Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#constructor
func (ba *Babylon) NewRay(origin *Vector3, direction *Vector3, opts *NewRayOpts) *Ray {
	if opts == nil {
		opts = &NewRayOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, origin.JSObject())
	args = append(args, direction.JSObject())

	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	p := ba.ctx.Get("Ray").New(args...)
	return RayFromJSObject(p, ba.ctx)
}

// CreateNew calls the CreateNew method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#createnew
func (r *Ray) CreateNew(x float64, y float64, viewportWidth float64, viewportHeight float64, world *Matrix, view *Matrix, projection *Matrix) *Ray {

	args := make([]interface{}, 0, 7+0)

	args = append(args, x)

	args = append(args, y)

	args = append(args, viewportWidth)

	args = append(args, viewportHeight)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if view == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, view.JSObject())
	}

	if projection == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, projection.JSObject())
	}

	retVal := r.p.Call("CreateNew", args...)
	return RayFromJSObject(retVal, r.ctx)
}

// RayCreateNewFromToOpts contains optional parameters for Ray.CreateNewFromTo.
type RayCreateNewFromToOpts struct {
	World *Matrix
}

// CreateNewFromTo calls the CreateNewFromTo method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#createnewfromto
func (r *Ray) CreateNewFromTo(origin *Vector3, end *Vector3, opts *RayCreateNewFromToOpts) *Ray {
	if opts == nil {
		opts = &RayCreateNewFromToOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if origin == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, origin.JSObject())
	}

	if end == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, end.JSObject())
	}

	if opts.World == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.World.JSObject())
	}

	retVal := r.p.Call("CreateNewFromTo", args...)
	return RayFromJSObject(retVal, r.ctx)
}

// IntersectionSegment calls the IntersectionSegment method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectionsegment
func (r *Ray) IntersectionSegment(sega *Vector3, segb *Vector3, threshold float64) float64 {

	args := make([]interface{}, 0, 3+0)

	if sega == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, sega.JSObject())
	}

	if segb == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, segb.JSObject())
	}

	args = append(args, threshold)

	retVal := r.p.Call("intersectionSegment", args...)
	return retVal.Float()
}

// RayIntersectsAxisOpts contains optional parameters for Ray.IntersectsAxis.
type RayIntersectsAxisOpts struct {
	Offset *float64
}

// IntersectsAxis calls the IntersectsAxis method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsaxis
func (r *Ray) IntersectsAxis(axis string, opts *RayIntersectsAxisOpts) *Vector3 {
	if opts == nil {
		opts = &RayIntersectsAxisOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, axis)

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	retVal := r.p.Call("intersectsAxis", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// RayIntersectsBoxOpts contains optional parameters for Ray.IntersectsBox.
type RayIntersectsBoxOpts struct {
	IntersectionTreshold *float64
}

// IntersectsBox calls the IntersectsBox method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsbox
func (r *Ray) IntersectsBox(box *BoundingBox, opts *RayIntersectsBoxOpts) bool {
	if opts == nil {
		opts = &RayIntersectsBoxOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if box == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, box.JSObject())
	}

	if opts.IntersectionTreshold == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IntersectionTreshold)
	}

	retVal := r.p.Call("intersectsBox", args...)
	return retVal.Bool()
}

// RayIntersectsBoxMinMaxOpts contains optional parameters for Ray.IntersectsBoxMinMax.
type RayIntersectsBoxMinMaxOpts struct {
	IntersectionTreshold *float64
}

// IntersectsBoxMinMax calls the IntersectsBoxMinMax method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsboxminmax
func (r *Ray) IntersectsBoxMinMax(minimum *Vector3, maximum *Vector3, opts *RayIntersectsBoxMinMaxOpts) bool {
	if opts == nil {
		opts = &RayIntersectsBoxMinMaxOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if minimum == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, minimum.JSObject())
	}

	if maximum == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, maximum.JSObject())
	}

	if opts.IntersectionTreshold == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IntersectionTreshold)
	}

	retVal := r.p.Call("intersectsBoxMinMax", args...)
	return retVal.Bool()
}

// RayIntersectsMeshOpts contains optional parameters for Ray.IntersectsMesh.
type RayIntersectsMeshOpts struct {
	FastCheck *bool
}

// IntersectsMesh calls the IntersectsMesh method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsmesh
func (r *Ray) IntersectsMesh(mesh *AbstractMesh, opts *RayIntersectsMeshOpts) *PickingInfo {
	if opts == nil {
		opts = &RayIntersectsMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.FastCheck == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FastCheck)
	}

	retVal := r.p.Call("intersectsMesh", args...)
	return PickingInfoFromJSObject(retVal, r.ctx)
}

// RayIntersectsMeshesOpts contains optional parameters for Ray.IntersectsMeshes.
type RayIntersectsMeshesOpts struct {
	FastCheck *bool
	Results   []js.Value
}

// IntersectsMeshes calls the IntersectsMeshes method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsmeshes
func (r *Ray) IntersectsMeshes(meshes []*AbstractMesh, opts *RayIntersectsMeshesOpts) []*PickingInfo {
	if opts == nil {
		opts = &RayIntersectsMeshesOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, AbstractMeshArrayToJSArray(meshes))

	if opts.FastCheck == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FastCheck)
	}
	if opts.Results == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Results)
	}

	retVal := r.p.Call("intersectsMeshes", args...)
	result := []*PickingInfo{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, PickingInfoFromJSObject(retVal.Index(ri), r.ctx))
	}
	return result
}

// IntersectsPlane calls the IntersectsPlane method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectsplane
func (r *Ray) IntersectsPlane(plane *Plane) float64 {

	args := make([]interface{}, 0, 1+0)

	if plane == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, plane.JSObject())
	}

	retVal := r.p.Call("intersectsPlane", args...)
	return retVal.Float()
}

// RayIntersectsSphereOpts contains optional parameters for Ray.IntersectsSphere.
type RayIntersectsSphereOpts struct {
	IntersectionTreshold *float64
}

// IntersectsSphere calls the IntersectsSphere method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectssphere
func (r *Ray) IntersectsSphere(sphere *BoundingSphere, opts *RayIntersectsSphereOpts) bool {
	if opts == nil {
		opts = &RayIntersectsSphereOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if sphere == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, sphere.JSObject())
	}

	if opts.IntersectionTreshold == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IntersectionTreshold)
	}

	retVal := r.p.Call("intersectsSphere", args...)
	return retVal.Bool()
}

// IntersectsTriangle calls the IntersectsTriangle method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#intersectstriangle
func (r *Ray) IntersectsTriangle(vertex0 *Vector3, vertex1 *Vector3, vertex2 *Vector3) js.Value {

	args := make([]interface{}, 0, 3+0)

	if vertex0 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vertex0.JSObject())
	}

	if vertex1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vertex1.JSObject())
	}

	if vertex2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vertex2.JSObject())
	}

	retVal := r.p.Call("intersectsTriangle", args...)
	return retVal
}

// Transform calls the Transform method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#transform
func (r *Ray) Transform(ray *Ray, matrix *Matrix) *Ray {

	args := make([]interface{}, 0, 2+0)

	if ray == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ray.JSObject())
	}

	if matrix == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, matrix.JSObject())
	}

	retVal := r.p.Call("Transform", args...)
	return RayFromJSObject(retVal, r.ctx)
}

// TransformToRef calls the TransformToRef method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#transformtoref
func (r *Ray) TransformToRef(ray *Ray, matrix *Matrix, result *Ray) {

	args := make([]interface{}, 0, 3+0)

	if ray == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ray.JSObject())
	}

	if matrix == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, matrix.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	r.p.Call("TransformToRef", args...)
}

// UnprojectRayToRef calls the UnprojectRayToRef method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#unprojectraytoref
func (r *Ray) UnprojectRayToRef(sourceX float64, sourceY float64, viewportWidth float64, viewportHeight float64, world *Matrix, view *Matrix, projection *Matrix) {

	args := make([]interface{}, 0, 7+0)

	args = append(args, sourceX)

	args = append(args, sourceY)

	args = append(args, viewportWidth)

	args = append(args, viewportHeight)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if view == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, view.JSObject())
	}

	if projection == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, projection.JSObject())
	}

	r.p.Call("unprojectRayToRef", args...)
}

// Update calls the Update method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#update
func (r *Ray) Update(x float64, y float64, viewportWidth float64, viewportHeight float64, world *Matrix, view *Matrix, projection *Matrix) *Ray {

	args := make([]interface{}, 0, 7+0)

	args = append(args, x)

	args = append(args, y)

	args = append(args, viewportWidth)

	args = append(args, viewportHeight)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if view == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, view.JSObject())
	}

	if projection == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, projection.JSObject())
	}

	retVal := r.p.Call("update", args...)
	return RayFromJSObject(retVal, r.ctx)
}

// Zero calls the Zero method on the Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#zero
func (r *Ray) Zero() *Ray {

	retVal := r.p.Call("Zero")
	return RayFromJSObject(retVal, r.ctx)
}

// Direction returns the Direction property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#direction
func (r *Ray) Direction() *Vector3 {
	retVal := r.p.Get("direction")
	return Vector3FromJSObject(retVal, r.ctx)
}

// SetDirection sets the Direction property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#direction
func (r *Ray) SetDirection(direction *Vector3) *Ray {
	r.p.Set("direction", direction.JSObject())
	return r
}

// Length returns the Length property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#length
func (r *Ray) Length() float64 {
	retVal := r.p.Get("length")
	return retVal.Float()
}

// SetLength sets the Length property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#length
func (r *Ray) SetLength(length float64) *Ray {
	r.p.Set("length", length)
	return r
}

// Origin returns the Origin property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#origin
func (r *Ray) Origin() *Vector3 {
	retVal := r.p.Get("origin")
	return Vector3FromJSObject(retVal, r.ctx)
}

// SetOrigin sets the Origin property of class Ray.
//
// https://doc.babylonjs.com/api/classes/babylon.ray#origin
func (r *Ray) SetOrigin(origin *Vector3) *Ray {
	r.p.Set("origin", origin.JSObject())
	return r
}
