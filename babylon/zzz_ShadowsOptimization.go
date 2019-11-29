// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShadowsOptimization represents a babylon.js ShadowsOptimization.
// Defines an optimization used to remove shadows

//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type ShadowsOptimization struct{ *SceneOptimization }

// JSObject returns the underlying js.Value.
func (s *ShadowsOptimization) JSObject() js.Value { return s.p }

// ShadowsOptimization returns a ShadowsOptimization JavaScript class.
func (b *Babylon) ShadowsOptimization() *ShadowsOptimization {
	p := b.ctx.Get("ShadowsOptimization")
	return ShadowsOptimizationFromJSObject(p)
}

// ShadowsOptimizationFromJSObject returns a wrapped ShadowsOptimization JavaScript class.
func ShadowsOptimizationFromJSObject(p js.Value) *ShadowsOptimization {
	return &ShadowsOptimization{SceneOptimizationFromJSObject(p)}
}

// NewShadowsOptimization returns a new ShadowsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowsoptimization
func (b *Babylon) NewShadowsOptimization(todo parameters) *ShadowsOptimization {
	p := b.ctx.Get("ShadowsOptimization").New(todo)
	return ShadowsOptimizationFromJSObject(p)
}

// TODO: methods