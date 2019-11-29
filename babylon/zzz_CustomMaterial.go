// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CustomMaterial represents a babylon.js CustomMaterial.
//
type CustomMaterial struct{ *StandardMaterial }

// JSObject returns the underlying js.Value.
func (c *CustomMaterial) JSObject() js.Value { return c.p }

// CustomMaterial returns a CustomMaterial JavaScript class.
func (b *Babylon) CustomMaterial() *CustomMaterial {
	p := b.ctx.Get("CustomMaterial")
	return CustomMaterialFromJSObject(p)
}

// CustomMaterialFromJSObject returns a wrapped CustomMaterial JavaScript class.
func CustomMaterialFromJSObject(p js.Value) *CustomMaterial {
	return &CustomMaterial{StandardMaterialFromJSObject(p)}
}

// NewCustomMaterial returns a new CustomMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.custommaterial
func (b *Babylon) NewCustomMaterial(todo parameters) *CustomMaterial {
	p := b.ctx.Get("CustomMaterial").New(todo)
	return CustomMaterialFromJSObject(p)
}

// TODO: methods
