// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PassPostProcess represents a babylon.js PassPostProcess.
// PassPostProcess which produces an output the same as it&amp;#39;s input
type PassPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (p *PassPostProcess) JSObject() js.Value { return p.p }

// PassPostProcess returns a PassPostProcess JavaScript class.
func (b *Babylon) PassPostProcess() *PassPostProcess {
	p := b.ctx.Get("PassPostProcess")
	return PassPostProcessFromJSObject(p)
}

// PassPostProcessFromJSObject returns a wrapped PassPostProcess JavaScript class.
func PassPostProcessFromJSObject(p js.Value) *PassPostProcess {
	return &PassPostProcess{PostProcessFromJSObject(p)}
}

// NewPassPostProcessOpts contains optional parameters for NewPassPostProcess.
type NewPassPostProcessOpts struct {
	Camera *Camera

	SamplingMode *float64

	Engine *Engine

	Reusable *bool

	TextureType *float64

	BlockCompilation *bool
}

// NewPassPostProcess returns a new PassPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.passpostprocess
func (b *Babylon) NewPassPostProcess(name string, options float64, opts *NewPassPostProcessOpts) *PassPostProcess {
	if opts == nil {
		opts = &NewPassPostProcessOpts{}
	}

	p := b.ctx.Get("PassPostProcess").New(name, options, opts.Camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType, opts.BlockCompilation.JSObject())
	return PassPostProcessFromJSObject(p)
}

// TODO: methods
