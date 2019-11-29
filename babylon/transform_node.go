package babylon

import "syscall/js"

// TransformNode represents a babylon.js TransformNode.
type TransformNode struct{ *Node }

// JSObject returns the underlying js.Value.
func (s *TransformNode) JSObject() js.Value { return s.p }

// TransformNode returns a TransformNode JavaScript class.
func (b *Babylon) TransformNode() *TransformNode {
	p := b.ctx.Get("TransformNode")
	return TransformNodeFromJSObject(p)
}

// TransformNodeFromJSObject returns a wrapped TransformNode JavaScript class.
func TransformNodeFromJSObject(p js.Value) *TransformNode {
	return &TransformNode{NodeFromJSObject(p)}
}

// NewTransformNode returns a new TransformNode object.
//
// https://doc.babylonjs.com/api/classes/babylon.node
func (b *Babylon) NewTransformNode(name string, scene *Scene) *TransformNode {
	p := b.ctx.Get("TransformNode").New(name, scene.JSObject())
	return TransformNodeFromJSObject(p)
}
