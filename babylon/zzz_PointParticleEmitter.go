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

// NewPointParticleEmitter returns a new PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter
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

	args = append(args, effect.JSObject())

	p.p.Call("applyToShader", args...)
}

// Clone calls the Clone method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#clone
func (p *PointParticleEmitter) Clone() *PointParticleEmitter {

	args := make([]interface{}, 0, 0+0)

	retVal := p.p.Call("clone", args...)
	return PointParticleEmitterFromJSObject(retVal, p.ctx)
}

// GetClassName calls the GetClassName method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#getclassname
func (p *PointParticleEmitter) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := p.p.Call("getClassName", args...)
	return retVal.String()
}

// GetEffectDefines calls the GetEffectDefines method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#geteffectdefines
func (p *PointParticleEmitter) GetEffectDefines() string {

	args := make([]interface{}, 0, 0+0)

	retVal := p.p.Call("getEffectDefines", args...)
	return retVal.String()
}

// Parse calls the Parse method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#parse
func (p *PointParticleEmitter) Parse(serializationObject interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, serializationObject)

	p.p.Call("parse", args...)
}

// Serialize calls the Serialize method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#serialize
func (p *PointParticleEmitter) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := p.p.Call("serialize", args...)
	return retVal
}

// StartDirectionFunction calls the StartDirectionFunction method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#startdirectionfunction
func (p *PointParticleEmitter) StartDirectionFunction(worldMatrix *Matrix, directionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, worldMatrix.JSObject())
	args = append(args, directionToUpdate.JSObject())
	args = append(args, particle.JSObject())

	p.p.Call("startDirectionFunction", args...)
}

// StartPositionFunction calls the StartPositionFunction method on the PointParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#startpositionfunction
func (p *PointParticleEmitter) StartPositionFunction(worldMatrix *Matrix, positionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, worldMatrix.JSObject())
	args = append(args, positionToUpdate.JSObject())
	args = append(args, particle.JSObject())

	p.p.Call("startPositionFunction", args...)
}

/*

// Direction1 returns the Direction1 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction1
func (p *PointParticleEmitter) Direction1(direction1 *Vector3) *PointParticleEmitter {
	p := ba.ctx.Get("PointParticleEmitter").New(direction1.JSObject())
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

// SetDirection1 sets the Direction1 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction1
func (p *PointParticleEmitter) SetDirection1(direction1 *Vector3) *PointParticleEmitter {
	p := ba.ctx.Get("PointParticleEmitter").New(direction1.JSObject())
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

// Direction2 returns the Direction2 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction2
func (p *PointParticleEmitter) Direction2(direction2 *Vector3) *PointParticleEmitter {
	p := ba.ctx.Get("PointParticleEmitter").New(direction2.JSObject())
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

// SetDirection2 sets the Direction2 property of class PointParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.pointparticleemitter#direction2
func (p *PointParticleEmitter) SetDirection2(direction2 *Vector3) *PointParticleEmitter {
	p := ba.ctx.Get("PointParticleEmitter").New(direction2.JSObject())
	return PointParticleEmitterFromJSObject(p, ba.ctx)
}

*/
