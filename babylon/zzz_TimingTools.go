// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TimingTools represents a babylon.js TimingTools.
// Class used to provide helper for timing
type TimingTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TimingTools) JSObject() js.Value { return t.p }

// TimingTools returns a TimingTools JavaScript class.
func (ba *Babylon) TimingTools() *TimingTools {
	p := ba.ctx.Get("TimingTools")
	return TimingToolsFromJSObject(p, ba.ctx)
}

// TimingToolsFromJSObject returns a wrapped TimingTools JavaScript class.
func TimingToolsFromJSObject(p js.Value, ctx js.Value) *TimingTools {
	return &TimingTools{p: p, ctx: ctx}
}

// TimingToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func TimingToolsArrayToJSArray(array []*TimingTools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// SetImmediate calls the SetImmediate method on the TimingTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.timingtools#setimmediate
func (t *TimingTools) SetImmediate(action JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(action))

	t.p.Call("SetImmediate", args...)
}
