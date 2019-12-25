// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CameraInputsManager represents a babylon.js CameraInputsManager.
// This represents the input manager used within a camera.
// It helps dealing with all the different kind of input attached to a camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type CameraInputsManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CameraInputsManager) JSObject() js.Value { return c.p }

// CameraInputsManager returns a CameraInputsManager JavaScript class.
func (ba *Babylon) CameraInputsManager() *CameraInputsManager {
	p := ba.ctx.Get("CameraInputsManager")
	return CameraInputsManagerFromJSObject(p, ba.ctx)
}

// CameraInputsManagerFromJSObject returns a wrapped CameraInputsManager JavaScript class.
func CameraInputsManagerFromJSObject(p js.Value, ctx js.Value) *CameraInputsManager {
	return &CameraInputsManager{p: p, ctx: ctx}
}

// CameraInputsManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func CameraInputsManagerArrayToJSArray(array []*CameraInputsManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCameraInputsManager returns a new CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#constructor
func (ba *Babylon) NewCameraInputsManager(camera *Camera) *CameraInputsManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, camera.JSObject())

	p := ba.ctx.Get("CameraInputsManager").New(args...)
	return CameraInputsManagerFromJSObject(p, ba.ctx)
}

// Add calls the Add method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#add
func (c *CameraInputsManager) Add(input *ICameraInput) {

	args := make([]interface{}, 0, 1+0)

	if input == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, input.JSObject())
	}

	c.p.Call("add", args...)
}

// CameraInputsManagerAttachElementOpts contains optional parameters for CameraInputsManager.AttachElement.
type CameraInputsManagerAttachElementOpts struct {
	NoPreventDefault *bool
}

// AttachElement calls the AttachElement method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attachelement
func (c *CameraInputsManager) AttachElement(element js.Value, opts *CameraInputsManagerAttachElementOpts) {
	if opts == nil {
		opts = &CameraInputsManagerAttachElementOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	c.p.Call("attachElement", args...)
}

// AttachInput calls the AttachInput method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attachinput
func (c *CameraInputsManager) AttachInput(input *ICameraInput) {

	args := make([]interface{}, 0, 1+0)

	if input == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, input.JSObject())
	}

	c.p.Call("attachInput", args...)
}

// Clear calls the Clear method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#clear
func (c *CameraInputsManager) Clear() {

	c.p.Call("clear")
}

// CameraInputsManagerDetachElementOpts contains optional parameters for CameraInputsManager.DetachElement.
type CameraInputsManagerDetachElementOpts struct {
	Disconnect *bool
}

// DetachElement calls the DetachElement method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#detachelement
func (c *CameraInputsManager) DetachElement(element js.Value, opts *CameraInputsManagerDetachElementOpts) {
	if opts == nil {
		opts = &CameraInputsManagerDetachElementOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.Disconnect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Disconnect)
	}

	c.p.Call("detachElement", args...)
}

// Parse calls the Parse method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#parse
func (c *CameraInputsManager) Parse(parsedCamera JSObject) {

	args := make([]interface{}, 0, 1+0)

	if parsedCamera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedCamera.JSObject())
	}

	c.p.Call("parse", args...)
}

// RebuildInputCheck calls the RebuildInputCheck method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#rebuildinputcheck
func (c *CameraInputsManager) RebuildInputCheck() {

	c.p.Call("rebuildInputCheck")
}

// Remove calls the Remove method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#remove
func (c *CameraInputsManager) Remove(inputToRemove *ICameraInput) {

	args := make([]interface{}, 0, 1+0)

	if inputToRemove == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, inputToRemove.JSObject())
	}

	c.p.Call("remove", args...)
}

// RemoveByType calls the RemoveByType method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#removebytype
func (c *CameraInputsManager) RemoveByType(inputType string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, inputType)

	c.p.Call("removeByType", args...)
}

// Serialize calls the Serialize method on the CameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#serialize
func (c *CameraInputsManager) Serialize(serializedCamera JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializedCamera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializedCamera.JSObject())
	}

	c.p.Call("serialize", args...)
}

// Attached returns the Attached property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attached
func (c *CameraInputsManager) Attached() *CameraInputsMap {
	retVal := c.p.Get("attached")
	return CameraInputsMapFromJSObject(retVal, c.ctx)
}

// SetAttached sets the Attached property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attached
func (c *CameraInputsManager) SetAttached(attached *CameraInputsMap) *CameraInputsManager {
	c.p.Set("attached", attached.JSObject())
	return c
}

// AttachedElement returns the AttachedElement property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attachedelement
func (c *CameraInputsManager) AttachedElement() js.Value {
	retVal := c.p.Get("attachedElement")
	return retVal
}

// SetAttachedElement sets the AttachedElement property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#attachedelement
func (c *CameraInputsManager) SetAttachedElement(attachedElement js.Value) *CameraInputsManager {
	c.p.Set("attachedElement", attachedElement)
	return c
}

// Camera returns the Camera property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#camera
func (c *CameraInputsManager) Camera() *Camera {
	retVal := c.p.Get("camera")
	return CameraFromJSObject(retVal, c.ctx)
}

// SetCamera sets the Camera property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#camera
func (c *CameraInputsManager) SetCamera(camera *Camera) *CameraInputsManager {
	c.p.Set("camera", camera.JSObject())
	return c
}

// CheckInputs returns the CheckInputs property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#checkinputs
func (c *CameraInputsManager) CheckInputs() js.Value {
	retVal := c.p.Get("checkInputs")
	return retVal
}

// SetCheckInputs sets the CheckInputs property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#checkinputs
func (c *CameraInputsManager) SetCheckInputs(checkInputs JSFunc) *CameraInputsManager {
	c.p.Set("checkInputs", js.FuncOf(checkInputs))
	return c
}

// NoPreventDefault returns the NoPreventDefault property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#nopreventdefault
func (c *CameraInputsManager) NoPreventDefault() bool {
	retVal := c.p.Get("noPreventDefault")
	return retVal.Bool()
}

// SetNoPreventDefault sets the NoPreventDefault property of class CameraInputsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.camerainputsmanager#nopreventdefault
func (c *CameraInputsManager) SetNoPreventDefault(noPreventDefault bool) *CameraInputsManager {
	c.p.Set("noPreventDefault", noPreventDefault)
	return c
}
