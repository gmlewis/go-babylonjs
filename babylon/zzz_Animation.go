// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Animation represents a babylon.js Animation.
// Class used to store any kind of animation
type Animation struct{}

// JSObject returns the underlying js.Value.
func (a *Animation) JSObject() js.Value { return a.p }

// Animation returns a Animation JavaScript class.
func (b *Babylon) Animation() *Animation {
	p := b.ctx.Get("Animation")
	return AnimationFromJSObject(p)
}

// AnimationFromJSObject returns a wrapped Animation JavaScript class.
func AnimationFromJSObject(p js.Value) *Animation {
	return &Animation{p: p}
}

// NewAnimation returns a new Animation object.
//
// https://doc.babylonjs.com/api/classes/babylon.animation
func (b *Babylon) NewAnimation(todo parameters) *Animation {
	p := b.ctx.Get("Animation").New(todo)
	return AnimationFromJSObject(p)
}

// TODO: methods
