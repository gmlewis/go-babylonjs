// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StereoscopicArcRotateCamera represents a babylon.js StereoscopicArcRotateCamera.
// Camera used to simulate stereoscopic rendering (based on ArcRotateCamera)
//
// See: http://doc.babylonjs.com/features/cameras
type StereoscopicArcRotateCamera struct {
	*ArcRotateCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StereoscopicArcRotateCamera) JSObject() js.Value { return s.p }

// StereoscopicArcRotateCamera returns a StereoscopicArcRotateCamera JavaScript class.
func (ba *Babylon) StereoscopicArcRotateCamera() *StereoscopicArcRotateCamera {
	p := ba.ctx.Get("StereoscopicArcRotateCamera")
	return StereoscopicArcRotateCameraFromJSObject(p, ba.ctx)
}

// StereoscopicArcRotateCameraFromJSObject returns a wrapped StereoscopicArcRotateCamera JavaScript class.
func StereoscopicArcRotateCameraFromJSObject(p js.Value, ctx js.Value) *StereoscopicArcRotateCamera {
	return &StereoscopicArcRotateCamera{ArcRotateCamera: ArcRotateCameraFromJSObject(p, ctx), ctx: ctx}
}

// StereoscopicArcRotateCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func StereoscopicArcRotateCameraArrayToJSArray(array []*StereoscopicArcRotateCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStereoscopicArcRotateCamera returns a new StereoscopicArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.stereoscopicarcrotatecamera
func (ba *Babylon) NewStereoscopicArcRotateCamera(name string, alpha float64, beta float64, radius float64, target *Vector3, interaxialDistance float64, isStereoscopicSideBySide bool, scene *Scene) *StereoscopicArcRotateCamera {

	args := make([]interface{}, 0, 8+0)

	args = append(args, name)
	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, radius)
	args = append(args, target.JSObject())
	args = append(args, interaxialDistance)
	args = append(args, isStereoscopicSideBySide)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("StereoscopicArcRotateCamera").New(args...)
	return StereoscopicArcRotateCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the StereoscopicArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.stereoscopicarcrotatecamera#getclassname
func (s *StereoscopicArcRotateCamera) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

/*

 */
