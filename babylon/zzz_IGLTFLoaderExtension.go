// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IGLTFLoaderExtension represents a babylon.js IGLTFLoaderExtension.
// Interface for extending the loader.
type IGLTFLoaderExtension struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IGLTFLoaderExtension) JSObject() js.Value { return i.p }

// IGLTFLoaderExtension returns a IGLTFLoaderExtension JavaScript class.
func (ba *Babylon) IGLTFLoaderExtension() *IGLTFLoaderExtension {
	p := ba.ctx.Get("IGLTFLoaderExtension")
	return IGLTFLoaderExtensionFromJSObject(p, ba.ctx)
}

// IGLTFLoaderExtensionFromJSObject returns a wrapped IGLTFLoaderExtension JavaScript class.
func IGLTFLoaderExtensionFromJSObject(p js.Value, ctx js.Value) *IGLTFLoaderExtension {
	return &IGLTFLoaderExtension{p: p, ctx: ctx}
}

// IGLTFLoaderExtensionArrayToJSArray returns a JavaScript Array for the wrapped array.
func IGLTFLoaderExtensionArrayToJSArray(array []*IGLTFLoaderExtension) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Enabled returns the Enabled property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#enabled
func (i *IGLTFLoaderExtension) Enabled() bool {
	retVal := i.p.Get("enabled")
	return retVal.Bool()
}

// SetEnabled sets the Enabled property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#enabled
func (i *IGLTFLoaderExtension) SetEnabled(enabled bool) *IGLTFLoaderExtension {
	i.p.Set("enabled", enabled)
	return i
}

// Name returns the Name property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#name
func (i *IGLTFLoaderExtension) Name() string {
	retVal := i.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#name
func (i *IGLTFLoaderExtension) SetName(name string) *IGLTFLoaderExtension {
	i.p.Set("name", name)
	return i
}

// Order returns the Order property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#order
func (i *IGLTFLoaderExtension) Order() float64 {
	retVal := i.p.Get("order")
	return retVal.Float()
}

// SetOrder sets the Order property of class IGLTFLoaderExtension.
//
// https://doc.babylonjs.com/api/classes/babylon.igltfloaderextension#order
func (i *IGLTFLoaderExtension) SetOrder(order float64) *IGLTFLoaderExtension {
	i.p.Set("order", order)
	return i
}
