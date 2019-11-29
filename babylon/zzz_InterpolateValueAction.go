// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InterpolateValueAction represents a babylon.js InterpolateValueAction.
// This defines an action responsible to change the value of a property
// by interpolating between its current value and the newly set one once triggered.

//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type InterpolateValueAction struct{ *Action }

// JSObject returns the underlying js.Value.
func (i *InterpolateValueAction) JSObject() js.Value { return i.p }

// InterpolateValueAction returns a InterpolateValueAction JavaScript class.
func (b *Babylon) InterpolateValueAction() *InterpolateValueAction {
	p := b.ctx.Get("InterpolateValueAction")
	return InterpolateValueActionFromJSObject(p)
}

// InterpolateValueActionFromJSObject returns a wrapped InterpolateValueAction JavaScript class.
func InterpolateValueActionFromJSObject(p js.Value) *InterpolateValueAction {
	return &InterpolateValueAction{ActionFromJSObject(p)}
}

// NewInterpolateValueAction returns a new InterpolateValueAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.interpolatevalueaction
func (b *Babylon) NewInterpolateValueAction(todo parameters) *InterpolateValueAction {
	p := b.ctx.Get("InterpolateValueAction").New(todo)
	return InterpolateValueActionFromJSObject(p)
}

// TODO: methods
