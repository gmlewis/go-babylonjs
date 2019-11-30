// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WindowsMotionController represents a babylon.js WindowsMotionController.
// Defines the WindowsMotionController object that the state of the windows motion controller
type WindowsMotionController struct{ *WebVRController }

// JSObject returns the underlying js.Value.
func (w *WindowsMotionController) JSObject() js.Value { return w.p }

// WindowsMotionController returns a WindowsMotionController JavaScript class.
func (b *Babylon) WindowsMotionController() *WindowsMotionController {
	p := b.ctx.Get("WindowsMotionController")
	return WindowsMotionControllerFromJSObject(p)
}

// WindowsMotionControllerFromJSObject returns a wrapped WindowsMotionController JavaScript class.
func WindowsMotionControllerFromJSObject(p js.Value) *WindowsMotionController {
	return &WindowsMotionController{WebVRControllerFromJSObject(p)}
}

// NewWindowsMotionController returns a new WindowsMotionController object.
//
// https://doc.babylonjs.com/api/classes/babylon.windowsmotioncontroller
func (b *Babylon) NewWindowsMotionController(vrGamepad interface{}) *WindowsMotionController {
	p := b.ctx.Get("WindowsMotionController").New(vrGamepad)
	return WindowsMotionControllerFromJSObject(p)
}

// TODO: methods
