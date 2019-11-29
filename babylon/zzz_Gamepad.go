// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Gamepad represents a babylon.js Gamepad.
// Represents a gamepad
type Gamepad struct{}

// JSObject returns the underlying js.Value.
func (g *Gamepad) JSObject() js.Value { return g.p }

// Gamepad returns a Gamepad JavaScript class.
func (b *Babylon) Gamepad() *Gamepad {
	p := b.ctx.Get("Gamepad")
	return GamepadFromJSObject(p)
}

// GamepadFromJSObject returns a wrapped Gamepad JavaScript class.
func GamepadFromJSObject(p js.Value) *Gamepad {
	return &Gamepad{p: p}
}

// NewGamepad returns a new Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad
func (b *Babylon) NewGamepad(todo parameters) *Gamepad {
	p := b.ctx.Get("Gamepad").New(todo)
	return GamepadFromJSObject(p)
}

// TODO: methods