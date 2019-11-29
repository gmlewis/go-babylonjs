// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterialConnectionPoint represents a babylon.js NodeMaterialConnectionPoint.
// Defines a connection point for a block
type NodeMaterialConnectionPoint struct{}

// JSObject returns the underlying js.Value.
func (n *NodeMaterialConnectionPoint) JSObject() js.Value { return n.p }

// NodeMaterialConnectionPoint returns a NodeMaterialConnectionPoint JavaScript class.
func (b *Babylon) NodeMaterialConnectionPoint() *NodeMaterialConnectionPoint {
	p := b.ctx.Get("NodeMaterialConnectionPoint")
	return NodeMaterialConnectionPointFromJSObject(p)
}

// NodeMaterialConnectionPointFromJSObject returns a wrapped NodeMaterialConnectionPoint JavaScript class.
func NodeMaterialConnectionPointFromJSObject(p js.Value) *NodeMaterialConnectionPoint {
	return &NodeMaterialConnectionPoint{p: p}
}

// NewNodeMaterialConnectionPoint returns a new NodeMaterialConnectionPoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerialconnectionpoint
func (b *Babylon) NewNodeMaterialConnectionPoint(todo parameters) *NodeMaterialConnectionPoint {
	p := b.ctx.Get("NodeMaterialConnectionPoint").New(todo)
	return NodeMaterialConnectionPointFromJSObject(p)
}

// TODO: methods