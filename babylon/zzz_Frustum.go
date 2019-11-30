// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Frustum represents a babylon.js Frustum.
// Represents a camera frustum
type Frustum struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (f *Frustum) JSObject() js.Value { return f.p }

// Frustum returns a Frustum JavaScript class.
func (b *Babylon) Frustum() *Frustum {
	p := b.ctx.Get("Frustum")
	return FrustumFromJSObject(p)
}

// FrustumFromJSObject returns a wrapped Frustum JavaScript class.
func FrustumFromJSObject(p js.Value) *Frustum {
	return &Frustum{p: p}
}

// TODO: methods
