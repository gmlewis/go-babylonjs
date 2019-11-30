// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FlyCameraKeyboardInput represents a babylon.js FlyCameraKeyboardInput.
// Listen to keyboard events to control the camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FlyCameraKeyboardInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FlyCameraKeyboardInput) JSObject() js.Value { return f.p }

// FlyCameraKeyboardInput returns a FlyCameraKeyboardInput JavaScript class.
func (ba *Babylon) FlyCameraKeyboardInput() *FlyCameraKeyboardInput {
	p := ba.ctx.Get("FlyCameraKeyboardInput")
	return FlyCameraKeyboardInputFromJSObject(p, ba.ctx)
}

// FlyCameraKeyboardInputFromJSObject returns a wrapped FlyCameraKeyboardInput JavaScript class.
func FlyCameraKeyboardInputFromJSObject(p js.Value, ctx js.Value) *FlyCameraKeyboardInput {
	return &FlyCameraKeyboardInput{p: p, ctx: ctx}
}

// TODO: methods
