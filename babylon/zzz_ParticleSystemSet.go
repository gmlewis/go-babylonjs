// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ParticleSystemSet represents a babylon.js ParticleSystemSet.
// Represents a set of particle systems working together to create a specific effect
type ParticleSystemSet struct{}

// JSObject returns the underlying js.Value.
func (p *ParticleSystemSet) JSObject() js.Value { return p.p }

// ParticleSystemSet returns a ParticleSystemSet JavaScript class.
func (b *Babylon) ParticleSystemSet() *ParticleSystemSet {
	p := b.ctx.Get("ParticleSystemSet")
	return ParticleSystemSetFromJSObject(p)
}

// ParticleSystemSetFromJSObject returns a wrapped ParticleSystemSet JavaScript class.
func ParticleSystemSetFromJSObject(p js.Value) *ParticleSystemSet {
	return &ParticleSystemSet{p: p}
}

// NewParticleSystemSet returns a new ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset
func (b *Babylon) NewParticleSystemSet(todo parameters) *ParticleSystemSet {
	p := b.ctx.Get("ParticleSystemSet").New(todo)
	return ParticleSystemSetFromJSObject(p)
}

// TODO: methods