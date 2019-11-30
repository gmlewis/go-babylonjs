// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FreeCameraInputsManager represents a babylon.js FreeCameraInputsManager.
// Default Inputs manager for the FreeCamera.
// It groups all the default supported inputs for ease of use.
// Interface representing a free camera inputs manager
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FreeCameraInputsManager struct {
	*CameraInputsManager
	*FreeCamera
}

// JSObject returns the underlying js.Value.
func (f *FreeCameraInputsManager) JSObject() js.Value { return f.p }

// FreeCameraInputsManager returns a FreeCameraInputsManager JavaScript class.
func (b *Babylon) FreeCameraInputsManager() *FreeCameraInputsManager {
	p := b.ctx.Get("FreeCameraInputsManager")
	return FreeCameraInputsManagerFromJSObject(p)
}

// FreeCameraInputsManagerFromJSObject returns a wrapped FreeCameraInputsManager JavaScript class.
func FreeCameraInputsManagerFromJSObject(p js.Value) *FreeCameraInputsManager {
	return &FreeCameraInputsManager{CameraInputsManagerFromJSObject(p)}
}

// NewFreeCameraInputsManager returns a new FreeCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerainputsmanager
func (b *Babylon) NewFreeCameraInputsManager(camera *FreeCamera) *FreeCameraInputsManager {
	p := b.ctx.Get("FreeCameraInputsManager").New(camera.JSObject())
	return FreeCameraInputsManagerFromJSObject(p)
}

// TODO: methods
