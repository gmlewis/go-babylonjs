// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRDefaultExperience represents a babylon.js WebXRDefaultExperience.
// Default experience which provides a similar setup to the previous webVRExperience
type WebXRDefaultExperience struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRDefaultExperience) JSObject() js.Value { return w.p }

// WebXRDefaultExperience returns a WebXRDefaultExperience JavaScript class.
func (ba *Babylon) WebXRDefaultExperience() *WebXRDefaultExperience {
	p := ba.ctx.Get("WebXRDefaultExperience")
	return WebXRDefaultExperienceFromJSObject(p, ba.ctx)
}

// WebXRDefaultExperienceFromJSObject returns a wrapped WebXRDefaultExperience JavaScript class.
func WebXRDefaultExperienceFromJSObject(p js.Value, ctx js.Value) *WebXRDefaultExperience {
	return &WebXRDefaultExperience{p: p, ctx: ctx}
}

// TODO: methods
