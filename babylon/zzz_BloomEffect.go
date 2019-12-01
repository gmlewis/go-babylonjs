// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BloomEffect represents a babylon.js BloomEffect.
// The bloom effect spreads bright areas of an image to simulate artifacts seen in cameras
type BloomEffect struct {
	*PostProcessRenderEffect
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BloomEffect) JSObject() js.Value { return b.p }

// BloomEffect returns a BloomEffect JavaScript class.
func (ba *Babylon) BloomEffect() *BloomEffect {
	p := ba.ctx.Get("BloomEffect")
	return BloomEffectFromJSObject(p, ba.ctx)
}

// BloomEffectFromJSObject returns a wrapped BloomEffect JavaScript class.
func BloomEffectFromJSObject(p js.Value, ctx js.Value) *BloomEffect {
	return &BloomEffect{PostProcessRenderEffect: PostProcessRenderEffectFromJSObject(p, ctx), ctx: ctx}
}

// NewBloomEffectOpts contains optional parameters for NewBloomEffect.
type NewBloomEffectOpts struct {
	PipelineTextureType *float64
	BlockCompilation    *bool
}

// NewBloomEffect returns a new BloomEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect
func (ba *Babylon) NewBloomEffect(scene *Scene, bloomScale float64, bloomWeight float64, bloomKernel float64, opts *NewBloomEffectOpts) *BloomEffect {
	if opts == nil {
		opts = &NewBloomEffectOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, scene.JSObject())
	args = append(args, bloomScale)
	args = append(args, bloomWeight)
	args = append(args, bloomKernel)

	if opts.PipelineTextureType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PipelineTextureType)
	}
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("BloomEffect").New(args...)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// DisposeEffects calls the DisposeEffects method on the BloomEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#disposeeffects
func (b *BloomEffect) DisposeEffects(camera *Camera) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, camera.JSObject())

	b.p.Call("disposeEffects", args...)
}

// BloomEffectGetPostProcessesOpts contains optional parameters for BloomEffect.GetPostProcesses.
type BloomEffectGetPostProcessesOpts struct {
	Camera *Camera
}

// GetPostProcesses calls the GetPostProcesses method on the BloomEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#getpostprocesses
func (b *BloomEffect) GetPostProcesses(opts *BloomEffectGetPostProcessesOpts) []js.Value {
	if opts == nil {
		opts = &BloomEffectGetPostProcessesOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	retVal := b.p.Call("getPostProcesses", args...)
	return retVal
}

/*

// IsSupported returns the IsSupported property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#issupported
func (b *BloomEffect) IsSupported(isSupported bool) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(isSupported)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#issupported
func (b *BloomEffect) SetIsSupported(isSupported bool) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(isSupported)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// Kernel returns the Kernel property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#kernel
func (b *BloomEffect) Kernel(kernel float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(kernel)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// SetKernel sets the Kernel property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#kernel
func (b *BloomEffect) SetKernel(kernel float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(kernel)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// Threshold returns the Threshold property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#threshold
func (b *BloomEffect) Threshold(threshold float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(threshold)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// SetThreshold sets the Threshold property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#threshold
func (b *BloomEffect) SetThreshold(threshold float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(threshold)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// Weight returns the Weight property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#weight
func (b *BloomEffect) Weight(weight float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(weight)
	return BloomEffectFromJSObject(p, ba.ctx)
}

// SetWeight sets the Weight property of class BloomEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.bloomeffect#weight
func (b *BloomEffect) SetWeight(weight float64) *BloomEffect {
	p := ba.ctx.Get("BloomEffect").New(weight)
	return BloomEffectFromJSObject(p, ba.ctx)
}

*/
