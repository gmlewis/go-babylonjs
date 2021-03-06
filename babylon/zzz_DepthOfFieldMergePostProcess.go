// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthOfFieldMergePostProcess represents a babylon.js DepthOfFieldMergePostProcess.
// The DepthOfFieldMergePostProcess merges blurred images with the original based on the values of the circle of confusion.
type DepthOfFieldMergePostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldMergePostProcess) JSObject() js.Value { return d.p }

// DepthOfFieldMergePostProcess returns a DepthOfFieldMergePostProcess JavaScript class.
func (ba *Babylon) DepthOfFieldMergePostProcess() *DepthOfFieldMergePostProcess {
	p := ba.ctx.Get("DepthOfFieldMergePostProcess")
	return DepthOfFieldMergePostProcessFromJSObject(p, ba.ctx)
}

// DepthOfFieldMergePostProcessFromJSObject returns a wrapped DepthOfFieldMergePostProcess JavaScript class.
func DepthOfFieldMergePostProcessFromJSObject(p js.Value, ctx js.Value) *DepthOfFieldMergePostProcess {
	return &DepthOfFieldMergePostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// DepthOfFieldMergePostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthOfFieldMergePostProcessArrayToJSArray(array []*DepthOfFieldMergePostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDepthOfFieldMergePostProcessOpts contains optional parameters for NewDepthOfFieldMergePostProcess.
type NewDepthOfFieldMergePostProcessOpts struct {
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	TextureType      *float64
	BlockCompilation *bool
}

// NewDepthOfFieldMergePostProcess returns a new DepthOfFieldMergePostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocess#constructor
func (ba *Babylon) NewDepthOfFieldMergePostProcess(name string, originalFromInput *PostProcess, circleOfConfusion *PostProcess, blurSteps []*PostProcess, options float64, camera *Camera, opts *NewDepthOfFieldMergePostProcessOpts) *DepthOfFieldMergePostProcess {
	if opts == nil {
		opts = &NewDepthOfFieldMergePostProcessOpts{}
	}

	args := make([]interface{}, 0, 6+5)

	args = append(args, name)
	args = append(args, originalFromInput.JSObject())
	args = append(args, circleOfConfusion.JSObject())
	args = append(args, blurSteps)
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
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("DepthOfFieldMergePostProcess").New(args...)
	return DepthOfFieldMergePostProcessFromJSObject(p, ba.ctx)
}

// DepthOfFieldMergePostProcessUpdateEffectOpts contains optional parameters for DepthOfFieldMergePostProcess.UpdateEffect.
type DepthOfFieldMergePostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        []string
	Samplers        []string
	IndexParameters interface{}
	OnCompiled      JSFunc
	OnError         JSFunc
}

// UpdateEffect calls the UpdateEffect method on the DepthOfFieldMergePostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocess#updateeffect
func (d *DepthOfFieldMergePostProcess) UpdateEffect(opts *DepthOfFieldMergePostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &DepthOfFieldMergePostProcessUpdateEffectOpts{}
	}

	args := make([]interface{}, 0, 0+6)

	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.Uniforms == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Uniforms)
	}
	if opts.Samplers == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Samplers)
	}
	if opts.IndexParameters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.IndexParameters)
	}
	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnCompiled) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnError) /* never freed! */)
	}

	d.p.Call("updateEffect", args...)
}
