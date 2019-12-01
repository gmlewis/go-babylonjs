// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TGATools represents a babylon.js TGATools.
// Based on jsTGALoader - Javascript loader for TGA file
// By Vincent Thibault
//
// See: http://blog.robrowser.com/javascript-tga-loader.html
type TGATools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TGATools) JSObject() js.Value { return t.p }

// TGATools returns a TGATools JavaScript class.
func (ba *Babylon) TGATools() *TGATools {
	p := ba.ctx.Get("TGATools")
	return TGAToolsFromJSObject(p, ba.ctx)
}

// TGAToolsFromJSObject returns a wrapped TGATools JavaScript class.
func TGAToolsFromJSObject(p js.Value, ctx js.Value) *TGATools {
	return &TGATools{p: p, ctx: ctx}
}

// TGAToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func TGAToolsArrayToJSArray(array []*TGATools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// GetTGAHeader calls the GetTGAHeader method on the TGATools object.
//
// https://doc.babylonjs.com/api/classes/babylon.tgatools#gettgaheader
func (t *TGATools) GetTGAHeader(data js.Value) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := t.p.Call("GetTGAHeader", args...)
	return retVal
}

/*

 */
