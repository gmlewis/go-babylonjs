// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DaydreamController represents a babylon.js DaydreamController.
// Google Daydream controller
type DaydreamController struct {
	*WebVRController
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DaydreamController) JSObject() js.Value { return d.p }

// DaydreamController returns a DaydreamController JavaScript class.
func (ba *Babylon) DaydreamController() *DaydreamController {
	p := ba.ctx.Get("DaydreamController")
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// DaydreamControllerFromJSObject returns a wrapped DaydreamController JavaScript class.
func DaydreamControllerFromJSObject(p js.Value, ctx js.Value) *DaydreamController {
	return &DaydreamController{WebVRController: WebVRControllerFromJSObject(p, ctx), ctx: ctx}
}

// DaydreamControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func DaydreamControllerArrayToJSArray(array []*DaydreamController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDaydreamController returns a new DaydreamController object.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller
func (ba *Babylon) NewDaydreamController(vrGamepad interface{}) *DaydreamController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vrGamepad)

	p := ba.ctx.Get("DaydreamController").New(args...)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// DaydreamControllerInitControllerMeshOpts contains optional parameters for DaydreamController.InitControllerMesh.
type DaydreamControllerInitControllerMeshOpts struct {
	MeshLoaded func()
}

// InitControllerMesh calls the InitControllerMesh method on the DaydreamController object.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#initcontrollermesh
func (d *DaydreamController) InitControllerMesh(scene *Scene, opts *DaydreamControllerInitControllerMeshOpts) {
	if opts == nil {
		opts = &DaydreamControllerInitControllerMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.MeshLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MeshLoaded)
	}

	d.p.Call("initControllerMesh", args...)
}

/*

// GAMEPAD_ID_PREFIX returns the GAMEPAD_ID_PREFIX property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#gamepad_id_prefix
func (d *DaydreamController) GAMEPAD_ID_PREFIX(GAMEPAD_ID_PREFIX string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(GAMEPAD_ID_PREFIX)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// SetGAMEPAD_ID_PREFIX sets the GAMEPAD_ID_PREFIX property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#gamepad_id_prefix
func (d *DaydreamController) SetGAMEPAD_ID_PREFIX(GAMEPAD_ID_PREFIX string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(GAMEPAD_ID_PREFIX)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// MODEL_BASE_URL returns the MODEL_BASE_URL property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#model_base_url
func (d *DaydreamController) MODEL_BASE_URL(MODEL_BASE_URL string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(MODEL_BASE_URL)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// SetMODEL_BASE_URL sets the MODEL_BASE_URL property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#model_base_url
func (d *DaydreamController) SetMODEL_BASE_URL(MODEL_BASE_URL string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(MODEL_BASE_URL)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// MODEL_FILENAME returns the MODEL_FILENAME property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#model_filename
func (d *DaydreamController) MODEL_FILENAME(MODEL_FILENAME string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(MODEL_FILENAME)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

// SetMODEL_FILENAME sets the MODEL_FILENAME property of class DaydreamController.
//
// https://doc.babylonjs.com/api/classes/babylon.daydreamcontroller#model_filename
func (d *DaydreamController) SetMODEL_FILENAME(MODEL_FILENAME string) *DaydreamController {
	p := ba.ctx.Get("DaydreamController").New(MODEL_FILENAME)
	return DaydreamControllerFromJSObject(p, ba.ctx)
}

*/
