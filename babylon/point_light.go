package babylon

import "syscall/js"

// PointLight represents a babylon.js PointLight.
type PointLight struct{ *ShadowLight }

// JSObject returns the underlying js.Value.
func (s *PointLight) JSObject() js.Value { return s.p }

// PointLight returns a PointLight JavaScript class.
func (b *Babylon) PointLight() *PointLight {
	p := b.ctx.Get("PointLight")
	return PointLightFromJSObject(p)
}

// PointLightFromJSObject returns a wrapped PointLight JavaScript class.
func PointLightFromJSObject(p js.Value) *PointLight {
	return &PointLight{ShadowLightFromJSObject(p)}
}

// NewPointLight returns a new PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.Pointlight
func (b *Babylon) NewPointLight(name string, position *Vector3, scene *Scene) *PointLight {
	p := b.ctx.Get("PointLight").New(name, position.JSObject(), scene.JSObject())
	return PointLightFromJSObject(p)
}
