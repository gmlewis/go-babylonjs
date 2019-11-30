// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VectorMergerBlock represents a babylon.js VectorMergerBlock.
// Block used to create a Vector2/3/4 out of individual inputs (one for each component)
type VectorMergerBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VectorMergerBlock) JSObject() js.Value { return v.p }

// VectorMergerBlock returns a VectorMergerBlock JavaScript class.
func (ba *Babylon) VectorMergerBlock() *VectorMergerBlock {
	p := ba.ctx.Get("VectorMergerBlock")
	return VectorMergerBlockFromJSObject(p, ba.ctx)
}

// VectorMergerBlockFromJSObject returns a wrapped VectorMergerBlock JavaScript class.
func VectorMergerBlockFromJSObject(p js.Value, ctx js.Value) *VectorMergerBlock {
	return &VectorMergerBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewVectorMergerBlock returns a new VectorMergerBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.vectormergerblock
func (ba *Babylon) NewVectorMergerBlock(name string) *VectorMergerBlock {
	p := ba.ctx.Get("VectorMergerBlock").New(name)
	return VectorMergerBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
