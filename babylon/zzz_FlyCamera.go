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
// https://doc.babylonjs.com/api/classes/babylon.flycamera#constructor
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

// AttachControl calls the AttachControl method on the FlyCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#attachcontrol
func (f *FlyCamera) AttachControl(element js.Value, noPreventDefault bool) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, element)

	args = append(args, noPreventDefault)

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

// AngularSensibility returns the AngularSensibility property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#angularsensibility
func (f *FlyCamera) AngularSensibility() float64 {
	retVal := f.p.Get("angularSensibility")
	return retVal.Float()
}

// SetAngularSensibility sets the AngularSensibility property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#angularsensibility
func (f *FlyCamera) SetAngularSensibility(angularSensibility float64) *FlyCamera {
	f.p.Set("angularSensibility", angularSensibility)
	return f
}

// ApplyGravity returns the ApplyGravity property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#applygravity
func (f *FlyCamera) ApplyGravity() bool {
	retVal := f.p.Get("applyGravity")
	return retVal.Bool()
}

// SetApplyGravity sets the ApplyGravity property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#applygravity
func (f *FlyCamera) SetApplyGravity(applyGravity bool) *FlyCamera {
	f.p.Set("applyGravity", applyGravity)
	return f
}

// BankedTurn returns the BankedTurn property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturn
func (f *FlyCamera) BankedTurn() bool {
	retVal := f.p.Get("bankedTurn")
	return retVal.Bool()
}

// SetBankedTurn sets the BankedTurn property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturn
func (f *FlyCamera) SetBankedTurn(bankedTurn bool) *FlyCamera {
	f.p.Set("bankedTurn", bankedTurn)
	return f
}

// BankedTurnLimit returns the BankedTurnLimit property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnlimit
func (f *FlyCamera) BankedTurnLimit() float64 {
	retVal := f.p.Get("bankedTurnLimit")
	return retVal.Float()
}

// SetBankedTurnLimit sets the BankedTurnLimit property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnlimit
func (f *FlyCamera) SetBankedTurnLimit(bankedTurnLimit float64) *FlyCamera {
	f.p.Set("bankedTurnLimit", bankedTurnLimit)
	return f
}

// BankedTurnMultiplier returns the BankedTurnMultiplier property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnmultiplier
func (f *FlyCamera) BankedTurnMultiplier() float64 {
	retVal := f.p.Get("bankedTurnMultiplier")
	return retVal.Float()
}

// SetBankedTurnMultiplier sets the BankedTurnMultiplier property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#bankedturnmultiplier
func (f *FlyCamera) SetBankedTurnMultiplier(bankedTurnMultiplier float64) *FlyCamera {
	f.p.Set("bankedTurnMultiplier", bankedTurnMultiplier)
	return f
}

// CameraDirection returns the CameraDirection property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#cameradirection
func (f *FlyCamera) CameraDirection() *Vector3 {
	retVal := f.p.Get("cameraDirection")
	return Vector3FromJSObject(retVal, f.ctx)
}

// SetCameraDirection sets the CameraDirection property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#cameradirection
func (f *FlyCamera) SetCameraDirection(cameraDirection *Vector3) *FlyCamera {
	f.p.Set("cameraDirection", cameraDirection.JSObject())
	return f
}

// CheckCollisions returns the CheckCollisions property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#checkcollisions
func (f *FlyCamera) CheckCollisions() bool {
	retVal := f.p.Get("checkCollisions")
	return retVal.Bool()
}

// SetCheckCollisions sets the CheckCollisions property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#checkcollisions
func (f *FlyCamera) SetCheckCollisions(checkCollisions bool) *FlyCamera {
	f.p.Set("checkCollisions", checkCollisions)
	return f
}

// CollisionMask returns the CollisionMask property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#collisionmask
func (f *FlyCamera) CollisionMask() float64 {
	retVal := f.p.Get("collisionMask")
	return retVal.Float()
}

// SetCollisionMask sets the CollisionMask property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#collisionmask
func (f *FlyCamera) SetCollisionMask(collisionMask float64) *FlyCamera {
	f.p.Set("collisionMask", collisionMask)
	return f
}

// Ellipsoid returns the Ellipsoid property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoid
func (f *FlyCamera) Ellipsoid() *Vector3 {
	retVal := f.p.Get("ellipsoid")
	return Vector3FromJSObject(retVal, f.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoid
func (f *FlyCamera) SetEllipsoid(ellipsoid *Vector3) *FlyCamera {
	f.p.Set("ellipsoid", ellipsoid.JSObject())
	return f
}

// EllipsoidOffset returns the EllipsoidOffset property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoidoffset
func (f *FlyCamera) EllipsoidOffset() *Vector3 {
	retVal := f.p.Get("ellipsoidOffset")
	return Vector3FromJSObject(retVal, f.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#ellipsoidoffset
func (f *FlyCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *FlyCamera {
	f.p.Set("ellipsoidOffset", ellipsoidOffset.JSObject())
	return f
}

// Inputs returns the Inputs property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#inputs
func (f *FlyCamera) Inputs() *FlyCameraInputsManager {
	retVal := f.p.Get("inputs")
	return FlyCameraInputsManagerFromJSObject(retVal, f.ctx)
}

// SetInputs sets the Inputs property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#inputs
func (f *FlyCamera) SetInputs(inputs *FlyCameraInputsManager) *FlyCamera {
	f.p.Set("inputs", inputs.JSObject())
	return f
}

// KeysBackward returns the KeysBackward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysbackward
func (f *FlyCamera) KeysBackward() []float64 {
	retVal := f.p.Get("keysBackward")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysBackward sets the KeysBackward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysbackward
func (f *FlyCamera) SetKeysBackward(keysBackward []float64) *FlyCamera {
	f.p.Set("keysBackward", keysBackward)
	return f
}

// KeysDown returns the KeysDown property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysdown
func (f *FlyCamera) KeysDown() []float64 {
	retVal := f.p.Get("keysDown")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysDown sets the KeysDown property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysdown
func (f *FlyCamera) SetKeysDown(keysDown []float64) *FlyCamera {
	f.p.Set("keysDown", keysDown)
	return f
}

// KeysForward returns the KeysForward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysforward
func (f *FlyCamera) KeysForward() []float64 {
	retVal := f.p.Get("keysForward")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysForward sets the KeysForward property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysforward
func (f *FlyCamera) SetKeysForward(keysForward []float64) *FlyCamera {
	f.p.Set("keysForward", keysForward)
	return f
}

// KeysLeft returns the KeysLeft property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysleft
func (f *FlyCamera) KeysLeft() []float64 {
	retVal := f.p.Get("keysLeft")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysLeft sets the KeysLeft property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysleft
func (f *FlyCamera) SetKeysLeft(keysLeft []float64) *FlyCamera {
	f.p.Set("keysLeft", keysLeft)
	return f
}

// KeysRight returns the KeysRight property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysright
func (f *FlyCamera) KeysRight() []float64 {
	retVal := f.p.Get("keysRight")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysRight sets the KeysRight property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysright
func (f *FlyCamera) SetKeysRight(keysRight []float64) *FlyCamera {
	f.p.Set("keysRight", keysRight)
	return f
}

// KeysUp returns the KeysUp property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysup
func (f *FlyCamera) KeysUp() []float64 {
	retVal := f.p.Get("keysUp")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysUp sets the KeysUp property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#keysup
func (f *FlyCamera) SetKeysUp(keysUp []float64) *FlyCamera {
	f.p.Set("keysUp", keysUp)
	return f
}

// OnCollide returns the OnCollide property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#oncollide
func (f *FlyCamera) OnCollide() js.Value {
	retVal := f.p.Get("onCollide")
	return retVal
}

// SetOnCollide sets the OnCollide property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#oncollide
func (f *FlyCamera) SetOnCollide(onCollide JSFunc) *FlyCamera {
	f.p.Set("onCollide", js.FuncOf(onCollide))
	return f
}

// RollCorrect returns the RollCorrect property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rollcorrect
func (f *FlyCamera) RollCorrect() float64 {
	retVal := f.p.Get("rollCorrect")
	return retVal.Float()
}

// SetRollCorrect sets the RollCorrect property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rollcorrect
func (f *FlyCamera) SetRollCorrect(rollCorrect float64) *FlyCamera {
	f.p.Set("rollCorrect", rollCorrect)
	return f
}

// RotationQuaternion returns the RotationQuaternion property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rotationquaternion
func (f *FlyCamera) RotationQuaternion() *Quaternion {
	retVal := f.p.Get("rotationQuaternion")
	return QuaternionFromJSObject(retVal, f.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#rotationquaternion
func (f *FlyCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *FlyCamera {
	f.p.Set("rotationQuaternion", rotationQuaternion.JSObject())
	return f
}

// _trackRoll returns the _trackRoll property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#_trackroll
func (f *FlyCamera) _trackRoll() float64 {
	retVal := f.p.Get("_trackRoll")
	return retVal.Float()
}

// Set_trackRoll sets the _trackRoll property of class FlyCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamera#_trackroll
func (f *FlyCamera) Set_trackRoll(_trackRoll float64) *FlyCamera {
	f.p.Set("_trackRoll", _trackRoll)
	return f
}
