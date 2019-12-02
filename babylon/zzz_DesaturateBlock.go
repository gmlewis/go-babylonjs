// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DesaturateBlock represents a babylon.js DesaturateBlock.
// Block used to desaturate a color
type DesaturateBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DesaturateBlock) JSObject() js.Value { return d.p }

// DesaturateBlock returns a DesaturateBlock JavaScript class.
func (ba *Babylon) DesaturateBlock() *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock")
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// DesaturateBlockFromJSObject returns a wrapped DesaturateBlock JavaScript class.
func DesaturateBlockFromJSObject(p js.Value, ctx js.Value) *DesaturateBlock {
	return &DesaturateBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// DesaturateBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func DesaturateBlockArrayToJSArray(array []*DesaturateBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDesaturateBlock returns a new DesaturateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock
func (ba *Babylon) NewDesaturateBlock(name string) *DesaturateBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("DesaturateBlock").New(args...)
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the DesaturateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#getclassname
func (d *DesaturateBlock) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

/*

// Color returns the Color property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#color
func (d *DesaturateBlock) Color(color *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(color.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// SetColor sets the Color property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#color
func (d *DesaturateBlock) SetColor(color *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(color.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// Level returns the Level property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#level
func (d *DesaturateBlock) Level(level *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(level.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// SetLevel sets the Level property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#level
func (d *DesaturateBlock) SetLevel(level *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(level.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#output
func (d *DesaturateBlock) Output(output *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(output.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class DesaturateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.desaturateblock#output
func (d *DesaturateBlock) SetOutput(output *NodeMaterialConnectionPoint) *DesaturateBlock {
	p := ba.ctx.Get("DesaturateBlock").New(output.JSObject())
	return DesaturateBlockFromJSObject(p, ba.ctx)
}

*/
