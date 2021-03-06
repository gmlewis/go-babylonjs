// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ExponentialEase represents a babylon.js ExponentialEase.
// Easing function with an exponential shape (see link below).
//
// See: https://easings.net/#easeInExpo
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

// ExponentialEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func ExponentialEaseArrayToJSArray(array []*ExponentialEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewExponentialEaseOpts contains optional parameters for NewExponentialEase.
type NewExponentialEaseOpts struct {
	Exponent *float64
}

// NewExponentialEase returns a new ExponentialEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#constructor
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

// Exponent returns the Exponent property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#exponent
func (e *ExponentialEase) Exponent() float64 {
	retVal := e.p.Get("exponent")
	return retVal.Float()
}

// SetExponent sets the Exponent property of class ExponentialEase.
//
// https://doc.babylonjs.com/api/classes/babylon.exponentialease#exponent
func (e *ExponentialEase) SetExponent(exponent float64) *ExponentialEase {
	e.p.Set("exponent", exponent)
	return e
}
