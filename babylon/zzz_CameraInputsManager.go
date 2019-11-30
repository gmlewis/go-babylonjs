// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CameraInputsManager represents a babylon.js CameraInputsManager.
// This represents the input manager used within a camera.
// It helps dealing with all the different kind of input attached to a camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type CameraInputsManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CameraInputsManager) JSObject() js.Value { return c.p }

// CameraInputsManager returns a CameraInputsManager JavaScript class.
func (ba *Babylon) CameraInputsManager() *CameraInputsManager {
	p := ba.ctx.Get("CameraInputsManager")
	return CameraInputsManagerFromJSObject(p, ba.ctx)
}

// CameraInputsManagerFromJSObject returns a wrapped CameraInputsManager JavaScript class.
func CameraInputsManagerFromJSObject(p js.Value, ctx js.Value) *CameraInputsManager {
	return &CameraInputsManager{p: p, ctx: ctx}
}

// NewCameraInputsManager returns a new CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager
func (ba *Babylon) NewCameraInputsManager(camera *Camera) *CameraInputsManager {
	p := ba.ctx.Get("CameraInputsManager").New(camera.JSObject())
	return CameraInputsManagerFromJSObject(p, ba.ctx)
}

// TODO: methods
