// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GeometryBufferRenderer represents a babylon.js GeometryBufferRenderer.
// This renderer is helpfull to fill one of the render target with a geometry buffer.
type GeometryBufferRenderer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GeometryBufferRenderer) JSObject() js.Value { return g.p }

// GeometryBufferRenderer returns a GeometryBufferRenderer JavaScript class.
func (ba *Babylon) GeometryBufferRenderer() *GeometryBufferRenderer {
	p := ba.ctx.Get("GeometryBufferRenderer")
	return GeometryBufferRendererFromJSObject(p, ba.ctx)
}

// GeometryBufferRendererFromJSObject returns a wrapped GeometryBufferRenderer JavaScript class.
func GeometryBufferRendererFromJSObject(p js.Value, ctx js.Value) *GeometryBufferRenderer {
	return &GeometryBufferRenderer{p: p, ctx: ctx}
}

// GeometryBufferRendererArrayToJSArray returns a JavaScript Array for the wrapped array.
func GeometryBufferRendererArrayToJSArray(array []*GeometryBufferRenderer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGeometryBufferRendererOpts contains optional parameters for NewGeometryBufferRenderer.
type NewGeometryBufferRendererOpts struct {
	Ratio *float64
}

// NewGeometryBufferRenderer returns a new GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer
func (ba *Babylon) NewGeometryBufferRenderer(scene *Scene, opts *NewGeometryBufferRendererOpts) *GeometryBufferRenderer {
	if opts == nil {
		opts = &NewGeometryBufferRendererOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.Ratio == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Ratio)
	}

	p := ba.ctx.Get("GeometryBufferRenderer").New(args...)
	return GeometryBufferRendererFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#dispose
func (g *GeometryBufferRenderer) Dispose() {

	g.p.Call("dispose")
}

// GetGBuffer calls the GetGBuffer method on the GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#getgbuffer
func (g *GeometryBufferRenderer) GetGBuffer() *MultiRenderTarget {

	retVal := g.p.Call("getGBuffer")
	return MultiRenderTargetFromJSObject(retVal, g.ctx)
}

// GetTextureIndex calls the GetTextureIndex method on the GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#gettextureindex
func (g *GeometryBufferRenderer) GetTextureIndex(textureType float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, textureType)

	retVal := g.p.Call("getTextureIndex", args...)
	return retVal.Float()
}

// IsReady calls the IsReady method on the GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#isready
func (g *GeometryBufferRenderer) IsReady(subMesh *SubMesh, useInstances bool) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, subMesh.JSObject())
	args = append(args, useInstances)

	retVal := g.p.Call("isReady", args...)
	return retVal.Bool()
}

// EnablePosition returns the EnablePosition property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#enableposition
func (g *GeometryBufferRenderer) EnablePosition() bool {
	retVal := g.p.Get("enablePosition")
	return retVal.Bool()
}

// SetEnablePosition sets the EnablePosition property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#enableposition
func (g *GeometryBufferRenderer) SetEnablePosition(enablePosition bool) *GeometryBufferRenderer {
	g.p.Set("enablePosition", enablePosition)
	return g
}

// EnableVelocity returns the EnableVelocity property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#enablevelocity
func (g *GeometryBufferRenderer) EnableVelocity() bool {
	retVal := g.p.Get("enableVelocity")
	return retVal.Bool()
}

// SetEnableVelocity sets the EnableVelocity property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#enablevelocity
func (g *GeometryBufferRenderer) SetEnableVelocity(enableVelocity bool) *GeometryBufferRenderer {
	g.p.Set("enableVelocity", enableVelocity)
	return g
}

// ExcludedSkinnedMeshesFromVelocity returns the ExcludedSkinnedMeshesFromVelocity property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#excludedskinnedmeshesfromvelocity
func (g *GeometryBufferRenderer) ExcludedSkinnedMeshesFromVelocity() []*AbstractMesh {
	retVal := g.p.Get("excludedSkinnedMeshesFromVelocity")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), g.ctx))
	}
	return result
}

// SetExcludedSkinnedMeshesFromVelocity sets the ExcludedSkinnedMeshesFromVelocity property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#excludedskinnedmeshesfromvelocity
func (g *GeometryBufferRenderer) SetExcludedSkinnedMeshesFromVelocity(excludedSkinnedMeshesFromVelocity []*AbstractMesh) *GeometryBufferRenderer {
	g.p.Set("excludedSkinnedMeshesFromVelocity", excludedSkinnedMeshesFromVelocity)
	return g
}

// IsSupported returns the IsSupported property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#issupported
func (g *GeometryBufferRenderer) IsSupported() bool {
	retVal := g.p.Get("isSupported")
	return retVal.Bool()
}

// SetIsSupported sets the IsSupported property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#issupported
func (g *GeometryBufferRenderer) SetIsSupported(isSupported bool) *GeometryBufferRenderer {
	g.p.Set("isSupported", isSupported)
	return g
}

// POSITION_TEXTURE_TYPE returns the POSITION_TEXTURE_TYPE property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#position_texture_type
func (g *GeometryBufferRenderer) POSITION_TEXTURE_TYPE() float64 {
	retVal := g.p.Get("POSITION_TEXTURE_TYPE")
	return retVal.Float()
}

// SetPOSITION_TEXTURE_TYPE sets the POSITION_TEXTURE_TYPE property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#position_texture_type
func (g *GeometryBufferRenderer) SetPOSITION_TEXTURE_TYPE(POSITION_TEXTURE_TYPE float64) *GeometryBufferRenderer {
	g.p.Set("POSITION_TEXTURE_TYPE", POSITION_TEXTURE_TYPE)
	return g
}

// Ratio returns the Ratio property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#ratio
func (g *GeometryBufferRenderer) Ratio() float64 {
	retVal := g.p.Get("ratio")
	return retVal.Float()
}

// SetRatio sets the Ratio property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#ratio
func (g *GeometryBufferRenderer) SetRatio(ratio float64) *GeometryBufferRenderer {
	g.p.Set("ratio", ratio)
	return g
}

// RenderList returns the RenderList property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#renderlist
func (g *GeometryBufferRenderer) RenderList() []*Mesh {
	retVal := g.p.Get("renderList")
	result := []*Mesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, MeshFromJSObject(retVal.Index(ri), g.ctx))
	}
	return result
}

// SetRenderList sets the RenderList property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#renderlist
func (g *GeometryBufferRenderer) SetRenderList(renderList []*Mesh) *GeometryBufferRenderer {
	g.p.Set("renderList", renderList)
	return g
}

// Samples returns the Samples property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#samples
func (g *GeometryBufferRenderer) Samples() float64 {
	retVal := g.p.Get("samples")
	return retVal.Float()
}

// SetSamples sets the Samples property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#samples
func (g *GeometryBufferRenderer) SetSamples(samples float64) *GeometryBufferRenderer {
	g.p.Set("samples", samples)
	return g
}

// Scene returns the Scene property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#scene
func (g *GeometryBufferRenderer) Scene() *Scene {
	retVal := g.p.Get("scene")
	return SceneFromJSObject(retVal, g.ctx)
}

// SetScene sets the Scene property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#scene
func (g *GeometryBufferRenderer) SetScene(scene *Scene) *GeometryBufferRenderer {
	g.p.Set("scene", scene.JSObject())
	return g
}

// VELOCITY_TEXTURE_TYPE returns the VELOCITY_TEXTURE_TYPE property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#velocity_texture_type
func (g *GeometryBufferRenderer) VELOCITY_TEXTURE_TYPE() float64 {
	retVal := g.p.Get("VELOCITY_TEXTURE_TYPE")
	return retVal.Float()
}

// SetVELOCITY_TEXTURE_TYPE sets the VELOCITY_TEXTURE_TYPE property of class GeometryBufferRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer#velocity_texture_type
func (g *GeometryBufferRenderer) SetVELOCITY_TEXTURE_TYPE(VELOCITY_TEXTURE_TYPE float64) *GeometryBufferRenderer {
	g.p.Set("VELOCITY_TEXTURE_TYPE", VELOCITY_TEXTURE_TYPE)
	return g
}
