// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DomManagement represents a babylon.js DomManagement.
// Sets of helpers dealing with the DOM and some of the recurrent functions needed in
// Babylon.js
type DomManagement struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DomManagement) JSObject() js.Value { return d.p }

// DomManagement returns a DomManagement JavaScript class.
func (ba *Babylon) DomManagement() *DomManagement {
	p := ba.ctx.Get("DomManagement")
	return DomManagementFromJSObject(p, ba.ctx)
}

// DomManagementFromJSObject returns a wrapped DomManagement JavaScript class.
func DomManagementFromJSObject(p js.Value, ctx js.Value) *DomManagement {
	return &DomManagement{p: p, ctx: ctx}
}

// DomManagementArrayToJSArray returns a JavaScript Array for the wrapped array.
func DomManagementArrayToJSArray(array []*DomManagement) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// GetDOMTextContent calls the GetDOMTextContent method on the DomManagement object.
//
// https://doc.babylonjs.com/api/classes/babylon.dommanagement#getdomtextcontent
func (d *DomManagement) GetDOMTextContent(element js.Value) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	retVal := d.p.Call("GetDOMTextContent", args...)
	return retVal.String()
}

// IsNavigatorAvailable calls the IsNavigatorAvailable method on the DomManagement object.
//
// https://doc.babylonjs.com/api/classes/babylon.dommanagement#isnavigatoravailable
func (d *DomManagement) IsNavigatorAvailable() bool {

	retVal := d.p.Call("IsNavigatorAvailable")
	return retVal.Bool()
}

// IsWindowObjectExist calls the IsWindowObjectExist method on the DomManagement object.
//
// https://doc.babylonjs.com/api/classes/babylon.dommanagement#iswindowobjectexist
func (d *DomManagement) IsWindowObjectExist() bool {

	retVal := d.p.Call("IsWindowObjectExist")
	return retVal.Bool()
}
