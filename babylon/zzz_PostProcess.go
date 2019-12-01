// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcess represents a babylon.js PostProcess.
// PostProcess can be used to apply a shader to a texture after it has been rendered
// See &lt;a href=&#34;https://doc.babylonjs.com/how_to/how_to_use_postprocesses&#34;&gt;https://doc.babylonjs.com/how_to/how_to_use_postprocesses&lt;/a&gt;
type PostProcess struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PostProcess) JSObject() js.Value { return p.p }

// PostProcess returns a PostProcess JavaScript class.
func (ba *Babylon) PostProcess() *PostProcess {
	p := ba.ctx.Get("PostProcess")
	return PostProcessFromJSObject(p, ba.ctx)
}

// PostProcessFromJSObject returns a wrapped PostProcess JavaScript class.
func PostProcessFromJSObject(p js.Value, ctx js.Value) *PostProcess {
	return &PostProcess{p: p, ctx: ctx}
}

// PostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func PostProcessArrayToJSArray(array []*PostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPostProcessOpts contains optional parameters for NewPostProcess.
type NewPostProcessOpts struct {
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	Defines          *string
	TextureType      *float64
	VertexUrl        *string
	IndexParameters  *interface{}
	BlockCompilation *bool
}

// NewPostProcess returns a new PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess
func (ba *Babylon) NewPostProcess(name string, fragmentUrl string, parameters string, samplers string, options float64, camera *Camera, opts *NewPostProcessOpts) *PostProcess {
	if opts == nil {
		opts = &NewPostProcessOpts{}
	}

	args := make([]interface{}, 0, 6+8)

	args = append(args, name)
	args = append(args, fragmentUrl)
	args = append(args, parameters)
	args = append(args, samplers)
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
	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.TextureType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureType)
	}
	if opts.VertexUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.VertexUrl)
	}
	if opts.IndexParameters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.IndexParameters)
	}
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("PostProcess").New(args...)
	return PostProcessFromJSObject(p, ba.ctx)
}

// PostProcessActivateOpts contains optional parameters for PostProcess.Activate.
type PostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#activate
func (p *PostProcess) Activate(camera *Camera, opts *PostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &PostProcessActivateOpts{}
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

	retVal := p.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, p.ctx)
}

// Apply calls the Apply method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#apply
func (p *PostProcess) Apply() *Effect {

	retVal := p.p.Call("apply")
	return EffectFromJSObject(retVal, p.ctx)
}

// PostProcessDisposeOpts contains optional parameters for PostProcess.Dispose.
type PostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#dispose
func (p *PostProcess) Dispose(opts *PostProcessDisposeOpts) {
	if opts == nil {
		opts = &PostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	p.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#getcamera
func (p *PostProcess) GetCamera() *Camera {

	retVal := p.p.Call("getCamera")
	return CameraFromJSObject(retVal, p.ctx)
}

// GetClassName calls the GetClassName method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#getclassname
func (p *PostProcess) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// GetEffect calls the GetEffect method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#geteffect
func (p *PostProcess) GetEffect() *Effect {

	retVal := p.p.Call("getEffect")
	return EffectFromJSObject(retVal, p.ctx)
}

// GetEffectName calls the GetEffectName method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#geteffectname
func (p *PostProcess) GetEffectName() string {

	retVal := p.p.Call("getEffectName")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#getengine
func (p *PostProcess) GetEngine() *Engine {

	retVal := p.p.Call("getEngine")
	return EngineFromJSObject(retVal, p.ctx)
}

// IsReady calls the IsReady method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#isready
func (p *PostProcess) IsReady() bool {

	retVal := p.p.Call("isReady")
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#isreusable
func (p *PostProcess) IsReusable() bool {

	retVal := p.p.Call("isReusable")
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#marktexturedirty
func (p *PostProcess) MarkTextureDirty() {

	p.p.Call("markTextureDirty")
}

// ShareOutputWith calls the ShareOutputWith method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#shareoutputwith
func (p *PostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := p.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, p.ctx)
}

// PostProcessUpdateEffectOpts contains optional parameters for PostProcess.UpdateEffect.
type PostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      *func()
	OnError         *func()
}

// UpdateEffect calls the UpdateEffect method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#updateeffect
func (p *PostProcess) UpdateEffect(opts *PostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &PostProcessUpdateEffectOpts{}
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

	p.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#useownoutput
func (p *PostProcess) UseOwnOutput() {

	p.p.Call("useOwnOutput")
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#adaptscaletocurrentviewport
func (p *PostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(adaptScaleToCurrentViewport)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#adaptscaletocurrentviewport
func (p *PostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(adaptScaleToCurrentViewport)
	return PostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alphaconstants
func (p *PostProcess) AlphaConstants(alphaConstants *Color4) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alphaConstants.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alphaconstants
func (p *PostProcess) SetAlphaConstants(alphaConstants *Color4) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alphaConstants.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alphamode
func (p *PostProcess) AlphaMode(alphaMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alphaMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alphamode
func (p *PostProcess) SetAlphaMode(alphaMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alphaMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alwaysforcepot
func (p *PostProcess) AlwaysForcePOT(alwaysForcePOT bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alwaysForcePOT)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#alwaysforcepot
func (p *PostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(alwaysForcePOT)
	return PostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#animations
func (p *PostProcess) Animations(animations *Animation) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(animations.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#animations
func (p *PostProcess) SetAnimations(animations *Animation) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(animations.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#aspectratio
func (p *PostProcess) AspectRatio(aspectRatio float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(aspectRatio)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#aspectratio
func (p *PostProcess) SetAspectRatio(aspectRatio float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(aspectRatio)
	return PostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#autoclear
func (p *PostProcess) AutoClear(autoClear bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(autoClear)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#autoclear
func (p *PostProcess) SetAutoClear(autoClear bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(autoClear)
	return PostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#clearcolor
func (p *PostProcess) ClearColor(clearColor *Color4) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(clearColor.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#clearcolor
func (p *PostProcess) SetClearColor(clearColor *Color4) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(clearColor.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#enablepixelperfectmode
func (p *PostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(enablePixelPerfectMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#enablepixelperfectmode
func (p *PostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(enablePixelPerfectMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#forcefullscreenviewport
func (p *PostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(forceFullscreenViewport)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#forcefullscreenviewport
func (p *PostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(forceFullscreenViewport)
	return PostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#height
func (p *PostProcess) Height(height float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(height)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#height
func (p *PostProcess) SetHeight(height float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(height)
	return PostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#inputtexture
func (p *PostProcess) InputTexture(inputTexture *InternalTexture) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(inputTexture.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#inputtexture
func (p *PostProcess) SetInputTexture(inputTexture *InternalTexture) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(inputTexture.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#inspectablecustomproperties
func (p *PostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(inspectableCustomProperties.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#inspectablecustomproperties
func (p *PostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(inspectableCustomProperties.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#issupported
func (p *PostProcess) IsSupported(isSupported bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(isSupported)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#issupported
func (p *PostProcess) SetIsSupported(isSupported bool) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(isSupported)
	return PostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#name
func (p *PostProcess) Name(name string) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(name)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#name
func (p *PostProcess) SetName(name string) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(name)
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onactivate
func (p *PostProcess) OnActivate(onActivate func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onactivate
func (p *PostProcess) SetOnActivate(onActivate func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onactivateobservable
func (p *PostProcess) OnActivateObservable(onActivateObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onActivateObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onactivateobservable
func (p *PostProcess) SetOnActivateObservable(onActivateObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onActivateObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onafterrender
func (p *PostProcess) OnAfterRender(onAfterRender func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onafterrender
func (p *PostProcess) SetOnAfterRender(onAfterRender func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onafterrenderobservable
func (p *PostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onAfterRenderObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onafterrenderobservable
func (p *PostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onAfterRenderObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onapply
func (p *PostProcess) OnApply(onApply func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onapply
func (p *PostProcess) SetOnApply(onApply func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onapplyobservable
func (p *PostProcess) OnApplyObservable(onApplyObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onApplyObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onapplyobservable
func (p *PostProcess) SetOnApplyObservable(onApplyObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onApplyObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onbeforerender
func (p *PostProcess) OnBeforeRender(onBeforeRender func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onbeforerender
func (p *PostProcess) SetOnBeforeRender(onBeforeRender func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onbeforerenderobservable
func (p *PostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onBeforeRenderObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onbeforerenderobservable
func (p *PostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onBeforeRenderObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onsizechanged
func (p *PostProcess) OnSizeChanged(onSizeChanged func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onsizechanged
func (p *PostProcess) SetOnSizeChanged(onSizeChanged func()) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return PostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onsizechangedobservable
func (p *PostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onSizeChangedObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#onsizechangedobservable
func (p *PostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(onSizeChangedObservable.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#rendertargetsamplingmode
func (p *PostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(renderTargetSamplingMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#rendertargetsamplingmode
func (p *PostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(renderTargetSamplingMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#samples
func (p *PostProcess) Samples(samples float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(samples)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#samples
func (p *PostProcess) SetSamples(samples float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(samples)
	return PostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#scalemode
func (p *PostProcess) ScaleMode(scaleMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(scaleMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#scalemode
func (p *PostProcess) SetScaleMode(scaleMode float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(scaleMode)
	return PostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#texelsize
func (p *PostProcess) TexelSize(texelSize *Vector2) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(texelSize.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#texelsize
func (p *PostProcess) SetTexelSize(texelSize *Vector2) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(texelSize.JSObject())
	return PostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#uniqueid
func (p *PostProcess) UniqueId(uniqueId float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(uniqueId)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#uniqueid
func (p *PostProcess) SetUniqueId(uniqueId float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(uniqueId)
	return PostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#width
func (p *PostProcess) Width(width float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(width)
	return PostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class PostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess#width
func (p *PostProcess) SetWidth(width float64) *PostProcess {
	p := ba.ctx.Get("PostProcess").New(width)
	return PostProcessFromJSObject(p, ba.ctx)
}

*/
