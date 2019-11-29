// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnimationEvent represents a babylon.js AnimationEvent.
// Composed of a frame, and an action function
type AnimationEvent struct{}

// JSObject returns the underlying js.Value.
func (a *AnimationEvent) JSObject() js.Value { return a.p }

// AnimationEvent returns a AnimationEvent JavaScript class.
func (b *Babylon) AnimationEvent() *AnimationEvent {
	p := b.ctx.Get("AnimationEvent")
	return AnimationEventFromJSObject(p)
}

// AnimationEventFromJSObject returns a wrapped AnimationEvent JavaScript class.
func AnimationEventFromJSObject(p js.Value) *AnimationEvent {
	return &AnimationEvent{p: p}
}

// NewAnimationEvent returns a new AnimationEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent
func (b *Babylon) NewAnimationEvent(todo parameters) *AnimationEvent {
	p := b.ctx.Get("AnimationEvent").New(todo)
	return AnimationEventFromJSObject(p)
}

// TODO: methods
