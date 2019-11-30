// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphereParticleEmitter represents a babylon.js SphereParticleEmitter.
// Particle emitter emitting particles from the inside of a sphere.
// It emits the particles alongside the sphere radius. The emission direction might be randomized.
type SphereParticleEmitter struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SphereParticleEmitter) JSObject() js.Value { return s.p }

// SphereParticleEmitter returns a SphereParticleEmitter JavaScript class.
func (b *Babylon) SphereParticleEmitter() *SphereParticleEmitter {
	p := b.ctx.Get("SphereParticleEmitter")
	return SphereParticleEmitterFromJSObject(p)
}

// SphereParticleEmitterFromJSObject returns a wrapped SphereParticleEmitter JavaScript class.
func SphereParticleEmitterFromJSObject(p js.Value) *SphereParticleEmitter {
	return &SphereParticleEmitter{p: p}
}

// NewSphereParticleEmitterOpts contains optional parameters for NewSphereParticleEmitter.
type NewSphereParticleEmitterOpts struct {
	Radius *float64

	RadiusRange *float64

	DirectionRandomizer *float64
}

// NewSphereParticleEmitter returns a new SphereParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphereparticleemitter
func (b *Babylon) NewSphereParticleEmitter(opts *NewSphereParticleEmitterOpts) *SphereParticleEmitter {
	if opts == nil {
		opts = &NewSphereParticleEmitterOpts{}
	}

	p := b.ctx.Get("SphereParticleEmitter").New(opts.Radius, opts.RadiusRange, opts.DirectionRandomizer)
	return SphereParticleEmitterFromJSObject(p)
}

// TODO: methods
