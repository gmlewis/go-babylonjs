// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRControllerTeleportation represents a babylon.js WebXRControllerTeleportation.
// Enables teleportation
type WebXRControllerTeleportation struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRControllerTeleportation) JSObject() js.Value { return w.p }

// WebXRControllerTeleportation returns a WebXRControllerTeleportation JavaScript class.
func (ba *Babylon) WebXRControllerTeleportation() *WebXRControllerTeleportation {
	p := ba.ctx.Get("WebXRControllerTeleportation")
	return WebXRControllerTeleportationFromJSObject(p, ba.ctx)
}

// WebXRControllerTeleportationFromJSObject returns a wrapped WebXRControllerTeleportation JavaScript class.
func WebXRControllerTeleportationFromJSObject(p js.Value, ctx js.Value) *WebXRControllerTeleportation {
	return &WebXRControllerTeleportation{p: p, ctx: ctx}
}

// WebXRControllerTeleportationArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebXRControllerTeleportationArrayToJSArray(array []*WebXRControllerTeleportation) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWebXRControllerTeleportationOpts contains optional parameters for NewWebXRControllerTeleportation.
type NewWebXRControllerTeleportationOpts struct {
	FloorMeshes []js.Value
}

// NewWebXRControllerTeleportation returns a new WebXRControllerTeleportation object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontrollerteleportation#constructor
func (ba *Babylon) NewWebXRControllerTeleportation(input *WebXRInput, opts *NewWebXRControllerTeleportationOpts) *WebXRControllerTeleportation {
	if opts == nil {
		opts = &NewWebXRControllerTeleportationOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, input.JSObject())

	if opts.FloorMeshes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FloorMeshes)
	}

	p := ba.ctx.Get("WebXRControllerTeleportation").New(args...)
	return WebXRControllerTeleportationFromJSObject(p, ba.ctx)
}
