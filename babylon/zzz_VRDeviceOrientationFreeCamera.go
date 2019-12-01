// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDeviceOrientationFreeCamera represents a babylon.js VRDeviceOrientationFreeCamera.
// Camera used to simulate VR rendering (based on FreeCamera)
//
// See: http://doc.babylonjs.com/babylon101/cameras#vr-device-orientation-cameras
type VRDeviceOrientationFreeCamera struct {
	*DeviceOrientationCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VRDeviceOrientationFreeCamera) JSObject() js.Value { return v.p }

// VRDeviceOrientationFreeCamera returns a VRDeviceOrientationFreeCamera JavaScript class.
func (ba *Babylon) VRDeviceOrientationFreeCamera() *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera")
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// VRDeviceOrientationFreeCameraFromJSObject returns a wrapped VRDeviceOrientationFreeCamera JavaScript class.
func VRDeviceOrientationFreeCameraFromJSObject(p js.Value, ctx js.Value) *VRDeviceOrientationFreeCamera {
	return &VRDeviceOrientationFreeCamera{DeviceOrientationCamera: DeviceOrientationCameraFromJSObject(p, ctx), ctx: ctx}
}

// VRDeviceOrientationFreeCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func VRDeviceOrientationFreeCameraArrayToJSArray(array []*VRDeviceOrientationFreeCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVRDeviceOrientationFreeCameraOpts contains optional parameters for NewVRDeviceOrientationFreeCamera.
type NewVRDeviceOrientationFreeCameraOpts struct {
	CompensateDistortion *bool
	VrCameraMetrics      *VRCameraMetrics
}

// NewVRDeviceOrientationFreeCamera returns a new VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera
func (ba *Babylon) NewVRDeviceOrientationFreeCamera(name string, position *Vector3, scene *Scene, opts *NewVRDeviceOrientationFreeCameraOpts) *VRDeviceOrientationFreeCamera {
	if opts == nil {
		opts = &NewVRDeviceOrientationFreeCameraOpts{}
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

	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(args...)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// VRDeviceOrientationFreeCameraAttachControlOpts contains optional parameters for VRDeviceOrientationFreeCamera.AttachControl.
type VRDeviceOrientationFreeCameraAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#attachcontrol
func (v *VRDeviceOrientationFreeCamera) AttachControl(element js.Value, opts *VRDeviceOrientationFreeCameraAttachControlOpts) {
	if opts == nil {
		opts = &VRDeviceOrientationFreeCameraAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	v.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#detachcontrol
func (v *VRDeviceOrientationFreeCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	v.p.Call("detachControl", args...)
}

// Dispose calls the Dispose method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#dispose
func (v *VRDeviceOrientationFreeCamera) Dispose() {

	v.p.Call("dispose")
}

// VRDeviceOrientationFreeCameraEnableHorizontalDraggingOpts contains optional parameters for VRDeviceOrientationFreeCamera.EnableHorizontalDragging.
type VRDeviceOrientationFreeCameraEnableHorizontalDraggingOpts struct {
	DragFactor *float64
}

// EnableHorizontalDragging calls the EnableHorizontalDragging method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#enablehorizontaldragging
func (v *VRDeviceOrientationFreeCamera) EnableHorizontalDragging(opts *VRDeviceOrientationFreeCameraEnableHorizontalDraggingOpts) {
	if opts == nil {
		opts = &VRDeviceOrientationFreeCameraEnableHorizontalDraggingOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DragFactor == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DragFactor)
	}

	v.p.Call("enableHorizontalDragging", args...)
}

// GetClassName calls the GetClassName method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#getclassname
func (v *VRDeviceOrientationFreeCamera) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// GetFrontPosition calls the GetFrontPosition method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#getfrontposition
func (v *VRDeviceOrientationFreeCamera) GetFrontPosition(distance float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	retVal := v.p.Call("getFrontPosition", args...)
	return Vector3FromJSObject(retVal, v.ctx)
}

// GetTarget calls the GetTarget method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#gettarget
func (v *VRDeviceOrientationFreeCamera) GetTarget() *Vector3 {

	retVal := v.p.Call("getTarget")
	return Vector3FromJSObject(retVal, v.ctx)
}

// VRDeviceOrientationFreeCameraResetToCurrentRotationOpts contains optional parameters for VRDeviceOrientationFreeCamera.ResetToCurrentRotation.
type VRDeviceOrientationFreeCameraResetToCurrentRotationOpts struct {
	Axis *Axis
}

// ResetToCurrentRotation calls the ResetToCurrentRotation method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#resettocurrentrotation
func (v *VRDeviceOrientationFreeCamera) ResetToCurrentRotation(opts *VRDeviceOrientationFreeCameraResetToCurrentRotationOpts) {
	if opts == nil {
		opts = &VRDeviceOrientationFreeCameraResetToCurrentRotationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Axis == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Axis.JSObject())
	}

	v.p.Call("resetToCurrentRotation", args...)
}

// SetTarget calls the SetTarget method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#settarget
func (v *VRDeviceOrientationFreeCamera) SetTarget(target *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	v.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the VRDeviceOrientationFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#storestate
func (v *VRDeviceOrientationFreeCamera) StoreState() *Camera {

	retVal := v.p.Call("storeState")
	return CameraFromJSObject(retVal, v.ctx)
}

/*

// AngularSensibility returns the AngularSensibility property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#angularsensibility
func (v *VRDeviceOrientationFreeCamera) AngularSensibility(angularSensibility float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(angularSensibility)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#angularsensibility
func (v *VRDeviceOrientationFreeCamera) SetAngularSensibility(angularSensibility float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(angularSensibility)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// ApplyGravity returns the ApplyGravity property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#applygravity
func (v *VRDeviceOrientationFreeCamera) ApplyGravity(applyGravity bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(applyGravity)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetApplyGravity sets the ApplyGravity property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#applygravity
func (v *VRDeviceOrientationFreeCamera) SetApplyGravity(applyGravity bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(applyGravity)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#cameradirection
func (v *VRDeviceOrientationFreeCamera) CameraDirection(cameraDirection *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(cameraDirection.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#cameradirection
func (v *VRDeviceOrientationFreeCamera) SetCameraDirection(cameraDirection *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(cameraDirection.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// CameraRotation returns the CameraRotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#camerarotation
func (v *VRDeviceOrientationFreeCamera) CameraRotation(cameraRotation *Vector2) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(cameraRotation.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraRotation sets the CameraRotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#camerarotation
func (v *VRDeviceOrientationFreeCamera) SetCameraRotation(cameraRotation *Vector2) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(cameraRotation.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// CheckCollisions returns the CheckCollisions property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#checkcollisions
func (v *VRDeviceOrientationFreeCamera) CheckCollisions(checkCollisions bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(checkCollisions)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetCheckCollisions sets the CheckCollisions property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#checkcollisions
func (v *VRDeviceOrientationFreeCamera) SetCheckCollisions(checkCollisions bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(checkCollisions)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// CollisionMask returns the CollisionMask property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#collisionmask
func (v *VRDeviceOrientationFreeCamera) CollisionMask(collisionMask float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(collisionMask)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetCollisionMask sets the CollisionMask property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#collisionmask
func (v *VRDeviceOrientationFreeCamera) SetCollisionMask(collisionMask float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(collisionMask)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// DisablePointerInputWhenUsingDeviceOrientation returns the DisablePointerInputWhenUsingDeviceOrientation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#disablepointerinputwhenusingdeviceorientation
func (v *VRDeviceOrientationFreeCamera) DisablePointerInputWhenUsingDeviceOrientation(disablePointerInputWhenUsingDeviceOrientation bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(disablePointerInputWhenUsingDeviceOrientation)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetDisablePointerInputWhenUsingDeviceOrientation sets the DisablePointerInputWhenUsingDeviceOrientation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#disablepointerinputwhenusingdeviceorientation
func (v *VRDeviceOrientationFreeCamera) SetDisablePointerInputWhenUsingDeviceOrientation(disablePointerInputWhenUsingDeviceOrientation bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(disablePointerInputWhenUsingDeviceOrientation)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// Ellipsoid returns the Ellipsoid property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#ellipsoid
func (v *VRDeviceOrientationFreeCamera) Ellipsoid(ellipsoid *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(ellipsoid.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#ellipsoid
func (v *VRDeviceOrientationFreeCamera) SetEllipsoid(ellipsoid *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(ellipsoid.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// EllipsoidOffset returns the EllipsoidOffset property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#ellipsoidoffset
func (v *VRDeviceOrientationFreeCamera) EllipsoidOffset(ellipsoidOffset *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(ellipsoidOffset.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#ellipsoidoffset
func (v *VRDeviceOrientationFreeCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(ellipsoidOffset.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#inputs
func (v *VRDeviceOrientationFreeCamera) Inputs(inputs *FreeCameraInputsManager) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(inputs.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#inputs
func (v *VRDeviceOrientationFreeCamera) SetInputs(inputs *FreeCameraInputsManager) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(inputs.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysdown
func (v *VRDeviceOrientationFreeCamera) KeysDown(keysDown float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysDown)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysdown
func (v *VRDeviceOrientationFreeCamera) SetKeysDown(keysDown float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysDown)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysleft
func (v *VRDeviceOrientationFreeCamera) KeysLeft(keysLeft float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysLeft)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysleft
func (v *VRDeviceOrientationFreeCamera) SetKeysLeft(keysLeft float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysLeft)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysright
func (v *VRDeviceOrientationFreeCamera) KeysRight(keysRight float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysRight)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysright
func (v *VRDeviceOrientationFreeCamera) SetKeysRight(keysRight float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysRight)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysup
func (v *VRDeviceOrientationFreeCamera) KeysUp(keysUp float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysUp)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#keysup
func (v *VRDeviceOrientationFreeCamera) SetKeysUp(keysUp float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(keysUp)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// LockedTarget returns the LockedTarget property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#lockedtarget
func (v *VRDeviceOrientationFreeCamera) LockedTarget(lockedTarget interface{}) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(lockedTarget)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetLockedTarget sets the LockedTarget property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#lockedtarget
func (v *VRDeviceOrientationFreeCamera) SetLockedTarget(lockedTarget interface{}) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(lockedTarget)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// NoRotationConstraint returns the NoRotationConstraint property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#norotationconstraint
func (v *VRDeviceOrientationFreeCamera) NoRotationConstraint(noRotationConstraint bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(noRotationConstraint)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetNoRotationConstraint sets the NoRotationConstraint property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#norotationconstraint
func (v *VRDeviceOrientationFreeCamera) SetNoRotationConstraint(noRotationConstraint bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(noRotationConstraint)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// OnCollide returns the OnCollide property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#oncollide
func (v *VRDeviceOrientationFreeCamera) OnCollide(onCollide func()) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnCollide sets the OnCollide property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#oncollide
func (v *VRDeviceOrientationFreeCamera) SetOnCollide(onCollide func()) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#rotation
func (v *VRDeviceOrientationFreeCamera) Rotation(rotation *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(rotation.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#rotation
func (v *VRDeviceOrientationFreeCamera) SetRotation(rotation *Vector3) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(rotation.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#rotationquaternion
func (v *VRDeviceOrientationFreeCamera) RotationQuaternion(rotationQuaternion *Quaternion) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(rotationQuaternion.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#rotationquaternion
func (v *VRDeviceOrientationFreeCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(rotationQuaternion.JSObject())
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// Speed returns the Speed property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#speed
func (v *VRDeviceOrientationFreeCamera) Speed(speed float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(speed)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetSpeed sets the Speed property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#speed
func (v *VRDeviceOrientationFreeCamera) SetSpeed(speed float64) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(speed)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// UpdateUpVectorFromRotation returns the UpdateUpVectorFromRotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#updateupvectorfromrotation
func (v *VRDeviceOrientationFreeCamera) UpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(updateUpVectorFromRotation)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

// SetUpdateUpVectorFromRotation sets the UpdateUpVectorFromRotation property of class VRDeviceOrientationFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationfreecamera#updateupvectorfromrotation
func (v *VRDeviceOrientationFreeCamera) SetUpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *VRDeviceOrientationFreeCamera {
	p := ba.ctx.Get("VRDeviceOrientationFreeCamera").New(updateUpVectorFromRotation)
	return VRDeviceOrientationFreeCameraFromJSObject(p, ba.ctx)
}

*/
