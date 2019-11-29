package babylon

import "syscall/js"

// Vector3 represents a babylon.js Vector3.
type Vector3 struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *Vector3) JSObject() js.Value { return s.p }

// Vector3 returns a Vector3 JavaScript class.
func (b *Babylon) Vector3() *Vector3 {
	p := b.ctx.Get("Vector3")
	return Vector3FromJSObject(p)
}

// Vector3FromJSObject returns a wrapped Vector3 JavaScript class.
func Vector3FromJSObject(p js.Value) *Vector3 {
	return &Vector3{p: p}
}

// NewVector3 returns a new Vector3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.Vector3
func (b *Babylon) NewVector3(x, y, z float64) *Vector3 {
	p := b.ctx.Get("Vector3").New(x, y, z)
	return Vector3FromJSObject(p)
}
