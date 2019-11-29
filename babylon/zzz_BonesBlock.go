// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BonesBlock represents a babylon.js BonesBlock.
// Block used to add support for vertex skinning (bones)
type BonesBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (b *BonesBlock) JSObject() js.Value { return b.p }

// BonesBlock returns a BonesBlock JavaScript class.
func (b *Babylon) BonesBlock() *BonesBlock {
	p := b.ctx.Get("BonesBlock")
	return BonesBlockFromJSObject(p)
}

// BonesBlockFromJSObject returns a wrapped BonesBlock JavaScript class.
func BonesBlockFromJSObject(p js.Value) *BonesBlock {
	return &BonesBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewBonesBlock returns a new BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock
func (b *Babylon) NewBonesBlock(todo parameters) *BonesBlock {
	p := b.ctx.Get("BonesBlock").New(todo)
	return BonesBlockFromJSObject(p)
}

// TODO: methods
