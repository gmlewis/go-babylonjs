// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoxParticleEmitter represents a babylon.js BoxParticleEmitter.
// Particle emitter emitting particles from the inside of a box.
// It emits the particles randomly between 2 given directions.
type BoxParticleEmitter struct{}

// JSObject returns the underlying js.Value.
func (b *BoxParticleEmitter) JSObject() js.Value { return b.p }

// BoxParticleEmitter returns a BoxParticleEmitter JavaScript class.
func (b *Babylon) BoxParticleEmitter() *BoxParticleEmitter {
	p := b.ctx.Get("BoxParticleEmitter")
	return BoxParticleEmitterFromJSObject(p)
}

// BoxParticleEmitterFromJSObject returns a wrapped BoxParticleEmitter JavaScript class.
func BoxParticleEmitterFromJSObject(p js.Value) *BoxParticleEmitter {
	return &BoxParticleEmitter{p: p}
}

// NewBoxParticleEmitter returns a new BoxParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.boxparticleemitter
func (b *Babylon) NewBoxParticleEmitter(todo parameters) *BoxParticleEmitter {
	p := b.ctx.Get("BoxParticleEmitter").New(todo)
	return BoxParticleEmitterFromJSObject(p)
}

// TODO: methods
