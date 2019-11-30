// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IcoSphereBuilder represents a babylon.js IcoSphereBuilder.
// Class containing static functions to help procedurally build meshes
type IcoSphereBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IcoSphereBuilder) JSObject() js.Value { return i.p }

// IcoSphereBuilder returns a IcoSphereBuilder JavaScript class.
func (ba *Babylon) IcoSphereBuilder() *IcoSphereBuilder {
	p := ba.ctx.Get("IcoSphereBuilder")
	return IcoSphereBuilderFromJSObject(p, ba.ctx)
}

// IcoSphereBuilderFromJSObject returns a wrapped IcoSphereBuilder JavaScript class.
func IcoSphereBuilderFromJSObject(p js.Value, ctx js.Value) *IcoSphereBuilder {
	return &IcoSphereBuilder{p: p, ctx: ctx}
}

// TODO: methods
