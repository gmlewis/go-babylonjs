// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Behavior represents a babylon.js Behavior.
// Interface used to define a behavior
type Behavior struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *Behavior) JSObject() js.Value { return b.p }

// Behavior returns a Behavior JavaScript class.
func (ba *Babylon) Behavior() *Behavior {
	p := ba.ctx.Get("Behavior")
	return BehaviorFromJSObject(p, ba.ctx)
}

// BehaviorFromJSObject returns a wrapped Behavior JavaScript class.
func BehaviorFromJSObject(p js.Value, ctx js.Value) *Behavior {
	return &Behavior{p: p, ctx: ctx}
}

// BehaviorArrayToJSArray returns a JavaScript Array for the wrapped array.
func BehaviorArrayToJSArray(array []*Behavior) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Attach calls the Attach method on the Behavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.behavior#attach
func (b *Behavior) Attach(target *T) {

	args := make([]interface{}, 0, 1+0)

	if target == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, target.JSObject())
	}

	b.p.Call("attach", args...)
}

// Detach calls the Detach method on the Behavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.behavior#detach
func (b *Behavior) Detach() {

	b.p.Call("detach")
}

// Init calls the Init method on the Behavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.behavior#init
func (b *Behavior) Init() {

	b.p.Call("init")
}

// Name returns the Name property of class Behavior.
//
// https://doc.babylonjs.com/api/classes/babylon.behavior#name
func (b *Behavior) Name() string {
	retVal := b.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Behavior.
//
// https://doc.babylonjs.com/api/classes/babylon.behavior#name
func (b *Behavior) SetName(name string) *Behavior {
	b.p.Set("name", name)
	return b
}
