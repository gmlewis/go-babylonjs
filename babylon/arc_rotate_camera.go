package babylon

import "syscall/js"

// ArcRotateCamera represents a babylon.js ArcRotateCamera.
type ArcRotateCamera struct{ *TargetCamera }

// JSObject returns the underlying js.Value.
func (a *ArcRotateCamera) JSObject() js.Value { return a.p }

// ArcRotateCamera returns a ArcRotateCamera JavaScript class.
func (b *Babylon) ArcRotateCamera() *ArcRotateCamera {
	p := b.ctx.Get("ArcRotateCamera")
	return ArcRotateCameraFromJSObject(p)
}

// ArcRotateCameraFromJSObject returns a wrapped ArcRotateCamera JavaScript class.
func ArcRotateCameraFromJSObject(p js.Value) *ArcRotateCamera {
	return &ArcRotateCamera{TargetCameraFromJSObject(p)}
}

// NewArcRotateCamera returns a new ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera
func (b *Babylon) NewArcRotateCamera(name string, alpha, beta, radius float64, target *Vector3, scene *Scene) *ArcRotateCamera {
	p := b.ctx.Get("ArcRotateCamera").New(name, alpha, beta, radius, target.JSObject(), scene.JSObject())
	return ArcRotateCameraFromJSObject(p)
}

// AttachControl calls the JavaScript method of the same name.
func (a *ArcRotateCamera) AttachControl(canvas js.Value, b bool) {
	a.p.Call("attachControl", canvas, b)
}
