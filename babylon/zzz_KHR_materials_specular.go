// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_materials_specular represents a babylon.js KHR_materials_specular.
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/pull/1677&#34;&gt;Proposed Specification&lt;/a&gt;
// &lt;a href=&#34;https://www.babylonjs-playground.com/frame.html#BNIZX6#4&#34;&gt;Playground Sample&lt;/a&gt;
// !!! Experimental Extension Subject to Changes !!!
type KHR_materials_specular struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KHR_materials_specular) JSObject() js.Value { return k.p }

// KHR_materials_specular returns a KHR_materials_specular JavaScript class.
func (ba *Babylon) KHR_materials_specular() *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular")
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// KHR_materials_specularFromJSObject returns a wrapped KHR_materials_specular JavaScript class.
func KHR_materials_specularFromJSObject(p js.Value, ctx js.Value) *KHR_materials_specular {
	return &KHR_materials_specular{p: p, ctx: ctx}
}

// KHR_materials_specularArrayToJSArray returns a JavaScript Array for the wrapped array.
func KHR_materials_specularArrayToJSArray(array []*KHR_materials_specular) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Enabled returns the Enabled property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#enabled
func (k *KHR_materials_specular) Enabled(enabled bool) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(enabled)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// SetEnabled sets the Enabled property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#enabled
func (k *KHR_materials_specular) SetEnabled(enabled bool) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(enabled)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#name
func (k *KHR_materials_specular) Name(name string) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(name)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#name
func (k *KHR_materials_specular) SetName(name string) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(name)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// Order returns the Order property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#order
func (k *KHR_materials_specular) Order(order float64) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(order)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

// SetOrder sets the Order property of class KHR_materials_specular.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_specular#order
func (k *KHR_materials_specular) SetOrder(order float64) *KHR_materials_specular {
	p := ba.ctx.Get("KHR_materials_specular").New(order)
	return KHR_materials_specularFromJSObject(p, ba.ctx)
}

*/
