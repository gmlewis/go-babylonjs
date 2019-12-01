// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebVRFreeCamera represents a babylon.js WebVRFreeCamera.
// This represents a WebVR camera.
// The WebVR camera is Babylon&amp;#39;s simple interface to interaction with Windows Mixed Reality, HTC Vive and Oculus Rift.
//
// See: http://doc.babylonjs.com/how_to/webvr_camera
type WebVRFreeCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebVRFreeCamera) JSObject() js.Value { return w.p }

// WebVRFreeCamera returns a WebVRFreeCamera JavaScript class.
func (ba *Babylon) WebVRFreeCamera() *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera")
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// WebVRFreeCameraFromJSObject returns a wrapped WebVRFreeCamera JavaScript class.
func WebVRFreeCameraFromJSObject(p js.Value, ctx js.Value) *WebVRFreeCamera {
	return &WebVRFreeCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// WebVRFreeCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebVRFreeCameraArrayToJSArray(array []*WebVRFreeCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWebVRFreeCameraOpts contains optional parameters for NewWebVRFreeCamera.
type NewWebVRFreeCameraOpts struct {
	WebVROptions js.Value
}

// NewWebVRFreeCamera returns a new WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera
func (ba *Babylon) NewWebVRFreeCamera(name string, position *Vector3, scene *Scene, opts *NewWebVRFreeCameraOpts) *WebVRFreeCamera {
	if opts == nil {
		opts = &NewWebVRFreeCameraOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	if opts.WebVROptions == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.WebVROptions)
	}

	p := ba.ctx.Get("WebVRFreeCamera").New(args...)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// WebVRFreeCameraAttachControlOpts contains optional parameters for WebVRFreeCamera.AttachControl.
type WebVRFreeCameraAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#attachcontrol
func (w *WebVRFreeCamera) AttachControl(element js.Value, opts *WebVRFreeCameraAttachControlOpts) {
	if opts == nil {
		opts = &WebVRFreeCameraAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	w.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#detachcontrol
func (w *WebVRFreeCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	w.p.Call("detachControl", args...)
}

// DeviceDistanceToRoomGround calls the DeviceDistanceToRoomGround method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicedistancetoroomground
func (w *WebVRFreeCamera) DeviceDistanceToRoomGround() float64 {

	retVal := w.p.Call("deviceDistanceToRoomGround")
	return retVal.Float()
}

// Dispose calls the Dispose method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#dispose
func (w *WebVRFreeCamera) Dispose() {

	w.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#getclassname
func (w *WebVRFreeCamera) GetClassName() string {

	retVal := w.p.Call("getClassName")
	return retVal.String()
}

// GetControllerByName calls the GetControllerByName method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#getcontrollerbyname
func (w *WebVRFreeCamera) GetControllerByName(name string) *WebVRController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := w.p.Call("getControllerByName", args...)
	return WebVRControllerFromJSObject(retVal, w.ctx)
}

// WebVRFreeCameraGetForwardRayOpts contains optional parameters for WebVRFreeCamera.GetForwardRay.
type WebVRFreeCameraGetForwardRayOpts struct {
	Length *float64
}

// GetForwardRay calls the GetForwardRay method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#getforwardray
func (w *WebVRFreeCamera) GetForwardRay(opts *WebVRFreeCameraGetForwardRayOpts) *Ray {
	if opts == nil {
		opts = &WebVRFreeCameraGetForwardRayOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	retVal := w.p.Call("getForwardRay", args...)
	return RayFromJSObject(retVal, w.ctx)
}

// GetFrontPosition calls the GetFrontPosition method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#getfrontposition
func (w *WebVRFreeCamera) GetFrontPosition(distance float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, distance)

	retVal := w.p.Call("getFrontPosition", args...)
	return Vector3FromJSObject(retVal, w.ctx)
}

// GetTarget calls the GetTarget method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#gettarget
func (w *WebVRFreeCamera) GetTarget() *Vector3 {

	retVal := w.p.Call("getTarget")
	return Vector3FromJSObject(retVal, w.ctx)
}

// InitControllers calls the InitControllers method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#initcontrollers
func (w *WebVRFreeCamera) InitControllers() {

	w.p.Call("initControllers")
}

// ResetToCurrentRotation calls the ResetToCurrentRotation method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#resettocurrentrotation
func (w *WebVRFreeCamera) ResetToCurrentRotation() {

	w.p.Call("resetToCurrentRotation")
}

// SetTarget calls the SetTarget method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#settarget
func (w *WebVRFreeCamera) SetTarget(target *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	w.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#storestate
func (w *WebVRFreeCamera) StoreState() *Camera {

	retVal := w.p.Call("storeState")
	return CameraFromJSObject(retVal, w.ctx)
}

// Update calls the Update method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#update
func (w *WebVRFreeCamera) Update() {

	w.p.Call("update")
}

// UpdateFromDevice calls the UpdateFromDevice method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#updatefromdevice
func (w *WebVRFreeCamera) UpdateFromDevice(poseData js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, poseData)

	w.p.Call("updateFromDevice", args...)
}

// WebVRFreeCameraUseStandingMatrixOpts contains optional parameters for WebVRFreeCamera.UseStandingMatrix.
type WebVRFreeCameraUseStandingMatrixOpts struct {
	Callback *func()
}

// UseStandingMatrix calls the UseStandingMatrix method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#usestandingmatrix
func (w *WebVRFreeCamera) UseStandingMatrix(opts *WebVRFreeCameraUseStandingMatrixOpts) {
	if opts == nil {
		opts = &WebVRFreeCameraUseStandingMatrixOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Callback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Callback)
	}

	w.p.Call("useStandingMatrix", args...)
}

// UseStandingMatrixAsync calls the UseStandingMatrixAsync method on the WebVRFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#usestandingmatrixasync
func (w *WebVRFreeCamera) UseStandingMatrixAsync() *Promise {

	retVal := w.p.Call("useStandingMatrixAsync")
	return PromiseFromJSObject(retVal, w.ctx)
}

/*

// AngularSensibility returns the AngularSensibility property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#angularsensibility
func (w *WebVRFreeCamera) AngularSensibility(angularSensibility float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(angularSensibility)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#angularsensibility
func (w *WebVRFreeCamera) SetAngularSensibility(angularSensibility float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(angularSensibility)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// ApplyGravity returns the ApplyGravity property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#applygravity
func (w *WebVRFreeCamera) ApplyGravity(applyGravity bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(applyGravity)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetApplyGravity sets the ApplyGravity property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#applygravity
func (w *WebVRFreeCamera) SetApplyGravity(applyGravity bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(applyGravity)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// CameraDirection returns the CameraDirection property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#cameradirection
func (w *WebVRFreeCamera) CameraDirection(cameraDirection *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(cameraDirection.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraDirection sets the CameraDirection property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#cameradirection
func (w *WebVRFreeCamera) SetCameraDirection(cameraDirection *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(cameraDirection.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// CameraRotation returns the CameraRotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#camerarotation
func (w *WebVRFreeCamera) CameraRotation(cameraRotation *Vector2) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(cameraRotation.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetCameraRotation sets the CameraRotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#camerarotation
func (w *WebVRFreeCamera) SetCameraRotation(cameraRotation *Vector2) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(cameraRotation.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// CheckCollisions returns the CheckCollisions property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#checkcollisions
func (w *WebVRFreeCamera) CheckCollisions(checkCollisions bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(checkCollisions)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetCheckCollisions sets the CheckCollisions property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#checkcollisions
func (w *WebVRFreeCamera) SetCheckCollisions(checkCollisions bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(checkCollisions)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// CollisionMask returns the CollisionMask property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#collisionmask
func (w *WebVRFreeCamera) CollisionMask(collisionMask float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(collisionMask)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetCollisionMask sets the CollisionMask property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#collisionmask
func (w *WebVRFreeCamera) SetCollisionMask(collisionMask float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(collisionMask)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// Controllers returns the Controllers property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#controllers
func (w *WebVRFreeCamera) Controllers(controllers []*WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(controllers)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetControllers sets the Controllers property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#controllers
func (w *WebVRFreeCamera) SetControllers(controllers []*WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(controllers)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// DevicePosition returns the DevicePosition property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#deviceposition
func (w *WebVRFreeCamera) DevicePosition(devicePosition *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(devicePosition.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetDevicePosition sets the DevicePosition property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#deviceposition
func (w *WebVRFreeCamera) SetDevicePosition(devicePosition *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(devicePosition.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// DeviceRotationQuaternion returns the DeviceRotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicerotationquaternion
func (w *WebVRFreeCamera) DeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(deviceRotationQuaternion.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetDeviceRotationQuaternion sets the DeviceRotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicerotationquaternion
func (w *WebVRFreeCamera) SetDeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(deviceRotationQuaternion.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// DeviceScaleFactor returns the DeviceScaleFactor property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicescalefactor
func (w *WebVRFreeCamera) DeviceScaleFactor(deviceScaleFactor float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(deviceScaleFactor)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetDeviceScaleFactor sets the DeviceScaleFactor property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicescalefactor
func (w *WebVRFreeCamera) SetDeviceScaleFactor(deviceScaleFactor float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(deviceScaleFactor)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// Ellipsoid returns the Ellipsoid property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#ellipsoid
func (w *WebVRFreeCamera) Ellipsoid(ellipsoid *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(ellipsoid.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoid sets the Ellipsoid property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#ellipsoid
func (w *WebVRFreeCamera) SetEllipsoid(ellipsoid *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(ellipsoid.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// EllipsoidOffset returns the EllipsoidOffset property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#ellipsoidoffset
func (w *WebVRFreeCamera) EllipsoidOffset(ellipsoidOffset *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(ellipsoidOffset.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetEllipsoidOffset sets the EllipsoidOffset property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#ellipsoidoffset
func (w *WebVRFreeCamera) SetEllipsoidOffset(ellipsoidOffset *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(ellipsoidOffset.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#inputs
func (w *WebVRFreeCamera) Inputs(inputs *FreeCameraInputsManager) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(inputs.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#inputs
func (w *WebVRFreeCamera) SetInputs(inputs *FreeCameraInputsManager) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(inputs.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysdown
func (w *WebVRFreeCamera) KeysDown(keysDown float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysDown)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysdown
func (w *WebVRFreeCamera) SetKeysDown(keysDown float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysDown)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysleft
func (w *WebVRFreeCamera) KeysLeft(keysLeft float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysLeft)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysleft
func (w *WebVRFreeCamera) SetKeysLeft(keysLeft float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysLeft)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysright
func (w *WebVRFreeCamera) KeysRight(keysRight float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysRight)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysright
func (w *WebVRFreeCamera) SetKeysRight(keysRight float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysRight)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysup
func (w *WebVRFreeCamera) KeysUp(keysUp float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysUp)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#keysup
func (w *WebVRFreeCamera) SetKeysUp(keysUp float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(keysUp)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// LeftController returns the LeftController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#leftcontroller
func (w *WebVRFreeCamera) LeftController(leftController *WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(leftController.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetLeftController sets the LeftController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#leftcontroller
func (w *WebVRFreeCamera) SetLeftController(leftController *WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(leftController.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// LockedTarget returns the LockedTarget property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#lockedtarget
func (w *WebVRFreeCamera) LockedTarget(lockedTarget interface{}) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(lockedTarget)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetLockedTarget sets the LockedTarget property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#lockedtarget
func (w *WebVRFreeCamera) SetLockedTarget(lockedTarget interface{}) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(lockedTarget)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// NoRotationConstraint returns the NoRotationConstraint property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#norotationconstraint
func (w *WebVRFreeCamera) NoRotationConstraint(noRotationConstraint bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(noRotationConstraint)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetNoRotationConstraint sets the NoRotationConstraint property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#norotationconstraint
func (w *WebVRFreeCamera) SetNoRotationConstraint(noRotationConstraint bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(noRotationConstraint)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// OnCollide returns the OnCollide property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncollide
func (w *WebVRFreeCamera) OnCollide(onCollide func()) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnCollide sets the OnCollide property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncollide
func (w *WebVRFreeCamera) SetOnCollide(onCollide func()) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onCollide(); return nil}))
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// OnControllerMeshLoadedObservable returns the OnControllerMeshLoadedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollermeshloadedobservable
func (w *WebVRFreeCamera) OnControllerMeshLoadedObservable(onControllerMeshLoadedObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onControllerMeshLoadedObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnControllerMeshLoadedObservable sets the OnControllerMeshLoadedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollermeshloadedobservable
func (w *WebVRFreeCamera) SetOnControllerMeshLoadedObservable(onControllerMeshLoadedObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onControllerMeshLoadedObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// OnControllersAttachedObservable returns the OnControllersAttachedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollersattachedobservable
func (w *WebVRFreeCamera) OnControllersAttachedObservable(onControllersAttachedObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onControllersAttachedObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnControllersAttachedObservable sets the OnControllersAttachedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollersattachedobservable
func (w *WebVRFreeCamera) SetOnControllersAttachedObservable(onControllersAttachedObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onControllersAttachedObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// OnPoseUpdatedFromDeviceObservable returns the OnPoseUpdatedFromDeviceObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#onposeupdatedfromdeviceobservable
func (w *WebVRFreeCamera) OnPoseUpdatedFromDeviceObservable(onPoseUpdatedFromDeviceObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onPoseUpdatedFromDeviceObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetOnPoseUpdatedFromDeviceObservable sets the OnPoseUpdatedFromDeviceObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#onposeupdatedfromdeviceobservable
func (w *WebVRFreeCamera) SetOnPoseUpdatedFromDeviceObservable(onPoseUpdatedFromDeviceObservable *Observable) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(onPoseUpdatedFromDeviceObservable.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// RawPose returns the RawPose property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rawpose
func (w *WebVRFreeCamera) RawPose(rawPose js.Value) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rawPose)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetRawPose sets the RawPose property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rawpose
func (w *WebVRFreeCamera) SetRawPose(rawPose js.Value) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rawPose)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// RigParenting returns the RigParenting property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rigparenting
func (w *WebVRFreeCamera) RigParenting(rigParenting bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rigParenting)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetRigParenting sets the RigParenting property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rigparenting
func (w *WebVRFreeCamera) SetRigParenting(rigParenting bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rigParenting)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// RightController returns the RightController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rightcontroller
func (w *WebVRFreeCamera) RightController(rightController *WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rightController.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetRightController sets the RightController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rightcontroller
func (w *WebVRFreeCamera) SetRightController(rightController *WebVRController) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rightController.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rotation
func (w *WebVRFreeCamera) Rotation(rotation *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rotation.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rotation
func (w *WebVRFreeCamera) SetRotation(rotation *Vector3) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rotation.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rotationquaternion
func (w *WebVRFreeCamera) RotationQuaternion(rotationQuaternion *Quaternion) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rotationQuaternion.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rotationquaternion
func (w *WebVRFreeCamera) SetRotationQuaternion(rotationQuaternion *Quaternion) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(rotationQuaternion.JSObject())
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// Speed returns the Speed property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#speed
func (w *WebVRFreeCamera) Speed(speed float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(speed)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetSpeed sets the Speed property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#speed
func (w *WebVRFreeCamera) SetSpeed(speed float64) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(speed)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// UpdateUpVectorFromRotation returns the UpdateUpVectorFromRotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#updateupvectorfromrotation
func (w *WebVRFreeCamera) UpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(updateUpVectorFromRotation)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

// SetUpdateUpVectorFromRotation sets the UpdateUpVectorFromRotation property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#updateupvectorfromrotation
func (w *WebVRFreeCamera) SetUpdateUpVectorFromRotation(updateUpVectorFromRotation bool) *WebVRFreeCamera {
	p := ba.ctx.Get("WebVRFreeCamera").New(updateUpVectorFromRotation)
	return WebVRFreeCameraFromJSObject(p, ba.ctx)
}

*/
