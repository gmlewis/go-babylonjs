// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXREnterExitUIButton represents a babylon.js WebXREnterExitUIButton.
// Button which can be used to enter a different mode of XR
type WebXREnterExitUIButton struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXREnterExitUIButton) JSObject() js.Value { return w.p }

// WebXREnterExitUIButton returns a WebXREnterExitUIButton JavaScript class.
func (ba *Babylon) WebXREnterExitUIButton() *WebXREnterExitUIButton {
	p := ba.ctx.Get("WebXREnterExitUIButton")
	return WebXREnterExitUIButtonFromJSObject(p, ba.ctx)
}

// WebXREnterExitUIButtonFromJSObject returns a wrapped WebXREnterExitUIButton JavaScript class.
func WebXREnterExitUIButtonFromJSObject(p js.Value, ctx js.Value) *WebXREnterExitUIButton {
	return &WebXREnterExitUIButton{p: p, ctx: ctx}
}

// NewWebXREnterExitUIButton returns a new WebXREnterExitUIButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexituibutton
func (ba *Babylon) NewWebXREnterExitUIButton(element js.Value, sessionMode js.Value, referenceSpaceType js.Value) *WebXREnterExitUIButton {
	p := ba.ctx.Get("WebXREnterExitUIButton").New(element, sessionMode, referenceSpaceType)
	return WebXREnterExitUIButtonFromJSObject(p, ba.ctx)
}

// TODO: methods
