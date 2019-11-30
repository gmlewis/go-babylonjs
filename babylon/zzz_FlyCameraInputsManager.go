// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FlyCameraInputsManager represents a babylon.js FlyCameraInputsManager.
// Default Inputs manager for the FlyCamera.
// It groups all the default supported inputs for ease of use.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FlyCameraInputsManager struct {
	*CameraInputsManager
	*FlyCamera
}

// JSObject returns the underlying js.Value.
func (f *FlyCameraInputsManager) JSObject() js.Value { return f.p }

// FlyCameraInputsManager returns a FlyCameraInputsManager JavaScript class.
func (b *Babylon) FlyCameraInputsManager() *FlyCameraInputsManager {
	p := b.ctx.Get("FlyCameraInputsManager")
	return FlyCameraInputsManagerFromJSObject(p)
}

// FlyCameraInputsManagerFromJSObject returns a wrapped FlyCameraInputsManager JavaScript class.
func FlyCameraInputsManagerFromJSObject(p js.Value) *FlyCameraInputsManager {
	return &FlyCameraInputsManager{CameraInputsManagerFromJSObject(p)}
}

// NewFlyCameraInputsManager returns a new FlyCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamerainputsmanager
func (b *Babylon) NewFlyCameraInputsManager(camera *FlyCamera) *FlyCameraInputsManager {
	p := b.ctx.Get("FlyCameraInputsManager").New(camera.JSObject())
	return FlyCameraInputsManagerFromJSObject(p)
}

// TODO: methods
