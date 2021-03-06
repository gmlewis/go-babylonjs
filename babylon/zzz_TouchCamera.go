// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TouchCamera represents a babylon.js TouchCamera.
// This represents a FPS type of camera controlled by touch.
// This is like a universal camera minus the Gamepad controls.
//
// See: http://doc.babylonjs.com/features/cameras#universal-camera
type TouchCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TouchCamera) JSObject() js.Value { return t.p }

// TouchCamera returns a TouchCamera JavaScript class.
func (ba *Babylon) TouchCamera() *TouchCamera {
	p := ba.ctx.Get("TouchCamera")
	return TouchCameraFromJSObject(p, ba.ctx)
}

// TouchCameraFromJSObject returns a wrapped TouchCamera JavaScript class.
func TouchCameraFromJSObject(p js.Value, ctx js.Value) *TouchCamera {
	return &TouchCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// TouchCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func TouchCameraArrayToJSArray(array []*TouchCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTouchCamera returns a new TouchCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#constructor
func (ba *Babylon) NewTouchCamera(name string, position *Vector3, scene *Scene) *TouchCamera {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("TouchCamera").New(args...)
	return TouchCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the TouchCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#getclassname
func (t *TouchCamera) GetClassName() string {

	retVal := t.p.Call("getClassName")
	return retVal.String()
}

// TouchAngularSensibility returns the TouchAngularSensibility property of class TouchCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#touchangularsensibility
func (t *TouchCamera) TouchAngularSensibility() float64 {
	retVal := t.p.Get("touchAngularSensibility")
	return retVal.Float()
}

// SetTouchAngularSensibility sets the TouchAngularSensibility property of class TouchCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#touchangularsensibility
func (t *TouchCamera) SetTouchAngularSensibility(touchAngularSensibility float64) *TouchCamera {
	t.p.Set("touchAngularSensibility", touchAngularSensibility)
	return t
}

// TouchMoveSensibility returns the TouchMoveSensibility property of class TouchCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#touchmovesensibility
func (t *TouchCamera) TouchMoveSensibility() float64 {
	retVal := t.p.Get("touchMoveSensibility")
	return retVal.Float()
}

// SetTouchMoveSensibility sets the TouchMoveSensibility property of class TouchCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera#touchmovesensibility
func (t *TouchCamera) SetTouchMoveSensibility(touchMoveSensibility float64) *TouchCamera {
	t.p.Set("touchMoveSensibility", touchMoveSensibility)
	return t
}
