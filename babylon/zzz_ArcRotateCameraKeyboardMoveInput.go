// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCameraKeyboardMoveInput represents a babylon.js ArcRotateCameraKeyboardMoveInput.
// Manage the keyboard inputs to control the movement of an arc rotate camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type ArcRotateCameraKeyboardMoveInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCameraKeyboardMoveInput) JSObject() js.Value { return a.p }

// ArcRotateCameraKeyboardMoveInput returns a ArcRotateCameraKeyboardMoveInput JavaScript class.
func (ba *Babylon) ArcRotateCameraKeyboardMoveInput() *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput")
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// ArcRotateCameraKeyboardMoveInputFromJSObject returns a wrapped ArcRotateCameraKeyboardMoveInput JavaScript class.
func ArcRotateCameraKeyboardMoveInputFromJSObject(p js.Value, ctx js.Value) *ArcRotateCameraKeyboardMoveInput {
	return &ArcRotateCameraKeyboardMoveInput{p: p, ctx: ctx}
}

// ArcRotateCameraKeyboardMoveInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArcRotateCameraKeyboardMoveInputArrayToJSArray(array []*ArcRotateCameraKeyboardMoveInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ArcRotateCameraKeyboardMoveInputAttachControlOpts contains optional parameters for ArcRotateCameraKeyboardMoveInput.AttachControl.
type ArcRotateCameraKeyboardMoveInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the ArcRotateCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#attachcontrol
func (a *ArcRotateCameraKeyboardMoveInput) AttachControl(element js.Value, opts *ArcRotateCameraKeyboardMoveInputAttachControlOpts) {
	if opts == nil {
		opts = &ArcRotateCameraKeyboardMoveInputAttachControlOpts{}
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

// CheckInputs calls the CheckInputs method on the ArcRotateCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#checkinputs
func (a *ArcRotateCameraKeyboardMoveInput) CheckInputs() {

	a.p.Call("checkInputs")
}

// DetachControl calls the DetachControl method on the ArcRotateCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#detachcontrol
func (a *ArcRotateCameraKeyboardMoveInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	a.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the ArcRotateCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#getclassname
func (a *ArcRotateCameraKeyboardMoveInput) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the ArcRotateCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#getsimplename
func (a *ArcRotateCameraKeyboardMoveInput) GetSimpleName() string {

	retVal := a.p.Call("getSimpleName")
	return retVal.String()
}

/*

// AngularSpeed returns the AngularSpeed property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#angularspeed
func (a *ArcRotateCameraKeyboardMoveInput) AngularSpeed(angularSpeed float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(angularSpeed)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetAngularSpeed sets the AngularSpeed property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#angularspeed
func (a *ArcRotateCameraKeyboardMoveInput) SetAngularSpeed(angularSpeed float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(angularSpeed)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#camera
func (a *ArcRotateCameraKeyboardMoveInput) Camera(camera *ArcRotateCamera) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(camera.JSObject())
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#camera
func (a *ArcRotateCameraKeyboardMoveInput) SetCamera(camera *ArcRotateCamera) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(camera.JSObject())
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// KeysDown returns the KeysDown property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysdown
func (a *ArcRotateCameraKeyboardMoveInput) KeysDown(keysDown float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysDown)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetKeysDown sets the KeysDown property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysdown
func (a *ArcRotateCameraKeyboardMoveInput) SetKeysDown(keysDown float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysDown)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// KeysLeft returns the KeysLeft property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysleft
func (a *ArcRotateCameraKeyboardMoveInput) KeysLeft(keysLeft float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysLeft)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetKeysLeft sets the KeysLeft property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysleft
func (a *ArcRotateCameraKeyboardMoveInput) SetKeysLeft(keysLeft float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysLeft)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// KeysReset returns the KeysReset property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysreset
func (a *ArcRotateCameraKeyboardMoveInput) KeysReset(keysReset float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysReset)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetKeysReset sets the KeysReset property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysreset
func (a *ArcRotateCameraKeyboardMoveInput) SetKeysReset(keysReset float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysReset)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// KeysRight returns the KeysRight property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysright
func (a *ArcRotateCameraKeyboardMoveInput) KeysRight(keysRight float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysRight)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetKeysRight sets the KeysRight property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysright
func (a *ArcRotateCameraKeyboardMoveInput) SetKeysRight(keysRight float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysRight)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// KeysUp returns the KeysUp property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysup
func (a *ArcRotateCameraKeyboardMoveInput) KeysUp(keysUp float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysUp)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetKeysUp sets the KeysUp property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#keysup
func (a *ArcRotateCameraKeyboardMoveInput) SetKeysUp(keysUp float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(keysUp)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// PanningSensibility returns the PanningSensibility property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#panningsensibility
func (a *ArcRotateCameraKeyboardMoveInput) PanningSensibility(panningSensibility float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(panningSensibility)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetPanningSensibility sets the PanningSensibility property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#panningsensibility
func (a *ArcRotateCameraKeyboardMoveInput) SetPanningSensibility(panningSensibility float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(panningSensibility)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// UseAltToZoom returns the UseAltToZoom property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#usealttozoom
func (a *ArcRotateCameraKeyboardMoveInput) UseAltToZoom(useAltToZoom bool) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(useAltToZoom)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetUseAltToZoom sets the UseAltToZoom property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#usealttozoom
func (a *ArcRotateCameraKeyboardMoveInput) SetUseAltToZoom(useAltToZoom bool) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(useAltToZoom)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// ZoomingSensibility returns the ZoomingSensibility property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#zoomingsensibility
func (a *ArcRotateCameraKeyboardMoveInput) ZoomingSensibility(zoomingSensibility float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(zoomingSensibility)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// SetZoomingSensibility sets the ZoomingSensibility property of class ArcRotateCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamerakeyboardmoveinput#zoomingsensibility
func (a *ArcRotateCameraKeyboardMoveInput) SetZoomingSensibility(zoomingSensibility float64) *ArcRotateCameraKeyboardMoveInput {
	p := ba.ctx.Get("ArcRotateCameraKeyboardMoveInput").New(zoomingSensibility)
	return ArcRotateCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

*/
