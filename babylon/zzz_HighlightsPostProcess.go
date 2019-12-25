// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HighlightsPostProcess represents a babylon.js HighlightsPostProcess.
// Extracts highlights from the image
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses
type HighlightsPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HighlightsPostProcess) JSObject() js.Value { return h.p }

// HighlightsPostProcess returns a HighlightsPostProcess JavaScript class.
func (ba *Babylon) HighlightsPostProcess() *HighlightsPostProcess {
	p := ba.ctx.Get("HighlightsPostProcess")
	return HighlightsPostProcessFromJSObject(p, ba.ctx)
}

// HighlightsPostProcessFromJSObject returns a wrapped HighlightsPostProcess JavaScript class.
func HighlightsPostProcessFromJSObject(p js.Value, ctx js.Value) *HighlightsPostProcess {
	return &HighlightsPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// HighlightsPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func HighlightsPostProcessArrayToJSArray(array []*HighlightsPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHighlightsPostProcessOpts contains optional parameters for NewHighlightsPostProcess.
type NewHighlightsPostProcessOpts struct {
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
	TextureType  *float64
}

// NewHighlightsPostProcess returns a new HighlightsPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.highlightspostprocess#constructor
func (ba *Babylon) NewHighlightsPostProcess(name string, options float64, camera *Camera, opts *NewHighlightsPostProcessOpts) *HighlightsPostProcess {
	if opts == nil {
		opts = &NewHighlightsPostProcessOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, name)
	args = append(args, options)
	args = append(args, camera.JSObject())

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Reusable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reusable)
	}
	if opts.TextureType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureType)
	}

	p := ba.ctx.Get("HighlightsPostProcess").New(args...)
	return HighlightsPostProcessFromJSObject(p, ba.ctx)
}
