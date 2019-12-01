// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderTargetCreationOptions represents a babylon.js RenderTargetCreationOptions.
// Define options used to create a render target texture
type RenderTargetCreationOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RenderTargetCreationOptions) JSObject() js.Value { return r.p }

// RenderTargetCreationOptions returns a RenderTargetCreationOptions JavaScript class.
func (ba *Babylon) RenderTargetCreationOptions() *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions")
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// RenderTargetCreationOptionsFromJSObject returns a wrapped RenderTargetCreationOptions JavaScript class.
func RenderTargetCreationOptionsFromJSObject(p js.Value, ctx js.Value) *RenderTargetCreationOptions {
	return &RenderTargetCreationOptions{p: p, ctx: ctx}
}

// RenderTargetCreationOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func RenderTargetCreationOptionsArrayToJSArray(array []*RenderTargetCreationOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Format returns the Format property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#format
func (r *RenderTargetCreationOptions) Format(format float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(format)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetFormat sets the Format property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#format
func (r *RenderTargetCreationOptions) SetFormat(format float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(format)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// GenerateDepthBuffer returns the GenerateDepthBuffer property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatedepthbuffer
func (r *RenderTargetCreationOptions) GenerateDepthBuffer(generateDepthBuffer bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateDepthBuffer)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetGenerateDepthBuffer sets the GenerateDepthBuffer property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatedepthbuffer
func (r *RenderTargetCreationOptions) SetGenerateDepthBuffer(generateDepthBuffer bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateDepthBuffer)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// GenerateMipMaps returns the GenerateMipMaps property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatemipmaps
func (r *RenderTargetCreationOptions) GenerateMipMaps(generateMipMaps bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateMipMaps)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetGenerateMipMaps sets the GenerateMipMaps property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatemipmaps
func (r *RenderTargetCreationOptions) SetGenerateMipMaps(generateMipMaps bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateMipMaps)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// GenerateStencilBuffer returns the GenerateStencilBuffer property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatestencilbuffer
func (r *RenderTargetCreationOptions) GenerateStencilBuffer(generateStencilBuffer bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateStencilBuffer)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetGenerateStencilBuffer sets the GenerateStencilBuffer property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#generatestencilbuffer
func (r *RenderTargetCreationOptions) SetGenerateStencilBuffer(generateStencilBuffer bool) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(generateStencilBuffer)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SamplingMode returns the SamplingMode property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#samplingmode
func (r *RenderTargetCreationOptions) SamplingMode(samplingMode float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(samplingMode)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetSamplingMode sets the SamplingMode property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#samplingmode
func (r *RenderTargetCreationOptions) SetSamplingMode(samplingMode float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(samplingMode)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#type
func (r *RenderTargetCreationOptions) Type(jsType float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(jsType)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class RenderTargetCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetcreationoptions#type
func (r *RenderTargetCreationOptions) SetType(jsType float64) *RenderTargetCreationOptions {
	p := ba.ctx.Get("RenderTargetCreationOptions").New(jsType)
	return RenderTargetCreationOptionsFromJSObject(p, ba.ctx)
}

*/
