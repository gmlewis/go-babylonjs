// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensRenderingPipeline represents a babylon.js LensRenderingPipeline.
// BABYLON.JS Chromatic Aberration GLSL Shader
// Author: Olivier Guyot
// Separates very slightly R, G and B colors on the edges of the screen
// Inspired by Francois Tarlier &amp;amp; Martins Upitis
type LensRenderingPipeline struct{ *PostProcessRenderPipeline }

// JSObject returns the underlying js.Value.
func (l *LensRenderingPipeline) JSObject() js.Value { return l.p }

// LensRenderingPipeline returns a LensRenderingPipeline JavaScript class.
func (b *Babylon) LensRenderingPipeline() *LensRenderingPipeline {
	p := b.ctx.Get("LensRenderingPipeline")
	return LensRenderingPipelineFromJSObject(p)
}

// LensRenderingPipelineFromJSObject returns a wrapped LensRenderingPipeline JavaScript class.
func LensRenderingPipelineFromJSObject(p js.Value) *LensRenderingPipeline {
	return &LensRenderingPipeline{PostProcessRenderPipelineFromJSObject(p)}
}

// NewLensRenderingPipelineOpts contains optional parameters for NewLensRenderingPipeline.
type NewLensRenderingPipelineOpts struct {
	Ratio *float64

	Cameras *Camera
}

// NewLensRenderingPipeline returns a new LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline
func (b *Babylon) NewLensRenderingPipeline(name string, parameters interface{}, scene *Scene, opts *NewLensRenderingPipelineOpts) *LensRenderingPipeline {
	if opts == nil {
		opts = &NewLensRenderingPipelineOpts{}
	}

	p := b.ctx.Get("LensRenderingPipeline").New(name, parameters, scene.JSObject(), opts.Ratio, opts.Cameras.JSObject())
	return LensRenderingPipelineFromJSObject(p)
}

// TODO: methods
