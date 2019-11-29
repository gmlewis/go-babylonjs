// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeyboardEventTypes represents a babylon.js KeyboardEventTypes.
// Gather the list of keyboard event types as constants.
type KeyboardEventTypes struct{}

// JSObject returns the underlying js.Value.
func (k *KeyboardEventTypes) JSObject() js.Value { return k.p }

// KeyboardEventTypes returns a KeyboardEventTypes JavaScript class.
func (b *Babylon) KeyboardEventTypes() *KeyboardEventTypes {
	p := b.ctx.Get("KeyboardEventTypes")
	return KeyboardEventTypesFromJSObject(p)
}

// KeyboardEventTypesFromJSObject returns a wrapped KeyboardEventTypes JavaScript class.
func KeyboardEventTypesFromJSObject(p js.Value) *KeyboardEventTypes {
	return &KeyboardEventTypes{p: p}
}

// NewKeyboardEventTypes returns a new KeyboardEventTypes object.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardeventtypes
func (b *Babylon) NewKeyboardEventTypes(todo parameters) *KeyboardEventTypes {
	p := b.ctx.Get("KeyboardEventTypes").New(todo)
	return KeyboardEventTypesFromJSObject(p)
}

// TODO: methods