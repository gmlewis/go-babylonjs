// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// QuinticEase represents a babylon.js QuinticEase.
// Easing function with a power of 5 shape (see link below).
//
// See: https://easings.net/#easeInQuint
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type QuinticEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (q *QuinticEase) JSObject() js.Value { return q.p }

// QuinticEase returns a QuinticEase JavaScript class.
func (ba *Babylon) QuinticEase() *QuinticEase {
	p := ba.ctx.Get("QuinticEase")
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// QuinticEaseFromJSObject returns a wrapped QuinticEase JavaScript class.
func QuinticEaseFromJSObject(p js.Value, ctx js.Value) *QuinticEase {
	return &QuinticEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// QuinticEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func QuinticEaseArrayToJSArray(array []*QuinticEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Ease calls the Ease method on the QuinticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#ease
func (q *QuinticEase) Ease(gradient float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := q.p.Call("ease", args...)
	return retVal.Float()
}

// GetEasingMode calls the GetEasingMode method on the QuinticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#geteasingmode
func (q *QuinticEase) GetEasingMode() float64 {

	retVal := q.p.Call("getEasingMode")
	return retVal.Float()
}

// SetEasingMode calls the SetEasingMode method on the QuinticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#seteasingmode
func (q *QuinticEase) SetEasingMode(easingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, easingMode)

	q.p.Call("setEasingMode", args...)
}

/*

// EASINGMODE_EASEIN returns the EASINGMODE_EASEIN property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easein
func (q *QuinticEase) EASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEIN)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEIN sets the EASINGMODE_EASEIN property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easein
func (q *QuinticEase) SetEASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEIN)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEINOUT returns the EASINGMODE_EASEINOUT property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easeinout
func (q *QuinticEase) EASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEINOUT)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEINOUT sets the EASINGMODE_EASEINOUT property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easeinout
func (q *QuinticEase) SetEASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEINOUT)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEOUT returns the EASINGMODE_EASEOUT property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easeout
func (q *QuinticEase) EASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEOUT)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEOUT sets the EASINGMODE_EASEOUT property of class QuinticEase.
//
// https://doc.babylonjs.com/api/classes/babylon.quinticease#easingmode_easeout
func (q *QuinticEase) SetEASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *QuinticEase {
	p := ba.ctx.Get("QuinticEase").New(EASINGMODE_EASEOUT)
	return QuinticEaseFromJSObject(p, ba.ctx)
}

*/
