// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebVRFreeCamera represents a babylon.js WebVRFreeCamera.
// This represents a WebVR camera.
// The WebVR camera is Babylon&#39;s simple interface to interaction with Windows Mixed Reality, HTC Vive and Oculus Rift.
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

	args = append(args, opts.WebVROptions)

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
	Callback func()
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
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.Callback(); return nil }) /* never freed! */)
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

// Controllers returns the Controllers property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#controllers
func (w *WebVRFreeCamera) Controllers() []*WebVRController {
	retVal := w.p.Get("controllers")
	result := []*WebVRController{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, WebVRControllerFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// SetControllers sets the Controllers property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#controllers
func (w *WebVRFreeCamera) SetControllers(controllers []*WebVRController) *WebVRFreeCamera {
	w.p.Set("controllers", controllers)
	return w
}

// DevicePosition returns the DevicePosition property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#deviceposition
func (w *WebVRFreeCamera) DevicePosition() *Vector3 {
	retVal := w.p.Get("devicePosition")
	return Vector3FromJSObject(retVal, w.ctx)
}

// SetDevicePosition sets the DevicePosition property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#deviceposition
func (w *WebVRFreeCamera) SetDevicePosition(devicePosition *Vector3) *WebVRFreeCamera {
	w.p.Set("devicePosition", devicePosition.JSObject())
	return w
}

// DeviceRotationQuaternion returns the DeviceRotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicerotationquaternion
func (w *WebVRFreeCamera) DeviceRotationQuaternion() *Quaternion {
	retVal := w.p.Get("deviceRotationQuaternion")
	return QuaternionFromJSObject(retVal, w.ctx)
}

// SetDeviceRotationQuaternion sets the DeviceRotationQuaternion property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicerotationquaternion
func (w *WebVRFreeCamera) SetDeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *WebVRFreeCamera {
	w.p.Set("deviceRotationQuaternion", deviceRotationQuaternion.JSObject())
	return w
}

// DeviceScaleFactor returns the DeviceScaleFactor property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicescalefactor
func (w *WebVRFreeCamera) DeviceScaleFactor() float64 {
	retVal := w.p.Get("deviceScaleFactor")
	return retVal.Float()
}

// SetDeviceScaleFactor sets the DeviceScaleFactor property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#devicescalefactor
func (w *WebVRFreeCamera) SetDeviceScaleFactor(deviceScaleFactor float64) *WebVRFreeCamera {
	w.p.Set("deviceScaleFactor", deviceScaleFactor)
	return w
}

// LeftController returns the LeftController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#leftcontroller
func (w *WebVRFreeCamera) LeftController() *WebVRController {
	retVal := w.p.Get("leftController")
	return WebVRControllerFromJSObject(retVal, w.ctx)
}

// SetLeftController sets the LeftController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#leftcontroller
func (w *WebVRFreeCamera) SetLeftController(leftController *WebVRController) *WebVRFreeCamera {
	w.p.Set("leftController", leftController.JSObject())
	return w
}

// OnControllerMeshLoadedObservable returns the OnControllerMeshLoadedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollermeshloadedobservable
func (w *WebVRFreeCamera) OnControllerMeshLoadedObservable() *Observable {
	retVal := w.p.Get("onControllerMeshLoadedObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetOnControllerMeshLoadedObservable sets the OnControllerMeshLoadedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollermeshloadedobservable
func (w *WebVRFreeCamera) SetOnControllerMeshLoadedObservable(onControllerMeshLoadedObservable *Observable) *WebVRFreeCamera {
	w.p.Set("onControllerMeshLoadedObservable", onControllerMeshLoadedObservable.JSObject())
	return w
}

// OnControllersAttachedObservable returns the OnControllersAttachedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollersattachedobservable
func (w *WebVRFreeCamera) OnControllersAttachedObservable() []*Observable {
	retVal := w.p.Get("onControllersAttachedObservable")
	result := []*Observable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ObservableFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// SetOnControllersAttachedObservable sets the OnControllersAttachedObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#oncontrollersattachedobservable
func (w *WebVRFreeCamera) SetOnControllersAttachedObservable(onControllersAttachedObservable []*Observable) *WebVRFreeCamera {
	w.p.Set("onControllersAttachedObservable", onControllersAttachedObservable)
	return w
}

// OnPoseUpdatedFromDeviceObservable returns the OnPoseUpdatedFromDeviceObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#onposeupdatedfromdeviceobservable
func (w *WebVRFreeCamera) OnPoseUpdatedFromDeviceObservable() *Observable {
	retVal := w.p.Get("onPoseUpdatedFromDeviceObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetOnPoseUpdatedFromDeviceObservable sets the OnPoseUpdatedFromDeviceObservable property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#onposeupdatedfromdeviceobservable
func (w *WebVRFreeCamera) SetOnPoseUpdatedFromDeviceObservable(onPoseUpdatedFromDeviceObservable *Observable) *WebVRFreeCamera {
	w.p.Set("onPoseUpdatedFromDeviceObservable", onPoseUpdatedFromDeviceObservable.JSObject())
	return w
}

// RawPose returns the RawPose property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rawpose
func (w *WebVRFreeCamera) RawPose() js.Value {
	retVal := w.p.Get("rawPose")
	return retVal
}

// SetRawPose sets the RawPose property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rawpose
func (w *WebVRFreeCamera) SetRawPose(rawPose js.Value) *WebVRFreeCamera {
	w.p.Set("rawPose", rawPose)
	return w
}

// RigParenting returns the RigParenting property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rigparenting
func (w *WebVRFreeCamera) RigParenting() bool {
	retVal := w.p.Get("rigParenting")
	return retVal.Bool()
}

// SetRigParenting sets the RigParenting property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rigparenting
func (w *WebVRFreeCamera) SetRigParenting(rigParenting bool) *WebVRFreeCamera {
	w.p.Set("rigParenting", rigParenting)
	return w
}

// RightController returns the RightController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rightcontroller
func (w *WebVRFreeCamera) RightController() *WebVRController {
	retVal := w.p.Get("rightController")
	return WebVRControllerFromJSObject(retVal, w.ctx)
}

// SetRightController sets the RightController property of class WebVRFreeCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.webvrfreecamera#rightcontroller
func (w *WebVRFreeCamera) SetRightController(rightController *WebVRController) *WebVRFreeCamera {
	w.p.Set("rightController", rightController.JSObject())
	return w
}
