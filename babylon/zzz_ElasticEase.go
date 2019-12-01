// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ElasticEase represents a babylon.js ElasticEase.
// Easing function with an elastic shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type ElasticEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *ElasticEase) JSObject() js.Value { return e.p }

// ElasticEase returns a ElasticEase JavaScript class.
func (ba *Babylon) ElasticEase() *ElasticEase {
	p := ba.ctx.Get("ElasticEase")
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// ElasticEaseFromJSObject returns a wrapped ElasticEase JavaScript class.
func ElasticEaseFromJSObject(p js.Value, ctx js.Value) *ElasticEase {
	return &ElasticEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// ElasticEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func ElasticEaseArrayToJSArray(array []*ElasticEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewElasticEaseOpts contains optional parameters for NewElasticEase.
type NewElasticEaseOpts struct {
	Oscillations *float64
	Springiness  *float64
}

// NewElasticEase returns a new ElasticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease
func (ba *Babylon) NewElasticEase(opts *NewElasticEaseOpts) *ElasticEase {
	if opts == nil {
		opts = &NewElasticEaseOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Oscillations == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Oscillations)
	}
	if opts.Springiness == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Springiness)
	}

	p := ba.ctx.Get("ElasticEase").New(args...)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// Ease calls the Ease method on the ElasticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#ease
func (e *ElasticEase) Ease(gradient float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := e.p.Call("ease", args...)
	return retVal.Float()
}

// GetEasingMode calls the GetEasingMode method on the ElasticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#geteasingmode
func (e *ElasticEase) GetEasingMode() float64 {

	retVal := e.p.Call("getEasingMode")
	return retVal.Float()
}

// SetEasingMode calls the SetEasingMode method on the ElasticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#seteasingmode
func (e *ElasticEase) SetEasingMode(easingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, easingMode)

	e.p.Call("setEasingMode", args...)
}

/*

// EASINGMODE_EASEIN returns the EASINGMODE_EASEIN property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easein
func (e *ElasticEase) EASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEIN)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEIN sets the EASINGMODE_EASEIN property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easein
func (e *ElasticEase) SetEASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEIN)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEINOUT returns the EASINGMODE_EASEINOUT property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easeinout
func (e *ElasticEase) EASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEINOUT)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEINOUT sets the EASINGMODE_EASEINOUT property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easeinout
func (e *ElasticEase) SetEASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEINOUT)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEOUT returns the EASINGMODE_EASEOUT property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easeout
func (e *ElasticEase) EASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEOUT)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEOUT sets the EASINGMODE_EASEOUT property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#easingmode_easeout
func (e *ElasticEase) SetEASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(EASINGMODE_EASEOUT)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// Oscillations returns the Oscillations property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#oscillations
func (e *ElasticEase) Oscillations(oscillations float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(oscillations)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// SetOscillations sets the Oscillations property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#oscillations
func (e *ElasticEase) SetOscillations(oscillations float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(oscillations)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// Springiness returns the Springiness property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#springiness
func (e *ElasticEase) Springiness(springiness float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(springiness)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

// SetSpringiness sets the Springiness property of class ElasticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease#springiness
func (e *ElasticEase) SetSpringiness(springiness float64) *ElasticEase {
	p := ba.ctx.Get("ElasticEase").New(springiness)
	return ElasticEaseFromJSObject(p, ba.ctx)
}

*/
