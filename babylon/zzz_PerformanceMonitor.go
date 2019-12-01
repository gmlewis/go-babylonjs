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
	FrameSampleSize *float64
}

// NewPerformanceMonitor returns a new PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor
func (ba *Babylon) NewPerformanceMonitor(opts *NewPerformanceMonitorOpts) *PerformanceMonitor {
	if opts == nil {
		opts = &NewPerformanceMonitorOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.FrameSampleSize == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FrameSampleSize)
	}

	p := ba.ctx.Get("PerformanceMonitor").New(args...)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// Disable calls the Disable method on the PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#disable
func (p *PerformanceMonitor) Disable() {

	args := make([]interface{}, 0, 0+0)

	p.p.Call("disable", args...)
}

// Enable calls the Enable method on the PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#enable
func (p *PerformanceMonitor) Enable() {

	args := make([]interface{}, 0, 0+0)

	p.p.Call("enable", args...)
}

// Reset calls the Reset method on the PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#reset
func (p *PerformanceMonitor) Reset() {

	args := make([]interface{}, 0, 0+0)

	p.p.Call("reset", args...)
}

// PerformanceMonitorSampleFrameOpts contains optional parameters for PerformanceMonitor.SampleFrame.
type PerformanceMonitorSampleFrameOpts struct {
	TimeMs *float64
}

// SampleFrame calls the SampleFrame method on the PerformanceMonitor object.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#sampleframe
func (p *PerformanceMonitor) SampleFrame(opts *PerformanceMonitorSampleFrameOpts) {
	if opts == nil {
		opts = &PerformanceMonitorSampleFrameOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.TimeMs == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TimeMs)
	}

	p.p.Call("sampleFrame", args...)
}

/*

// AverageFPS returns the AverageFPS property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averagefps
func (p *PerformanceMonitor) AverageFPS(averageFPS float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFPS)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetAverageFPS sets the AverageFPS property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averagefps
func (p *PerformanceMonitor) SetAverageFPS(averageFPS float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFPS)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// AverageFrameTime returns the AverageFrameTime property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averageframetime
func (p *PerformanceMonitor) AverageFrameTime(averageFrameTime float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFrameTime)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetAverageFrameTime sets the AverageFrameTime property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averageframetime
func (p *PerformanceMonitor) SetAverageFrameTime(averageFrameTime float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFrameTime)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// AverageFrameTimeVariance returns the AverageFrameTimeVariance property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averageframetimevariance
func (p *PerformanceMonitor) AverageFrameTimeVariance(averageFrameTimeVariance float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFrameTimeVariance)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetAverageFrameTimeVariance sets the AverageFrameTimeVariance property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#averageframetimevariance
func (p *PerformanceMonitor) SetAverageFrameTimeVariance(averageFrameTimeVariance float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(averageFrameTimeVariance)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// InstantaneousFPS returns the InstantaneousFPS property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#instantaneousfps
func (p *PerformanceMonitor) InstantaneousFPS(instantaneousFPS float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(instantaneousFPS)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetInstantaneousFPS sets the InstantaneousFPS property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#instantaneousfps
func (p *PerformanceMonitor) SetInstantaneousFPS(instantaneousFPS float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(instantaneousFPS)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// InstantaneousFrameTime returns the InstantaneousFrameTime property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#instantaneousframetime
func (p *PerformanceMonitor) InstantaneousFrameTime(instantaneousFrameTime float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(instantaneousFrameTime)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetInstantaneousFrameTime sets the InstantaneousFrameTime property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#instantaneousframetime
func (p *PerformanceMonitor) SetInstantaneousFrameTime(instantaneousFrameTime float64) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(instantaneousFrameTime)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// IsEnabled returns the IsEnabled property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#isenabled
func (p *PerformanceMonitor) IsEnabled(isEnabled bool) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(isEnabled)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetIsEnabled sets the IsEnabled property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#isenabled
func (p *PerformanceMonitor) SetIsEnabled(isEnabled bool) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(isEnabled)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// IsSaturated returns the IsSaturated property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#issaturated
func (p *PerformanceMonitor) IsSaturated(isSaturated bool) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(isSaturated)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

// SetIsSaturated sets the IsSaturated property of class PerformanceMonitor.
//
// https://doc.babylonjs.com/api/classes/babylon.performancemonitor#issaturated
func (p *PerformanceMonitor) SetIsSaturated(isSaturated bool) *PerformanceMonitor {
	p := ba.ctx.Get("PerformanceMonitor").New(isSaturated)
	return PerformanceMonitorFromJSObject(p, ba.ctx)
}

*/
