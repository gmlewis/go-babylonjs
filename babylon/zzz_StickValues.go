// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StickValues represents a babylon.js StickValues.
// Represents a gamepad control stick position
type StickValues struct{}

// JSObject returns the underlying js.Value.
func (s *StickValues) JSObject() js.Value { return s.p }

// StickValues returns a StickValues JavaScript class.
func (b *Babylon) StickValues() *StickValues {
	p := b.ctx.Get("StickValues")
	return StickValuesFromJSObject(p)
}

// StickValuesFromJSObject returns a wrapped StickValues JavaScript class.
func StickValuesFromJSObject(p js.Value) *StickValues {
	return &StickValues{p: p}
}

// NewStickValues returns a new StickValues object.
//
// https://doc.babylonjs.com/api/classes/babylon.stickvalues
func (b *Babylon) NewStickValues(todo parameters) *StickValues {
	p := b.ctx.Get("StickValues").New(todo)
	return StickValuesFromJSObject(p)
}

// TODO: methods
