// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VertexOutputBlock represents a babylon.js VertexOutputBlock.
// Block used to output the vertex position
type VertexOutputBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VertexOutputBlock) JSObject() js.Value { return v.p }

// VertexOutputBlock returns a VertexOutputBlock JavaScript class.
func (ba *Babylon) VertexOutputBlock() *VertexOutputBlock {
	p := ba.ctx.Get("VertexOutputBlock")
	return VertexOutputBlockFromJSObject(p, ba.ctx)
}

// VertexOutputBlockFromJSObject returns a wrapped VertexOutputBlock JavaScript class.
func VertexOutputBlockFromJSObject(p js.Value, ctx js.Value) *VertexOutputBlock {
	return &VertexOutputBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewVertexOutputBlock returns a new VertexOutputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexoutputblock
func (ba *Babylon) NewVertexOutputBlock(name string) *VertexOutputBlock {
	p := ba.ctx.Get("VertexOutputBlock").New(name)
	return VertexOutputBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
