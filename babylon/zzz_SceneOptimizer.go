// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneOptimizer represents a babylon.js SceneOptimizer.
// Class used to run optimizations in order to reach a target frame rate

//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type SceneOptimizer struct{}

// JSObject returns the underlying js.Value.
func (s *SceneOptimizer) JSObject() js.Value { return s.p }

// SceneOptimizer returns a SceneOptimizer JavaScript class.
func (b *Babylon) SceneOptimizer() *SceneOptimizer {
	p := b.ctx.Get("SceneOptimizer")
	return SceneOptimizerFromJSObject(p)
}

// SceneOptimizerFromJSObject returns a wrapped SceneOptimizer JavaScript class.
func SceneOptimizerFromJSObject(p js.Value) *SceneOptimizer {
	return &SceneOptimizer{p: p}
}

// NewSceneOptimizer returns a new SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer
func (b *Babylon) NewSceneOptimizer(todo parameters) *SceneOptimizer {
	p := b.ctx.Get("SceneOptimizer").New(todo)
	return SceneOptimizerFromJSObject(p)
}

// TODO: methods