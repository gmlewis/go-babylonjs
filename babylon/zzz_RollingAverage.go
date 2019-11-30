// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RollingAverage represents a babylon.js RollingAverage.
// RollingAverage
//
// Utility to efficiently compute the rolling average and variance over a sliding window of samples
type RollingAverage struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (r *RollingAverage) JSObject() js.Value { return r.p }

// RollingAverage returns a RollingAverage JavaScript class.
func (b *Babylon) RollingAverage() *RollingAverage {
	p := b.ctx.Get("RollingAverage")
	return RollingAverageFromJSObject(p)
}

// RollingAverageFromJSObject returns a wrapped RollingAverage JavaScript class.
func RollingAverageFromJSObject(p js.Value) *RollingAverage {
	return &RollingAverage{p: p}
}

// NewRollingAverage returns a new RollingAverage object.
//
// https://doc.babylonjs.com/api/classes/babylon.rollingaverage
func (b *Babylon) NewRollingAverage(length float64) *RollingAverage {
	p := b.ctx.Get("RollingAverage").New(length)
	return RollingAverageFromJSObject(p)
}

// TODO: methods
