// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ExponentialEase represents a babylon.js ExponentialEase.
// Easing function with an exponential shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type ExponentialEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *ExponentialEase) JSObject() js.Value { return e.p }

// ExponentialEase returns a ExponentialEase JavaScript class.
func (ba *Babylon) ExponentialEase() *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase")
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// ExponentialEaseFromJSObject returns a wrapped ExponentialEase JavaScript class.
func ExponentialEaseFromJSObject(p js.Value, ctx js.Value) *ExponentialEase {
	return &ExponentialEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// NewExponentialEaseOpts contains optional parameters for NewExponentialEase.
type NewExponentialEaseOpts struct {
	Exponent *float64
}

// NewExponentialEase returns a new ExponentialEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease
func (ba *Babylon) NewExponentialEase(opts *NewExponentialEaseOpts) *ExponentialEase {
	if opts == nil {
		opts = &NewExponentialEaseOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Exponent == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Exponent)
	}

	p := ba.ctx.Get("ExponentialEase").New(args...)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// Ease calls the Ease method on the ExponentialEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#ease
func (e *ExponentialEase) Ease(gradient float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := e.p.Call("ease", args...)
	return retVal.Float()
}

// GetEasingMode calls the GetEasingMode method on the ExponentialEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#geteasingmode
func (e *ExponentialEase) GetEasingMode() float64 {

	args := make([]interface{}, 0, 0+0)

	retVal := e.p.Call("getEasingMode", args...)
	return retVal.Float()
}

// SetEasingMode calls the SetEasingMode method on the ExponentialEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#seteasingmode
func (e *ExponentialEase) SetEasingMode(easingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, easingMode)

	e.p.Call("setEasingMode", args...)
}

/*

// EASINGMODE_EASEIN returns the EASINGMODE_EASEIN property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easein
func (e *ExponentialEase) EASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEIN)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEIN sets the EASINGMODE_EASEIN property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easein
func (e *ExponentialEase) SetEASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEIN)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEINOUT returns the EASINGMODE_EASEINOUT property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easeinout
func (e *ExponentialEase) EASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEINOUT)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEINOUT sets the EASINGMODE_EASEINOUT property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easeinout
func (e *ExponentialEase) SetEASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEINOUT)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEOUT returns the EASINGMODE_EASEOUT property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easeout
func (e *ExponentialEase) EASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEOUT)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEOUT sets the EASINGMODE_EASEOUT property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#easingmode_easeout
func (e *ExponentialEase) SetEASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(EASINGMODE_EASEOUT)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// Exponent returns the Exponent property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#exponent
func (e *ExponentialEase) Exponent(exponent float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(exponent)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

// SetExponent sets the Exponent property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#exponent
func (e *ExponentialEase) SetExponent(exponent float64) *ExponentialEase {
	p := ba.ctx.Get("ExponentialEase").New(exponent)
	return ExponentialEaseFromJSObject(p, ba.ctx)
}

*/
