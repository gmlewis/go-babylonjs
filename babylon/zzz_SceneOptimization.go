// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneOptimization represents a babylon.js SceneOptimization.
// Defines the root class used to create scene optimization to use with SceneOptimizer

//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type SceneOptimization struct{}

// JSObject returns the underlying js.Value.
func (s *SceneOptimization) JSObject() js.Value { return s.p }

// SceneOptimization returns a SceneOptimization JavaScript class.
func (b *Babylon) SceneOptimization() *SceneOptimization {
	p := b.ctx.Get("SceneOptimization")
	return SceneOptimizationFromJSObject(p)
}

// SceneOptimizationFromJSObject returns a wrapped SceneOptimization JavaScript class.
func SceneOptimizationFromJSObject(p js.Value) *SceneOptimization {
	return &SceneOptimization{p: p}
}

// NewSceneOptimization returns a new SceneOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimization
func (b *Babylon) NewSceneOptimization(todo parameters) *SceneOptimization {
	p := b.ctx.Get("SceneOptimization").New(todo)
	return SceneOptimizationFromJSObject(p)
}

// TODO: methods
