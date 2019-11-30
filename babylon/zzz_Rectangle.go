// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Rectangle represents a babylon.js Rectangle.
// Class used to create rectangle container
type Rectangle struct{ *Container }

// JSObject returns the underlying js.Value.
func (r *Rectangle) JSObject() js.Value { return r.p }

// Rectangle returns a Rectangle JavaScript class.
func (b *Babylon) Rectangle() *Rectangle {
	p := b.ctx.Get("Rectangle")
	return RectangleFromJSObject(p)
}

// RectangleFromJSObject returns a wrapped Rectangle JavaScript class.
func RectangleFromJSObject(p js.Value) *Rectangle {
	return &Rectangle{ContainerFromJSObject(p)}
}

// NewRectangleOpts contains optional parameters for NewRectangle.
type NewRectangleOpts struct {
	Name *string
}

// NewRectangle returns a new Rectangle object.
//
// https://doc.babylonjs.com/api/classes/babylon.rectangle
func (b *Babylon) NewRectangle(opts *NewRectangleOpts) *Rectangle {
	if opts == nil {
		opts = &NewRectangleOpts{}
	}

	p := b.ctx.Get("Rectangle").New(opts.Name)
	return RectangleFromJSObject(p)
}

// TODO: methods
