// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderingManager represents a babylon.js RenderingManager.
// This is the manager responsible of all the rendering for meshes sprites and particles.
// It is enable to manage the different groups as well as the different necessary sort functions.
// This should not be used directly aside of the few static configurations
type RenderingManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (r *RenderingManager) JSObject() js.Value { return r.p }

// RenderingManager returns a RenderingManager JavaScript class.
func (b *Babylon) RenderingManager() *RenderingManager {
	p := b.ctx.Get("RenderingManager")
	return RenderingManagerFromJSObject(p)
}

// RenderingManagerFromJSObject returns a wrapped RenderingManager JavaScript class.
func RenderingManagerFromJSObject(p js.Value) *RenderingManager {
	return &RenderingManager{p: p}
}

// NewRenderingManager returns a new RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager
func (b *Babylon) NewRenderingManager(scene *Scene) *RenderingManager {
	p := b.ctx.Get("RenderingManager").New(scene.JSObject())
	return RenderingManagerFromJSObject(p)
}

// TODO: methods
