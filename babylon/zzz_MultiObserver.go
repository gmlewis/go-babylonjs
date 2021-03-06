// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiObserver represents a babylon.js MultiObserver.
// Represent a list of observers registered to multiple Observables object.
type MultiObserver struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiObserver) JSObject() js.Value { return m.p }

// MultiObserver returns a MultiObserver JavaScript class.
func (ba *Babylon) MultiObserver() *MultiObserver {
	p := ba.ctx.Get("MultiObserver")
	return MultiObserverFromJSObject(p, ba.ctx)
}

// MultiObserverFromJSObject returns a wrapped MultiObserver JavaScript class.
func MultiObserverFromJSObject(p js.Value, ctx js.Value) *MultiObserver {
	return &MultiObserver{p: p, ctx: ctx}
}

// MultiObserverArrayToJSArray returns a JavaScript Array for the wrapped array.
func MultiObserverArrayToJSArray(array []*MultiObserver) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Dispose calls the Dispose method on the MultiObserver object.
//
// https://doc.babylonjs.com/api/classes/babylon.multiobserver#dispose
func (m *MultiObserver) Dispose() {

	m.p.Call("dispose")
}
