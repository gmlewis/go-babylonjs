// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsGravitationalFieldEvent represents a babylon.js PhysicsGravitationalFieldEvent.
// Represents a gravitational field event
type PhysicsGravitationalFieldEvent struct{}

// JSObject returns the underlying js.Value.
func (p *PhysicsGravitationalFieldEvent) JSObject() js.Value { return p.p }

// PhysicsGravitationalFieldEvent returns a PhysicsGravitationalFieldEvent JavaScript class.
func (b *Babylon) PhysicsGravitationalFieldEvent() *PhysicsGravitationalFieldEvent {
	p := b.ctx.Get("PhysicsGravitationalFieldEvent")
	return PhysicsGravitationalFieldEventFromJSObject(p)
}

// PhysicsGravitationalFieldEventFromJSObject returns a wrapped PhysicsGravitationalFieldEvent JavaScript class.
func PhysicsGravitationalFieldEventFromJSObject(p js.Value) *PhysicsGravitationalFieldEvent {
	return &PhysicsGravitationalFieldEvent{p: p}
}

// NewPhysicsGravitationalFieldEvent returns a new PhysicsGravitationalFieldEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsgravitationalfieldevent
func (b *Babylon) NewPhysicsGravitationalFieldEvent(todo parameters) *PhysicsGravitationalFieldEvent {
	p := b.ctx.Get("PhysicsGravitationalFieldEvent").New(todo)
	return PhysicsGravitationalFieldEventFromJSObject(p)
}

// TODO: methods
