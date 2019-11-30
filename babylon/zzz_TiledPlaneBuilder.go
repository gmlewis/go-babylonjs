// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TiledPlaneBuilder represents a babylon.js TiledPlaneBuilder.
// Class containing static functions to help procedurally build meshes
type TiledPlaneBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TiledPlaneBuilder) JSObject() js.Value { return t.p }

// TiledPlaneBuilder returns a TiledPlaneBuilder JavaScript class.
func (ba *Babylon) TiledPlaneBuilder() *TiledPlaneBuilder {
	p := ba.ctx.Get("TiledPlaneBuilder")
	return TiledPlaneBuilderFromJSObject(p, ba.ctx)
}

// TiledPlaneBuilderFromJSObject returns a wrapped TiledPlaneBuilder JavaScript class.
func TiledPlaneBuilderFromJSObject(p js.Value, ctx js.Value) *TiledPlaneBuilder {
	return &TiledPlaneBuilder{p: p, ctx: ctx}
}

// TODO: methods
