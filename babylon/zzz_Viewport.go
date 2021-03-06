// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Viewport represents a babylon.js Viewport.
// Class used to represent a viewport on screen
type Viewport struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *Viewport) JSObject() js.Value { return v.p }

// Viewport returns a Viewport JavaScript class.
func (ba *Babylon) Viewport() *Viewport {
	p := ba.ctx.Get("Viewport")
	return ViewportFromJSObject(p, ba.ctx)
}

// ViewportFromJSObject returns a wrapped Viewport JavaScript class.
func ViewportFromJSObject(p js.Value, ctx js.Value) *Viewport {
	return &Viewport{p: p, ctx: ctx}
}

// ViewportArrayToJSArray returns a JavaScript Array for the wrapped array.
func ViewportArrayToJSArray(array []*Viewport) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewViewport returns a new Viewport object.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#constructor
func (ba *Babylon) NewViewport(x float64, y float64, width float64, height float64) *Viewport {

	args := make([]interface{}, 0, 4+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, width)
	args = append(args, height)

	p := ba.ctx.Get("Viewport").New(args...)
	return ViewportFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the Viewport object.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#clone
func (v *Viewport) Clone() *Viewport {

	retVal := v.p.Call("clone")
	return ViewportFromJSObject(retVal, v.ctx)
}

// ToGlobal calls the ToGlobal method on the Viewport object.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#toglobal
func (v *Viewport) ToGlobal(renderWidth float64, renderHeight float64) *Viewport {

	args := make([]interface{}, 0, 2+0)

	args = append(args, renderWidth)

	args = append(args, renderHeight)

	retVal := v.p.Call("toGlobal", args...)
	return ViewportFromJSObject(retVal, v.ctx)
}

// ToGlobalToRef calls the ToGlobalToRef method on the Viewport object.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#toglobaltoref
func (v *Viewport) ToGlobalToRef(renderWidth float64, renderHeight float64, ref *Viewport) *Viewport {

	args := make([]interface{}, 0, 3+0)

	args = append(args, renderWidth)

	args = append(args, renderHeight)

	if ref == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ref.JSObject())
	}

	retVal := v.p.Call("toGlobalToRef", args...)
	return ViewportFromJSObject(retVal, v.ctx)
}

// Height returns the Height property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#height
func (v *Viewport) Height() float64 {
	retVal := v.p.Get("height")
	return retVal.Float()
}

// SetHeight sets the Height property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#height
func (v *Viewport) SetHeight(height float64) *Viewport {
	v.p.Set("height", height)
	return v
}

// Width returns the Width property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#width
func (v *Viewport) Width() float64 {
	retVal := v.p.Get("width")
	return retVal.Float()
}

// SetWidth sets the Width property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#width
func (v *Viewport) SetWidth(width float64) *Viewport {
	v.p.Set("width", width)
	return v
}

// X returns the X property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#x
func (v *Viewport) X() float64 {
	retVal := v.p.Get("x")
	return retVal.Float()
}

// SetX sets the X property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#x
func (v *Viewport) SetX(x float64) *Viewport {
	v.p.Set("x", x)
	return v
}

// Y returns the Y property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#y
func (v *Viewport) Y() float64 {
	retVal := v.p.Get("y")
	return retVal.Float()
}

// SetY sets the Y property of class Viewport.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport#y
func (v *Viewport) SetY(y float64) *Viewport {
	v.p.Set("y", y)
	return v
}
