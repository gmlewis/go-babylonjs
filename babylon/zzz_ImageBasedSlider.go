// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageBasedSlider represents a babylon.js ImageBasedSlider.
// Class used to create slider controls based on images
type ImageBasedSlider struct{ *BaseSlider }

// JSObject returns the underlying js.Value.
func (i *ImageBasedSlider) JSObject() js.Value { return i.p }

// ImageBasedSlider returns a ImageBasedSlider JavaScript class.
func (b *Babylon) ImageBasedSlider() *ImageBasedSlider {
	p := b.ctx.Get("ImageBasedSlider")
	return ImageBasedSliderFromJSObject(p)
}

// ImageBasedSliderFromJSObject returns a wrapped ImageBasedSlider JavaScript class.
func ImageBasedSliderFromJSObject(p js.Value) *ImageBasedSlider {
	return &ImageBasedSlider{BaseSliderFromJSObject(p)}
}

// NewImageBasedSlider returns a new ImageBasedSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.imagebasedslider
func (b *Babylon) NewImageBasedSlider(todo parameters) *ImageBasedSlider {
	p := b.ctx.Get("ImageBasedSlider").New(todo)
	return ImageBasedSliderFromJSObject(p)
}

// TODO: methods
