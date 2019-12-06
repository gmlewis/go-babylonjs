// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Rectangle represents a babylon.js Rectangle.
// Class used to create rectangle container
type Rectangle struct {
	*Container
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *Rectangle) JSObject() js.Value { return r.p }

// Rectangle returns a Rectangle JavaScript class.
func (gui *GUI) Rectangle() *Rectangle {
	p := gui.ctx.Get("Rectangle")
	return RectangleFromJSObject(p, gui.ctx)
}

// RectangleFromJSObject returns a wrapped Rectangle JavaScript class.
func RectangleFromJSObject(p js.Value, ctx js.Value) *Rectangle {
	return &Rectangle{Container: ContainerFromJSObject(p, ctx), ctx: ctx}
}

// RectangleArrayToJSArray returns a JavaScript Array for the wrapped array.
func RectangleArrayToJSArray(array []*Rectangle) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRectangleOpts contains optional parameters for NewRectangle.
type NewRectangleOpts struct {
	Name *string
}

// NewRectangle returns a new Rectangle object.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle
func (gui *GUI) NewRectangle(opts *NewRectangleOpts) *Rectangle {
	if opts == nil {
		opts = &NewRectangleOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("Rectangle").New(args...)
	return RectangleFromJSObject(p, gui.ctx)
}

// CornerRadius returns the CornerRadius property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#cornerradius
func (r *Rectangle) CornerRadius() float64 {
	retVal := r.p.Get("cornerRadius")
	return retVal.Float()
}

// SetCornerRadius sets the CornerRadius property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#cornerradius
func (r *Rectangle) SetCornerRadius(cornerRadius float64) *Rectangle {
	r.p.Set("cornerRadius", cornerRadius)
	return r
}

// Name returns the Name property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#name
func (r *Rectangle) Name() string {
	retVal := r.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#name
func (r *Rectangle) SetName(name string) *Rectangle {
	r.p.Set("name", name)
	return r
}

// Thickness returns the Thickness property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#thickness
func (r *Rectangle) Thickness() float64 {
	retVal := r.p.Get("thickness")
	return retVal.Float()
}

// SetThickness sets the Thickness property of class Rectangle.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle#thickness
func (r *Rectangle) SetThickness(thickness float64) *Rectangle {
	r.p.Set("thickness", thickness)
	return r
}
