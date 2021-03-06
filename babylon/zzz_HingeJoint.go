// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HingeJoint represents a babylon.js HingeJoint.
// This class represents a single physics Hinge-Joint
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type HingeJoint struct {
	*MotorEnabledJoint
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HingeJoint) JSObject() js.Value { return h.p }

// HingeJoint returns a HingeJoint JavaScript class.
func (ba *Babylon) HingeJoint() *HingeJoint {
	p := ba.ctx.Get("HingeJoint")
	return HingeJointFromJSObject(p, ba.ctx)
}

// HingeJointFromJSObject returns a wrapped HingeJoint JavaScript class.
func HingeJointFromJSObject(p js.Value, ctx js.Value) *HingeJoint {
	return &HingeJoint{MotorEnabledJoint: MotorEnabledJointFromJSObject(p, ctx), ctx: ctx}
}

// HingeJointArrayToJSArray returns a JavaScript Array for the wrapped array.
func HingeJointArrayToJSArray(array []*HingeJoint) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHingeJoint returns a new HingeJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hingejoint#constructor
func (ba *Babylon) NewHingeJoint(jointData js.Value) *HingeJoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, jointData)

	p := ba.ctx.Get("HingeJoint").New(args...)
	return HingeJointFromJSObject(p, ba.ctx)
}

// HingeJointSetLimitOpts contains optional parameters for HingeJoint.SetLimit.
type HingeJointSetLimitOpts struct {
	LowerLimit *float64
}

// SetLimit calls the SetLimit method on the HingeJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hingejoint#setlimit
func (h *HingeJoint) SetLimit(upperLimit float64, opts *HingeJointSetLimitOpts) {
	if opts == nil {
		opts = &HingeJointSetLimitOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, upperLimit)

	if opts.LowerLimit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LowerLimit)
	}

	h.p.Call("setLimit", args...)
}

// HingeJointSetMotorOpts contains optional parameters for HingeJoint.SetMotor.
type HingeJointSetMotorOpts struct {
	Force    *float64
	MaxForce *float64
}

// SetMotor calls the SetMotor method on the HingeJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.hingejoint#setmotor
func (h *HingeJoint) SetMotor(opts *HingeJointSetMotorOpts) {
	if opts == nil {
		opts = &HingeJointSetMotorOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Force == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Force)
	}
	if opts.MaxForce == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxForce)
	}

	h.p.Call("setMotor", args...)
}
