// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArrayTools represents a babylon.js ArrayTools.
// Class containing a set of static utilities functions for arrays.
type ArrayTools struct{}

// JSObject returns the underlying js.Value.
func (a *ArrayTools) JSObject() js.Value { return a.p }

// ArrayTools returns a ArrayTools JavaScript class.
func (b *Babylon) ArrayTools() *ArrayTools {
	p := b.ctx.Get("ArrayTools")
	return ArrayToolsFromJSObject(p)
}

// ArrayToolsFromJSObject returns a wrapped ArrayTools JavaScript class.
func ArrayToolsFromJSObject(p js.Value) *ArrayTools {
	return &ArrayTools{p: p}
}

// NewArrayTools returns a new ArrayTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.arraytools
func (b *Babylon) NewArrayTools(todo parameters) *ArrayTools {
	p := b.ctx.Get("ArrayTools").New(todo)
	return ArrayToolsFromJSObject(p)
}

// TODO: methods