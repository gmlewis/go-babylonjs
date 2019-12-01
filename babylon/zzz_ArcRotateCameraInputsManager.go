// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCameraInputsManager represents a babylon.js ArcRotateCameraInputsManager.
// Default Inputs manager for the ArcRotateCamera.
// It groups all the default supported inputs for ease of use.
// Interface representing an arc rotate camera inputs manager
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type ArcRotateCameraInputsManager struct {
	*CameraInputsManager
	*ArcRotateCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCameraInputsManager) JSObject() js.Value { return a.p }

// ArcRotateCameraInputsManager returns a ArcRotateCameraInputsManager JavaScript class.
func (ba *Babylon) ArcRotateCameraInputsManager() *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager")
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// ArcRotateCameraInputsManagerFromJSObject returns a wrapped ArcRotateCameraInputsManager JavaScript class.
func ArcRotateCameraInputsManagerFromJSObject(p js.Value, ctx js.Value) *ArcRotateCameraInputsManager {
	return &ArcRotateCameraInputsManager{CameraInputsManager: CameraInputsManagerFromJSObject(p, ctx), ArcRotateCamera: ArcRotateCameraFromJSObject(p, ctx), ctx: ctx}
}

// ArcRotateCameraInputsManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArcRotateCameraInputsManagerArrayToJSArray(array []*ArcRotateCameraInputsManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewArcRotateCameraInputsManager returns a new ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager
func (ba *Babylon) NewArcRotateCameraInputsManager(camera *ArcRotateCamera) *ArcRotateCameraInputsManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, camera.JSObject())

	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(args...)
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// Add calls the Add method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#add
func (a *ArcRotateCameraInputsManager) Add(input js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, input)

	a.p.Call("add", args...)
}

// AddGamepad calls the AddGamepad method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#addgamepad
func (a *ArcRotateCameraInputsManager) AddGamepad() *ArcRotateCameraInputsManager {

	retVal := a.p.Call("addGamepad")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// AddKeyboard calls the AddKeyboard method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#addkeyboard
func (a *ArcRotateCameraInputsManager) AddKeyboard() *ArcRotateCameraInputsManager {

	retVal := a.p.Call("addKeyboard")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// AddMouseWheel calls the AddMouseWheel method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#addmousewheel
func (a *ArcRotateCameraInputsManager) AddMouseWheel() *ArcRotateCameraInputsManager {

	retVal := a.p.Call("addMouseWheel")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// AddPointers calls the AddPointers method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#addpointers
func (a *ArcRotateCameraInputsManager) AddPointers() *ArcRotateCameraInputsManager {

	retVal := a.p.Call("addPointers")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// AddVRDeviceOrientation calls the AddVRDeviceOrientation method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#addvrdeviceorientation
func (a *ArcRotateCameraInputsManager) AddVRDeviceOrientation() *ArcRotateCameraInputsManager {

	retVal := a.p.Call("addVRDeviceOrientation")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// ArcRotateCameraInputsManagerAttachElementOpts contains optional parameters for ArcRotateCameraInputsManager.AttachElement.
type ArcRotateCameraInputsManagerAttachElementOpts struct {
	NoPreventDefault *bool
}

// AttachElement calls the AttachElement method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attachelement
func (a *ArcRotateCameraInputsManager) AttachElement(element js.Value, opts *ArcRotateCameraInputsManagerAttachElementOpts) {
	if opts == nil {
		opts = &ArcRotateCameraInputsManagerAttachElementOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	a.p.Call("attachElement", args...)
}

// AttachInput calls the AttachInput method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attachinput
func (a *ArcRotateCameraInputsManager) AttachInput(input js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, input)

	a.p.Call("attachInput", args...)
}

// Clear calls the Clear method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#clear
func (a *ArcRotateCameraInputsManager) Clear() {

	a.p.Call("clear")
}

// ArcRotateCameraInputsManagerDetachElementOpts contains optional parameters for ArcRotateCameraInputsManager.DetachElement.
type ArcRotateCameraInputsManagerDetachElementOpts struct {
	Disconnect *bool
}

// DetachElement calls the DetachElement method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#detachelement
func (a *ArcRotateCameraInputsManager) DetachElement(element js.Value, opts *ArcRotateCameraInputsManagerDetachElementOpts) {
	if opts == nil {
		opts = &ArcRotateCameraInputsManagerDetachElementOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.Disconnect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Disconnect)
	}

	a.p.Call("detachElement", args...)
}

// Parse calls the Parse method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#parse
func (a *ArcRotateCameraInputsManager) Parse(parsedCamera interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, parsedCamera)

	a.p.Call("parse", args...)
}

// RebuildInputCheck calls the RebuildInputCheck method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#rebuildinputcheck
func (a *ArcRotateCameraInputsManager) RebuildInputCheck() {

	a.p.Call("rebuildInputCheck")
}

// Remove calls the Remove method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#remove
func (a *ArcRotateCameraInputsManager) Remove(inputToRemove js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, inputToRemove)

	a.p.Call("remove", args...)
}

// RemoveByType calls the RemoveByType method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#removebytype
func (a *ArcRotateCameraInputsManager) RemoveByType(inputType string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, inputType)

	a.p.Call("removeByType", args...)
}

// Serialize calls the Serialize method on the ArcRotateCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#serialize
func (a *ArcRotateCameraInputsManager) Serialize(serializedCamera interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, serializedCamera)

	a.p.Call("serialize", args...)
}

/*

// Attached returns the Attached property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attached
func (a *ArcRotateCameraInputsManager) Attached(attached *CameraInputsMap) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(attached.JSObject())
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// SetAttached sets the Attached property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attached
func (a *ArcRotateCameraInputsManager) SetAttached(attached *CameraInputsMap) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(attached.JSObject())
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// AttachedElement returns the AttachedElement property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attachedelement
func (a *ArcRotateCameraInputsManager) AttachedElement(attachedElement js.Value) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(attachedElement)
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// SetAttachedElement sets the AttachedElement property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#attachedelement
func (a *ArcRotateCameraInputsManager) SetAttachedElement(attachedElement js.Value) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(attachedElement)
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#camera
func (a *ArcRotateCameraInputsManager) Camera(camera *ArcRotateCamera) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(camera.JSObject())
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#camera
func (a *ArcRotateCameraInputsManager) SetCamera(camera *ArcRotateCamera) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(camera.JSObject())
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// CheckInputs returns the CheckInputs property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#checkinputs
func (a *ArcRotateCameraInputsManager) CheckInputs(checkInputs func()) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {checkInputs(); return nil}))
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// SetCheckInputs sets the CheckInputs property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#checkinputs
func (a *ArcRotateCameraInputsManager) SetCheckInputs(checkInputs func()) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {checkInputs(); return nil}))
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// NoPreventDefault returns the NoPreventDefault property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#nopreventdefault
func (a *ArcRotateCameraInputsManager) NoPreventDefault(noPreventDefault bool) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(noPreventDefault)
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

// SetNoPreventDefault sets the NoPreventDefault property of class ArcRotateCameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerainputsmanager#nopreventdefault
func (a *ArcRotateCameraInputsManager) SetNoPreventDefault(noPreventDefault bool) *ArcRotateCameraInputsManager {
	p := ba.ctx.Get("ArcRotateCameraInputsManager").New(noPreventDefault)
	return ArcRotateCameraInputsManagerFromJSObject(p, ba.ctx)
}

*/
