// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDeviceOrientationGamepadCamera represents a babylon.js VRDeviceOrientationGamepadCamera.
// Camera used to simulate VR rendering (based on VRDeviceOrientationFreeCamera)
//
// See: http://doc.babylonjs.com/babylon101/cameras#vr-device-orientation-cameras
type VRDeviceOrientationGamepadCamera struct {
	*VRDeviceOrientationFreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VRDeviceOrientationGamepadCamera) JSObject() js.Value { return v.p }

// VRDeviceOrientationGamepadCamera returns a VRDeviceOrientationGamepadCamera JavaScript class.
func (ba *Babylon) VRDeviceOrientationGamepadCamera() *VRDeviceOrientationGamepadCamera {
	p := ba.ctx.Get("VRDeviceOrientationGamepadCamera")
	return VRDeviceOrientationGamepadCameraFromJSObject(p, ba.ctx)
}

// VRDeviceOrientationGamepadCameraFromJSObject returns a wrapped VRDeviceOrientationGamepadCamera JavaScript class.
func VRDeviceOrientationGamepadCameraFromJSObject(p js.Value, ctx js.Value) *VRDeviceOrientationGamepadCamera {
	return &VRDeviceOrientationGamepadCamera{VRDeviceOrientationFreeCamera: VRDeviceOrientationFreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// VRDeviceOrientationGamepadCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func VRDeviceOrientationGamepadCameraArrayToJSArray(array []*VRDeviceOrientationGamepadCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVRDeviceOrientationGamepadCameraOpts contains optional parameters for NewVRDeviceOrientationGamepadCamera.
type NewVRDeviceOrientationGamepadCameraOpts struct {
	CompensateDistortion *bool
	VrCameraMetrics      *VRCameraMetrics
}

// NewVRDeviceOrientationGamepadCamera returns a new VRDeviceOrientationGamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationgamepadcamera
func (ba *Babylon) NewVRDeviceOrientationGamepadCamera(name string, position *Vector3, scene *Scene, opts *NewVRDeviceOrientationGamepadCameraOpts) *VRDeviceOrientationGamepadCamera {
	if opts == nil {
		opts = &NewVRDeviceOrientationGamepadCameraOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	if opts.CompensateDistortion == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CompensateDistortion)
	}
	if opts.VrCameraMetrics == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.VrCameraMetrics.JSObject())
	}

	p := ba.ctx.Get("VRDeviceOrientationGamepadCamera").New(args...)
	return VRDeviceOrientationGamepadCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the VRDeviceOrientationGamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationgamepadcamera#getclassname
func (v *VRDeviceOrientationGamepadCamera) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

/*

 */
