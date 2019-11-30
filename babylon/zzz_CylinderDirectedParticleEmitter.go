// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderDirectedParticleEmitter represents a babylon.js CylinderDirectedParticleEmitter.
// Particle emitter emitting particles from the inside of a cylinder.
// It emits the particles randomly between two vectors.
type CylinderDirectedParticleEmitter struct{ *CylinderParticleEmitter }

// JSObject returns the underlying js.Value.
func (c *CylinderDirectedParticleEmitter) JSObject() js.Value { return c.p }

// CylinderDirectedParticleEmitter returns a CylinderDirectedParticleEmitter JavaScript class.
func (b *Babylon) CylinderDirectedParticleEmitter() *CylinderDirectedParticleEmitter {
	p := b.ctx.Get("CylinderDirectedParticleEmitter")
	return CylinderDirectedParticleEmitterFromJSObject(p)
}

// CylinderDirectedParticleEmitterFromJSObject returns a wrapped CylinderDirectedParticleEmitter JavaScript class.
func CylinderDirectedParticleEmitterFromJSObject(p js.Value) *CylinderDirectedParticleEmitter {
	return &CylinderDirectedParticleEmitter{CylinderParticleEmitterFromJSObject(p)}
}

// NewCylinderDirectedParticleEmitterOpts contains optional parameters for NewCylinderDirectedParticleEmitter.
type NewCylinderDirectedParticleEmitterOpts struct {
	Radius *float64

	Height *float64

	RadiusRange *float64

	Direction1 *Vector3

	Direction2 *Vector3
}

// NewCylinderDirectedParticleEmitter returns a new CylinderDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderdirectedparticleemitter
func (b *Babylon) NewCylinderDirectedParticleEmitter(opts *NewCylinderDirectedParticleEmitterOpts) *CylinderDirectedParticleEmitter {
	if opts == nil {
		opts = &NewCylinderDirectedParticleEmitterOpts{}
	}

	p := b.ctx.Get("CylinderDirectedParticleEmitter").New(opts.Radius, opts.Height, opts.RadiusRange, opts.Direction1.JSObject(), opts.Direction2.JSObject())
	return CylinderDirectedParticleEmitterFromJSObject(p)
}

// TODO: methods
