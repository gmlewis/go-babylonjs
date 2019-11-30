// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Scene represents a babylon.js Scene.
// Represents a scene to be rendered by the engine.
//
// See: http://doc.babylonjs.com/features/scene
type Scene struct {
	*AbstractScene
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Scene) JSObject() js.Value { return s.p }

// Scene returns a Scene JavaScript class.
func (ba *Babylon) Scene() *Scene {
	p := ba.ctx.Get("Scene")
	return SceneFromJSObject(p, ba.ctx)
}

// SceneFromJSObject returns a wrapped Scene JavaScript class.
func SceneFromJSObject(p js.Value, ctx js.Value) *Scene {
	return &Scene{AbstractScene: AbstractSceneFromJSObject(p, ctx), ctx: ctx}
}

// NewSceneOpts contains optional parameters for NewScene.
type NewSceneOpts struct {
	Options *JSValue
}

// NewScene returns a new Scene object.
//
// https://doc.babylonjs.com/api/classes/babylon.scene
func (ba *Babylon) NewScene(engine *Engine, opts *NewSceneOpts) *Scene {
	if opts == nil {
		opts = &NewSceneOpts{}
	}

	p := ba.ctx.Get("Scene").New(engine.JSObject(), opts.Options.JSObject())
	return SceneFromJSObject(p, ba.ctx)
}

// TODO: methods
