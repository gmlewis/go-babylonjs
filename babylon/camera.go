package babylon

/*
// Camera represents a babylon.js Camera.
type Camera struct{ *Node }

// JSObject returns the underlying js.Value.
func (s *Camera) JSObject() js.Value { return s.p }

// Camera returns a Camera JavaScript class.
func (b *Babylon) Camera() *Camera {
	p := b.ctx.Get("Camera")
	return CameraFromJSObject(p)
}

// CameraFromJSObject returns a wrapped Camera JavaScript class.
func CameraFromJSObject(p js.Value) *Camera {
	return &Camera{NodeFromJSObject(p)}
}

// NewCamera returns a new Camera object.
//
// https://doc.babylonjs.com/api/classes/babylon.camera
func (b *Babylon) NewCamera(name string, position *Vector3, scene *Scene) *Camera {
	p := b.ctx.Get("Camera").New(name, position.JSObject(), scene.JSObject())
	return CameraFromJSObject(p)
}
*/
