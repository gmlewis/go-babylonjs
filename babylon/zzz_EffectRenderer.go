// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectRenderer represents a babylon.js EffectRenderer.
// Helper class to render one or more effects
type EffectRenderer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EffectRenderer) JSObject() js.Value { return e.p }

// EffectRenderer returns a EffectRenderer JavaScript class.
func (ba *Babylon) EffectRenderer() *EffectRenderer {
	p := ba.ctx.Get("EffectRenderer")
	return EffectRendererFromJSObject(p, ba.ctx)
}

// EffectRendererFromJSObject returns a wrapped EffectRenderer JavaScript class.
func EffectRendererFromJSObject(p js.Value, ctx js.Value) *EffectRenderer {
	return &EffectRenderer{p: p, ctx: ctx}
}

// EffectRendererArrayToJSArray returns a JavaScript Array for the wrapped array.
func EffectRendererArrayToJSArray(array []*EffectRenderer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEffectRendererOpts contains optional parameters for NewEffectRenderer.
type NewEffectRendererOpts struct {
	Options *IEffectRendererOptions
}

// NewEffectRenderer returns a new EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#constructor
func (ba *Babylon) NewEffectRenderer(engine *ThinEngine, opts *NewEffectRendererOpts) *EffectRenderer {
	if opts == nil {
		opts = &NewEffectRendererOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, engine.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("EffectRenderer").New(args...)
	return EffectRendererFromJSObject(p, ba.ctx)
}

// ApplyEffectWrapper calls the ApplyEffectWrapper method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#applyeffectwrapper
func (e *EffectRenderer) ApplyEffectWrapper(effectWrapper *EffectWrapper) {

	args := make([]interface{}, 0, 1+0)

	if effectWrapper == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effectWrapper.JSObject())
	}

	e.p.Call("applyEffectWrapper", args...)
}

// BindBuffers calls the BindBuffers method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#bindbuffers
func (e *EffectRenderer) BindBuffers(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	e.p.Call("bindBuffers", args...)
}

// Dispose calls the Dispose method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#dispose
func (e *EffectRenderer) Dispose() {

	e.p.Call("dispose")
}

// Draw calls the Draw method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#draw
func (e *EffectRenderer) Draw() {

	e.p.Call("draw")
}

// EffectRendererRenderOpts contains optional parameters for EffectRenderer.Render.
type EffectRendererRenderOpts struct {
	OutputTexture *Texture
}

// Render calls the Render method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#render
func (e *EffectRenderer) Render(effectWrappers []*EffectWrapper, opts *EffectRendererRenderOpts) {
	if opts == nil {
		opts = &EffectRendererRenderOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, EffectWrapperArrayToJSArray(effectWrappers))

	if opts.OutputTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OutputTexture.JSObject())
	}

	e.p.Call("render", args...)
}

// EffectRendererSetViewportOpts contains optional parameters for EffectRenderer.SetViewport.
type EffectRendererSetViewportOpts struct {
	Viewport *Viewport
}

// SetViewport calls the SetViewport method on the EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer#setviewport
func (e *EffectRenderer) SetViewport(opts *EffectRendererSetViewportOpts) {
	if opts == nil {
		opts = &EffectRendererSetViewportOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Viewport == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Viewport.JSObject())
	}

	e.p.Call("setViewport", args...)
}
