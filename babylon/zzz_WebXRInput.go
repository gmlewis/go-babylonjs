// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRInput represents a babylon.js WebXRInput.
// XR input used to track XR inputs such as controllers/rays
type WebXRInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRInput) JSObject() js.Value { return w.p }

// WebXRInput returns a WebXRInput JavaScript class.
func (ba *Babylon) WebXRInput() *WebXRInput {
	p := ba.ctx.Get("WebXRInput")
	return WebXRInputFromJSObject(p, ba.ctx)
}

// WebXRInputFromJSObject returns a wrapped WebXRInput JavaScript class.
func WebXRInputFromJSObject(p js.Value, ctx js.Value) *WebXRInput {
	return &WebXRInput{p: p, ctx: ctx}
}

// WebXRInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebXRInputArrayToJSArray(array []*WebXRInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWebXRInput returns a new WebXRInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#constructor
func (ba *Babylon) NewWebXRInput(baseExperience *WebXRExperienceHelper) *WebXRInput {

	args := make([]interface{}, 0, 1+0)

	args = append(args, baseExperience.JSObject())

	p := ba.ctx.Get("WebXRInput").New(args...)
	return WebXRInputFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the WebXRInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#dispose
func (w *WebXRInput) Dispose() {

	w.p.Call("dispose")
}

// BaseExperience returns the BaseExperience property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#baseexperience
func (w *WebXRInput) BaseExperience() *WebXRExperienceHelper {
	retVal := w.p.Get("baseExperience")
	return WebXRExperienceHelperFromJSObject(retVal, w.ctx)
}

// SetBaseExperience sets the BaseExperience property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#baseexperience
func (w *WebXRInput) SetBaseExperience(baseExperience *WebXRExperienceHelper) *WebXRInput {
	w.p.Set("baseExperience", baseExperience.JSObject())
	return w
}

// Controllers returns the Controllers property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#controllers
func (w *WebXRInput) Controllers() []*WebXRController {
	retVal := w.p.Get("controllers")
	result := []*WebXRController{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, WebXRControllerFromJSObject(retVal.Index(ri), w.ctx))
	}
	return result
}

// SetControllers sets the Controllers property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#controllers
func (w *WebXRInput) SetControllers(controllers []*WebXRController) *WebXRInput {
	w.p.Set("controllers", controllers)
	return w
}

// OnControllerAddedObservable returns the OnControllerAddedObservable property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#oncontrolleraddedobservable
func (w *WebXRInput) OnControllerAddedObservable() *Observable {
	retVal := w.p.Get("onControllerAddedObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetOnControllerAddedObservable sets the OnControllerAddedObservable property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#oncontrolleraddedobservable
func (w *WebXRInput) SetOnControllerAddedObservable(onControllerAddedObservable *Observable) *WebXRInput {
	w.p.Set("onControllerAddedObservable", onControllerAddedObservable.JSObject())
	return w
}

// OnControllerRemovedObservable returns the OnControllerRemovedObservable property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#oncontrollerremovedobservable
func (w *WebXRInput) OnControllerRemovedObservable() *Observable {
	retVal := w.p.Get("onControllerRemovedObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetOnControllerRemovedObservable sets the OnControllerRemovedObservable property of class WebXRInput.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrinput#oncontrollerremovedobservable
func (w *WebXRInput) SetOnControllerRemovedObservable(onControllerRemovedObservable *Observable) *WebXRInput {
	w.p.Set("onControllerRemovedObservable", onControllerRemovedObservable.JSObject())
	return w
}
