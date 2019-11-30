// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRController represents a babylon.js WebXRController.
// Represents an XR input
type WebXRController struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRController) JSObject() js.Value { return w.p }

// WebXRController returns a WebXRController JavaScript class.
func (ba *Babylon) WebXRController() *WebXRController {
	p := ba.ctx.Get("WebXRController")
	return WebXRControllerFromJSObject(p, ba.ctx)
}

// WebXRControllerFromJSObject returns a wrapped WebXRController JavaScript class.
func WebXRControllerFromJSObject(p js.Value, ctx js.Value) *WebXRController {
	return &WebXRController{p: p, ctx: ctx}
}

// NewWebXRControllerOpts contains optional parameters for NewWebXRController.
type NewWebXRControllerOpts struct {
	ParentContainer *AbstractMesh
}

// NewWebXRController returns a new WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller
func (ba *Babylon) NewWebXRController(scene *Scene, inputSource js.Value, opts *NewWebXRControllerOpts) *WebXRController {
	if opts == nil {
		opts = &NewWebXRControllerOpts{}
	}

	p := ba.ctx.Get("WebXRController").New(scene.JSObject(), inputSource, opts.ParentContainer.JSObject())
	return WebXRControllerFromJSObject(p, ba.ctx)
}

// TODO: methods
