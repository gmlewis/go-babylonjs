// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PosterizeBlock represents a babylon.js PosterizeBlock.
// Block used to posterize a value
//
// See: https://en.wikipedia.org/wiki/Posterization
type PosterizeBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PosterizeBlock) JSObject() js.Value { return p.p }

// PosterizeBlock returns a PosterizeBlock JavaScript class.
func (ba *Babylon) PosterizeBlock() *PosterizeBlock {
	p := ba.ctx.Get("PosterizeBlock")
	return PosterizeBlockFromJSObject(p, ba.ctx)
}

// PosterizeBlockFromJSObject returns a wrapped PosterizeBlock JavaScript class.
func PosterizeBlockFromJSObject(p js.Value, ctx js.Value) *PosterizeBlock {
	return &PosterizeBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// PosterizeBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func PosterizeBlockArrayToJSArray(array []*PosterizeBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPosterizeBlock returns a new PosterizeBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#constructor
func (ba *Babylon) NewPosterizeBlock(name string) *PosterizeBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("PosterizeBlock").New(args...)
	return PosterizeBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the PosterizeBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#getclassname
func (p *PosterizeBlock) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// Output returns the Output property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#output
func (p *PosterizeBlock) Output() *NodeMaterialConnectionPoint {
	retVal := p.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, p.ctx)
}

// SetOutput sets the Output property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#output
func (p *PosterizeBlock) SetOutput(output *NodeMaterialConnectionPoint) *PosterizeBlock {
	p.p.Set("output", output.JSObject())
	return p
}

// Steps returns the Steps property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#steps
func (p *PosterizeBlock) Steps() *NodeMaterialConnectionPoint {
	retVal := p.p.Get("steps")
	return NodeMaterialConnectionPointFromJSObject(retVal, p.ctx)
}

// SetSteps sets the Steps property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#steps
func (p *PosterizeBlock) SetSteps(steps *NodeMaterialConnectionPoint) *PosterizeBlock {
	p.p.Set("steps", steps.JSObject())
	return p
}

// Value returns the Value property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#value
func (p *PosterizeBlock) Value() *NodeMaterialConnectionPoint {
	retVal := p.p.Get("value")
	return NodeMaterialConnectionPointFromJSObject(retVal, p.ctx)
}

// SetValue sets the Value property of class PosterizeBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.posterizeblock#value
func (p *PosterizeBlock) SetValue(value *NodeMaterialConnectionPoint) *PosterizeBlock {
	p.p.Set("value", value.JSObject())
	return p
}
