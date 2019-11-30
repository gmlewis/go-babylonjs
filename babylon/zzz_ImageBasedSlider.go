// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageBasedSlider represents a babylon.js ImageBasedSlider.
// Class used to create slider controls based on images
type ImageBasedSlider struct {
	*BaseSlider
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ImageBasedSlider) JSObject() js.Value { return i.p }

// ImageBasedSlider returns a ImageBasedSlider JavaScript class.
func (ba *Babylon) ImageBasedSlider() *ImageBasedSlider {
	p := ba.ctx.Get("ImageBasedSlider")
	return ImageBasedSliderFromJSObject(p, ba.ctx)
}

// ImageBasedSliderFromJSObject returns a wrapped ImageBasedSlider JavaScript class.
func ImageBasedSliderFromJSObject(p js.Value, ctx js.Value) *ImageBasedSlider {
	return &ImageBasedSlider{BaseSlider: BaseSliderFromJSObject(p, ctx), ctx: ctx}
}

// NewImageBasedSliderOpts contains optional parameters for NewImageBasedSlider.
type NewImageBasedSliderOpts struct {
	Name *JSString
}

// NewImageBasedSlider returns a new ImageBasedSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.imagebasedslider
func (ba *Babylon) NewImageBasedSlider(opts *NewImageBasedSliderOpts) *ImageBasedSlider {
	if opts == nil {
		opts = &NewImageBasedSliderOpts{}
	}

	p := ba.ctx.Get("ImageBasedSlider").New(opts.Name.JSObject())
	return ImageBasedSliderFromJSObject(p, ba.ctx)
}

// TODO: methods
