// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GUID represents a babylon.js GUID.
// Class used to manipulate GUIDs
type GUID struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GUID) JSObject() js.Value { return g.p }

// GUID returns a GUID JavaScript class.
func (gui *GUI) GUID() *GUID {
	p := gui.ctx.Get("GUID")
	return GUIDFromJSObject(p, gui.ctx)
}

// GUIDFromJSObject returns a wrapped GUID JavaScript class.
func GUIDFromJSObject(p js.Value, ctx js.Value) *GUID {
	return &GUID{p: p, ctx: ctx}
}

// GUIDArrayToJSArray returns a JavaScript Array for the wrapped array.
func GUIDArrayToJSArray(array []*GUID) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// RandomId calls the RandomId method on the GUID object.
//
// https://doc.babylonjs.com/api/classes/babylon.guid#randomid
func (g *GUID) RandomId() string {

	retVal := g.p.Call("RandomId")
	return retVal.String()
}
