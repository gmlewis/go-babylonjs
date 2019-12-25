// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GamepadCamera represents a babylon.js GamepadCamera.
// This represents a FPS type of camera. This is only here for back compat purpose.
// Please use the UniversalCamera instead as both are identical.
//
// See: http://doc.babylonjs.com/features/cameras#universal-camera
type GamepadCamera struct {
	*UniversalCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GamepadCamera) JSObject() js.Value { return g.p }

// GamepadCamera returns a GamepadCamera JavaScript class.
func (ba *Babylon) GamepadCamera() *GamepadCamera {
	p := ba.ctx.Get("GamepadCamera")
	return GamepadCameraFromJSObject(p, ba.ctx)
}

// GamepadCameraFromJSObject returns a wrapped GamepadCamera JavaScript class.
func GamepadCameraFromJSObject(p js.Value, ctx js.Value) *GamepadCamera {
	return &GamepadCamera{UniversalCamera: UniversalCameraFromJSObject(p, ctx), ctx: ctx}
}

// GamepadCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func GamepadCameraArrayToJSArray(array []*GamepadCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGamepadCamera returns a new GamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadcamera#constructor
func (ba *Babylon) NewGamepadCamera(name string, position *Vector3, scene *Scene) *GamepadCamera {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("GamepadCamera").New(args...)
	return GamepadCameraFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the GamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadcamera#getclassname
func (g *GamepadCamera) GetClassName() string {

	retVal := g.p.Call("getClassName")
	return retVal.String()
}
