// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PowerEase represents a babylon.js PowerEase.
// Easing function with a power shape (see link below).
//
// See: https://easings.net/#easeInQuad
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type PowerEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PowerEase) JSObject() js.Value { return p.p }

// PowerEase returns a PowerEase JavaScript class.
func (ba *Babylon) PowerEase() *PowerEase {
	p := ba.ctx.Get("PowerEase")
	return PowerEaseFromJSObject(p, ba.ctx)
}

// PowerEaseFromJSObject returns a wrapped PowerEase JavaScript class.
func PowerEaseFromJSObject(p js.Value, ctx js.Value) *PowerEase {
	return &PowerEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// PowerEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func PowerEaseArrayToJSArray(array []*PowerEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPowerEaseOpts contains optional parameters for NewPowerEase.
type NewPowerEaseOpts struct {
	Power *float64
}

// NewPowerEase returns a new PowerEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.powerease
func (ba *Babylon) NewPowerEase(opts *NewPowerEaseOpts) *PowerEase {
	if opts == nil {
		opts = &NewPowerEaseOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Power == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Power)
	}

	p := ba.ctx.Get("PowerEase").New(args...)
	return PowerEaseFromJSObject(p, ba.ctx)
}

/*

// Power returns the Power property of class PowerEase.
//
// https://doc.babylonjs.com/api/classes/babylon.powerease#power
func (p *PowerEase) Power(power float64) *PowerEase {
	p := ba.ctx.Get("PowerEase").New(power)
	return PowerEaseFromJSObject(p, ba.ctx)
}

// SetPower sets the Power property of class PowerEase.
//
// https://doc.babylonjs.com/api/classes/babylon.powerease#power
func (p *PowerEase) SetPower(power float64) *PowerEase {
	p := ba.ctx.Get("PowerEase").New(power)
	return PowerEaseFromJSObject(p, ba.ctx)
}

*/
