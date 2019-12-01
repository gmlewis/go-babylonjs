// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OculusTouchController represents a babylon.js OculusTouchController.
// Oculus Touch Controller
type OculusTouchController struct {
	*WebVRController
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *OculusTouchController) JSObject() js.Value { return o.p }

// OculusTouchController returns a OculusTouchController JavaScript class.
func (ba *Babylon) OculusTouchController() *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController")
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OculusTouchControllerFromJSObject returns a wrapped OculusTouchController JavaScript class.
func OculusTouchControllerFromJSObject(p js.Value, ctx js.Value) *OculusTouchController {
	return &OculusTouchController{WebVRController: WebVRControllerFromJSObject(p, ctx), ctx: ctx}
}

// OculusTouchControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func OculusTouchControllerArrayToJSArray(array []*OculusTouchController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewOculusTouchController returns a new OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller
func (ba *Babylon) NewOculusTouchController(vrGamepad interface{}) *OculusTouchController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vrGamepad)

	p := ba.ctx.Get("OculusTouchController").New(args...)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// AttachToMesh calls the AttachToMesh method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#attachtomesh
func (o *OculusTouchController) AttachToMesh(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	o.p.Call("attachToMesh", args...)
}

// AttachToPoseControlledCamera calls the AttachToPoseControlledCamera method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#attachtoposecontrolledcamera
func (o *OculusTouchController) AttachToPoseControlledCamera(camera *TargetCamera) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, camera.JSObject())

	o.p.Call("attachToPoseControlledCamera", args...)
}

// Dispose calls the Dispose method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#dispose
func (o *OculusTouchController) Dispose() {

	o.p.Call("dispose")
}

// OculusTouchControllerGetForwardRayOpts contains optional parameters for OculusTouchController.GetForwardRay.
type OculusTouchControllerGetForwardRayOpts struct {
	Length *float64
}

// GetForwardRay calls the GetForwardRay method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#getforwardray
func (o *OculusTouchController) GetForwardRay(opts *OculusTouchControllerGetForwardRayOpts) *Ray {
	if opts == nil {
		opts = &OculusTouchControllerGetForwardRayOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	retVal := o.p.Call("getForwardRay", args...)
	return RayFromJSObject(retVal, o.ctx)
}

// OculusTouchControllerInitControllerMeshOpts contains optional parameters for OculusTouchController.InitControllerMesh.
type OculusTouchControllerInitControllerMeshOpts struct {
	MeshLoaded *func()
}

// InitControllerMesh calls the InitControllerMesh method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#initcontrollermesh
func (o *OculusTouchController) InitControllerMesh(scene *Scene, opts *OculusTouchControllerInitControllerMeshOpts) {
	if opts == nil {
		opts = &OculusTouchControllerInitControllerMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.MeshLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MeshLoaded)
	}

	o.p.Call("initControllerMesh", args...)
}

// OnButtonStateChange calls the OnButtonStateChange method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onbuttonstatechange
func (o *OculusTouchController) OnButtonStateChange(callback func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	o.p.Call("onButtonStateChange", args...)
}

// Onleftstickchanged calls the Onleftstickchanged method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onleftstickchanged
func (o *OculusTouchController) Onleftstickchanged(callback func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	o.p.Call("onleftstickchanged", args...)
}

// Onrightstickchanged calls the Onrightstickchanged method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onrightstickchanged
func (o *OculusTouchController) Onrightstickchanged(callback func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	o.p.Call("onrightstickchanged", args...)
}

// Update calls the Update method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#update
func (o *OculusTouchController) Update() {

	o.p.Call("update")
}

// UpdateFromDevice calls the UpdateFromDevice method on the OculusTouchController object.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#updatefromdevice
func (o *OculusTouchController) UpdateFromDevice(poseData js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, poseData)

	o.p.Call("updateFromDevice", args...)
}

/*

// BrowserGamepad returns the BrowserGamepad property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#browsergamepad
func (o *OculusTouchController) BrowserGamepad(browserGamepad interface{}) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(browserGamepad)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetBrowserGamepad sets the BrowserGamepad property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#browsergamepad
func (o *OculusTouchController) SetBrowserGamepad(browserGamepad interface{}) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(browserGamepad)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// ControllerType returns the ControllerType property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#controllertype
func (o *OculusTouchController) ControllerType(controllerType *PoseEnabledControllerType) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(controllerType.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetControllerType sets the ControllerType property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#controllertype
func (o *OculusTouchController) SetControllerType(controllerType *PoseEnabledControllerType) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(controllerType.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// DUALSHOCK returns the DUALSHOCK property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#dualshock
func (o *OculusTouchController) DUALSHOCK(DUALSHOCK float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(DUALSHOCK)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetDUALSHOCK sets the DUALSHOCK property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#dualshock
func (o *OculusTouchController) SetDUALSHOCK(DUALSHOCK float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(DUALSHOCK)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// DefaultModel returns the DefaultModel property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#defaultmodel
func (o *OculusTouchController) DefaultModel(defaultModel *AbstractMesh) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(defaultModel.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetDefaultModel sets the DefaultModel property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#defaultmodel
func (o *OculusTouchController) SetDefaultModel(defaultModel *AbstractMesh) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(defaultModel.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// DevicePosition returns the DevicePosition property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#deviceposition
func (o *OculusTouchController) DevicePosition(devicePosition *Vector3) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(devicePosition.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetDevicePosition sets the DevicePosition property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#deviceposition
func (o *OculusTouchController) SetDevicePosition(devicePosition *Vector3) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(devicePosition.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// DeviceRotationQuaternion returns the DeviceRotationQuaternion property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#devicerotationquaternion
func (o *OculusTouchController) DeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(deviceRotationQuaternion.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetDeviceRotationQuaternion sets the DeviceRotationQuaternion property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#devicerotationquaternion
func (o *OculusTouchController) SetDeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(deviceRotationQuaternion.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// DeviceScaleFactor returns the DeviceScaleFactor property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#devicescalefactor
func (o *OculusTouchController) DeviceScaleFactor(deviceScaleFactor float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(deviceScaleFactor)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetDeviceScaleFactor sets the DeviceScaleFactor property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#devicescalefactor
func (o *OculusTouchController) SetDeviceScaleFactor(deviceScaleFactor float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(deviceScaleFactor)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// GAMEPAD returns the GAMEPAD property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#gamepad
func (o *OculusTouchController) GAMEPAD(GAMEPAD float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(GAMEPAD)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetGAMEPAD sets the GAMEPAD property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#gamepad
func (o *OculusTouchController) SetGAMEPAD(GAMEPAD float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(GAMEPAD)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// GENERIC returns the GENERIC property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#generic
func (o *OculusTouchController) GENERIC(GENERIC float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(GENERIC)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetGENERIC sets the GENERIC property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#generic
func (o *OculusTouchController) SetGENERIC(GENERIC float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(GENERIC)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Hand returns the Hand property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#hand
func (o *OculusTouchController) Hand(hand string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(hand)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetHand sets the Hand property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#hand
func (o *OculusTouchController) SetHand(hand string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(hand)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Id returns the Id property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#id
func (o *OculusTouchController) Id(id string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(id)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetId sets the Id property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#id
func (o *OculusTouchController) SetId(id string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(id)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Index returns the Index property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#index
func (o *OculusTouchController) Index(index float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(index)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetIndex sets the Index property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#index
func (o *OculusTouchController) SetIndex(index float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(index)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// IsConnected returns the IsConnected property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#isconnected
func (o *OculusTouchController) IsConnected(isConnected bool) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(isConnected)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetIsConnected sets the IsConnected property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#isconnected
func (o *OculusTouchController) SetIsConnected(isConnected bool) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(isConnected)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// IsXR returns the IsXR property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#isxr
func (o *OculusTouchController) IsXR(isXR bool) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(isXR)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetIsXR sets the IsXR property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#isxr
func (o *OculusTouchController) SetIsXR(isXR bool) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(isXR)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// LeftStick returns the LeftStick property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#leftstick
func (o *OculusTouchController) LeftStick(leftStick *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(leftStick.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetLeftStick sets the LeftStick property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#leftstick
func (o *OculusTouchController) SetLeftStick(leftStick *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(leftStick.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// MODEL_BASE_URL returns the MODEL_BASE_URL property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_base_url
func (o *OculusTouchController) MODEL_BASE_URL(MODEL_BASE_URL string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_BASE_URL)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetMODEL_BASE_URL sets the MODEL_BASE_URL property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_base_url
func (o *OculusTouchController) SetMODEL_BASE_URL(MODEL_BASE_URL string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_BASE_URL)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// MODEL_LEFT_FILENAME returns the MODEL_LEFT_FILENAME property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_left_filename
func (o *OculusTouchController) MODEL_LEFT_FILENAME(MODEL_LEFT_FILENAME string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_LEFT_FILENAME)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetMODEL_LEFT_FILENAME sets the MODEL_LEFT_FILENAME property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_left_filename
func (o *OculusTouchController) SetMODEL_LEFT_FILENAME(MODEL_LEFT_FILENAME string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_LEFT_FILENAME)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// MODEL_RIGHT_FILENAME returns the MODEL_RIGHT_FILENAME property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_right_filename
func (o *OculusTouchController) MODEL_RIGHT_FILENAME(MODEL_RIGHT_FILENAME string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_RIGHT_FILENAME)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetMODEL_RIGHT_FILENAME sets the MODEL_RIGHT_FILENAME property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#model_right_filename
func (o *OculusTouchController) SetMODEL_RIGHT_FILENAME(MODEL_RIGHT_FILENAME string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(MODEL_RIGHT_FILENAME)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Mesh returns the Mesh property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#mesh
func (o *OculusTouchController) Mesh(mesh *AbstractMesh) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(mesh.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetMesh sets the Mesh property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#mesh
func (o *OculusTouchController) SetMesh(mesh *AbstractMesh) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(mesh.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnAButtonStateChangedObservable returns the OnAButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onabuttonstatechangedobservable
func (o *OculusTouchController) OnAButtonStateChangedObservable(onAButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onAButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnAButtonStateChangedObservable sets the OnAButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onabuttonstatechangedobservable
func (o *OculusTouchController) SetOnAButtonStateChangedObservable(onAButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onAButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnBButtonStateChangedObservable returns the OnBButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onbbuttonstatechangedobservable
func (o *OculusTouchController) OnBButtonStateChangedObservable(onBButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onBButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnBButtonStateChangedObservable sets the OnBButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onbbuttonstatechangedobservable
func (o *OculusTouchController) SetOnBButtonStateChangedObservable(onBButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onBButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnMainButtonStateChangedObservable returns the OnMainButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onmainbuttonstatechangedobservable
func (o *OculusTouchController) OnMainButtonStateChangedObservable(onMainButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onMainButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnMainButtonStateChangedObservable sets the OnMainButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onmainbuttonstatechangedobservable
func (o *OculusTouchController) SetOnMainButtonStateChangedObservable(onMainButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onMainButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnPadStateChangedObservable returns the OnPadStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onpadstatechangedobservable
func (o *OculusTouchController) OnPadStateChangedObservable(onPadStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onPadStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnPadStateChangedObservable sets the OnPadStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onpadstatechangedobservable
func (o *OculusTouchController) SetOnPadStateChangedObservable(onPadStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onPadStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnPadValuesChangedObservable returns the OnPadValuesChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onpadvalueschangedobservable
func (o *OculusTouchController) OnPadValuesChangedObservable(onPadValuesChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onPadValuesChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnPadValuesChangedObservable sets the OnPadValuesChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onpadvalueschangedobservable
func (o *OculusTouchController) SetOnPadValuesChangedObservable(onPadValuesChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onPadValuesChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnSecondaryButtonStateChangedObservable returns the OnSecondaryButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onsecondarybuttonstatechangedobservable
func (o *OculusTouchController) OnSecondaryButtonStateChangedObservable(onSecondaryButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onSecondaryButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnSecondaryButtonStateChangedObservable sets the OnSecondaryButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onsecondarybuttonstatechangedobservable
func (o *OculusTouchController) SetOnSecondaryButtonStateChangedObservable(onSecondaryButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onSecondaryButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnSecondaryTriggerStateChangedObservable returns the OnSecondaryTriggerStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onsecondarytriggerstatechangedobservable
func (o *OculusTouchController) OnSecondaryTriggerStateChangedObservable(onSecondaryTriggerStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onSecondaryTriggerStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnSecondaryTriggerStateChangedObservable sets the OnSecondaryTriggerStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onsecondarytriggerstatechangedobservable
func (o *OculusTouchController) SetOnSecondaryTriggerStateChangedObservable(onSecondaryTriggerStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onSecondaryTriggerStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnThumbRestChangedObservable returns the OnThumbRestChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onthumbrestchangedobservable
func (o *OculusTouchController) OnThumbRestChangedObservable(onThumbRestChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onThumbRestChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnThumbRestChangedObservable sets the OnThumbRestChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onthumbrestchangedobservable
func (o *OculusTouchController) SetOnThumbRestChangedObservable(onThumbRestChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onThumbRestChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnTriggerStateChangedObservable returns the OnTriggerStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#ontriggerstatechangedobservable
func (o *OculusTouchController) OnTriggerStateChangedObservable(onTriggerStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onTriggerStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnTriggerStateChangedObservable sets the OnTriggerStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#ontriggerstatechangedobservable
func (o *OculusTouchController) SetOnTriggerStateChangedObservable(onTriggerStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onTriggerStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnXButtonStateChangedObservable returns the OnXButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onxbuttonstatechangedobservable
func (o *OculusTouchController) OnXButtonStateChangedObservable(onXButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onXButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnXButtonStateChangedObservable sets the OnXButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onxbuttonstatechangedobservable
func (o *OculusTouchController) SetOnXButtonStateChangedObservable(onXButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onXButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// OnYButtonStateChangedObservable returns the OnYButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onybuttonstatechangedobservable
func (o *OculusTouchController) OnYButtonStateChangedObservable(onYButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onYButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetOnYButtonStateChangedObservable sets the OnYButtonStateChangedObservable property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#onybuttonstatechangedobservable
func (o *OculusTouchController) SetOnYButtonStateChangedObservable(onYButtonStateChangedObservable *Observable) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(onYButtonStateChangedObservable.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// POINTING_POSE returns the POINTING_POSE property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pointing_pose
func (o *OculusTouchController) POINTING_POSE(POINTING_POSE string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(POINTING_POSE)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetPOINTING_POSE sets the POINTING_POSE property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pointing_pose
func (o *OculusTouchController) SetPOINTING_POSE(POINTING_POSE string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(POINTING_POSE)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// POSE_ENABLED returns the POSE_ENABLED property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pose_enabled
func (o *OculusTouchController) POSE_ENABLED(POSE_ENABLED float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(POSE_ENABLED)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetPOSE_ENABLED sets the POSE_ENABLED property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pose_enabled
func (o *OculusTouchController) SetPOSE_ENABLED(POSE_ENABLED float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(POSE_ENABLED)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Pad returns the Pad property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pad
func (o *OculusTouchController) Pad(pad *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(pad.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetPad sets the Pad property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#pad
func (o *OculusTouchController) SetPad(pad *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(pad.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Position returns the Position property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#position
func (o *OculusTouchController) Position(position *Vector3) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(position.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetPosition sets the Position property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#position
func (o *OculusTouchController) SetPosition(position *Vector3) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(position.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// QUEST_MODEL_BASE_URL returns the QUEST_MODEL_BASE_URL property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#quest_model_base_url
func (o *OculusTouchController) QUEST_MODEL_BASE_URL(QUEST_MODEL_BASE_URL string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(QUEST_MODEL_BASE_URL)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetQUEST_MODEL_BASE_URL sets the QUEST_MODEL_BASE_URL property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#quest_model_base_url
func (o *OculusTouchController) SetQUEST_MODEL_BASE_URL(QUEST_MODEL_BASE_URL string) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(QUEST_MODEL_BASE_URL)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// RawPose returns the RawPose property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rawpose
func (o *OculusTouchController) RawPose(rawPose js.Value) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rawPose)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetRawPose sets the RawPose property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rawpose
func (o *OculusTouchController) SetRawPose(rawPose js.Value) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rawPose)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// RightStick returns the RightStick property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rightstick
func (o *OculusTouchController) RightStick(rightStick *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rightStick.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetRightStick sets the RightStick property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rightstick
func (o *OculusTouchController) SetRightStick(rightStick *StickValues) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rightStick.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rotationquaternion
func (o *OculusTouchController) RotationQuaternion(rotationQuaternion *Quaternion) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rotationQuaternion.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#rotationquaternion
func (o *OculusTouchController) SetRotationQuaternion(rotationQuaternion *Quaternion) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(rotationQuaternion.JSObject())
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#type
func (o *OculusTouchController) Type(jsType float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(jsType)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#type
func (o *OculusTouchController) SetType(jsType float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(jsType)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// XBOX returns the XBOX property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#xbox
func (o *OculusTouchController) XBOX(XBOX float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(XBOX)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

// SetXBOX sets the XBOX property of class OculusTouchController.
//
// https://doc.babylonjs.com/api/classes/babylon.oculustouchcontroller#xbox
func (o *OculusTouchController) SetXBOX(XBOX float64) *OculusTouchController {
	p := ba.ctx.Get("OculusTouchController").New(XBOX)
	return OculusTouchControllerFromJSObject(p, ba.ctx)
}

*/
