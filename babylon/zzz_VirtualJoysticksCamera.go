// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VirtualJoysticksCamera represents a babylon.js VirtualJoysticksCamera.
// This represents a free type of camera. It can be useful in First Person Shooter game for instance.
// It is identical to the Free Camera and simply adds by default a virtual joystick.
// Virtual Joysticks are on-screen 2D graphics that are used to control the camera or other scene items.
//
// See: http://doc.babylonjs.com/features/cameras#virtual-joysticks-camera
type VirtualJoysticksCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VirtualJoysticksCamera) JSObject() js.Value { return v.p }

// VirtualJoysticksCamera returns a VirtualJoysticksCamera JavaScript class.
func (ba *Babylon) VirtualJoysticksCamera() *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera")
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// VirtualJoysticksCameraFromJSObject returns a wrapped VirtualJoysticksCamera JavaScript class.
func VirtualJoysticksCameraFromJSObject(p js.Value, ctx js.Value) *VirtualJoysticksCamera {
	return &VirtualJoysticksCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// NewVirtualJoysticksCamera returns a new VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera
func (ba *Babylon) NewVirtualJoysticksCamera(name string, position *Vector3, scene *Scene) *VirtualJoysticksCamera {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("VirtualJoysticksCamera").New(args...)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// VirtualJoysticksCameraAttachControlOpts contains optional parameters for VirtualJoysticksCamera.AttachControl.
type VirtualJoysticksCameraAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#attachcontrol
func (v *VirtualJoysticksCamera) AttachControl(element js.Value, opts *VirtualJoysticksCameraAttachControlOpts) {
	if opts == nil {
		opts = &VirtualJoysticksCameraAttachControlOpts{}
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

// DetachControl calls the DetachControl method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#detachcontrol
func (v *VirtualJoysticksCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	v.p.Call("detachControl", args...)
}

// Dispose calls the Dispose method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#dispose
func (v *VirtualJoysticksCamera) Dispose() {

	args := make([]interface{}, 0, 0+0)

	v.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#getclassname
func (v *VirtualJoysticksCamera) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := v.p.Call("getClassName", args...)
	return retVal.String()
}

// GetFrontPosition calls the GetFrontPosition method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#getfrontposition
func (v *VirtualJoysticksCamera) GetFrontPosition(distance float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	retVal := v.p.Call("getFrontPosition", args...)
	return Vector3FromJSObject(retVal, v.ctx)
}

// GetTarget calls the GetTarget method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#gettarget
func (v *VirtualJoysticksCamera) GetTarget() *Vector3 {

	args := make([]interface{}, 0, 0+0)

	retVal := v.p.Call("getTarget", args...)
	return Vector3FromJSObject(retVal, v.ctx)
}

// SetTarget calls the SetTarget method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#settarget
func (v *VirtualJoysticksCamera) SetTarget(target *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	v.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#storestate
func (v *VirtualJoysticksCamera) StoreState() *Camera {

	args := make([]interface{}, 0, 0+0)

	retVal := v.p.Call("storeState", args...)
	return CameraFromJSObject(retVal, v.ctx)
}

/*

// AngularSensibility returns the AngularSensibility property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#angularsensibility
func (v *VirtualJoysticksCamera) AngularSensibility(angularSensibility float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(angularSensibility)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#angularsensibility
func (v *VirtualJoysticksCamera) SetAngularSensibility(angularSensibility float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(angularSensibility)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// ApplyGravity returns the ApplyGravity property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#applygravity
func (v *VirtualJoysticksCamera) ApplyGravity(applyGravity bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(applyGravity)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetApplyGravity sets the ApplyGravity property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#applygravity
func (v *VirtualJoysticksCamera) SetApplyGravity(applyGravity bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(applyGravity)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#cameradirection
func (v *VirtualJoysticksCamera) CameraDirection(cameraDirection *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(cameraDirection.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#cameradirection
func (v *VirtualJoysticksCamera) SetCameraDirection(cameraDirection *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(cameraDirection.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// CameraRotation returns the CameraRotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#camerarotation
func (v *VirtualJoysticksCamera) CameraRotation(cameraRotation *Vector2) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(cameraRotation.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetCameraRotation sets the CameraRotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#camerarotation
func (v *VirtualJoysticksCamera) SetCameraRotation(cameraRotation *Vector2) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(cameraRotation.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// CheckCollisions returns the CheckCollisions property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#checkcollisions
func (v *VirtualJoysticksCamera) CheckCollisions(checkCollisions bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(checkCollisions)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetCheckCollisions sets the CheckCollisions property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#checkcollisions
func (v *VirtualJoysticksCamera) SetCheckCollisions(checkCollisions bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(checkCollisions)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// CollisionMask returns the CollisionMask property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#collisionmask
func (v *VirtualJoysticksCamera) CollisionMask(collisionMask float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(collisionMask)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetCollisionMask sets the CollisionMask property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#collisionmask
func (v *VirtualJoysticksCamera) SetCollisionMask(collisionMask float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(collisionMask)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// Ellipsoid returns the Ellipsoid property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#ellipsoid
func (v *VirtualJoysticksCamera) Ellipsoid(ellipsoid *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(ellipsoid.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#ellipsoid
func (v *VirtualJoysticksCamera) SetEllipsoid(ellipsoid *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(ellipsoid.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// EllipsoidOffset returns the EllipsoidOffset property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#ellipsoidoffset
func (v *VirtualJoysticksCamera) EllipsoidOffset(ellipsoidOffset *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(ellipsoidOffset.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#ellipsoidoffset
func (v *VirtualJoysticksCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(ellipsoidOffset.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#inputs
func (v *VirtualJoysticksCamera) Inputs(inputs *FreeCameraInputsManager) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(inputs.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#inputs
func (v *VirtualJoysticksCamera) SetInputs(inputs *FreeCameraInputsManager) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(inputs.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysdown
func (v *VirtualJoysticksCamera) KeysDown(keysDown float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysDown)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysdown
func (v *VirtualJoysticksCamera) SetKeysDown(keysDown float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysDown)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysleft
func (v *VirtualJoysticksCamera) KeysLeft(keysLeft float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysLeft)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysleft
func (v *VirtualJoysticksCamera) SetKeysLeft(keysLeft float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysLeft)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysright
func (v *VirtualJoysticksCamera) KeysRight(keysRight float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysRight)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysright
func (v *VirtualJoysticksCamera) SetKeysRight(keysRight float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysRight)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysup
func (v *VirtualJoysticksCamera) KeysUp(keysUp float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysUp)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#keysup
func (v *VirtualJoysticksCamera) SetKeysUp(keysUp float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(keysUp)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// LockedTarget returns the LockedTarget property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#lockedtarget
func (v *VirtualJoysticksCamera) LockedTarget(lockedTarget interface{}) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(lockedTarget)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetLockedTarget sets the LockedTarget property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#lockedtarget
func (v *VirtualJoysticksCamera) SetLockedTarget(lockedTarget interface{}) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(lockedTarget)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// NoRotationConstraint returns the NoRotationConstraint property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#norotationconstraint
func (v *VirtualJoysticksCamera) NoRotationConstraint(noRotationConstraint bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(noRotationConstraint)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetNoRotationConstraint sets the NoRotationConstraint property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#norotationconstraint
func (v *VirtualJoysticksCamera) SetNoRotationConstraint(noRotationConstraint bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(noRotationConstraint)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// OnCollide returns the OnCollide property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#oncollide
func (v *VirtualJoysticksCamera) OnCollide(onCollide func()) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(onCollide)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetOnCollide sets the OnCollide property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#oncollide
func (v *VirtualJoysticksCamera) SetOnCollide(onCollide func()) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(onCollide)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#rotation
func (v *VirtualJoysticksCamera) Rotation(rotation *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(rotation.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#rotation
func (v *VirtualJoysticksCamera) SetRotation(rotation *Vector3) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(rotation.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#rotationquaternion
func (v *VirtualJoysticksCamera) RotationQuaternion(rotationQuaternion *Quaternion) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(rotationQuaternion.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#rotationquaternion
func (v *VirtualJoysticksCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(rotationQuaternion.JSObject())
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// Speed returns the Speed property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#speed
func (v *VirtualJoysticksCamera) Speed(speed float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(speed)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetSpeed sets the Speed property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#speed
func (v *VirtualJoysticksCamera) SetSpeed(speed float64) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(speed)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// UpdateUpVectorFromRotation returns the UpdateUpVectorFromRotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#updateupvectorfromrotation
func (v *VirtualJoysticksCamera) UpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(updateUpVectorFromRotation)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

// SetUpdateUpVectorFromRotation sets the UpdateUpVectorFromRotation property of class VirtualJoysticksCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera#updateupvectorfromrotation
func (v *VirtualJoysticksCamera) SetUpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(updateUpVectorFromRotation)
	return VirtualJoysticksCameraFromJSObject(p, ba.ctx)
}

*/
