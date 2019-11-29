// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterial represents a babylon.js NodeMaterial.
// Class used to create a node based material built by assembling shader blocks
type NodeMaterial struct{}

// JSObject returns the underlying js.Value.
func (n *NodeMaterial) JSObject() js.Value { return n.p }

// NodeMaterial returns a NodeMaterial JavaScript class.
func (b *Babylon) NodeMaterial() *NodeMaterial {
	p := b.ctx.Get("NodeMaterial")
	return NodeMaterialFromJSObject(p)
}

// NodeMaterialFromJSObject returns a wrapped NodeMaterial JavaScript class.
func NodeMaterialFromJSObject(p js.Value) *NodeMaterial {
	return &NodeMaterial{p: p}
}

// NewNodeMaterial returns a new NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial
func (b *Babylon) NewNodeMaterial(todo parameters) *NodeMaterial {
	p := b.ctx.Get("NodeMaterial").New(todo)
	return NodeMaterialFromJSObject(p)
}

// TODO: methods
