// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CircleEase represents a babylon.js CircleEase.
// Easing function with a circle shape (see link below).
//
// See: https://easings.net/#easeInCirc
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type CircleEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CircleEase) JSObject() js.Value { return c.p }

// CircleEase returns a CircleEase JavaScript class.
func (ba *Babylon) CircleEase() *CircleEase {
	p := ba.ctx.Get("CircleEase")
	return CircleEaseFromJSObject(p, ba.ctx)
}

// CircleEaseFromJSObject returns a wrapped CircleEase JavaScript class.
func CircleEaseFromJSObject(p js.Value, ctx js.Value) *CircleEase {
	return &CircleEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// CircleEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func CircleEaseArrayToJSArray(array []*CircleEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
