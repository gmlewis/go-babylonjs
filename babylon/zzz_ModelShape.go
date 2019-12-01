// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ModelShape represents a babylon.js ModelShape.
// Represents the shape of the model used by one particle of a solid particle system.
// SPS internal tool, don&amp;#39;t use it manually.
type ModelShape struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *ModelShape) JSObject() js.Value { return m.p }

// ModelShape returns a ModelShape JavaScript class.
func (ba *Babylon) ModelShape() *ModelShape {
	p := ba.ctx.Get("ModelShape")
	return ModelShapeFromJSObject(p, ba.ctx)
}

// ModelShapeFromJSObject returns a wrapped ModelShape JavaScript class.
func ModelShapeFromJSObject(p js.Value, ctx js.Value) *ModelShape {
	return &ModelShape{p: p, ctx: ctx}
}

// ModelShapeArrayToJSArray returns a JavaScript Array for the wrapped array.
func ModelShapeArrayToJSArray(array []*ModelShape) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

 */
