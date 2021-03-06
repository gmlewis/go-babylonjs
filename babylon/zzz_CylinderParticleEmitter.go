// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderParticleEmitter represents a babylon.js CylinderParticleEmitter.
// Particle emitter emitting particles from the inside of a cylinder.
// It emits the particles alongside the cylinder radius. The emission direction might be randomized.
type CylinderParticleEmitter struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CylinderParticleEmitter) JSObject() js.Value { return c.p }

// CylinderParticleEmitter returns a CylinderParticleEmitter JavaScript class.
func (ba *Babylon) CylinderParticleEmitter() *CylinderParticleEmitter {
	p := ba.ctx.Get("CylinderParticleEmitter")
	return CylinderParticleEmitterFromJSObject(p, ba.ctx)
}

// CylinderParticleEmitterFromJSObject returns a wrapped CylinderParticleEmitter JavaScript class.
func CylinderParticleEmitterFromJSObject(p js.Value, ctx js.Value) *CylinderParticleEmitter {
	return &CylinderParticleEmitter{p: p, ctx: ctx}
}

// CylinderParticleEmitterArrayToJSArray returns a JavaScript Array for the wrapped array.
func CylinderParticleEmitterArrayToJSArray(array []*CylinderParticleEmitter) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCylinderParticleEmitterOpts contains optional parameters for NewCylinderParticleEmitter.
type NewCylinderParticleEmitterOpts struct {
	Radius              *float64
	Height              *float64
	RadiusRange         *float64
	DirectionRandomizer *float64
}

// NewCylinderParticleEmitter returns a new CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#constructor
func (ba *Babylon) NewCylinderParticleEmitter(opts *NewCylinderParticleEmitterOpts) *CylinderParticleEmitter {
	if opts == nil {
		opts = &NewCylinderParticleEmitterOpts{}
	}

	args := make([]interface{}, 0, 0+4)

	if opts.Radius == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Radius)
	}
	if opts.Height == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Height)
	}
	if opts.RadiusRange == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RadiusRange)
	}
	if opts.DirectionRandomizer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DirectionRandomizer)
	}

	p := ba.ctx.Get("CylinderParticleEmitter").New(args...)
	return CylinderParticleEmitterFromJSObject(p, ba.ctx)
}

// ApplyToShader calls the ApplyToShader method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#applytoshader
func (c *CylinderParticleEmitter) ApplyToShader(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	c.p.Call("applyToShader", args...)
}

// Clone calls the Clone method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#clone
func (c *CylinderParticleEmitter) Clone() *CylinderParticleEmitter {

	retVal := c.p.Call("clone")
	return CylinderParticleEmitterFromJSObject(retVal, c.ctx)
}

// GetClassName calls the GetClassName method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#getclassname
func (c *CylinderParticleEmitter) GetClassName() string {

	retVal := c.p.Call("getClassName")
	return retVal.String()
}

// GetEffectDefines calls the GetEffectDefines method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#geteffectdefines
func (c *CylinderParticleEmitter) GetEffectDefines() string {

	retVal := c.p.Call("getEffectDefines")
	return retVal.String()
}

// Parse calls the Parse method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#parse
func (c *CylinderParticleEmitter) Parse(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	c.p.Call("parse", args...)
}

// Serialize calls the Serialize method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#serialize
func (c *CylinderParticleEmitter) Serialize() js.Value {

	retVal := c.p.Call("serialize")
	return retVal
}

// StartDirectionFunction calls the StartDirectionFunction method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#startdirectionfunction
func (c *CylinderParticleEmitter) StartDirectionFunction(worldMatrix *Matrix, directionToUpdate *Vector3, particle *Particle) {

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

	c.p.Call("startDirectionFunction", args...)
}

// StartPositionFunction calls the StartPositionFunction method on the CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#startpositionfunction
func (c *CylinderParticleEmitter) StartPositionFunction(worldMatrix *Matrix, positionToUpdate *Vector3, particle *Particle) {

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

	c.p.Call("startPositionFunction", args...)
}

// DirectionRandomizer returns the DirectionRandomizer property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#directionrandomizer
func (c *CylinderParticleEmitter) DirectionRandomizer() float64 {
	retVal := c.p.Get("directionRandomizer")
	return retVal.Float()
}

// SetDirectionRandomizer sets the DirectionRandomizer property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#directionrandomizer
func (c *CylinderParticleEmitter) SetDirectionRandomizer(directionRandomizer float64) *CylinderParticleEmitter {
	c.p.Set("directionRandomizer", directionRandomizer)
	return c
}

// Height returns the Height property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#height
func (c *CylinderParticleEmitter) Height() float64 {
	retVal := c.p.Get("height")
	return retVal.Float()
}

// SetHeight sets the Height property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#height
func (c *CylinderParticleEmitter) SetHeight(height float64) *CylinderParticleEmitter {
	c.p.Set("height", height)
	return c
}

// Radius returns the Radius property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#radius
func (c *CylinderParticleEmitter) Radius() float64 {
	retVal := c.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#radius
func (c *CylinderParticleEmitter) SetRadius(radius float64) *CylinderParticleEmitter {
	c.p.Set("radius", radius)
	return c
}

// RadiusRange returns the RadiusRange property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#radiusrange
func (c *CylinderParticleEmitter) RadiusRange() float64 {
	retVal := c.p.Get("radiusRange")
	return retVal.Float()
}

// SetRadiusRange sets the RadiusRange property of class CylinderParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter#radiusrange
func (c *CylinderParticleEmitter) SetRadiusRange(radiusRange float64) *CylinderParticleEmitter {
	c.p.Set("radiusRange", radiusRange)
	return c
}
