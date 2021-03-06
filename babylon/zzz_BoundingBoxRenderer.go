// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingBoxRenderer represents a babylon.js BoundingBoxRenderer.
// Component responsible of rendering the bounding box of the meshes in a scene.
// This is usually used through the mesh.showBoundingBox or the scene.forceShowBoundingBoxes properties
type BoundingBoxRenderer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BoundingBoxRenderer) JSObject() js.Value { return b.p }

// BoundingBoxRenderer returns a BoundingBoxRenderer JavaScript class.
func (ba *Babylon) BoundingBoxRenderer() *BoundingBoxRenderer {
	p := ba.ctx.Get("BoundingBoxRenderer")
	return BoundingBoxRendererFromJSObject(p, ba.ctx)
}

// BoundingBoxRendererFromJSObject returns a wrapped BoundingBoxRenderer JavaScript class.
func BoundingBoxRendererFromJSObject(p js.Value, ctx js.Value) *BoundingBoxRenderer {
	return &BoundingBoxRenderer{p: p, ctx: ctx}
}

// BoundingBoxRendererArrayToJSArray returns a JavaScript Array for the wrapped array.
func BoundingBoxRendererArrayToJSArray(array []*BoundingBoxRenderer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBoundingBoxRenderer returns a new BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#constructor
func (ba *Babylon) NewBoundingBoxRenderer(scene *Scene) *BoundingBoxRenderer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("BoundingBoxRenderer").New(args...)
	return BoundingBoxRendererFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#dispose
func (b *BoundingBoxRenderer) Dispose() {

	b.p.Call("dispose")
}

// Rebuild calls the Rebuild method on the BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#rebuild
func (b *BoundingBoxRenderer) Rebuild() {

	b.p.Call("rebuild")
}

// Register calls the Register method on the BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#register
func (b *BoundingBoxRenderer) Register() {

	b.p.Call("register")
}

// Render calls the Render method on the BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#render
func (b *BoundingBoxRenderer) Render(renderingGroupId float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, renderingGroupId)

	b.p.Call("render", args...)
}

// RenderOcclusionBoundingBox calls the RenderOcclusionBoundingBox method on the BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#renderocclusionboundingbox
func (b *BoundingBoxRenderer) RenderOcclusionBoundingBox(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	b.p.Call("renderOcclusionBoundingBox", args...)
}

// BackColor returns the BackColor property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#backcolor
func (b *BoundingBoxRenderer) BackColor() *Color3 {
	retVal := b.p.Get("backColor")
	return Color3FromJSObject(retVal, b.ctx)
}

// SetBackColor sets the BackColor property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#backcolor
func (b *BoundingBoxRenderer) SetBackColor(backColor *Color3) *BoundingBoxRenderer {
	b.p.Set("backColor", backColor.JSObject())
	return b
}

// FrontColor returns the FrontColor property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#frontcolor
func (b *BoundingBoxRenderer) FrontColor() *Color3 {
	retVal := b.p.Get("frontColor")
	return Color3FromJSObject(retVal, b.ctx)
}

// SetFrontColor sets the FrontColor property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#frontcolor
func (b *BoundingBoxRenderer) SetFrontColor(frontColor *Color3) *BoundingBoxRenderer {
	b.p.Set("frontColor", frontColor.JSObject())
	return b
}

// Name returns the Name property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#name
func (b *BoundingBoxRenderer) Name() string {
	retVal := b.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#name
func (b *BoundingBoxRenderer) SetName(name string) *BoundingBoxRenderer {
	b.p.Set("name", name)
	return b
}

// Scene returns the Scene property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#scene
func (b *BoundingBoxRenderer) Scene() *Scene {
	retVal := b.p.Get("scene")
	return SceneFromJSObject(retVal, b.ctx)
}

// SetScene sets the Scene property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#scene
func (b *BoundingBoxRenderer) SetScene(scene *Scene) *BoundingBoxRenderer {
	b.p.Set("scene", scene.JSObject())
	return b
}

// ShowBackLines returns the ShowBackLines property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#showbacklines
func (b *BoundingBoxRenderer) ShowBackLines() bool {
	retVal := b.p.Get("showBackLines")
	return retVal.Bool()
}

// SetShowBackLines sets the ShowBackLines property of class BoundingBoxRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer#showbacklines
func (b *BoundingBoxRenderer) SetShowBackLines(showBackLines bool) *BoundingBoxRenderer {
	b.p.Set("showBackLines", showBackLines)
	return b
}
