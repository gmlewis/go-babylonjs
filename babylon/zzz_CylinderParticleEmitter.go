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

// NewCylinderParticleEmitterOpts contains optional parameters for NewCylinderParticleEmitter.
type NewCylinderParticleEmitterOpts struct {
	Radius *JSFloat64

	Height *JSFloat64

	RadiusRange *JSFloat64

	DirectionRandomizer *JSFloat64
}

// NewCylinderParticleEmitter returns a new CylinderParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderparticleemitter
func (ba *Babylon) NewCylinderParticleEmitter(opts *NewCylinderParticleEmitterOpts) *CylinderParticleEmitter {
	if opts == nil {
		opts = &NewCylinderParticleEmitterOpts{}
	}

	p := ba.ctx.Get("CylinderParticleEmitter").New(opts.Radius.JSObject(), opts.Height.JSObject(), opts.RadiusRange.JSObject(), opts.DirectionRandomizer.JSObject())
	return CylinderParticleEmitterFromJSObject(p, ba.ctx)
}

// TODO: methods
