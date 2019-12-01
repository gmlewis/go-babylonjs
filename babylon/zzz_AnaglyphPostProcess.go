// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphPostProcess represents a babylon.js AnaglyphPostProcess.
// Postprocess used to generate anaglyphic rendering
type AnaglyphPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnaglyphPostProcess) JSObject() js.Value { return a.p }

// AnaglyphPostProcess returns a AnaglyphPostProcess JavaScript class.
func (ba *Babylon) AnaglyphPostProcess() *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess")
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AnaglyphPostProcessFromJSObject returns a wrapped AnaglyphPostProcess JavaScript class.
func AnaglyphPostProcessFromJSObject(p js.Value, ctx js.Value) *AnaglyphPostProcess {
	return &AnaglyphPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// NewAnaglyphPostProcessOpts contains optional parameters for NewAnaglyphPostProcess.
type NewAnaglyphPostProcessOpts struct {
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
}

// NewAnaglyphPostProcess returns a new AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess
func (ba *Babylon) NewAnaglyphPostProcess(name string, options float64, rigCameras *Camera, opts *NewAnaglyphPostProcessOpts) *AnaglyphPostProcess {
	if opts == nil {
		opts = &NewAnaglyphPostProcessOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, name)
	args = append(args, options)
	args = append(args, rigCameras.JSObject())

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Reusable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reusable)
	}

	p := ba.ctx.Get("AnaglyphPostProcess").New(args...)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AnaglyphPostProcessActivateOpts contains optional parameters for AnaglyphPostProcess.Activate.
type AnaglyphPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#activate
func (a *AnaglyphPostProcess) Activate(camera *Camera, opts *AnaglyphPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &AnaglyphPostProcessActivateOpts{}
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

	retVal := a.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, a.ctx)
}

// Apply calls the Apply method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#apply
func (a *AnaglyphPostProcess) Apply() *Effect {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("apply", args...)
	return EffectFromJSObject(retVal, a.ctx)
}

// AnaglyphPostProcessDisposeOpts contains optional parameters for AnaglyphPostProcess.Dispose.
type AnaglyphPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#dispose
func (a *AnaglyphPostProcess) Dispose(opts *AnaglyphPostProcessDisposeOpts) {
	if opts == nil {
		opts = &AnaglyphPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	a.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#getcamera
func (a *AnaglyphPostProcess) GetCamera() *Camera {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getCamera", args...)
	return CameraFromJSObject(retVal, a.ctx)
}

// GetClassName calls the GetClassName method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#getclassname
func (a *AnaglyphPostProcess) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getClassName", args...)
	return retVal.String()
}

// GetEffect calls the GetEffect method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#geteffect
func (a *AnaglyphPostProcess) GetEffect() *Effect {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getEffect", args...)
	return EffectFromJSObject(retVal, a.ctx)
}

// GetEffectName calls the GetEffectName method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#geteffectname
func (a *AnaglyphPostProcess) GetEffectName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getEffectName", args...)
	return retVal.String()
}

// GetEngine calls the GetEngine method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#getengine
func (a *AnaglyphPostProcess) GetEngine() *Engine {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getEngine", args...)
	return EngineFromJSObject(retVal, a.ctx)
}

// IsReady calls the IsReady method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#isready
func (a *AnaglyphPostProcess) IsReady() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("isReady", args...)
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#isreusable
func (a *AnaglyphPostProcess) IsReusable() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("isReusable", args...)
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#marktexturedirty
func (a *AnaglyphPostProcess) MarkTextureDirty() {

	args := make([]interface{}, 0, 0+0)

	a.p.Call("markTextureDirty", args...)
}

// ShareOutputWith calls the ShareOutputWith method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#shareoutputwith
func (a *AnaglyphPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := a.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, a.ctx)
}

// AnaglyphPostProcessUpdateEffectOpts contains optional parameters for AnaglyphPostProcess.UpdateEffect.
type AnaglyphPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      *func()
	OnError         *func()
}

// UpdateEffect calls the UpdateEffect method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#updateeffect
func (a *AnaglyphPostProcess) UpdateEffect(opts *AnaglyphPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &AnaglyphPostProcessUpdateEffectOpts{}
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

	a.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the AnaglyphPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#useownoutput
func (a *AnaglyphPostProcess) UseOwnOutput() {

	args := make([]interface{}, 0, 0+0)

	a.p.Call("useOwnOutput", args...)
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#adaptscaletocurrentviewport
func (a *AnaglyphPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(adaptScaleToCurrentViewport)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#adaptscaletocurrentviewport
func (a *AnaglyphPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(adaptScaleToCurrentViewport)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alphaconstants
func (a *AnaglyphPostProcess) AlphaConstants(alphaConstants *Color4) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alphaConstants.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alphaconstants
func (a *AnaglyphPostProcess) SetAlphaConstants(alphaConstants *Color4) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alphaConstants.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alphamode
func (a *AnaglyphPostProcess) AlphaMode(alphaMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alphaMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alphamode
func (a *AnaglyphPostProcess) SetAlphaMode(alphaMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alphaMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alwaysforcepot
func (a *AnaglyphPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alwaysForcePOT)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#alwaysforcepot
func (a *AnaglyphPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(alwaysForcePOT)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#animations
func (a *AnaglyphPostProcess) Animations(animations *Animation) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(animations.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#animations
func (a *AnaglyphPostProcess) SetAnimations(animations *Animation) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(animations.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#aspectratio
func (a *AnaglyphPostProcess) AspectRatio(aspectRatio float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(aspectRatio)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#aspectratio
func (a *AnaglyphPostProcess) SetAspectRatio(aspectRatio float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(aspectRatio)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#autoclear
func (a *AnaglyphPostProcess) AutoClear(autoClear bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(autoClear)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#autoclear
func (a *AnaglyphPostProcess) SetAutoClear(autoClear bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(autoClear)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#clearcolor
func (a *AnaglyphPostProcess) ClearColor(clearColor *Color4) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(clearColor.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#clearcolor
func (a *AnaglyphPostProcess) SetClearColor(clearColor *Color4) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(clearColor.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#enablepixelperfectmode
func (a *AnaglyphPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(enablePixelPerfectMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#enablepixelperfectmode
func (a *AnaglyphPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(enablePixelPerfectMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#forcefullscreenviewport
func (a *AnaglyphPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(forceFullscreenViewport)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#forcefullscreenviewport
func (a *AnaglyphPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(forceFullscreenViewport)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#height
func (a *AnaglyphPostProcess) Height(height float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(height)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#height
func (a *AnaglyphPostProcess) SetHeight(height float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(height)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#inputtexture
func (a *AnaglyphPostProcess) InputTexture(inputTexture *InternalTexture) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(inputTexture.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#inputtexture
func (a *AnaglyphPostProcess) SetInputTexture(inputTexture *InternalTexture) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(inputTexture.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#inspectablecustomproperties
func (a *AnaglyphPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(inspectableCustomProperties.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#inspectablecustomproperties
func (a *AnaglyphPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(inspectableCustomProperties.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#issupported
func (a *AnaglyphPostProcess) IsSupported(isSupported bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(isSupported)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#issupported
func (a *AnaglyphPostProcess) SetIsSupported(isSupported bool) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(isSupported)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#name
func (a *AnaglyphPostProcess) Name(name string) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(name)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#name
func (a *AnaglyphPostProcess) SetName(name string) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(name)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onactivate
func (a *AnaglyphPostProcess) OnActivate(onActivate func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onActivate)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onactivate
func (a *AnaglyphPostProcess) SetOnActivate(onActivate func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onActivate)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onactivateobservable
func (a *AnaglyphPostProcess) OnActivateObservable(onActivateObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onActivateObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onactivateobservable
func (a *AnaglyphPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onActivateObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onafterrender
func (a *AnaglyphPostProcess) OnAfterRender(onAfterRender func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onAfterRender)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onafterrender
func (a *AnaglyphPostProcess) SetOnAfterRender(onAfterRender func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onAfterRender)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onafterrenderobservable
func (a *AnaglyphPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onAfterRenderObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onafterrenderobservable
func (a *AnaglyphPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onAfterRenderObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onapply
func (a *AnaglyphPostProcess) OnApply(onApply func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onApply)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onapply
func (a *AnaglyphPostProcess) SetOnApply(onApply func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onApply)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onapplyobservable
func (a *AnaglyphPostProcess) OnApplyObservable(onApplyObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onApplyObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onapplyobservable
func (a *AnaglyphPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onApplyObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onbeforerender
func (a *AnaglyphPostProcess) OnBeforeRender(onBeforeRender func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onBeforeRender)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onbeforerender
func (a *AnaglyphPostProcess) SetOnBeforeRender(onBeforeRender func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onBeforeRender)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onbeforerenderobservable
func (a *AnaglyphPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onBeforeRenderObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onbeforerenderobservable
func (a *AnaglyphPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onBeforeRenderObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onsizechanged
func (a *AnaglyphPostProcess) OnSizeChanged(onSizeChanged func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onSizeChanged)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onsizechanged
func (a *AnaglyphPostProcess) SetOnSizeChanged(onSizeChanged func()) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onSizeChanged)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onsizechangedobservable
func (a *AnaglyphPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onSizeChangedObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#onsizechangedobservable
func (a *AnaglyphPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(onSizeChangedObservable.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#rendertargetsamplingmode
func (a *AnaglyphPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(renderTargetSamplingMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#rendertargetsamplingmode
func (a *AnaglyphPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(renderTargetSamplingMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#samples
func (a *AnaglyphPostProcess) Samples(samples float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(samples)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#samples
func (a *AnaglyphPostProcess) SetSamples(samples float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(samples)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#scalemode
func (a *AnaglyphPostProcess) ScaleMode(scaleMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(scaleMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#scalemode
func (a *AnaglyphPostProcess) SetScaleMode(scaleMode float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(scaleMode)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#texelsize
func (a *AnaglyphPostProcess) TexelSize(texelSize *Vector2) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(texelSize.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#texelsize
func (a *AnaglyphPostProcess) SetTexelSize(texelSize *Vector2) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(texelSize.JSObject())
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#uniqueid
func (a *AnaglyphPostProcess) UniqueId(uniqueId float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(uniqueId)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#uniqueid
func (a *AnaglyphPostProcess) SetUniqueId(uniqueId float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(uniqueId)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#width
func (a *AnaglyphPostProcess) Width(width float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(width)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class AnaglyphPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphpostprocess#width
func (a *AnaglyphPostProcess) SetWidth(width float64) *AnaglyphPostProcess {
	p := ba.ctx.Get("AnaglyphPostProcess").New(width)
	return AnaglyphPostProcessFromJSObject(p, ba.ctx)
}

*/
