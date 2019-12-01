package babylon

import (
	"syscall/js"
)

// T represents a babylon.js T.
type T struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *T) JSObject() js.Value { return a.p }

// T returns a T JavaScript class.
func (ba *Babylon) T() *T {
	p := ba.ctx.Get("T")
	return TFromJSObject(p, ba.ctx)
}

// TFromJSObject returns a wrapped T JavaScript class.
func TFromJSObject(p js.Value, ctx js.Value) *T {
	return &T{p: p, ctx: ctx}
}

// TArrayToJSArray returns a JavaScript Array for the wrapped array.
func TArrayToJSArray(array []*T) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
