// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ElasticEase represents a babylon.js ElasticEase.
// Easing function with an elastic shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type ElasticEase struct{ *EasingFunction }

// JSObject returns the underlying js.Value.
func (e *ElasticEase) JSObject() js.Value { return e.p }

// ElasticEase returns a ElasticEase JavaScript class.
func (b *Babylon) ElasticEase() *ElasticEase {
	p := b.ctx.Get("ElasticEase")
	return ElasticEaseFromJSObject(p)
}

// ElasticEaseFromJSObject returns a wrapped ElasticEase JavaScript class.
func ElasticEaseFromJSObject(p js.Value) *ElasticEase {
	return &ElasticEase{EasingFunctionFromJSObject(p)}
}

// NewElasticEaseOpts contains optional parameters for NewElasticEase.
type NewElasticEaseOpts struct {
	Oscillations *float64

	Springiness *float64
}

// NewElasticEase returns a new ElasticEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.elasticease
func (b *Babylon) NewElasticEase(opts *NewElasticEaseOpts) *ElasticEase {
	if opts == nil {
		opts = &NewElasticEaseOpts{}
	}

	p := b.ctx.Get("ElasticEase").New(opts.Oscillations, opts.Springiness)
	return ElasticEaseFromJSObject(p)
}

// TODO: methods
