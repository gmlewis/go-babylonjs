// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// INode represents a babylon.js INode.
// Loader interface with additional members.
type INode struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *INode) JSObject() js.Value { return i.p }

// INode returns a INode JavaScript class.
func (ba *Babylon) INode() *INode {
	p := ba.ctx.Get("INode")
	return INodeFromJSObject(p, ba.ctx)
}

// INodeFromJSObject returns a wrapped INode JavaScript class.
func INodeFromJSObject(p js.Value, ctx js.Value) *INode {
	return &INode{p: p, ctx: ctx}
}

// INodeArrayToJSArray returns a JavaScript Array for the wrapped array.
func INodeArrayToJSArray(array []*INode) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Parent returns the Parent property of class INode.
//
// https://doc.babylonjs.com/api/classes/babylon.inode#parent
func (i *INode) Parent(parent *INode) *INode {
	p := ba.ctx.Get("INode").New(parent.JSObject())
	return INodeFromJSObject(p, ba.ctx)
}

// SetParent sets the Parent property of class INode.
//
// https://doc.babylonjs.com/api/classes/babylon.inode#parent
func (i *INode) SetParent(parent *INode) *INode {
	p := ba.ctx.Get("INode").New(parent.JSObject())
	return INodeFromJSObject(p, ba.ctx)
}

*/
