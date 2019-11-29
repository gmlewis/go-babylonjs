// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRDefaultExperienceOptions represents a babylon.js WebXRDefaultExperienceOptions.
// Options for the default xr helper
type WebXRDefaultExperienceOptions struct{}

// JSObject returns the underlying js.Value.
func (w *WebXRDefaultExperienceOptions) JSObject() js.Value { return w.p }

// WebXRDefaultExperienceOptions returns a WebXRDefaultExperienceOptions JavaScript class.
func (b *Babylon) WebXRDefaultExperienceOptions() *WebXRDefaultExperienceOptions {
	p := b.ctx.Get("WebXRDefaultExperienceOptions")
	return WebXRDefaultExperienceOptionsFromJSObject(p)
}

// WebXRDefaultExperienceOptionsFromJSObject returns a wrapped WebXRDefaultExperienceOptions JavaScript class.
func WebXRDefaultExperienceOptionsFromJSObject(p js.Value) *WebXRDefaultExperienceOptions {
	return &WebXRDefaultExperienceOptions{p: p}
}

// NewWebXRDefaultExperienceOptions returns a new WebXRDefaultExperienceOptions object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrdefaultexperienceoptions
func (b *Babylon) NewWebXRDefaultExperienceOptions(todo parameters) *WebXRDefaultExperienceOptions {
	p := b.ctx.Get("WebXRDefaultExperienceOptions").New(todo)
	return WebXRDefaultExperienceOptionsFromJSObject(p)
}

// TODO: methods
