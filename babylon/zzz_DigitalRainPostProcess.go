// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DigitalRainPostProcess represents a babylon.js DigitalRainPostProcess.
// DigitalRainPostProcess helps rendering everithing in digital rain.
//
// Simmply add it to your scene and let the nerd that lives in you have fun.
// Example usage: var pp = new DigitalRainPostProcess(&amp;quot;digitalRain&amp;quot;, &amp;quot;20px Monospace&amp;quot;, camera);
type DigitalRainPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DigitalRainPostProcess) JSObject() js.Value { return d.p }

// DigitalRainPostProcess returns a DigitalRainPostProcess JavaScript class.
func (ba *Babylon) DigitalRainPostProcess() *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess")
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// DigitalRainPostProcessFromJSObject returns a wrapped DigitalRainPostProcess JavaScript class.
func DigitalRainPostProcessFromJSObject(p js.Value, ctx js.Value) *DigitalRainPostProcess {
	return &DigitalRainPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// DigitalRainPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func DigitalRainPostProcessArrayToJSArray(array []*DigitalRainPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDigitalRainPostProcessOpts contains optional parameters for NewDigitalRainPostProcess.
type NewDigitalRainPostProcessOpts struct {
	Options *string
}

// NewDigitalRainPostProcess returns a new DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess
func (ba *Babylon) NewDigitalRainPostProcess(name string, camera *Camera, opts *NewDigitalRainPostProcessOpts) *DigitalRainPostProcess {
	if opts == nil {
		opts = &NewDigitalRainPostProcessOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, camera.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Options)
	}

	p := ba.ctx.Get("DigitalRainPostProcess").New(args...)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// DigitalRainPostProcessActivateOpts contains optional parameters for DigitalRainPostProcess.Activate.
type DigitalRainPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#activate
func (d *DigitalRainPostProcess) Activate(camera *Camera, opts *DigitalRainPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &DigitalRainPostProcessActivateOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, camera.JSObject())

	if opts.SourceTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SourceTexture.JSObject())
	}
	if opts.ForceDepthStencil == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDepthStencil)
	}

	retVal := d.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, d.ctx)
}

// Apply calls the Apply method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#apply
func (d *DigitalRainPostProcess) Apply() *Effect {

	retVal := d.p.Call("apply")
	return EffectFromJSObject(retVal, d.ctx)
}

// DigitalRainPostProcessDisposeOpts contains optional parameters for DigitalRainPostProcess.Dispose.
type DigitalRainPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#dispose
func (d *DigitalRainPostProcess) Dispose(opts *DigitalRainPostProcessDisposeOpts) {
	if opts == nil {
		opts = &DigitalRainPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	d.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#getcamera
func (d *DigitalRainPostProcess) GetCamera() *Camera {

	retVal := d.p.Call("getCamera")
	return CameraFromJSObject(retVal, d.ctx)
}

// GetClassName calls the GetClassName method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#getclassname
func (d *DigitalRainPostProcess) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// GetEffect calls the GetEffect method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#geteffect
func (d *DigitalRainPostProcess) GetEffect() *Effect {

	retVal := d.p.Call("getEffect")
	return EffectFromJSObject(retVal, d.ctx)
}

// GetEffectName calls the GetEffectName method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#geteffectname
func (d *DigitalRainPostProcess) GetEffectName() string {

	retVal := d.p.Call("getEffectName")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#getengine
func (d *DigitalRainPostProcess) GetEngine() *Engine {

	retVal := d.p.Call("getEngine")
	return EngineFromJSObject(retVal, d.ctx)
}

// IsReady calls the IsReady method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#isready
func (d *DigitalRainPostProcess) IsReady() bool {

	retVal := d.p.Call("isReady")
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#isreusable
func (d *DigitalRainPostProcess) IsReusable() bool {

	retVal := d.p.Call("isReusable")
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#marktexturedirty
func (d *DigitalRainPostProcess) MarkTextureDirty() {

	d.p.Call("markTextureDirty")
}

// ShareOutputWith calls the ShareOutputWith method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#shareoutputwith
func (d *DigitalRainPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := d.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, d.ctx)
}

// DigitalRainPostProcessUpdateEffectOpts contains optional parameters for DigitalRainPostProcess.UpdateEffect.
type DigitalRainPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      *func()
	OnError         *func()
}

// UpdateEffect calls the UpdateEffect method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#updateeffect
func (d *DigitalRainPostProcess) UpdateEffect(opts *DigitalRainPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &DigitalRainPostProcessUpdateEffectOpts{}
	}

	args := make([]interface{}, 0, 0+6)

	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.Uniforms == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Uniforms)
	}
	if opts.Samplers == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Samplers)
	}
	if opts.IndexParameters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.IndexParameters)
	}
	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnCompiled)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	d.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the DigitalRainPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#useownoutput
func (d *DigitalRainPostProcess) UseOwnOutput() {

	d.p.Call("useOwnOutput")
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#adaptscaletocurrentviewport
func (d *DigitalRainPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(adaptScaleToCurrentViewport)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#adaptscaletocurrentviewport
func (d *DigitalRainPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(adaptScaleToCurrentViewport)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alphaconstants
func (d *DigitalRainPostProcess) AlphaConstants(alphaConstants *Color4) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alphaConstants.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alphaconstants
func (d *DigitalRainPostProcess) SetAlphaConstants(alphaConstants *Color4) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alphaConstants.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alphamode
func (d *DigitalRainPostProcess) AlphaMode(alphaMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alphaMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alphamode
func (d *DigitalRainPostProcess) SetAlphaMode(alphaMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alphaMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alwaysforcepot
func (d *DigitalRainPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alwaysForcePOT)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#alwaysforcepot
func (d *DigitalRainPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(alwaysForcePOT)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#animations
func (d *DigitalRainPostProcess) Animations(animations *Animation) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(animations.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#animations
func (d *DigitalRainPostProcess) SetAnimations(animations *Animation) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(animations.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#aspectratio
func (d *DigitalRainPostProcess) AspectRatio(aspectRatio float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(aspectRatio)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#aspectratio
func (d *DigitalRainPostProcess) SetAspectRatio(aspectRatio float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(aspectRatio)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#autoclear
func (d *DigitalRainPostProcess) AutoClear(autoClear bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(autoClear)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#autoclear
func (d *DigitalRainPostProcess) SetAutoClear(autoClear bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(autoClear)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#clearcolor
func (d *DigitalRainPostProcess) ClearColor(clearColor *Color4) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(clearColor.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#clearcolor
func (d *DigitalRainPostProcess) SetClearColor(clearColor *Color4) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(clearColor.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#enablepixelperfectmode
func (d *DigitalRainPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(enablePixelPerfectMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#enablepixelperfectmode
func (d *DigitalRainPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(enablePixelPerfectMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#forcefullscreenviewport
func (d *DigitalRainPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(forceFullscreenViewport)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#forcefullscreenviewport
func (d *DigitalRainPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(forceFullscreenViewport)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#height
func (d *DigitalRainPostProcess) Height(height float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(height)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#height
func (d *DigitalRainPostProcess) SetHeight(height float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(height)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#inputtexture
func (d *DigitalRainPostProcess) InputTexture(inputTexture *InternalTexture) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(inputTexture.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#inputtexture
func (d *DigitalRainPostProcess) SetInputTexture(inputTexture *InternalTexture) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(inputTexture.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#inspectablecustomproperties
func (d *DigitalRainPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(inspectableCustomProperties.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#inspectablecustomproperties
func (d *DigitalRainPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(inspectableCustomProperties.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#issupported
func (d *DigitalRainPostProcess) IsSupported(isSupported bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(isSupported)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#issupported
func (d *DigitalRainPostProcess) SetIsSupported(isSupported bool) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(isSupported)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// MixToNormal returns the MixToNormal property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#mixtonormal
func (d *DigitalRainPostProcess) MixToNormal(mixToNormal float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(mixToNormal)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetMixToNormal sets the MixToNormal property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#mixtonormal
func (d *DigitalRainPostProcess) SetMixToNormal(mixToNormal float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(mixToNormal)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// MixToTile returns the MixToTile property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#mixtotile
func (d *DigitalRainPostProcess) MixToTile(mixToTile float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(mixToTile)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetMixToTile sets the MixToTile property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#mixtotile
func (d *DigitalRainPostProcess) SetMixToTile(mixToTile float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(mixToTile)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#name
func (d *DigitalRainPostProcess) Name(name string) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(name)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#name
func (d *DigitalRainPostProcess) SetName(name string) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(name)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onactivate
func (d *DigitalRainPostProcess) OnActivate(onActivate func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onactivate
func (d *DigitalRainPostProcess) SetOnActivate(onActivate func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onactivateobservable
func (d *DigitalRainPostProcess) OnActivateObservable(onActivateObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onActivateObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onactivateobservable
func (d *DigitalRainPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onActivateObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onafterrender
func (d *DigitalRainPostProcess) OnAfterRender(onAfterRender func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onafterrender
func (d *DigitalRainPostProcess) SetOnAfterRender(onAfterRender func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onafterrenderobservable
func (d *DigitalRainPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onAfterRenderObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onafterrenderobservable
func (d *DigitalRainPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onAfterRenderObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onapply
func (d *DigitalRainPostProcess) OnApply(onApply func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onapply
func (d *DigitalRainPostProcess) SetOnApply(onApply func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onapplyobservable
func (d *DigitalRainPostProcess) OnApplyObservable(onApplyObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onApplyObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onapplyobservable
func (d *DigitalRainPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onApplyObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onbeforerender
func (d *DigitalRainPostProcess) OnBeforeRender(onBeforeRender func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onbeforerender
func (d *DigitalRainPostProcess) SetOnBeforeRender(onBeforeRender func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onbeforerenderobservable
func (d *DigitalRainPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onBeforeRenderObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onbeforerenderobservable
func (d *DigitalRainPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onBeforeRenderObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onsizechanged
func (d *DigitalRainPostProcess) OnSizeChanged(onSizeChanged func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onsizechanged
func (d *DigitalRainPostProcess) SetOnSizeChanged(onSizeChanged func()) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onsizechangedobservable
func (d *DigitalRainPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onSizeChangedObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#onsizechangedobservable
func (d *DigitalRainPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(onSizeChangedObservable.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#rendertargetsamplingmode
func (d *DigitalRainPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(renderTargetSamplingMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#rendertargetsamplingmode
func (d *DigitalRainPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(renderTargetSamplingMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#samples
func (d *DigitalRainPostProcess) Samples(samples float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(samples)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#samples
func (d *DigitalRainPostProcess) SetSamples(samples float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(samples)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#scalemode
func (d *DigitalRainPostProcess) ScaleMode(scaleMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(scaleMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#scalemode
func (d *DigitalRainPostProcess) SetScaleMode(scaleMode float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(scaleMode)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#texelsize
func (d *DigitalRainPostProcess) TexelSize(texelSize *Vector2) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(texelSize.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#texelsize
func (d *DigitalRainPostProcess) SetTexelSize(texelSize *Vector2) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(texelSize.JSObject())
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#uniqueid
func (d *DigitalRainPostProcess) UniqueId(uniqueId float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(uniqueId)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#uniqueid
func (d *DigitalRainPostProcess) SetUniqueId(uniqueId float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(uniqueId)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#width
func (d *DigitalRainPostProcess) Width(width float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(width)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class DigitalRainPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainpostprocess#width
func (d *DigitalRainPostProcess) SetWidth(width float64) *DigitalRainPostProcess {
	p := ba.ctx.Get("DigitalRainPostProcess").New(width)
	return DigitalRainPostProcessFromJSObject(p, ba.ctx)
}

*/
