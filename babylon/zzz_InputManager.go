// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputManager represents a babylon.js InputManager.
// Class used to manage all inputs for the scene.
type InputManager struct{}

// JSObject returns the underlying js.Value.
func (i *InputManager) JSObject() js.Value { return i.p }

// InputManager returns a InputManager JavaScript class.
func (b *Babylon) InputManager() *InputManager {
	p := b.ctx.Get("InputManager")
	return InputManagerFromJSObject(p)
}

// InputManagerFromJSObject returns a wrapped InputManager JavaScript class.
func InputManagerFromJSObject(p js.Value) *InputManager {
	return &InputManager{p: p}
}

// NewInputManager returns a new InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager
func (b *Babylon) NewInputManager(todo parameters) *InputManager {
	p := b.ctx.Get("InputManager").New(todo)
	return InputManagerFromJSObject(p)
}

// TODO: methods