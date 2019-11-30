// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DualShockPad represents a babylon.js DualShockPad.
// Defines a DualShock gamepad
type DualShockPad struct {
	*Gamepad
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DualShockPad) JSObject() js.Value { return d.p }

// DualShockPad returns a DualShockPad JavaScript class.
func (ba *Babylon) DualShockPad() *DualShockPad {
	p := ba.ctx.Get("DualShockPad")
	return DualShockPadFromJSObject(p, ba.ctx)
}

// DualShockPadFromJSObject returns a wrapped DualShockPad JavaScript class.
func DualShockPadFromJSObject(p js.Value, ctx js.Value) *DualShockPad {
	return &DualShockPad{Gamepad: GamepadFromJSObject(p, ctx), ctx: ctx}
}

// NewDualShockPad returns a new DualShockPad object.
//
// https://doc.babylonjs.com/api/classes/babylon.dualshockpad
func (ba *Babylon) NewDualShockPad(id string, index float64, gamepad interface{}) *DualShockPad {
	p := ba.ctx.Get("DualShockPad").New(id, index, gamepad)
	return DualShockPadFromJSObject(p, ba.ctx)
}

// TODO: methods
