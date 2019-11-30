// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SimplicationQueueSceneComponent represents a babylon.js SimplicationQueueSceneComponent.
// Defines the simplification queue scene component responsible to help scheduling the various simplification task
// created in a scene
type SimplicationQueueSceneComponent struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SimplicationQueueSceneComponent) JSObject() js.Value { return s.p }

// SimplicationQueueSceneComponent returns a SimplicationQueueSceneComponent JavaScript class.
func (b *Babylon) SimplicationQueueSceneComponent() *SimplicationQueueSceneComponent {
	p := b.ctx.Get("SimplicationQueueSceneComponent")
	return SimplicationQueueSceneComponentFromJSObject(p)
}

// SimplicationQueueSceneComponentFromJSObject returns a wrapped SimplicationQueueSceneComponent JavaScript class.
func SimplicationQueueSceneComponentFromJSObject(p js.Value) *SimplicationQueueSceneComponent {
	return &SimplicationQueueSceneComponent{p: p}
}

// NewSimplicationQueueSceneComponent returns a new SimplicationQueueSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplicationqueuescenecomponent
func (b *Babylon) NewSimplicationQueueSceneComponent(scene *Scene) *SimplicationQueueSceneComponent {
	p := b.ctx.Get("SimplicationQueueSceneComponent").New(scene.JSObject())
	return SimplicationQueueSceneComponentFromJSObject(p)
}

// TODO: methods
