// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCameraPointersInput represents a babylon.js ArcRotateCameraPointersInput.
// Manage the pointers inputs to control an arc rotate camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type ArcRotateCameraPointersInput struct {
	*BaseCameraPointersInput
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCameraPointersInput) JSObject() js.Value { return a.p }

// ArcRotateCameraPointersInput returns a ArcRotateCameraPointersInput JavaScript class.
func (ba *Babylon) ArcRotateCameraPointersInput() *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput")
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// ArcRotateCameraPointersInputFromJSObject returns a wrapped ArcRotateCameraPointersInput JavaScript class.
func ArcRotateCameraPointersInputFromJSObject(p js.Value, ctx js.Value) *ArcRotateCameraPointersInput {
	return &ArcRotateCameraPointersInput{BaseCameraPointersInput: BaseCameraPointersInputFromJSObject(p, ctx), ctx: ctx}
}

// ArcRotateCameraPointersInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArcRotateCameraPointersInputArrayToJSArray(array []*ArcRotateCameraPointersInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ArcRotateCameraPointersInputAttachControlOpts contains optional parameters for ArcRotateCameraPointersInput.AttachControl.
type ArcRotateCameraPointersInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the ArcRotateCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#attachcontrol
func (a *ArcRotateCameraPointersInput) AttachControl(element js.Value, opts *ArcRotateCameraPointersInputAttachControlOpts) {
	if opts == nil {
		opts = &ArcRotateCameraPointersInputAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	a.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the ArcRotateCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#detachcontrol
func (a *ArcRotateCameraPointersInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	a.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the ArcRotateCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#getclassname
func (a *ArcRotateCameraPointersInput) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the ArcRotateCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#getsimplename
func (a *ArcRotateCameraPointersInput) GetSimpleName() string {

	retVal := a.p.Call("getSimpleName")
	return retVal.String()
}

/*

// AngularSensibilityX returns the AngularSensibilityX property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#angularsensibilityx
func (a *ArcRotateCameraPointersInput) AngularSensibilityX(angularSensibilityX float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(angularSensibilityX)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAngularSensibilityX sets the AngularSensibilityX property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#angularsensibilityx
func (a *ArcRotateCameraPointersInput) SetAngularSensibilityX(angularSensibilityX float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(angularSensibilityX)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// AngularSensibilityY returns the AngularSensibilityY property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#angularsensibilityy
func (a *ArcRotateCameraPointersInput) AngularSensibilityY(angularSensibilityY float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(angularSensibilityY)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetAngularSensibilityY sets the AngularSensibilityY property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#angularsensibilityy
func (a *ArcRotateCameraPointersInput) SetAngularSensibilityY(angularSensibilityY float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(angularSensibilityY)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// Buttons returns the Buttons property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#buttons
func (a *ArcRotateCameraPointersInput) Buttons(buttons float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(buttons)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetButtons sets the Buttons property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#buttons
func (a *ArcRotateCameraPointersInput) SetButtons(buttons float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(buttons)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#camera
func (a *ArcRotateCameraPointersInput) Camera(camera *ArcRotateCamera) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(camera.JSObject())
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#camera
func (a *ArcRotateCameraPointersInput) SetCamera(camera *ArcRotateCamera) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(camera.JSObject())
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// MultiTouchPanAndZoom returns the MultiTouchPanAndZoom property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#multitouchpanandzoom
func (a *ArcRotateCameraPointersInput) MultiTouchPanAndZoom(multiTouchPanAndZoom bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(multiTouchPanAndZoom)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetMultiTouchPanAndZoom sets the MultiTouchPanAndZoom property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#multitouchpanandzoom
func (a *ArcRotateCameraPointersInput) SetMultiTouchPanAndZoom(multiTouchPanAndZoom bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(multiTouchPanAndZoom)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// MultiTouchPanning returns the MultiTouchPanning property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#multitouchpanning
func (a *ArcRotateCameraPointersInput) MultiTouchPanning(multiTouchPanning bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(multiTouchPanning)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetMultiTouchPanning sets the MultiTouchPanning property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#multitouchpanning
func (a *ArcRotateCameraPointersInput) SetMultiTouchPanning(multiTouchPanning bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(multiTouchPanning)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// PanningSensibility returns the PanningSensibility property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#panningsensibility
func (a *ArcRotateCameraPointersInput) PanningSensibility(panningSensibility float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(panningSensibility)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPanningSensibility sets the PanningSensibility property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#panningsensibility
func (a *ArcRotateCameraPointersInput) SetPanningSensibility(panningSensibility float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(panningSensibility)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// PinchDeltaPercentage returns the PinchDeltaPercentage property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchdeltapercentage
func (a *ArcRotateCameraPointersInput) PinchDeltaPercentage(pinchDeltaPercentage float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchDeltaPercentage)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPinchDeltaPercentage sets the PinchDeltaPercentage property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchdeltapercentage
func (a *ArcRotateCameraPointersInput) SetPinchDeltaPercentage(pinchDeltaPercentage float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchDeltaPercentage)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// PinchInwards returns the PinchInwards property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchinwards
func (a *ArcRotateCameraPointersInput) PinchInwards(pinchInwards bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchInwards)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPinchInwards sets the PinchInwards property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchinwards
func (a *ArcRotateCameraPointersInput) SetPinchInwards(pinchInwards bool) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchInwards)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// PinchPrecision returns the PinchPrecision property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchprecision
func (a *ArcRotateCameraPointersInput) PinchPrecision(pinchPrecision float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchPrecision)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

// SetPinchPrecision sets the PinchPrecision property of class ArcRotateCameraPointersInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerapointersinput#pinchprecision
func (a *ArcRotateCameraPointersInput) SetPinchPrecision(pinchPrecision float64) *ArcRotateCameraPointersInput {
	p := ba.ctx.Get("ArcRotateCameraPointersInput").New(pinchPrecision)
	return ArcRotateCameraPointersInputFromJSObject(p, ba.ctx)
}

*/
