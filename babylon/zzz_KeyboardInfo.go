// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeyboardInfo represents a babylon.js KeyboardInfo.
// This class is used to store keyboard related info for the onKeyboardObservable event.
type KeyboardInfo struct{}

// JSObject returns the underlying js.Value.
func (k *KeyboardInfo) JSObject() js.Value { return k.p }

// KeyboardInfo returns a KeyboardInfo JavaScript class.
func (b *Babylon) KeyboardInfo() *KeyboardInfo {
	p := b.ctx.Get("KeyboardInfo")
	return KeyboardInfoFromJSObject(p)
}

// KeyboardInfoFromJSObject returns a wrapped KeyboardInfo JavaScript class.
func KeyboardInfoFromJSObject(p js.Value) *KeyboardInfo {
	return &KeyboardInfo{p: p}
}

// NewKeyboardInfo returns a new KeyboardInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo
func (b *Babylon) NewKeyboardInfo(todo parameters) *KeyboardInfo {
	p := b.ctx.Get("KeyboardInfo").New(todo)
	return KeyboardInfoFromJSObject(p)
}

// TODO: methods