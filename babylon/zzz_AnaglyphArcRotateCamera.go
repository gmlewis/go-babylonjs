// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphArcRotateCamera represents a babylon.js AnaglyphArcRotateCamera.
// Camera used to simulate anaglyphic rendering (based on ArcRotateCamera)

//
// See: http://doc.babylonjs.com/features/cameras#anaglyph-cameras
type AnaglyphArcRotateCamera struct{ *ArcRotateCamera }

// JSObject returns the underlying js.Value.
func (a *AnaglyphArcRotateCamera) JSObject() js.Value { return a.p }

// AnaglyphArcRotateCamera returns a AnaglyphArcRotateCamera JavaScript class.
func (b *Babylon) AnaglyphArcRotateCamera() *AnaglyphArcRotateCamera {
	p := b.ctx.Get("AnaglyphArcRotateCamera")
	return AnaglyphArcRotateCameraFromJSObject(p)
}

// AnaglyphArcRotateCameraFromJSObject returns a wrapped AnaglyphArcRotateCamera JavaScript class.
func AnaglyphArcRotateCameraFromJSObject(p js.Value) *AnaglyphArcRotateCamera {
	return &AnaglyphArcRotateCamera{ArcRotateCameraFromJSObject(p)}
}

// NewAnaglyphArcRotateCamera returns a new AnaglyphArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglypharcrotatecamera
func (b *Babylon) NewAnaglyphArcRotateCamera(todo parameters) *AnaglyphArcRotateCamera {
	p := b.ctx.Get("AnaglyphArcRotateCamera").New(todo)
	return AnaglyphArcRotateCameraFromJSObject(p)
}

// TODO: methods