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

// NewTouchCamera returns a new TouchCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.touchcamera
func (ba *Babylon) NewTouchCamera(name string, position *Vector3, scene *Scene) *TouchCamera {
	p := ba.ctx.Get("TouchCamera").New(name, position.JSObject(), scene.JSObject())
	return TouchCameraFromJSObject(p, ba.ctx)
}

// TODO: methods
