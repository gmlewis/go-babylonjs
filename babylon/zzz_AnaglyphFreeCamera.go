// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphFreeCamera represents a babylon.js AnaglyphFreeCamera.
// Camera used to simulate anaglyphic rendering (based on FreeCamera)
//
// See: http://doc.babylonjs.com/features/cameras#anaglyph-cameras
type AnaglyphFreeCamera struct {
	*FreeCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnaglyphFreeCamera) JSObject() js.Value { return a.p }

// AnaglyphFreeCamera returns a AnaglyphFreeCamera JavaScript class.
func (ba *Babylon) AnaglyphFreeCamera() *AnaglyphFreeCamera {
	p := ba.ctx.Get("AnaglyphFreeCamera")
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// AnaglyphFreeCameraFromJSObject returns a wrapped AnaglyphFreeCamera JavaScript class.
func AnaglyphFreeCameraFromJSObject(p js.Value, ctx js.Value) *AnaglyphFreeCamera {
	return &AnaglyphFreeCamera{FreeCamera: FreeCameraFromJSObject(p, ctx), ctx: ctx}
}

// AnaglyphFreeCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnaglyphFreeCameraArrayToJSArray(array []*AnaglyphFreeCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnaglyphFreeCamera returns a new AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera
func (ba *Babylon) NewAnaglyphFreeCamera(name string, position *Vector3, interaxialDistance float64, scene *Scene) *AnaglyphFreeCamera {

	args := make([]interface{}, 0, 4+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, interaxialDistance)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("AnaglyphFreeCamera").New(args...)
	return AnaglyphFreeCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the AnaglyphFreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphfreecamera#getclassname
func (a *AnaglyphFreeCamera) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

/*

 */
