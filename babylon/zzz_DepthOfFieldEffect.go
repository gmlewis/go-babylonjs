// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthOfFieldEffect represents a babylon.js DepthOfFieldEffect.
// The depth of field effect applies a blur to objects that are closer or further from where the camera is focusing.
type DepthOfFieldEffect struct {
	*PostProcessRenderEffect
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldEffect) JSObject() js.Value { return d.p }

// DepthOfFieldEffect returns a DepthOfFieldEffect JavaScript class.
func (ba *Babylon) DepthOfFieldEffect() *DepthOfFieldEffect {
	p := ba.ctx.Get("DepthOfFieldEffect")
	return DepthOfFieldEffectFromJSObject(p, ba.ctx)
}

// DepthOfFieldEffectFromJSObject returns a wrapped DepthOfFieldEffect JavaScript class.
func DepthOfFieldEffectFromJSObject(p js.Value, ctx js.Value) *DepthOfFieldEffect {
	return &DepthOfFieldEffect{PostProcessRenderEffect: PostProcessRenderEffectFromJSObject(p, ctx), ctx: ctx}
}

// DepthOfFieldEffectArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthOfFieldEffectArrayToJSArray(array []*DepthOfFieldEffect) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDepthOfFieldEffectOpts contains optional parameters for NewDepthOfFieldEffect.
type NewDepthOfFieldEffectOpts struct {
	BlurLevel           *DepthOfFieldEffectBlurLevel
	PipelineTextureType *float64
	BlockCompilation    *bool
}

// NewDepthOfFieldEffect returns a new DepthOfFieldEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#constructor
func (ba *Babylon) NewDepthOfFieldEffect(scene *Scene, depthTexture *RenderTargetTexture, opts *NewDepthOfFieldEffectOpts) *DepthOfFieldEffect {
	if opts == nil {
		opts = &NewDepthOfFieldEffectOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, scene.JSObject())
	args = append(args, depthTexture.JSObject())

	if opts.BlurLevel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.BlurLevel.JSObject())
	}
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

	p := ba.ctx.Get("DepthOfFieldEffect").New(args...)
	return DepthOfFieldEffectFromJSObject(p, ba.ctx)
}

// DisposeEffects calls the DisposeEffects method on the DepthOfFieldEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#disposeeffects
func (d *DepthOfFieldEffect) DisposeEffects(camera *Camera) {

	args := make([]interface{}, 0, 1+0)

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	d.p.Call("disposeEffects", args...)
}

// GetClassName calls the GetClassName method on the DepthOfFieldEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#getclassname
func (d *DepthOfFieldEffect) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// DepthTexture returns the DepthTexture property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#depthtexture
func (d *DepthOfFieldEffect) DepthTexture() *RenderTargetTexture {
	retVal := d.p.Get("depthTexture")
	return RenderTargetTextureFromJSObject(retVal, d.ctx)
}

// SetDepthTexture sets the DepthTexture property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#depthtexture
func (d *DepthOfFieldEffect) SetDepthTexture(depthTexture *RenderTargetTexture) *DepthOfFieldEffect {
	d.p.Set("depthTexture", depthTexture.JSObject())
	return d
}

// FStop returns the FStop property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#fstop
func (d *DepthOfFieldEffect) FStop() float64 {
	retVal := d.p.Get("fStop")
	return retVal.Float()
}

// SetFStop sets the FStop property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#fstop
func (d *DepthOfFieldEffect) SetFStop(fStop float64) *DepthOfFieldEffect {
	d.p.Set("fStop", fStop)
	return d
}

// FocalLength returns the FocalLength property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#focallength
func (d *DepthOfFieldEffect) FocalLength() float64 {
	retVal := d.p.Get("focalLength")
	return retVal.Float()
}

// SetFocalLength sets the FocalLength property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#focallength
func (d *DepthOfFieldEffect) SetFocalLength(focalLength float64) *DepthOfFieldEffect {
	d.p.Set("focalLength", focalLength)
	return d
}

// FocusDistance returns the FocusDistance property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#focusdistance
func (d *DepthOfFieldEffect) FocusDistance() float64 {
	retVal := d.p.Get("focusDistance")
	return retVal.Float()
}

// SetFocusDistance sets the FocusDistance property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#focusdistance
func (d *DepthOfFieldEffect) SetFocusDistance(focusDistance float64) *DepthOfFieldEffect {
	d.p.Set("focusDistance", focusDistance)
	return d
}

// LensSize returns the LensSize property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#lenssize
func (d *DepthOfFieldEffect) LensSize() float64 {
	retVal := d.p.Get("lensSize")
	return retVal.Float()
}

// SetLensSize sets the LensSize property of class DepthOfFieldEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect#lenssize
func (d *DepthOfFieldEffect) SetLensSize(lensSize float64) *DepthOfFieldEffect {
	d.p.Set("lensSize", lensSize)
	return d
}
