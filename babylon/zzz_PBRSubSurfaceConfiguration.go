// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRSubSurfaceConfiguration represents a babylon.js PBRSubSurfaceConfiguration.
// Define the code related to the sub surface parameters of the pbr material.
type PBRSubSurfaceConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRSubSurfaceConfiguration) JSObject() js.Value { return p.p }

// PBRSubSurfaceConfiguration returns a PBRSubSurfaceConfiguration JavaScript class.
func (ba *Babylon) PBRSubSurfaceConfiguration() *PBRSubSurfaceConfiguration {
	p := ba.ctx.Get("PBRSubSurfaceConfiguration")
	return PBRSubSurfaceConfigurationFromJSObject(p, ba.ctx)
}

// PBRSubSurfaceConfigurationFromJSObject returns a wrapped PBRSubSurfaceConfiguration JavaScript class.
func PBRSubSurfaceConfigurationFromJSObject(p js.Value, ctx js.Value) *PBRSubSurfaceConfiguration {
	return &PBRSubSurfaceConfiguration{p: p, ctx: ctx}
}

// PBRSubSurfaceConfigurationArrayToJSArray returns a JavaScript Array for the wrapped array.
func PBRSubSurfaceConfigurationArrayToJSArray(array []*PBRSubSurfaceConfiguration) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPBRSubSurfaceConfiguration returns a new PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration
func (ba *Babylon) NewPBRSubSurfaceConfiguration(markAllSubMeshesAsTexturesDirty func()) *PBRSubSurfaceConfiguration {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { markAllSubMeshesAsTexturesDirty(); return nil }))

	p := ba.ctx.Get("PBRSubSurfaceConfiguration").New(args...)
	return PBRSubSurfaceConfigurationFromJSObject(p, ba.ctx)
}

// AddFallbacks calls the AddFallbacks method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#addfallbacks
func (p *PBRSubSurfaceConfiguration) AddFallbacks(defines js.Value, fallbacks *EffectFallbacks, currentRank float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, defines)
	args = append(args, fallbacks.JSObject())
	args = append(args, currentRank)

	retVal := p.p.Call("AddFallbacks", args...)
	return retVal.Float()
}

// AddSamplers calls the AddSamplers method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#addsamplers
func (p *PBRSubSurfaceConfiguration) AddSamplers(samplers []string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, samplers)

	p.p.Call("AddSamplers", args...)
}

// AddUniforms calls the AddUniforms method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#adduniforms
func (p *PBRSubSurfaceConfiguration) AddUniforms(uniforms []string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uniforms)

	p.p.Call("AddUniforms", args...)
}

// BindForSubMesh calls the BindForSubMesh method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#bindforsubmesh
func (p *PBRSubSurfaceConfiguration) BindForSubMesh(uniformBuffer *UniformBuffer, scene *Scene, engine *Engine, isFrozen bool, lodBasedMicrosurface bool) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, uniformBuffer.JSObject())
	args = append(args, scene.JSObject())
	args = append(args, engine.JSObject())
	args = append(args, isFrozen)
	args = append(args, lodBasedMicrosurface)

	p.p.Call("bindForSubMesh", args...)
}

// CopyTo calls the CopyTo method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#copyto
func (p *PBRSubSurfaceConfiguration) CopyTo(configuration *PBRSubSurfaceConfiguration) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, configuration.JSObject())

	p.p.Call("copyTo", args...)
}

// PBRSubSurfaceConfigurationDisposeOpts contains optional parameters for PBRSubSurfaceConfiguration.Dispose.
type PBRSubSurfaceConfigurationDisposeOpts struct {
	ForceDisposeTextures *bool
}

// Dispose calls the Dispose method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#dispose
func (p *PBRSubSurfaceConfiguration) Dispose(opts *PBRSubSurfaceConfigurationDisposeOpts) {
	if opts == nil {
		opts = &PBRSubSurfaceConfigurationDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeTextures)
	}

	p.p.Call("dispose", args...)
}

// FillRenderTargetTextures calls the FillRenderTargetTextures method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#fillrendertargettextures
func (p *PBRSubSurfaceConfiguration) FillRenderTargetTextures(renderTargets *SmartArray) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, renderTargets.JSObject())

	p.p.Call("fillRenderTargetTextures", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#getactivetextures
func (p *PBRSubSurfaceConfiguration) GetActiveTextures(activeTextures []*BaseTexture) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, BaseTextureArrayToJSArray(activeTextures))

	p.p.Call("getActiveTextures", args...)
}

// GetAnimatables calls the GetAnimatables method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#getanimatables
func (p *PBRSubSurfaceConfiguration) GetAnimatables(animatables []*IAnimatable) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, IAnimatableArrayToJSArray(animatables))

	p.p.Call("getAnimatables", args...)
}

// GetClassName calls the GetClassName method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#getclassname
func (p *PBRSubSurfaceConfiguration) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// HasRenderTargetTextures calls the HasRenderTargetTextures method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#hasrendertargettextures
func (p *PBRSubSurfaceConfiguration) HasRenderTargetTextures() bool {

	retVal := p.p.Call("hasRenderTargetTextures")
	return retVal.Bool()
}

// HasTexture calls the HasTexture method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#hastexture
func (p *PBRSubSurfaceConfiguration) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, texture.JSObject())

	retVal := p.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#isreadyforsubmesh
func (p *PBRSubSurfaceConfiguration) IsReadyForSubMesh(defines js.Value, scene *Scene) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)
	args = append(args, scene.JSObject())

	retVal := p.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// Parse calls the Parse method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#parse
func (p *PBRSubSurfaceConfiguration) Parse(source interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, source)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	p.p.Call("parse", args...)
}

// PrepareDefines calls the PrepareDefines method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#preparedefines
func (p *PBRSubSurfaceConfiguration) PrepareDefines(defines js.Value, scene *Scene) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)
	args = append(args, scene.JSObject())

	p.p.Call("prepareDefines", args...)
}

// PrepareUniformBuffer calls the PrepareUniformBuffer method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#prepareuniformbuffer
func (p *PBRSubSurfaceConfiguration) PrepareUniformBuffer(uniformBuffer *UniformBuffer) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uniformBuffer.JSObject())

	p.p.Call("PrepareUniformBuffer", args...)
}

// Serialize calls the Serialize method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#serialize
func (p *PBRSubSurfaceConfiguration) Serialize() interface{} {

	retVal := p.p.Call("serialize")
	return retVal
}

// Unbind calls the Unbind method on the PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#unbind
func (p *PBRSubSurfaceConfiguration) Unbind(activeEffect *Effect) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, activeEffect.JSObject())

	retVal := p.p.Call("unbind", args...)
	return retVal.Bool()
}

// DiffusionDistance returns the DiffusionDistance property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#diffusiondistance
func (p *PBRSubSurfaceConfiguration) DiffusionDistance() *Color3 {
	retVal := p.p.Get("diffusionDistance")
	return Color3FromJSObject(retVal, p.ctx)
}

// SetDiffusionDistance sets the DiffusionDistance property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#diffusiondistance
func (p *PBRSubSurfaceConfiguration) SetDiffusionDistance(diffusionDistance *Color3) *PBRSubSurfaceConfiguration {
	p.p.Set("diffusionDistance", diffusionDistance.JSObject())
	return p
}

// DisableAlphaBlending returns the DisableAlphaBlending property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#disablealphablending
func (p *PBRSubSurfaceConfiguration) DisableAlphaBlending() bool {
	retVal := p.p.Get("disableAlphaBlending")
	return retVal.Bool()
}

// SetDisableAlphaBlending sets the DisableAlphaBlending property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#disablealphablending
func (p *PBRSubSurfaceConfiguration) SetDisableAlphaBlending(disableAlphaBlending bool) *PBRSubSurfaceConfiguration {
	p.p.Set("disableAlphaBlending", disableAlphaBlending)
	return p
}

// IndexOfRefraction returns the IndexOfRefraction property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#indexofrefraction
func (p *PBRSubSurfaceConfiguration) IndexOfRefraction() float64 {
	retVal := p.p.Get("indexOfRefraction")
	return retVal.Float()
}

// SetIndexOfRefraction sets the IndexOfRefraction property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#indexofrefraction
func (p *PBRSubSurfaceConfiguration) SetIndexOfRefraction(indexOfRefraction float64) *PBRSubSurfaceConfiguration {
	p.p.Set("indexOfRefraction", indexOfRefraction)
	return p
}

// InvertRefractionY returns the InvertRefractionY property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#invertrefractiony
func (p *PBRSubSurfaceConfiguration) InvertRefractionY() bool {
	retVal := p.p.Get("invertRefractionY")
	return retVal.Bool()
}

// SetInvertRefractionY sets the InvertRefractionY property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#invertrefractiony
func (p *PBRSubSurfaceConfiguration) SetInvertRefractionY(invertRefractionY bool) *PBRSubSurfaceConfiguration {
	p.p.Set("invertRefractionY", invertRefractionY)
	return p
}

// IsRefractionEnabled returns the IsRefractionEnabled property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#isrefractionenabled
func (p *PBRSubSurfaceConfiguration) IsRefractionEnabled() bool {
	retVal := p.p.Get("isRefractionEnabled")
	return retVal.Bool()
}

// SetIsRefractionEnabled sets the IsRefractionEnabled property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#isrefractionenabled
func (p *PBRSubSurfaceConfiguration) SetIsRefractionEnabled(isRefractionEnabled bool) *PBRSubSurfaceConfiguration {
	p.p.Set("isRefractionEnabled", isRefractionEnabled)
	return p
}

// IsTranslucencyEnabled returns the IsTranslucencyEnabled property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#istranslucencyenabled
func (p *PBRSubSurfaceConfiguration) IsTranslucencyEnabled() bool {
	retVal := p.p.Get("isTranslucencyEnabled")
	return retVal.Bool()
}

// SetIsTranslucencyEnabled sets the IsTranslucencyEnabled property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#istranslucencyenabled
func (p *PBRSubSurfaceConfiguration) SetIsTranslucencyEnabled(isTranslucencyEnabled bool) *PBRSubSurfaceConfiguration {
	p.p.Set("isTranslucencyEnabled", isTranslucencyEnabled)
	return p
}

// LinkRefractionWithTransparency returns the LinkRefractionWithTransparency property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#linkrefractionwithtransparency
func (p *PBRSubSurfaceConfiguration) LinkRefractionWithTransparency() bool {
	retVal := p.p.Get("linkRefractionWithTransparency")
	return retVal.Bool()
}

// SetLinkRefractionWithTransparency sets the LinkRefractionWithTransparency property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#linkrefractionwithtransparency
func (p *PBRSubSurfaceConfiguration) SetLinkRefractionWithTransparency(linkRefractionWithTransparency bool) *PBRSubSurfaceConfiguration {
	p.p.Set("linkRefractionWithTransparency", linkRefractionWithTransparency)
	return p
}

// MaximumThickness returns the MaximumThickness property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#maximumthickness
func (p *PBRSubSurfaceConfiguration) MaximumThickness() float64 {
	retVal := p.p.Get("maximumThickness")
	return retVal.Float()
}

// SetMaximumThickness sets the MaximumThickness property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#maximumthickness
func (p *PBRSubSurfaceConfiguration) SetMaximumThickness(maximumThickness float64) *PBRSubSurfaceConfiguration {
	p.p.Set("maximumThickness", maximumThickness)
	return p
}

// MinimumThickness returns the MinimumThickness property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#minimumthickness
func (p *PBRSubSurfaceConfiguration) MinimumThickness() float64 {
	retVal := p.p.Get("minimumThickness")
	return retVal.Float()
}

// SetMinimumThickness sets the MinimumThickness property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#minimumthickness
func (p *PBRSubSurfaceConfiguration) SetMinimumThickness(minimumThickness float64) *PBRSubSurfaceConfiguration {
	p.p.Set("minimumThickness", minimumThickness)
	return p
}

// RefractionIntensity returns the RefractionIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#refractionintensity
func (p *PBRSubSurfaceConfiguration) RefractionIntensity() float64 {
	retVal := p.p.Get("refractionIntensity")
	return retVal.Float()
}

// SetRefractionIntensity sets the RefractionIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#refractionintensity
func (p *PBRSubSurfaceConfiguration) SetRefractionIntensity(refractionIntensity float64) *PBRSubSurfaceConfiguration {
	p.p.Set("refractionIntensity", refractionIntensity)
	return p
}

// RefractionTexture returns the RefractionTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#refractiontexture
func (p *PBRSubSurfaceConfiguration) RefractionTexture() *BaseTexture {
	retVal := p.p.Get("refractionTexture")
	return BaseTextureFromJSObject(retVal, p.ctx)
}

// SetRefractionTexture sets the RefractionTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#refractiontexture
func (p *PBRSubSurfaceConfiguration) SetRefractionTexture(refractionTexture *BaseTexture) *PBRSubSurfaceConfiguration {
	p.p.Set("refractionTexture", refractionTexture.JSObject())
	return p
}

// ScatteringIntensity returns the ScatteringIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#scatteringintensity
func (p *PBRSubSurfaceConfiguration) ScatteringIntensity() float64 {
	retVal := p.p.Get("scatteringIntensity")
	return retVal.Float()
}

// SetScatteringIntensity sets the ScatteringIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#scatteringintensity
func (p *PBRSubSurfaceConfiguration) SetScatteringIntensity(scatteringIntensity float64) *PBRSubSurfaceConfiguration {
	p.p.Set("scatteringIntensity", scatteringIntensity)
	return p
}

// ThicknessTexture returns the ThicknessTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#thicknesstexture
func (p *PBRSubSurfaceConfiguration) ThicknessTexture() *BaseTexture {
	retVal := p.p.Get("thicknessTexture")
	return BaseTextureFromJSObject(retVal, p.ctx)
}

// SetThicknessTexture sets the ThicknessTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#thicknesstexture
func (p *PBRSubSurfaceConfiguration) SetThicknessTexture(thicknessTexture *BaseTexture) *PBRSubSurfaceConfiguration {
	p.p.Set("thicknessTexture", thicknessTexture.JSObject())
	return p
}

// TintColor returns the TintColor property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#tintcolor
func (p *PBRSubSurfaceConfiguration) TintColor() *Color3 {
	retVal := p.p.Get("tintColor")
	return Color3FromJSObject(retVal, p.ctx)
}

// SetTintColor sets the TintColor property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#tintcolor
func (p *PBRSubSurfaceConfiguration) SetTintColor(tintColor *Color3) *PBRSubSurfaceConfiguration {
	p.p.Set("tintColor", tintColor.JSObject())
	return p
}

// TintColorAtDistance returns the TintColorAtDistance property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#tintcoloratdistance
func (p *PBRSubSurfaceConfiguration) TintColorAtDistance() float64 {
	retVal := p.p.Get("tintColorAtDistance")
	return retVal.Float()
}

// SetTintColorAtDistance sets the TintColorAtDistance property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#tintcoloratdistance
func (p *PBRSubSurfaceConfiguration) SetTintColorAtDistance(tintColorAtDistance float64) *PBRSubSurfaceConfiguration {
	p.p.Set("tintColorAtDistance", tintColorAtDistance)
	return p
}

// TranslucencyIntensity returns the TranslucencyIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#translucencyintensity
func (p *PBRSubSurfaceConfiguration) TranslucencyIntensity() float64 {
	retVal := p.p.Get("translucencyIntensity")
	return retVal.Float()
}

// SetTranslucencyIntensity sets the TranslucencyIntensity property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#translucencyintensity
func (p *PBRSubSurfaceConfiguration) SetTranslucencyIntensity(translucencyIntensity float64) *PBRSubSurfaceConfiguration {
	p.p.Set("translucencyIntensity", translucencyIntensity)
	return p
}

// UseMaskFromThicknessTexture returns the UseMaskFromThicknessTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#usemaskfromthicknesstexture
func (p *PBRSubSurfaceConfiguration) UseMaskFromThicknessTexture() bool {
	retVal := p.p.Get("useMaskFromThicknessTexture")
	return retVal.Bool()
}

// SetUseMaskFromThicknessTexture sets the UseMaskFromThicknessTexture property of class PBRSubSurfaceConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration#usemaskfromthicknesstexture
func (p *PBRSubSurfaceConfiguration) SetUseMaskFromThicknessTexture(useMaskFromThicknessTexture bool) *PBRSubSurfaceConfiguration {
	p.p.Set("useMaskFromThicknessTexture", useMaskFromThicknessTexture)
	return p
}
