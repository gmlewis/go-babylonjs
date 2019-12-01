// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GlowLayer represents a babylon.js GlowLayer.
// The glow layer Helps adding a glow effect around the emissive parts of a mesh.
//
// Documentation: &lt;a href=&#34;https://doc.babylonjs.com/how_to/glow_layer&#34;&gt;https://doc.babylonjs.com/how_to/glow_layer&lt;/a&gt;
type GlowLayer struct {
	*EffectLayer
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GlowLayer) JSObject() js.Value { return g.p }

// GlowLayer returns a GlowLayer JavaScript class.
func (ba *Babylon) GlowLayer() *GlowLayer {
	p := ba.ctx.Get("GlowLayer")
	return GlowLayerFromJSObject(p, ba.ctx)
}

// GlowLayerFromJSObject returns a wrapped GlowLayer JavaScript class.
func GlowLayerFromJSObject(p js.Value, ctx js.Value) *GlowLayer {
	return &GlowLayer{EffectLayer: EffectLayerFromJSObject(p, ctx), ctx: ctx}
}

// GlowLayerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GlowLayerArrayToJSArray(array []*GlowLayer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGlowLayerOpts contains optional parameters for NewGlowLayer.
type NewGlowLayerOpts struct {
	Options js.Value
}

// NewGlowLayer returns a new GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer
func (ba *Babylon) NewGlowLayer(name string, scene *Scene, opts *NewGlowLayerOpts) *GlowLayer {
	if opts == nil {
		opts = &NewGlowLayerOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, scene.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	p := ba.ctx.Get("GlowLayer").New(args...)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// AddExcludedMesh calls the AddExcludedMesh method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#addexcludedmesh
func (g *GlowLayer) AddExcludedMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("addExcludedMesh", args...)
}

// AddIncludedOnlyMesh calls the AddIncludedOnlyMesh method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#addincludedonlymesh
func (g *GlowLayer) AddIncludedOnlyMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("addIncludedOnlyMesh", args...)
}

// Dispose calls the Dispose method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#dispose
func (g *GlowLayer) Dispose() {

	g.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#getclassname
func (g *GlowLayer) GetClassName() string {

	retVal := g.p.Call("getClassName")
	return retVal.String()
}

// GetEffectName calls the GetEffectName method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#geteffectname
func (g *GlowLayer) GetEffectName() string {

	retVal := g.p.Call("getEffectName")
	return retVal.String()
}

// HasMesh calls the HasMesh method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#hasmesh
func (g *GlowLayer) HasMesh(mesh *AbstractMesh) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := g.p.Call("hasMesh", args...)
	return retVal.Bool()
}

// IsReady calls the IsReady method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#isready
func (g *GlowLayer) IsReady(subMesh *SubMesh, useInstances bool) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, subMesh.JSObject())
	args = append(args, useInstances)

	retVal := g.p.Call("isReady", args...)
	return retVal.Bool()
}

// NeedStencil calls the NeedStencil method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#needstencil
func (g *GlowLayer) NeedStencil() bool {

	retVal := g.p.Call("needStencil")
	return retVal.Bool()
}

// Parse calls the Parse method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#parse
func (g *GlowLayer) Parse(parsedGlowLayer interface{}, scene *Scene, rootUrl string) *GlowLayer {

	args := make([]interface{}, 0, 3+0)

	args = append(args, parsedGlowLayer)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := g.p.Call("Parse", args...)
	return GlowLayerFromJSObject(retVal, g.ctx)
}

// ReferenceMeshToUseItsOwnMaterial calls the ReferenceMeshToUseItsOwnMaterial method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#referencemeshtouseitsownmaterial
func (g *GlowLayer) ReferenceMeshToUseItsOwnMaterial(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("referenceMeshToUseItsOwnMaterial", args...)
}

// RemoveExcludedMesh calls the RemoveExcludedMesh method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#removeexcludedmesh
func (g *GlowLayer) RemoveExcludedMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("removeExcludedMesh", args...)
}

// RemoveIncludedOnlyMesh calls the RemoveIncludedOnlyMesh method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#removeincludedonlymesh
func (g *GlowLayer) RemoveIncludedOnlyMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("removeIncludedOnlyMesh", args...)
}

// Render calls the Render method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#render
func (g *GlowLayer) Render() {

	g.p.Call("render")
}

// Serialize calls the Serialize method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#serialize
func (g *GlowLayer) Serialize() interface{} {

	retVal := g.p.Call("serialize")
	return retVal
}

// ShouldRender calls the ShouldRender method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#shouldrender
func (g *GlowLayer) ShouldRender() bool {

	retVal := g.p.Call("shouldRender")
	return retVal.Bool()
}

// UnReferenceMeshFromUsingItsOwnMaterial calls the UnReferenceMeshFromUsingItsOwnMaterial method on the GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#unreferencemeshfromusingitsownmaterial
func (g *GlowLayer) UnReferenceMeshFromUsingItsOwnMaterial(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("unReferenceMeshFromUsingItsOwnMaterial", args...)
}

/*

// BlurKernelSize returns the BlurKernelSize property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#blurkernelsize
func (g *GlowLayer) BlurKernelSize(blurKernelSize float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(blurKernelSize)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetBlurKernelSize sets the BlurKernelSize property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#blurkernelsize
func (g *GlowLayer) SetBlurKernelSize(blurKernelSize float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(blurKernelSize)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#camera
func (g *GlowLayer) Camera(camera *Camera) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(camera.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#camera
func (g *GlowLayer) SetCamera(camera *Camera) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(camera.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// CustomEmissiveColorSelector returns the CustomEmissiveColorSelector property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#customemissivecolorselector
func (g *GlowLayer) CustomEmissiveColorSelector(customEmissiveColorSelector func()) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {customEmissiveColorSelector(); return nil}))
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetCustomEmissiveColorSelector sets the CustomEmissiveColorSelector property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#customemissivecolorselector
func (g *GlowLayer) SetCustomEmissiveColorSelector(customEmissiveColorSelector func()) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {customEmissiveColorSelector(); return nil}))
	return GlowLayerFromJSObject(p, ba.ctx)
}

// CustomEmissiveTextureSelector returns the CustomEmissiveTextureSelector property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#customemissivetextureselector
func (g *GlowLayer) CustomEmissiveTextureSelector(customEmissiveTextureSelector func()) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {customEmissiveTextureSelector(); return nil}))
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetCustomEmissiveTextureSelector sets the CustomEmissiveTextureSelector property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#customemissivetextureselector
func (g *GlowLayer) SetCustomEmissiveTextureSelector(customEmissiveTextureSelector func()) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {customEmissiveTextureSelector(); return nil}))
	return GlowLayerFromJSObject(p, ba.ctx)
}

// DefaultBlurKernelSize returns the DefaultBlurKernelSize property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#defaultblurkernelsize
func (g *GlowLayer) DefaultBlurKernelSize(DefaultBlurKernelSize float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(DefaultBlurKernelSize)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetDefaultBlurKernelSize sets the DefaultBlurKernelSize property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#defaultblurkernelsize
func (g *GlowLayer) SetDefaultBlurKernelSize(DefaultBlurKernelSize float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(DefaultBlurKernelSize)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// DefaultTextureRatio returns the DefaultTextureRatio property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#defaulttextureratio
func (g *GlowLayer) DefaultTextureRatio(DefaultTextureRatio float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(DefaultTextureRatio)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetDefaultTextureRatio sets the DefaultTextureRatio property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#defaulttextureratio
func (g *GlowLayer) SetDefaultTextureRatio(DefaultTextureRatio float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(DefaultTextureRatio)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// EffectName returns the EffectName property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#effectname
func (g *GlowLayer) EffectName(EffectName string) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(EffectName)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetEffectName sets the EffectName property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#effectname
func (g *GlowLayer) SetEffectName(EffectName string) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(EffectName)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// Intensity returns the Intensity property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#intensity
func (g *GlowLayer) Intensity(intensity float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(intensity)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetIntensity sets the Intensity property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#intensity
func (g *GlowLayer) SetIntensity(intensity float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(intensity)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// IsEnabled returns the IsEnabled property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#isenabled
func (g *GlowLayer) IsEnabled(isEnabled bool) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(isEnabled)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetIsEnabled sets the IsEnabled property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#isenabled
func (g *GlowLayer) SetIsEnabled(isEnabled bool) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(isEnabled)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#name
func (g *GlowLayer) Name(name string) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(name)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#name
func (g *GlowLayer) SetName(name string) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(name)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// NeutralColor returns the NeutralColor property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#neutralcolor
func (g *GlowLayer) NeutralColor(neutralColor *Color4) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(neutralColor.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetNeutralColor sets the NeutralColor property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#neutralcolor
func (g *GlowLayer) SetNeutralColor(neutralColor *Color4) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(neutralColor.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnAfterComposeObservable returns the OnAfterComposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onaftercomposeobservable
func (g *GlowLayer) OnAfterComposeObservable(onAfterComposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onAfterComposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnAfterComposeObservable sets the OnAfterComposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onaftercomposeobservable
func (g *GlowLayer) SetOnAfterComposeObservable(onAfterComposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onAfterComposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnAfterRenderMeshToEffect returns the OnAfterRenderMeshToEffect property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onafterrendermeshtoeffect
func (g *GlowLayer) OnAfterRenderMeshToEffect(onAfterRenderMeshToEffect *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onAfterRenderMeshToEffect.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderMeshToEffect sets the OnAfterRenderMeshToEffect property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onafterrendermeshtoeffect
func (g *GlowLayer) SetOnAfterRenderMeshToEffect(onAfterRenderMeshToEffect *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onAfterRenderMeshToEffect.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnBeforeComposeObservable returns the OnBeforeComposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforecomposeobservable
func (g *GlowLayer) OnBeforeComposeObservable(onBeforeComposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeComposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnBeforeComposeObservable sets the OnBeforeComposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforecomposeobservable
func (g *GlowLayer) SetOnBeforeComposeObservable(onBeforeComposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeComposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnBeforeRenderMainTextureObservable returns the OnBeforeRenderMainTextureObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforerendermaintextureobservable
func (g *GlowLayer) OnBeforeRenderMainTextureObservable(onBeforeRenderMainTextureObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeRenderMainTextureObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderMainTextureObservable sets the OnBeforeRenderMainTextureObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforerendermaintextureobservable
func (g *GlowLayer) SetOnBeforeRenderMainTextureObservable(onBeforeRenderMainTextureObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeRenderMainTextureObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnBeforeRenderMeshToEffect returns the OnBeforeRenderMeshToEffect property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforerendermeshtoeffect
func (g *GlowLayer) OnBeforeRenderMeshToEffect(onBeforeRenderMeshToEffect *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeRenderMeshToEffect.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderMeshToEffect sets the OnBeforeRenderMeshToEffect property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onbeforerendermeshtoeffect
func (g *GlowLayer) SetOnBeforeRenderMeshToEffect(onBeforeRenderMeshToEffect *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onBeforeRenderMeshToEffect.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnDisposeObservable returns the OnDisposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#ondisposeobservable
func (g *GlowLayer) OnDisposeObservable(onDisposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onDisposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#ondisposeobservable
func (g *GlowLayer) SetOnDisposeObservable(onDisposeObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onDisposeObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onsizechangedobservable
func (g *GlowLayer) OnSizeChangedObservable(onSizeChangedObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onSizeChangedObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#onsizechangedobservable
func (g *GlowLayer) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(onSizeChangedObservable.JSObject())
	return GlowLayerFromJSObject(p, ba.ctx)
}

// RenderingGroupId returns the RenderingGroupId property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#renderinggroupid
func (g *GlowLayer) RenderingGroupId(renderingGroupId float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(renderingGroupId)
	return GlowLayerFromJSObject(p, ba.ctx)
}

// SetRenderingGroupId sets the RenderingGroupId property of class GlowLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer#renderinggroupid
func (g *GlowLayer) SetRenderingGroupId(renderingGroupId float64) *GlowLayer {
	p := ba.ctx.Get("GlowLayer").New(renderingGroupId)
	return GlowLayerFromJSObject(p, ba.ctx)
}

*/
