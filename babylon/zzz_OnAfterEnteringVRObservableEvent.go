// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OnAfterEnteringVRObservableEvent represents a babylon.js OnAfterEnteringVRObservableEvent.
// Event containing information after VR has been entered
type OnAfterEnteringVRObservableEvent struct{}

// JSObject returns the underlying js.Value.
func (o *OnAfterEnteringVRObservableEvent) JSObject() js.Value { return o.p }

// OnAfterEnteringVRObservableEvent returns a OnAfterEnteringVRObservableEvent JavaScript class.
func (b *Babylon) OnAfterEnteringVRObservableEvent() *OnAfterEnteringVRObservableEvent {
	p := b.ctx.Get("OnAfterEnteringVRObservableEvent")
	return OnAfterEnteringVRObservableEventFromJSObject(p)
}

// OnAfterEnteringVRObservableEventFromJSObject returns a wrapped OnAfterEnteringVRObservableEvent JavaScript class.
func OnAfterEnteringVRObservableEventFromJSObject(p js.Value) *OnAfterEnteringVRObservableEvent {
	return &OnAfterEnteringVRObservableEvent{p: p}
}

// NewOnAfterEnteringVRObservableEvent returns a new OnAfterEnteringVRObservableEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.onafterenteringvrobservableevent
func (b *Babylon) NewOnAfterEnteringVRObservableEvent(todo parameters) *OnAfterEnteringVRObservableEvent {
	p := b.ctx.Get("OnAfterEnteringVRObservableEvent").New(todo)
	return OnAfterEnteringVRObservableEventFromJSObject(p)
}

// TODO: methods
