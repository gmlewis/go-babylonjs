// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DistanceJoint represents a babylon.js DistanceJoint.
// A class representing a physics distance joint

//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type DistanceJoint struct{ *PhysicsJoint }

// JSObject returns the underlying js.Value.
func (d *DistanceJoint) JSObject() js.Value { return d.p }

// DistanceJoint returns a DistanceJoint JavaScript class.
func (b *Babylon) DistanceJoint() *DistanceJoint {
	p := b.ctx.Get("DistanceJoint")
	return DistanceJointFromJSObject(p)
}

// DistanceJointFromJSObject returns a wrapped DistanceJoint JavaScript class.
func DistanceJointFromJSObject(p js.Value) *DistanceJoint {
	return &DistanceJoint{PhysicsJointFromJSObject(p)}
}

// NewDistanceJoint returns a new DistanceJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.distancejoint
func (b *Babylon) NewDistanceJoint(todo parameters) *DistanceJoint {
	p := b.ctx.Get("DistanceJoint").New(todo)
	return DistanceJointFromJSObject(p)
}

// TODO: methods
