// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsVortexEvent represents a babylon.js PhysicsVortexEvent.
// Represents a physics vortex event
type PhysicsVortexEvent struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PhysicsVortexEvent) JSObject() js.Value { return p.p }

// PhysicsVortexEvent returns a PhysicsVortexEvent JavaScript class.
func (b *Babylon) PhysicsVortexEvent() *PhysicsVortexEvent {
	p := b.ctx.Get("PhysicsVortexEvent")
	return PhysicsVortexEventFromJSObject(p)
}

// PhysicsVortexEventFromJSObject returns a wrapped PhysicsVortexEvent JavaScript class.
func PhysicsVortexEventFromJSObject(p js.Value) *PhysicsVortexEvent {
	return &PhysicsVortexEvent{p: p}
}

// NewPhysicsVortexEvent returns a new PhysicsVortexEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexevent
func (b *Babylon) NewPhysicsVortexEvent(_scene *Scene, _origin *Vector3, _options *PhysicsVortexEventOptions) *PhysicsVortexEvent {
	p := b.ctx.Get("PhysicsVortexEvent").New(_scene.JSObject(), _origin.JSObject(), _options.JSObject())
	return PhysicsVortexEventFromJSObject(p)
}

// TODO: methods
