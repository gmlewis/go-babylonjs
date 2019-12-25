// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Effect represents a babylon.js Effect.
// Effect containing vertex and fragment shader that can be executed on an object.
type Effect struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *Effect) JSObject() js.Value { return e.p }

// Effect returns a Effect JavaScript class.
func (ba *Babylon) Effect() *Effect {
	p := ba.ctx.Get("Effect")
	return EffectFromJSObject(p, ba.ctx)
}

// EffectFromJSObject returns a wrapped Effect JavaScript class.
func EffectFromJSObject(p js.Value, ctx js.Value) *Effect {
	return &Effect{p: p, ctx: ctx}
}

// EffectArrayToJSArray returns a JavaScript Array for the wrapped array.
func EffectArrayToJSArray(array []*Effect) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEffectOpts contains optional parameters for NewEffect.
type NewEffectOpts struct {
	Samplers        []string
	Engine          *ThinEngine
	Defines         *string
	Fallbacks       *IEffectFallbacks
	OnCompiled      JSFunc
	OnError         JSFunc
	IndexParameters interface{}
}

// NewEffect returns a new Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#constructor
func (ba *Babylon) NewEffect(baseName JSObject, attributesNamesOrOptions []string, uniformsNamesOrEngine []string, opts *NewEffectOpts) *Effect {
	if opts == nil {
		opts = &NewEffectOpts{}
	}

	args := make([]interface{}, 0, 3+7)

	args = append(args, baseName.JSObject())
	args = append(args, attributesNamesOrOptions)
	args = append(args, uniformsNamesOrEngine)

	if opts.Samplers == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Samplers)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.Fallbacks == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Fallbacks.JSObject())
	}
	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnCompiled) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnError) /* never freed! */)
	}
	if opts.IndexParameters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.IndexParameters)
	}

	p := ba.ctx.Get("Effect").New(args...)
	return EffectFromJSObject(p, ba.ctx)
}

// AllFallbacksProcessed calls the AllFallbacksProcessed method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#allfallbacksprocessed
func (e *Effect) AllFallbacksProcessed() bool {

	retVal := e.p.Call("allFallbacksProcessed")
	return retVal.Bool()
}

// BindUniformBlock calls the BindUniformBlock method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#binduniformblock
func (e *Effect) BindUniformBlock(blockName string, index float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, blockName)

	args = append(args, index)

	e.p.Call("bindUniformBlock", args...)
}

// BindUniformBuffer calls the BindUniformBuffer method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#binduniformbuffer
func (e *Effect) BindUniformBuffer(buffer *DataBuffer, name string) {

	args := make([]interface{}, 0, 2+0)

	if buffer == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, buffer.JSObject())
	}

	args = append(args, name)

	e.p.Call("bindUniformBuffer", args...)
}

// Dispose calls the Dispose method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#dispose
func (e *Effect) Dispose() {

	e.p.Call("dispose")
}

// ExecuteWhenCompiled calls the ExecuteWhenCompiled method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#executewhencompiled
func (e *Effect) ExecuteWhenCompiled(jsFunc JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(jsFunc))

	e.p.Call("executeWhenCompiled", args...)
}

// GetAttributeLocation calls the GetAttributeLocation method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getattributelocation
func (e *Effect) GetAttributeLocation(index float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := e.p.Call("getAttributeLocation", args...)
	return retVal.Float()
}

// GetAttributeLocationByName calls the GetAttributeLocationByName method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getattributelocationbyname
func (e *Effect) GetAttributeLocationByName(name string) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := e.p.Call("getAttributeLocationByName", args...)
	return retVal.Float()
}

// GetAttributesCount calls the GetAttributesCount method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getattributescount
func (e *Effect) GetAttributesCount() float64 {

	retVal := e.p.Call("getAttributesCount")
	return retVal.Float()
}

// GetAttributesNames calls the GetAttributesNames method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getattributesnames
func (e *Effect) GetAttributesNames() []string {

	retVal := e.p.Call("getAttributesNames")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// GetCompilationError calls the GetCompilationError method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getcompilationerror
func (e *Effect) GetCompilationError() string {

	retVal := e.p.Call("getCompilationError")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getengine
func (e *Effect) GetEngine() *Engine {

	retVal := e.p.Call("getEngine")
	return EngineFromJSObject(retVal, e.ctx)
}

// GetPipelineContext calls the GetPipelineContext method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getpipelinecontext
func (e *Effect) GetPipelineContext() *IPipelineContext {

	retVal := e.p.Call("getPipelineContext")
	return IPipelineContextFromJSObject(retVal, e.ctx)
}

// GetSamplers calls the GetSamplers method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getsamplers
func (e *Effect) GetSamplers() []string {

	retVal := e.p.Call("getSamplers")
	result := []string{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).String())
	}
	return result
}

// GetUniform calls the GetUniform method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getuniform
func (e *Effect) GetUniform(uniformName string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uniformName)

	retVal := e.p.Call("getUniform", args...)
	return retVal
}

// GetUniformIndex calls the GetUniformIndex method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#getuniformindex
func (e *Effect) GetUniformIndex(uniformName string) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uniformName)

	retVal := e.p.Call("getUniformIndex", args...)
	return retVal.Float()
}

// IsReady calls the IsReady method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#isready
func (e *Effect) IsReady() bool {

	retVal := e.p.Call("isReady")
	return retVal.Bool()
}

// EffectRegisterShaderOpts contains optional parameters for Effect.RegisterShader.
type EffectRegisterShaderOpts struct {
	PixelShader  *string
	VertexShader *string
}

// RegisterShader calls the RegisterShader method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#registershader
func (e *Effect) RegisterShader(name string, opts *EffectRegisterShaderOpts) {
	if opts == nil {
		opts = &EffectRegisterShaderOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, name)

	if opts.PixelShader == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PixelShader)
	}
	if opts.VertexShader == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.VertexShader)
	}

	e.p.Call("RegisterShader", args...)
}

// ResetCache calls the ResetCache method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#resetcache
func (e *Effect) ResetCache() {

	e.p.Call("ResetCache")
}

// SetArray calls the SetArray method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setarray
func (e *Effect) SetArray(uniformName string, array []float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setArray", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetArray2 calls the SetArray2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setarray2
func (e *Effect) SetArray2(uniformName string, array []float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setArray2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetArray3 calls the SetArray3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setarray3
func (e *Effect) SetArray3(uniformName string, array []float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setArray3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetArray4 calls the SetArray4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setarray4
func (e *Effect) SetArray4(uniformName string, array []float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setArray4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetBool calls the SetBool method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setbool
func (e *Effect) SetBool(uniformName string, bool bool) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, bool)

	retVal := e.p.Call("setBool", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetColor3 calls the SetColor3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setcolor3
func (e *Effect) SetColor3(uniformName string, color3 js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, color3)

	retVal := e.p.Call("setColor3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetColor4 calls the SetColor4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setcolor4
func (e *Effect) SetColor4(uniformName string, color3 js.Value, alpha float64) *Effect {

	args := make([]interface{}, 0, 3+0)

	args = append(args, uniformName)

	args = append(args, color3)

	args = append(args, alpha)

	retVal := e.p.Call("setColor4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetDepthStencilTexture calls the SetDepthStencilTexture method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setdepthstenciltexture
func (e *Effect) SetDepthStencilTexture(channel string, texture *RenderTargetTexture) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, channel)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	e.p.Call("setDepthStencilTexture", args...)
}

// SetDirectColor4 calls the SetDirectColor4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setdirectcolor4
func (e *Effect) SetDirectColor4(uniformName string, color4 js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, color4)

	retVal := e.p.Call("setDirectColor4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloat calls the SetFloat method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloat
func (e *Effect) SetFloat(uniformName string, value float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, value)

	retVal := e.p.Call("setFloat", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloat2 calls the SetFloat2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloat2
func (e *Effect) SetFloat2(uniformName string, x float64, y float64) *Effect {

	args := make([]interface{}, 0, 3+0)

	args = append(args, uniformName)

	args = append(args, x)

	args = append(args, y)

	retVal := e.p.Call("setFloat2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloat3 calls the SetFloat3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloat3
func (e *Effect) SetFloat3(uniformName string, x float64, y float64, z float64) *Effect {

	args := make([]interface{}, 0, 4+0)

	args = append(args, uniformName)

	args = append(args, x)

	args = append(args, y)

	args = append(args, z)

	retVal := e.p.Call("setFloat3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloat4 calls the SetFloat4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloat4
func (e *Effect) SetFloat4(uniformName string, x float64, y float64, z float64, w float64) *Effect {

	args := make([]interface{}, 0, 5+0)

	args = append(args, uniformName)

	args = append(args, x)

	args = append(args, y)

	args = append(args, z)

	args = append(args, w)

	retVal := e.p.Call("setFloat4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloatArray calls the SetFloatArray method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloatarray
func (e *Effect) SetFloatArray(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setFloatArray", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloatArray2 calls the SetFloatArray2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloatarray2
func (e *Effect) SetFloatArray2(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setFloatArray2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloatArray3 calls the SetFloatArray3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloatarray3
func (e *Effect) SetFloatArray3(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setFloatArray3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetFloatArray4 calls the SetFloatArray4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setfloatarray4
func (e *Effect) SetFloatArray4(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setFloatArray4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetInt calls the SetInt method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setint
func (e *Effect) SetInt(uniformName string, value float64) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, value)

	retVal := e.p.Call("setInt", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetIntArray calls the SetIntArray method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setintarray
func (e *Effect) SetIntArray(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setIntArray", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetIntArray2 calls the SetIntArray2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setintarray2
func (e *Effect) SetIntArray2(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setIntArray2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetIntArray3 calls the SetIntArray3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setintarray3
func (e *Effect) SetIntArray3(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setIntArray3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetIntArray4 calls the SetIntArray4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setintarray4
func (e *Effect) SetIntArray4(uniformName string, array js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, array)

	retVal := e.p.Call("setIntArray4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetMatrices calls the SetMatrices method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setmatrices
func (e *Effect) SetMatrices(uniformName string, matrices js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, matrices)

	retVal := e.p.Call("setMatrices", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetMatrix calls the SetMatrix method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setmatrix
func (e *Effect) SetMatrix(uniformName string, matrix js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, matrix)

	retVal := e.p.Call("setMatrix", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetMatrix2x2 calls the SetMatrix2x2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setmatrix2x2
func (e *Effect) SetMatrix2x2(uniformName string, matrix js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, matrix)

	retVal := e.p.Call("setMatrix2x2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetMatrix3x3 calls the SetMatrix3x3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setmatrix3x3
func (e *Effect) SetMatrix3x3(uniformName string, matrix js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, matrix)

	retVal := e.p.Call("setMatrix3x3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetTexture calls the SetTexture method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#settexture
func (e *Effect) SetTexture(channel string, texture *BaseTexture) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, channel)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	e.p.Call("setTexture", args...)
}

// SetTextureArray calls the SetTextureArray method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#settexturearray
func (e *Effect) SetTextureArray(channel string, textures []*BaseTexture) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, channel)

	args = append(args, BaseTextureArrayToJSArray(textures))

	e.p.Call("setTextureArray", args...)
}

// SetTextureFromPostProcess calls the SetTextureFromPostProcess method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#settexturefrompostprocess
func (e *Effect) SetTextureFromPostProcess(channel string, postProcess *PostProcess) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, channel)

	if postProcess == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, postProcess.JSObject())
	}

	e.p.Call("setTextureFromPostProcess", args...)
}

// SetTextureFromPostProcessOutput calls the SetTextureFromPostProcessOutput method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#settexturefrompostprocessoutput
func (e *Effect) SetTextureFromPostProcessOutput(channel string, postProcess *PostProcess) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, channel)

	if postProcess == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, postProcess.JSObject())
	}

	e.p.Call("setTextureFromPostProcessOutput", args...)
}

// SetVector2 calls the SetVector2 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setvector2
func (e *Effect) SetVector2(uniformName string, vector2 js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, vector2)

	retVal := e.p.Call("setVector2", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetVector3 calls the SetVector3 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setvector3
func (e *Effect) SetVector3(uniformName string, vector3 js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, vector3)

	retVal := e.p.Call("setVector3", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// SetVector4 calls the SetVector4 method on the Effect object.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#setvector4
func (e *Effect) SetVector4(uniformName string, vector4 js.Value) *Effect {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniformName)

	args = append(args, vector4)

	retVal := e.p.Call("setVector4", args...)
	return EffectFromJSObject(retVal, e.ctx)
}

// Defines returns the Defines property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#defines
func (e *Effect) Defines() string {
	retVal := e.p.Get("defines")
	return retVal.String()
}

// SetDefines sets the Defines property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#defines
func (e *Effect) SetDefines(defines string) *Effect {
	e.p.Set("defines", defines)
	return e
}

// IncludesShadersStore returns the IncludesShadersStore property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#includesshadersstore
func (e *Effect) IncludesShadersStore() js.Value {
	retVal := e.p.Get("IncludesShadersStore")
	return retVal
}

// SetIncludesShadersStore sets the IncludesShadersStore property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#includesshadersstore
func (e *Effect) SetIncludesShadersStore(IncludesShadersStore js.Value) *Effect {
	e.p.Set("IncludesShadersStore", IncludesShadersStore)
	return e
}

// IsSupported returns the IsSupported property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#issupported
func (e *Effect) IsSupported() bool {
	retVal := e.p.Get("isSupported")
	return retVal.Bool()
}

// SetIsSupported sets the IsSupported property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#issupported
func (e *Effect) SetIsSupported(isSupported bool) *Effect {
	e.p.Set("isSupported", isSupported)
	return e
}

// Key returns the Key property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#key
func (e *Effect) Key() string {
	retVal := e.p.Get("key")
	return retVal.String()
}

// SetKey sets the Key property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#key
func (e *Effect) SetKey(key string) *Effect {
	e.p.Set("key", key)
	return e
}

// Name returns the Name property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#name
func (e *Effect) Name() js.Value {
	retVal := e.p.Get("name")
	return retVal
}

// SetName sets the Name property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#name
func (e *Effect) SetName(name JSObject) *Effect {
	e.p.Set("name", name.JSObject())
	return e
}

// OnBind returns the OnBind property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onbind
func (e *Effect) OnBind() js.Value {
	retVal := e.p.Get("onBind")
	return retVal
}

// SetOnBind sets the OnBind property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onbind
func (e *Effect) SetOnBind(onBind JSFunc) *Effect {
	e.p.Set("onBind", js.FuncOf(onBind))
	return e
}

// OnBindObservable returns the OnBindObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onbindobservable
func (e *Effect) OnBindObservable() *Observable {
	retVal := e.p.Get("onBindObservable")
	return ObservableFromJSObject(retVal, e.ctx)
}

// SetOnBindObservable sets the OnBindObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onbindobservable
func (e *Effect) SetOnBindObservable(onBindObservable *Observable) *Effect {
	e.p.Set("onBindObservable", onBindObservable.JSObject())
	return e
}

// OnCompileObservable returns the OnCompileObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#oncompileobservable
func (e *Effect) OnCompileObservable() *Observable {
	retVal := e.p.Get("onCompileObservable")
	return ObservableFromJSObject(retVal, e.ctx)
}

// SetOnCompileObservable sets the OnCompileObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#oncompileobservable
func (e *Effect) SetOnCompileObservable(onCompileObservable *Observable) *Effect {
	e.p.Set("onCompileObservable", onCompileObservable.JSObject())
	return e
}

// OnCompiled returns the OnCompiled property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#oncompiled
func (e *Effect) OnCompiled() js.Value {
	retVal := e.p.Get("onCompiled")
	return retVal
}

// SetOnCompiled sets the OnCompiled property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#oncompiled
func (e *Effect) SetOnCompiled(onCompiled JSFunc) *Effect {
	e.p.Set("onCompiled", js.FuncOf(onCompiled))
	return e
}

// OnError returns the OnError property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onerror
func (e *Effect) OnError() js.Value {
	retVal := e.p.Get("onError")
	return retVal
}

// SetOnError sets the OnError property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onerror
func (e *Effect) SetOnError(onError JSFunc) *Effect {
	e.p.Set("onError", js.FuncOf(onError))
	return e
}

// OnErrorObservable returns the OnErrorObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onerrorobservable
func (e *Effect) OnErrorObservable() *Observable {
	retVal := e.p.Get("onErrorObservable")
	return ObservableFromJSObject(retVal, e.ctx)
}

// SetOnErrorObservable sets the OnErrorObservable property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#onerrorobservable
func (e *Effect) SetOnErrorObservable(onErrorObservable *Observable) *Effect {
	e.p.Set("onErrorObservable", onErrorObservable.JSObject())
	return e
}

// ShadersRepository returns the ShadersRepository property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#shadersrepository
func (e *Effect) ShadersRepository() string {
	retVal := e.p.Get("ShadersRepository")
	return retVal.String()
}

// SetShadersRepository sets the ShadersRepository property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#shadersrepository
func (e *Effect) SetShadersRepository(ShadersRepository string) *Effect {
	e.p.Set("ShadersRepository", ShadersRepository)
	return e
}

// ShadersStore returns the ShadersStore property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#shadersstore
func (e *Effect) ShadersStore() js.Value {
	retVal := e.p.Get("ShadersStore")
	return retVal
}

// SetShadersStore sets the ShadersStore property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#shadersstore
func (e *Effect) SetShadersStore(ShadersStore js.Value) *Effect {
	e.p.Set("ShadersStore", ShadersStore)
	return e
}

// UniqueId returns the UniqueId property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#uniqueid
func (e *Effect) UniqueId() float64 {
	retVal := e.p.Get("uniqueId")
	return retVal.Float()
}

// SetUniqueId sets the UniqueId property of class Effect.
//
// https://doc.babylonjs.com/api/classes/babylon.effect#uniqueid
func (e *Effect) SetUniqueId(uniqueId float64) *Effect {
	e.p.Set("uniqueId", uniqueId)
	return e
}
