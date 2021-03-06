// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SolidParticleSystem represents a babylon.js SolidParticleSystem.
// The SPS is a single updatable mesh. The solid particles are simply separate parts or faces fo this big mesh.
// As it is just a mesh, the SPS has all the same properties than any other BJS mesh : not more, not less. It can be scaled, rotated, translated, enlighted, textured, moved, etc.
// The SPS is also a particle system. It provides some methods to manage the particles.
// However it is behavior agnostic. This means it has no emitter, no particle physics, no particle recycler. You have to implement your own behavior.
//
// Full documentation here : <a href="http://doc.babylonjs.com/how_to/Solid_Particle_System">http://doc.babylonjs.com/how_to/Solid_Particle_System</a>
type SolidParticleSystem struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SolidParticleSystem) JSObject() js.Value { return s.p }

// SolidParticleSystem returns a SolidParticleSystem JavaScript class.
func (ba *Babylon) SolidParticleSystem() *SolidParticleSystem {
	p := ba.ctx.Get("SolidParticleSystem")
	return SolidParticleSystemFromJSObject(p, ba.ctx)
}

// SolidParticleSystemFromJSObject returns a wrapped SolidParticleSystem JavaScript class.
func SolidParticleSystemFromJSObject(p js.Value, ctx js.Value) *SolidParticleSystem {
	return &SolidParticleSystem{p: p, ctx: ctx}
}

// SolidParticleSystemArrayToJSArray returns a JavaScript Array for the wrapped array.
func SolidParticleSystemArrayToJSArray(array []*SolidParticleSystem) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSolidParticleSystemOpts contains optional parameters for NewSolidParticleSystem.
type NewSolidParticleSystemOpts struct {
	Options map[string]interface{}
}

// NewSolidParticleSystem returns a new SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#constructor
func (ba *Babylon) NewSolidParticleSystem(name string, scene *Scene, opts *NewSolidParticleSystemOpts) *SolidParticleSystem {
	if opts == nil {
		opts = &NewSolidParticleSystemOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, scene.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	p := ba.ctx.Get("SolidParticleSystem").New(args...)
	return SolidParticleSystemFromJSObject(p, ba.ctx)
}

// SolidParticleSystemAddShapeOpts contains optional parameters for SolidParticleSystem.AddShape.
type SolidParticleSystemAddShapeOpts struct {
	Options map[string]interface{}
}

// AddShape calls the AddShape method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#addshape
func (s *SolidParticleSystem) AddShape(mesh *Mesh, nb float64, opts *SolidParticleSystemAddShapeOpts) float64 {
	if opts == nil {
		opts = &SolidParticleSystemAddShapeOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	args = append(args, nb)

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("addShape", args...)
	return retVal.Float()
}

// SolidParticleSystemAfterUpdateParticlesOpts contains optional parameters for SolidParticleSystem.AfterUpdateParticles.
type SolidParticleSystemAfterUpdateParticlesOpts struct {
	Start  *float64
	Stop   *float64
	Update *bool
}

// AfterUpdateParticles calls the AfterUpdateParticles method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#afterupdateparticles
func (s *SolidParticleSystem) AfterUpdateParticles(opts *SolidParticleSystemAfterUpdateParticlesOpts) {
	if opts == nil {
		opts = &SolidParticleSystemAfterUpdateParticlesOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Start == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Start)
	}
	if opts.Stop == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stop)
	}
	if opts.Update == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Update)
	}

	s.p.Call("afterUpdateParticles", args...)
}

// SolidParticleSystemBeforeUpdateParticlesOpts contains optional parameters for SolidParticleSystem.BeforeUpdateParticles.
type SolidParticleSystemBeforeUpdateParticlesOpts struct {
	Start  *float64
	Stop   *float64
	Update *bool
}

// BeforeUpdateParticles calls the BeforeUpdateParticles method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#beforeupdateparticles
func (s *SolidParticleSystem) BeforeUpdateParticles(opts *SolidParticleSystemBeforeUpdateParticlesOpts) {
	if opts == nil {
		opts = &SolidParticleSystemBeforeUpdateParticlesOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Start == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Start)
	}
	if opts.Stop == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stop)
	}
	if opts.Update == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Update)
	}

	s.p.Call("beforeUpdateParticles", args...)
}

// BuildMesh calls the BuildMesh method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#buildmesh
func (s *SolidParticleSystem) BuildMesh() *Mesh {

	retVal := s.p.Call("buildMesh")
	return MeshFromJSObject(retVal, s.ctx)
}

// ComputeSubMeshes calls the ComputeSubMeshes method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computesubmeshes
func (s *SolidParticleSystem) ComputeSubMeshes() *SolidParticleSystem {

	retVal := s.p.Call("computeSubMeshes")
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// SolidParticleSystemDigestOpts contains optional parameters for SolidParticleSystem.Digest.
type SolidParticleSystemDigestOpts struct {
	Options map[string]interface{}
}

// Digest calls the Digest method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#digest
func (s *SolidParticleSystem) Digest(mesh *Mesh, opts *SolidParticleSystemDigestOpts) *SolidParticleSystem {
	if opts == nil {
		opts = &SolidParticleSystemDigestOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("digest", args...)
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#dispose
func (s *SolidParticleSystem) Dispose() {

	s.p.Call("dispose")
}

// GetParticleById calls the GetParticleById method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#getparticlebyid
func (s *SolidParticleSystem) GetParticleById(id float64) *SolidParticle {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := s.p.Call("getParticleById", args...)
	return SolidParticleFromJSObject(retVal, s.ctx)
}

// GetParticlesByShapeId calls the GetParticlesByShapeId method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#getparticlesbyshapeid
func (s *SolidParticleSystem) GetParticlesByShapeId(shapeId float64) []*SolidParticle {

	args := make([]interface{}, 0, 1+0)

	args = append(args, shapeId)

	retVal := s.p.Call("getParticlesByShapeId", args...)
	result := []*SolidParticle{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SolidParticleFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// GetParticlesByShapeIdToRef calls the GetParticlesByShapeIdToRef method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#getparticlesbyshapeidtoref
func (s *SolidParticleSystem) GetParticlesByShapeIdToRef(shapeId float64, ref []*SolidParticle) *SolidParticleSystem {

	args := make([]interface{}, 0, 2+0)

	args = append(args, shapeId)

	args = append(args, SolidParticleArrayToJSArray(ref))

	retVal := s.p.Call("getParticlesByShapeIdToRef", args...)
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// InitParticles calls the InitParticles method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#initparticles
func (s *SolidParticleSystem) InitParticles() {

	s.p.Call("initParticles")
}

// InsertParticlesFromArray calls the InsertParticlesFromArray method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#insertparticlesfromarray
func (s *SolidParticleSystem) InsertParticlesFromArray(solidParticleArray []*SolidParticle) *SolidParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, SolidParticleArrayToJSArray(solidParticleArray))

	retVal := s.p.Call("insertParticlesFromArray", args...)
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// SolidParticleSystemRebuildMeshOpts contains optional parameters for SolidParticleSystem.RebuildMesh.
type SolidParticleSystemRebuildMeshOpts struct {
	Reset *bool
}

// RebuildMesh calls the RebuildMesh method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#rebuildmesh
func (s *SolidParticleSystem) RebuildMesh(opts *SolidParticleSystemRebuildMeshOpts) *SolidParticleSystem {
	if opts == nil {
		opts = &SolidParticleSystemRebuildMeshOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Reset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reset)
	}

	retVal := s.p.Call("rebuildMesh", args...)
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// RecycleParticle calls the RecycleParticle method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#recycleparticle
func (s *SolidParticleSystem) RecycleParticle(particle *SolidParticle) *SolidParticle {

	args := make([]interface{}, 0, 1+0)

	if particle == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, particle.JSObject())
	}

	retVal := s.p.Call("recycleParticle", args...)
	return SolidParticleFromJSObject(retVal, s.ctx)
}

// RefreshVisibleSize calls the RefreshVisibleSize method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#refreshvisiblesize
func (s *SolidParticleSystem) RefreshVisibleSize() *SolidParticleSystem {

	retVal := s.p.Call("refreshVisibleSize")
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// RemoveParticles calls the RemoveParticles method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#removeparticles
func (s *SolidParticleSystem) RemoveParticles(start float64, end float64) []*SolidParticle {

	args := make([]interface{}, 0, 2+0)

	args = append(args, start)

	args = append(args, end)

	retVal := s.p.Call("removeParticles", args...)
	result := []*SolidParticle{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SolidParticleFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetMultiMaterial calls the SetMultiMaterial method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#setmultimaterial
func (s *SolidParticleSystem) SetMultiMaterial(materials []*Material) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, MaterialArrayToJSArray(materials))

	s.p.Call("setMultiMaterial", args...)
}

// SolidParticleSystemSetParticlesOpts contains optional parameters for SolidParticleSystem.SetParticles.
type SolidParticleSystemSetParticlesOpts struct {
	Start  *float64
	End    *float64
	Update *bool
}

// SetParticles calls the SetParticles method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#setparticles
func (s *SolidParticleSystem) SetParticles(opts *SolidParticleSystemSetParticlesOpts) *SolidParticleSystem {
	if opts == nil {
		opts = &SolidParticleSystemSetParticlesOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Start == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Start)
	}
	if opts.End == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.End)
	}
	if opts.Update == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Update)
	}

	retVal := s.p.Call("setParticles", args...)
	return SolidParticleSystemFromJSObject(retVal, s.ctx)
}

// SetVisibilityBox calls the SetVisibilityBox method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#setvisibilitybox
func (s *SolidParticleSystem) SetVisibilityBox(size float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, size)

	s.p.Call("setVisibilityBox", args...)
}

// UpdateParticle calls the UpdateParticle method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#updateparticle
func (s *SolidParticleSystem) UpdateParticle(particle *SolidParticle) *SolidParticle {

	args := make([]interface{}, 0, 1+0)

	if particle == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, particle.JSObject())
	}

	retVal := s.p.Call("updateParticle", args...)
	return SolidParticleFromJSObject(retVal, s.ctx)
}

// UpdateParticleVertex calls the UpdateParticleVertex method on the SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#updateparticlevertex
func (s *SolidParticleSystem) UpdateParticleVertex(particle *SolidParticle, vertex *Vector3, pt float64) *Vector3 {

	args := make([]interface{}, 0, 3+0)

	if particle == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, particle.JSObject())
	}

	if vertex == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vertex.JSObject())
	}

	args = append(args, pt)

	retVal := s.p.Call("updateParticleVertex", args...)
	return Vector3FromJSObject(retVal, s.ctx)
}

// AutoUpdateSubMeshes returns the AutoUpdateSubMeshes property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#autoupdatesubmeshes
func (s *SolidParticleSystem) AutoUpdateSubMeshes() bool {
	retVal := s.p.Get("autoUpdateSubMeshes")
	return retVal.Bool()
}

// SetAutoUpdateSubMeshes sets the AutoUpdateSubMeshes property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#autoupdatesubmeshes
func (s *SolidParticleSystem) SetAutoUpdateSubMeshes(autoUpdateSubMeshes bool) *SolidParticleSystem {
	s.p.Set("autoUpdateSubMeshes", autoUpdateSubMeshes)
	return s
}

// Billboard returns the Billboard property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#billboard
func (s *SolidParticleSystem) Billboard() bool {
	retVal := s.p.Get("billboard")
	return retVal.Bool()
}

// SetBillboard sets the Billboard property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#billboard
func (s *SolidParticleSystem) SetBillboard(billboard bool) *SolidParticleSystem {
	s.p.Set("billboard", billboard)
	return s
}

// ComputeBoundingBox returns the ComputeBoundingBox property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeboundingbox
func (s *SolidParticleSystem) ComputeBoundingBox() bool {
	retVal := s.p.Get("computeBoundingBox")
	return retVal.Bool()
}

// SetComputeBoundingBox sets the ComputeBoundingBox property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeboundingbox
func (s *SolidParticleSystem) SetComputeBoundingBox(computeBoundingBox bool) *SolidParticleSystem {
	s.p.Set("computeBoundingBox", computeBoundingBox)
	return s
}

// ComputeParticleColor returns the ComputeParticleColor property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlecolor
func (s *SolidParticleSystem) ComputeParticleColor() bool {
	retVal := s.p.Get("computeParticleColor")
	return retVal.Bool()
}

// SetComputeParticleColor sets the ComputeParticleColor property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlecolor
func (s *SolidParticleSystem) SetComputeParticleColor(computeParticleColor bool) *SolidParticleSystem {
	s.p.Set("computeParticleColor", computeParticleColor)
	return s
}

// ComputeParticleRotation returns the ComputeParticleRotation property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlerotation
func (s *SolidParticleSystem) ComputeParticleRotation() bool {
	retVal := s.p.Get("computeParticleRotation")
	return retVal.Bool()
}

// SetComputeParticleRotation sets the ComputeParticleRotation property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlerotation
func (s *SolidParticleSystem) SetComputeParticleRotation(computeParticleRotation bool) *SolidParticleSystem {
	s.p.Set("computeParticleRotation", computeParticleRotation)
	return s
}

// ComputeParticleTexture returns the ComputeParticleTexture property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticletexture
func (s *SolidParticleSystem) ComputeParticleTexture() bool {
	retVal := s.p.Get("computeParticleTexture")
	return retVal.Bool()
}

// SetComputeParticleTexture sets the ComputeParticleTexture property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticletexture
func (s *SolidParticleSystem) SetComputeParticleTexture(computeParticleTexture bool) *SolidParticleSystem {
	s.p.Set("computeParticleTexture", computeParticleTexture)
	return s
}

// ComputeParticleVertex returns the ComputeParticleVertex property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlevertex
func (s *SolidParticleSystem) ComputeParticleVertex() bool {
	retVal := s.p.Get("computeParticleVertex")
	return retVal.Bool()
}

// SetComputeParticleVertex sets the ComputeParticleVertex property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#computeparticlevertex
func (s *SolidParticleSystem) SetComputeParticleVertex(computeParticleVertex bool) *SolidParticleSystem {
	s.p.Set("computeParticleVertex", computeParticleVertex)
	return s
}

// Counter returns the Counter property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#counter
func (s *SolidParticleSystem) Counter() float64 {
	retVal := s.p.Get("counter")
	return retVal.Float()
}

// SetCounter sets the Counter property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#counter
func (s *SolidParticleSystem) SetCounter(counter float64) *SolidParticleSystem {
	s.p.Set("counter", counter)
	return s
}

// DepthSortParticles returns the DepthSortParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#depthsortparticles
func (s *SolidParticleSystem) DepthSortParticles() bool {
	retVal := s.p.Get("depthSortParticles")
	return retVal.Bool()
}

// SetDepthSortParticles sets the DepthSortParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#depthsortparticles
func (s *SolidParticleSystem) SetDepthSortParticles(depthSortParticles bool) *SolidParticleSystem {
	s.p.Set("depthSortParticles", depthSortParticles)
	return s
}

// DepthSortedParticles returns the DepthSortedParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#depthsortedparticles
func (s *SolidParticleSystem) DepthSortedParticles() js.Value {
	retVal := s.p.Get("depthSortedParticles")
	return retVal
}

// SetDepthSortedParticles sets the DepthSortedParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#depthsortedparticles
func (s *SolidParticleSystem) SetDepthSortedParticles(depthSortedParticles js.Value) *SolidParticleSystem {
	s.p.Set("depthSortedParticles", depthSortedParticles)
	return s
}

// Expandable returns the Expandable property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#expandable
func (s *SolidParticleSystem) Expandable() bool {
	retVal := s.p.Get("expandable")
	return retVal.Bool()
}

// SetExpandable sets the Expandable property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#expandable
func (s *SolidParticleSystem) SetExpandable(expandable bool) *SolidParticleSystem {
	s.p.Set("expandable", expandable)
	return s
}

// IsAlwaysVisible returns the IsAlwaysVisible property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#isalwaysvisible
func (s *SolidParticleSystem) IsAlwaysVisible() bool {
	retVal := s.p.Get("isAlwaysVisible")
	return retVal.Bool()
}

// SetIsAlwaysVisible sets the IsAlwaysVisible property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#isalwaysvisible
func (s *SolidParticleSystem) SetIsAlwaysVisible(isAlwaysVisible bool) *SolidParticleSystem {
	s.p.Set("isAlwaysVisible", isAlwaysVisible)
	return s
}

// IsVisibilityBoxLocked returns the IsVisibilityBoxLocked property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#isvisibilityboxlocked
func (s *SolidParticleSystem) IsVisibilityBoxLocked() bool {
	retVal := s.p.Get("isVisibilityBoxLocked")
	return retVal.Bool()
}

// SetIsVisibilityBoxLocked sets the IsVisibilityBoxLocked property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#isvisibilityboxlocked
func (s *SolidParticleSystem) SetIsVisibilityBoxLocked(isVisibilityBoxLocked bool) *SolidParticleSystem {
	s.p.Set("isVisibilityBoxLocked", isVisibilityBoxLocked)
	return s
}

// Materials returns the Materials property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#materials
func (s *SolidParticleSystem) Materials() []*Material {
	retVal := s.p.Get("materials")
	result := []*Material{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, MaterialFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetMaterials sets the Materials property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#materials
func (s *SolidParticleSystem) SetMaterials(materials []*Material) *SolidParticleSystem {
	s.p.Set("materials", materials)
	return s
}

// Mesh returns the Mesh property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#mesh
func (s *SolidParticleSystem) Mesh() *Mesh {
	retVal := s.p.Get("mesh")
	return MeshFromJSObject(retVal, s.ctx)
}

// SetMesh sets the Mesh property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#mesh
func (s *SolidParticleSystem) SetMesh(mesh *Mesh) *SolidParticleSystem {
	s.p.Set("mesh", mesh.JSObject())
	return s
}

// Multimaterial returns the Multimaterial property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#multimaterial
func (s *SolidParticleSystem) Multimaterial() *MultiMaterial {
	retVal := s.p.Get("multimaterial")
	return MultiMaterialFromJSObject(retVal, s.ctx)
}

// SetMultimaterial sets the Multimaterial property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#multimaterial
func (s *SolidParticleSystem) SetMultimaterial(multimaterial *MultiMaterial) *SolidParticleSystem {
	s.p.Set("multimaterial", multimaterial.JSObject())
	return s
}

// MultimaterialEnabled returns the MultimaterialEnabled property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#multimaterialenabled
func (s *SolidParticleSystem) MultimaterialEnabled() bool {
	retVal := s.p.Get("multimaterialEnabled")
	return retVal.Bool()
}

// SetMultimaterialEnabled sets the MultimaterialEnabled property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#multimaterialenabled
func (s *SolidParticleSystem) SetMultimaterialEnabled(multimaterialEnabled bool) *SolidParticleSystem {
	s.p.Set("multimaterialEnabled", multimaterialEnabled)
	return s
}

// Name returns the Name property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#name
func (s *SolidParticleSystem) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#name
func (s *SolidParticleSystem) SetName(name string) *SolidParticleSystem {
	s.p.Set("name", name)
	return s
}

// NbParticles returns the NbParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#nbparticles
func (s *SolidParticleSystem) NbParticles() float64 {
	retVal := s.p.Get("nbParticles")
	return retVal.Float()
}

// SetNbParticles sets the NbParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#nbparticles
func (s *SolidParticleSystem) SetNbParticles(nbParticles float64) *SolidParticleSystem {
	s.p.Set("nbParticles", nbParticles)
	return s
}

// Particles returns the Particles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#particles
func (s *SolidParticleSystem) Particles() []*SolidParticle {
	retVal := s.p.Get("particles")
	result := []*SolidParticle{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SolidParticleFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// PickedParticles returns the PickedParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#pickedparticles
func (s *SolidParticleSystem) PickedParticles() js.Value {
	retVal := s.p.Get("pickedParticles")
	return retVal
}

// SetPickedParticles sets the PickedParticles property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#pickedparticles
func (s *SolidParticleSystem) SetPickedParticles(pickedParticles js.Value) *SolidParticleSystem {
	s.p.Set("pickedParticles", pickedParticles)
	return s
}

// RecomputeNormals returns the RecomputeNormals property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#recomputenormals
func (s *SolidParticleSystem) RecomputeNormals() bool {
	retVal := s.p.Get("recomputeNormals")
	return retVal.Bool()
}

// SetRecomputeNormals sets the RecomputeNormals property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#recomputenormals
func (s *SolidParticleSystem) SetRecomputeNormals(recomputeNormals bool) *SolidParticleSystem {
	s.p.Set("recomputeNormals", recomputeNormals)
	return s
}

// UseModelMaterial returns the UseModelMaterial property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#usemodelmaterial
func (s *SolidParticleSystem) UseModelMaterial() bool {
	retVal := s.p.Get("useModelMaterial")
	return retVal.Bool()
}

// SetUseModelMaterial sets the UseModelMaterial property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#usemodelmaterial
func (s *SolidParticleSystem) SetUseModelMaterial(useModelMaterial bool) *SolidParticleSystem {
	s.p.Set("useModelMaterial", useModelMaterial)
	return s
}

// Vars returns the Vars property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#vars
func (s *SolidParticleSystem) Vars() js.Value {
	retVal := s.p.Get("vars")
	return retVal
}

// SetVars sets the Vars property of class SolidParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem#vars
func (s *SolidParticleSystem) SetVars(vars JSObject) *SolidParticleSystem {
	s.p.Set("vars", vars.JSObject())
	return s
}
