// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FollowCameraPointersInput represents a babylon.js FollowCameraPointersInput.
// Manage the pointers inputs to control an follow camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FollowCameraPointersInput struct {
	*BaseCameraPointersInput
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FollowCameraPointersInput) JSObject() js.Value { return f.p }

// FollowCameraPointersInput returns a FollowCameraPointersInput JavaScript class.
func (ba *Babylon) FollowCameraPointersInput() *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput")
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// FollowCameraPointersInputFromJSObject returns a wrapped FollowCameraPointersInput JavaScript class.
func FollowCameraPointersInputFromJSObject(p js.Value, ctx js.Value) *FollowCameraPointersInput {
	return &FollowCameraPointersInput{BaseCameraPointersInput: BaseCameraPointersInputFromJSObject(p, ctx), ctx: ctx}
}

// FollowCameraPointersInputAttachControlOpts contains optional parameters for FollowCameraPointersInput.AttachControl.
type FollowCameraPointersInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FollowCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#attachcontrol
func (f *FollowCameraPointersInput) AttachControl(element js.Value, opts *FollowCameraPointersInputAttachControlOpts) {
	if opts == nil {
		opts = &FollowCameraPointersInputAttachControlOpts{}
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

// DetachControl calls the DetachControl method on the FollowCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#detachcontrol
func (f *FollowCameraPointersInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the FollowCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#getclassname
func (f *FollowCameraPointersInput) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getClassName", args...)
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the FollowCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#getsimplename
func (f *FollowCameraPointersInput) GetSimpleName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := f.p.Call("getSimpleName", args...)
	return retVal.String()
}

/*

// AngularSensibilityX returns the AngularSensibilityX property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#angularsensibilityx
func (f *FollowCameraPointersInput) AngularSensibilityX(angularSensibilityX float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(angularSensibilityX)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAngularSensibilityX sets the AngularSensibilityX property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#angularsensibilityx
func (f *FollowCameraPointersInput) SetAngularSensibilityX(angularSensibilityX float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(angularSensibilityX)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AngularSensibilityY returns the AngularSensibilityY property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#angularsensibilityy
func (f *FollowCameraPointersInput) AngularSensibilityY(angularSensibilityY float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(angularSensibilityY)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAngularSensibilityY sets the AngularSensibilityY property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#angularsensibilityy
func (f *FollowCameraPointersInput) SetAngularSensibilityY(angularSensibilityY float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(angularSensibilityY)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisPinchControlHeight returns the AxisPinchControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolheight
func (f *FollowCameraPointersInput) AxisPinchControlHeight(axisPinchControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisPinchControlHeight sets the AxisPinchControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolheight
func (f *FollowCameraPointersInput) SetAxisPinchControlHeight(axisPinchControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisPinchControlRadius returns the AxisPinchControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolradius
func (f *FollowCameraPointersInput) AxisPinchControlRadius(axisPinchControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisPinchControlRadius sets the AxisPinchControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolradius
func (f *FollowCameraPointersInput) SetAxisPinchControlRadius(axisPinchControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisPinchControlRotation returns the AxisPinchControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolrotation
func (f *FollowCameraPointersInput) AxisPinchControlRotation(axisPinchControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisPinchControlRotation sets the AxisPinchControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axispinchcontrolrotation
func (f *FollowCameraPointersInput) SetAxisPinchControlRotation(axisPinchControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisPinchControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisXControlHeight returns the AxisXControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolheight
func (f *FollowCameraPointersInput) AxisXControlHeight(axisXControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisXControlHeight sets the AxisXControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolheight
func (f *FollowCameraPointersInput) SetAxisXControlHeight(axisXControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisXControlRadius returns the AxisXControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolradius
func (f *FollowCameraPointersInput) AxisXControlRadius(axisXControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisXControlRadius sets the AxisXControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolradius
func (f *FollowCameraPointersInput) SetAxisXControlRadius(axisXControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisXControlRotation returns the AxisXControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolrotation
func (f *FollowCameraPointersInput) AxisXControlRotation(axisXControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisXControlRotation sets the AxisXControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisxcontrolrotation
func (f *FollowCameraPointersInput) SetAxisXControlRotation(axisXControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisXControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisYControlHeight returns the AxisYControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolheight
func (f *FollowCameraPointersInput) AxisYControlHeight(axisYControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisYControlHeight sets the AxisYControlHeight property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolheight
func (f *FollowCameraPointersInput) SetAxisYControlHeight(axisYControlHeight bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlHeight)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisYControlRadius returns the AxisYControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolradius
func (f *FollowCameraPointersInput) AxisYControlRadius(axisYControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisYControlRadius sets the AxisYControlRadius property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolradius
func (f *FollowCameraPointersInput) SetAxisYControlRadius(axisYControlRadius bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlRadius)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// AxisYControlRotation returns the AxisYControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolrotation
func (f *FollowCameraPointersInput) AxisYControlRotation(axisYControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAxisYControlRotation sets the AxisYControlRotation property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#axisycontrolrotation
func (f *FollowCameraPointersInput) SetAxisYControlRotation(axisYControlRotation bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(axisYControlRotation)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// Buttons returns the Buttons property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#buttons
func (f *FollowCameraPointersInput) Buttons(buttons float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(buttons)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetButtons sets the Buttons property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#buttons
func (f *FollowCameraPointersInput) SetButtons(buttons float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(buttons)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#camera
func (f *FollowCameraPointersInput) Camera(camera *FollowCamera) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(camera.JSObject())
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#camera
func (f *FollowCameraPointersInput) SetCamera(camera *FollowCamera) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(camera.JSObject())
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// PinchDeltaPercentage returns the PinchDeltaPercentage property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#pinchdeltapercentage
func (f *FollowCameraPointersInput) PinchDeltaPercentage(pinchDeltaPercentage float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(pinchDeltaPercentage)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPinchDeltaPercentage sets the PinchDeltaPercentage property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#pinchdeltapercentage
func (f *FollowCameraPointersInput) SetPinchDeltaPercentage(pinchDeltaPercentage float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(pinchDeltaPercentage)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// PinchPrecision returns the PinchPrecision property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#pinchprecision
func (f *FollowCameraPointersInput) PinchPrecision(pinchPrecision float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(pinchPrecision)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPinchPrecision sets the PinchPrecision property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#pinchprecision
func (f *FollowCameraPointersInput) SetPinchPrecision(pinchPrecision float64) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(pinchPrecision)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// WarningEnable returns the WarningEnable property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#warningenable
func (f *FollowCameraPointersInput) WarningEnable(warningEnable bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(warningEnable)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetWarningEnable sets the WarningEnable property of class FollowCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerapointersinput#warningenable
func (f *FollowCameraPointersInput) SetWarningEnable(warningEnable bool) *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput").New(warningEnable)
	return FollowCameraPointersInputFromJSObject(p, ba.ctx)
}

*/
