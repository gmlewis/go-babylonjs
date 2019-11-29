// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthRendererSceneComponent represents a babylon.js DepthRendererSceneComponent.
// Defines the Depth Renderer scene component responsible to manage a depth buffer useful
// in several rendering techniques.
type DepthRendererSceneComponent struct{}

// JSObject returns the underlying js.Value.
func (d *DepthRendererSceneComponent) JSObject() js.Value { return d.p }

// DepthRendererSceneComponent returns a DepthRendererSceneComponent JavaScript class.
func (b *Babylon) DepthRendererSceneComponent() *DepthRendererSceneComponent {
	p := b.ctx.Get("DepthRendererSceneComponent")
	return DepthRendererSceneComponentFromJSObject(p)
}

// DepthRendererSceneComponentFromJSObject returns a wrapped DepthRendererSceneComponent JavaScript class.
func DepthRendererSceneComponentFromJSObject(p js.Value) *DepthRendererSceneComponent {
	return &DepthRendererSceneComponent{p: p}
}

// NewDepthRendererSceneComponent returns a new DepthRendererSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrendererscenecomponent
func (b *Babylon) NewDepthRendererSceneComponent(todo parameters) *DepthRendererSceneComponent {
	p := b.ctx.Get("DepthRendererSceneComponent").New(todo)
	return DepthRendererSceneComponentFromJSObject(p)
}

// TODO: methods