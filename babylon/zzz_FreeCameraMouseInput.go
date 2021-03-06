// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FreeCameraMouseInput represents a babylon.js FreeCameraMouseInput.
// Manage the mouse inputs to control the movement of a free camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FreeCameraMouseInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FreeCameraMouseInput) JSObject() js.Value { return f.p }

// FreeCameraMouseInput returns a FreeCameraMouseInput JavaScript class.
func (ba *Babylon) FreeCameraMouseInput() *FreeCameraMouseInput {
	p := ba.ctx.Get("FreeCameraMouseInput")
	return FreeCameraMouseInputFromJSObject(p, ba.ctx)
}

// FreeCameraMouseInputFromJSObject returns a wrapped FreeCameraMouseInput JavaScript class.
func FreeCameraMouseInputFromJSObject(p js.Value, ctx js.Value) *FreeCameraMouseInput {
	return &FreeCameraMouseInput{p: p, ctx: ctx}
}

// FreeCameraMouseInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func FreeCameraMouseInputArrayToJSArray(array []*FreeCameraMouseInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFreeCameraMouseInputOpts contains optional parameters for NewFreeCameraMouseInput.
type NewFreeCameraMouseInputOpts struct {
	TouchEnabled *bool
}

// NewFreeCameraMouseInput returns a new FreeCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#constructor
func (ba *Babylon) NewFreeCameraMouseInput(opts *NewFreeCameraMouseInputOpts) *FreeCameraMouseInput {
	if opts == nil {
		opts = &NewFreeCameraMouseInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.TouchEnabled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TouchEnabled)
	}

	p := ba.ctx.Get("FreeCameraMouseInput").New(args...)
	return FreeCameraMouseInputFromJSObject(p, ba.ctx)
}

// FreeCameraMouseInputAttachControlOpts contains optional parameters for FreeCameraMouseInput.AttachControl.
type FreeCameraMouseInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FreeCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#attachcontrol
func (f *FreeCameraMouseInput) AttachControl(element js.Value, opts *FreeCameraMouseInputAttachControlOpts) {
	if opts == nil {
		opts = &FreeCameraMouseInputAttachControlOpts{}
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

// DetachControl calls the DetachControl method on the FreeCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#detachcontrol
func (f *FreeCameraMouseInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the FreeCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#getclassname
func (f *FreeCameraMouseInput) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the FreeCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#getsimplename
func (f *FreeCameraMouseInput) GetSimpleName() string {

	retVal := f.p.Call("getSimpleName")
	return retVal.String()
}

// AngularSensibility returns the AngularSensibility property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#angularsensibility
func (f *FreeCameraMouseInput) AngularSensibility() float64 {
	retVal := f.p.Get("angularSensibility")
	return retVal.Float()
}

// SetAngularSensibility sets the AngularSensibility property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#angularsensibility
func (f *FreeCameraMouseInput) SetAngularSensibility(angularSensibility float64) *FreeCameraMouseInput {
	f.p.Set("angularSensibility", angularSensibility)
	return f
}

// Buttons returns the Buttons property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#buttons
func (f *FreeCameraMouseInput) Buttons() []float64 {
	retVal := f.p.Get("buttons")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetButtons sets the Buttons property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#buttons
func (f *FreeCameraMouseInput) SetButtons(buttons []float64) *FreeCameraMouseInput {
	f.p.Set("buttons", buttons)
	return f
}

// Camera returns the Camera property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#camera
func (f *FreeCameraMouseInput) Camera() *FreeCamera {
	retVal := f.p.Get("camera")
	return FreeCameraFromJSObject(retVal, f.ctx)
}

// SetCamera sets the Camera property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#camera
func (f *FreeCameraMouseInput) SetCamera(camera *FreeCamera) *FreeCameraMouseInput {
	f.p.Set("camera", camera.JSObject())
	return f
}

// OnPointerMovedObservable returns the OnPointerMovedObservable property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#onpointermovedobservable
func (f *FreeCameraMouseInput) OnPointerMovedObservable() *Observable {
	retVal := f.p.Get("onPointerMovedObservable")
	return ObservableFromJSObject(retVal, f.ctx)
}

// SetOnPointerMovedObservable sets the OnPointerMovedObservable property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#onpointermovedobservable
func (f *FreeCameraMouseInput) SetOnPointerMovedObservable(onPointerMovedObservable *Observable) *FreeCameraMouseInput {
	f.p.Set("onPointerMovedObservable", onPointerMovedObservable.JSObject())
	return f
}

// TouchEnabled returns the TouchEnabled property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#touchenabled
func (f *FreeCameraMouseInput) TouchEnabled() bool {
	retVal := f.p.Get("touchEnabled")
	return retVal.Bool()
}

// SetTouchEnabled sets the TouchEnabled property of class FreeCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecameramouseinput#touchenabled
func (f *FreeCameraMouseInput) SetTouchEnabled(touchEnabled bool) *FreeCameraMouseInput {
	f.p.Set("touchEnabled", touchEnabled)
	return f
}
