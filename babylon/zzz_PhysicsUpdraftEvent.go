// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsUpdraftEvent represents a babylon.js PhysicsUpdraftEvent.
// Represents a physics updraft event
type PhysicsUpdraftEvent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsUpdraftEvent) JSObject() js.Value { return p.p }

// PhysicsUpdraftEvent returns a PhysicsUpdraftEvent JavaScript class.
func (ba *Babylon) PhysicsUpdraftEvent() *PhysicsUpdraftEvent {
	p := ba.ctx.Get("PhysicsUpdraftEvent")
	return PhysicsUpdraftEventFromJSObject(p, ba.ctx)
}

// PhysicsUpdraftEventFromJSObject returns a wrapped PhysicsUpdraftEvent JavaScript class.
func PhysicsUpdraftEventFromJSObject(p js.Value, ctx js.Value) *PhysicsUpdraftEvent {
	return &PhysicsUpdraftEvent{p: p, ctx: ctx}
}

// PhysicsUpdraftEventArrayToJSArray returns a JavaScript Array for the wrapped array.
func PhysicsUpdraftEventArrayToJSArray(array []*PhysicsUpdraftEvent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPhysicsUpdraftEvent returns a new PhysicsUpdraftEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsupdraftevent#constructor
func (ba *Babylon) NewPhysicsUpdraftEvent(_scene *Scene, _origin *Vector3, _options *PhysicsUpdraftEventOptions) *PhysicsUpdraftEvent {

	args := make([]interface{}, 0, 3+0)

	args = append(args, _scene.JSObject())
	args = append(args, _origin.JSObject())
	args = append(args, _options.JSObject())

	p := ba.ctx.Get("PhysicsUpdraftEvent").New(args...)
	return PhysicsUpdraftEventFromJSObject(p, ba.ctx)
}

// Disable calls the Disable method on the PhysicsUpdraftEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsupdraftevent#disable
func (p *PhysicsUpdraftEvent) Disable() {

	p.p.Call("disable")
}

// PhysicsUpdraftEventDisposeOpts contains optional parameters for PhysicsUpdraftEvent.Dispose.
type PhysicsUpdraftEventDisposeOpts struct {
	Force *bool
}

// Dispose calls the Dispose method on the PhysicsUpdraftEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsupdraftevent#dispose
func (p *PhysicsUpdraftEvent) Dispose(opts *PhysicsUpdraftEventDisposeOpts) {
	if opts == nil {
		opts = &PhysicsUpdraftEventDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Force == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Force)
	}

	p.p.Call("dispose", args...)
}

// Enable calls the Enable method on the PhysicsUpdraftEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsupdraftevent#enable
func (p *PhysicsUpdraftEvent) Enable() {

	p.p.Call("enable")
}

// GetData calls the GetData method on the PhysicsUpdraftEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsupdraftevent#getdata
func (p *PhysicsUpdraftEvent) GetData() *PhysicsUpdraftEventData {

	retVal := p.p.Call("getData")
	return PhysicsUpdraftEventDataFromJSObject(retVal, p.ctx)
}
