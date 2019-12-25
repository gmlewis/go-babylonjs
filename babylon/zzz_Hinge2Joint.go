// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Hinge2Joint represents a babylon.js Hinge2Joint.
// This class represents a dual hinge physics joint (same as wheel joint)
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type Hinge2Joint struct {
	*MotorEnabledJoint
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *Hinge2Joint) JSObject() js.Value { return h.p }

// Hinge2Joint returns a Hinge2Joint JavaScript class.
func (ba *Babylon) Hinge2Joint() *Hinge2Joint {
	p := ba.ctx.Get("Hinge2Joint")
	return Hinge2JointFromJSObject(p, ba.ctx)
}

// Hinge2JointFromJSObject returns a wrapped Hinge2Joint JavaScript class.
func Hinge2JointFromJSObject(p js.Value, ctx js.Value) *Hinge2Joint {
	return &Hinge2Joint{MotorEnabledJoint: MotorEnabledJointFromJSObject(p, ctx), ctx: ctx}
}

// Hinge2JointArrayToJSArray returns a JavaScript Array for the wrapped array.
func Hinge2JointArrayToJSArray(array []*Hinge2Joint) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHinge2Joint returns a new Hinge2Joint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hinge2joint#constructor
func (ba *Babylon) NewHinge2Joint(jointData js.Value) *Hinge2Joint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, jointData)

	p := ba.ctx.Get("Hinge2Joint").New(args...)
	return Hinge2JointFromJSObject(p, ba.ctx)
}

// Hinge2JointSetLimitOpts contains optional parameters for Hinge2Joint.SetLimit.
type Hinge2JointSetLimitOpts struct {
	LowerLimit *float64
	MotorIndex *float64
}

// SetLimit calls the SetLimit method on the Hinge2Joint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hinge2joint#setlimit
func (h *Hinge2Joint) SetLimit(upperLimit float64, opts *Hinge2JointSetLimitOpts) {
	if opts == nil {
		opts = &Hinge2JointSetLimitOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, upperLimit)

	if opts.LowerLimit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LowerLimit)
	}
	if opts.MotorIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MotorIndex)
	}

	h.p.Call("setLimit", args...)
}

// Hinge2JointSetMotorOpts contains optional parameters for Hinge2Joint.SetMotor.
type Hinge2JointSetMotorOpts struct {
	TargetSpeed *float64
	MaxForce    *float64
	MotorIndex  *float64
}

// SetMotor calls the SetMotor method on the Hinge2Joint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hinge2joint#setmotor
func (h *Hinge2Joint) SetMotor(opts *Hinge2JointSetMotorOpts) {
	if opts == nil {
		opts = &Hinge2JointSetMotorOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.TargetSpeed == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TargetSpeed)
	}
	if opts.MaxForce == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxForce)
	}
	if opts.MotorIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MotorIndex)
	}

	h.p.Call("setMotor", args...)
}
