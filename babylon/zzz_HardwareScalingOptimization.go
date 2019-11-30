// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HardwareScalingOptimization represents a babylon.js HardwareScalingOptimization.
// Defines an optimization used to increase or decrease the rendering resolution
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type HardwareScalingOptimization struct{ *SceneOptimization }

// JSObject returns the underlying js.Value.
func (h *HardwareScalingOptimization) JSObject() js.Value { return h.p }

// HardwareScalingOptimization returns a HardwareScalingOptimization JavaScript class.
func (b *Babylon) HardwareScalingOptimization() *HardwareScalingOptimization {
	p := b.ctx.Get("HardwareScalingOptimization")
	return HardwareScalingOptimizationFromJSObject(p)
}

// HardwareScalingOptimizationFromJSObject returns a wrapped HardwareScalingOptimization JavaScript class.
func HardwareScalingOptimizationFromJSObject(p js.Value) *HardwareScalingOptimization {
	return &HardwareScalingOptimization{SceneOptimizationFromJSObject(p)}
}

// NewHardwareScalingOptimizationOpts contains optional parameters for NewHardwareScalingOptimization.
type NewHardwareScalingOptimizationOpts struct {
	Priority *float64

	MaximumScale *float64

	Step *float64
}

// NewHardwareScalingOptimization returns a new HardwareScalingOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization
func (b *Babylon) NewHardwareScalingOptimization(opts *NewHardwareScalingOptimizationOpts) *HardwareScalingOptimization {
	if opts == nil {
		opts = &NewHardwareScalingOptimizationOpts{}
	}

	p := b.ctx.Get("HardwareScalingOptimization").New(opts.Priority, opts.MaximumScale, opts.Step)
	return HardwareScalingOptimizationFromJSObject(p)
}

// TODO: methods
