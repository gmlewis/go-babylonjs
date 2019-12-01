// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsRaycastResult represents a babylon.js PhysicsRaycastResult.
// Holds the data for the raycast result
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type PhysicsRaycastResult struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsRaycastResult) JSObject() js.Value { return p.p }

// PhysicsRaycastResult returns a PhysicsRaycastResult JavaScript class.
func (ba *Babylon) PhysicsRaycastResult() *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult")
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// PhysicsRaycastResultFromJSObject returns a wrapped PhysicsRaycastResult JavaScript class.
func PhysicsRaycastResultFromJSObject(p js.Value, ctx js.Value) *PhysicsRaycastResult {
	return &PhysicsRaycastResult{p: p, ctx: ctx}
}

// PhysicsRaycastResultArrayToJSArray returns a JavaScript Array for the wrapped array.
func PhysicsRaycastResultArrayToJSArray(array []*PhysicsRaycastResult) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CalculateHitDistance calls the CalculateHitDistance method on the PhysicsRaycastResult object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#calculatehitdistance
func (p *PhysicsRaycastResult) CalculateHitDistance() {

	p.p.Call("calculateHitDistance")
}

// PhysicsRaycastResultResetOpts contains optional parameters for PhysicsRaycastResult.Reset.
type PhysicsRaycastResultResetOpts struct {
	From *Vector3
	To   *Vector3
}

// Reset calls the Reset method on the PhysicsRaycastResult object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#reset
func (p *PhysicsRaycastResult) Reset(opts *PhysicsRaycastResultResetOpts) {
	if opts == nil {
		opts = &PhysicsRaycastResultResetOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.From == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.From.JSObject())
	}
	if opts.To == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.To.JSObject())
	}

	p.p.Call("reset", args...)
}

// SetHitData calls the SetHitData method on the PhysicsRaycastResult object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#sethitdata
func (p *PhysicsRaycastResult) SetHitData(hitNormalWorld *IXYZ, hitPointWorld *IXYZ) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, hitNormalWorld.JSObject())
	args = append(args, hitPointWorld.JSObject())

	p.p.Call("setHitData", args...)
}

// SetHitDistance calls the SetHitDistance method on the PhysicsRaycastResult object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#sethitdistance
func (p *PhysicsRaycastResult) SetHitDistance(distance float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	p.p.Call("setHitDistance", args...)
}

/*

// HasHit returns the HasHit property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hashit
func (p *PhysicsRaycastResult) HasHit(hasHit bool) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hasHit)
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetHasHit sets the HasHit property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hashit
func (p *PhysicsRaycastResult) SetHasHit(hasHit bool) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hasHit)
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// HitDistance returns the HitDistance property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitdistance
func (p *PhysicsRaycastResult) HitDistance(hitDistance float64) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitDistance)
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetHitDistance sets the HitDistance property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitdistance
func (p *PhysicsRaycastResult) SetHitDistance(hitDistance float64) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitDistance)
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// HitNormalWorld returns the HitNormalWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitnormalworld
func (p *PhysicsRaycastResult) HitNormalWorld(hitNormalWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitNormalWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetHitNormalWorld sets the HitNormalWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitnormalworld
func (p *PhysicsRaycastResult) SetHitNormalWorld(hitNormalWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitNormalWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// HitPointWorld returns the HitPointWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitpointworld
func (p *PhysicsRaycastResult) HitPointWorld(hitPointWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitPointWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetHitPointWorld sets the HitPointWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#hitpointworld
func (p *PhysicsRaycastResult) SetHitPointWorld(hitPointWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(hitPointWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// RayFromWorld returns the RayFromWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#rayfromworld
func (p *PhysicsRaycastResult) RayFromWorld(rayFromWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(rayFromWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetRayFromWorld sets the RayFromWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#rayfromworld
func (p *PhysicsRaycastResult) SetRayFromWorld(rayFromWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(rayFromWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// RayToWorld returns the RayToWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#raytoworld
func (p *PhysicsRaycastResult) RayToWorld(rayToWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(rayToWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

// SetRayToWorld sets the RayToWorld property of class PhysicsRaycastResult.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsraycastresult#raytoworld
func (p *PhysicsRaycastResult) SetRayToWorld(rayToWorld *Vector3) *PhysicsRaycastResult {
	p := ba.ctx.Get("PhysicsRaycastResult").New(rayToWorld.JSObject())
	return PhysicsRaycastResultFromJSObject(p, ba.ctx)
}

*/
