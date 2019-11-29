// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ClipboardEventTypes represents a babylon.js ClipboardEventTypes.
// Gather the list of clipboard event types as constants.
type ClipboardEventTypes struct{}

// JSObject returns the underlying js.Value.
func (c *ClipboardEventTypes) JSObject() js.Value { return c.p }

// ClipboardEventTypes returns a ClipboardEventTypes JavaScript class.
func (b *Babylon) ClipboardEventTypes() *ClipboardEventTypes {
	p := b.ctx.Get("ClipboardEventTypes")
	return ClipboardEventTypesFromJSObject(p)
}

// ClipboardEventTypesFromJSObject returns a wrapped ClipboardEventTypes JavaScript class.
func ClipboardEventTypesFromJSObject(p js.Value) *ClipboardEventTypes {
	return &ClipboardEventTypes{p: p}
}

// NewClipboardEventTypes returns a new ClipboardEventTypes object.
//
// https://doc.babylonjs.com/api/classes/babylon.clipboardeventtypes
func (b *Babylon) NewClipboardEventTypes(todo parameters) *ClipboardEventTypes {
	p := b.ctx.Get("ClipboardEventTypes").New(todo)
	return ClipboardEventTypesFromJSObject(p)
}

// TODO: methods
