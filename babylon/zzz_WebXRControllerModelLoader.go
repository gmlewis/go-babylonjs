// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRControllerModelLoader represents a babylon.js WebXRControllerModelLoader.
// Loads a controller model and adds it as a child of the WebXRControllers grip when the controller is created
type WebXRControllerModelLoader struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (w *WebXRControllerModelLoader) JSObject() js.Value { return w.p }

// WebXRControllerModelLoader returns a WebXRControllerModelLoader JavaScript class.
func (b *Babylon) WebXRControllerModelLoader() *WebXRControllerModelLoader {
	p := b.ctx.Get("WebXRControllerModelLoader")
	return WebXRControllerModelLoaderFromJSObject(p)
}

// WebXRControllerModelLoaderFromJSObject returns a wrapped WebXRControllerModelLoader JavaScript class.
func WebXRControllerModelLoaderFromJSObject(p js.Value) *WebXRControllerModelLoader {
	return &WebXRControllerModelLoader{p: p}
}

// NewWebXRControllerModelLoader returns a new WebXRControllerModelLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontrollermodelloader
func (b *Babylon) NewWebXRControllerModelLoader(input *WebXRInput) *WebXRControllerModelLoader {
	p := b.ctx.Get("WebXRControllerModelLoader").New(input.JSObject())
	return WebXRControllerModelLoaderFromJSObject(p)
}

// TODO: methods
