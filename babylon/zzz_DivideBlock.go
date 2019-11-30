// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DivideBlock represents a babylon.js DivideBlock.
// Block used to divide 2 vectors
type DivideBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DivideBlock) JSObject() js.Value { return d.p }

// DivideBlock returns a DivideBlock JavaScript class.
func (ba *Babylon) DivideBlock() *DivideBlock {
	p := ba.ctx.Get("DivideBlock")
	return DivideBlockFromJSObject(p, ba.ctx)
}

// DivideBlockFromJSObject returns a wrapped DivideBlock JavaScript class.
func DivideBlockFromJSObject(p js.Value, ctx js.Value) *DivideBlock {
	return &DivideBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewDivideBlock returns a new DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock
func (ba *Babylon) NewDivideBlock(name string) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(name)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
