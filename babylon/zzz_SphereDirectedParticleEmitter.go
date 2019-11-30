// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphereDirectedParticleEmitter represents a babylon.js SphereDirectedParticleEmitter.
// Particle emitter emitting particles from the inside of a sphere.
// It emits the particles randomly between two vectors.
type SphereDirectedParticleEmitter struct {
	*SphereParticleEmitter
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SphereDirectedParticleEmitter) JSObject() js.Value { return s.p }

// SphereDirectedParticleEmitter returns a SphereDirectedParticleEmitter JavaScript class.
func (ba *Babylon) SphereDirectedParticleEmitter() *SphereDirectedParticleEmitter {
	p := ba.ctx.Get("SphereDirectedParticleEmitter")
	return SphereDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// SphereDirectedParticleEmitterFromJSObject returns a wrapped SphereDirectedParticleEmitter JavaScript class.
func SphereDirectedParticleEmitterFromJSObject(p js.Value, ctx js.Value) *SphereDirectedParticleEmitter {
	return &SphereDirectedParticleEmitter{SphereParticleEmitter: SphereParticleEmitterFromJSObject(p, ctx), ctx: ctx}
}

// NewSphereDirectedParticleEmitterOpts contains optional parameters for NewSphereDirectedParticleEmitter.
type NewSphereDirectedParticleEmitterOpts struct {
	Radius *JSFloat64

	Direction1 *Vector3

	Direction2 *Vector3
}

// NewSphereDirectedParticleEmitter returns a new SphereDirectedParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.spheredirectedparticleemitter
func (ba *Babylon) NewSphereDirectedParticleEmitter(opts *NewSphereDirectedParticleEmitterOpts) *SphereDirectedParticleEmitter {
	if opts == nil {
		opts = &NewSphereDirectedParticleEmitterOpts{}
	}

	p := ba.ctx.Get("SphereDirectedParticleEmitter").New(opts.Radius.JSObject(), opts.Direction1.JSObject(), opts.Direction2.JSObject())
	return SphereDirectedParticleEmitterFromJSObject(p, ba.ctx)
}

// TODO: methods
