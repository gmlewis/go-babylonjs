// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShadowGenerator represents a babylon.js ShadowGenerator.
// Default implementation IShadowGenerator.
// This is the main object responsible of generating shadows in the framework.
// Documentation: <a href="https://doc.babylonjs.com/babylon101/shadows">https://doc.babylonjs.com/babylon101/shadows</a>
type ShadowGenerator struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ShadowGenerator) JSObject() js.Value { return s.p }

// ShadowGenerator returns a ShadowGenerator JavaScript class.
func (ba *Babylon) ShadowGenerator() *ShadowGenerator {
	p := ba.ctx.Get("ShadowGenerator")
	return ShadowGeneratorFromJSObject(p, ba.ctx)
}

// ShadowGeneratorFromJSObject returns a wrapped ShadowGenerator JavaScript class.
func ShadowGeneratorFromJSObject(p js.Value, ctx js.Value) *ShadowGenerator {
	return &ShadowGenerator{p: p, ctx: ctx}
}

// ShadowGeneratorArrayToJSArray returns a JavaScript Array for the wrapped array.
func ShadowGeneratorArrayToJSArray(array []*ShadowGenerator) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewShadowGeneratorOpts contains optional parameters for NewShadowGenerator.
type NewShadowGeneratorOpts struct {
	UsefulFloatFirst *bool
}

// NewShadowGenerator returns a new ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#constructor
func (ba *Babylon) NewShadowGenerator(mapSize float64, light *IShadowLight, opts *NewShadowGeneratorOpts) *ShadowGenerator {
	if opts == nil {
		opts = &NewShadowGeneratorOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, mapSize)
	args = append(args, light.JSObject())

	if opts.UsefulFloatFirst == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UsefulFloatFirst)
	}

	p := ba.ctx.Get("ShadowGenerator").New(args...)
	return ShadowGeneratorFromJSObject(p, ba.ctx)
}

// ShadowGeneratorAddShadowCasterOpts contains optional parameters for ShadowGenerator.AddShadowCaster.
type ShadowGeneratorAddShadowCasterOpts struct {
	IncludeDescendants *bool
}

// AddShadowCaster calls the AddShadowCaster method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#addshadowcaster
func (s *ShadowGenerator) AddShadowCaster(mesh *AbstractMesh, opts *ShadowGeneratorAddShadowCasterOpts) *ShadowGenerator {
	if opts == nil {
		opts = &ShadowGeneratorAddShadowCasterOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.IncludeDescendants == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IncludeDescendants)
	}

	retVal := s.p.Call("addShadowCaster", args...)
	return ShadowGeneratorFromJSObject(retVal, s.ctx)
}

// BindShadowLight calls the BindShadowLight method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#bindshadowlight
func (s *ShadowGenerator) BindShadowLight(lightIndex string, effect *Effect) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, lightIndex)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	s.p.Call("bindShadowLight", args...)
}

// Dispose calls the Dispose method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#dispose
func (s *ShadowGenerator) Dispose() {

	s.p.Call("dispose")
}

// ShadowGeneratorForceCompilationOpts contains optional parameters for ShadowGenerator.ForceCompilation.
type ShadowGeneratorForceCompilationOpts struct {
	OnCompiled JSFunc
	Options    map[string]interface{}
}

// ForceCompilation calls the ForceCompilation method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#forcecompilation
func (s *ShadowGenerator) ForceCompilation(opts *ShadowGeneratorForceCompilationOpts) {
	if opts == nil {
		opts = &ShadowGeneratorForceCompilationOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnCompiled) /* never freed! */)
	}
	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	s.p.Call("forceCompilation", args...)
}

// ShadowGeneratorForceCompilationAsyncOpts contains optional parameters for ShadowGenerator.ForceCompilationAsync.
type ShadowGeneratorForceCompilationAsyncOpts struct {
	Options map[string]interface{}
}

// ForceCompilationAsync calls the ForceCompilationAsync method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#forcecompilationasync
func (s *ShadowGenerator) ForceCompilationAsync(opts *ShadowGeneratorForceCompilationAsyncOpts) *Promise {
	if opts == nil {
		opts = &ShadowGeneratorForceCompilationAsyncOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("forceCompilationAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// GetClassName calls the GetClassName method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#getclassname
func (s *ShadowGenerator) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// GetDarkness calls the GetDarkness method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#getdarkness
func (s *ShadowGenerator) GetDarkness() float64 {

	retVal := s.p.Call("getDarkness")
	return retVal.Float()
}

// GetLight calls the GetLight method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#getlight
func (s *ShadowGenerator) GetLight() *IShadowLight {

	retVal := s.p.Call("getLight")
	return IShadowLightFromJSObject(retVal, s.ctx)
}

// GetShadowMap calls the GetShadowMap method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#getshadowmap
func (s *ShadowGenerator) GetShadowMap() *RenderTargetTexture {

	retVal := s.p.Call("getShadowMap")
	return RenderTargetTextureFromJSObject(retVal, s.ctx)
}

// GetShadowMapForRendering calls the GetShadowMapForRendering method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#getshadowmapforrendering
func (s *ShadowGenerator) GetShadowMapForRendering() *RenderTargetTexture {

	retVal := s.p.Call("getShadowMapForRendering")
	return RenderTargetTextureFromJSObject(retVal, s.ctx)
}

// GetTransformMatrix calls the GetTransformMatrix method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#gettransformmatrix
func (s *ShadowGenerator) GetTransformMatrix() *Matrix {

	retVal := s.p.Call("getTransformMatrix")
	return MatrixFromJSObject(retVal, s.ctx)
}

// IsReady calls the IsReady method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#isready
func (s *ShadowGenerator) IsReady(subMesh *SubMesh, useInstances bool) bool {

	args := make([]interface{}, 0, 2+0)

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	args = append(args, useInstances)

	retVal := s.p.Call("isReady", args...)
	return retVal.Bool()
}

// Parse calls the Parse method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#parse
func (s *ShadowGenerator) Parse(parsedShadowGenerator JSObject, scene *Scene) *ShadowGenerator {

	args := make([]interface{}, 0, 2+0)

	if parsedShadowGenerator == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedShadowGenerator.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := s.p.Call("Parse", args...)
	return ShadowGeneratorFromJSObject(retVal, s.ctx)
}

// PrepareDefines calls the PrepareDefines method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#preparedefines
func (s *ShadowGenerator) PrepareDefines(defines JSObject, lightIndex float64) {

	args := make([]interface{}, 0, 2+0)

	if defines == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, defines.JSObject())
	}

	args = append(args, lightIndex)

	s.p.Call("prepareDefines", args...)
}

// RecreateShadowMap calls the RecreateShadowMap method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#recreateshadowmap
func (s *ShadowGenerator) RecreateShadowMap() {

	s.p.Call("recreateShadowMap")
}

// ShadowGeneratorRemoveShadowCasterOpts contains optional parameters for ShadowGenerator.RemoveShadowCaster.
type ShadowGeneratorRemoveShadowCasterOpts struct {
	IncludeDescendants *bool
}

// RemoveShadowCaster calls the RemoveShadowCaster method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#removeshadowcaster
func (s *ShadowGenerator) RemoveShadowCaster(mesh *AbstractMesh, opts *ShadowGeneratorRemoveShadowCasterOpts) *ShadowGenerator {
	if opts == nil {
		opts = &ShadowGeneratorRemoveShadowCasterOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.IncludeDescendants == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IncludeDescendants)
	}

	retVal := s.p.Call("removeShadowCaster", args...)
	return ShadowGeneratorFromJSObject(retVal, s.ctx)
}

// Serialize calls the Serialize method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#serialize
func (s *ShadowGenerator) Serialize() js.Value {

	retVal := s.p.Call("serialize")
	return retVal
}

// SetDarkness calls the SetDarkness method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#setdarkness
func (s *ShadowGenerator) SetDarkness(darkness float64) *ShadowGenerator {

	args := make([]interface{}, 0, 1+0)

	args = append(args, darkness)

	retVal := s.p.Call("setDarkness", args...)
	return ShadowGeneratorFromJSObject(retVal, s.ctx)
}

// SetTransparencyShadow calls the SetTransparencyShadow method on the ShadowGenerator object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#settransparencyshadow
func (s *ShadowGenerator) SetTransparencyShadow(transparent bool) *ShadowGenerator {

	args := make([]interface{}, 0, 1+0)

	args = append(args, transparent)

	retVal := s.p.Call("setTransparencyShadow", args...)
	return ShadowGeneratorFromJSObject(retVal, s.ctx)
}

// Bias returns the Bias property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#bias
func (s *ShadowGenerator) Bias() float64 {
	retVal := s.p.Get("bias")
	return retVal.Float()
}

// SetBias sets the Bias property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#bias
func (s *ShadowGenerator) SetBias(bias float64) *ShadowGenerator {
	s.p.Set("bias", bias)
	return s
}

// BlurBoxOffset returns the BlurBoxOffset property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurboxoffset
func (s *ShadowGenerator) BlurBoxOffset() float64 {
	retVal := s.p.Get("blurBoxOffset")
	return retVal.Float()
}

// SetBlurBoxOffset sets the BlurBoxOffset property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurboxoffset
func (s *ShadowGenerator) SetBlurBoxOffset(blurBoxOffset float64) *ShadowGenerator {
	s.p.Set("blurBoxOffset", blurBoxOffset)
	return s
}

// BlurKernel returns the BlurKernel property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurkernel
func (s *ShadowGenerator) BlurKernel() float64 {
	retVal := s.p.Get("blurKernel")
	return retVal.Float()
}

// SetBlurKernel sets the BlurKernel property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurkernel
func (s *ShadowGenerator) SetBlurKernel(blurKernel float64) *ShadowGenerator {
	s.p.Set("blurKernel", blurKernel)
	return s
}

// BlurScale returns the BlurScale property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurscale
func (s *ShadowGenerator) BlurScale() float64 {
	retVal := s.p.Get("blurScale")
	return retVal.Float()
}

// SetBlurScale sets the BlurScale property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#blurscale
func (s *ShadowGenerator) SetBlurScale(blurScale float64) *ShadowGenerator {
	s.p.Set("blurScale", blurScale)
	return s
}

// ContactHardeningLightSizeUVRatio returns the ContactHardeningLightSizeUVRatio property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#contacthardeninglightsizeuvratio
func (s *ShadowGenerator) ContactHardeningLightSizeUVRatio() float64 {
	retVal := s.p.Get("contactHardeningLightSizeUVRatio")
	return retVal.Float()
}

// SetContactHardeningLightSizeUVRatio sets the ContactHardeningLightSizeUVRatio property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#contacthardeninglightsizeuvratio
func (s *ShadowGenerator) SetContactHardeningLightSizeUVRatio(contactHardeningLightSizeUVRatio float64) *ShadowGenerator {
	s.p.Set("contactHardeningLightSizeUVRatio", contactHardeningLightSizeUVRatio)
	return s
}

// CustomShaderOptions returns the CustomShaderOptions property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#customshaderoptions
func (s *ShadowGenerator) CustomShaderOptions() *ICustomShaderOptions {
	retVal := s.p.Get("customShaderOptions")
	return ICustomShaderOptionsFromJSObject(retVal, s.ctx)
}

// SetCustomShaderOptions sets the CustomShaderOptions property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#customshaderoptions
func (s *ShadowGenerator) SetCustomShaderOptions(customShaderOptions *ICustomShaderOptions) *ShadowGenerator {
	s.p.Set("customShaderOptions", customShaderOptions.JSObject())
	return s
}

// Darkness returns the Darkness property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#darkness
func (s *ShadowGenerator) Darkness() float64 {
	retVal := s.p.Get("darkness")
	return retVal.Float()
}

// DepthScale returns the DepthScale property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#depthscale
func (s *ShadowGenerator) DepthScale() float64 {
	retVal := s.p.Get("depthScale")
	return retVal.Float()
}

// SetDepthScale sets the DepthScale property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#depthscale
func (s *ShadowGenerator) SetDepthScale(depthScale float64) *ShadowGenerator {
	s.p.Set("depthScale", depthScale)
	return s
}

// FILTER_BLURCLOSEEXPONENTIALSHADOWMAP returns the FILTER_BLURCLOSEEXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_blurcloseexponentialshadowmap
func (s *ShadowGenerator) FILTER_BLURCLOSEEXPONENTIALSHADOWMAP() float64 {
	retVal := s.p.Get("FILTER_BLURCLOSEEXPONENTIALSHADOWMAP")
	return retVal.Float()
}

// SetFILTER_BLURCLOSEEXPONENTIALSHADOWMAP sets the FILTER_BLURCLOSEEXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_blurcloseexponentialshadowmap
func (s *ShadowGenerator) SetFILTER_BLURCLOSEEXPONENTIALSHADOWMAP(FILTER_BLURCLOSEEXPONENTIALSHADOWMAP float64) *ShadowGenerator {
	s.p.Set("FILTER_BLURCLOSEEXPONENTIALSHADOWMAP", FILTER_BLURCLOSEEXPONENTIALSHADOWMAP)
	return s
}

// FILTER_BLUREXPONENTIALSHADOWMAP returns the FILTER_BLUREXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_blurexponentialshadowmap
func (s *ShadowGenerator) FILTER_BLUREXPONENTIALSHADOWMAP() float64 {
	retVal := s.p.Get("FILTER_BLUREXPONENTIALSHADOWMAP")
	return retVal.Float()
}

// SetFILTER_BLUREXPONENTIALSHADOWMAP sets the FILTER_BLUREXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_blurexponentialshadowmap
func (s *ShadowGenerator) SetFILTER_BLUREXPONENTIALSHADOWMAP(FILTER_BLUREXPONENTIALSHADOWMAP float64) *ShadowGenerator {
	s.p.Set("FILTER_BLUREXPONENTIALSHADOWMAP", FILTER_BLUREXPONENTIALSHADOWMAP)
	return s
}

// FILTER_CLOSEEXPONENTIALSHADOWMAP returns the FILTER_CLOSEEXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_closeexponentialshadowmap
func (s *ShadowGenerator) FILTER_CLOSEEXPONENTIALSHADOWMAP() float64 {
	retVal := s.p.Get("FILTER_CLOSEEXPONENTIALSHADOWMAP")
	return retVal.Float()
}

// SetFILTER_CLOSEEXPONENTIALSHADOWMAP sets the FILTER_CLOSEEXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_closeexponentialshadowmap
func (s *ShadowGenerator) SetFILTER_CLOSEEXPONENTIALSHADOWMAP(FILTER_CLOSEEXPONENTIALSHADOWMAP float64) *ShadowGenerator {
	s.p.Set("FILTER_CLOSEEXPONENTIALSHADOWMAP", FILTER_CLOSEEXPONENTIALSHADOWMAP)
	return s
}

// FILTER_EXPONENTIALSHADOWMAP returns the FILTER_EXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_exponentialshadowmap
func (s *ShadowGenerator) FILTER_EXPONENTIALSHADOWMAP() float64 {
	retVal := s.p.Get("FILTER_EXPONENTIALSHADOWMAP")
	return retVal.Float()
}

// SetFILTER_EXPONENTIALSHADOWMAP sets the FILTER_EXPONENTIALSHADOWMAP property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_exponentialshadowmap
func (s *ShadowGenerator) SetFILTER_EXPONENTIALSHADOWMAP(FILTER_EXPONENTIALSHADOWMAP float64) *ShadowGenerator {
	s.p.Set("FILTER_EXPONENTIALSHADOWMAP", FILTER_EXPONENTIALSHADOWMAP)
	return s
}

// FILTER_NONE returns the FILTER_NONE property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_none
func (s *ShadowGenerator) FILTER_NONE() float64 {
	retVal := s.p.Get("FILTER_NONE")
	return retVal.Float()
}

// SetFILTER_NONE sets the FILTER_NONE property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_none
func (s *ShadowGenerator) SetFILTER_NONE(FILTER_NONE float64) *ShadowGenerator {
	s.p.Set("FILTER_NONE", FILTER_NONE)
	return s
}

// FILTER_PCF returns the FILTER_PCF property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_pcf
func (s *ShadowGenerator) FILTER_PCF() float64 {
	retVal := s.p.Get("FILTER_PCF")
	return retVal.Float()
}

// SetFILTER_PCF sets the FILTER_PCF property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_pcf
func (s *ShadowGenerator) SetFILTER_PCF(FILTER_PCF float64) *ShadowGenerator {
	s.p.Set("FILTER_PCF", FILTER_PCF)
	return s
}

// FILTER_PCSS returns the FILTER_PCSS property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_pcss
func (s *ShadowGenerator) FILTER_PCSS() float64 {
	retVal := s.p.Get("FILTER_PCSS")
	return retVal.Float()
}

// SetFILTER_PCSS sets the FILTER_PCSS property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_pcss
func (s *ShadowGenerator) SetFILTER_PCSS(FILTER_PCSS float64) *ShadowGenerator {
	s.p.Set("FILTER_PCSS", FILTER_PCSS)
	return s
}

// FILTER_POISSONSAMPLING returns the FILTER_POISSONSAMPLING property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_poissonsampling
func (s *ShadowGenerator) FILTER_POISSONSAMPLING() float64 {
	retVal := s.p.Get("FILTER_POISSONSAMPLING")
	return retVal.Float()
}

// SetFILTER_POISSONSAMPLING sets the FILTER_POISSONSAMPLING property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter_poissonsampling
func (s *ShadowGenerator) SetFILTER_POISSONSAMPLING(FILTER_POISSONSAMPLING float64) *ShadowGenerator {
	s.p.Set("FILTER_POISSONSAMPLING", FILTER_POISSONSAMPLING)
	return s
}

// Filter returns the Filter property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter
func (s *ShadowGenerator) Filter() float64 {
	retVal := s.p.Get("filter")
	return retVal.Float()
}

// SetFilter sets the Filter property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filter
func (s *ShadowGenerator) SetFilter(filter float64) *ShadowGenerator {
	s.p.Set("filter", filter)
	return s
}

// FilteringQuality returns the FilteringQuality property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filteringquality
func (s *ShadowGenerator) FilteringQuality() float64 {
	retVal := s.p.Get("filteringQuality")
	return retVal.Float()
}

// SetFilteringQuality sets the FilteringQuality property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#filteringquality
func (s *ShadowGenerator) SetFilteringQuality(filteringQuality float64) *ShadowGenerator {
	s.p.Set("filteringQuality", filteringQuality)
	return s
}

// ForceBackFacesOnly returns the ForceBackFacesOnly property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#forcebackfacesonly
func (s *ShadowGenerator) ForceBackFacesOnly() bool {
	retVal := s.p.Get("forceBackFacesOnly")
	return retVal.Bool()
}

// SetForceBackFacesOnly sets the ForceBackFacesOnly property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#forcebackfacesonly
func (s *ShadowGenerator) SetForceBackFacesOnly(forceBackFacesOnly bool) *ShadowGenerator {
	s.p.Set("forceBackFacesOnly", forceBackFacesOnly)
	return s
}

// FrustumEdgeFalloff returns the FrustumEdgeFalloff property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#frustumedgefalloff
func (s *ShadowGenerator) FrustumEdgeFalloff() float64 {
	retVal := s.p.Get("frustumEdgeFalloff")
	return retVal.Float()
}

// SetFrustumEdgeFalloff sets the FrustumEdgeFalloff property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#frustumedgefalloff
func (s *ShadowGenerator) SetFrustumEdgeFalloff(frustumEdgeFalloff float64) *ShadowGenerator {
	s.p.Set("frustumEdgeFalloff", frustumEdgeFalloff)
	return s
}

// NormalBias returns the NormalBias property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#normalbias
func (s *ShadowGenerator) NormalBias() float64 {
	retVal := s.p.Get("normalBias")
	return retVal.Float()
}

// SetNormalBias sets the NormalBias property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#normalbias
func (s *ShadowGenerator) SetNormalBias(normalBias float64) *ShadowGenerator {
	s.p.Set("normalBias", normalBias)
	return s
}

// OnAfterShadowMapRenderMeshObservable returns the OnAfterShadowMapRenderMeshObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onaftershadowmaprendermeshobservable
func (s *ShadowGenerator) OnAfterShadowMapRenderMeshObservable() *Observable {
	retVal := s.p.Get("onAfterShadowMapRenderMeshObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnAfterShadowMapRenderMeshObservable sets the OnAfterShadowMapRenderMeshObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onaftershadowmaprendermeshobservable
func (s *ShadowGenerator) SetOnAfterShadowMapRenderMeshObservable(onAfterShadowMapRenderMeshObservable *Observable) *ShadowGenerator {
	s.p.Set("onAfterShadowMapRenderMeshObservable", onAfterShadowMapRenderMeshObservable.JSObject())
	return s
}

// OnAfterShadowMapRenderObservable returns the OnAfterShadowMapRenderObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onaftershadowmaprenderobservable
func (s *ShadowGenerator) OnAfterShadowMapRenderObservable() *Observable {
	retVal := s.p.Get("onAfterShadowMapRenderObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnAfterShadowMapRenderObservable sets the OnAfterShadowMapRenderObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onaftershadowmaprenderobservable
func (s *ShadowGenerator) SetOnAfterShadowMapRenderObservable(onAfterShadowMapRenderObservable *Observable) *ShadowGenerator {
	s.p.Set("onAfterShadowMapRenderObservable", onAfterShadowMapRenderObservable.JSObject())
	return s
}

// OnBeforeShadowMapRenderMeshObservable returns the OnBeforeShadowMapRenderMeshObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onbeforeshadowmaprendermeshobservable
func (s *ShadowGenerator) OnBeforeShadowMapRenderMeshObservable() *Observable {
	retVal := s.p.Get("onBeforeShadowMapRenderMeshObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnBeforeShadowMapRenderMeshObservable sets the OnBeforeShadowMapRenderMeshObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onbeforeshadowmaprendermeshobservable
func (s *ShadowGenerator) SetOnBeforeShadowMapRenderMeshObservable(onBeforeShadowMapRenderMeshObservable *Observable) *ShadowGenerator {
	s.p.Set("onBeforeShadowMapRenderMeshObservable", onBeforeShadowMapRenderMeshObservable.JSObject())
	return s
}

// OnBeforeShadowMapRenderObservable returns the OnBeforeShadowMapRenderObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onbeforeshadowmaprenderobservable
func (s *ShadowGenerator) OnBeforeShadowMapRenderObservable() *Observable {
	retVal := s.p.Get("onBeforeShadowMapRenderObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnBeforeShadowMapRenderObservable sets the OnBeforeShadowMapRenderObservable property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#onbeforeshadowmaprenderobservable
func (s *ShadowGenerator) SetOnBeforeShadowMapRenderObservable(onBeforeShadowMapRenderObservable *Observable) *ShadowGenerator {
	s.p.Set("onBeforeShadowMapRenderObservable", onBeforeShadowMapRenderObservable.JSObject())
	return s
}

// QUALITY_HIGH returns the QUALITY_HIGH property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_high
func (s *ShadowGenerator) QUALITY_HIGH() float64 {
	retVal := s.p.Get("QUALITY_HIGH")
	return retVal.Float()
}

// SetQUALITY_HIGH sets the QUALITY_HIGH property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_high
func (s *ShadowGenerator) SetQUALITY_HIGH(QUALITY_HIGH float64) *ShadowGenerator {
	s.p.Set("QUALITY_HIGH", QUALITY_HIGH)
	return s
}

// QUALITY_LOW returns the QUALITY_LOW property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_low
func (s *ShadowGenerator) QUALITY_LOW() float64 {
	retVal := s.p.Get("QUALITY_LOW")
	return retVal.Float()
}

// SetQUALITY_LOW sets the QUALITY_LOW property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_low
func (s *ShadowGenerator) SetQUALITY_LOW(QUALITY_LOW float64) *ShadowGenerator {
	s.p.Set("QUALITY_LOW", QUALITY_LOW)
	return s
}

// QUALITY_MEDIUM returns the QUALITY_MEDIUM property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_medium
func (s *ShadowGenerator) QUALITY_MEDIUM() float64 {
	retVal := s.p.Get("QUALITY_MEDIUM")
	return retVal.Float()
}

// SetQUALITY_MEDIUM sets the QUALITY_MEDIUM property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#quality_medium
func (s *ShadowGenerator) SetQUALITY_MEDIUM(QUALITY_MEDIUM float64) *ShadowGenerator {
	s.p.Set("QUALITY_MEDIUM", QUALITY_MEDIUM)
	return s
}

// TransparencyShadow returns the TransparencyShadow property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#transparencyshadow
func (s *ShadowGenerator) TransparencyShadow() bool {
	retVal := s.p.Get("transparencyShadow")
	return retVal.Bool()
}

// UseBlurCloseExponentialShadowMap returns the UseBlurCloseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useblurcloseexponentialshadowmap
func (s *ShadowGenerator) UseBlurCloseExponentialShadowMap() bool {
	retVal := s.p.Get("useBlurCloseExponentialShadowMap")
	return retVal.Bool()
}

// SetUseBlurCloseExponentialShadowMap sets the UseBlurCloseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useblurcloseexponentialshadowmap
func (s *ShadowGenerator) SetUseBlurCloseExponentialShadowMap(useBlurCloseExponentialShadowMap bool) *ShadowGenerator {
	s.p.Set("useBlurCloseExponentialShadowMap", useBlurCloseExponentialShadowMap)
	return s
}

// UseBlurExponentialShadowMap returns the UseBlurExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useblurexponentialshadowmap
func (s *ShadowGenerator) UseBlurExponentialShadowMap() bool {
	retVal := s.p.Get("useBlurExponentialShadowMap")
	return retVal.Bool()
}

// SetUseBlurExponentialShadowMap sets the UseBlurExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useblurexponentialshadowmap
func (s *ShadowGenerator) SetUseBlurExponentialShadowMap(useBlurExponentialShadowMap bool) *ShadowGenerator {
	s.p.Set("useBlurExponentialShadowMap", useBlurExponentialShadowMap)
	return s
}

// UseCloseExponentialShadowMap returns the UseCloseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usecloseexponentialshadowmap
func (s *ShadowGenerator) UseCloseExponentialShadowMap() bool {
	retVal := s.p.Get("useCloseExponentialShadowMap")
	return retVal.Bool()
}

// SetUseCloseExponentialShadowMap sets the UseCloseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usecloseexponentialshadowmap
func (s *ShadowGenerator) SetUseCloseExponentialShadowMap(useCloseExponentialShadowMap bool) *ShadowGenerator {
	s.p.Set("useCloseExponentialShadowMap", useCloseExponentialShadowMap)
	return s
}

// UseContactHardeningShadow returns the UseContactHardeningShadow property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usecontacthardeningshadow
func (s *ShadowGenerator) UseContactHardeningShadow() bool {
	retVal := s.p.Get("useContactHardeningShadow")
	return retVal.Bool()
}

// SetUseContactHardeningShadow sets the UseContactHardeningShadow property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usecontacthardeningshadow
func (s *ShadowGenerator) SetUseContactHardeningShadow(useContactHardeningShadow bool) *ShadowGenerator {
	s.p.Set("useContactHardeningShadow", useContactHardeningShadow)
	return s
}

// UseExponentialShadowMap returns the UseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useexponentialshadowmap
func (s *ShadowGenerator) UseExponentialShadowMap() bool {
	retVal := s.p.Get("useExponentialShadowMap")
	return retVal.Bool()
}

// SetUseExponentialShadowMap sets the UseExponentialShadowMap property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#useexponentialshadowmap
func (s *ShadowGenerator) SetUseExponentialShadowMap(useExponentialShadowMap bool) *ShadowGenerator {
	s.p.Set("useExponentialShadowMap", useExponentialShadowMap)
	return s
}

// UseKernelBlur returns the UseKernelBlur property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usekernelblur
func (s *ShadowGenerator) UseKernelBlur() bool {
	retVal := s.p.Get("useKernelBlur")
	return retVal.Bool()
}

// SetUseKernelBlur sets the UseKernelBlur property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usekernelblur
func (s *ShadowGenerator) SetUseKernelBlur(useKernelBlur bool) *ShadowGenerator {
	s.p.Set("useKernelBlur", useKernelBlur)
	return s
}

// UsePercentageCloserFiltering returns the UsePercentageCloserFiltering property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usepercentagecloserfiltering
func (s *ShadowGenerator) UsePercentageCloserFiltering() bool {
	retVal := s.p.Get("usePercentageCloserFiltering")
	return retVal.Bool()
}

// SetUsePercentageCloserFiltering sets the UsePercentageCloserFiltering property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usepercentagecloserfiltering
func (s *ShadowGenerator) SetUsePercentageCloserFiltering(usePercentageCloserFiltering bool) *ShadowGenerator {
	s.p.Set("usePercentageCloserFiltering", usePercentageCloserFiltering)
	return s
}

// UsePoissonSampling returns the UsePoissonSampling property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usepoissonsampling
func (s *ShadowGenerator) UsePoissonSampling() bool {
	retVal := s.p.Get("usePoissonSampling")
	return retVal.Bool()
}

// SetUsePoissonSampling sets the UsePoissonSampling property of class ShadowGenerator.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgenerator#usepoissonsampling
func (s *ShadowGenerator) SetUsePoissonSampling(usePoissonSampling bool) *ShadowGenerator {
	s.p.Set("usePoissonSampling", usePoissonSampling)
	return s
}
