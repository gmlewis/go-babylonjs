// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessRenderPipeline represents a babylon.js PostProcessRenderPipeline.
// PostProcessRenderPipeline

//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocessrenderpipeline
type PostProcessRenderPipeline struct{}

// JSObject returns the underlying js.Value.
func (p *PostProcessRenderPipeline) JSObject() js.Value { return p.p }

// PostProcessRenderPipeline returns a PostProcessRenderPipeline JavaScript class.
func (b *Babylon) PostProcessRenderPipeline() *PostProcessRenderPipeline {
	p := b.ctx.Get("PostProcessRenderPipeline")
	return PostProcessRenderPipelineFromJSObject(p)
}

// PostProcessRenderPipelineFromJSObject returns a wrapped PostProcessRenderPipeline JavaScript class.
func PostProcessRenderPipelineFromJSObject(p js.Value) *PostProcessRenderPipeline {
	return &PostProcessRenderPipeline{p: p}
}

// NewPostProcessRenderPipeline returns a new PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline
func (b *Babylon) NewPostProcessRenderPipeline(todo parameters) *PostProcessRenderPipeline {
	p := b.ctx.Get("PostProcessRenderPipeline").New(todo)
	return PostProcessRenderPipelineFromJSObject(p)
}

// TODO: methods