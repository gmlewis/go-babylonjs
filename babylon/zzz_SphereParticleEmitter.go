// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphereParticleEmitter represents a babylon.js SphereParticleEmitter.
// Particle emitter emitting particles from the inside of a sphere.
// It emits the particles alongside the sphere radius. The emission direction might be randomized.
type SphereParticleEmitter struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SphereParticleEmitter) JSObject() js.Value { return s.p }

// SphereParticleEmitter returns a SphereParticleEmitter JavaScript class.
func (ba *Babylon) SphereParticleEmitter() *SphereParticleEmitter {
	p := ba.ctx.Get("SphereParticleEmitter")
	return SphereParticleEmitterFromJSObject(p, ba.ctx)
}

// SphereParticleEmitterFromJSObject returns a wrapped SphereParticleEmitter JavaScript class.
func SphereParticleEmitterFromJSObject(p js.Value, ctx js.Value) *SphereParticleEmitter {
	return &SphereParticleEmitter{p: p, ctx: ctx}
}

// NewSphereParticleEmitterOpts contains optional parameters for NewSphereParticleEmitter.
type NewSphereParticleEmitterOpts struct {
	Radius *JSFloat64

	RadiusRange *JSFloat64

	DirectionRandomizer *JSFloat64
}

// NewSphereParticleEmitter returns a new SphereParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphereparticleemitter
func (ba *Babylon) NewSphereParticleEmitter(opts *NewSphereParticleEmitterOpts) *SphereParticleEmitter {
	if opts == nil {
		opts = &NewSphereParticleEmitterOpts{}
	}

	p := ba.ctx.Get("SphereParticleEmitter").New(opts.Radius.JSObject(), opts.RadiusRange.JSObject(), opts.DirectionRandomizer.JSObject())
	return SphereParticleEmitterFromJSObject(p, ba.ctx)
}

// TODO: methods
