// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Slider represents a babylon.js Slider.
// Class used to create slider controls
type Slider struct {
	*BaseSlider
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Slider) JSObject() js.Value { return s.p }

// Slider returns a Slider JavaScript class.
func (ba *Babylon) Slider() *Slider {
	p := ba.ctx.Get("Slider")
	return SliderFromJSObject(p, ba.ctx)
}

// SliderFromJSObject returns a wrapped Slider JavaScript class.
func SliderFromJSObject(p js.Value, ctx js.Value) *Slider {
	return &Slider{BaseSlider: BaseSliderFromJSObject(p, ctx), ctx: ctx}
}

// SliderArrayToJSArray returns a JavaScript Array for the wrapped array.
func SliderArrayToJSArray(array []*Slider) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSliderOpts contains optional parameters for NewSlider.
type NewSliderOpts struct {
	Name *string
}

// NewSlider returns a new Slider object.
//
// https://doc.babylonjs.com/api/classes/babylon.slider
func (ba *Babylon) NewSlider(opts *NewSliderOpts) *Slider {
	if opts == nil {
		opts = &NewSliderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Slider").New(args...)
	return SliderFromJSObject(p, ba.ctx)
}

// Slider_drawOpts contains optional parameters for Slider._draw.
type Slider_drawOpts struct {
	InvalidatedRectangle *Measure
}

// _draw calls the _draw method on the Slider object.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#_draw
func (s *Slider) _draw(context js.Value, opts *Slider_drawOpts) {
	if opts == nil {
		opts = &Slider_drawOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, context)

	if opts.InvalidatedRectangle == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.InvalidatedRectangle.JSObject())
	}

	s.p.Call("_draw", args...)
}

/*

// Background returns the Background property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#background
func (s *Slider) Background(background string) *Slider {
	p := ba.ctx.Get("Slider").New(background)
	return SliderFromJSObject(p, ba.ctx)
}

// SetBackground sets the Background property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#background
func (s *Slider) SetBackground(background string) *Slider {
	p := ba.ctx.Get("Slider").New(background)
	return SliderFromJSObject(p, ba.ctx)
}

// BorderColor returns the BorderColor property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#bordercolor
func (s *Slider) BorderColor(borderColor string) *Slider {
	p := ba.ctx.Get("Slider").New(borderColor)
	return SliderFromJSObject(p, ba.ctx)
}

// SetBorderColor sets the BorderColor property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#bordercolor
func (s *Slider) SetBorderColor(borderColor string) *Slider {
	p := ba.ctx.Get("Slider").New(borderColor)
	return SliderFromJSObject(p, ba.ctx)
}

// DisplayValueBar returns the DisplayValueBar property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#displayvaluebar
func (s *Slider) DisplayValueBar(displayValueBar bool) *Slider {
	p := ba.ctx.Get("Slider").New(displayValueBar)
	return SliderFromJSObject(p, ba.ctx)
}

// SetDisplayValueBar sets the DisplayValueBar property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#displayvaluebar
func (s *Slider) SetDisplayValueBar(displayValueBar bool) *Slider {
	p := ba.ctx.Get("Slider").New(displayValueBar)
	return SliderFromJSObject(p, ba.ctx)
}

// IsThumbCircle returns the IsThumbCircle property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#isthumbcircle
func (s *Slider) IsThumbCircle(isThumbCircle bool) *Slider {
	p := ba.ctx.Get("Slider").New(isThumbCircle)
	return SliderFromJSObject(p, ba.ctx)
}

// SetIsThumbCircle sets the IsThumbCircle property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#isthumbcircle
func (s *Slider) SetIsThumbCircle(isThumbCircle bool) *Slider {
	p := ba.ctx.Get("Slider").New(isThumbCircle)
	return SliderFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#name
func (s *Slider) Name(name string) *Slider {
	p := ba.ctx.Get("Slider").New(name)
	return SliderFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.slider#name
func (s *Slider) SetName(name string) *Slider {
	p := ba.ctx.Get("Slider").New(name)
	return SliderFromJSObject(p, ba.ctx)
}

*/
