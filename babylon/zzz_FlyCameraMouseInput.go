// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FlyCameraMouseInput represents a babylon.js FlyCameraMouseInput.
// Listen to mouse events to control the camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FlyCameraMouseInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FlyCameraMouseInput) JSObject() js.Value { return f.p }

// FlyCameraMouseInput returns a FlyCameraMouseInput JavaScript class.
func (ba *Babylon) FlyCameraMouseInput() *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput")
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// FlyCameraMouseInputFromJSObject returns a wrapped FlyCameraMouseInput JavaScript class.
func FlyCameraMouseInputFromJSObject(p js.Value, ctx js.Value) *FlyCameraMouseInput {
	return &FlyCameraMouseInput{p: p, ctx: ctx}
}

// NewFlyCameraMouseInputOpts contains optional parameters for NewFlyCameraMouseInput.
type NewFlyCameraMouseInputOpts struct {
	TouchEnabled *bool
}

// NewFlyCameraMouseInput returns a new FlyCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput
func (ba *Babylon) NewFlyCameraMouseInput(opts *NewFlyCameraMouseInputOpts) *FlyCameraMouseInput {
	if opts == nil {
		opts = &NewFlyCameraMouseInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.TouchEnabled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TouchEnabled)
	}

	p := ba.ctx.Get("FlyCameraMouseInput").New(args...)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// FlyCameraMouseInputAttachControlOpts contains optional parameters for FlyCameraMouseInput.AttachControl.
type FlyCameraMouseInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FlyCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#attachcontrol
func (f *FlyCameraMouseInput) AttachControl(element js.Value, opts *FlyCameraMouseInputAttachControlOpts) {
	if opts == nil {
		opts = &FlyCameraMouseInputAttachControlOpts{}
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

// DetachControl calls the DetachControl method on the FlyCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#detachcontrol
func (f *FlyCameraMouseInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the FlyCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#getclassname
func (f *FlyCameraMouseInput) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getClassName", args...)
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the FlyCameraMouseInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#getsimplename
func (f *FlyCameraMouseInput) GetSimpleName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getSimpleName", args...)
	return retVal.String()
}

/*

// ActiveButton returns the ActiveButton property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#activebutton
func (f *FlyCameraMouseInput) ActiveButton(activeButton float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(activeButton)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetActiveButton sets the ActiveButton property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#activebutton
func (f *FlyCameraMouseInput) SetActiveButton(activeButton float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(activeButton)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// AngularSensibility returns the AngularSensibility property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#angularsensibility
func (f *FlyCameraMouseInput) AngularSensibility(angularSensibility float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(angularSensibility)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetAngularSensibility sets the AngularSensibility property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#angularsensibility
func (f *FlyCameraMouseInput) SetAngularSensibility(angularSensibility float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(angularSensibility)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// Buttons returns the Buttons property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttons
func (f *FlyCameraMouseInput) Buttons(buttons float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttons)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetButtons sets the Buttons property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttons
func (f *FlyCameraMouseInput) SetButtons(buttons float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttons)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// ButtonsPitch returns the ButtonsPitch property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonspitch
func (f *FlyCameraMouseInput) ButtonsPitch(buttonsPitch float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsPitch)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetButtonsPitch sets the ButtonsPitch property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonspitch
func (f *FlyCameraMouseInput) SetButtonsPitch(buttonsPitch float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsPitch)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// ButtonsRoll returns the ButtonsRoll property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonsroll
func (f *FlyCameraMouseInput) ButtonsRoll(buttonsRoll float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsRoll)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetButtonsRoll sets the ButtonsRoll property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonsroll
func (f *FlyCameraMouseInput) SetButtonsRoll(buttonsRoll float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsRoll)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// ButtonsYaw returns the ButtonsYaw property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonsyaw
func (f *FlyCameraMouseInput) ButtonsYaw(buttonsYaw float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsYaw)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetButtonsYaw sets the ButtonsYaw property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#buttonsyaw
func (f *FlyCameraMouseInput) SetButtonsYaw(buttonsYaw float64) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(buttonsYaw)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#camera
func (f *FlyCameraMouseInput) Camera(camera *FlyCamera) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(camera.JSObject())
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#camera
func (f *FlyCameraMouseInput) SetCamera(camera *FlyCamera) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(camera.JSObject())
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// TouchEnabled returns the TouchEnabled property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#touchenabled
func (f *FlyCameraMouseInput) TouchEnabled(touchEnabled bool) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(touchEnabled)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

// SetTouchEnabled sets the TouchEnabled property of class FlyCameraMouseInput.
//
// https://doc.babylonjs.com/api/classes/babylon.flycameramouseinput#touchenabled
func (f *FlyCameraMouseInput) SetTouchEnabled(touchEnabled bool) *FlyCameraMouseInput {
	p := ba.ctx.Get("FlyCameraMouseInput").New(touchEnabled)
	return FlyCameraMouseInputFromJSObject(p, ba.ctx)
}

*/
