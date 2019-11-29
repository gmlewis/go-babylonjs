// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GPUParticleSystem represents a babylon.js GPUParticleSystem.
// This represents a GPU particle system in Babylon
// This is the fastest particle system in Babylon as it uses the GPU to update the individual particle data

//
// See: https://www.babylonjs-playground.com/#PU4WYI#4
type GPUParticleSystem struct{ *BaseParticleSystem }

// JSObject returns the underlying js.Value.
func (g *GPUParticleSystem) JSObject() js.Value { return g.p }

// GPUParticleSystem returns a GPUParticleSystem JavaScript class.
func (b *Babylon) GPUParticleSystem() *GPUParticleSystem {
	p := b.ctx.Get("GPUParticleSystem")
	return GPUParticleSystemFromJSObject(p)
}

// GPUParticleSystemFromJSObject returns a wrapped GPUParticleSystem JavaScript class.
func GPUParticleSystemFromJSObject(p js.Value) *GPUParticleSystem {
	return &GPUParticleSystem{BaseParticleSystemFromJSObject(p)}
}

// NewGPUParticleSystem returns a new GPUParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.gpuparticlesystem
func (b *Babylon) NewGPUParticleSystem(todo parameters) *GPUParticleSystem {
	p := b.ctx.Get("GPUParticleSystem").New(todo)
	return GPUParticleSystemFromJSObject(p)
}

// TODO: methods
