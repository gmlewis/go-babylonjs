// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsJoint represents a babylon.js PhysicsJoint.
// This is a holder class for the physics joint created by the physics plugin
// It holds a set of functions to control the underlying joint
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type PhysicsJoint struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsJoint) JSObject() js.Value { return p.p }

// PhysicsJoint returns a PhysicsJoint JavaScript class.
func (ba *Babylon) PhysicsJoint() *PhysicsJoint {
	p := ba.ctx.Get("PhysicsJoint")
	return PhysicsJointFromJSObject(p, ba.ctx)
}

// PhysicsJointFromJSObject returns a wrapped PhysicsJoint JavaScript class.
func PhysicsJointFromJSObject(p js.Value, ctx js.Value) *PhysicsJoint {
	return &PhysicsJoint{p: p, ctx: ctx}
}

// NewPhysicsJoint returns a new PhysicsJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsjoint
func (ba *Babylon) NewPhysicsJoint(jsType float64, jointData js.Value) *PhysicsJoint {
	p := ba.ctx.Get("PhysicsJoint").New(jsType, jointData)
	return PhysicsJointFromJSObject(p, ba.ctx)
}

// TODO: methods
