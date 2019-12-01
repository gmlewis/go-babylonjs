// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HardwareScalingOptimization represents a babylon.js HardwareScalingOptimization.
// Defines an optimization used to increase or decrease the rendering resolution
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type HardwareScalingOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HardwareScalingOptimization) JSObject() js.Value { return h.p }

// HardwareScalingOptimization returns a HardwareScalingOptimization JavaScript class.
func (ba *Babylon) HardwareScalingOptimization() *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization")
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// HardwareScalingOptimizationFromJSObject returns a wrapped HardwareScalingOptimization JavaScript class.
func HardwareScalingOptimizationFromJSObject(p js.Value, ctx js.Value) *HardwareScalingOptimization {
	return &HardwareScalingOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// HardwareScalingOptimizationArrayToJSArray returns a JavaScript Array for the wrapped array.
func HardwareScalingOptimizationArrayToJSArray(array []*HardwareScalingOptimization) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHardwareScalingOptimizationOpts contains optional parameters for NewHardwareScalingOptimization.
type NewHardwareScalingOptimizationOpts struct {
	Priority     *float64
	MaximumScale *float64
	Step         *float64
}

// NewHardwareScalingOptimization returns a new HardwareScalingOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization
func (ba *Babylon) NewHardwareScalingOptimization(opts *NewHardwareScalingOptimizationOpts) *HardwareScalingOptimization {
	if opts == nil {
		opts = &NewHardwareScalingOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}
	if opts.MaximumScale == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaximumScale)
	}
	if opts.Step == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Step)
	}

	p := ba.ctx.Get("HardwareScalingOptimization").New(args...)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the HardwareScalingOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#apply
func (h *HardwareScalingOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, optimizer.JSObject())

	retVal := h.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the HardwareScalingOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#getdescription
func (h *HardwareScalingOptimization) GetDescription() string {

	retVal := h.p.Call("getDescription")
	return retVal.String()
}

/*

// MaximumScale returns the MaximumScale property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#maximumscale
func (h *HardwareScalingOptimization) MaximumScale(maximumScale float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(maximumScale)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// SetMaximumScale sets the MaximumScale property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#maximumscale
func (h *HardwareScalingOptimization) SetMaximumScale(maximumScale float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(maximumScale)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// Priority returns the Priority property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#priority
func (h *HardwareScalingOptimization) Priority(priority float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(priority)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// SetPriority sets the Priority property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#priority
func (h *HardwareScalingOptimization) SetPriority(priority float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(priority)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// Step returns the Step property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#step
func (h *HardwareScalingOptimization) Step(step float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(step)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

// SetStep sets the Step property of class HardwareScalingOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization#step
func (h *HardwareScalingOptimization) SetStep(step float64) *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization").New(step)
	return HardwareScalingOptimizationFromJSObject(p, ba.ctx)
}

*/
