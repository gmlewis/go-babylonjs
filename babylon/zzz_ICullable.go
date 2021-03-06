// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ICullable represents a babylon.js ICullable.
// Interface for cullable objects
//
// See: https://doc.babylonjs.com/babylon101/materials#back-face-culling
type ICullable struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ICullable) JSObject() js.Value { return i.p }

// ICullable returns a ICullable JavaScript class.
func (ba *Babylon) ICullable() *ICullable {
	p := ba.ctx.Get("ICullable")
	return ICullableFromJSObject(p, ba.ctx)
}

// ICullableFromJSObject returns a wrapped ICullable JavaScript class.
func ICullableFromJSObject(p js.Value, ctx js.Value) *ICullable {
	return &ICullable{p: p, ctx: ctx}
}

// ICullableArrayToJSArray returns a JavaScript Array for the wrapped array.
func ICullableArrayToJSArray(array []*ICullable) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// IsCompletelyInFrustum calls the IsCompletelyInFrustum method on the ICullable object.
//
// https://doc.babylonjs.com/api/classes/babylon.icullable#iscompletelyinfrustum
func (i *ICullable) IsCompletelyInFrustum(frustumPlanes []*Plane) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	retVal := i.p.Call("isCompletelyInFrustum", args...)
	return retVal.Bool()
}

// IsInFrustum calls the IsInFrustum method on the ICullable object.
//
// https://doc.babylonjs.com/api/classes/babylon.icullable#isinfrustum
func (i *ICullable) IsInFrustum(frustumPlanes []*Plane) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	retVal := i.p.Call("isInFrustum", args...)
	return retVal.Bool()
}
