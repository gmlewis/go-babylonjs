// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXREnterExitUIButton represents a babylon.js WebXREnterExitUIButton.
// Button which can be used to enter a different mode of XR
type WebXREnterExitUIButton struct{}

// JSObject returns the underlying js.Value.
func (w *WebXREnterExitUIButton) JSObject() js.Value { return w.p }

// WebXREnterExitUIButton returns a WebXREnterExitUIButton JavaScript class.
func (b *Babylon) WebXREnterExitUIButton() *WebXREnterExitUIButton {
	p := b.ctx.Get("WebXREnterExitUIButton")
	return WebXREnterExitUIButtonFromJSObject(p)
}

// WebXREnterExitUIButtonFromJSObject returns a wrapped WebXREnterExitUIButton JavaScript class.
func WebXREnterExitUIButtonFromJSObject(p js.Value) *WebXREnterExitUIButton {
	return &WebXREnterExitUIButton{p: p}
}

// NewWebXREnterExitUIButton returns a new WebXREnterExitUIButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexituibutton
func (b *Babylon) NewWebXREnterExitUIButton(todo parameters) *WebXREnterExitUIButton {
	p := b.ctx.Get("WebXREnterExitUIButton").New(todo)
	return WebXREnterExitUIButtonFromJSObject(p)
}

// TODO: methods