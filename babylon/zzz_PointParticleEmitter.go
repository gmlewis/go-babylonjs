// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointParticleEmitter represents a babylon.js PointParticleEmitter.
// Particle emitter emitting particles from a point.
// It emits the particles randomly between 2 given directions.
type PointParticleEmitter struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PointParticleEmitter) JSObject() js.Value { return p.p }

// PointParticleEmitter returns a PointParticleEmitter JavaScript class.
func (ba *Babylon) PointParticleEmitter() *PointParticleEmitter {
	p := ba.ctx.Get("PointParticleEmitter")
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

// PointParticleEmitterFromJSObject returns a wrapped PointParticleEmitter JavaScript class.
func PointParticleEmitterFromJSObject(p js.Value, ctx js.Value) *PointParticleEmitter {
	return &PointParticleEmitter{p: p, ctx: ctx}
}

// PointParticleEmitterArrayToJSArray returns a JavaScript Array for the wrapped array.
func PointParticleEmitterArrayToJSArray(array []*PointParticleEmitter) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPointParticleEmitter returns a new PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#constructor
func (ba *Babylon) NewPointParticleEmitter() *PointParticleEmitter {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("PointParticleEmitter").New(args...)
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

// ApplyToShader calls the ApplyToShader method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#applytoshader
func (p *PointParticleEmitter) ApplyToShader(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	p.p.Call("applyToShader", args...)
}

// Clone calls the Clone method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#clone
func (p *PointParticleEmitter) Clone() *PointParticleEmitter {

	retVal := p.p.Call("clone")
	return PointParticleEmitterFromJSObject(retVal, p.ctx)
}

// GetClassName calls the GetClassName method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#getclassname
func (p *PointParticleEmitter) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// GetEffectDefines calls the GetEffectDefines method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#geteffectdefines
func (p *PointParticleEmitter) GetEffectDefines() string {

	retVal := p.p.Call("getEffectDefines")
	return retVal.String()
}

// Parse calls the Parse method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#parse
func (p *PointParticleEmitter) Parse(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	p.p.Call("parse", args...)
}

// Serialize calls the Serialize method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#serialize
func (p *PointParticleEmitter) Serialize() js.Value {

	retVal := p.p.Call("serialize")
	return retVal
}

// StartDirectionFunction calls the StartDirectionFunction method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#startdirectionfunction
func (p *PointParticleEmitter) StartDirectionFunction(worldMatrix *Matrix, directionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	if worldMatrix == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, worldMatrix.JSObject())
	}

	if directionToUpdate == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, directionToUpdate.JSObject())
	}

	if particle == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, particle.JSObject())
	}

	p.p.Call("startDirectionFunction", args...)
}

// StartPositionFunction calls the StartPositionFunction method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#startpositionfunction
func (p *PointParticleEmitter) StartPositionFunction(worldMatrix *Matrix, positionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	if worldMatrix == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, worldMatrix.JSObject())
	}

	if positionToUpdate == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, positionToUpdate.JSObject())
	}

	if particle == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, particle.JSObject())
	}

	p.p.Call("startPositionFunction", args...)
}

// Direction1 returns the Direction1 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction1
func (p *PointParticleEmitter) Direction1() *Vector3 {
	retVal := p.p.Get("direction1")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetDirection1 sets the Direction1 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction1
func (p *PointParticleEmitter) SetDirection1(direction1 *Vector3) *PointParticleEmitter {
	p.p.Set("direction1", direction1.JSObject())
	return p
}

// Direction2 returns the Direction2 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction2
func (p *PointParticleEmitter) Direction2() *Vector3 {
	retVal := p.p.Get("direction2")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetDirection2 sets the Direction2 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction2
func (p *PointParticleEmitter) SetDirection2(direction2 *Vector3) *PointParticleEmitter {
	p.p.Set("direction2", direction2.JSObject())
	return p
}
