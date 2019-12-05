// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GearVRController represents a babylon.js GearVRController.
// Gear VR Controller
type GearVRController struct {
	*WebVRController
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GearVRController) JSObject() js.Value { return g.p }

// GearVRController returns a GearVRController JavaScript class.
func (ba *Babylon) GearVRController() *GearVRController {
	p := ba.ctx.Get("GearVRController")
	return GearVRControllerFromJSObject(p, ba.ctx)
}

// GearVRControllerFromJSObject returns a wrapped GearVRController JavaScript class.
func GearVRControllerFromJSObject(p js.Value, ctx js.Value) *GearVRController {
	return &GearVRController{WebVRController: WebVRControllerFromJSObject(p, ctx), ctx: ctx}
}

// GearVRControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GearVRControllerArrayToJSArray(array []*GearVRController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGearVRController returns a new GearVRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller
func (ba *Babylon) NewGearVRController(vrGamepad interface{}) *GearVRController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vrGamepad)

	p := ba.ctx.Get("GearVRController").New(args...)
	return GearVRControllerFromJSObject(p, ba.ctx)
}

// GearVRControllerInitControllerMeshOpts contains optional parameters for GearVRController.InitControllerMesh.
type GearVRControllerInitControllerMeshOpts struct {
	MeshLoaded func()
}

// InitControllerMesh calls the InitControllerMesh method on the GearVRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#initcontrollermesh
func (g *GearVRController) InitControllerMesh(scene *Scene, opts *GearVRControllerInitControllerMeshOpts) {
	if opts == nil {
		opts = &GearVRControllerInitControllerMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.MeshLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.MeshLoaded(); return nil }) /* never freed! */)
	}

	g.p.Call("initControllerMesh", args...)
}

// GAMEPAD_ID_PREFIX returns the GAMEPAD_ID_PREFIX property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#gamepad_id_prefix
func (g *GearVRController) GAMEPAD_ID_PREFIX() string {
	retVal := g.p.Get("GAMEPAD_ID_PREFIX")
	return retVal.String()
}

// SetGAMEPAD_ID_PREFIX sets the GAMEPAD_ID_PREFIX property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#gamepad_id_prefix
func (g *GearVRController) SetGAMEPAD_ID_PREFIX(GAMEPAD_ID_PREFIX string) *GearVRController {
	g.p.Set("GAMEPAD_ID_PREFIX", GAMEPAD_ID_PREFIX)
	return g
}

// MODEL_BASE_URL returns the MODEL_BASE_URL property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#model_base_url
func (g *GearVRController) MODEL_BASE_URL() string {
	retVal := g.p.Get("MODEL_BASE_URL")
	return retVal.String()
}

// SetMODEL_BASE_URL sets the MODEL_BASE_URL property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#model_base_url
func (g *GearVRController) SetMODEL_BASE_URL(MODEL_BASE_URL string) *GearVRController {
	g.p.Set("MODEL_BASE_URL", MODEL_BASE_URL)
	return g
}

// MODEL_FILENAME returns the MODEL_FILENAME property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#model_filename
func (g *GearVRController) MODEL_FILENAME() string {
	retVal := g.p.Get("MODEL_FILENAME")
	return retVal.String()
}

// SetMODEL_FILENAME sets the MODEL_FILENAME property of class GearVRController.
//
// https://doc.babylonjs.com/api/classes/babylon.gearvrcontroller#model_filename
func (g *GearVRController) SetMODEL_FILENAME(MODEL_FILENAME string) *GearVRController {
	g.p.Set("MODEL_FILENAME", MODEL_FILENAME)
	return g
}
