// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ActionEvent represents a babylon.js ActionEvent.
// ActionEvent is the event being sent when an action is triggered.
type ActionEvent struct{}

// JSObject returns the underlying js.Value.
func (a *ActionEvent) JSObject() js.Value { return a.p }

// ActionEvent returns a ActionEvent JavaScript class.
func (b *Babylon) ActionEvent() *ActionEvent {
	p := b.ctx.Get("ActionEvent")
	return ActionEventFromJSObject(p)
}

// ActionEventFromJSObject returns a wrapped ActionEvent JavaScript class.
func ActionEventFromJSObject(p js.Value) *ActionEvent {
	return &ActionEvent{p: p}
}

// NewActionEvent returns a new ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent
func (b *Babylon) NewActionEvent(todo parameters) *ActionEvent {
	p := b.ctx.Get("ActionEvent").New(todo)
	return ActionEventFromJSObject(p)
}

// TODO: methods
