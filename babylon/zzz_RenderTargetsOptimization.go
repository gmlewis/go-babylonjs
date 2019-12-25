// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderTargetsOptimization represents a babylon.js RenderTargetsOptimization.
// Defines an optimization used to turn render targets off
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type RenderTargetsOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RenderTargetsOptimization) JSObject() js.Value { return r.p }

// RenderTargetsOptimization returns a RenderTargetsOptimization JavaScript class.
func (ba *Babylon) RenderTargetsOptimization() *RenderTargetsOptimization {
	p := ba.ctx.Get("RenderTargetsOptimization")
	return RenderTargetsOptimizationFromJSObject(p, ba.ctx)
}

// RenderTargetsOptimizationFromJSObject returns a wrapped RenderTargetsOptimization JavaScript class.
func RenderTargetsOptimizationFromJSObject(p js.Value, ctx js.Value) *RenderTargetsOptimization {
	return &RenderTargetsOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// RenderTargetsOptimizationArrayToJSArray returns a JavaScript Array for the wrapped array.
func RenderTargetsOptimizationArrayToJSArray(array []*RenderTargetsOptimization) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRenderTargetsOptimizationOpts contains optional parameters for NewRenderTargetsOptimization.
type NewRenderTargetsOptimizationOpts struct {
	Priority *float64
}

// NewRenderTargetsOptimization returns a new RenderTargetsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetsoptimization#constructor
func (ba *Babylon) NewRenderTargetsOptimization(opts *NewRenderTargetsOptimizationOpts) *RenderTargetsOptimization {
	if opts == nil {
		opts = &NewRenderTargetsOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}

	p := ba.ctx.Get("RenderTargetsOptimization").New(args...)
	return RenderTargetsOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the RenderTargetsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetsoptimization#apply
func (r *RenderTargetsOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

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

	retVal := r.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the RenderTargetsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargetsoptimization#getdescription
func (r *RenderTargetsOptimization) GetDescription() string {

	retVal := r.p.Call("getDescription")
	return retVal.String()
}
