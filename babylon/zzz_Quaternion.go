// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Quaternion represents a babylon.js Quaternion.
// Class used to store quaternion data
//
// See: http://doc.babylonjs.com/features/position,_rotation,_scaling
type Quaternion struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (q *Quaternion) JSObject() js.Value { return q.p }

// Quaternion returns a Quaternion JavaScript class.
func (b *Babylon) Quaternion() *Quaternion {
	p := b.ctx.Get("Quaternion")
	return QuaternionFromJSObject(p)
}

// QuaternionFromJSObject returns a wrapped Quaternion JavaScript class.
func QuaternionFromJSObject(p js.Value) *Quaternion {
	return &Quaternion{p: p}
}

// NewQuaternionOpts contains optional parameters for NewQuaternion.
type NewQuaternionOpts struct {
	X *float64

	Y *float64

	Z *float64

	W *float64
}

// NewQuaternion returns a new Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion
func (b *Babylon) NewQuaternion(opts *NewQuaternionOpts) *Quaternion {
	if opts == nil {
		opts = &NewQuaternionOpts{}
	}

	p := b.ctx.Get("Quaternion").New(opts.X, opts.Y, opts.Z, opts.W)
	return QuaternionFromJSObject(p)
}

// TODO: methods
