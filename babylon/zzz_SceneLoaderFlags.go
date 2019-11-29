// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneLoaderFlags represents a babylon.js SceneLoaderFlags.
// Class used to represent data loading progression
type SceneLoaderFlags struct{}

// JSObject returns the underlying js.Value.
func (s *SceneLoaderFlags) JSObject() js.Value { return s.p }

// SceneLoaderFlags returns a SceneLoaderFlags JavaScript class.
func (b *Babylon) SceneLoaderFlags() *SceneLoaderFlags {
	p := b.ctx.Get("SceneLoaderFlags")
	return SceneLoaderFlagsFromJSObject(p)
}

// SceneLoaderFlagsFromJSObject returns a wrapped SceneLoaderFlags JavaScript class.
func SceneLoaderFlagsFromJSObject(p js.Value) *SceneLoaderFlags {
	return &SceneLoaderFlags{p: p}
}

// NewSceneLoaderFlags returns a new SceneLoaderFlags object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderflags
func (b *Babylon) NewSceneLoaderFlags(todo parameters) *SceneLoaderFlags {
	p := b.ctx.Get("SceneLoaderFlags").New(todo)
	return SceneLoaderFlagsFromJSObject(p)
}

// TODO: methods
