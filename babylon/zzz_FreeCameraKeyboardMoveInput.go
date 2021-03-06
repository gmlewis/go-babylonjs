// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FreeCameraKeyboardMoveInput represents a babylon.js FreeCameraKeyboardMoveInput.
// Manage the keyboard inputs to control the movement of a free camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FreeCameraKeyboardMoveInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FreeCameraKeyboardMoveInput) JSObject() js.Value { return f.p }

// FreeCameraKeyboardMoveInput returns a FreeCameraKeyboardMoveInput JavaScript class.
func (ba *Babylon) FreeCameraKeyboardMoveInput() *FreeCameraKeyboardMoveInput {
	p := ba.ctx.Get("FreeCameraKeyboardMoveInput")
	return FreeCameraKeyboardMoveInputFromJSObject(p, ba.ctx)
}

// FreeCameraKeyboardMoveInputFromJSObject returns a wrapped FreeCameraKeyboardMoveInput JavaScript class.
func FreeCameraKeyboardMoveInputFromJSObject(p js.Value, ctx js.Value) *FreeCameraKeyboardMoveInput {
	return &FreeCameraKeyboardMoveInput{p: p, ctx: ctx}
}

// FreeCameraKeyboardMoveInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func FreeCameraKeyboardMoveInputArrayToJSArray(array []*FreeCameraKeyboardMoveInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// FreeCameraKeyboardMoveInputAttachControlOpts contains optional parameters for FreeCameraKeyboardMoveInput.AttachControl.
type FreeCameraKeyboardMoveInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the FreeCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#attachcontrol
func (f *FreeCameraKeyboardMoveInput) AttachControl(element js.Value, opts *FreeCameraKeyboardMoveInputAttachControlOpts) {
	if opts == nil {
		opts = &FreeCameraKeyboardMoveInputAttachControlOpts{}
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

// CheckInputs calls the CheckInputs method on the FreeCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#checkinputs
func (f *FreeCameraKeyboardMoveInput) CheckInputs() {

	f.p.Call("checkInputs")
}

// DetachControl calls the DetachControl method on the FreeCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#detachcontrol
func (f *FreeCameraKeyboardMoveInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	f.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the FreeCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#getclassname
func (f *FreeCameraKeyboardMoveInput) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the FreeCameraKeyboardMoveInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#getsimplename
func (f *FreeCameraKeyboardMoveInput) GetSimpleName() string {

	retVal := f.p.Call("getSimpleName")
	return retVal.String()
}

// Camera returns the Camera property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#camera
func (f *FreeCameraKeyboardMoveInput) Camera() *FreeCamera {
	retVal := f.p.Get("camera")
	return FreeCameraFromJSObject(retVal, f.ctx)
}

// SetCamera sets the Camera property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#camera
func (f *FreeCameraKeyboardMoveInput) SetCamera(camera *FreeCamera) *FreeCameraKeyboardMoveInput {
	f.p.Set("camera", camera.JSObject())
	return f
}

// KeysDown returns the KeysDown property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysdown
func (f *FreeCameraKeyboardMoveInput) KeysDown() []float64 {
	retVal := f.p.Get("keysDown")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysDown sets the KeysDown property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysdown
func (f *FreeCameraKeyboardMoveInput) SetKeysDown(keysDown []float64) *FreeCameraKeyboardMoveInput {
	f.p.Set("keysDown", keysDown)
	return f
}

// KeysLeft returns the KeysLeft property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysleft
func (f *FreeCameraKeyboardMoveInput) KeysLeft() []float64 {
	retVal := f.p.Get("keysLeft")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysLeft sets the KeysLeft property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysleft
func (f *FreeCameraKeyboardMoveInput) SetKeysLeft(keysLeft []float64) *FreeCameraKeyboardMoveInput {
	f.p.Set("keysLeft", keysLeft)
	return f
}

// KeysRight returns the KeysRight property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysright
func (f *FreeCameraKeyboardMoveInput) KeysRight() []float64 {
	retVal := f.p.Get("keysRight")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysRight sets the KeysRight property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysright
func (f *FreeCameraKeyboardMoveInput) SetKeysRight(keysRight []float64) *FreeCameraKeyboardMoveInput {
	f.p.Set("keysRight", keysRight)
	return f
}

// KeysUp returns the KeysUp property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysup
func (f *FreeCameraKeyboardMoveInput) KeysUp() []float64 {
	retVal := f.p.Get("keysUp")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysUp sets the KeysUp property of class FreeCameraKeyboardMoveInput.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamerakeyboardmoveinput#keysup
func (f *FreeCameraKeyboardMoveInput) SetKeysUp(keysUp []float64) *FreeCameraKeyboardMoveInput {
	f.p.Set("keysUp", keysUp)
	return f
}
