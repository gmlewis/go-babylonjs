// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EasingFunction represents a babylon.js EasingFunction.
// Base class used for every default easing function.
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type EasingFunction struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EasingFunction) JSObject() js.Value { return e.p }

// EasingFunction returns a EasingFunction JavaScript class.
func (ba *Babylon) EasingFunction() *EasingFunction {
	p := ba.ctx.Get("EasingFunction")
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// EasingFunctionFromJSObject returns a wrapped EasingFunction JavaScript class.
func EasingFunctionFromJSObject(p js.Value, ctx js.Value) *EasingFunction {
	return &EasingFunction{p: p, ctx: ctx}
}

// EasingFunctionArrayToJSArray returns a JavaScript Array for the wrapped array.
func EasingFunctionArrayToJSArray(array []*EasingFunction) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Ease calls the Ease method on the EasingFunction object.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#ease
func (e *EasingFunction) Ease(gradient float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := e.p.Call("ease", args...)
	return retVal.Float()
}

// GetEasingMode calls the GetEasingMode method on the EasingFunction object.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#geteasingmode
func (e *EasingFunction) GetEasingMode() float64 {

	retVal := e.p.Call("getEasingMode")
	return retVal.Float()
}

// SetEasingMode calls the SetEasingMode method on the EasingFunction object.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#seteasingmode
func (e *EasingFunction) SetEasingMode(easingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, easingMode)

	e.p.Call("setEasingMode", args...)
}

/*

// EASINGMODE_EASEIN returns the EASINGMODE_EASEIN property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easein
func (e *EasingFunction) EASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEIN)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEIN sets the EASINGMODE_EASEIN property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easein
func (e *EasingFunction) SetEASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEIN)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEINOUT returns the EASINGMODE_EASEINOUT property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easeinout
func (e *EasingFunction) EASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEINOUT)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEINOUT sets the EASINGMODE_EASEINOUT property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easeinout
func (e *EasingFunction) SetEASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEINOUT)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEOUT returns the EASINGMODE_EASEOUT property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easeout
func (e *EasingFunction) EASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEOUT)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEOUT sets the EASINGMODE_EASEOUT property of class EasingFunction.
//
// https://doc.babylonjs.com/api/classes/babylon.easingfunction#easingmode_easeout
func (e *EasingFunction) SetEASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *EasingFunction {
	p := ba.ctx.Get("EasingFunction").New(EASINGMODE_EASEOUT)
	return EasingFunctionFromJSObject(p, ba.ctx)
}

*/
