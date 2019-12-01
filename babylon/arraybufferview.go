package babylon

import (
	"syscall/js"
)

// ArrayBufferView represents a babylon.js ArrayBufferView.
type ArrayBufferView struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArrayBufferView) JSObject() js.Value { return a.p }

// ArrayBufferView returns a ArrayBufferView JavaScript class.
func (ba *Babylon) ArrayBufferView() *ArrayBufferView {
	p := ba.ctx.Get("ArrayBufferView")
	return ArrayBufferViewFromJSObject(p, ba.ctx)
}

// ArrayBufferViewFromJSObject returns a wrapped ArrayBufferView JavaScript class.
func ArrayBufferViewFromJSObject(p js.Value, ctx js.Value) *ArrayBufferView {
	return &ArrayBufferView{p: p, ctx: ctx}
}

// ArrayBufferViewArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArrayBufferViewArrayToJSArray(array []*ArrayBufferView) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
