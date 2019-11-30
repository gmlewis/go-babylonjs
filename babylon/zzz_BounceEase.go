// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BounceEase represents a babylon.js BounceEase.
// Easing function with a bouncing shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type BounceEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BounceEase) JSObject() js.Value { return b.p }

// BounceEase returns a BounceEase JavaScript class.
func (ba *Babylon) BounceEase() *BounceEase {
	p := ba.ctx.Get("BounceEase")
	return BounceEaseFromJSObject(p, ba.ctx)
}

// BounceEaseFromJSObject returns a wrapped BounceEase JavaScript class.
func BounceEaseFromJSObject(p js.Value, ctx js.Value) *BounceEase {
	return &BounceEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// NewBounceEaseOpts contains optional parameters for NewBounceEase.
type NewBounceEaseOpts struct {
	Bounces *JSFloat64

	Bounciness *JSFloat64
}

// NewBounceEase returns a new BounceEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.bounceease
func (ba *Babylon) NewBounceEase(opts *NewBounceEaseOpts) *BounceEase {
	if opts == nil {
		opts = &NewBounceEaseOpts{}
	}

	p := ba.ctx.Get("BounceEase").New(opts.Bounces.JSObject(), opts.Bounciness.JSObject())
	return BounceEaseFromJSObject(p, ba.ctx)
}

// TODO: methods
