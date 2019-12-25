// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiPointerScaleBehavior represents a babylon.js MultiPointerScaleBehavior.
// A behavior that when attached to a mesh will allow the mesh to be scaled
type MultiPointerScaleBehavior struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiPointerScaleBehavior) JSObject() js.Value { return m.p }

// MultiPointerScaleBehavior returns a MultiPointerScaleBehavior JavaScript class.
func (ba *Babylon) MultiPointerScaleBehavior() *MultiPointerScaleBehavior {
	p := ba.ctx.Get("MultiPointerScaleBehavior")
	return MultiPointerScaleBehaviorFromJSObject(p, ba.ctx)
}

// MultiPointerScaleBehaviorFromJSObject returns a wrapped MultiPointerScaleBehavior JavaScript class.
func MultiPointerScaleBehaviorFromJSObject(p js.Value, ctx js.Value) *MultiPointerScaleBehavior {
	return &MultiPointerScaleBehavior{p: p, ctx: ctx}
}

// MultiPointerScaleBehaviorArrayToJSArray returns a JavaScript Array for the wrapped array.
func MultiPointerScaleBehaviorArrayToJSArray(array []*MultiPointerScaleBehavior) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMultiPointerScaleBehavior returns a new MultiPointerScaleBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#constructor
func (ba *Babylon) NewMultiPointerScaleBehavior() *MultiPointerScaleBehavior {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("MultiPointerScaleBehavior").New(args...)
	return MultiPointerScaleBehaviorFromJSObject(p, ba.ctx)
}

// Attach calls the Attach method on the MultiPointerScaleBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#attach
func (m *MultiPointerScaleBehavior) Attach(ownerNode *Mesh) {

	args := make([]interface{}, 0, 1+0)

	if ownerNode == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ownerNode.JSObject())
	}

	m.p.Call("attach", args...)
}

// Detach calls the Detach method on the MultiPointerScaleBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#detach
func (m *MultiPointerScaleBehavior) Detach() {

	m.p.Call("detach")
}

// Init calls the Init method on the MultiPointerScaleBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#init
func (m *MultiPointerScaleBehavior) Init() {

	m.p.Call("init")
}

// Name returns the Name property of class MultiPointerScaleBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#name
func (m *MultiPointerScaleBehavior) Name() string {
	retVal := m.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class MultiPointerScaleBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.multipointerscalebehavior#name
func (m *MultiPointerScaleBehavior) SetName(name string) *MultiPointerScaleBehavior {
	m.p.Set("name", name)
	return m
}
