// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShapeBuilder represents a babylon.js ShapeBuilder.
// Class containing static functions to help procedurally build meshes
type ShapeBuilder struct{}

// JSObject returns the underlying js.Value.
func (s *ShapeBuilder) JSObject() js.Value { return s.p }

// ShapeBuilder returns a ShapeBuilder JavaScript class.
func (b *Babylon) ShapeBuilder() *ShapeBuilder {
	p := b.ctx.Get("ShapeBuilder")
	return ShapeBuilderFromJSObject(p)
}

// ShapeBuilderFromJSObject returns a wrapped ShapeBuilder JavaScript class.
func ShapeBuilderFromJSObject(p js.Value) *ShapeBuilder {
	return &ShapeBuilder{p: p}
}

// NewShapeBuilder returns a new ShapeBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.shapebuilder
func (b *Babylon) NewShapeBuilder(todo parameters) *ShapeBuilder {
	p := b.ctx.Get("ShapeBuilder").New(todo)
	return ShapeBuilderFromJSObject(p)
}

// TODO: methods
