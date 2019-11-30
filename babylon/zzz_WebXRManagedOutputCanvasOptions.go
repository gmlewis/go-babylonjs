// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRManagedOutputCanvasOptions represents a babylon.js WebXRManagedOutputCanvasOptions.
// COnfiguration object for WebXR output canvas
type WebXRManagedOutputCanvasOptions struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (w *WebXRManagedOutputCanvasOptions) JSObject() js.Value { return w.p }

// WebXRManagedOutputCanvasOptions returns a WebXRManagedOutputCanvasOptions JavaScript class.
func (b *Babylon) WebXRManagedOutputCanvasOptions() *WebXRManagedOutputCanvasOptions {
	p := b.ctx.Get("WebXRManagedOutputCanvasOptions")
	return WebXRManagedOutputCanvasOptionsFromJSObject(p)
}

// WebXRManagedOutputCanvasOptionsFromJSObject returns a wrapped WebXRManagedOutputCanvasOptions JavaScript class.
func WebXRManagedOutputCanvasOptionsFromJSObject(p js.Value) *WebXRManagedOutputCanvasOptions {
	return &WebXRManagedOutputCanvasOptions{p: p}
}

// TODO: methods
