// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// QuarticEase represents a babylon.js QuarticEase.
// Easing function with a power of 4 shape (see link below).
//
// See: https://easings.net/#easeInQuart
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type QuarticEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (q *QuarticEase) JSObject() js.Value { return q.p }

// QuarticEase returns a QuarticEase JavaScript class.
func (ba *Babylon) QuarticEase() *QuarticEase {
	p := ba.ctx.Get("QuarticEase")
	return QuarticEaseFromJSObject(p, ba.ctx)
}

// QuarticEaseFromJSObject returns a wrapped QuarticEase JavaScript class.
func QuarticEaseFromJSObject(p js.Value, ctx js.Value) *QuarticEase {
	return &QuarticEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// QuarticEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func QuarticEaseArrayToJSArray(array []*QuarticEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

 */
