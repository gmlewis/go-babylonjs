// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GeometryBufferRendererSceneComponent represents a babylon.js GeometryBufferRendererSceneComponent.
// Defines the Geometry Buffer scene component responsible to manage a G-Buffer useful
// in several rendering techniques.
type GeometryBufferRendererSceneComponent struct{}

// JSObject returns the underlying js.Value.
func (g *GeometryBufferRendererSceneComponent) JSObject() js.Value { return g.p }

// GeometryBufferRendererSceneComponent returns a GeometryBufferRendererSceneComponent JavaScript class.
func (b *Babylon) GeometryBufferRendererSceneComponent() *GeometryBufferRendererSceneComponent {
	p := b.ctx.Get("GeometryBufferRendererSceneComponent")
	return GeometryBufferRendererSceneComponentFromJSObject(p)
}

// GeometryBufferRendererSceneComponentFromJSObject returns a wrapped GeometryBufferRendererSceneComponent JavaScript class.
func GeometryBufferRendererSceneComponentFromJSObject(p js.Value) *GeometryBufferRendererSceneComponent {
	return &GeometryBufferRendererSceneComponent{p: p}
}

// NewGeometryBufferRendererSceneComponent returns a new GeometryBufferRendererSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrendererscenecomponent
func (b *Babylon) NewGeometryBufferRendererSceneComponent(todo parameters) *GeometryBufferRendererSceneComponent {
	p := b.ctx.Get("GeometryBufferRendererSceneComponent").New(todo)
	return GeometryBufferRendererSceneComponentFromJSObject(p)
}

// TODO: methods