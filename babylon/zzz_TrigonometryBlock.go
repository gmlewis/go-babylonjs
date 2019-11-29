// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TrigonometryBlock represents a babylon.js TrigonometryBlock.
// Block used to apply trigonometry operation to floats
type TrigonometryBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (t *TrigonometryBlock) JSObject() js.Value { return t.p }

// TrigonometryBlock returns a TrigonometryBlock JavaScript class.
func (b *Babylon) TrigonometryBlock() *TrigonometryBlock {
	p := b.ctx.Get("TrigonometryBlock")
	return TrigonometryBlockFromJSObject(p)
}

// TrigonometryBlockFromJSObject returns a wrapped TrigonometryBlock JavaScript class.
func TrigonometryBlockFromJSObject(p js.Value) *TrigonometryBlock {
	return &TrigonometryBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewTrigonometryBlock returns a new TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock
func (b *Babylon) NewTrigonometryBlock(todo parameters) *TrigonometryBlock {
	p := b.ctx.Get("TrigonometryBlock").New(todo)
	return TrigonometryBlockFromJSObject(p)
}

// TODO: methods
