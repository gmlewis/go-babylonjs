// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HighlightsPostProcess represents a babylon.js HighlightsPostProcess.
// Extracts highlights from the image
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses
type HighlightsPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (h *HighlightsPostProcess) JSObject() js.Value { return h.p }

// HighlightsPostProcess returns a HighlightsPostProcess JavaScript class.
func (b *Babylon) HighlightsPostProcess() *HighlightsPostProcess {
	p := b.ctx.Get("HighlightsPostProcess")
	return HighlightsPostProcessFromJSObject(p)
}

// HighlightsPostProcessFromJSObject returns a wrapped HighlightsPostProcess JavaScript class.
func HighlightsPostProcessFromJSObject(p js.Value) *HighlightsPostProcess {
	return &HighlightsPostProcess{PostProcessFromJSObject(p)}
}

// NewHighlightsPostProcessOpts contains optional parameters for NewHighlightsPostProcess.
type NewHighlightsPostProcessOpts struct {
	SamplingMode *float64

	Engine *Engine

	Reusable *bool

	TextureType *float64
}

// NewHighlightsPostProcess returns a new HighlightsPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.highlightspostprocess
func (b *Babylon) NewHighlightsPostProcess(name string, options float64, camera *Camera, opts *NewHighlightsPostProcessOpts) *HighlightsPostProcess {
	if opts == nil {
		opts = &NewHighlightsPostProcessOpts{}
	}

	p := b.ctx.Get("HighlightsPostProcess").New(name, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType)
	return HighlightsPostProcessFromJSObject(p)
}

// TODO: methods
