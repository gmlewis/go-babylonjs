// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NegateBlock represents a babylon.js NegateBlock.
// Block used to get negative version of a value (i.e. x * -1)
type NegateBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NegateBlock) JSObject() js.Value { return n.p }

// NegateBlock returns a NegateBlock JavaScript class.
func (ba *Babylon) NegateBlock() *NegateBlock {
	p := ba.ctx.Get("NegateBlock")
	return NegateBlockFromJSObject(p, ba.ctx)
}

// NegateBlockFromJSObject returns a wrapped NegateBlock JavaScript class.
func NegateBlockFromJSObject(p js.Value, ctx js.Value) *NegateBlock {
	return &NegateBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NegateBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func NegateBlockArrayToJSArray(array []*NegateBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNegateBlock returns a new NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#constructor
func (ba *Babylon) NewNegateBlock(name string) *NegateBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("NegateBlock").New(args...)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getclassname
func (n *NegateBlock) GetClassName() string {

	retVal := n.p.Call("getClassName")
	return retVal.String()
}

// Output returns the Output property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#output
func (n *NegateBlock) Output() *NodeMaterialConnectionPoint {
	retVal := n.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// SetOutput sets the Output property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#output
func (n *NegateBlock) SetOutput(output *NodeMaterialConnectionPoint) *NegateBlock {
	n.p.Set("output", output.JSObject())
	return n
}

// Value returns the Value property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#value
func (n *NegateBlock) Value() *NodeMaterialConnectionPoint {
	retVal := n.p.Get("value")
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// SetValue sets the Value property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#value
func (n *NegateBlock) SetValue(value *NodeMaterialConnectionPoint) *NegateBlock {
	n.p.Set("value", value.JSObject())
	return n
}
