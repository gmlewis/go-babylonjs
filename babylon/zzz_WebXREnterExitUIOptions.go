// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXREnterExitUIOptions represents a babylon.js WebXREnterExitUIOptions.
// Options to create the webXR UI
type WebXREnterExitUIOptions struct{}

// JSObject returns the underlying js.Value.
func (w *WebXREnterExitUIOptions) JSObject() js.Value { return w.p }

// WebXREnterExitUIOptions returns a WebXREnterExitUIOptions JavaScript class.
func (b *Babylon) WebXREnterExitUIOptions() *WebXREnterExitUIOptions {
	p := b.ctx.Get("WebXREnterExitUIOptions")
	return WebXREnterExitUIOptionsFromJSObject(p)
}

// WebXREnterExitUIOptionsFromJSObject returns a wrapped WebXREnterExitUIOptions JavaScript class.
func WebXREnterExitUIOptionsFromJSObject(p js.Value) *WebXREnterExitUIOptions {
	return &WebXREnterExitUIOptions{p: p}
}

// NewWebXREnterExitUIOptions returns a new WebXREnterExitUIOptions object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexituioptions
func (b *Babylon) NewWebXREnterExitUIOptions(todo parameters) *WebXREnterExitUIOptions {
	p := b.ctx.Get("WebXREnterExitUIOptions").New(todo)
	return WebXREnterExitUIOptionsFromJSObject(p)
}

// TODO: methods