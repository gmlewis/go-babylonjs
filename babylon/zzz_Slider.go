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
func (gui *GUI) Slider() *Slider {
	p := gui.ctx.Get("Slider")
	return SliderFromJSObject(p, gui.ctx)
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
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#constructor
func (gui *GUI) NewSlider(opts *NewSliderOpts) *Slider {
	if opts == nil {
		opts = &NewSliderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("Slider").New(args...)
	return SliderFromJSObject(p, gui.ctx)
}

// Slider_drawOpts contains optional parameters for Slider._draw.
type Slider_drawOpts struct {
	InvalidatedRectangle *Measure
}

// _draw calls the _draw method on the Slider object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#_draw
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

// Background returns the Background property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#background
func (s *Slider) Background() string {
	retVal := s.p.Get("background")
	return retVal.String()
}

// SetBackground sets the Background property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#background
func (s *Slider) SetBackground(background string) *Slider {
	s.p.Set("background", background)
	return s
}

// BorderColor returns the BorderColor property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#bordercolor
func (s *Slider) BorderColor() string {
	retVal := s.p.Get("borderColor")
	return retVal.String()
}

// SetBorderColor sets the BorderColor property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#bordercolor
func (s *Slider) SetBorderColor(borderColor string) *Slider {
	s.p.Set("borderColor", borderColor)
	return s
}

// DisplayValueBar returns the DisplayValueBar property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#displayvaluebar
func (s *Slider) DisplayValueBar() bool {
	retVal := s.p.Get("displayValueBar")
	return retVal.Bool()
}

// SetDisplayValueBar sets the DisplayValueBar property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#displayvaluebar
func (s *Slider) SetDisplayValueBar(displayValueBar bool) *Slider {
	s.p.Set("displayValueBar", displayValueBar)
	return s
}

// IsThumbCircle returns the IsThumbCircle property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#isthumbcircle
func (s *Slider) IsThumbCircle() bool {
	retVal := s.p.Get("isThumbCircle")
	return retVal.Bool()
}

// SetIsThumbCircle sets the IsThumbCircle property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#isthumbcircle
func (s *Slider) SetIsThumbCircle(isThumbCircle bool) *Slider {
	s.p.Set("isThumbCircle", isThumbCircle)
	return s
}

// Name returns the Name property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#name
func (s *Slider) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Slider.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.slider#name
func (s *Slider) SetName(name string) *Slider {
	s.p.Set("name", name)
	return s
}
