// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DistanceBlock represents a babylon.js DistanceBlock.
// Block used to get the distance between 2 values
type DistanceBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DistanceBlock) JSObject() js.Value { return d.p }

// DistanceBlock returns a DistanceBlock JavaScript class.
func (ba *Babylon) DistanceBlock() *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock")
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// DistanceBlockFromJSObject returns a wrapped DistanceBlock JavaScript class.
func DistanceBlockFromJSObject(p js.Value, ctx js.Value) *DistanceBlock {
	return &DistanceBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// DistanceBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func DistanceBlockArrayToJSArray(array []*DistanceBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDistanceBlock returns a new DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#constructor
func (ba *Babylon) NewDistanceBlock(name string) *DistanceBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("DistanceBlock").New(args...)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getclassname
func (d *DistanceBlock) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// Left returns the Left property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#left
func (d *DistanceBlock) Left() *NodeMaterialConnectionPoint {
	retVal := d.p.Get("left")
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// SetLeft sets the Left property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#left
func (d *DistanceBlock) SetLeft(left *NodeMaterialConnectionPoint) *DistanceBlock {
	d.p.Set("left", left.JSObject())
	return d
}

// Output returns the Output property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#output
func (d *DistanceBlock) Output() *NodeMaterialConnectionPoint {
	retVal := d.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// SetOutput sets the Output property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#output
func (d *DistanceBlock) SetOutput(output *NodeMaterialConnectionPoint) *DistanceBlock {
	d.p.Set("output", output.JSObject())
	return d
}

// Right returns the Right property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#right
func (d *DistanceBlock) Right() *NodeMaterialConnectionPoint {
	retVal := d.p.Get("right")
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// SetRight sets the Right property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#right
func (d *DistanceBlock) SetRight(right *NodeMaterialConnectionPoint) *DistanceBlock {
	d.p.Set("right", right.JSObject())
	return d
}
