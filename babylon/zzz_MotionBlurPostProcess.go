// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MotionBlurPostProcess represents a babylon.js MotionBlurPostProcess.
// The Motion Blur Post Process which blurs an image based on the objects velocity in scene.
// Velocity can be affected by each object&amp;#39;s rotation, position and scale depending on the transformation speed.
// As an example, all you have to do is to create the post-process:
// var mb = new BABYLON.MotionBlurPostProcess(
// &amp;#39;mb&amp;#39;, // The name of the effect.
// scene, // The scene containing the objects to blur according to their velocity.
// 1.0, // The required width/height ratio to downsize to before computing the render pass.
// camera // The camera to apply the render pass to.
// );
// Then, all objects moving, rotating and/or scaling will be blurred depending on the transformation speed.
type MotionBlurPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MotionBlurPostProcess) JSObject() js.Value { return m.p }

// MotionBlurPostProcess returns a MotionBlurPostProcess JavaScript class.
func (ba *Babylon) MotionBlurPostProcess() *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess")
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// MotionBlurPostProcessFromJSObject returns a wrapped MotionBlurPostProcess JavaScript class.
func MotionBlurPostProcessFromJSObject(p js.Value, ctx js.Value) *MotionBlurPostProcess {
	return &MotionBlurPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// NewMotionBlurPostProcessOpts contains optional parameters for NewMotionBlurPostProcess.
type NewMotionBlurPostProcessOpts struct {
	SamplingMode *JSFloat64

	Engine *Engine

	Reusable *JSBool

	TextureType *JSFloat64

	BlockCompilation *JSBool
}

// NewMotionBlurPostProcess returns a new MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess
func (ba *Babylon) NewMotionBlurPostProcess(name string, scene *Scene, options float64, camera *Camera, opts *NewMotionBlurPostProcessOpts) *MotionBlurPostProcess {
	if opts == nil {
		opts = &NewMotionBlurPostProcessOpts{}
	}

	p := ba.ctx.Get("MotionBlurPostProcess").New(name, scene.JSObject(), options, camera.JSObject(), opts.SamplingMode.JSObject(), opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType.JSObject(), opts.BlockCompilation.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// TODO: methods
