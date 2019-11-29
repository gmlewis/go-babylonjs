// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PositionNormalVertex represents a babylon.js PositionNormalVertex.
// Contains position and normal vectors for a vertex
type PositionNormalVertex struct{}

// JSObject returns the underlying js.Value.
func (p *PositionNormalVertex) JSObject() js.Value { return p.p }

// PositionNormalVertex returns a PositionNormalVertex JavaScript class.
func (b *Babylon) PositionNormalVertex() *PositionNormalVertex {
	p := b.ctx.Get("PositionNormalVertex")
	return PositionNormalVertexFromJSObject(p)
}

// PositionNormalVertexFromJSObject returns a wrapped PositionNormalVertex JavaScript class.
func PositionNormalVertexFromJSObject(p js.Value) *PositionNormalVertex {
	return &PositionNormalVertex{p: p}
}

// NewPositionNormalVertex returns a new PositionNormalVertex object.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormalvertex
func (b *Babylon) NewPositionNormalVertex(todo parameters) *PositionNormalVertex {
	p := b.ctx.Get("PositionNormalVertex").New(todo)
	return PositionNormalVertexFromJSObject(p)
}

// TODO: methods
