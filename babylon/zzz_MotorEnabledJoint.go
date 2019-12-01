// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MotorEnabledJoint represents a babylon.js MotorEnabledJoint.
// Represents a Motor-Enabled Joint
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type MotorEnabledJoint struct {
	*PhysicsJoint
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MotorEnabledJoint) JSObject() js.Value { return m.p }

// MotorEnabledJoint returns a MotorEnabledJoint JavaScript class.
func (ba *Babylon) MotorEnabledJoint() *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint")
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// MotorEnabledJointFromJSObject returns a wrapped MotorEnabledJoint JavaScript class.
func MotorEnabledJointFromJSObject(p js.Value, ctx js.Value) *MotorEnabledJoint {
	return &MotorEnabledJoint{PhysicsJoint: PhysicsJointFromJSObject(p, ctx), ctx: ctx}
}

// MotorEnabledJointArrayToJSArray returns a JavaScript Array for the wrapped array.
func MotorEnabledJointArrayToJSArray(array []*MotorEnabledJoint) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMotorEnabledJoint returns a new MotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint
func (ba *Babylon) NewMotorEnabledJoint(jsType float64, jointData js.Value) *MotorEnabledJoint {

	args := make([]interface{}, 0, 2+0)

	args = append(args, jsType)
	args = append(args, jointData)

	p := ba.ctx.Get("MotorEnabledJoint").New(args...)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// ExecuteNativeFunction calls the ExecuteNativeFunction method on the MotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#executenativefunction
func (m *MotorEnabledJoint) ExecuteNativeFunction(jsFunc func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { jsFunc(); return nil }))

	m.p.Call("executeNativeFunction", args...)
}

// MotorEnabledJointSetLimitOpts contains optional parameters for MotorEnabledJoint.SetLimit.
type MotorEnabledJointSetLimitOpts struct {
	LowerLimit *float64
}

// SetLimit calls the SetLimit method on the MotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#setlimit
func (m *MotorEnabledJoint) SetLimit(upperLimit float64, opts *MotorEnabledJointSetLimitOpts) {
	if opts == nil {
		opts = &MotorEnabledJointSetLimitOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, upperLimit)

	if opts.LowerLimit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LowerLimit)
	}

	m.p.Call("setLimit", args...)
}

// MotorEnabledJointSetMotorOpts contains optional parameters for MotorEnabledJoint.SetMotor.
type MotorEnabledJointSetMotorOpts struct {
	Force    *float64
	MaxForce *float64
}

// SetMotor calls the SetMotor method on the MotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#setmotor
func (m *MotorEnabledJoint) SetMotor(opts *MotorEnabledJointSetMotorOpts) {
	if opts == nil {
		opts = &MotorEnabledJointSetMotorOpts{}
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

	m.p.Call("setMotor", args...)
}

/*

// BallAndSocketJoint returns the BallAndSocketJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#ballandsocketjoint
func (m *MotorEnabledJoint) BallAndSocketJoint(BallAndSocketJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(BallAndSocketJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetBallAndSocketJoint sets the BallAndSocketJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#ballandsocketjoint
func (m *MotorEnabledJoint) SetBallAndSocketJoint(BallAndSocketJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(BallAndSocketJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// DistanceJoint returns the DistanceJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#distancejoint
func (m *MotorEnabledJoint) DistanceJoint(DistanceJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(DistanceJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetDistanceJoint sets the DistanceJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#distancejoint
func (m *MotorEnabledJoint) SetDistanceJoint(DistanceJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(DistanceJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// Hinge2Joint returns the Hinge2Joint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#hinge2joint
func (m *MotorEnabledJoint) Hinge2Joint(Hinge2Joint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(Hinge2Joint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetHinge2Joint sets the Hinge2Joint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#hinge2joint
func (m *MotorEnabledJoint) SetHinge2Joint(Hinge2Joint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(Hinge2Joint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// HingeJoint returns the HingeJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#hingejoint
func (m *MotorEnabledJoint) HingeJoint(HingeJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(HingeJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetHingeJoint sets the HingeJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#hingejoint
func (m *MotorEnabledJoint) SetHingeJoint(HingeJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(HingeJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// JointData returns the JointData property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#jointdata
func (m *MotorEnabledJoint) JointData(jointData js.Value) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(jointData)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetJointData sets the JointData property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#jointdata
func (m *MotorEnabledJoint) SetJointData(jointData js.Value) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(jointData)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// LockJoint returns the LockJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#lockjoint
func (m *MotorEnabledJoint) LockJoint(LockJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(LockJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetLockJoint sets the LockJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#lockjoint
func (m *MotorEnabledJoint) SetLockJoint(LockJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(LockJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// PhysicsJoint returns the PhysicsJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#physicsjoint
func (m *MotorEnabledJoint) PhysicsJoint(physicsJoint interface{}) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(physicsJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetPhysicsJoint sets the PhysicsJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#physicsjoint
func (m *MotorEnabledJoint) SetPhysicsJoint(physicsJoint interface{}) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(physicsJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// PhysicsPlugin returns the PhysicsPlugin property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#physicsplugin
func (m *MotorEnabledJoint) PhysicsPlugin(physicsPlugin js.Value) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(physicsPlugin)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetPhysicsPlugin sets the PhysicsPlugin property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#physicsplugin
func (m *MotorEnabledJoint) SetPhysicsPlugin(physicsPlugin js.Value) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(physicsPlugin)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// PointToPointJoint returns the PointToPointJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#pointtopointjoint
func (m *MotorEnabledJoint) PointToPointJoint(PointToPointJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(PointToPointJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetPointToPointJoint sets the PointToPointJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#pointtopointjoint
func (m *MotorEnabledJoint) SetPointToPointJoint(PointToPointJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(PointToPointJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// PrismaticJoint returns the PrismaticJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#prismaticjoint
func (m *MotorEnabledJoint) PrismaticJoint(PrismaticJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(PrismaticJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetPrismaticJoint sets the PrismaticJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#prismaticjoint
func (m *MotorEnabledJoint) SetPrismaticJoint(PrismaticJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(PrismaticJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SliderJoint returns the SliderJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#sliderjoint
func (m *MotorEnabledJoint) SliderJoint(SliderJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(SliderJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetSliderJoint sets the SliderJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#sliderjoint
func (m *MotorEnabledJoint) SetSliderJoint(SliderJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(SliderJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SpringJoint returns the SpringJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#springjoint
func (m *MotorEnabledJoint) SpringJoint(SpringJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(SpringJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetSpringJoint sets the SpringJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#springjoint
func (m *MotorEnabledJoint) SetSpringJoint(SpringJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(SpringJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#type
func (m *MotorEnabledJoint) Type(jsType float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(jsType)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#type
func (m *MotorEnabledJoint) SetType(jsType float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(jsType)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// UniversalJoint returns the UniversalJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#universaljoint
func (m *MotorEnabledJoint) UniversalJoint(UniversalJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(UniversalJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetUniversalJoint sets the UniversalJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#universaljoint
func (m *MotorEnabledJoint) SetUniversalJoint(UniversalJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(UniversalJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// WheelJoint returns the WheelJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#wheeljoint
func (m *MotorEnabledJoint) WheelJoint(WheelJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(WheelJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetWheelJoint sets the WheelJoint property of class MotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.motorenabledjoint#wheeljoint
func (m *MotorEnabledJoint) SetWheelJoint(WheelJoint float64) *MotorEnabledJoint {
	p := ba.ctx.Get("MotorEnabledJoint").New(WheelJoint)
	return MotorEnabledJointFromJSObject(p, ba.ctx)
}

*/
