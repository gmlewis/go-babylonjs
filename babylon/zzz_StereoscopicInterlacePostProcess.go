// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StereoscopicInterlacePostProcess represents a babylon.js StereoscopicInterlacePostProcess.
// StereoscopicInterlacePostProcess used to render stereo views from a rigged camera
type StereoscopicInterlacePostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StereoscopicInterlacePostProcess) JSObject() js.Value { return s.p }

// StereoscopicInterlacePostProcess returns a StereoscopicInterlacePostProcess JavaScript class.
func (ba *Babylon) StereoscopicInterlacePostProcess() *StereoscopicInterlacePostProcess {
	p := ba.ctx.Get("StereoscopicInterlacePostProcess")
	return StereoscopicInterlacePostProcessFromJSObject(p, ba.ctx)
}

// StereoscopicInterlacePostProcessFromJSObject returns a wrapped StereoscopicInterlacePostProcess JavaScript class.
func StereoscopicInterlacePostProcessFromJSObject(p js.Value, ctx js.Value) *StereoscopicInterlacePostProcess {
	return &StereoscopicInterlacePostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// StereoscopicInterlacePostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func StereoscopicInterlacePostProcessArrayToJSArray(array []*StereoscopicInterlacePostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStereoscopicInterlacePostProcessOpts contains optional parameters for NewStereoscopicInterlacePostProcess.
type NewStereoscopicInterlacePostProcessOpts struct {
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
}

// NewStereoscopicInterlacePostProcess returns a new StereoscopicInterlacePostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.stereoscopicinterlacepostprocess#constructor
func (ba *Babylon) NewStereoscopicInterlacePostProcess(name string, rigCameras []*Camera, isStereoscopicHoriz bool, opts *NewStereoscopicInterlacePostProcessOpts) *StereoscopicInterlacePostProcess {
	if opts == nil {
		opts = &NewStereoscopicInterlacePostProcessOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, name)
	args = append(args, rigCameras)
	args = append(args, isStereoscopicHoriz)

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

	p := ba.ctx.Get("StereoscopicInterlacePostProcess").New(args...)
	return StereoscopicInterlacePostProcessFromJSObject(p, ba.ctx)
}
