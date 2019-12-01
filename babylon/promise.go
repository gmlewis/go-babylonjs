package babylon

import (
	"syscall/js"
)

// Promise represents a babylon.js Promise.
type Promise struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *Promise) JSObject() js.Value { return a.p }

// Promise returns a Promise JavaScript class.
func (ba *Babylon) Promise() *Promise {
	p := ba.ctx.Get("Promise")
	return PromiseFromJSObject(p, ba.ctx)
}

// PromiseFromJSObject returns a wrapped Promise JavaScript class.
func PromiseFromJSObject(p js.Value, ctx js.Value) *Promise {
	return &Promise{p: p, ctx: ctx}
}

// PromiseArrayToJSArray returns a JavaScript Array for the wrapped array.
func PromiseArrayToJSArray(array []*Promise) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
