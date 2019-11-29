// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ClampBlock represents a babylon.js ClampBlock.
// Block used to clamp a float
type ClampBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (c *ClampBlock) JSObject() js.Value { return c.p }

// ClampBlock returns a ClampBlock JavaScript class.
func (b *Babylon) ClampBlock() *ClampBlock {
	p := b.ctx.Get("ClampBlock")
	return ClampBlockFromJSObject(p)
}

// ClampBlockFromJSObject returns a wrapped ClampBlock JavaScript class.
func ClampBlockFromJSObject(p js.Value) *ClampBlock {
	return &ClampBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewClampBlock returns a new ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock
func (b *Babylon) NewClampBlock(todo parameters) *ClampBlock {
	p := b.ctx.Get("ClampBlock").New(todo)
	return ClampBlockFromJSObject(p)
}

// TODO: methods