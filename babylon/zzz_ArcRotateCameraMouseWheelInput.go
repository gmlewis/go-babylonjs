// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCameraMouseWheelInput represents a babylon.js ArcRotateCameraMouseWheelInput.
// Manage the mouse wheel inputs to control an arc rotate camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type ArcRotateCameraMouseWheelInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCameraMouseWheelInput) JSObject() js.Value { return a.p }

// ArcRotateCameraMouseWheelInput returns a ArcRotateCameraMouseWheelInput JavaScript class.
func (ba *Babylon) ArcRotateCameraMouseWheelInput() *ArcRotateCameraMouseWheelInput {
	p := ba.ctx.Get("ArcRotateCameraMouseWheelInput")
	return ArcRotateCameraMouseWheelInputFromJSObject(p, ba.ctx)
}

// ArcRotateCameraMouseWheelInputFromJSObject returns a wrapped ArcRotateCameraMouseWheelInput JavaScript class.
func ArcRotateCameraMouseWheelInputFromJSObject(p js.Value, ctx js.Value) *ArcRotateCameraMouseWheelInput {
	return &ArcRotateCameraMouseWheelInput{p: p, ctx: ctx}
}

// TODO: methods
