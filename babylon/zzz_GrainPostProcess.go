// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GrainPostProcess represents a babylon.js GrainPostProcess.
// The GrainPostProcess adds noise to the image at mid luminance levels
type GrainPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (g *GrainPostProcess) JSObject() js.Value { return g.p }

// GrainPostProcess returns a GrainPostProcess JavaScript class.
func (b *Babylon) GrainPostProcess() *GrainPostProcess {
	p := b.ctx.Get("GrainPostProcess")
	return GrainPostProcessFromJSObject(p)
}

// GrainPostProcessFromJSObject returns a wrapped GrainPostProcess JavaScript class.
func GrainPostProcessFromJSObject(p js.Value) *GrainPostProcess {
	return &GrainPostProcess{PostProcessFromJSObject(p)}
}

// NewGrainPostProcessOpts contains optional parameters for NewGrainPostProcess.
type NewGrainPostProcessOpts struct {
	SamplingMode *float64

	Engine *Engine

	Reusable *bool

	TextureType *float64

	BlockCompilation *bool
}

// NewGrainPostProcess returns a new GrainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.grainpostprocess
func (b *Babylon) NewGrainPostProcess(name string, options float64, camera *Camera, opts *NewGrainPostProcessOpts) *GrainPostProcess {
	if opts == nil {
		opts = &NewGrainPostProcessOpts{}
	}

	p := b.ctx.Get("GrainPostProcess").New(name, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType, opts.BlockCompilation.JSObject())
	return GrainPostProcessFromJSObject(p)
}

// TODO: methods
