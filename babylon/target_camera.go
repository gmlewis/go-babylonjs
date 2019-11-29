package babylon

import "syscall/js"

// TargetCamera represents a babylon.js TargetCamera.
type TargetCamera struct{ *Camera }

// JSObject returns the underlying js.Value.
func (s *TargetCamera) JSObject() js.Value { return s.p }

// TargetCamera returns a TargetCamera JavaScript class.
func (b *Babylon) TargetCamera() *TargetCamera {
	p := b.ctx.Get("TargetCamera")
	return TargetCameraFromJSObject(p)
}

// TargetCameraFromJSObject returns a wrapped TargetCamera JavaScript class.
func TargetCameraFromJSObject(p js.Value) *TargetCamera {
	return &TargetCamera{CameraFromJSObject(p)}
}

// NewTargetCamera returns a new TargetCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.Targetcamera
func (b *Babylon) NewTargetCamera(name string, position *Vector3, scene *Scene) *TargetCamera {
	p := b.ctx.Get("TargetCamera").New(name, position.JSObject(), scene.JSObject())
	return TargetCameraFromJSObject(p)
}
