// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRController represents a babylon.js WebXRController.
// Represents an XR input
type WebXRController struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (w *WebXRController) JSObject() js.Value { return w.p }

// WebXRController returns a WebXRController JavaScript class.
func (b *Babylon) WebXRController() *WebXRController {
	p := b.ctx.Get("WebXRController")
	return WebXRControllerFromJSObject(p)
}

// WebXRControllerFromJSObject returns a wrapped WebXRController JavaScript class.
func WebXRControllerFromJSObject(p js.Value) *WebXRController {
	return &WebXRController{p: p}
}

// NewWebXRControllerOpts contains optional parameters for NewWebXRController.
type NewWebXRControllerOpts struct {
	ParentContainer *AbstractMesh
}

// NewWebXRController returns a new WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller
func (b *Babylon) NewWebXRController(scene *Scene, inputSource *XRInputSource, opts *NewWebXRControllerOpts) *WebXRController {
	if opts == nil {
		opts = &NewWebXRControllerOpts{}
	}

	p := b.ctx.Get("WebXRController").New(scene.JSObject(), inputSource.JSObject(), opts.ParentContainer.JSObject())
	return WebXRControllerFromJSObject(p)
}

// TODO: methods
