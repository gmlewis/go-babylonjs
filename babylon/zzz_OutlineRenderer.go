// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OutlineRenderer represents a babylon.js OutlineRenderer.
// This class is responsible to draw bothe outline/overlay of meshes.
// It should not be used directly but through the available method on mesh.
type OutlineRenderer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (o *OutlineRenderer) JSObject() js.Value { return o.p }

// OutlineRenderer returns a OutlineRenderer JavaScript class.
func (b *Babylon) OutlineRenderer() *OutlineRenderer {
	p := b.ctx.Get("OutlineRenderer")
	return OutlineRendererFromJSObject(p)
}

// OutlineRendererFromJSObject returns a wrapped OutlineRenderer JavaScript class.
func OutlineRendererFromJSObject(p js.Value) *OutlineRenderer {
	return &OutlineRenderer{p: p}
}

// NewOutlineRenderer returns a new OutlineRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.outlinerenderer
func (b *Babylon) NewOutlineRenderer(scene *Scene) *OutlineRenderer {
	p := b.ctx.Get("OutlineRenderer").New(scene.JSObject())
	return OutlineRendererFromJSObject(p)
}

// TODO: methods
