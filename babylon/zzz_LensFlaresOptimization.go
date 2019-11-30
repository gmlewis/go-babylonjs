// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensFlaresOptimization represents a babylon.js LensFlaresOptimization.
// Defines an optimization used to turn lens flares off
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type LensFlaresOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LensFlaresOptimization) JSObject() js.Value { return l.p }

// LensFlaresOptimization returns a LensFlaresOptimization JavaScript class.
func (ba *Babylon) LensFlaresOptimization() *LensFlaresOptimization {
	p := ba.ctx.Get("LensFlaresOptimization")
	return LensFlaresOptimizationFromJSObject(p, ba.ctx)
}

// LensFlaresOptimizationFromJSObject returns a wrapped LensFlaresOptimization JavaScript class.
func LensFlaresOptimizationFromJSObject(p js.Value, ctx js.Value) *LensFlaresOptimization {
	return &LensFlaresOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// NewLensFlaresOptimizationOpts contains optional parameters for NewLensFlaresOptimization.
type NewLensFlaresOptimizationOpts struct {
	Priority *JSFloat64
}

// NewLensFlaresOptimization returns a new LensFlaresOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresoptimization
func (ba *Babylon) NewLensFlaresOptimization(opts *NewLensFlaresOptimizationOpts) *LensFlaresOptimization {
	if opts == nil {
		opts = &NewLensFlaresOptimizationOpts{}
	}

	p := ba.ctx.Get("LensFlaresOptimization").New(opts.Priority.JSObject())
	return LensFlaresOptimizationFromJSObject(p, ba.ctx)
}

// TODO: methods
