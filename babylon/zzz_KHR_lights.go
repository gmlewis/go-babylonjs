// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_lights represents a babylon.js KHR_lights.
// <a href="https://github.com/KhronosGroup/glTF/blob/master/extensions/2.0/Khronos/KHR_lights_punctual/README.md">Specification</a>
type KHR_lights struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KHR_lights) JSObject() js.Value { return k.p }

// KHR_lights returns a KHR_lights JavaScript class.
func (ba *Babylon) KHR_lights() *KHR_lights {
	p := ba.ctx.Get("KHR_lights")
	return KHR_lightsFromJSObject(p, ba.ctx)
}

// KHR_lightsFromJSObject returns a wrapped KHR_lights JavaScript class.
func KHR_lightsFromJSObject(p js.Value, ctx js.Value) *KHR_lights {
	return &KHR_lights{p: p, ctx: ctx}
}

// KHR_lightsArrayToJSArray returns a JavaScript Array for the wrapped array.
func KHR_lightsArrayToJSArray(array []*KHR_lights) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Enabled returns the Enabled property of class KHR_lights.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_lights#enabled
func (k *KHR_lights) Enabled() bool {
	retVal := k.p.Get("enabled")
	return retVal.Bool()
}

// SetEnabled sets the Enabled property of class KHR_lights.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_lights#enabled
func (k *KHR_lights) SetEnabled(enabled bool) *KHR_lights {
	k.p.Set("enabled", enabled)
	return k
}

// Name returns the Name property of class KHR_lights.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_lights#name
func (k *KHR_lights) Name() string {
	retVal := k.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class KHR_lights.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_lights#name
func (k *KHR_lights) SetName(name string) *KHR_lights {
	k.p.Set("name", name)
	return k
}
