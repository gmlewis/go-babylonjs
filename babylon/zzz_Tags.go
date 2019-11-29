// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Tags represents a babylon.js Tags.
// Class used to store custom tags
type Tags struct{}

// JSObject returns the underlying js.Value.
func (t *Tags) JSObject() js.Value { return t.p }

// Tags returns a Tags JavaScript class.
func (b *Babylon) Tags() *Tags {
	p := b.ctx.Get("Tags")
	return TagsFromJSObject(p)
}

// TagsFromJSObject returns a wrapped Tags JavaScript class.
func TagsFromJSObject(p js.Value) *Tags {
	return &Tags{p: p}
}

// NewTags returns a new Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags
func (b *Babylon) NewTags(todo parameters) *Tags {
	p := b.ctx.Get("Tags").New(todo)
	return TagsFromJSObject(p)
}

// TODO: methods
