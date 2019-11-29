// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DerivativeBlock represents a babylon.js DerivativeBlock.
// Block used to get the derivative value on x and y of a given input
type DerivativeBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (d *DerivativeBlock) JSObject() js.Value { return d.p }

// DerivativeBlock returns a DerivativeBlock JavaScript class.
func (b *Babylon) DerivativeBlock() *DerivativeBlock {
	p := b.ctx.Get("DerivativeBlock")
	return DerivativeBlockFromJSObject(p)
}

// DerivativeBlockFromJSObject returns a wrapped DerivativeBlock JavaScript class.
func DerivativeBlockFromJSObject(p js.Value) *DerivativeBlock {
	return &DerivativeBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewDerivativeBlock returns a new DerivativeBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.derivativeblock
func (b *Babylon) NewDerivativeBlock(todo parameters) *DerivativeBlock {
	p := b.ctx.Get("DerivativeBlock").New(todo)
	return DerivativeBlockFromJSObject(p)
}

// TODO: methods
