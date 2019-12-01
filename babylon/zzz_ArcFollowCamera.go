// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcFollowCamera represents a babylon.js ArcFollowCamera.
// Arc Rotate version of the follow camera.
// It still follows a Defined mesh but in an Arc Rotate Camera fashion.
//
// See: http://doc.babylonjs.com/features/cameras#follow-camera
type ArcFollowCamera struct {
	*TargetCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcFollowCamera) JSObject() js.Value { return a.p }

// ArcFollowCamera returns a ArcFollowCamera JavaScript class.
func (ba *Babylon) ArcFollowCamera() *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera")
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// ArcFollowCameraFromJSObject returns a wrapped ArcFollowCamera JavaScript class.
func ArcFollowCameraFromJSObject(p js.Value, ctx js.Value) *ArcFollowCamera {
	return &ArcFollowCamera{TargetCamera: TargetCameraFromJSObject(p, ctx), ctx: ctx}
}

// NewArcFollowCamera returns a new ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera
func (ba *Babylon) NewArcFollowCamera(name string, alpha float64, beta float64, radius float64, target *AbstractMesh, scene *Scene) *ArcFollowCamera {

	args := make([]interface{}, 0, 6+0)

	args = append(args, name)
	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, radius)
	args = append(args, target.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("ArcFollowCamera").New(args...)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#getclassname
func (a *ArcFollowCamera) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getClassName", args...)
	return retVal.String()
}

// GetFrontPosition calls the GetFrontPosition method on the ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#getfrontposition
func (a *ArcFollowCamera) GetFrontPosition(distance float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	retVal := a.p.Call("getFrontPosition", args...)
	return Vector3FromJSObject(retVal, a.ctx)
}

// GetTarget calls the GetTarget method on the ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#gettarget
func (a *ArcFollowCamera) GetTarget() *Vector3 {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("getTarget", args...)
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetTarget calls the SetTarget method on the ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#settarget
func (a *ArcFollowCamera) SetTarget(target *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	a.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the ArcFollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#storestate
func (a *ArcFollowCamera) StoreState() *Camera {

	args := make([]interface{}, 0, 0+0)

	retVal := a.p.Call("storeState", args...)
	return CameraFromJSObject(retVal, a.ctx)
}

/*

// Alpha returns the Alpha property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#alpha
func (a *ArcFollowCamera) Alpha(alpha float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(alpha)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetAlpha sets the Alpha property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#alpha
func (a *ArcFollowCamera) SetAlpha(alpha float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(alpha)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// Beta returns the Beta property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#beta
func (a *ArcFollowCamera) Beta(beta float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(beta)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetBeta sets the Beta property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#beta
func (a *ArcFollowCamera) SetBeta(beta float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(beta)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#cameradirection
func (a *ArcFollowCamera) CameraDirection(cameraDirection *Vector3) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(cameraDirection.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#cameradirection
func (a *ArcFollowCamera) SetCameraDirection(cameraDirection *Vector3) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(cameraDirection.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// CameraRotation returns the CameraRotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#camerarotation
func (a *ArcFollowCamera) CameraRotation(cameraRotation *Vector2) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(cameraRotation.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetCameraRotation sets the CameraRotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#camerarotation
func (a *ArcFollowCamera) SetCameraRotation(cameraRotation *Vector2) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(cameraRotation.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// LockedTarget returns the LockedTarget property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#lockedtarget
func (a *ArcFollowCamera) LockedTarget(lockedTarget interface{}) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(lockedTarget)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetLockedTarget sets the LockedTarget property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#lockedtarget
func (a *ArcFollowCamera) SetLockedTarget(lockedTarget interface{}) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(lockedTarget)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// NoRotationConstraint returns the NoRotationConstraint property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#norotationconstraint
func (a *ArcFollowCamera) NoRotationConstraint(noRotationConstraint bool) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(noRotationConstraint)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetNoRotationConstraint sets the NoRotationConstraint property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#norotationconstraint
func (a *ArcFollowCamera) SetNoRotationConstraint(noRotationConstraint bool) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(noRotationConstraint)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// Radius returns the Radius property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#radius
func (a *ArcFollowCamera) Radius(radius float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(radius)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetRadius sets the Radius property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#radius
func (a *ArcFollowCamera) SetRadius(radius float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(radius)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#rotation
func (a *ArcFollowCamera) Rotation(rotation *Vector3) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(rotation.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#rotation
func (a *ArcFollowCamera) SetRotation(rotation *Vector3) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(rotation.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#rotationquaternion
func (a *ArcFollowCamera) RotationQuaternion(rotationQuaternion *Quaternion) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(rotationQuaternion.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#rotationquaternion
func (a *ArcFollowCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(rotationQuaternion.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// Speed returns the Speed property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#speed
func (a *ArcFollowCamera) Speed(speed float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(speed)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetSpeed sets the Speed property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#speed
func (a *ArcFollowCamera) SetSpeed(speed float64) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(speed)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#target
func (a *ArcFollowCamera) Target(target *AbstractMesh) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(target.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#target
func (a *ArcFollowCamera) SetTarget(target *AbstractMesh) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(target.JSObject())
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// UpdateUpVectorFromRotation returns the UpdateUpVectorFromRotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#updateupvectorfromrotation
func (a *ArcFollowCamera) UpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(updateUpVectorFromRotation)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

// SetUpdateUpVectorFromRotation sets the UpdateUpVectorFromRotation property of class ArcFollowCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcfollowcamera#updateupvectorfromrotation
func (a *ArcFollowCamera) SetUpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *ArcFollowCamera {
	p := ba.ctx.Get("ArcFollowCamera").New(updateUpVectorFromRotation)
	return ArcFollowCameraFromJSObject(p, ba.ctx)
}

*/
