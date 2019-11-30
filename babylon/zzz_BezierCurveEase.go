// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BezierCurveEase represents a babylon.js BezierCurveEase.
// Easing function with a bezier shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type BezierCurveEase struct{ *EasingFunction }

// JSObject returns the underlying js.Value.
func (b *BezierCurveEase) JSObject() js.Value { return b.p }

// BezierCurveEase returns a BezierCurveEase JavaScript class.
func (b *Babylon) BezierCurveEase() *BezierCurveEase {
	p := b.ctx.Get("BezierCurveEase")
	return BezierCurveEaseFromJSObject(p)
}

// BezierCurveEaseFromJSObject returns a wrapped BezierCurveEase JavaScript class.
func BezierCurveEaseFromJSObject(p js.Value) *BezierCurveEase {
	return &BezierCurveEase{EasingFunctionFromJSObject(p)}
}

// NewBezierCurveEaseOpts contains optional parameters for NewBezierCurveEase.
type NewBezierCurveEaseOpts struct {
	X1 *float64

	Y1 *float64

	X2 *float64

	Y2 *float64
}

// NewBezierCurveEase returns a new BezierCurveEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.beziercurveease
func (b *Babylon) NewBezierCurveEase(opts *NewBezierCurveEaseOpts) *BezierCurveEase {
	if opts == nil {
		opts = &NewBezierCurveEaseOpts{}
	}

	p := b.ctx.Get("BezierCurveEase").New(opts.X1, opts.Y1, opts.X2, opts.Y2)
	return BezierCurveEaseFromJSObject(p)
}

// TODO: methods
