// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RetryStrategy represents a babylon.js RetryStrategy.
// Class used to define a retry strategy when error happens while loading assets
type RetryStrategy struct{}

// JSObject returns the underlying js.Value.
func (r *RetryStrategy) JSObject() js.Value { return r.p }

// RetryStrategy returns a RetryStrategy JavaScript class.
func (b *Babylon) RetryStrategy() *RetryStrategy {
	p := b.ctx.Get("RetryStrategy")
	return RetryStrategyFromJSObject(p)
}

// RetryStrategyFromJSObject returns a wrapped RetryStrategy JavaScript class.
func RetryStrategyFromJSObject(p js.Value) *RetryStrategy {
	return &RetryStrategy{p: p}
}

// NewRetryStrategy returns a new RetryStrategy object.
//
// https://doc.babylonjs.com/api/classes/babylon.retrystrategy
func (b *Babylon) NewRetryStrategy(todo parameters) *RetryStrategy {
	p := b.ctx.Get("RetryStrategy").New(todo)
	return RetryStrategyFromJSObject(p)
}

// TODO: methods
