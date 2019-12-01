// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphFreeCamera represents a babylon.js AnaglyphFreeCamera.
// Camera used to simulate anaglyphic rendering (based on FreeCamera)
//
// See: http://doc.babylonjs.com/features/cameras#anaglyph-cameras
type AnaglyphFreeCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnaglyphFreeCamera) JSObject() js.Value { return a.p }

// AnaglyphFreeCamera returns a AnaglyphFreeCamera JavaScript class.
func (ba *Babylon) AnaglyphFreeCamera() *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera")
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// AnaglyphFreeCameraFromJSObject returns a wrapped AnaglyphFreeCamera JavaScript class.
func AnaglyphFreeCameraFromJSObject(p js.Value, ctx js.Value) *AnaglyphFreeCamera {
	return &AnaglyphFreeCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// AnaglyphFreeCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnaglyphFreeCameraArrayToJSArray(array []*AnaglyphFreeCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnaglyphFreeCamera returns a new AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera
func (ba *Babylon) NewAnaglyphFreeCamera(name string, position *Vector3, interaxialDistance float64, scene *Scene) *AnaglyphFreeCamera {

	args := make([]interface{}, 0, 4+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, interaxialDistance)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("AnaglyphFreeCamera").New(args...)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// AnaglyphFreeCameraAttachControlOpts contains optional parameters for AnaglyphFreeCamera.AttachControl.
type AnaglyphFreeCameraAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#attachcontrol
func (a *AnaglyphFreeCamera) AttachControl(element js.Value, opts *AnaglyphFreeCameraAttachControlOpts) {
	if opts == nil {
		opts = &AnaglyphFreeCameraAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	a.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#detachcontrol
func (a *AnaglyphFreeCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	a.p.Call("detachControl", args...)
}

// Dispose calls the Dispose method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#dispose
func (a *AnaglyphFreeCamera) Dispose() {

	a.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#getclassname
func (a *AnaglyphFreeCamera) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// GetFrontPosition calls the GetFrontPosition method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#getfrontposition
func (a *AnaglyphFreeCamera) GetFrontPosition(distance float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	retVal := a.p.Call("getFrontPosition", args...)
	return Vector3FromJSObject(retVal, a.ctx)
}

// GetTarget calls the GetTarget method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#gettarget
func (a *AnaglyphFreeCamera) GetTarget() *Vector3 {

	retVal := a.p.Call("getTarget")
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetTarget calls the SetTarget method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#settarget
func (a *AnaglyphFreeCamera) SetTarget(target *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	a.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#storestate
func (a *AnaglyphFreeCamera) StoreState() *Camera {

	retVal := a.p.Call("storeState")
	return CameraFromJSObject(retVal, a.ctx)
}

/*

// AngularSensibility returns the AngularSensibility property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#angularsensibility
func (a *AnaglyphFreeCamera) AngularSensibility(angularSensibility float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(angularSensibility)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#angularsensibility
func (a *AnaglyphFreeCamera) SetAngularSensibility(angularSensibility float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(angularSensibility)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// ApplyGravity returns the ApplyGravity property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#applygravity
func (a *AnaglyphFreeCamera) ApplyGravity(applyGravity bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(applyGravity)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetApplyGravity sets the ApplyGravity property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#applygravity
func (a *AnaglyphFreeCamera) SetApplyGravity(applyGravity bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(applyGravity)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#cameradirection
func (a *AnaglyphFreeCamera) CameraDirection(cameraDirection *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(cameraDirection.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#cameradirection
func (a *AnaglyphFreeCamera) SetCameraDirection(cameraDirection *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(cameraDirection.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// CameraRotation returns the CameraRotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#camerarotation
func (a *AnaglyphFreeCamera) CameraRotation(cameraRotation *Vector2) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(cameraRotation.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraRotation sets the CameraRotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#camerarotation
func (a *AnaglyphFreeCamera) SetCameraRotation(cameraRotation *Vector2) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(cameraRotation.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// CheckCollisions returns the CheckCollisions property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#checkcollisions
func (a *AnaglyphFreeCamera) CheckCollisions(checkCollisions bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(checkCollisions)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetCheckCollisions sets the CheckCollisions property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#checkcollisions
func (a *AnaglyphFreeCamera) SetCheckCollisions(checkCollisions bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(checkCollisions)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// CollisionMask returns the CollisionMask property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#collisionmask
func (a *AnaglyphFreeCamera) CollisionMask(collisionMask float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(collisionMask)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetCollisionMask sets the CollisionMask property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#collisionmask
func (a *AnaglyphFreeCamera) SetCollisionMask(collisionMask float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(collisionMask)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// Ellipsoid returns the Ellipsoid property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#ellipsoid
func (a *AnaglyphFreeCamera) Ellipsoid(ellipsoid *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(ellipsoid.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#ellipsoid
func (a *AnaglyphFreeCamera) SetEllipsoid(ellipsoid *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(ellipsoid.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// EllipsoidOffset returns the EllipsoidOffset property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#ellipsoidoffset
func (a *AnaglyphFreeCamera) EllipsoidOffset(ellipsoidOffset *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(ellipsoidOffset.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#ellipsoidoffset
func (a *AnaglyphFreeCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(ellipsoidOffset.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#inputs
func (a *AnaglyphFreeCamera) Inputs(inputs *FreeCameraInputsManager) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(inputs.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#inputs
func (a *AnaglyphFreeCamera) SetInputs(inputs *FreeCameraInputsManager) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(inputs.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysdown
func (a *AnaglyphFreeCamera) KeysDown(keysDown float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysDown)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysdown
func (a *AnaglyphFreeCamera) SetKeysDown(keysDown float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysDown)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysleft
func (a *AnaglyphFreeCamera) KeysLeft(keysLeft float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysLeft)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysleft
func (a *AnaglyphFreeCamera) SetKeysLeft(keysLeft float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysLeft)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysright
func (a *AnaglyphFreeCamera) KeysRight(keysRight float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysRight)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysright
func (a *AnaglyphFreeCamera) SetKeysRight(keysRight float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysRight)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysup
func (a *AnaglyphFreeCamera) KeysUp(keysUp float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysUp)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#keysup
func (a *AnaglyphFreeCamera) SetKeysUp(keysUp float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(keysUp)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// LockedTarget returns the LockedTarget property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#lockedtarget
func (a *AnaglyphFreeCamera) LockedTarget(lockedTarget interface{}) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(lockedTarget)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetLockedTarget sets the LockedTarget property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#lockedtarget
func (a *AnaglyphFreeCamera) SetLockedTarget(lockedTarget interface{}) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(lockedTarget)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// NoRotationConstraint returns the NoRotationConstraint property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#norotationconstraint
func (a *AnaglyphFreeCamera) NoRotationConstraint(noRotationConstraint bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(noRotationConstraint)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetNoRotationConstraint sets the NoRotationConstraint property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#norotationconstraint
func (a *AnaglyphFreeCamera) SetNoRotationConstraint(noRotationConstraint bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(noRotationConstraint)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// OnCollide returns the OnCollide property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#oncollide
func (a *AnaglyphFreeCamera) OnCollide(onCollide func()) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnCollide sets the OnCollide property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#oncollide
func (a *AnaglyphFreeCamera) SetOnCollide(onCollide func()) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#rotation
func (a *AnaglyphFreeCamera) Rotation(rotation *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(rotation.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#rotation
func (a *AnaglyphFreeCamera) SetRotation(rotation *Vector3) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(rotation.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#rotationquaternion
func (a *AnaglyphFreeCamera) RotationQuaternion(rotationQuaternion *Quaternion) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(rotationQuaternion.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#rotationquaternion
func (a *AnaglyphFreeCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(rotationQuaternion.JSObject())
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// Speed returns the Speed property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#speed
func (a *AnaglyphFreeCamera) Speed(speed float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(speed)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetSpeed sets the Speed property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#speed
func (a *AnaglyphFreeCamera) SetSpeed(speed float64) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(speed)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// UpdateUpVectorFromRotation returns the UpdateUpVectorFromRotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#updateupvectorfromrotation
func (a *AnaglyphFreeCamera) UpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(updateUpVectorFromRotation)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// SetUpdateUpVectorFromRotation sets the UpdateUpVectorFromRotation property of class AnaglyphFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#updateupvectorfromrotation
func (a *AnaglyphFreeCamera) SetUpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera").New(updateUpVectorFromRotation)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

*/
