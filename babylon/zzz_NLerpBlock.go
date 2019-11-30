// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NLerpBlock represents a babylon.js NLerpBlock.
// Block used to normalize lerp between 2 values
type NLerpBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (n *NLerpBlock) JSObject() js.Value { return n.p }

// NLerpBlock returns a NLerpBlock JavaScript class.
func (b *Babylon) NLerpBlock() *NLerpBlock {
	p := b.ctx.Get("NLerpBlock")
	return NLerpBlockFromJSObject(p)
}

// NLerpBlockFromJSObject returns a wrapped NLerpBlock JavaScript class.
func NLerpBlockFromJSObject(p js.Value) *NLerpBlock {
	return &NLerpBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewNLerpBlock returns a new NLerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.nlerpblock
func (b *Babylon) NewNLerpBlock(name string) *NLerpBlock {
	p := b.ctx.Get("NLerpBlock").New(name)
	return NLerpBlockFromJSObject(p)
}

// TODO: methods
