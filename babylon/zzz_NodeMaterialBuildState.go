// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterialBuildState represents a babylon.js NodeMaterialBuildState.
// Class used to store node based material build state
type NodeMaterialBuildState struct{}

// JSObject returns the underlying js.Value.
func (n *NodeMaterialBuildState) JSObject() js.Value { return n.p }

// NodeMaterialBuildState returns a NodeMaterialBuildState JavaScript class.
func (b *Babylon) NodeMaterialBuildState() *NodeMaterialBuildState {
	p := b.ctx.Get("NodeMaterialBuildState")
	return NodeMaterialBuildStateFromJSObject(p)
}

// NodeMaterialBuildStateFromJSObject returns a wrapped NodeMaterialBuildState JavaScript class.
func NodeMaterialBuildStateFromJSObject(p js.Value) *NodeMaterialBuildState {
	return &NodeMaterialBuildState{p: p}
}

// NewNodeMaterialBuildState returns a new NodeMaterialBuildState object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerialbuildstate
func (b *Babylon) NewNodeMaterialBuildState(todo parameters) *NodeMaterialBuildState {
	p := b.ctx.Get("NodeMaterialBuildState").New(todo)
	return NodeMaterialBuildStateFromJSObject(p)
}

// TODO: methods