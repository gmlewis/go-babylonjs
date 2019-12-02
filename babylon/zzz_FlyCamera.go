// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FlyCamera represents a babylon.js FlyCamera.
// This is a flying camera, designed for 3D movement and rotation in all directions,
// such as in a 3D Space Shooter or a Flight Simulator.
type FlyCamera struct {
	*TargetCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FlyCamera) JSObject() js.Value { return f.p }

// FlyCamera returns a FlyCamera JavaScript class.
func (ba *Babylon) FlyCamera() *FlyCamera {
	p := ba.ctx.Get("FlyCamera")
	return FlyCameraFromJSObject(p, ba.ctx)
}

// FlyCameraFromJSObject returns a wrapped FlyCamera JavaScript class.
func FlyCameraFromJSObject(p js.Value, ctx js.Value) *FlyCamera {
	return &FlyCamera{TargetCamera: TargetCameraFromJSObject(p, ctx), ctx: ctx}
}

// FlyCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func FlyCameraArrayToJSArray(array []*FlyCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFlyCameraOpts contains optional parameters for NewFlyCamera.
type NewFlyCameraOpts struct {
	SetActiveOnSceneIfNoneActive *bool
}

// NewFlyCamera returns a new FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera
func (ba *Babylon) NewFlyCamera(name string, position *Vector3, scene *Scene, opts *NewFlyCameraOpts) *FlyCamera {
	if opts == nil {
		opts = &NewFlyCameraOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	if opts.SetActiveOnSceneIfNoneActive == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SetActiveOnSceneIfNoneActive)
	}

	p := ba.ctx.Get("FlyCamera").New(args...)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// FlyCameraAttachControlOpts contains optional parameters for FlyCamera.AttachControl.
type FlyCameraAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#attachcontrol
func (f *FlyCamera) AttachControl(element js.Value, opts *FlyCameraAttachControlOpts) {
	if opts == nil {
		opts = &FlyCameraAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	f.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#detachcontrol
func (f *FlyCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// Dispose calls the Dispose method on the FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#dispose
func (f *FlyCamera) Dispose() {

	f.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#getclassname
func (f *FlyCamera) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

/*

// AngularSensibility returns the AngularSensibility property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#angularsensibility
func (f *FlyCamera) AngularSensibility(angularSensibility float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(angularSensibility)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#angularsensibility
func (f *FlyCamera) SetAngularSensibility(angularSensibility float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(angularSensibility)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// ApplyGravity returns the ApplyGravity property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#applygravity
func (f *FlyCamera) ApplyGravity(applyGravity bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(applyGravity)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetApplyGravity sets the ApplyGravity property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#applygravity
func (f *FlyCamera) SetApplyGravity(applyGravity bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(applyGravity)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// BankedTurn returns the BankedTurn property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturn
func (f *FlyCamera) BankedTurn(bankedTurn bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurn)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetBankedTurn sets the BankedTurn property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturn
func (f *FlyCamera) SetBankedTurn(bankedTurn bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurn)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// BankedTurnLimit returns the BankedTurnLimit property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnlimit
func (f *FlyCamera) BankedTurnLimit(bankedTurnLimit float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurnLimit)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetBankedTurnLimit sets the BankedTurnLimit property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnlimit
func (f *FlyCamera) SetBankedTurnLimit(bankedTurnLimit float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurnLimit)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// BankedTurnMultiplier returns the BankedTurnMultiplier property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnmultiplier
func (f *FlyCamera) BankedTurnMultiplier(bankedTurnMultiplier float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurnMultiplier)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetBankedTurnMultiplier sets the BankedTurnMultiplier property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnmultiplier
func (f *FlyCamera) SetBankedTurnMultiplier(bankedTurnMultiplier float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(bankedTurnMultiplier)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#cameradirection
func (f *FlyCamera) CameraDirection(cameraDirection *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(cameraDirection.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#cameradirection
func (f *FlyCamera) SetCameraDirection(cameraDirection *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(cameraDirection.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// CheckCollisions returns the CheckCollisions property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#checkcollisions
func (f *FlyCamera) CheckCollisions(checkCollisions bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(checkCollisions)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetCheckCollisions sets the CheckCollisions property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#checkcollisions
func (f *FlyCamera) SetCheckCollisions(checkCollisions bool) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(checkCollisions)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// CollisionMask returns the CollisionMask property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#collisionmask
func (f *FlyCamera) CollisionMask(collisionMask float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(collisionMask)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetCollisionMask sets the CollisionMask property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#collisionmask
func (f *FlyCamera) SetCollisionMask(collisionMask float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(collisionMask)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// Ellipsoid returns the Ellipsoid property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoid
func (f *FlyCamera) Ellipsoid(ellipsoid *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(ellipsoid.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoid
func (f *FlyCamera) SetEllipsoid(ellipsoid *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(ellipsoid.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// EllipsoidOffset returns the EllipsoidOffset property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoidoffset
func (f *FlyCamera) EllipsoidOffset(ellipsoidOffset *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(ellipsoidOffset.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoidoffset
func (f *FlyCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(ellipsoidOffset.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#inputs
func (f *FlyCamera) Inputs(inputs *FlyCameraInputsManager) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(inputs.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#inputs
func (f *FlyCamera) SetInputs(inputs *FlyCameraInputsManager) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(inputs.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysBackward returns the KeysBackward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysbackward
func (f *FlyCamera) KeysBackward(keysBackward float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysBackward)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysBackward sets the KeysBackward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysbackward
func (f *FlyCamera) SetKeysBackward(keysBackward float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysBackward)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysdown
func (f *FlyCamera) KeysDown(keysDown float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysDown)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysdown
func (f *FlyCamera) SetKeysDown(keysDown float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysDown)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysForward returns the KeysForward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysforward
func (f *FlyCamera) KeysForward(keysForward float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysForward)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysForward sets the KeysForward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysforward
func (f *FlyCamera) SetKeysForward(keysForward float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysForward)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysleft
func (f *FlyCamera) KeysLeft(keysLeft float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysLeft)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysleft
func (f *FlyCamera) SetKeysLeft(keysLeft float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysLeft)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysright
func (f *FlyCamera) KeysRight(keysRight float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysRight)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysright
func (f *FlyCamera) SetKeysRight(keysRight float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysRight)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysup
func (f *FlyCamera) KeysUp(keysUp float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysUp)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysup
func (f *FlyCamera) SetKeysUp(keysUp float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(keysUp)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// OnCollide returns the OnCollide property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#oncollide
func (f *FlyCamera) OnCollide(onCollide func()) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetOnCollide sets the OnCollide property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#oncollide
func (f *FlyCamera) SetOnCollide(onCollide func()) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return FlyCameraFromJSObject(p, ba.ctx)
}

// RollCorrect returns the RollCorrect property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rollcorrect
func (f *FlyCamera) RollCorrect(rollCorrect float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(rollCorrect)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetRollCorrect sets the RollCorrect property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rollcorrect
func (f *FlyCamera) SetRollCorrect(rollCorrect float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(rollCorrect)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rotationquaternion
func (f *FlyCamera) RotationQuaternion(rotationQuaternion *Quaternion) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(rotationQuaternion.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rotationquaternion
func (f *FlyCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(rotationQuaternion.JSObject())
	return FlyCameraFromJSObject(p, ba.ctx)
}

// _trackRoll returns the _trackRoll property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#_trackroll
func (f *FlyCamera) _trackRoll(_trackRoll float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(_trackRoll)
	return FlyCameraFromJSObject(p, ba.ctx)
}

// Set_trackRoll sets the _trackRoll property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#_trackroll
func (f *FlyCamera) Set_trackRoll(_trackRoll float64) *FlyCamera {
	p := ba.ctx.Get("FlyCamera").New(_trackRoll)
	return FlyCameraFromJSObject(p, ba.ctx)
}

*/
