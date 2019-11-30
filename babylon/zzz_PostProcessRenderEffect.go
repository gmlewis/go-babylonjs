// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessRenderEffect represents a babylon.js PostProcessRenderEffect.
// This represents a set of one or more post processes in Babylon.
// A post process can be used to apply a shader to a texture after it is rendered.
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocessrenderpipeline
type PostProcessRenderEffect struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PostProcessRenderEffect) JSObject() js.Value { return p.p }

// PostProcessRenderEffect returns a PostProcessRenderEffect JavaScript class.
func (ba *Babylon) PostProcessRenderEffect() *PostProcessRenderEffect {
	p := ba.ctx.Get("PostProcessRenderEffect")
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

// PostProcessRenderEffectFromJSObject returns a wrapped PostProcessRenderEffect JavaScript class.
func PostProcessRenderEffectFromJSObject(p js.Value, ctx js.Value) *PostProcessRenderEffect {
	return &PostProcessRenderEffect{p: p, ctx: ctx}
}

// NewPostProcessRenderEffectOpts contains optional parameters for NewPostProcessRenderEffect.
type NewPostProcessRenderEffectOpts struct {
	SingleInstance *JSBool
}

// NewPostProcessRenderEffect returns a new PostProcessRenderEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrendereffect
func (ba *Babylon) NewPostProcessRenderEffect(engine *Engine, name string, getPostProcesses func(), opts *NewPostProcessRenderEffectOpts) *PostProcessRenderEffect {
	if opts == nil {
		opts = &NewPostProcessRenderEffectOpts{}
	}

	p := ba.ctx.Get("PostProcessRenderEffect").New(engine.JSObject(), name, getPostProcesses, opts.SingleInstance.JSObject())
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

// TODO: methods
