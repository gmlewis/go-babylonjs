// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BaseSlider represents a babylon.js BaseSlider.
// Class used to create slider controls
type BaseSlider struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BaseSlider) JSObject() js.Value { return b.p }

// BaseSlider returns a BaseSlider JavaScript class.
func (ba *Babylon) BaseSlider() *BaseSlider {
	p := ba.ctx.Get("BaseSlider")
	return BaseSliderFromJSObject(p, ba.ctx)
}

// BaseSliderFromJSObject returns a wrapped BaseSlider JavaScript class.
func BaseSliderFromJSObject(p js.Value, ctx js.Value) *BaseSlider {
	return &BaseSlider{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// BaseSliderArrayToJSArray returns a JavaScript Array for the wrapped array.
func BaseSliderArrayToJSArray(array []*BaseSlider) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBaseSliderOpts contains optional parameters for NewBaseSlider.
type NewBaseSliderOpts struct {
	Name *string
}

// NewBaseSlider returns a new BaseSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider
func (gui *GUI) NewBaseSlider(opts *NewBaseSliderOpts) *BaseSlider {
	if opts == nil {
		opts = &NewBaseSliderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("BaseSlider").New(args...)
	return BaseSliderFromJSObject(p, gui.ctx)
}

// _onPointerDown calls the _onPointerDown method on the BaseSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#_onpointerdown
func (b *BaseSlider) _onPointerDown(target *Control, coordinates *Vector2, pointerId float64, buttonIndex float64) bool {

	args := make([]interface{}, 0, 4+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)
	args = append(args, buttonIndex)

	retVal := b.p.Call("_onPointerDown", args...)
	return retVal.Bool()
}

// _onPointerMove calls the _onPointerMove method on the BaseSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#_onpointermove
func (b *BaseSlider) _onPointerMove(target *Control, coordinates *Vector2, pointerId float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)

	b.p.Call("_onPointerMove", args...)
}

// _onPointerUp calls the _onPointerUp method on the BaseSlider object.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#_onpointerup
func (b *BaseSlider) _onPointerUp(target *Control, coordinates *Vector2, pointerId float64, buttonIndex float64, notifyClick bool) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)
	args = append(args, buttonIndex)
	args = append(args, notifyClick)

	b.p.Call("_onPointerUp", args...)
}

// BarOffset returns the BarOffset property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#baroffset
func (b *BaseSlider) BarOffset() string {
	retVal := b.p.Get("barOffset")
	return retVal.String()
}

// SetBarOffset sets the BarOffset property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#baroffset
func (b *BaseSlider) SetBarOffset(barOffset string) *BaseSlider {
	b.p.Set("barOffset", barOffset)
	return b
}

// BarOffsetInPixels returns the BarOffsetInPixels property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#baroffsetinpixels
func (b *BaseSlider) BarOffsetInPixels() float64 {
	retVal := b.p.Get("barOffsetInPixels")
	return retVal.Float()
}

// SetBarOffsetInPixels sets the BarOffsetInPixels property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#baroffsetinpixels
func (b *BaseSlider) SetBarOffsetInPixels(barOffsetInPixels float64) *BaseSlider {
	b.p.Set("barOffsetInPixels", barOffsetInPixels)
	return b
}

// DisplayThumb returns the DisplayThumb property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#displaythumb
func (b *BaseSlider) DisplayThumb() bool {
	retVal := b.p.Get("displayThumb")
	return retVal.Bool()
}

// SetDisplayThumb sets the DisplayThumb property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#displaythumb
func (b *BaseSlider) SetDisplayThumb(displayThumb bool) *BaseSlider {
	b.p.Set("displayThumb", displayThumb)
	return b
}

// IsThumbClamped returns the IsThumbClamped property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#isthumbclamped
func (b *BaseSlider) IsThumbClamped() bool {
	retVal := b.p.Get("isThumbClamped")
	return retVal.Bool()
}

// SetIsThumbClamped sets the IsThumbClamped property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#isthumbclamped
func (b *BaseSlider) SetIsThumbClamped(isThumbClamped bool) *BaseSlider {
	b.p.Set("isThumbClamped", isThumbClamped)
	return b
}

// IsVertical returns the IsVertical property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#isvertical
func (b *BaseSlider) IsVertical() bool {
	retVal := b.p.Get("isVertical")
	return retVal.Bool()
}

// SetIsVertical sets the IsVertical property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#isvertical
func (b *BaseSlider) SetIsVertical(isVertical bool) *BaseSlider {
	b.p.Set("isVertical", isVertical)
	return b
}

// Maximum returns the Maximum property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#maximum
func (b *BaseSlider) Maximum() float64 {
	retVal := b.p.Get("maximum")
	return retVal.Float()
}

// SetMaximum sets the Maximum property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#maximum
func (b *BaseSlider) SetMaximum(maximum float64) *BaseSlider {
	b.p.Set("maximum", maximum)
	return b
}

// Minimum returns the Minimum property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#minimum
func (b *BaseSlider) Minimum() float64 {
	retVal := b.p.Get("minimum")
	return retVal.Float()
}

// SetMinimum sets the Minimum property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#minimum
func (b *BaseSlider) SetMinimum(minimum float64) *BaseSlider {
	b.p.Set("minimum", minimum)
	return b
}

// Name returns the Name property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#name
func (b *BaseSlider) Name() string {
	retVal := b.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#name
func (b *BaseSlider) SetName(name string) *BaseSlider {
	b.p.Set("name", name)
	return b
}

// OnValueChangedObservable returns the OnValueChangedObservable property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#onvaluechangedobservable
func (b *BaseSlider) OnValueChangedObservable() *Observable {
	retVal := b.p.Get("onValueChangedObservable")
	return ObservableFromJSObject(retVal, b.ctx)
}

// SetOnValueChangedObservable sets the OnValueChangedObservable property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#onvaluechangedobservable
func (b *BaseSlider) SetOnValueChangedObservable(onValueChangedObservable *Observable) *BaseSlider {
	b.p.Set("onValueChangedObservable", onValueChangedObservable.JSObject())
	return b
}

// Step returns the Step property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#step
func (b *BaseSlider) Step() float64 {
	retVal := b.p.Get("step")
	return retVal.Float()
}

// SetStep sets the Step property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#step
func (b *BaseSlider) SetStep(step float64) *BaseSlider {
	b.p.Set("step", step)
	return b
}

// ThumbWidth returns the ThumbWidth property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#thumbwidth
func (b *BaseSlider) ThumbWidth() string {
	retVal := b.p.Get("thumbWidth")
	return retVal.String()
}

// SetThumbWidth sets the ThumbWidth property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#thumbwidth
func (b *BaseSlider) SetThumbWidth(thumbWidth string) *BaseSlider {
	b.p.Set("thumbWidth", thumbWidth)
	return b
}

// ThumbWidthInPixels returns the ThumbWidthInPixels property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#thumbwidthinpixels
func (b *BaseSlider) ThumbWidthInPixels() float64 {
	retVal := b.p.Get("thumbWidthInPixels")
	return retVal.Float()
}

// SetThumbWidthInPixels sets the ThumbWidthInPixels property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#thumbwidthinpixels
func (b *BaseSlider) SetThumbWidthInPixels(thumbWidthInPixels float64) *BaseSlider {
	b.p.Set("thumbWidthInPixels", thumbWidthInPixels)
	return b
}

// Value returns the Value property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#value
func (b *BaseSlider) Value() float64 {
	retVal := b.p.Get("value")
	return retVal.Float()
}

// SetValue sets the Value property of class BaseSlider.
//
// https://doc.babylonjs.com/api/classes/babylon.baseslider#value
func (b *BaseSlider) SetValue(value float64) *BaseSlider {
	b.p.Set("value", value)
	return b
}
