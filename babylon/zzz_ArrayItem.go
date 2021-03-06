// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArrayItem represents a babylon.js ArrayItem.
// Helper class for working with arrays when loading the glTF asset
type ArrayItem struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArrayItem) JSObject() js.Value { return a.p }

// ArrayItem returns a ArrayItem JavaScript class.
func (ba *Babylon) ArrayItem() *ArrayItem {
	p := ba.ctx.Get("ArrayItem")
	return ArrayItemFromJSObject(p, ba.ctx)
}

// ArrayItemFromJSObject returns a wrapped ArrayItem JavaScript class.
func ArrayItemFromJSObject(p js.Value, ctx js.Value) *ArrayItem {
	return &ArrayItem{p: p, ctx: ctx}
}

// ArrayItemArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArrayItemArrayToJSArray(array []*ArrayItem) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ArrayItemAssignOpts contains optional parameters for ArrayItem.Assign.
type ArrayItemAssignOpts struct {
	Array []*IArrayItem
}

// Assign calls the Assign method on the ArrayItem object.
//
// https://doc.babylonjs.com/api/classes/babylon.arrayitem#assign
func (a *ArrayItem) Assign(opts *ArrayItemAssignOpts) {
	if opts == nil {
		opts = &ArrayItemAssignOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Array == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, IArrayItemArrayToJSArray(opts.Array))
	}

	a.p.Call("Assign", args...)
}
