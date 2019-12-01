// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CubicEase represents a babylon.js CubicEase.
// Easing function with a power of 3 shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type CubicEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CubicEase) JSObject() js.Value { return c.p }

// CubicEase returns a CubicEase JavaScript class.
func (ba *Babylon) CubicEase() *CubicEase {
	p := ba.ctx.Get("CubicEase")
	return CubicEaseFromJSObject(p, ba.ctx)
}

// CubicEaseFromJSObject returns a wrapped CubicEase JavaScript class.
func CubicEaseFromJSObject(p js.Value, ctx js.Value) *CubicEase {
	return &CubicEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// CubicEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func CubicEaseArrayToJSArray(array []*CubicEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Ease calls the Ease method on the CubicEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#ease
func (c *CubicEase) Ease(gradient float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := c.p.Call("ease", args...)
	return retVal.Float()
}

// GetEasingMode calls the GetEasingMode method on the CubicEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#geteasingmode
func (c *CubicEase) GetEasingMode() float64 {

	retVal := c.p.Call("getEasingMode")
	return retVal.Float()
}

// SetEasingMode calls the SetEasingMode method on the CubicEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#seteasingmode
func (c *CubicEase) SetEasingMode(easingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, easingMode)

	c.p.Call("setEasingMode", args...)
}

/*

// EASINGMODE_EASEIN returns the EASINGMODE_EASEIN property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easein
func (c *CubicEase) EASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEIN)
	return CubicEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEIN sets the EASINGMODE_EASEIN property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easein
func (c *CubicEase) SetEASINGMODE_EASEIN(EASINGMODE_EASEIN float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEIN)
	return CubicEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEINOUT returns the EASINGMODE_EASEINOUT property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easeinout
func (c *CubicEase) EASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEINOUT)
	return CubicEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEINOUT sets the EASINGMODE_EASEINOUT property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easeinout
func (c *CubicEase) SetEASINGMODE_EASEINOUT(EASINGMODE_EASEINOUT float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEINOUT)
	return CubicEaseFromJSObject(p, ba.ctx)
}

// EASINGMODE_EASEOUT returns the EASINGMODE_EASEOUT property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easeout
func (c *CubicEase) EASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEOUT)
	return CubicEaseFromJSObject(p, ba.ctx)
}

// SetEASINGMODE_EASEOUT sets the EASINGMODE_EASEOUT property of class CubicEase.
//
// https://doc.babylonjs.com/api/classes/babylon.cubicease#easingmode_easeout
func (c *CubicEase) SetEASINGMODE_EASEOUT(EASINGMODE_EASEOUT float64) *CubicEase {
	p := ba.ctx.Get("CubicEase").New(EASINGMODE_EASEOUT)
	return CubicEaseFromJSObject(p, ba.ctx)
}

*/
