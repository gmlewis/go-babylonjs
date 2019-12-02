// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SliderGroup represents a babylon.js SliderGroup.
// Class used to create a SliderGroup
// which contains groups of slider buttons
type SliderGroup struct {
	*SelectorGroup
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SliderGroup) JSObject() js.Value { return s.p }

// SliderGroup returns a SliderGroup JavaScript class.
func (ba *Babylon) SliderGroup() *SliderGroup {
	p := ba.ctx.Get("SliderGroup")
	return SliderGroupFromJSObject(p, ba.ctx)
}

// SliderGroupFromJSObject returns a wrapped SliderGroup JavaScript class.
func SliderGroupFromJSObject(p js.Value, ctx js.Value) *SliderGroup {
	return &SliderGroup{SelectorGroup: SelectorGroupFromJSObject(p, ctx), ctx: ctx}
}

// SliderGroupArrayToJSArray returns a JavaScript Array for the wrapped array.
func SliderGroupArrayToJSArray(array []*SliderGroup) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSliderGroup returns a new SliderGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.slidergroup
func (ba *Babylon) NewSliderGroup(name string) *SliderGroup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SliderGroup").New(args...)
	return SliderGroupFromJSObject(p, ba.ctx)
}

// SliderGroupAddSliderOpts contains optional parameters for SliderGroup.AddSlider.
type SliderGroupAddSliderOpts struct {
	Func          func()
	Unit          *string
	Min           *float64
	Max           *float64
	Value         *float64
	OnValueChange func()
}

// AddSlider calls the AddSlider method on the SliderGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.slidergroup#addslider
func (s *SliderGroup) AddSlider(label string, opts *SliderGroupAddSliderOpts) {
	if opts == nil {
		opts = &SliderGroupAddSliderOpts{}
	}

	args := make([]interface{}, 0, 1+6)

	args = append(args, label)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Func)
	}
	if opts.Unit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Unit)
	}
	if opts.Min == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Min)
	}
	if opts.Max == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Max)
	}
	if opts.Value == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Value)
	}
	if opts.OnValueChange == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnValueChange)
	}

	s.p.Call("addSlider", args...)
}

/*

 */
