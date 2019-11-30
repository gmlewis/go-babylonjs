// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterialBlock represents a babylon.js NodeMaterialBlock.
// Defines a block that can be used inside a node based material
type NodeMaterialBlock struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NodeMaterialBlock) JSObject() js.Value { return n.p }

// NodeMaterialBlock returns a NodeMaterialBlock JavaScript class.
func (ba *Babylon) NodeMaterialBlock() *NodeMaterialBlock {
	p := ba.ctx.Get("NodeMaterialBlock")
	return NodeMaterialBlockFromJSObject(p, ba.ctx)
}

// NodeMaterialBlockFromJSObject returns a wrapped NodeMaterialBlock JavaScript class.
func NodeMaterialBlockFromJSObject(p js.Value, ctx js.Value) *NodeMaterialBlock {
	return &NodeMaterialBlock{p: p, ctx: ctx}
}

// NewNodeMaterialBlockOpts contains optional parameters for NewNodeMaterialBlock.
type NewNodeMaterialBlockOpts struct {
	Target *JSValue

	IsFinalMerger *JSBool

	IsInput *JSBool
}

// NewNodeMaterialBlock returns a new NodeMaterialBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerialblock
func (ba *Babylon) NewNodeMaterialBlock(name string, opts *NewNodeMaterialBlockOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &NewNodeMaterialBlockOpts{}
	}

	p := ba.ctx.Get("NodeMaterialBlock").New(name, opts.Target.JSObject(), opts.IsFinalMerger.JSObject(), opts.IsInput.JSObject())
	return NodeMaterialBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
