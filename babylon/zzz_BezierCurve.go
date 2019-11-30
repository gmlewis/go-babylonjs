// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BezierCurve represents a babylon.js BezierCurve.
// Class used to represent a Bezier curve
type BezierCurve struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BezierCurve) JSObject() js.Value { return b.p }

// BezierCurve returns a BezierCurve JavaScript class.
func (ba *Babylon) BezierCurve() *BezierCurve {
	p := ba.ctx.Get("BezierCurve")
	return BezierCurveFromJSObject(p, ba.ctx)
}

// BezierCurveFromJSObject returns a wrapped BezierCurve JavaScript class.
func BezierCurveFromJSObject(p js.Value, ctx js.Value) *BezierCurve {
	return &BezierCurve{p: p, ctx: ctx}
}

// TODO: methods
