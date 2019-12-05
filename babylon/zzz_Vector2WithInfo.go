// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Vector2WithInfo represents a babylon.js Vector2WithInfo.
// Class used to transport BABYLON.Vector2 information for pointer events
type Vector2WithInfo struct {
	*Vector2
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *Vector2WithInfo) JSObject() js.Value { return v.p }

// Vector2WithInfo returns a Vector2WithInfo JavaScript class.
func (ba *Babylon) Vector2WithInfo() *Vector2WithInfo {
	p := ba.ctx.Get("Vector2WithInfo")
	return Vector2WithInfoFromJSObject(p, ba.ctx)
}

// Vector2WithInfoFromJSObject returns a wrapped Vector2WithInfo JavaScript class.
func Vector2WithInfoFromJSObject(p js.Value, ctx js.Value) *Vector2WithInfo {
	return &Vector2WithInfo{Vector2: Vector2FromJSObject(p, ctx), ctx: ctx}
}

// Vector2WithInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func Vector2WithInfoArrayToJSArray(array []*Vector2WithInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVector2WithInfoOpts contains optional parameters for NewVector2WithInfo.
type NewVector2WithInfoOpts struct {
	ButtonIndex *float64
}

// NewVector2WithInfo returns a new Vector2WithInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2withinfo
func (gui *GUI) NewVector2WithInfo(source *Vector2, opts *NewVector2WithInfoOpts) *Vector2WithInfo {
	if opts == nil {
		opts = &NewVector2WithInfoOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, source.JSObject())

	if opts.ButtonIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ButtonIndex)
	}

	p := gui.ctx.Get("Vector2WithInfo").New(args...)
	return Vector2WithInfoFromJSObject(p, gui.ctx)
}

// ButtonIndex returns the ButtonIndex property of class Vector2WithInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2withinfo#buttonindex
func (v *Vector2WithInfo) ButtonIndex() float64 {
	retVal := v.p.Get("buttonIndex")
	return retVal.Float()
}

// SetButtonIndex sets the ButtonIndex property of class Vector2WithInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2withinfo#buttonindex
func (v *Vector2WithInfo) SetButtonIndex(buttonIndex float64) *Vector2WithInfo {
	v.p.Set("buttonIndex", buttonIndex)
	return v
}
