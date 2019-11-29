// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRDefaultExperience represents a babylon.js WebXRDefaultExperience.
// Default experience which provides a similar setup to the previous webVRExperience
type WebXRDefaultExperience struct{}

// JSObject returns the underlying js.Value.
func (w *WebXRDefaultExperience) JSObject() js.Value { return w.p }

// WebXRDefaultExperience returns a WebXRDefaultExperience JavaScript class.
func (b *Babylon) WebXRDefaultExperience() *WebXRDefaultExperience {
	p := b.ctx.Get("WebXRDefaultExperience")
	return WebXRDefaultExperienceFromJSObject(p)
}

// WebXRDefaultExperienceFromJSObject returns a wrapped WebXRDefaultExperience JavaScript class.
func WebXRDefaultExperienceFromJSObject(p js.Value) *WebXRDefaultExperience {
	return &WebXRDefaultExperience{p: p}
}

// NewWebXRDefaultExperience returns a new WebXRDefaultExperience object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrdefaultexperience
func (b *Babylon) NewWebXRDefaultExperience(todo parameters) *WebXRDefaultExperience {
	p := b.ctx.Get("WebXRDefaultExperience").New(todo)
	return WebXRDefaultExperienceFromJSObject(p)
}

// TODO: methods