// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDeviceOrientationArcRotateCamera represents a babylon.js VRDeviceOrientationArcRotateCamera.
// Camera used to simulate VR rendering (based on ArcRotateCamera)

//
// See: http://doc.babylonjs.com/babylon101/cameras#vr-device-orientation-cameras
type VRDeviceOrientationArcRotateCamera struct{ *ArcRotateCamera }

// JSObject returns the underlying js.Value.
func (v *VRDeviceOrientationArcRotateCamera) JSObject() js.Value { return v.p }

// VRDeviceOrientationArcRotateCamera returns a VRDeviceOrientationArcRotateCamera JavaScript class.
func (b *Babylon) VRDeviceOrientationArcRotateCamera() *VRDeviceOrientationArcRotateCamera {
	p := b.ctx.Get("VRDeviceOrientationArcRotateCamera")
	return VRDeviceOrientationArcRotateCameraFromJSObject(p)
}

// VRDeviceOrientationArcRotateCameraFromJSObject returns a wrapped VRDeviceOrientationArcRotateCamera JavaScript class.
func VRDeviceOrientationArcRotateCameraFromJSObject(p js.Value) *VRDeviceOrientationArcRotateCamera {
	return &VRDeviceOrientationArcRotateCamera{ArcRotateCameraFromJSObject(p)}
}

// NewVRDeviceOrientationArcRotateCamera returns a new VRDeviceOrientationArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationarcrotatecamera
func (b *Babylon) NewVRDeviceOrientationArcRotateCamera(todo parameters) *VRDeviceOrientationArcRotateCamera {
	p := b.ctx.Get("VRDeviceOrientationArcRotateCamera").New(todo)
	return VRDeviceOrientationArcRotateCameraFromJSObject(p)
}

// TODO: methods