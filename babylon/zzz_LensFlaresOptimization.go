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

// LensFlaresOptimizationArrayToJSArray returns a JavaScript Array for the wrapped array.
func LensFlaresOptimizationArrayToJSArray(array []*LensFlaresOptimization) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLensFlaresOptimizationOpts contains optional parameters for NewLensFlaresOptimization.
type NewLensFlaresOptimizationOpts struct {
	Priority *float64
}

// NewLensFlaresOptimization returns a new LensFlaresOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresoptimization#constructor
func (ba *Babylon) NewLensFlaresOptimization(opts *NewLensFlaresOptimizationOpts) *LensFlaresOptimization {
	if opts == nil {
		opts = &NewLensFlaresOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}

	p := ba.ctx.Get("LensFlaresOptimization").New(args...)
	return LensFlaresOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the LensFlaresOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresoptimization#apply
func (l *LensFlaresOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

	args := make([]interface{}, 0, 2+0)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if optimizer == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, optimizer.JSObject())
	}

	retVal := l.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the LensFlaresOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresoptimization#getdescription
func (l *LensFlaresOptimization) GetDescription() string {

	retVal := l.p.Call("getDescription")
	return retVal.String()
}
