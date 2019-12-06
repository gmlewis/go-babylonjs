// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MaterialHelper represents a babylon.js MaterialHelper.
// &quot;Static Class&quot; containing the most commonly used helper while dealing with material for
// rendering purpose.
//
// It contains the basic tools to help defining defines, binding uniform for the common part of the materials.
type MaterialHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MaterialHelper) JSObject() js.Value { return m.p }

// MaterialHelper returns a MaterialHelper JavaScript class.
func (ba *Babylon) MaterialHelper() *MaterialHelper {
	p := ba.ctx.Get("MaterialHelper")
	return MaterialHelperFromJSObject(p, ba.ctx)
}

// MaterialHelperFromJSObject returns a wrapped MaterialHelper JavaScript class.
func MaterialHelperFromJSObject(p js.Value, ctx js.Value) *MaterialHelper {
	return &MaterialHelper{p: p, ctx: ctx}
}

// MaterialHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func MaterialHelperArrayToJSArray(array []*MaterialHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// MaterialHelperBindBonesParametersOpts contains optional parameters for MaterialHelper.BindBonesParameters.
type MaterialHelperBindBonesParametersOpts struct {
	Mesh   *AbstractMesh
	Effect *Effect
}

// BindBonesParameters calls the BindBonesParameters method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindbonesparameters
func (m *MaterialHelper) BindBonesParameters(opts *MaterialHelperBindBonesParametersOpts) {
	if opts == nil {
		opts = &MaterialHelperBindBonesParametersOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}
	if opts.Effect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Effect.JSObject())
	}

	m.p.Call("BindBonesParameters", args...)
}

// BindClipPlane calls the BindClipPlane method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindclipplane
func (m *MaterialHelper) BindClipPlane(effect *Effect, scene *Scene) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, scene.JSObject())

	m.p.Call("BindClipPlane", args...)
}

// BindEyePosition calls the BindEyePosition method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindeyeposition
func (m *MaterialHelper) BindEyePosition(effect *Effect, scene *Scene) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, scene.JSObject())

	m.p.Call("BindEyePosition", args...)
}

// MaterialHelperBindFogParametersOpts contains optional parameters for MaterialHelper.BindFogParameters.
type MaterialHelperBindFogParametersOpts struct {
	LinearSpace *bool
}

// BindFogParameters calls the BindFogParameters method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindfogparameters
func (m *MaterialHelper) BindFogParameters(scene *Scene, mesh *AbstractMesh, effect *Effect, opts *MaterialHelperBindFogParametersOpts) {
	if opts == nil {
		opts = &MaterialHelperBindFogParametersOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, scene.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, effect.JSObject())

	if opts.LinearSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LinearSpace)
	}

	m.p.Call("BindFogParameters", args...)
}

// MaterialHelperBindLightOpts contains optional parameters for MaterialHelper.BindLight.
type MaterialHelperBindLightOpts struct {
	UsePhysicalLightFalloff *bool
	RebuildInParallel       *bool
}

// BindLight calls the BindLight method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindlight
func (m *MaterialHelper) BindLight(light *Light, lightIndex float64, scene *Scene, effect *Effect, useSpecular bool, opts *MaterialHelperBindLightOpts) {
	if opts == nil {
		opts = &MaterialHelperBindLightOpts{}
	}

	args := make([]interface{}, 0, 5+2)

	args = append(args, light.JSObject())
	args = append(args, lightIndex)
	args = append(args, scene.JSObject())
	args = append(args, effect.JSObject())
	args = append(args, useSpecular)

	if opts.UsePhysicalLightFalloff == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UsePhysicalLightFalloff)
	}
	if opts.RebuildInParallel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RebuildInParallel)
	}

	m.p.Call("BindLight", args...)
}

// BindLightProperties calls the BindLightProperties method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindlightproperties
func (m *MaterialHelper) BindLightProperties(light *Light, effect *Effect, lightIndex float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, light.JSObject())
	args = append(args, effect.JSObject())
	args = append(args, lightIndex)

	m.p.Call("BindLightProperties", args...)
}

// MaterialHelperBindLightsOpts contains optional parameters for MaterialHelper.BindLights.
type MaterialHelperBindLightsOpts struct {
	MaxSimultaneousLights   *float64
	UsePhysicalLightFalloff *bool
	RebuildInParallel       *bool
}

// BindLights calls the BindLights method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindlights
func (m *MaterialHelper) BindLights(scene *Scene, mesh *AbstractMesh, effect *Effect, defines interface{}, opts *MaterialHelperBindLightsOpts) {
	if opts == nil {
		opts = &MaterialHelperBindLightsOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, scene.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, effect.JSObject())
	args = append(args, defines)

	if opts.MaxSimultaneousLights == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxSimultaneousLights)
	}
	if opts.UsePhysicalLightFalloff == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UsePhysicalLightFalloff)
	}
	if opts.RebuildInParallel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RebuildInParallel)
	}

	m.p.Call("BindLights", args...)
}

// BindLogDepth calls the BindLogDepth method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindlogdepth
func (m *MaterialHelper) BindLogDepth(defines interface{}, effect *Effect, scene *Scene) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, defines)
	args = append(args, effect.JSObject())
	args = append(args, scene.JSObject())

	m.p.Call("BindLogDepth", args...)
}

// BindMorphTargetParameters calls the BindMorphTargetParameters method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindmorphtargetparameters
func (m *MaterialHelper) BindMorphTargetParameters(abstractMesh *AbstractMesh, effect *Effect) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, abstractMesh.JSObject())
	args = append(args, effect.JSObject())

	m.p.Call("BindMorphTargetParameters", args...)
}

// BindTextureMatrix calls the BindTextureMatrix method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#bindtexturematrix
func (m *MaterialHelper) BindTextureMatrix(texture *BaseTexture, uniformBuffer *UniformBuffer, key string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, texture.JSObject())
	args = append(args, uniformBuffer.JSObject())
	args = append(args, key)

	m.p.Call("BindTextureMatrix", args...)
}

// GetFogState calls the GetFogState method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#getfogstate
func (m *MaterialHelper) GetFogState(mesh *AbstractMesh, scene *Scene) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, scene.JSObject())

	retVal := m.p.Call("GetFogState", args...)
	return retVal.Bool()
}

// MaterialHelperHandleFallbacksForShadowsOpts contains optional parameters for MaterialHelper.HandleFallbacksForShadows.
type MaterialHelperHandleFallbacksForShadowsOpts struct {
	MaxSimultaneousLights *float64
	Rank                  *float64
}

// HandleFallbacksForShadows calls the HandleFallbacksForShadows method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#handlefallbacksforshadows
func (m *MaterialHelper) HandleFallbacksForShadows(defines interface{}, fallbacks *EffectFallbacks, opts *MaterialHelperHandleFallbacksForShadowsOpts) float64 {
	if opts == nil {
		opts = &MaterialHelperHandleFallbacksForShadowsOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, defines)
	args = append(args, fallbacks.JSObject())

	if opts.MaxSimultaneousLights == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxSimultaneousLights)
	}
	if opts.Rank == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Rank)
	}

	retVal := m.p.Call("HandleFallbacksForShadows", args...)
	return retVal.Float()
}

// PrepareAttributesForBones calls the PrepareAttributesForBones method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareattributesforbones
func (m *MaterialHelper) PrepareAttributesForBones(attribs []string, mesh *AbstractMesh, defines interface{}, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, attribs)
	args = append(args, mesh.JSObject())
	args = append(args, defines)
	args = append(args, fallbacks.JSObject())

	m.p.Call("PrepareAttributesForBones", args...)
}

// PrepareAttributesForInstances calls the PrepareAttributesForInstances method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareattributesforinstances
func (m *MaterialHelper) PrepareAttributesForInstances(attribs []string, defines *MaterialDefines) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, attribs)
	args = append(args, defines.JSObject())

	m.p.Call("PrepareAttributesForInstances", args...)
}

// PrepareAttributesForMorphTargets calls the PrepareAttributesForMorphTargets method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareattributesformorphtargets
func (m *MaterialHelper) PrepareAttributesForMorphTargets(attribs []string, mesh *AbstractMesh, defines interface{}) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, attribs)
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	m.p.Call("PrepareAttributesForMorphTargets", args...)
}

// PrepareAttributesForMorphTargetsInfluencers calls the PrepareAttributesForMorphTargetsInfluencers method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareattributesformorphtargetsinfluencers
func (m *MaterialHelper) PrepareAttributesForMorphTargetsInfluencers(attribs []string, mesh *AbstractMesh, influencers float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, attribs)
	args = append(args, mesh.JSObject())
	args = append(args, influencers)

	m.p.Call("PrepareAttributesForMorphTargetsInfluencers", args...)
}

// MaterialHelperPrepareDefinesForAttributesOpts contains optional parameters for MaterialHelper.PrepareDefinesForAttributes.
type MaterialHelperPrepareDefinesForAttributesOpts struct {
	UseMorphTargets *bool
	UseVertexAlpha  *bool
}

// PrepareDefinesForAttributes calls the PrepareDefinesForAttributes method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesforattributes
func (m *MaterialHelper) PrepareDefinesForAttributes(mesh *AbstractMesh, defines interface{}, useVertexColor bool, useBones bool, opts *MaterialHelperPrepareDefinesForAttributesOpts) bool {
	if opts == nil {
		opts = &MaterialHelperPrepareDefinesForAttributesOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, mesh.JSObject())
	args = append(args, defines)
	args = append(args, useVertexColor)
	args = append(args, useBones)

	if opts.UseMorphTargets == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseMorphTargets)
	}
	if opts.UseVertexAlpha == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseVertexAlpha)
	}

	retVal := m.p.Call("PrepareDefinesForAttributes", args...)
	return retVal.Bool()
}

// PrepareDefinesForBones calls the PrepareDefinesForBones method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesforbones
func (m *MaterialHelper) PrepareDefinesForBones(mesh *AbstractMesh, defines interface{}) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, defines)

	m.p.Call("PrepareDefinesForBones", args...)
}

// MaterialHelperPrepareDefinesForFrameBoundValuesOpts contains optional parameters for MaterialHelper.PrepareDefinesForFrameBoundValues.
type MaterialHelperPrepareDefinesForFrameBoundValuesOpts struct {
	UseClipPlane *bool
}

// PrepareDefinesForFrameBoundValues calls the PrepareDefinesForFrameBoundValues method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesforframeboundvalues
func (m *MaterialHelper) PrepareDefinesForFrameBoundValues(scene *Scene, engine *Engine, defines interface{}, useInstances bool, opts *MaterialHelperPrepareDefinesForFrameBoundValuesOpts) {
	if opts == nil {
		opts = &MaterialHelperPrepareDefinesForFrameBoundValuesOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, scene.JSObject())
	args = append(args, engine.JSObject())
	args = append(args, defines)
	args = append(args, useInstances)

	if opts.UseClipPlane == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseClipPlane)
	}

	m.p.Call("PrepareDefinesForFrameBoundValues", args...)
}

// PrepareDefinesForLight calls the PrepareDefinesForLight method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesforlight
func (m *MaterialHelper) PrepareDefinesForLight(scene *Scene, mesh *AbstractMesh, light *Light, lightIndex float64, defines interface{}, specularSupported bool, state js.Value) {

	args := make([]interface{}, 0, 7+0)

	args = append(args, scene.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, light.JSObject())
	args = append(args, lightIndex)
	args = append(args, defines)
	args = append(args, specularSupported)
	args = append(args, state)

	m.p.Call("PrepareDefinesForLight", args...)
}

// MaterialHelperPrepareDefinesForLightsOpts contains optional parameters for MaterialHelper.PrepareDefinesForLights.
type MaterialHelperPrepareDefinesForLightsOpts struct {
	MaxSimultaneousLights *float64
	DisableLighting       *bool
}

// PrepareDefinesForLights calls the PrepareDefinesForLights method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesforlights
func (m *MaterialHelper) PrepareDefinesForLights(scene *Scene, mesh *AbstractMesh, defines interface{}, specularSupported bool, opts *MaterialHelperPrepareDefinesForLightsOpts) bool {
	if opts == nil {
		opts = &MaterialHelperPrepareDefinesForLightsOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, scene.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)
	args = append(args, specularSupported)

	if opts.MaxSimultaneousLights == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxSimultaneousLights)
	}
	if opts.DisableLighting == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableLighting)
	}

	retVal := m.p.Call("PrepareDefinesForLights", args...)
	return retVal.Bool()
}

// PrepareDefinesForMergedUV calls the PrepareDefinesForMergedUV method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesformergeduv
func (m *MaterialHelper) PrepareDefinesForMergedUV(texture *BaseTexture, defines interface{}, key string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, texture.JSObject())
	args = append(args, defines)
	args = append(args, key)

	m.p.Call("PrepareDefinesForMergedUV", args...)
}

// PrepareDefinesForMisc calls the PrepareDefinesForMisc method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesformisc
func (m *MaterialHelper) PrepareDefinesForMisc(mesh *AbstractMesh, scene *Scene, useLogarithmicDepth bool, pointsCloud bool, fogEnabled bool, alphaTest bool, defines interface{}) {

	args := make([]interface{}, 0, 7+0)

	args = append(args, mesh.JSObject())
	args = append(args, scene.JSObject())
	args = append(args, useLogarithmicDepth)
	args = append(args, pointsCloud)
	args = append(args, fogEnabled)
	args = append(args, alphaTest)
	args = append(args, defines)

	m.p.Call("PrepareDefinesForMisc", args...)
}

// PrepareDefinesForMorphTargets calls the PrepareDefinesForMorphTargets method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesformorphtargets
func (m *MaterialHelper) PrepareDefinesForMorphTargets(mesh *AbstractMesh, defines interface{}) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, defines)

	m.p.Call("PrepareDefinesForMorphTargets", args...)
}

// PrepareDefinesForMultiview calls the PrepareDefinesForMultiview method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#preparedefinesformultiview
func (m *MaterialHelper) PrepareDefinesForMultiview(scene *Scene, defines interface{}) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, defines)

	m.p.Call("PrepareDefinesForMultiview", args...)
}

// MaterialHelperPrepareUniformsAndSamplersForLightOpts contains optional parameters for MaterialHelper.PrepareUniformsAndSamplersForLight.
type MaterialHelperPrepareUniformsAndSamplersForLightOpts struct {
	ProjectedLightTexture *interface{}
	UniformBuffersList    []string
}

// PrepareUniformsAndSamplersForLight calls the PrepareUniformsAndSamplersForLight method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareuniformsandsamplersforlight
func (m *MaterialHelper) PrepareUniformsAndSamplersForLight(lightIndex float64, uniformsList []string, samplersList []string, opts *MaterialHelperPrepareUniformsAndSamplersForLightOpts) {
	if opts == nil {
		opts = &MaterialHelperPrepareUniformsAndSamplersForLightOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, lightIndex)
	args = append(args, uniformsList)
	args = append(args, samplersList)

	if opts.ProjectedLightTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ProjectedLightTexture)
	}
	if opts.UniformBuffersList == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.UniformBuffersList)
	}

	m.p.Call("PrepareUniformsAndSamplersForLight", args...)
}

// MaterialHelperPrepareUniformsAndSamplersListOpts contains optional parameters for MaterialHelper.PrepareUniformsAndSamplersList.
type MaterialHelperPrepareUniformsAndSamplersListOpts struct {
	SamplersList          []string
	Defines               *interface{}
	MaxSimultaneousLights *float64
}

// PrepareUniformsAndSamplersList calls the PrepareUniformsAndSamplersList method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#prepareuniformsandsamplerslist
func (m *MaterialHelper) PrepareUniformsAndSamplersList(uniformsListOrOptions []string, opts *MaterialHelperPrepareUniformsAndSamplersListOpts) {
	if opts == nil {
		opts = &MaterialHelperPrepareUniformsAndSamplersListOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, uniformsListOrOptions)

	if opts.SamplersList == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SamplersList)
	}
	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Defines)
	}
	if opts.MaxSimultaneousLights == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxSimultaneousLights)
	}

	m.p.Call("PrepareUniformsAndSamplersList", args...)
}

// PushAttributesForInstances calls the PushAttributesForInstances method on the MaterialHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialhelper#pushattributesforinstances
func (m *MaterialHelper) PushAttributesForInstances(attribs []string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, attribs)

	m.p.Call("PushAttributesForInstances", args...)
}
