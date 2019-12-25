// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StereoscopicFreeCamera represents a babylon.js StereoscopicFreeCamera.
// Camera used to simulate stereoscopic rendering (based on FreeCamera)
//
// See: http://doc.babylonjs.com/features/cameras
type StereoscopicFreeCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StereoscopicFreeCamera) JSObject() js.Value { return s.p }

// StereoscopicFreeCamera returns a StereoscopicFreeCamera JavaScript class.
func (ba *Babylon) StereoscopicFreeCamera() *StereoscopicFreeCamera {
	p := ba.ctx.Get("StereoscopicFreeCamera")
	return StereoscopicFreeCameraFromJSObject(p, ba.ctx)
}

// StereoscopicFreeCameraFromJSObject returns a wrapped StereoscopicFreeCamera JavaScript class.
func StereoscopicFreeCameraFromJSObject(p js.Value, ctx js.Value) *StereoscopicFreeCamera {
	return &StereoscopicFreeCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// StereoscopicFreeCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func StereoscopicFreeCameraArrayToJSArray(array []*StereoscopicFreeCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStereoscopicFreeCamera returns a new StereoscopicFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.stereoscopicfreecamera#constructor
func (ba *Babylon) NewStereoscopicFreeCamera(name string, position *Vector3, interaxialDistance float64, isStereoscopicSideBySide bool, scene *Scene) *StereoscopicFreeCamera {

	args := make([]interface{}, 0, 5+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, interaxialDistance)
	args = append(args, isStereoscopicSideBySide)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("StereoscopicFreeCamera").New(args...)
	return StereoscopicFreeCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the StereoscopicFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.stereoscopicfreecamera#getclassname
func (s *StereoscopicFreeCamera) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}
