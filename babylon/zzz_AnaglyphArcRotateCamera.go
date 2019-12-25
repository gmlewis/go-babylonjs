// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphArcRotateCamera represents a babylon.js AnaglyphArcRotateCamera.
// Camera used to simulate anaglyphic rendering (based on ArcRotateCamera)
//
// See: http://doc.babylonjs.com/features/cameras#anaglyph-cameras
type AnaglyphArcRotateCamera struct {
	*ArcRotateCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnaglyphArcRotateCamera) JSObject() js.Value { return a.p }

// AnaglyphArcRotateCamera returns a AnaglyphArcRotateCamera JavaScript class.
func (ba *Babylon) AnaglyphArcRotateCamera() *AnaglyphArcRotateCamera {
	p := ba.ctx.Get("AnaglyphArcRotateCamera")
	return AnaglyphArcRotateCameraFromJSObject(p, ba.ctx)
}

// AnaglyphArcRotateCameraFromJSObject returns a wrapped AnaglyphArcRotateCamera JavaScript class.
func AnaglyphArcRotateCameraFromJSObject(p js.Value, ctx js.Value) *AnaglyphArcRotateCamera {
	return &AnaglyphArcRotateCamera{ArcRotateCamera: ArcRotateCameraFromJSObject(p, ctx), ctx: ctx}
}

// AnaglyphArcRotateCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnaglyphArcRotateCameraArrayToJSArray(array []*AnaglyphArcRotateCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnaglyphArcRotateCamera returns a new AnaglyphArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglypharcrotatecamera#constructor
func (ba *Babylon) NewAnaglyphArcRotateCamera(name string, alpha float64, beta float64, radius float64, target *Vector3, interaxialDistance float64, scene *Scene) *AnaglyphArcRotateCamera {

	args := make([]interface{}, 0, 7+0)

	args = append(args, name)
	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, radius)
	args = append(args, target.JSObject())
	args = append(args, interaxialDistance)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("AnaglyphArcRotateCamera").New(args...)
	return AnaglyphArcRotateCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the AnaglyphArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglypharcrotatecamera#getclassname
func (a *AnaglyphArcRotateCamera) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}
