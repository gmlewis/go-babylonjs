// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LerpBlock represents a babylon.js LerpBlock.
// Block used to lerp between 2 values
type LerpBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (l *LerpBlock) JSObject() js.Value { return l.p }

// LerpBlock returns a LerpBlock JavaScript class.
func (b *Babylon) LerpBlock() *LerpBlock {
	p := b.ctx.Get("LerpBlock")
	return LerpBlockFromJSObject(p)
}

// LerpBlockFromJSObject returns a wrapped LerpBlock JavaScript class.
func LerpBlockFromJSObject(p js.Value) *LerpBlock {
	return &LerpBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewLerpBlock returns a new LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock
func (b *Babylon) NewLerpBlock(todo parameters) *LerpBlock {
	p := b.ctx.Get("LerpBlock").New(todo)
	return LerpBlockFromJSObject(p)
}

// TODO: methods