// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FreeCameraTouchInput represents a babylon.js FreeCameraTouchInput.
// Manage the touch inputs to control the movement of a free camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FreeCameraTouchInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FreeCameraTouchInput) JSObject() js.Value { return f.p }

// FreeCameraTouchInput returns a FreeCameraTouchInput JavaScript class.
func (ba *Babylon) FreeCameraTouchInput() *FreeCameraTouchInput {
	p := ba.ctx.Get("FreeCameraTouchInput")
	return FreeCameraTouchInputFromJSObject(p, ba.ctx)
}

// FreeCameraTouchInputFromJSObject returns a wrapped FreeCameraTouchInput JavaScript class.
func FreeCameraTouchInputFromJSObject(p js.Value, ctx js.Value) *FreeCameraTouchInput {
	return &FreeCameraTouchInput{p: p, ctx: ctx}
}

// FreeCameraTouchInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func FreeCameraTouchInputArrayToJSArray(array []*FreeCameraTouchInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// FreeCameraTouchInputAttachControlOpts contains optional parameters for FreeCameraTouchInput.AttachControl.
type FreeCameraTouchInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FreeCameraTouchInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#attachcontrol
func (f *FreeCameraTouchInput) AttachControl(element js.Value, opts *FreeCameraTouchInputAttachControlOpts) {
	if opts == nil {
		opts = &FreeCameraTouchInputAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	f.p.Call("attachControl", args...)
}

// CheckInputs calls the CheckInputs method on the FreeCameraTouchInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#checkinputs
func (f *FreeCameraTouchInput) CheckInputs() {

	f.p.Call("checkInputs")
}

// DetachControl calls the DetachControl method on the FreeCameraTouchInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#detachcontrol
func (f *FreeCameraTouchInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the FreeCameraTouchInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#getclassname
func (f *FreeCameraTouchInput) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the FreeCameraTouchInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#getsimplename
func (f *FreeCameraTouchInput) GetSimpleName() string {

	retVal := f.p.Call("getSimpleName")
	return retVal.String()
}

// Camera returns the Camera property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#camera
func (f *FreeCameraTouchInput) Camera() *FreeCamera {
	retVal := f.p.Get("camera")
	return FreeCameraFromJSObject(retVal, f.ctx)
}

// SetCamera sets the Camera property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#camera
func (f *FreeCameraTouchInput) SetCamera(camera *FreeCamera) *FreeCameraTouchInput {
	f.p.Set("camera", camera.JSObject())
	return f
}

// TouchAngularSensibility returns the TouchAngularSensibility property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#touchangularsensibility
func (f *FreeCameraTouchInput) TouchAngularSensibility() float64 {
	retVal := f.p.Get("touchAngularSensibility")
	return retVal.Float()
}

// SetTouchAngularSensibility sets the TouchAngularSensibility property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#touchangularsensibility
func (f *FreeCameraTouchInput) SetTouchAngularSensibility(touchAngularSensibility float64) *FreeCameraTouchInput {
	f.p.Set("touchAngularSensibility", touchAngularSensibility)
	return f
}

// TouchMoveSensibility returns the TouchMoveSensibility property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#touchmovesensibility
func (f *FreeCameraTouchInput) TouchMoveSensibility() float64 {
	retVal := f.p.Get("touchMoveSensibility")
	return retVal.Float()
}

// SetTouchMoveSensibility sets the TouchMoveSensibility property of class FreeCameraTouchInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameratouchinput#touchmovesensibility
func (f *FreeCameraTouchInput) SetTouchMoveSensibility(touchMoveSensibility float64) *FreeCameraTouchInput {
	f.p.Set("touchMoveSensibility", touchMoveSensibility)
	return f
}
