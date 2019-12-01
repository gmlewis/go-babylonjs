// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RecastJSPlugin represents a babylon.js RecastJSPlugin.
// RecastJS navigation plugin
type RecastJSPlugin struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RecastJSPlugin) JSObject() js.Value { return r.p }

// RecastJSPlugin returns a RecastJSPlugin JavaScript class.
func (ba *Babylon) RecastJSPlugin() *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin")
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// RecastJSPluginFromJSObject returns a wrapped RecastJSPlugin JavaScript class.
func RecastJSPluginFromJSObject(p js.Value, ctx js.Value) *RecastJSPlugin {
	return &RecastJSPlugin{p: p, ctx: ctx}
}

// NewRecastJSPluginOpts contains optional parameters for NewRecastJSPlugin.
type NewRecastJSPluginOpts struct {
	RecastInjection *interface{}
}

// NewRecastJSPlugin returns a new RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin
func (ba *Babylon) NewRecastJSPlugin(opts *NewRecastJSPluginOpts) *RecastJSPlugin {
	if opts == nil {
		opts = &NewRecastJSPluginOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.RecastInjection == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.RecastInjection)
	}

	p := ba.ctx.Get("RecastJSPlugin").New(args...)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// ComputePath calls the ComputePath method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#computepath
func (r *RecastJSPlugin) ComputePath(start *Vector3, end *Vector3) *Vector3 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, start.JSObject())
	args = append(args, end.JSObject())

	retVal := r.p.Call("computePath", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// CreateCrowd calls the CreateCrowd method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#createcrowd
func (r *RecastJSPlugin) CreateCrowd(maxAgents float64, maxAgentRadius float64, scene *Scene) *ICrowd {

	args := make([]interface{}, 0, 3+0)

	args = append(args, maxAgents)
	args = append(args, maxAgentRadius)
	args = append(args, scene.JSObject())

	retVal := r.p.Call("createCrowd", args...)
	return ICrowdFromJSObject(retVal, r.ctx)
}

// CreateDebugNavMesh calls the CreateDebugNavMesh method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#createdebugnavmesh
func (r *RecastJSPlugin) CreateDebugNavMesh(scene *Scene) *Mesh {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	retVal := r.p.Call("createDebugNavMesh", args...)
	return MeshFromJSObject(retVal, r.ctx)
}

// CreateMavMesh calls the CreateMavMesh method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#createmavmesh
func (r *RecastJSPlugin) CreateMavMesh(meshes []js.Value, parameters *INavMeshParameters) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, meshes)
	args = append(args, parameters.JSObject())

	r.p.Call("createMavMesh", args...)
}

// Dispose calls the Dispose method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#dispose
func (r *RecastJSPlugin) Dispose() {

	args := make([]interface{}, 0, 0+0)

	r.p.Call("dispose", args...)
}

// GetClosestPoint calls the GetClosestPoint method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#getclosestpoint
func (r *RecastJSPlugin) GetClosestPoint(position *Vector3) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, position.JSObject())

	retVal := r.p.Call("getClosestPoint", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// GetDefaultQueryExtent calls the GetDefaultQueryExtent method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#getdefaultqueryextent
func (r *RecastJSPlugin) GetDefaultQueryExtent() *Vector3 {

	args := make([]interface{}, 0, 0+0)

	retVal := r.p.Call("getDefaultQueryExtent", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// GetRandomPointAround calls the GetRandomPointAround method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#getrandompointaround
func (r *RecastJSPlugin) GetRandomPointAround(position *Vector3, maxRadius float64) *Vector3 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, position.JSObject())
	args = append(args, maxRadius)

	retVal := r.p.Call("getRandomPointAround", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// IsSupported calls the IsSupported method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#issupported
func (r *RecastJSPlugin) IsSupported() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := r.p.Call("isSupported", args...)
	return retVal.Bool()
}

// MoveAlong calls the MoveAlong method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#movealong
func (r *RecastJSPlugin) MoveAlong(position *Vector3, destination *Vector3) *Vector3 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, position.JSObject())
	args = append(args, destination.JSObject())

	retVal := r.p.Call("moveAlong", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// SetDefaultQueryExtent calls the SetDefaultQueryExtent method on the RecastJSPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#setdefaultqueryextent
func (r *RecastJSPlugin) SetDefaultQueryExtent(extent *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, extent.JSObject())

	r.p.Call("setDefaultQueryExtent", args...)
}

/*

// BjsRECAST returns the BjsRECAST property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#bjsrecast
func (r *RecastJSPlugin) BjsRECAST(bjsRECAST interface{}) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(bjsRECAST)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// SetBjsRECAST sets the BjsRECAST property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#bjsrecast
func (r *RecastJSPlugin) SetBjsRECAST(bjsRECAST interface{}) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(bjsRECAST)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#name
func (r *RecastJSPlugin) Name(name string) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(name)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#name
func (r *RecastJSPlugin) SetName(name string) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(name)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// NavMesh returns the NavMesh property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#navmesh
func (r *RecastJSPlugin) NavMesh(navMesh interface{}) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(navMesh)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

// SetNavMesh sets the NavMesh property of class RecastJSPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjsplugin#navmesh
func (r *RecastJSPlugin) SetNavMesh(navMesh interface{}) *RecastJSPlugin {
	p := ba.ctx.Get("RecastJSPlugin").New(navMesh)
	return RecastJSPluginFromJSObject(p, ba.ctx)
}

*/
