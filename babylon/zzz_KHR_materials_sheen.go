// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_materials_sheen represents a babylon.js KHR_materials_sheen.
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/pull/1688&#34;&gt;Proposed Specification&lt;/a&gt;
// &lt;a href=&#34;https://www.babylonjs-playground.com/frame.html#BNIZX6#4&#34;&gt;Playground Sample&lt;/a&gt;
// !!! Experimental Extension Subject to Changes !!!
type KHR_materials_sheen struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KHR_materials_sheen) JSObject() js.Value { return k.p }

// KHR_materials_sheen returns a KHR_materials_sheen JavaScript class.
func (ba *Babylon) KHR_materials_sheen() *KHR_materials_sheen {
	p := ba.ctx.Get("KHR_materials_sheen")
	return KHR_materials_sheenFromJSObject(p, ba.ctx)
}

// KHR_materials_sheenFromJSObject returns a wrapped KHR_materials_sheen JavaScript class.
func KHR_materials_sheenFromJSObject(p js.Value, ctx js.Value) *KHR_materials_sheen {
	return &KHR_materials_sheen{p: p, ctx: ctx}
}

// TODO: methods
