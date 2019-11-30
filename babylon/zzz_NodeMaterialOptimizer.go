// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterialOptimizer represents a babylon.js NodeMaterialOptimizer.
// Root class for all node material optimizers
type NodeMaterialOptimizer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (n *NodeMaterialOptimizer) JSObject() js.Value { return n.p }

// NodeMaterialOptimizer returns a NodeMaterialOptimizer JavaScript class.
func (b *Babylon) NodeMaterialOptimizer() *NodeMaterialOptimizer {
	p := b.ctx.Get("NodeMaterialOptimizer")
	return NodeMaterialOptimizerFromJSObject(p)
}

// NodeMaterialOptimizerFromJSObject returns a wrapped NodeMaterialOptimizer JavaScript class.
func NodeMaterialOptimizerFromJSObject(p js.Value) *NodeMaterialOptimizer {
	return &NodeMaterialOptimizer{p: p}
}

// TODO: methods
