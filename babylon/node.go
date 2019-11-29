package babylon

import "syscall/js"

// Node represents a babylon.js Node.
type Node struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *Node) JSObject() js.Value { return s.p }

// Node returns a Node JavaScript class.
func (b *Babylon) Node() *Node {
	p := b.ctx.Get("Node")
	return NodeFromJSObject(p)
}

// NodeFromJSObject returns a wrapped Node JavaScript class.
func NodeFromJSObject(p js.Value) *Node {
	return &Node{p: p}
}

// NewNode returns a new Node object.
//
// https://doc.babylonjs.com/api/classes/babylon.node
func (b *Babylon) NewNode(name string, scene *Scene) *Node {
	p := b.ctx.Get("Node").New(name, scene.JSObject())
	return NodeFromJSObject(p)
}
