// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GamepadManager represents a babylon.js GamepadManager.
// Manager for handling gamepads
type GamepadManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GamepadManager) JSObject() js.Value { return g.p }

// GamepadManager returns a GamepadManager JavaScript class.
func (ba *Babylon) GamepadManager() *GamepadManager {
	p := ba.ctx.Get("GamepadManager")
	return GamepadManagerFromJSObject(p, ba.ctx)
}

// GamepadManagerFromJSObject returns a wrapped GamepadManager JavaScript class.
func GamepadManagerFromJSObject(p js.Value, ctx js.Value) *GamepadManager {
	return &GamepadManager{p: p, ctx: ctx}
}

// NewGamepadManagerOpts contains optional parameters for NewGamepadManager.
type NewGamepadManagerOpts struct {
	_scene *Scene
}

// NewGamepadManager returns a new GamepadManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager
func (ba *Babylon) NewGamepadManager(opts *NewGamepadManagerOpts) *GamepadManager {
	if opts == nil {
		opts = &NewGamepadManagerOpts{}
	}

	p := ba.ctx.Get("GamepadManager").New(opts._scene.JSObject())
	return GamepadManagerFromJSObject(p, ba.ctx)
}

// TODO: methods
