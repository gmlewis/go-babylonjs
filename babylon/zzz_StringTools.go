// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StringTools represents a babylon.js StringTools.
// Helper to manipulate strings
type StringTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StringTools) JSObject() js.Value { return s.p }

// StringTools returns a StringTools JavaScript class.
func (ba *Babylon) StringTools() *StringTools {
	p := ba.ctx.Get("StringTools")
	return StringToolsFromJSObject(p, ba.ctx)
}

// StringToolsFromJSObject returns a wrapped StringTools JavaScript class.
func StringToolsFromJSObject(p js.Value, ctx js.Value) *StringTools {
	return &StringTools{p: p, ctx: ctx}
}

// TODO: methods
