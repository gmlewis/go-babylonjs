// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Scene represents a babylon.js Scene.
// Represents a scene to be rendered by the engine.

//
// See: http://doc.babylonjs.com/features/scene
type Scene struct{ *AbstractScene }

// JSObject returns the underlying js.Value.
func (s *Scene) JSObject() js.Value { return s.p }

// Scene returns a Scene JavaScript class.
func (b *Babylon) Scene() *Scene {
	p := b.ctx.Get("Scene")
	return SceneFromJSObject(p)
}

// SceneFromJSObject returns a wrapped Scene JavaScript class.
func SceneFromJSObject(p js.Value) *Scene {
	return &Scene{AbstractSceneFromJSObject(p)}
}

// NewScene returns a new Scene object.
//
// https://doc.babylonjs.com/api/classes/babylon.scene
func (b *Babylon) NewScene(todo parameters) *Scene {
	p := b.ctx.Get("Scene").New(todo)
	return SceneFromJSObject(p)
}

// TODO: methods