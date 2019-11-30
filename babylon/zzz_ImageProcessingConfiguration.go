// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageProcessingConfiguration represents a babylon.js ImageProcessingConfiguration.
// This groups together the common properties used for image processing either in direct forward pass
// or through post processing effect depending on the use of the image processing pipeline in your scene
// or not.
type ImageProcessingConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ImageProcessingConfiguration) JSObject() js.Value { return i.p }

// ImageProcessingConfiguration returns a ImageProcessingConfiguration JavaScript class.
func (ba *Babylon) ImageProcessingConfiguration() *ImageProcessingConfiguration {
	p := ba.ctx.Get("ImageProcessingConfiguration")
	return ImageProcessingConfigurationFromJSObject(p, ba.ctx)
}

// ImageProcessingConfigurationFromJSObject returns a wrapped ImageProcessingConfiguration JavaScript class.
func ImageProcessingConfigurationFromJSObject(p js.Value, ctx js.Value) *ImageProcessingConfiguration {
	return &ImageProcessingConfiguration{p: p, ctx: ctx}
}

// TODO: methods
