// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ReflectBlock represents a babylon.js ReflectBlock.
// Block used to get the reflected vector from a direction and a normal
type ReflectBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (r *ReflectBlock) JSObject() js.Value { return r.p }

// ReflectBlock returns a ReflectBlock JavaScript class.
func (b *Babylon) ReflectBlock() *ReflectBlock {
	p := b.ctx.Get("ReflectBlock")
	return ReflectBlockFromJSObject(p)
}

// ReflectBlockFromJSObject returns a wrapped ReflectBlock JavaScript class.
func ReflectBlockFromJSObject(p js.Value) *ReflectBlock {
	return &ReflectBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewReflectBlock returns a new ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock
func (b *Babylon) NewReflectBlock(todo parameters) *ReflectBlock {
	p := b.ctx.Get("ReflectBlock").New(todo)
	return ReflectBlockFromJSObject(p)
}

// TODO: methods