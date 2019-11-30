// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageProcessingPostProcess represents a babylon.js ImageProcessingPostProcess.
// ImageProcessingPostProcess
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses#imageprocessing
type ImageProcessingPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (i *ImageProcessingPostProcess) JSObject() js.Value { return i.p }

// ImageProcessingPostProcess returns a ImageProcessingPostProcess JavaScript class.
func (b *Babylon) ImageProcessingPostProcess() *ImageProcessingPostProcess {
	p := b.ctx.Get("ImageProcessingPostProcess")
	return ImageProcessingPostProcessFromJSObject(p)
}

// ImageProcessingPostProcessFromJSObject returns a wrapped ImageProcessingPostProcess JavaScript class.
func ImageProcessingPostProcessFromJSObject(p js.Value) *ImageProcessingPostProcess {
	return &ImageProcessingPostProcess{PostProcessFromJSObject(p)}
}

// NewImageProcessingPostProcessOpts contains optional parameters for NewImageProcessingPostProcess.
type NewImageProcessingPostProcessOpts struct {
	Camera *Camera

	SamplingMode *float64

	Engine *Engine

	Reusable *bool

	TextureType *float64

	ImageProcessingConfiguration *ImageProcessingConfiguration
}

// NewImageProcessingPostProcess returns a new ImageProcessingPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess
func (b *Babylon) NewImageProcessingPostProcess(name string, options float64, opts *NewImageProcessingPostProcessOpts) *ImageProcessingPostProcess {
	if opts == nil {
		opts = &NewImageProcessingPostProcessOpts{}
	}

	p := b.ctx.Get("ImageProcessingPostProcess").New(name, options, opts.Camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType, opts.ImageProcessingConfiguration.JSObject())
	return ImageProcessingPostProcessFromJSObject(p)
}

// TODO: methods
