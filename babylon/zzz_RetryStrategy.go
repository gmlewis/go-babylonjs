// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RetryStrategy represents a babylon.js RetryStrategy.
// Class used to define a retry strategy when error happens while loading assets
type RetryStrategy struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RetryStrategy) JSObject() js.Value { return r.p }

// RetryStrategy returns a RetryStrategy JavaScript class.
func (ba *Babylon) RetryStrategy() *RetryStrategy {
	p := ba.ctx.Get("RetryStrategy")
	return RetryStrategyFromJSObject(p, ba.ctx)
}

// RetryStrategyFromJSObject returns a wrapped RetryStrategy JavaScript class.
func RetryStrategyFromJSObject(p js.Value, ctx js.Value) *RetryStrategy {
	return &RetryStrategy{p: p, ctx: ctx}
}

// RetryStrategyExponentialBackoffOpts contains optional parameters for RetryStrategy.ExponentialBackoff.
type RetryStrategyExponentialBackoffOpts struct {
	MaxRetries   *float64
	BaseInterval *float64
}

// ExponentialBackoff calls the ExponentialBackoff method on the RetryStrategy object.
//
// https://doc.babylonjs.com/api/classes/babylon.retrystrategy#exponentialbackoff
func (r *RetryStrategy) ExponentialBackoff(opts *RetryStrategyExponentialBackoffOpts) func() {
	if opts == nil {
		opts = &RetryStrategyExponentialBackoffOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.MaxRetries == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxRetries)
	}
	if opts.BaseInterval == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BaseInterval)
	}

	retVal := r.p.Call("ExponentialBackoff", args...)
	return retVal
}

/*

 */
