// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MaterialFlags represents a babylon.js MaterialFlags.
// This groups all the flags used to control the materials channel.
type MaterialFlags struct{}

// JSObject returns the underlying js.Value.
func (m *MaterialFlags) JSObject() js.Value { return m.p }

// MaterialFlags returns a MaterialFlags JavaScript class.
func (b *Babylon) MaterialFlags() *MaterialFlags {
	p := b.ctx.Get("MaterialFlags")
	return MaterialFlagsFromJSObject(p)
}

// MaterialFlagsFromJSObject returns a wrapped MaterialFlags JavaScript class.
func MaterialFlagsFromJSObject(p js.Value) *MaterialFlags {
	return &MaterialFlags{p: p}
}

// NewMaterialFlags returns a new MaterialFlags object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags
func (b *Babylon) NewMaterialFlags(todo parameters) *MaterialFlags {
	p := b.ctx.Get("MaterialFlags").New(todo)
	return MaterialFlagsFromJSObject(p)
}

// TODO: methods
