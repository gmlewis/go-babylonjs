// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDeviceOrientationGamepadCamera represents a babylon.js VRDeviceOrientationGamepadCamera.
// Camera used to simulate VR rendering (based on VRDeviceOrientationFreeCamera)

//
// See: http://doc.babylonjs.com/babylon101/cameras#vr-device-orientation-cameras
type VRDeviceOrientationGamepadCamera struct{ *VRDeviceOrientationFreeCamera }

// JSObject returns the underlying js.Value.
func (v *VRDeviceOrientationGamepadCamera) JSObject() js.Value { return v.p }

// VRDeviceOrientationGamepadCamera returns a VRDeviceOrientationGamepadCamera JavaScript class.
func (b *Babylon) VRDeviceOrientationGamepadCamera() *VRDeviceOrientationGamepadCamera {
	p := b.ctx.Get("VRDeviceOrientationGamepadCamera")
	return VRDeviceOrientationGamepadCameraFromJSObject(p)
}

// VRDeviceOrientationGamepadCameraFromJSObject returns a wrapped VRDeviceOrientationGamepadCamera JavaScript class.
func VRDeviceOrientationGamepadCameraFromJSObject(p js.Value) *VRDeviceOrientationGamepadCamera {
	return &VRDeviceOrientationGamepadCamera{VRDeviceOrientationFreeCameraFromJSObject(p)}
}

// NewVRDeviceOrientationGamepadCamera returns a new VRDeviceOrientationGamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationgamepadcamera
func (b *Babylon) NewVRDeviceOrientationGamepadCamera(todo parameters) *VRDeviceOrientationGamepadCamera {
	p := b.ctx.Get("VRDeviceOrientationGamepadCamera").New(todo)
	return VRDeviceOrientationGamepadCameraFromJSObject(p)
}

// TODO: methods
