// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Constants represents a babylon.js Constants.
// Defines the cross module used constants to avoid circular dependncies
type Constants struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (c *Constants) JSObject() js.Value { return c.p }

// Constants returns a Constants JavaScript class.
func (b *Babylon) Constants() *Constants {
	p := b.ctx.Get("Constants")
	return ConstantsFromJSObject(p)
}

// ConstantsFromJSObject returns a wrapped Constants JavaScript class.
func ConstantsFromJSObject(p js.Value) *Constants {
	return &Constants{p: p}
}

// TODO: methods
