// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SelectorGroup represents a babylon.js SelectorGroup.
// Class used to create a RadioGroup
// which contains groups of radio buttons
type SelectorGroup struct{}

// JSObject returns the underlying js.Value.
func (s *SelectorGroup) JSObject() js.Value { return s.p }

// SelectorGroup returns a SelectorGroup JavaScript class.
func (b *Babylon) SelectorGroup() *SelectorGroup {
	p := b.ctx.Get("SelectorGroup")
	return SelectorGroupFromJSObject(p)
}

// SelectorGroupFromJSObject returns a wrapped SelectorGroup JavaScript class.
func SelectorGroupFromJSObject(p js.Value) *SelectorGroup {
	return &SelectorGroup{p: p}
}

// NewSelectorGroup returns a new SelectorGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup
func (b *Babylon) NewSelectorGroup(todo parameters) *SelectorGroup {
	p := b.ctx.Get("SelectorGroup").New(todo)
	return SelectorGroupFromJSObject(p)
}

// TODO: methods
