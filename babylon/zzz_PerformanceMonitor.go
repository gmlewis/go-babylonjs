// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PerformanceMonitor represents a babylon.js PerformanceMonitor.
// Performance monitor tracks rolling average frame-time and frame-time variance over a user defined sliding-window
type PerformanceMonitor struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PerformanceMonitor) JSObject() js.Value { return p.p }

// PerformanceMonitor returns a PerformanceMonitor JavaScript class.
func (ba *Babylon) PerformanceMonitor() *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor")
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// PerformanceMonitorFromJSObject returns a wrapped PerformanceMonitor JavaScript class.
func PerformanceMonitorFromJSObject(p js.Value, ctx js.Value) *PerformanceMonitor {
	return &PerformanceMonitor{p: p, ctx: ctx}
}

// NewPerformanceMonitorOpts contains optional parameters for NewPerformanceMonitor.
type NewPerformanceMonitorOpts struct {
	FrameSampleSize *JSFloat64
}

// NewPerformanceMonitor returns a new PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor
func (ba *Babylon) NewPerformanceMonitor(opts *NewPerformanceMonitorOpts) *PerformanceMonitor {
	if opts == nil {
		opts = &NewPerformanceMonitorOpts{}
	}

	p := ba.ctx.Get("PerformanceMonitor").New(opts.FrameSampleSize.JSObject())
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// TODO: methods
