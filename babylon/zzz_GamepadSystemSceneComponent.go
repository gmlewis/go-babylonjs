// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GamepadSystemSceneComponent represents a babylon.js GamepadSystemSceneComponent.
// Defines the gamepad scene component responsible to manage gamepads in a given scene
type GamepadSystemSceneComponent struct{}

// JSObject returns the underlying js.Value.
func (g *GamepadSystemSceneComponent) JSObject() js.Value { return g.p }

// GamepadSystemSceneComponent returns a GamepadSystemSceneComponent JavaScript class.
func (b *Babylon) GamepadSystemSceneComponent() *GamepadSystemSceneComponent {
	p := b.ctx.Get("GamepadSystemSceneComponent")
	return GamepadSystemSceneComponentFromJSObject(p)
}

// GamepadSystemSceneComponentFromJSObject returns a wrapped GamepadSystemSceneComponent JavaScript class.
func GamepadSystemSceneComponentFromJSObject(p js.Value) *GamepadSystemSceneComponent {
	return &GamepadSystemSceneComponent{p: p}
}

// NewGamepadSystemSceneComponent returns a new GamepadSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadsystemscenecomponent
func (b *Babylon) NewGamepadSystemSceneComponent(todo parameters) *GamepadSystemSceneComponent {
	p := b.ctx.Get("GamepadSystemSceneComponent").New(todo)
	return GamepadSystemSceneComponentFromJSObject(p)
}

// TODO: methods
