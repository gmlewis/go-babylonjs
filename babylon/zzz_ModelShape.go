// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ModelShape represents a babylon.js ModelShape.
// Represents the shape of the model used by one particle of a solid particle system.
// SPS internal tool, don&amp;#39;t use it manually.
type ModelShape struct{}

// JSObject returns the underlying js.Value.
func (m *ModelShape) JSObject() js.Value { return m.p }

// ModelShape returns a ModelShape JavaScript class.
func (b *Babylon) ModelShape() *ModelShape {
	p := b.ctx.Get("ModelShape")
	return ModelShapeFromJSObject(p)
}

// ModelShapeFromJSObject returns a wrapped ModelShape JavaScript class.
func ModelShapeFromJSObject(p js.Value) *ModelShape {
	return &ModelShape{p: p}
}

// NewModelShape returns a new ModelShape object.
//
// https://doc.babylonjs.com/api/classes/babylon.modelshape
func (b *Babylon) NewModelShape(todo parameters) *ModelShape {
	p := b.ctx.Get("ModelShape").New(todo)
	return ModelShapeFromJSObject(p)
}

// TODO: methods
