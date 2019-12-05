// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ViveController represents a babylon.js ViveController.
// Vive Controller
type ViveController struct {
	*WebVRController
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *ViveController) JSObject() js.Value { return v.p }

// ViveController returns a ViveController JavaScript class.
func (ba *Babylon) ViveController() *ViveController {
	p := ba.ctx.Get("ViveController")
	return ViveControllerFromJSObject(p, ba.ctx)
}

// ViveControllerFromJSObject returns a wrapped ViveController JavaScript class.
func ViveControllerFromJSObject(p js.Value, ctx js.Value) *ViveController {
	return &ViveController{WebVRController: WebVRControllerFromJSObject(p, ctx), ctx: ctx}
}

// ViveControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ViveControllerArrayToJSArray(array []*ViveController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewViveController returns a new ViveController object.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller
func (ba *Babylon) NewViveController(vrGamepad interface{}) *ViveController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vrGamepad)

	p := ba.ctx.Get("ViveController").New(args...)
	return ViveControllerFromJSObject(p, ba.ctx)
}

// ViveControllerInitControllerMeshOpts contains optional parameters for ViveController.InitControllerMesh.
type ViveControllerInitControllerMeshOpts struct {
	MeshLoaded func()
}

// InitControllerMesh calls the InitControllerMesh method on the ViveController object.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#initcontrollermesh
func (v *ViveController) InitControllerMesh(scene *Scene, opts *ViveControllerInitControllerMeshOpts) {
	if opts == nil {
		opts = &ViveControllerInitControllerMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.MeshLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.MeshLoaded(); return nil }) /* never freed! */)
	}

	v.p.Call("initControllerMesh", args...)
}

// MODEL_BASE_URL returns the MODEL_BASE_URL property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#model_base_url
func (v *ViveController) MODEL_BASE_URL() string {
	retVal := v.p.Get("MODEL_BASE_URL")
	return retVal.String()
}

// SetMODEL_BASE_URL sets the MODEL_BASE_URL property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#model_base_url
func (v *ViveController) SetMODEL_BASE_URL(MODEL_BASE_URL string) *ViveController {
	v.p.Set("MODEL_BASE_URL", MODEL_BASE_URL)
	return v
}

// MODEL_FILENAME returns the MODEL_FILENAME property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#model_filename
func (v *ViveController) MODEL_FILENAME() string {
	retVal := v.p.Get("MODEL_FILENAME")
	return retVal.String()
}

// SetMODEL_FILENAME sets the MODEL_FILENAME property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#model_filename
func (v *ViveController) SetMODEL_FILENAME(MODEL_FILENAME string) *ViveController {
	v.p.Set("MODEL_FILENAME", MODEL_FILENAME)
	return v
}

// OnLeftButtonStateChangedObservable returns the OnLeftButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onleftbuttonstatechangedobservable
func (v *ViveController) OnLeftButtonStateChangedObservable() *Observable {
	retVal := v.p.Get("onLeftButtonStateChangedObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnLeftButtonStateChangedObservable sets the OnLeftButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onleftbuttonstatechangedobservable
func (v *ViveController) SetOnLeftButtonStateChangedObservable(onLeftButtonStateChangedObservable *Observable) *ViveController {
	v.p.Set("onLeftButtonStateChangedObservable", onLeftButtonStateChangedObservable.JSObject())
	return v
}

// OnMenuButtonStateChangedObservable returns the OnMenuButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onmenubuttonstatechangedobservable
func (v *ViveController) OnMenuButtonStateChangedObservable() *Observable {
	retVal := v.p.Get("onMenuButtonStateChangedObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnMenuButtonStateChangedObservable sets the OnMenuButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onmenubuttonstatechangedobservable
func (v *ViveController) SetOnMenuButtonStateChangedObservable(onMenuButtonStateChangedObservable *Observable) *ViveController {
	v.p.Set("onMenuButtonStateChangedObservable", onMenuButtonStateChangedObservable.JSObject())
	return v
}

// OnRightButtonStateChangedObservable returns the OnRightButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onrightbuttonstatechangedobservable
func (v *ViveController) OnRightButtonStateChangedObservable() *Observable {
	retVal := v.p.Get("onRightButtonStateChangedObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnRightButtonStateChangedObservable sets the OnRightButtonStateChangedObservable property of class ViveController.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller#onrightbuttonstatechangedobservable
func (v *ViveController) SetOnRightButtonStateChangedObservable(onRightButtonStateChangedObservable *Observable) *ViveController {
	v.p.Set("onRightButtonStateChangedObservable", onRightButtonStateChangedObservable.JSObject())
	return v
}
