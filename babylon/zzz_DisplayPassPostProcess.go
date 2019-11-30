// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DisplayPassPostProcess represents a babylon.js DisplayPassPostProcess.
// DisplayPassPostProcess which produces an output the same as it&amp;#39;s input
type DisplayPassPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (d *DisplayPassPostProcess) JSObject() js.Value { return d.p }

// DisplayPassPostProcess returns a DisplayPassPostProcess JavaScript class.
func (b *Babylon) DisplayPassPostProcess() *DisplayPassPostProcess {
	p := b.ctx.Get("DisplayPassPostProcess")
	return DisplayPassPostProcessFromJSObject(p)
}

// DisplayPassPostProcessFromJSObject returns a wrapped DisplayPassPostProcess JavaScript class.
func DisplayPassPostProcessFromJSObject(p js.Value) *DisplayPassPostProcess {
	return &DisplayPassPostProcess{PostProcessFromJSObject(p)}
}

// NewDisplayPassPostProcessOpts contains optional parameters for NewDisplayPassPostProcess.
type NewDisplayPassPostProcessOpts struct {
	SamplingMode *float64

	Engine *Engine

	Reusable *bool
}

// NewDisplayPassPostProcess returns a new DisplayPassPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.displaypasspostprocess
func (b *Babylon) NewDisplayPassPostProcess(name string, options float64, camera *Camera, opts *NewDisplayPassPostProcessOpts) *DisplayPassPostProcess {
	if opts == nil {
		opts = &NewDisplayPassPostProcessOpts{}
	}

	p := b.ctx.Get("DisplayPassPostProcess").New(name, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject())
	return DisplayPassPostProcessFromJSObject(p)
}

// TODO: methods
