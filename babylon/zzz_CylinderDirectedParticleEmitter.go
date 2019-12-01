// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderDirectedParticleEmitter represents a babylon.js CylinderDirectedParticleEmitter.
// Particle emitter emitting particles from the inside of a cylinder.
// It emits the particles randomly between two vectors.
type CylinderDirectedParticleEmitter struct {
	*CylinderParticleEmitter
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CylinderDirectedParticleEmitter) JSObject() js.Value { return c.p }

// CylinderDirectedParticleEmitter returns a CylinderDirectedParticleEmitter JavaScript class.
func (ba *Babylon) CylinderDirectedParticleEmitter() *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter")
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// CylinderDirectedParticleEmitterFromJSObject returns a wrapped CylinderDirectedParticleEmitter JavaScript class.
func CylinderDirectedParticleEmitterFromJSObject(p js.Value, ctx js.Value) *CylinderDirectedParticleEmitter {
	return &CylinderDirectedParticleEmitter{CylinderParticleEmitter: CylinderParticleEmitterFromJSObject(p, ctx), ctx: ctx}
}

// NewCylinderDirectedParticleEmitterOpts contains optional parameters for NewCylinderDirectedParticleEmitter.
type NewCylinderDirectedParticleEmitterOpts struct {
	Radius      *float64
	Height      *float64
	RadiusRange *float64
	Direction1  *Vector3
	Direction2  *Vector3
}

// NewCylinderDirectedParticleEmitter returns a new CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter
func (ba *Babylon) NewCylinderDirectedParticleEmitter(opts *NewCylinderDirectedParticleEmitterOpts) *CylinderDirectedParticleEmitter {
	if opts == nil {
		opts = &NewCylinderDirectedParticleEmitterOpts{}
	}

	args := make([]interface{}, 0, 0+5)

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
	if opts.Direction1 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Direction1.JSObject())
	}
	if opts.Direction2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Direction2.JSObject())
	}

	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(args...)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// ApplyToShader calls the ApplyToShader method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#applytoshader
func (c *CylinderDirectedParticleEmitter) ApplyToShader(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, effect.JSObject())

	c.p.Call("applyToShader", args...)
}

// Clone calls the Clone method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#clone
func (c *CylinderDirectedParticleEmitter) Clone() *CylinderDirectedParticleEmitter {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("clone", args...)
	return CylinderDirectedParticleEmitterFromJSObject(retVal, c.ctx)
}

// GetClassName calls the GetClassName method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#getclassname
func (c *CylinderDirectedParticleEmitter) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("getClassName", args...)
	return retVal.String()
}

// GetEffectDefines calls the GetEffectDefines method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#geteffectdefines
func (c *CylinderDirectedParticleEmitter) GetEffectDefines() string {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("getEffectDefines", args...)
	return retVal.String()
}

// Parse calls the Parse method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#parse
func (c *CylinderDirectedParticleEmitter) Parse(serializationObject interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, serializationObject)

	c.p.Call("parse", args...)
}

// Serialize calls the Serialize method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#serialize
func (c *CylinderDirectedParticleEmitter) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("serialize", args...)
	return retVal
}

// StartDirectionFunction calls the StartDirectionFunction method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#startdirectionfunction
func (c *CylinderDirectedParticleEmitter) StartDirectionFunction(worldMatrix *Matrix, directionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, worldMatrix.JSObject())
	args = append(args, directionToUpdate.JSObject())
	args = append(args, particle.JSObject())

	c.p.Call("startDirectionFunction", args...)
}

// StartPositionFunction calls the StartPositionFunction method on the CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#startpositionfunction
func (c *CylinderDirectedParticleEmitter) StartPositionFunction(worldMatrix *Matrix, positionToUpdate *Vector3, particle *Particle) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, worldMatrix.JSObject())
	args = append(args, positionToUpdate.JSObject())
	args = append(args, particle.JSObject())

	c.p.Call("startPositionFunction", args...)
}

/*

// Direction1 returns the Direction1 property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#direction1
func (c *CylinderDirectedParticleEmitter) Direction1(direction1 *Vector3) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(direction1.JSObject())
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetDirection1 sets the Direction1 property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#direction1
func (c *CylinderDirectedParticleEmitter) SetDirection1(direction1 *Vector3) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(direction1.JSObject())
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// Direction2 returns the Direction2 property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#direction2
func (c *CylinderDirectedParticleEmitter) Direction2(direction2 *Vector3) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(direction2.JSObject())
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetDirection2 sets the Direction2 property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#direction2
func (c *CylinderDirectedParticleEmitter) SetDirection2(direction2 *Vector3) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(direction2.JSObject())
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// DirectionRandomizer returns the DirectionRandomizer property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#directionrandomizer
func (c *CylinderDirectedParticleEmitter) DirectionRandomizer(directionRandomizer float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(directionRandomizer)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetDirectionRandomizer sets the DirectionRandomizer property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#directionrandomizer
func (c *CylinderDirectedParticleEmitter) SetDirectionRandomizer(directionRandomizer float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(directionRandomizer)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#height
func (c *CylinderDirectedParticleEmitter) Height(height float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(height)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#height
func (c *CylinderDirectedParticleEmitter) SetHeight(height float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(height)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// Radius returns the Radius property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#radius
func (c *CylinderDirectedParticleEmitter) Radius(radius float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(radius)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetRadius sets the Radius property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#radius
func (c *CylinderDirectedParticleEmitter) SetRadius(radius float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(radius)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// RadiusRange returns the RadiusRange property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#radiusrange
func (c *CylinderDirectedParticleEmitter) RadiusRange(radiusRange float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(radiusRange)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SetRadiusRange sets the RadiusRange property of class CylinderDirectedParticleEmitter.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter#radiusrange
func (c *CylinderDirectedParticleEmitter) SetRadiusRange(radiusRange float64) *CylinderDirectedParticleEmitter {
	p := ba.ctx.Get("CylinderDirectedParticleEmitter").New(radiusRange)
	return CylinderDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

*/
