// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRCamera represents a babylon.js WebXRCamera.
// WebXR Camera which holds the views for the xrSession
//
// See: https://doc.babylonjs.com/how_to/webxr
type WebXRCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRCamera) JSObject() js.Value { return w.p }

// WebXRCamera returns a WebXRCamera JavaScript class.
func (ba *Babylon) WebXRCamera() *WebXRCamera {
	p := ba.ctx.Get("WebXRCamera")
	return WebXRCameraFromJSObject(p, ba.ctx)
}

// WebXRCameraFromJSObject returns a wrapped WebXRCamera JavaScript class.
func WebXRCameraFromJSObject(p js.Value, ctx js.Value) *WebXRCamera {
	return &WebXRCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// NewWebXRCamera returns a new WebXRCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcamera
func (ba *Babylon) NewWebXRCamera(name string, scene *Scene) *WebXRCamera {
	p := ba.ctx.Get("WebXRCamera").New(name, scene.JSObject())
	return WebXRCameraFromJSObject(p, ba.ctx)
}

// TODO: methods
