// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TransformNode represents a babylon.js TransformNode.
// A TransformNode is an object that is not rendered but can be used as a center of transformation. This can decrease memory usage and increase rendering speed compared to using an empty mesh as a parent and is less complicated than using a pivot matrix.

//
// See: https://doc.babylonjs.com/how_to/transformnode
type TransformNode struct{ *Node }

// JSObject returns the underlying js.Value.
func (t *TransformNode) JSObject() js.Value { return t.p }

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
// https://doc.babylonjs.com/api/classes/babylon.transformnode
func (b *Babylon) NewTransformNode(todo parameters) *TransformNode {
	p := b.ctx.Get("TransformNode").New(todo)
	return TransformNodeFromJSObject(p)
}

// TODO: methods
