// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingInfo represents a babylon.js BoundingInfo.
// Info for a bounding data of a mesh
type BoundingInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BoundingInfo) JSObject() js.Value { return b.p }

// BoundingInfo returns a BoundingInfo JavaScript class.
func (ba *Babylon) BoundingInfo() *BoundingInfo {
	p := ba.ctx.Get("BoundingInfo")
	return BoundingInfoFromJSObject(p, ba.ctx)
}

// BoundingInfoFromJSObject returns a wrapped BoundingInfo JavaScript class.
func BoundingInfoFromJSObject(p js.Value, ctx js.Value) *BoundingInfo {
	return &BoundingInfo{p: p, ctx: ctx}
}

// BoundingInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func BoundingInfoArrayToJSArray(array []*BoundingInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBoundingInfoOpts contains optional parameters for NewBoundingInfo.
type NewBoundingInfoOpts struct {
	WorldMatrix *Matrix
}

// NewBoundingInfo returns a new BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#constructor
func (ba *Babylon) NewBoundingInfo(minimum *Vector3, maximum *Vector3, opts *NewBoundingInfoOpts) *BoundingInfo {
	if opts == nil {
		opts = &NewBoundingInfoOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, minimum.JSObject())
	args = append(args, maximum.JSObject())

	if opts.WorldMatrix == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.WorldMatrix.JSObject())
	}

	p := ba.ctx.Get("BoundingInfo").New(args...)
	return BoundingInfoFromJSObject(p, ba.ctx)
}

// CenterOn calls the CenterOn method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#centeron
func (b *BoundingInfo) CenterOn(center *Vector3, extend *Vector3) *BoundingInfo {

	args := make([]interface{}, 0, 2+0)

	if center == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, center.JSObject())
	}

	if extend == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, extend.JSObject())
	}

	retVal := b.p.Call("centerOn", args...)
	return BoundingInfoFromJSObject(retVal, b.ctx)
}

// Intersects calls the Intersects method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#intersects
func (b *BoundingInfo) Intersects(boundingInfo *BoundingInfo, precise bool) bool {

	args := make([]interface{}, 0, 2+0)

	if boundingInfo == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, boundingInfo.JSObject())
	}

	args = append(args, precise)

	retVal := b.p.Call("intersects", args...)
	return retVal.Bool()
}

// IntersectsPoint calls the IntersectsPoint method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#intersectspoint
func (b *BoundingInfo) IntersectsPoint(point *Vector3) bool {

	args := make([]interface{}, 0, 1+0)

	if point == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, point.JSObject())
	}

	retVal := b.p.Call("intersectsPoint", args...)
	return retVal.Bool()
}

// IsCompletelyInFrustum calls the IsCompletelyInFrustum method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#iscompletelyinfrustum
func (b *BoundingInfo) IsCompletelyInFrustum(frustumPlanes []*Plane) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	retVal := b.p.Call("isCompletelyInFrustum", args...)
	return retVal.Bool()
}

// BoundingInfoIsInFrustumOpts contains optional parameters for BoundingInfo.IsInFrustum.
type BoundingInfoIsInFrustumOpts struct {
	Strategy *float64
}

// IsInFrustum calls the IsInFrustum method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#isinfrustum
func (b *BoundingInfo) IsInFrustum(frustumPlanes []*Plane, opts *BoundingInfoIsInFrustumOpts) bool {
	if opts == nil {
		opts = &BoundingInfoIsInFrustumOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	if opts.Strategy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strategy)
	}

	retVal := b.p.Call("isInFrustum", args...)
	return retVal.Bool()
}

// BoundingInfoReConstructOpts contains optional parameters for BoundingInfo.ReConstruct.
type BoundingInfoReConstructOpts struct {
	WorldMatrix *Matrix
}

// ReConstruct calls the ReConstruct method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#reconstruct
func (b *BoundingInfo) ReConstruct(min *Vector3, max *Vector3, opts *BoundingInfoReConstructOpts) {
	if opts == nil {
		opts = &BoundingInfoReConstructOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if min == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, min.JSObject())
	}

	if max == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, max.JSObject())
	}

	if opts.WorldMatrix == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.WorldMatrix.JSObject())
	}

	b.p.Call("reConstruct", args...)
}

// Scale calls the Scale method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#scale
func (b *BoundingInfo) Scale(factor float64) *BoundingInfo {

	args := make([]interface{}, 0, 1+0)

	args = append(args, factor)

	retVal := b.p.Call("scale", args...)
	return BoundingInfoFromJSObject(retVal, b.ctx)
}

// Update calls the Update method on the BoundingInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#update
func (b *BoundingInfo) Update(world *Matrix) {

	args := make([]interface{}, 0, 1+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	b.p.Call("update", args...)
}

// BoundingBox returns the BoundingBox property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#boundingbox
func (b *BoundingInfo) BoundingBox() *BoundingBox {
	retVal := b.p.Get("boundingBox")
	return BoundingBoxFromJSObject(retVal, b.ctx)
}

// SetBoundingBox sets the BoundingBox property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#boundingbox
func (b *BoundingInfo) SetBoundingBox(boundingBox *BoundingBox) *BoundingInfo {
	b.p.Set("boundingBox", boundingBox.JSObject())
	return b
}

// BoundingSphere returns the BoundingSphere property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#boundingsphere
func (b *BoundingInfo) BoundingSphere() *BoundingSphere {
	retVal := b.p.Get("boundingSphere")
	return BoundingSphereFromJSObject(retVal, b.ctx)
}

// SetBoundingSphere sets the BoundingSphere property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#boundingsphere
func (b *BoundingInfo) SetBoundingSphere(boundingSphere *BoundingSphere) *BoundingInfo {
	b.p.Set("boundingSphere", boundingSphere.JSObject())
	return b
}

// DiagonalLength returns the DiagonalLength property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#diagonallength
func (b *BoundingInfo) DiagonalLength() float64 {
	retVal := b.p.Get("diagonalLength")
	return retVal.Float()
}

// SetDiagonalLength sets the DiagonalLength property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#diagonallength
func (b *BoundingInfo) SetDiagonalLength(diagonalLength float64) *BoundingInfo {
	b.p.Set("diagonalLength", diagonalLength)
	return b
}

// IsLocked returns the IsLocked property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#islocked
func (b *BoundingInfo) IsLocked() bool {
	retVal := b.p.Get("isLocked")
	return retVal.Bool()
}

// SetIsLocked sets the IsLocked property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#islocked
func (b *BoundingInfo) SetIsLocked(isLocked bool) *BoundingInfo {
	b.p.Set("isLocked", isLocked)
	return b
}

// Maximum returns the Maximum property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#maximum
func (b *BoundingInfo) Maximum() *Vector3 {
	retVal := b.p.Get("maximum")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetMaximum sets the Maximum property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#maximum
func (b *BoundingInfo) SetMaximum(maximum *Vector3) *BoundingInfo {
	b.p.Set("maximum", maximum.JSObject())
	return b
}

// Minimum returns the Minimum property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#minimum
func (b *BoundingInfo) Minimum() *Vector3 {
	retVal := b.p.Get("minimum")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetMinimum sets the Minimum property of class BoundingInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.boundinginfo#minimum
func (b *BoundingInfo) SetMinimum(minimum *Vector3) *BoundingInfo {
	b.p.Set("minimum", minimum.JSObject())
	return b
}
