// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VolumetricLightScatteringPostProcess represents a babylon.js VolumetricLightScatteringPostProcess.
// Inspired by <a href="http://http.developer.nvidia.com/GPUGems3/gpugems3_ch13.html">http://http.developer.nvidia.com/GPUGems3/gpugems3_ch13.html</a>
type VolumetricLightScatteringPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VolumetricLightScatteringPostProcess) JSObject() js.Value { return v.p }

// VolumetricLightScatteringPostProcess returns a VolumetricLightScatteringPostProcess JavaScript class.
func (ba *Babylon) VolumetricLightScatteringPostProcess() *VolumetricLightScatteringPostProcess {
	p := ba.ctx.Get("VolumetricLightScatteringPostProcess")
	return VolumetricLightScatteringPostProcessFromJSObject(p, ba.ctx)
}

// VolumetricLightScatteringPostProcessFromJSObject returns a wrapped VolumetricLightScatteringPostProcess JavaScript class.
func VolumetricLightScatteringPostProcessFromJSObject(p js.Value, ctx js.Value) *VolumetricLightScatteringPostProcess {
	return &VolumetricLightScatteringPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// VolumetricLightScatteringPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func VolumetricLightScatteringPostProcessArrayToJSArray(array []*VolumetricLightScatteringPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVolumetricLightScatteringPostProcessOpts contains optional parameters for NewVolumetricLightScatteringPostProcess.
type NewVolumetricLightScatteringPostProcessOpts struct {
	Mesh         *Mesh
	Samples      *float64
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
	Scene        *Scene
}

// NewVolumetricLightScatteringPostProcess returns a new VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#constructor
func (ba *Babylon) NewVolumetricLightScatteringPostProcess(name string, ratio JSObject, camera *Camera, opts *NewVolumetricLightScatteringPostProcessOpts) *VolumetricLightScatteringPostProcess {
	if opts == nil {
		opts = &NewVolumetricLightScatteringPostProcessOpts{}
	}

	args := make([]interface{}, 0, 3+6)

	args = append(args, name)
	args = append(args, ratio.JSObject())
	args = append(args, camera.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}
	if opts.Samples == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Samples)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Reusable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reusable)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	p := ba.ctx.Get("VolumetricLightScatteringPostProcess").New(args...)
	return VolumetricLightScatteringPostProcessFromJSObject(p, ba.ctx)
}

// CreateDefaultMesh calls the CreateDefaultMesh method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#createdefaultmesh
func (v *VolumetricLightScatteringPostProcess) CreateDefaultMesh(name string, scene *Scene) *Mesh {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := v.p.Call("CreateDefaultMesh", args...)
	return MeshFromJSObject(retVal, v.ctx)
}

// Dispose calls the Dispose method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#dispose
func (v *VolumetricLightScatteringPostProcess) Dispose(camera *Camera) {

	args := make([]interface{}, 0, 1+0)

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	v.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#getclassname
func (v *VolumetricLightScatteringPostProcess) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// GetCustomMeshPosition calls the GetCustomMeshPosition method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#getcustommeshposition
func (v *VolumetricLightScatteringPostProcess) GetCustomMeshPosition() *Vector3 {

	retVal := v.p.Call("getCustomMeshPosition")
	return Vector3FromJSObject(retVal, v.ctx)
}

// GetPass calls the GetPass method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#getpass
func (v *VolumetricLightScatteringPostProcess) GetPass() *RenderTargetTexture {

	retVal := v.p.Call("getPass")
	return RenderTargetTextureFromJSObject(retVal, v.ctx)
}

// SetCustomMeshPosition calls the SetCustomMeshPosition method on the VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#setcustommeshposition
func (v *VolumetricLightScatteringPostProcess) SetCustomMeshPosition(position *Vector3) {

	args := make([]interface{}, 0, 1+0)

	if position == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, position.JSObject())
	}

	v.p.Call("setCustomMeshPosition", args...)
}

// AttachedNode returns the AttachedNode property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#attachednode
func (v *VolumetricLightScatteringPostProcess) AttachedNode() js.Value {
	retVal := v.p.Get("attachedNode")
	return retVal
}

// SetAttachedNode sets the AttachedNode property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#attachednode
func (v *VolumetricLightScatteringPostProcess) SetAttachedNode(attachedNode js.Value) *VolumetricLightScatteringPostProcess {
	v.p.Set("attachedNode", attachedNode)
	return v
}

// CustomMeshPosition returns the CustomMeshPosition property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#custommeshposition
func (v *VolumetricLightScatteringPostProcess) CustomMeshPosition() *Vector3 {
	retVal := v.p.Get("customMeshPosition")
	return Vector3FromJSObject(retVal, v.ctx)
}

// Decay returns the Decay property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#decay
func (v *VolumetricLightScatteringPostProcess) Decay() float64 {
	retVal := v.p.Get("decay")
	return retVal.Float()
}

// SetDecay sets the Decay property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#decay
func (v *VolumetricLightScatteringPostProcess) SetDecay(decay float64) *VolumetricLightScatteringPostProcess {
	v.p.Set("decay", decay)
	return v
}

// Density returns the Density property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#density
func (v *VolumetricLightScatteringPostProcess) Density() float64 {
	retVal := v.p.Get("density")
	return retVal.Float()
}

// SetDensity sets the Density property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#density
func (v *VolumetricLightScatteringPostProcess) SetDensity(density float64) *VolumetricLightScatteringPostProcess {
	v.p.Set("density", density)
	return v
}

// ExcludedMeshes returns the ExcludedMeshes property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#excludedmeshes
func (v *VolumetricLightScatteringPostProcess) ExcludedMeshes() []*AbstractMesh {
	retVal := v.p.Get("excludedMeshes")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), v.ctx))
	}
	return result
}

// SetExcludedMeshes sets the ExcludedMeshes property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#excludedmeshes
func (v *VolumetricLightScatteringPostProcess) SetExcludedMeshes(excludedMeshes []*AbstractMesh) *VolumetricLightScatteringPostProcess {
	v.p.Set("excludedMeshes", excludedMeshes)
	return v
}

// Exposure returns the Exposure property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#exposure
func (v *VolumetricLightScatteringPostProcess) Exposure() float64 {
	retVal := v.p.Get("exposure")
	return retVal.Float()
}

// SetExposure sets the Exposure property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#exposure
func (v *VolumetricLightScatteringPostProcess) SetExposure(exposure float64) *VolumetricLightScatteringPostProcess {
	v.p.Set("exposure", exposure)
	return v
}

// Invert returns the Invert property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#invert
func (v *VolumetricLightScatteringPostProcess) Invert() bool {
	retVal := v.p.Get("invert")
	return retVal.Bool()
}

// SetInvert sets the Invert property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#invert
func (v *VolumetricLightScatteringPostProcess) SetInvert(invert bool) *VolumetricLightScatteringPostProcess {
	v.p.Set("invert", invert)
	return v
}

// Mesh returns the Mesh property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#mesh
func (v *VolumetricLightScatteringPostProcess) Mesh() *Mesh {
	retVal := v.p.Get("mesh")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetMesh sets the Mesh property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#mesh
func (v *VolumetricLightScatteringPostProcess) SetMesh(mesh *Mesh) *VolumetricLightScatteringPostProcess {
	v.p.Set("mesh", mesh.JSObject())
	return v
}

// UseCustomMeshPosition returns the UseCustomMeshPosition property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#usecustommeshposition
func (v *VolumetricLightScatteringPostProcess) UseCustomMeshPosition() bool {
	retVal := v.p.Get("useCustomMeshPosition")
	return retVal.Bool()
}

// SetUseCustomMeshPosition sets the UseCustomMeshPosition property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#usecustommeshposition
func (v *VolumetricLightScatteringPostProcess) SetUseCustomMeshPosition(useCustomMeshPosition bool) *VolumetricLightScatteringPostProcess {
	v.p.Set("useCustomMeshPosition", useCustomMeshPosition)
	return v
}

// Weight returns the Weight property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#weight
func (v *VolumetricLightScatteringPostProcess) Weight() float64 {
	retVal := v.p.Get("weight")
	return retVal.Float()
}

// SetWeight sets the Weight property of class VolumetricLightScatteringPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess#weight
func (v *VolumetricLightScatteringPostProcess) SetWeight(weight float64) *VolumetricLightScatteringPostProcess {
	v.p.Set("weight", weight)
	return v
}
