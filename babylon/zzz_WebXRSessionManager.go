// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRSessionManager represents a babylon.js WebXRSessionManager.
// Manages an XRSession to work with Babylon&amp;#39;s engine
//
// See: https://doc.babylonjs.com/how_to/webxr
type WebXRSessionManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (w *WebXRSessionManager) JSObject() js.Value { return w.p }

// WebXRSessionManager returns a WebXRSessionManager JavaScript class.
func (b *Babylon) WebXRSessionManager() *WebXRSessionManager {
	p := b.ctx.Get("WebXRSessionManager")
	return WebXRSessionManagerFromJSObject(p)
}

// WebXRSessionManagerFromJSObject returns a wrapped WebXRSessionManager JavaScript class.
func WebXRSessionManagerFromJSObject(p js.Value) *WebXRSessionManager {
	return &WebXRSessionManager{p: p}
}

// NewWebXRSessionManager returns a new WebXRSessionManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrsessionmanager
func (b *Babylon) NewWebXRSessionManager(scene *Scene) *WebXRSessionManager {
	p := b.ctx.Get("WebXRSessionManager").New(scene.JSObject())
	return WebXRSessionManagerFromJSObject(p)
}

// TODO: methods
