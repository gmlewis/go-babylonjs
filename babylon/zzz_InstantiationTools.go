// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InstantiationTools represents a babylon.js InstantiationTools.
// Class used to enable instatition of objects by class name
type InstantiationTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InstantiationTools) JSObject() js.Value { return i.p }

// InstantiationTools returns a InstantiationTools JavaScript class.
func (ba *Babylon) InstantiationTools() *InstantiationTools {
	p := ba.ctx.Get("InstantiationTools")
	return InstantiationToolsFromJSObject(p, ba.ctx)
}

// InstantiationToolsFromJSObject returns a wrapped InstantiationTools JavaScript class.
func InstantiationToolsFromJSObject(p js.Value, ctx js.Value) *InstantiationTools {
	return &InstantiationTools{p: p, ctx: ctx}
}

// InstantiationToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func InstantiationToolsArrayToJSArray(array []*InstantiationTools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Instantiate calls the Instantiate method on the InstantiationTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiationtools#instantiate
func (i *InstantiationTools) Instantiate(className string) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, className)

	retVal := i.p.Call("Instantiate", args...)
	return retVal
}

/*

// RegisteredExternalClasses returns the RegisteredExternalClasses property of class InstantiationTools.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiationtools#registeredexternalclasses
func (i *InstantiationTools) RegisteredExternalClasses(RegisteredExternalClasses js.Value) *InstantiationTools {
	p := ba.ctx.Get("InstantiationTools").New(RegisteredExternalClasses)
	return InstantiationToolsFromJSObject(p, ba.ctx)
}

// SetRegisteredExternalClasses sets the RegisteredExternalClasses property of class InstantiationTools.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiationtools#registeredexternalclasses
func (i *InstantiationTools) SetRegisteredExternalClasses(RegisteredExternalClasses js.Value) *InstantiationTools {
	p := ba.ctx.Get("InstantiationTools").New(RegisteredExternalClasses)
	return InstantiationToolsFromJSObject(p, ba.ctx)
}

*/
