// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SubMesh represents a babylon.js SubMesh.
// Defines a subdivision inside a mesh
type SubMesh struct {
	*BaseSubMesh
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SubMesh) JSObject() js.Value { return s.p }

// SubMesh returns a SubMesh JavaScript class.
func (ba *Babylon) SubMesh() *SubMesh {
	p := ba.ctx.Get("SubMesh")
	return SubMeshFromJSObject(p, ba.ctx)
}

// SubMeshFromJSObject returns a wrapped SubMesh JavaScript class.
func SubMeshFromJSObject(p js.Value, ctx js.Value) *SubMesh {
	return &SubMesh{BaseSubMesh: BaseSubMeshFromJSObject(p, ctx), ctx: ctx}
}

// SubMeshArrayToJSArray returns a JavaScript Array for the wrapped array.
func SubMeshArrayToJSArray(array []*SubMesh) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSubMeshOpts contains optional parameters for NewSubMesh.
type NewSubMeshOpts struct {
	RenderingMesh     *Mesh
	CreateBoundingBox *bool
}

// NewSubMesh returns a new SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh
func (ba *Babylon) NewSubMesh(materialIndex float64, verticesStart float64, verticesCount float64, indexStart float64, indexCount float64, mesh *AbstractMesh, opts *NewSubMeshOpts) *SubMesh {
	if opts == nil {
		opts = &NewSubMeshOpts{}
	}

	args := make([]interface{}, 0, 6+2)

	args = append(args, materialIndex)
	args = append(args, verticesStart)
	args = append(args, verticesCount)
	args = append(args, indexStart)
	args = append(args, indexCount)
	args = append(args, mesh.JSObject())

	if opts.RenderingMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.RenderingMesh.JSObject())
	}
	if opts.CreateBoundingBox == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CreateBoundingBox)
	}

	p := ba.ctx.Get("SubMesh").New(args...)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SubMeshAddToMeshOpts contains optional parameters for SubMesh.AddToMesh.
type SubMeshAddToMeshOpts struct {
	RenderingMesh     *Mesh
	CreateBoundingBox *bool
}

// AddToMesh calls the AddToMesh method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#addtomesh
func (s *SubMesh) AddToMesh(materialIndex float64, verticesStart float64, verticesCount float64, indexStart float64, indexCount float64, mesh *AbstractMesh, opts *SubMeshAddToMeshOpts) *SubMesh {
	if opts == nil {
		opts = &SubMeshAddToMeshOpts{}
	}

	args := make([]interface{}, 0, 6+2)

	args = append(args, materialIndex)
	args = append(args, verticesStart)
	args = append(args, verticesCount)
	args = append(args, indexStart)
	args = append(args, indexCount)
	args = append(args, mesh.JSObject())

	if opts.RenderingMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.RenderingMesh.JSObject())
	}
	if opts.CreateBoundingBox == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CreateBoundingBox)
	}

	retVal := s.p.Call("AddToMesh", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// CanIntersects calls the CanIntersects method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#canintersects
func (s *SubMesh) CanIntersects(ray *Ray) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ray.JSObject())

	retVal := s.p.Call("canIntersects", args...)
	return retVal.Bool()
}

// SubMeshCloneOpts contains optional parameters for SubMesh.Clone.
type SubMeshCloneOpts struct {
	NewRenderingMesh *Mesh
}

// Clone calls the Clone method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#clone
func (s *SubMesh) Clone(newMesh *AbstractMesh, opts *SubMeshCloneOpts) *SubMesh {
	if opts == nil {
		opts = &SubMeshCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, newMesh.JSObject())

	if opts.NewRenderingMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.NewRenderingMesh.JSObject())
	}

	retVal := s.p.Call("clone", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// SubMeshCreateFromIndicesOpts contains optional parameters for SubMesh.CreateFromIndices.
type SubMeshCreateFromIndicesOpts struct {
	RenderingMesh *Mesh
}

// CreateFromIndices calls the CreateFromIndices method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#createfromindices
func (s *SubMesh) CreateFromIndices(materialIndex float64, startIndex float64, indexCount float64, mesh *AbstractMesh, opts *SubMeshCreateFromIndicesOpts) *SubMesh {
	if opts == nil {
		opts = &SubMeshCreateFromIndicesOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, materialIndex)
	args = append(args, startIndex)
	args = append(args, indexCount)
	args = append(args, mesh.JSObject())

	if opts.RenderingMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.RenderingMesh.JSObject())
	}

	retVal := s.p.Call("CreateFromIndices", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#dispose
func (s *SubMesh) Dispose() {

	s.p.Call("dispose")
}

// GetBoundingInfo calls the GetBoundingInfo method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#getboundinginfo
func (s *SubMesh) GetBoundingInfo() *BoundingInfo {

	retVal := s.p.Call("getBoundingInfo")
	return BoundingInfoFromJSObject(retVal, s.ctx)
}

// GetClassName calls the GetClassName method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#getclassname
func (s *SubMesh) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// GetMaterial calls the GetMaterial method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#getmaterial
func (s *SubMesh) GetMaterial() *Material {

	retVal := s.p.Call("getMaterial")
	return MaterialFromJSObject(retVal, s.ctx)
}

// GetMesh calls the GetMesh method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#getmesh
func (s *SubMesh) GetMesh() *AbstractMesh {

	retVal := s.p.Call("getMesh")
	return AbstractMeshFromJSObject(retVal, s.ctx)
}

// GetRenderingMesh calls the GetRenderingMesh method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#getrenderingmesh
func (s *SubMesh) GetRenderingMesh() *Mesh {

	retVal := s.p.Call("getRenderingMesh")
	return MeshFromJSObject(retVal, s.ctx)
}

// SubMeshIntersectsOpts contains optional parameters for SubMesh.Intersects.
type SubMeshIntersectsOpts struct {
	FastCheck         *bool
	TrianglePredicate js.Value
}

// Intersects calls the Intersects method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#intersects
func (s *SubMesh) Intersects(ray *Ray, positions *Vector3, indices js.Value, opts *SubMeshIntersectsOpts) js.Value {
	if opts == nil {
		opts = &SubMeshIntersectsOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, ray.JSObject())
	args = append(args, positions.JSObject())
	args = append(args, indices)

	if opts.FastCheck == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FastCheck)
	}
	args = append(args, opts.TrianglePredicate)

	retVal := s.p.Call("intersects", args...)
	return retVal
}

// IsCompletelyInFrustum calls the IsCompletelyInFrustum method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#iscompletelyinfrustum
func (s *SubMesh) IsCompletelyInFrustum(frustumPlanes *Plane) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, frustumPlanes.JSObject())

	retVal := s.p.Call("isCompletelyInFrustum", args...)
	return retVal.Bool()
}

// IsInFrustum calls the IsInFrustum method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#isinfrustum
func (s *SubMesh) IsInFrustum(frustumPlanes *Plane) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, frustumPlanes.JSObject())

	retVal := s.p.Call("isInFrustum", args...)
	return retVal.Bool()
}

// SubMeshRefreshBoundingInfoOpts contains optional parameters for SubMesh.RefreshBoundingInfo.
type SubMeshRefreshBoundingInfoOpts struct {
	Data js.Value
}

// RefreshBoundingInfo calls the RefreshBoundingInfo method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#refreshboundinginfo
func (s *SubMesh) RefreshBoundingInfo(opts *SubMeshRefreshBoundingInfoOpts) *SubMesh {
	if opts == nil {
		opts = &SubMeshRefreshBoundingInfoOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.Data)

	retVal := s.p.Call("refreshBoundingInfo", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// Render calls the Render method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#render
func (s *SubMesh) Render(enableAlphaMode bool) *SubMesh {

	args := make([]interface{}, 0, 1+0)

	args = append(args, enableAlphaMode)

	retVal := s.p.Call("render", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// SetBoundingInfo calls the SetBoundingInfo method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#setboundinginfo
func (s *SubMesh) SetBoundingInfo(boundingInfo *BoundingInfo) *SubMesh {

	args := make([]interface{}, 0, 1+0)

	args = append(args, boundingInfo.JSObject())

	retVal := s.p.Call("setBoundingInfo", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

// UpdateBoundingInfo calls the UpdateBoundingInfo method on the SubMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#updateboundinginfo
func (s *SubMesh) UpdateBoundingInfo(world *Matrix) *SubMesh {

	args := make([]interface{}, 0, 1+0)

	args = append(args, world.JSObject())

	retVal := s.p.Call("updateBoundingInfo", args...)
	return SubMeshFromJSObject(retVal, s.ctx)
}

/*

// IndexCount returns the IndexCount property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#indexcount
func (s *SubMesh) IndexCount(indexCount float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(indexCount)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetIndexCount sets the IndexCount property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#indexcount
func (s *SubMesh) SetIndexCount(indexCount float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(indexCount)
	return SubMeshFromJSObject(p, ba.ctx)
}

// IndexStart returns the IndexStart property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#indexstart
func (s *SubMesh) IndexStart(indexStart float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(indexStart)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetIndexStart sets the IndexStart property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#indexstart
func (s *SubMesh) SetIndexStart(indexStart float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(indexStart)
	return SubMeshFromJSObject(p, ba.ctx)
}

// IsGlobal returns the IsGlobal property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#isglobal
func (s *SubMesh) IsGlobal(IsGlobal bool) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(IsGlobal)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetIsGlobal sets the IsGlobal property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#isglobal
func (s *SubMesh) SetIsGlobal(IsGlobal bool) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(IsGlobal)
	return SubMeshFromJSObject(p, ba.ctx)
}

// MaterialIndex returns the MaterialIndex property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#materialindex
func (s *SubMesh) MaterialIndex(materialIndex float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(materialIndex)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetMaterialIndex sets the MaterialIndex property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#materialindex
func (s *SubMesh) SetMaterialIndex(materialIndex float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(materialIndex)
	return SubMeshFromJSObject(p, ba.ctx)
}

// VerticesCount returns the VerticesCount property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#verticescount
func (s *SubMesh) VerticesCount(verticesCount float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(verticesCount)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetVerticesCount sets the VerticesCount property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#verticescount
func (s *SubMesh) SetVerticesCount(verticesCount float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(verticesCount)
	return SubMeshFromJSObject(p, ba.ctx)
}

// VerticesStart returns the VerticesStart property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#verticesstart
func (s *SubMesh) VerticesStart(verticesStart float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(verticesStart)
	return SubMeshFromJSObject(p, ba.ctx)
}

// SetVerticesStart sets the VerticesStart property of class SubMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.submesh#verticesstart
func (s *SubMesh) SetVerticesStart(verticesStart float64) *SubMesh {
	p := ba.ctx.Get("SubMesh").New(verticesStart)
	return SubMeshFromJSObject(p, ba.ctx)
}

*/
