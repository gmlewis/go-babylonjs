// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RefractionPostProcess represents a babylon.js RefractionPostProcess.
// Post process which applies a refractin texture
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses#refraction
type RefractionPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (r *RefractionPostProcess) JSObject() js.Value { return r.p }

// RefractionPostProcess returns a RefractionPostProcess JavaScript class.
func (b *Babylon) RefractionPostProcess() *RefractionPostProcess {
	p := b.ctx.Get("RefractionPostProcess")
	return RefractionPostProcessFromJSObject(p)
}

// RefractionPostProcessFromJSObject returns a wrapped RefractionPostProcess JavaScript class.
func RefractionPostProcessFromJSObject(p js.Value) *RefractionPostProcess {
	return &RefractionPostProcess{PostProcessFromJSObject(p)}
}

// NewRefractionPostProcessOpts contains optional parameters for NewRefractionPostProcess.
type NewRefractionPostProcessOpts struct {
	SamplingMode *float64

	Engine *Engine

	Reusable *bool
}

// NewRefractionPostProcess returns a new RefractionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess
func (b *Babylon) NewRefractionPostProcess(name string, refractionTextureUrl string, color *Color3, depth float64, colorLevel float64, options float64, camera *Camera, opts *NewRefractionPostProcessOpts) *RefractionPostProcess {
	if opts == nil {
		opts = &NewRefractionPostProcessOpts{}
	}

	p := b.ctx.Get("RefractionPostProcess").New(name, refractionTextureUrl, color.JSObject(), depth, colorLevel, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject())
	return RefractionPostProcessFromJSObject(p)
}

// TODO: methods
