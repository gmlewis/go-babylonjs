// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EnvironmentHelper represents a babylon.js EnvironmentHelper.
// The Environment helper class can be used to add a fully featuread none expensive background to your scene.
// It includes by default a skybox and a ground relying on the BackgroundMaterial.
// It also helps with the default setup of your imageProcessing configuration.
type EnvironmentHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EnvironmentHelper) JSObject() js.Value { return e.p }

// EnvironmentHelper returns a EnvironmentHelper JavaScript class.
func (ba *Babylon) EnvironmentHelper() *EnvironmentHelper {
	p := ba.ctx.Get("EnvironmentHelper")
	return EnvironmentHelperFromJSObject(p, ba.ctx)
}

// EnvironmentHelperFromJSObject returns a wrapped EnvironmentHelper JavaScript class.
func EnvironmentHelperFromJSObject(p js.Value, ctx js.Value) *EnvironmentHelper {
	return &EnvironmentHelper{p: p, ctx: ctx}
}

// EnvironmentHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func EnvironmentHelperArrayToJSArray(array []*EnvironmentHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEnvironmentHelper returns a new EnvironmentHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#constructor
func (ba *Babylon) NewEnvironmentHelper(options *IEnvironmentHelperOptions, scene *Scene) *EnvironmentHelper {

	args := make([]interface{}, 0, 2+0)

	args = append(args, options.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("EnvironmentHelper").New(args...)
	return EnvironmentHelperFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the EnvironmentHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#dispose
func (e *EnvironmentHelper) Dispose() {

	e.p.Call("dispose")
}

// SetMainColor calls the SetMainColor method on the EnvironmentHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#setmaincolor
func (e *EnvironmentHelper) SetMainColor(color *Color3) {

	args := make([]interface{}, 0, 1+0)

	if color == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color.JSObject())
	}

	e.p.Call("setMainColor", args...)
}

// UpdateOptions calls the UpdateOptions method on the EnvironmentHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#updateoptions
func (e *EnvironmentHelper) UpdateOptions(options *IEnvironmentHelperOptions) {

	args := make([]interface{}, 0, 1+0)

	if options == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, options.JSObject())
	}

	e.p.Call("updateOptions", args...)
}

// Ground returns the Ground property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#ground
func (e *EnvironmentHelper) Ground() *Mesh {
	retVal := e.p.Get("ground")
	return MeshFromJSObject(retVal, e.ctx)
}

// SetGround sets the Ground property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#ground
func (e *EnvironmentHelper) SetGround(ground *Mesh) *EnvironmentHelper {
	e.p.Set("ground", ground.JSObject())
	return e
}

// GroundMaterial returns the GroundMaterial property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmaterial
func (e *EnvironmentHelper) GroundMaterial() *BackgroundMaterial {
	retVal := e.p.Get("groundMaterial")
	return BackgroundMaterialFromJSObject(retVal, e.ctx)
}

// SetGroundMaterial sets the GroundMaterial property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmaterial
func (e *EnvironmentHelper) SetGroundMaterial(groundMaterial *BackgroundMaterial) *EnvironmentHelper {
	e.p.Set("groundMaterial", groundMaterial.JSObject())
	return e
}

// GroundMirror returns the GroundMirror property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmirror
func (e *EnvironmentHelper) GroundMirror() *MirrorTexture {
	retVal := e.p.Get("groundMirror")
	return MirrorTextureFromJSObject(retVal, e.ctx)
}

// SetGroundMirror sets the GroundMirror property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmirror
func (e *EnvironmentHelper) SetGroundMirror(groundMirror *MirrorTexture) *EnvironmentHelper {
	e.p.Set("groundMirror", groundMirror.JSObject())
	return e
}

// GroundMirrorRenderList returns the GroundMirrorRenderList property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmirrorrenderlist
func (e *EnvironmentHelper) GroundMirrorRenderList() []*AbstractMesh {
	retVal := e.p.Get("groundMirrorRenderList")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), e.ctx))
	}
	return result
}

// SetGroundMirrorRenderList sets the GroundMirrorRenderList property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundmirrorrenderlist
func (e *EnvironmentHelper) SetGroundMirrorRenderList(groundMirrorRenderList []*AbstractMesh) *EnvironmentHelper {
	e.p.Set("groundMirrorRenderList", groundMirrorRenderList)
	return e
}

// GroundTexture returns the GroundTexture property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundtexture
func (e *EnvironmentHelper) GroundTexture() *BaseTexture {
	retVal := e.p.Get("groundTexture")
	return BaseTextureFromJSObject(retVal, e.ctx)
}

// SetGroundTexture sets the GroundTexture property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#groundtexture
func (e *EnvironmentHelper) SetGroundTexture(groundTexture *BaseTexture) *EnvironmentHelper {
	e.p.Set("groundTexture", groundTexture.JSObject())
	return e
}

// OnErrorObservable returns the OnErrorObservable property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#onerrorobservable
func (e *EnvironmentHelper) OnErrorObservable() *Observable {
	retVal := e.p.Get("onErrorObservable")
	return ObservableFromJSObject(retVal, e.ctx)
}

// SetOnErrorObservable sets the OnErrorObservable property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#onerrorobservable
func (e *EnvironmentHelper) SetOnErrorObservable(onErrorObservable *Observable) *EnvironmentHelper {
	e.p.Set("onErrorObservable", onErrorObservable.JSObject())
	return e
}

// RootMesh returns the RootMesh property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#rootmesh
func (e *EnvironmentHelper) RootMesh() *Mesh {
	retVal := e.p.Get("rootMesh")
	return MeshFromJSObject(retVal, e.ctx)
}

// SetRootMesh sets the RootMesh property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#rootmesh
func (e *EnvironmentHelper) SetRootMesh(rootMesh *Mesh) *EnvironmentHelper {
	e.p.Set("rootMesh", rootMesh.JSObject())
	return e
}

// Skybox returns the Skybox property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skybox
func (e *EnvironmentHelper) Skybox() *Mesh {
	retVal := e.p.Get("skybox")
	return MeshFromJSObject(retVal, e.ctx)
}

// SetSkybox sets the Skybox property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skybox
func (e *EnvironmentHelper) SetSkybox(skybox *Mesh) *EnvironmentHelper {
	e.p.Set("skybox", skybox.JSObject())
	return e
}

// SkyboxMaterial returns the SkyboxMaterial property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skyboxmaterial
func (e *EnvironmentHelper) SkyboxMaterial() *BackgroundMaterial {
	retVal := e.p.Get("skyboxMaterial")
	return BackgroundMaterialFromJSObject(retVal, e.ctx)
}

// SetSkyboxMaterial sets the SkyboxMaterial property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skyboxmaterial
func (e *EnvironmentHelper) SetSkyboxMaterial(skyboxMaterial *BackgroundMaterial) *EnvironmentHelper {
	e.p.Set("skyboxMaterial", skyboxMaterial.JSObject())
	return e
}

// SkyboxTexture returns the SkyboxTexture property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skyboxtexture
func (e *EnvironmentHelper) SkyboxTexture() *BaseTexture {
	retVal := e.p.Get("skyboxTexture")
	return BaseTextureFromJSObject(retVal, e.ctx)
}

// SetSkyboxTexture sets the SkyboxTexture property of class EnvironmentHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper#skyboxtexture
func (e *EnvironmentHelper) SetSkyboxTexture(skyboxTexture *BaseTexture) *EnvironmentHelper {
	e.p.Set("skyboxTexture", skyboxTexture.JSObject())
	return e
}
