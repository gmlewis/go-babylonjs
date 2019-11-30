// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessRenderPipeline represents a babylon.js PostProcessRenderPipeline.
// PostProcessRenderPipeline
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocessrenderpipeline
type PostProcessRenderPipeline struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PostProcessRenderPipeline) JSObject() js.Value { return p.p }

// PostProcessRenderPipeline returns a PostProcessRenderPipeline JavaScript class.
func (ba *Babylon) PostProcessRenderPipeline() *PostProcessRenderPipeline {
	p := ba.ctx.Get("PostProcessRenderPipeline")
	return PostProcessRenderPipelineFromJSObject(p, ba.ctx)
}

// PostProcessRenderPipelineFromJSObject returns a wrapped PostProcessRenderPipeline JavaScript class.
func PostProcessRenderPipelineFromJSObject(p js.Value, ctx js.Value) *PostProcessRenderPipeline {
	return &PostProcessRenderPipeline{p: p, ctx: ctx}
}

// NewPostProcessRenderPipeline returns a new PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline
func (ba *Babylon) NewPostProcessRenderPipeline(engine *Engine, name string) *PostProcessRenderPipeline {
	p := ba.ctx.Get("PostProcessRenderPipeline").New(engine.JSObject(), name)
	return PostProcessRenderPipelineFromJSObject(p, ba.ctx)
}

// TODO: methods
