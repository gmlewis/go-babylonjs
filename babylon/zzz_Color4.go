// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Color4 represents a babylon.js Color4.
// Class used to hold a RBGA color
type Color4 struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Color4) JSObject() js.Value { return c.p }

// Color4 returns a Color4 JavaScript class.
func (ba *Babylon) Color4() *Color4 {
	p := ba.ctx.Get("Color4")
	return Color4FromJSObject(p, ba.ctx)
}

// Color4FromJSObject returns a wrapped Color4 JavaScript class.
func Color4FromJSObject(p js.Value, ctx js.Value) *Color4 {
	return &Color4{p: p, ctx: ctx}
}

// NewColor4Opts contains optional parameters for NewColor4.
type NewColor4Opts struct {
	R *JSFloat64

	G *JSFloat64

	B *JSFloat64

	A *JSFloat64
}

// NewColor4 returns a new Color4 object.
//
// https://doc.babylonjs.com/api/classes/babylon.color4
func (ba *Babylon) NewColor4(opts *NewColor4Opts) *Color4 {
	if opts == nil {
		opts = &NewColor4Opts{}
	}

	p := ba.ctx.Get("Color4").New(opts.R.JSObject(), opts.G.JSObject(), opts.B.JSObject(), opts.A.JSObject())
	return Color4FromJSObject(p, ba.ctx)
}

// TODO: methods
