// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SineEase represents a babylon.js SineEase.
// Easing function with a sin shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type SineEase struct{ *EasingFunction }

// JSObject returns the underlying js.Value.
func (s *SineEase) JSObject() js.Value { return s.p }

// SineEase returns a SineEase JavaScript class.
func (b *Babylon) SineEase() *SineEase {
	p := b.ctx.Get("SineEase")
	return SineEaseFromJSObject(p)
}

// SineEaseFromJSObject returns a wrapped SineEase JavaScript class.
func SineEaseFromJSObject(p js.Value) *SineEase {
	return &SineEase{EasingFunctionFromJSObject(p)}
}

// TODO: methods
