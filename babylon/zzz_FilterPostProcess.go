// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FilterPostProcess represents a babylon.js FilterPostProcess.
// Applies a kernel filter to the image
type FilterPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FilterPostProcess) JSObject() js.Value { return f.p }

// FilterPostProcess returns a FilterPostProcess JavaScript class.
func (ba *Babylon) FilterPostProcess() *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess")
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// FilterPostProcessFromJSObject returns a wrapped FilterPostProcess JavaScript class.
func FilterPostProcessFromJSObject(p js.Value, ctx js.Value) *FilterPostProcess {
	return &FilterPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// NewFilterPostProcessOpts contains optional parameters for NewFilterPostProcess.
type NewFilterPostProcessOpts struct {
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
}

// NewFilterPostProcess returns a new FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess
func (ba *Babylon) NewFilterPostProcess(name string, kernelMatrix *Matrix, options float64, camera *Camera, opts *NewFilterPostProcessOpts) *FilterPostProcess {
	if opts == nil {
		opts = &NewFilterPostProcessOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, name)
	args = append(args, kernelMatrix.JSObject())
	args = append(args, options)
	args = append(args, camera.JSObject())

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

	p := ba.ctx.Get("FilterPostProcess").New(args...)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// FilterPostProcessActivateOpts contains optional parameters for FilterPostProcess.Activate.
type FilterPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#activate
func (f *FilterPostProcess) Activate(camera *Camera, opts *FilterPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &FilterPostProcessActivateOpts{}
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

	retVal := f.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, f.ctx)
}

// Apply calls the Apply method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#apply
func (f *FilterPostProcess) Apply() *Effect {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("apply", args...)
	return EffectFromJSObject(retVal, f.ctx)
}

// FilterPostProcessDisposeOpts contains optional parameters for FilterPostProcess.Dispose.
type FilterPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#dispose
func (f *FilterPostProcess) Dispose(opts *FilterPostProcessDisposeOpts) {
	if opts == nil {
		opts = &FilterPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	f.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#getcamera
func (f *FilterPostProcess) GetCamera() *Camera {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getCamera", args...)
	return CameraFromJSObject(retVal, f.ctx)
}

// GetClassName calls the GetClassName method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#getclassname
func (f *FilterPostProcess) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getClassName", args...)
	return retVal.String()
}

// GetEffect calls the GetEffect method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#geteffect
func (f *FilterPostProcess) GetEffect() *Effect {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getEffect", args...)
	return EffectFromJSObject(retVal, f.ctx)
}

// GetEffectName calls the GetEffectName method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#geteffectname
func (f *FilterPostProcess) GetEffectName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getEffectName", args...)
	return retVal.String()
}

// GetEngine calls the GetEngine method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#getengine
func (f *FilterPostProcess) GetEngine() *Engine {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getEngine", args...)
	return EngineFromJSObject(retVal, f.ctx)
}

// IsReady calls the IsReady method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#isready
func (f *FilterPostProcess) IsReady() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("isReady", args...)
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#isreusable
func (f *FilterPostProcess) IsReusable() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("isReusable", args...)
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#marktexturedirty
func (f *FilterPostProcess) MarkTextureDirty() {

	args := make([]interface{}, 0, 0+0)

	f.p.Call("markTextureDirty", args...)
}

// ShareOutputWith calls the ShareOutputWith method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#shareoutputwith
func (f *FilterPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := f.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, f.ctx)
}

// FilterPostProcessUpdateEffectOpts contains optional parameters for FilterPostProcess.UpdateEffect.
type FilterPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      *func()
	OnError         *func()
}

// UpdateEffect calls the UpdateEffect method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#updateeffect
func (f *FilterPostProcess) UpdateEffect(opts *FilterPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &FilterPostProcessUpdateEffectOpts{}
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

	f.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the FilterPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#useownoutput
func (f *FilterPostProcess) UseOwnOutput() {

	args := make([]interface{}, 0, 0+0)

	f.p.Call("useOwnOutput", args...)
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#adaptscaletocurrentviewport
func (f *FilterPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(adaptScaleToCurrentViewport)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#adaptscaletocurrentviewport
func (f *FilterPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(adaptScaleToCurrentViewport)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alphaconstants
func (f *FilterPostProcess) AlphaConstants(alphaConstants *Color4) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alphaConstants.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alphaconstants
func (f *FilterPostProcess) SetAlphaConstants(alphaConstants *Color4) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alphaConstants.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alphamode
func (f *FilterPostProcess) AlphaMode(alphaMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alphaMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alphamode
func (f *FilterPostProcess) SetAlphaMode(alphaMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alphaMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alwaysforcepot
func (f *FilterPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alwaysForcePOT)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#alwaysforcepot
func (f *FilterPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(alwaysForcePOT)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#animations
func (f *FilterPostProcess) Animations(animations *Animation) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(animations.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#animations
func (f *FilterPostProcess) SetAnimations(animations *Animation) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(animations.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#aspectratio
func (f *FilterPostProcess) AspectRatio(aspectRatio float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(aspectRatio)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#aspectratio
func (f *FilterPostProcess) SetAspectRatio(aspectRatio float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(aspectRatio)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#autoclear
func (f *FilterPostProcess) AutoClear(autoClear bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(autoClear)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#autoclear
func (f *FilterPostProcess) SetAutoClear(autoClear bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(autoClear)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#clearcolor
func (f *FilterPostProcess) ClearColor(clearColor *Color4) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(clearColor.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#clearcolor
func (f *FilterPostProcess) SetClearColor(clearColor *Color4) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(clearColor.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#enablepixelperfectmode
func (f *FilterPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(enablePixelPerfectMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#enablepixelperfectmode
func (f *FilterPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(enablePixelPerfectMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#forcefullscreenviewport
func (f *FilterPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(forceFullscreenViewport)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#forcefullscreenviewport
func (f *FilterPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(forceFullscreenViewport)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#height
func (f *FilterPostProcess) Height(height float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(height)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#height
func (f *FilterPostProcess) SetHeight(height float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(height)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#inputtexture
func (f *FilterPostProcess) InputTexture(inputTexture *InternalTexture) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(inputTexture.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#inputtexture
func (f *FilterPostProcess) SetInputTexture(inputTexture *InternalTexture) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(inputTexture.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#inspectablecustomproperties
func (f *FilterPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(inspectableCustomProperties.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#inspectablecustomproperties
func (f *FilterPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(inspectableCustomProperties.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#issupported
func (f *FilterPostProcess) IsSupported(isSupported bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(isSupported)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#issupported
func (f *FilterPostProcess) SetIsSupported(isSupported bool) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(isSupported)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// KernelMatrix returns the KernelMatrix property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#kernelmatrix
func (f *FilterPostProcess) KernelMatrix(kernelMatrix *Matrix) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(kernelMatrix.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetKernelMatrix sets the KernelMatrix property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#kernelmatrix
func (f *FilterPostProcess) SetKernelMatrix(kernelMatrix *Matrix) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(kernelMatrix.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#name
func (f *FilterPostProcess) Name(name string) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(name)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#name
func (f *FilterPostProcess) SetName(name string) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(name)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onactivate
func (f *FilterPostProcess) OnActivate(onActivate func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onActivate)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onactivate
func (f *FilterPostProcess) SetOnActivate(onActivate func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onActivate)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onactivateobservable
func (f *FilterPostProcess) OnActivateObservable(onActivateObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onActivateObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onactivateobservable
func (f *FilterPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onActivateObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onafterrender
func (f *FilterPostProcess) OnAfterRender(onAfterRender func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onAfterRender)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onafterrender
func (f *FilterPostProcess) SetOnAfterRender(onAfterRender func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onAfterRender)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onafterrenderobservable
func (f *FilterPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onAfterRenderObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onafterrenderobservable
func (f *FilterPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onAfterRenderObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onapply
func (f *FilterPostProcess) OnApply(onApply func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onApply)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onapply
func (f *FilterPostProcess) SetOnApply(onApply func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onApply)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onapplyobservable
func (f *FilterPostProcess) OnApplyObservable(onApplyObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onApplyObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onapplyobservable
func (f *FilterPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onApplyObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onbeforerender
func (f *FilterPostProcess) OnBeforeRender(onBeforeRender func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onBeforeRender)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onbeforerender
func (f *FilterPostProcess) SetOnBeforeRender(onBeforeRender func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onBeforeRender)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onbeforerenderobservable
func (f *FilterPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onBeforeRenderObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onbeforerenderobservable
func (f *FilterPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onBeforeRenderObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onsizechanged
func (f *FilterPostProcess) OnSizeChanged(onSizeChanged func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onSizeChanged)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onsizechanged
func (f *FilterPostProcess) SetOnSizeChanged(onSizeChanged func()) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onSizeChanged)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onsizechangedobservable
func (f *FilterPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onSizeChangedObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#onsizechangedobservable
func (f *FilterPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(onSizeChangedObservable.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#rendertargetsamplingmode
func (f *FilterPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(renderTargetSamplingMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#rendertargetsamplingmode
func (f *FilterPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(renderTargetSamplingMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#samples
func (f *FilterPostProcess) Samples(samples float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(samples)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#samples
func (f *FilterPostProcess) SetSamples(samples float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(samples)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#scalemode
func (f *FilterPostProcess) ScaleMode(scaleMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(scaleMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#scalemode
func (f *FilterPostProcess) SetScaleMode(scaleMode float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(scaleMode)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#texelsize
func (f *FilterPostProcess) TexelSize(texelSize *Vector2) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(texelSize.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#texelsize
func (f *FilterPostProcess) SetTexelSize(texelSize *Vector2) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(texelSize.JSObject())
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#uniqueid
func (f *FilterPostProcess) UniqueId(uniqueId float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(uniqueId)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#uniqueid
func (f *FilterPostProcess) SetUniqueId(uniqueId float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(uniqueId)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#width
func (f *FilterPostProcess) Width(width float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(width)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class FilterPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.filterpostprocess#width
func (f *FilterPostProcess) SetWidth(width float64) *FilterPostProcess {
	p := ba.ctx.Get("FilterPostProcess").New(width)
	return FilterPostProcessFromJSObject(p, ba.ctx)
}

*/
