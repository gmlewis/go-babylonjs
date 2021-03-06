// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthOfFieldBlurPostProcess represents a babylon.js DepthOfFieldBlurPostProcess.
// The DepthOfFieldBlurPostProcess applied a blur in a give direction.
// This blur differs from the standard BlurPostProcess as it attempts to avoid blurring pixels
// based on samples that have a large difference in distance than the center pixel.
// See section 2.6.2 <a href="http://fileadmin.cs.lth.se/cs/education/edan35/lectures/12dof.pdf">http://fileadmin.cs.lth.se/cs/education/edan35/lectures/12dof.pdf</a>
type DepthOfFieldBlurPostProcess struct {
	*BlurPostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldBlurPostProcess) JSObject() js.Value { return d.p }

// DepthOfFieldBlurPostProcess returns a DepthOfFieldBlurPostProcess JavaScript class.
func (ba *Babylon) DepthOfFieldBlurPostProcess() *DepthOfFieldBlurPostProcess {
	p := ba.ctx.Get("DepthOfFieldBlurPostProcess")
	return DepthOfFieldBlurPostProcessFromJSObject(p, ba.ctx)
}

// DepthOfFieldBlurPostProcessFromJSObject returns a wrapped DepthOfFieldBlurPostProcess JavaScript class.
func DepthOfFieldBlurPostProcessFromJSObject(p js.Value, ctx js.Value) *DepthOfFieldBlurPostProcess {
	return &DepthOfFieldBlurPostProcess{BlurPostProcess: BlurPostProcessFromJSObject(p, ctx), ctx: ctx}
}

// DepthOfFieldBlurPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthOfFieldBlurPostProcessArrayToJSArray(array []*DepthOfFieldBlurPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDepthOfFieldBlurPostProcessOpts contains optional parameters for NewDepthOfFieldBlurPostProcess.
type NewDepthOfFieldBlurPostProcessOpts struct {
	ImageToBlur      *PostProcess
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	TextureType      *float64
	BlockCompilation *bool
}

// NewDepthOfFieldBlurPostProcess returns a new DepthOfFieldBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldblurpostprocess#constructor
func (ba *Babylon) NewDepthOfFieldBlurPostProcess(name string, scene *Scene, direction *Vector2, kernel float64, options float64, camera *Camera, circleOfConfusion *PostProcess, opts *NewDepthOfFieldBlurPostProcessOpts) *DepthOfFieldBlurPostProcess {
	if opts == nil {
		opts = &NewDepthOfFieldBlurPostProcessOpts{}
	}

	args := make([]interface{}, 0, 7+6)

	args = append(args, name)
	args = append(args, scene.JSObject())
	args = append(args, direction.JSObject())
	args = append(args, kernel)
	args = append(args, options)
	args = append(args, camera.JSObject())
	args = append(args, circleOfConfusion.JSObject())

	if opts.ImageToBlur == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ImageToBlur.JSObject())
	}
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

	p := ba.ctx.Get("DepthOfFieldBlurPostProcess").New(args...)
	return DepthOfFieldBlurPostProcessFromJSObject(p, ba.ctx)
}

// Direction returns the Direction property of class DepthOfFieldBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldblurpostprocess#direction
func (d *DepthOfFieldBlurPostProcess) Direction() *Vector2 {
	retVal := d.p.Get("direction")
	return Vector2FromJSObject(retVal, d.ctx)
}

// SetDirection sets the Direction property of class DepthOfFieldBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldblurpostprocess#direction
func (d *DepthOfFieldBlurPostProcess) SetDirection(direction *Vector2) *DepthOfFieldBlurPostProcess {
	d.p.Set("direction", direction.JSObject())
	return d
}
