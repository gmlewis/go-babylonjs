// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_texture_transform represents a babylon.js KHR_texture_transform.
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/blob/master/extensions/2.0/Khronos/KHR_texture_transform/README.md&#34;&gt;Specification&lt;/a&gt;
type KHR_texture_transform struct{}

// JSObject returns the underlying js.Value.
func (k *KHR_texture_transform) JSObject() js.Value { return k.p }

// KHR_texture_transform returns a KHR_texture_transform JavaScript class.
func (b *Babylon) KHR_texture_transform() *KHR_texture_transform {
	p := b.ctx.Get("KHR_texture_transform")
	return KHR_texture_transformFromJSObject(p)
}

// KHR_texture_transformFromJSObject returns a wrapped KHR_texture_transform JavaScript class.
func KHR_texture_transformFromJSObject(p js.Value) *KHR_texture_transform {
	return &KHR_texture_transform{p: p}
}

// NewKHR_texture_transform returns a new KHR_texture_transform object.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_texture_transform
func (b *Babylon) NewKHR_texture_transform(todo parameters) *KHR_texture_transform {
	p := b.ctx.Get("KHR_texture_transform").New(todo)
	return KHR_texture_transformFromJSObject(p)
}

// TODO: methods
