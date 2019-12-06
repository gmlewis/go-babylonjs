// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IEffectCreationOptions represents a babylon.js IEffectCreationOptions.
// Options to be used when creating an effect.
type IEffectCreationOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IEffectCreationOptions) JSObject() js.Value { return i.p }

// IEffectCreationOptions returns a IEffectCreationOptions JavaScript class.
func (ba *Babylon) IEffectCreationOptions() *IEffectCreationOptions {
	p := ba.ctx.Get("IEffectCreationOptions")
	return IEffectCreationOptionsFromJSObject(p, ba.ctx)
}

// IEffectCreationOptionsFromJSObject returns a wrapped IEffectCreationOptions JavaScript class.
func IEffectCreationOptionsFromJSObject(p js.Value, ctx js.Value) *IEffectCreationOptions {
	return &IEffectCreationOptions{p: p, ctx: ctx}
}

// IEffectCreationOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IEffectCreationOptionsArrayToJSArray(array []*IEffectCreationOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Attributes returns the Attributes property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#attributes
func (i *IEffectCreationOptions) Attributes() []string {
	retVal := i.p.Get("attributes")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// SetAttributes sets the Attributes property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#attributes
func (i *IEffectCreationOptions) SetAttributes(attributes []string) *IEffectCreationOptions {
	i.p.Set("attributes", attributes)
	return i
}

// Defines returns the Defines property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#defines
func (i *IEffectCreationOptions) Defines() interface{} {
	retVal := i.p.Get("defines")
	return retVal
}

// SetDefines sets the Defines property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#defines
func (i *IEffectCreationOptions) SetDefines(defines interface{}) *IEffectCreationOptions {
	i.p.Set("defines", defines)
	return i
}

// Fallbacks returns the Fallbacks property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#fallbacks
func (i *IEffectCreationOptions) Fallbacks() *IEffectFallbacks {
	retVal := i.p.Get("fallbacks")
	return IEffectFallbacksFromJSObject(retVal, i.ctx)
}

// SetFallbacks sets the Fallbacks property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#fallbacks
func (i *IEffectCreationOptions) SetFallbacks(fallbacks *IEffectFallbacks) *IEffectCreationOptions {
	i.p.Set("fallbacks", fallbacks.JSObject())
	return i
}

// IndexParameters returns the IndexParameters property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#indexparameters
func (i *IEffectCreationOptions) IndexParameters() interface{} {
	retVal := i.p.Get("indexParameters")
	return retVal
}

// SetIndexParameters sets the IndexParameters property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#indexparameters
func (i *IEffectCreationOptions) SetIndexParameters(indexParameters interface{}) *IEffectCreationOptions {
	i.p.Set("indexParameters", indexParameters)
	return i
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#maxsimultaneouslights
func (i *IEffectCreationOptions) MaxSimultaneousLights() float64 {
	retVal := i.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#maxsimultaneouslights
func (i *IEffectCreationOptions) SetMaxSimultaneousLights(maxSimultaneousLights float64) *IEffectCreationOptions {
	i.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return i
}

// OnCompiled returns the OnCompiled property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#oncompiled
func (i *IEffectCreationOptions) OnCompiled() js.Value {
	retVal := i.p.Get("onCompiled")
	return retVal
}

// SetOnCompiled sets the OnCompiled property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#oncompiled
func (i *IEffectCreationOptions) SetOnCompiled(onCompiled func()) *IEffectCreationOptions {
	i.p.Set("onCompiled", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onCompiled(); return nil }))
	return i
}

// OnError returns the OnError property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#onerror
func (i *IEffectCreationOptions) OnError() js.Value {
	retVal := i.p.Get("onError")
	return retVal
}

// SetOnError sets the OnError property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#onerror
func (i *IEffectCreationOptions) SetOnError(onError func()) *IEffectCreationOptions {
	i.p.Set("onError", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))
	return i
}

// Samplers returns the Samplers property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#samplers
func (i *IEffectCreationOptions) Samplers() []string {
	retVal := i.p.Get("samplers")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// SetSamplers sets the Samplers property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#samplers
func (i *IEffectCreationOptions) SetSamplers(samplers []string) *IEffectCreationOptions {
	i.p.Set("samplers", samplers)
	return i
}

// TransformFeedbackVaryings returns the TransformFeedbackVaryings property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#transformfeedbackvaryings
func (i *IEffectCreationOptions) TransformFeedbackVaryings() []string {
	retVal := i.p.Get("transformFeedbackVaryings")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// SetTransformFeedbackVaryings sets the TransformFeedbackVaryings property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#transformfeedbackvaryings
func (i *IEffectCreationOptions) SetTransformFeedbackVaryings(transformFeedbackVaryings []string) *IEffectCreationOptions {
	i.p.Set("transformFeedbackVaryings", transformFeedbackVaryings)
	return i
}

// UniformBuffersNames returns the UniformBuffersNames property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#uniformbuffersnames
func (i *IEffectCreationOptions) UniformBuffersNames() []string {
	retVal := i.p.Get("uniformBuffersNames")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// SetUniformBuffersNames sets the UniformBuffersNames property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#uniformbuffersnames
func (i *IEffectCreationOptions) SetUniformBuffersNames(uniformBuffersNames []string) *IEffectCreationOptions {
	i.p.Set("uniformBuffersNames", uniformBuffersNames)
	return i
}

// UniformsNames returns the UniformsNames property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#uniformsnames
func (i *IEffectCreationOptions) UniformsNames() []string {
	retVal := i.p.Get("uniformsNames")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// SetUniformsNames sets the UniformsNames property of class IEffectCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectcreationoptions#uniformsnames
func (i *IEffectCreationOptions) SetUniformsNames(uniformsNames []string) *IEffectCreationOptions {
	i.p.Set("uniformsNames", uniformsNames)
	return i
}
