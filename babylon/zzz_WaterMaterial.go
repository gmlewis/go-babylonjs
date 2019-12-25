// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WaterMaterial represents a babylon.js WaterMaterial.
//
type WaterMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WaterMaterial) JSObject() js.Value { return w.p }

// WaterMaterial returns a WaterMaterial JavaScript class.
func (ba *Babylon) WaterMaterial() *WaterMaterial {
	p := ba.ctx.Get("WaterMaterial")
	return WaterMaterialFromJSObject(p, ba.ctx)
}

// WaterMaterialFromJSObject returns a wrapped WaterMaterial JavaScript class.
func WaterMaterialFromJSObject(p js.Value, ctx js.Value) *WaterMaterial {
	return &WaterMaterial{p: p, ctx: ctx}
}

// WaterMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func WaterMaterialArrayToJSArray(array []*WaterMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWaterMaterialOpts contains optional parameters for NewWaterMaterial.
type NewWaterMaterialOpts struct {
	RenderTargetSize *Vector2
}

// NewWaterMaterial returns a new WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#constructor
func (ba *Babylon) NewWaterMaterial(name string, scene *Scene, opts *NewWaterMaterialOpts) *WaterMaterial {
	if opts == nil {
		opts = &NewWaterMaterialOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, scene.JSObject())

	if opts.RenderTargetSize == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.RenderTargetSize.JSObject())
	}

	p := ba.ctx.Get("WaterMaterial").New(args...)
	return WaterMaterialFromJSObject(p, ba.ctx)
}

// AddToRenderList calls the AddToRenderList method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#addtorenderlist
func (w *WaterMaterial) AddToRenderList(node JSObject) {

	args := make([]interface{}, 0, 1+0)

	if node == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, node.JSObject())
	}

	w.p.Call("addToRenderList", args...)
}

// BindForSubMesh calls the BindForSubMesh method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bindforsubmesh
func (w *WaterMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	w.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#clone
func (w *WaterMaterial) Clone(name string) *WaterMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := w.p.Call("clone", args...)
	return WaterMaterialFromJSObject(retVal, w.ctx)
}

// CreateDefaultMesh calls the CreateDefaultMesh method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#createdefaultmesh
func (w *WaterMaterial) CreateDefaultMesh(name string, scene *Scene) *Mesh {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := w.p.Call("CreateDefaultMesh", args...)
	return MeshFromJSObject(retVal, w.ctx)
}

// WaterMaterialDisposeOpts contains optional parameters for WaterMaterial.Dispose.
type WaterMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#dispose
func (w *WaterMaterial) Dispose(opts *WaterMaterialDisposeOpts) {
	if opts == nil {
		opts = &WaterMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	w.p.Call("dispose", args...)
}

// EnableRenderTargets calls the EnableRenderTargets method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#enablerendertargets
func (w *WaterMaterial) EnableRenderTargets(enable bool) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, enable)

	w.p.Call("enableRenderTargets", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#getactivetextures
func (w *WaterMaterial) GetActiveTextures() []*BaseTexture {

	retVal := w.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#getalphatesttexture
func (w *WaterMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := w.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, w.ctx)
}

// GetAnimatables calls the GetAnimatables method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#getanimatables
func (w *WaterMaterial) GetAnimatables() []*IAnimatable {

	retVal := w.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#getclassname
func (w *WaterMaterial) GetClassName() string {

	retVal := w.p.Call("getClassName")
	return retVal.String()
}

// GetRenderList calls the GetRenderList method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#getrenderlist
func (w *WaterMaterial) GetRenderList() []*AbstractMesh {

	retVal := w.p.Call("getRenderList")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// HasTexture calls the HasTexture method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#hastexture
func (w *WaterMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := w.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// WaterMaterialIsReadyForSubMeshOpts contains optional parameters for WaterMaterial.IsReadyForSubMesh.
type WaterMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#isreadyforsubmesh
func (w *WaterMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *WaterMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &WaterMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := w.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#needalphablending
func (w *WaterMaterial) NeedAlphaBlending() bool {

	retVal := w.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#needalphatesting
func (w *WaterMaterial) NeedAlphaTesting() bool {

	retVal := w.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#parse
func (w *WaterMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *WaterMaterial {

	args := make([]interface{}, 0, 3+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := w.p.Call("Parse", args...)
	return WaterMaterialFromJSObject(retVal, w.ctx)
}

// Serialize calls the Serialize method on the WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#serialize
func (w *WaterMaterial) Serialize() js.Value {

	retVal := w.p.Call("serialize")
	return retVal
}

// BumpAffectsReflection returns the BumpAffectsReflection property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpaffectsreflection
func (w *WaterMaterial) BumpAffectsReflection() bool {
	retVal := w.p.Get("bumpAffectsReflection")
	return retVal.Bool()
}

// SetBumpAffectsReflection sets the BumpAffectsReflection property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpaffectsreflection
func (w *WaterMaterial) SetBumpAffectsReflection(bumpAffectsReflection bool) *WaterMaterial {
	w.p.Set("bumpAffectsReflection", bumpAffectsReflection)
	return w
}

// BumpHeight returns the BumpHeight property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpheight
func (w *WaterMaterial) BumpHeight() float64 {
	retVal := w.p.Get("bumpHeight")
	return retVal.Float()
}

// SetBumpHeight sets the BumpHeight property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpheight
func (w *WaterMaterial) SetBumpHeight(bumpHeight float64) *WaterMaterial {
	w.p.Set("bumpHeight", bumpHeight)
	return w
}

// BumpSuperimpose returns the BumpSuperimpose property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpsuperimpose
func (w *WaterMaterial) BumpSuperimpose() bool {
	retVal := w.p.Get("bumpSuperimpose")
	return retVal.Bool()
}

// SetBumpSuperimpose sets the BumpSuperimpose property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumpsuperimpose
func (w *WaterMaterial) SetBumpSuperimpose(bumpSuperimpose bool) *WaterMaterial {
	w.p.Set("bumpSuperimpose", bumpSuperimpose)
	return w
}

// BumpTexture returns the BumpTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumptexture
func (w *WaterMaterial) BumpTexture() *BaseTexture {
	retVal := w.p.Get("bumpTexture")
	return BaseTextureFromJSObject(retVal, w.ctx)
}

// SetBumpTexture sets the BumpTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#bumptexture
func (w *WaterMaterial) SetBumpTexture(bumpTexture *BaseTexture) *WaterMaterial {
	w.p.Set("bumpTexture", bumpTexture.JSObject())
	return w
}

// ColorBlendFactor returns the ColorBlendFactor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#colorblendfactor
func (w *WaterMaterial) ColorBlendFactor() float64 {
	retVal := w.p.Get("colorBlendFactor")
	return retVal.Float()
}

// SetColorBlendFactor sets the ColorBlendFactor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#colorblendfactor
func (w *WaterMaterial) SetColorBlendFactor(colorBlendFactor float64) *WaterMaterial {
	w.p.Set("colorBlendFactor", colorBlendFactor)
	return w
}

// ColorBlendFactor2 returns the ColorBlendFactor2 property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#colorblendfactor2
func (w *WaterMaterial) ColorBlendFactor2() float64 {
	retVal := w.p.Get("colorBlendFactor2")
	return retVal.Float()
}

// SetColorBlendFactor2 sets the ColorBlendFactor2 property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#colorblendfactor2
func (w *WaterMaterial) SetColorBlendFactor2(colorBlendFactor2 float64) *WaterMaterial {
	w.p.Set("colorBlendFactor2", colorBlendFactor2)
	return w
}

// DiffuseColor returns the DiffuseColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#diffusecolor
func (w *WaterMaterial) DiffuseColor() *Color3 {
	retVal := w.p.Get("diffuseColor")
	return Color3FromJSObject(retVal, w.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#diffusecolor
func (w *WaterMaterial) SetDiffuseColor(diffuseColor *Color3) *WaterMaterial {
	w.p.Set("diffuseColor", diffuseColor.JSObject())
	return w
}

// DisableClipPlane returns the DisableClipPlane property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#disableclipplane
func (w *WaterMaterial) DisableClipPlane() bool {
	retVal := w.p.Get("disableClipPlane")
	return retVal.Bool()
}

// SetDisableClipPlane sets the DisableClipPlane property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#disableclipplane
func (w *WaterMaterial) SetDisableClipPlane(disableClipPlane bool) *WaterMaterial {
	w.p.Set("disableClipPlane", disableClipPlane)
	return w
}

// DisableLighting returns the DisableLighting property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#disablelighting
func (w *WaterMaterial) DisableLighting() bool {
	retVal := w.p.Get("disableLighting")
	return retVal.Bool()
}

// SetDisableLighting sets the DisableLighting property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#disablelighting
func (w *WaterMaterial) SetDisableLighting(disableLighting bool) *WaterMaterial {
	w.p.Set("disableLighting", disableLighting)
	return w
}

// FresnelSeparate returns the FresnelSeparate property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#fresnelseparate
func (w *WaterMaterial) FresnelSeparate() bool {
	retVal := w.p.Get("fresnelSeparate")
	return retVal.Bool()
}

// SetFresnelSeparate sets the FresnelSeparate property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#fresnelseparate
func (w *WaterMaterial) SetFresnelSeparate(fresnelSeparate bool) *WaterMaterial {
	w.p.Set("fresnelSeparate", fresnelSeparate)
	return w
}

// HasRenderTargetTextures returns the HasRenderTargetTextures property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#hasrendertargettextures
func (w *WaterMaterial) HasRenderTargetTextures() bool {
	retVal := w.p.Get("hasRenderTargetTextures")
	return retVal.Bool()
}

// SetHasRenderTargetTextures sets the HasRenderTargetTextures property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#hasrendertargettextures
func (w *WaterMaterial) SetHasRenderTargetTextures(hasRenderTargetTextures bool) *WaterMaterial {
	w.p.Set("hasRenderTargetTextures", hasRenderTargetTextures)
	return w
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#maxsimultaneouslights
func (w *WaterMaterial) MaxSimultaneousLights() float64 {
	retVal := w.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#maxsimultaneouslights
func (w *WaterMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *WaterMaterial {
	w.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return w
}

// ReflectionTexture returns the ReflectionTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#reflectiontexture
func (w *WaterMaterial) ReflectionTexture() *RenderTargetTexture {
	retVal := w.p.Get("reflectionTexture")
	return RenderTargetTextureFromJSObject(retVal, w.ctx)
}

// SetReflectionTexture sets the ReflectionTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#reflectiontexture
func (w *WaterMaterial) SetReflectionTexture(reflectionTexture *RenderTargetTexture) *WaterMaterial {
	w.p.Set("reflectionTexture", reflectionTexture.JSObject())
	return w
}

// RefractionTexture returns the RefractionTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#refractiontexture
func (w *WaterMaterial) RefractionTexture() *RenderTargetTexture {
	retVal := w.p.Get("refractionTexture")
	return RenderTargetTextureFromJSObject(retVal, w.ctx)
}

// SetRefractionTexture sets the RefractionTexture property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#refractiontexture
func (w *WaterMaterial) SetRefractionTexture(refractionTexture *RenderTargetTexture) *WaterMaterial {
	w.p.Set("refractionTexture", refractionTexture.JSObject())
	return w
}

// RenderTargetSize returns the RenderTargetSize property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#rendertargetsize
func (w *WaterMaterial) RenderTargetSize() *Vector2 {
	retVal := w.p.Get("renderTargetSize")
	return Vector2FromJSObject(retVal, w.ctx)
}

// SetRenderTargetSize sets the RenderTargetSize property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#rendertargetsize
func (w *WaterMaterial) SetRenderTargetSize(renderTargetSize *Vector2) *WaterMaterial {
	w.p.Set("renderTargetSize", renderTargetSize.JSObject())
	return w
}

// RenderTargetsEnabled returns the RenderTargetsEnabled property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#rendertargetsenabled
func (w *WaterMaterial) RenderTargetsEnabled() bool {
	retVal := w.p.Get("renderTargetsEnabled")
	return retVal.Bool()
}

// SetRenderTargetsEnabled sets the RenderTargetsEnabled property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#rendertargetsenabled
func (w *WaterMaterial) SetRenderTargetsEnabled(renderTargetsEnabled bool) *WaterMaterial {
	w.p.Set("renderTargetsEnabled", renderTargetsEnabled)
	return w
}

// SpecularColor returns the SpecularColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#specularcolor
func (w *WaterMaterial) SpecularColor() *Color3 {
	retVal := w.p.Get("specularColor")
	return Color3FromJSObject(retVal, w.ctx)
}

// SetSpecularColor sets the SpecularColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#specularcolor
func (w *WaterMaterial) SetSpecularColor(specularColor *Color3) *WaterMaterial {
	w.p.Set("specularColor", specularColor.JSObject())
	return w
}

// SpecularPower returns the SpecularPower property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#specularpower
func (w *WaterMaterial) SpecularPower() float64 {
	retVal := w.p.Get("specularPower")
	return retVal.Float()
}

// SetSpecularPower sets the SpecularPower property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#specularpower
func (w *WaterMaterial) SetSpecularPower(specularPower float64) *WaterMaterial {
	w.p.Set("specularPower", specularPower)
	return w
}

// UseLogarithmicDepth returns the UseLogarithmicDepth property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#uselogarithmicdepth
func (w *WaterMaterial) UseLogarithmicDepth() bool {
	retVal := w.p.Get("useLogarithmicDepth")
	return retVal.Bool()
}

// SetUseLogarithmicDepth sets the UseLogarithmicDepth property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#uselogarithmicdepth
func (w *WaterMaterial) SetUseLogarithmicDepth(useLogarithmicDepth bool) *WaterMaterial {
	w.p.Set("useLogarithmicDepth", useLogarithmicDepth)
	return w
}

// WaterColor returns the WaterColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#watercolor
func (w *WaterMaterial) WaterColor() *Color3 {
	retVal := w.p.Get("waterColor")
	return Color3FromJSObject(retVal, w.ctx)
}

// SetWaterColor sets the WaterColor property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#watercolor
func (w *WaterMaterial) SetWaterColor(waterColor *Color3) *WaterMaterial {
	w.p.Set("waterColor", waterColor.JSObject())
	return w
}

// WaterColor2 returns the WaterColor2 property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#watercolor2
func (w *WaterMaterial) WaterColor2() *Color3 {
	retVal := w.p.Get("waterColor2")
	return Color3FromJSObject(retVal, w.ctx)
}

// SetWaterColor2 sets the WaterColor2 property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#watercolor2
func (w *WaterMaterial) SetWaterColor2(waterColor2 *Color3) *WaterMaterial {
	w.p.Set("waterColor2", waterColor2.JSObject())
	return w
}

// WaveHeight returns the WaveHeight property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#waveheight
func (w *WaterMaterial) WaveHeight() float64 {
	retVal := w.p.Get("waveHeight")
	return retVal.Float()
}

// SetWaveHeight sets the WaveHeight property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#waveheight
func (w *WaterMaterial) SetWaveHeight(waveHeight float64) *WaterMaterial {
	w.p.Set("waveHeight", waveHeight)
	return w
}

// WaveLength returns the WaveLength property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#wavelength
func (w *WaterMaterial) WaveLength() float64 {
	retVal := w.p.Get("waveLength")
	return retVal.Float()
}

// SetWaveLength sets the WaveLength property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#wavelength
func (w *WaterMaterial) SetWaveLength(waveLength float64) *WaterMaterial {
	w.p.Set("waveLength", waveLength)
	return w
}

// WaveSpeed returns the WaveSpeed property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#wavespeed
func (w *WaterMaterial) WaveSpeed() float64 {
	retVal := w.p.Get("waveSpeed")
	return retVal.Float()
}

// SetWaveSpeed sets the WaveSpeed property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#wavespeed
func (w *WaterMaterial) SetWaveSpeed(waveSpeed float64) *WaterMaterial {
	w.p.Set("waveSpeed", waveSpeed)
	return w
}

// WindDirection returns the WindDirection property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#winddirection
func (w *WaterMaterial) WindDirection() *Vector2 {
	retVal := w.p.Get("windDirection")
	return Vector2FromJSObject(retVal, w.ctx)
}

// SetWindDirection sets the WindDirection property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#winddirection
func (w *WaterMaterial) SetWindDirection(windDirection *Vector2) *WaterMaterial {
	w.p.Set("windDirection", windDirection.JSObject())
	return w
}

// WindForce returns the WindForce property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#windforce
func (w *WaterMaterial) WindForce() float64 {
	retVal := w.p.Get("windForce")
	return retVal.Float()
}

// SetWindForce sets the WindForce property of class WaterMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial#windforce
func (w *WaterMaterial) SetWindForce(windForce float64) *WaterMaterial {
	w.p.Set("windForce", windForce)
	return w
}
