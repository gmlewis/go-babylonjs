// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// UtilityLayerRenderer represents a babylon.js UtilityLayerRenderer.
// Renders a layer on top of an existing scene
type UtilityLayerRenderer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (u *UtilityLayerRenderer) JSObject() js.Value { return u.p }

// UtilityLayerRenderer returns a UtilityLayerRenderer JavaScript class.
func (ba *Babylon) UtilityLayerRenderer() *UtilityLayerRenderer {
	p := ba.ctx.Get("UtilityLayerRenderer")
	return UtilityLayerRendererFromJSObject(p, ba.ctx)
}

// UtilityLayerRendererFromJSObject returns a wrapped UtilityLayerRenderer JavaScript class.
func UtilityLayerRendererFromJSObject(p js.Value, ctx js.Value) *UtilityLayerRenderer {
	return &UtilityLayerRenderer{p: p, ctx: ctx}
}

// UtilityLayerRendererArrayToJSArray returns a JavaScript Array for the wrapped array.
func UtilityLayerRendererArrayToJSArray(array []*UtilityLayerRenderer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewUtilityLayerRendererOpts contains optional parameters for NewUtilityLayerRenderer.
type NewUtilityLayerRendererOpts struct {
	HandleEvents *bool
}

// NewUtilityLayerRenderer returns a new UtilityLayerRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#constructor
func (ba *Babylon) NewUtilityLayerRenderer(originalScene *Scene, opts *NewUtilityLayerRendererOpts) *UtilityLayerRenderer {
	if opts == nil {
		opts = &NewUtilityLayerRendererOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, originalScene.JSObject())

	if opts.HandleEvents == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.HandleEvents)
	}

	p := ba.ctx.Get("UtilityLayerRenderer").New(args...)
	return UtilityLayerRendererFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the UtilityLayerRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#dispose
func (u *UtilityLayerRenderer) Dispose() {

	u.p.Call("dispose")
}

// GetRenderCamera calls the GetRenderCamera method on the UtilityLayerRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#getrendercamera
func (u *UtilityLayerRenderer) GetRenderCamera() *Camera {

	retVal := u.p.Call("getRenderCamera")
	return CameraFromJSObject(retVal, u.ctx)
}

// Render calls the Render method on the UtilityLayerRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#render
func (u *UtilityLayerRenderer) Render() {

	u.p.Call("render")
}

// SetRenderCamera calls the SetRenderCamera method on the UtilityLayerRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#setrendercamera
func (u *UtilityLayerRenderer) SetRenderCamera(cam *Camera) {

	args := make([]interface{}, 0, 1+0)

	if cam == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, cam.JSObject())
	}

	u.p.Call("setRenderCamera", args...)
}

// DefaultKeepDepthUtilityLayer returns the DefaultKeepDepthUtilityLayer property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#defaultkeepdepthutilitylayer
func (u *UtilityLayerRenderer) DefaultKeepDepthUtilityLayer() *UtilityLayerRenderer {
	retVal := u.p.Get("DefaultKeepDepthUtilityLayer")
	return UtilityLayerRendererFromJSObject(retVal, u.ctx)
}

// SetDefaultKeepDepthUtilityLayer sets the DefaultKeepDepthUtilityLayer property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#defaultkeepdepthutilitylayer
func (u *UtilityLayerRenderer) SetDefaultKeepDepthUtilityLayer(DefaultKeepDepthUtilityLayer *UtilityLayerRenderer) *UtilityLayerRenderer {
	u.p.Set("DefaultKeepDepthUtilityLayer", DefaultKeepDepthUtilityLayer.JSObject())
	return u
}

// DefaultUtilityLayer returns the DefaultUtilityLayer property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#defaultutilitylayer
func (u *UtilityLayerRenderer) DefaultUtilityLayer() *UtilityLayerRenderer {
	retVal := u.p.Get("DefaultUtilityLayer")
	return UtilityLayerRendererFromJSObject(retVal, u.ctx)
}

// SetDefaultUtilityLayer sets the DefaultUtilityLayer property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#defaultutilitylayer
func (u *UtilityLayerRenderer) SetDefaultUtilityLayer(DefaultUtilityLayer *UtilityLayerRenderer) *UtilityLayerRenderer {
	u.p.Set("DefaultUtilityLayer", DefaultUtilityLayer.JSObject())
	return u
}

// MainSceneTrackerPredicate returns the MainSceneTrackerPredicate property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#mainscenetrackerpredicate
func (u *UtilityLayerRenderer) MainSceneTrackerPredicate() js.Value {
	retVal := u.p.Get("mainSceneTrackerPredicate")
	return retVal
}

// SetMainSceneTrackerPredicate sets the MainSceneTrackerPredicate property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#mainscenetrackerpredicate
func (u *UtilityLayerRenderer) SetMainSceneTrackerPredicate(mainSceneTrackerPredicate JSFunc) *UtilityLayerRenderer {
	u.p.Set("mainSceneTrackerPredicate", js.FuncOf(mainSceneTrackerPredicate))
	return u
}

// OnPointerOutObservable returns the OnPointerOutObservable property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#onpointeroutobservable
func (u *UtilityLayerRenderer) OnPointerOutObservable() *Observable {
	retVal := u.p.Get("onPointerOutObservable")
	return ObservableFromJSObject(retVal, u.ctx)
}

// SetOnPointerOutObservable sets the OnPointerOutObservable property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#onpointeroutobservable
func (u *UtilityLayerRenderer) SetOnPointerOutObservable(onPointerOutObservable *Observable) *UtilityLayerRenderer {
	u.p.Set("onPointerOutObservable", onPointerOutObservable.JSObject())
	return u
}

// OnlyCheckPointerDownEvents returns the OnlyCheckPointerDownEvents property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#onlycheckpointerdownevents
func (u *UtilityLayerRenderer) OnlyCheckPointerDownEvents() bool {
	retVal := u.p.Get("onlyCheckPointerDownEvents")
	return retVal.Bool()
}

// SetOnlyCheckPointerDownEvents sets the OnlyCheckPointerDownEvents property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#onlycheckpointerdownevents
func (u *UtilityLayerRenderer) SetOnlyCheckPointerDownEvents(onlyCheckPointerDownEvents bool) *UtilityLayerRenderer {
	u.p.Set("onlyCheckPointerDownEvents", onlyCheckPointerDownEvents)
	return u
}

// OriginalScene returns the OriginalScene property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#originalscene
func (u *UtilityLayerRenderer) OriginalScene() *Scene {
	retVal := u.p.Get("originalScene")
	return SceneFromJSObject(retVal, u.ctx)
}

// SetOriginalScene sets the OriginalScene property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#originalscene
func (u *UtilityLayerRenderer) SetOriginalScene(originalScene *Scene) *UtilityLayerRenderer {
	u.p.Set("originalScene", originalScene.JSObject())
	return u
}

// PickUtilitySceneFirst returns the PickUtilitySceneFirst property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#pickutilityscenefirst
func (u *UtilityLayerRenderer) PickUtilitySceneFirst() bool {
	retVal := u.p.Get("pickUtilitySceneFirst")
	return retVal.Bool()
}

// SetPickUtilitySceneFirst sets the PickUtilitySceneFirst property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#pickutilityscenefirst
func (u *UtilityLayerRenderer) SetPickUtilitySceneFirst(pickUtilitySceneFirst bool) *UtilityLayerRenderer {
	u.p.Set("pickUtilitySceneFirst", pickUtilitySceneFirst)
	return u
}

// ProcessAllEvents returns the ProcessAllEvents property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#processallevents
func (u *UtilityLayerRenderer) ProcessAllEvents() bool {
	retVal := u.p.Get("processAllEvents")
	return retVal.Bool()
}

// SetProcessAllEvents sets the ProcessAllEvents property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#processallevents
func (u *UtilityLayerRenderer) SetProcessAllEvents(processAllEvents bool) *UtilityLayerRenderer {
	u.p.Set("processAllEvents", processAllEvents)
	return u
}

// ShouldRender returns the ShouldRender property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#shouldrender
func (u *UtilityLayerRenderer) ShouldRender() bool {
	retVal := u.p.Get("shouldRender")
	return retVal.Bool()
}

// SetShouldRender sets the ShouldRender property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#shouldrender
func (u *UtilityLayerRenderer) SetShouldRender(shouldRender bool) *UtilityLayerRenderer {
	u.p.Set("shouldRender", shouldRender)
	return u
}

// UtilityLayerScene returns the UtilityLayerScene property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#utilitylayerscene
func (u *UtilityLayerRenderer) UtilityLayerScene() *Scene {
	retVal := u.p.Get("utilityLayerScene")
	return SceneFromJSObject(retVal, u.ctx)
}

// SetUtilityLayerScene sets the UtilityLayerScene property of class UtilityLayerRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.utilitylayerrenderer#utilitylayerscene
func (u *UtilityLayerRenderer) SetUtilityLayerScene(utilityLayerScene *Scene) *UtilityLayerRenderer {
	u.p.Set("utilityLayerScene", utilityLayerScene.JSObject())
	return u
}
