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

// TODO: methods
