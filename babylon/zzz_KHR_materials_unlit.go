// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_materials_unlit represents a babylon.js KHR_materials_unlit.
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/tree/master/extensions/2.0/Khronos/KHR_materials_unlit&#34;&gt;Specification&lt;/a&gt;
type KHR_materials_unlit struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (k *KHR_materials_unlit) JSObject() js.Value { return k.p }

// KHR_materials_unlit returns a KHR_materials_unlit JavaScript class.
func (b *Babylon) KHR_materials_unlit() *KHR_materials_unlit {
	p := b.ctx.Get("KHR_materials_unlit")
	return KHR_materials_unlitFromJSObject(p)
}

// KHR_materials_unlitFromJSObject returns a wrapped KHR_materials_unlit JavaScript class.
func KHR_materials_unlitFromJSObject(p js.Value) *KHR_materials_unlit {
	return &KHR_materials_unlit{p: p}
}

// TODO: methods
