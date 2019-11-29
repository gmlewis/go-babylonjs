// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphereParticleEmitter represents a babylon.js SphereParticleEmitter.
// Particle emitter emitting particles from the inside of a sphere.
// It emits the particles alongside the sphere radius. The emission direction might be randomized.
type SphereParticleEmitter struct{}

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

// NewSphereParticleEmitter returns a new SphereParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphereparticleemitter
func (b *Babylon) NewSphereParticleEmitter(todo parameters) *SphereParticleEmitter {
	p := b.ctx.Get("SphereParticleEmitter").New(todo)
	return SphereParticleEmitterFromJSObject(p)
}

// TODO: methods