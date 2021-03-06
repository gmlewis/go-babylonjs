// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SimplexPerlin3DBlock represents a babylon.js SimplexPerlin3DBlock.
// block used to Generate a Simplex Perlin 3d Noise Pattern
type SimplexPerlin3DBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SimplexPerlin3DBlock) JSObject() js.Value { return s.p }

// SimplexPerlin3DBlock returns a SimplexPerlin3DBlock JavaScript class.
func (ba *Babylon) SimplexPerlin3DBlock() *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock")
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SimplexPerlin3DBlockFromJSObject returns a wrapped SimplexPerlin3DBlock JavaScript class.
func SimplexPerlin3DBlockFromJSObject(p js.Value, ctx js.Value) *SimplexPerlin3DBlock {
	return &SimplexPerlin3DBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// SimplexPerlin3DBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func SimplexPerlin3DBlockArrayToJSArray(array []*SimplexPerlin3DBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSimplexPerlin3DBlock returns a new SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#constructor
func (ba *Babylon) NewSimplexPerlin3DBlock(name string) *SimplexPerlin3DBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SimplexPerlin3DBlock").New(args...)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getclassname
func (s *SimplexPerlin3DBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// Output returns the Output property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#output
func (s *SimplexPerlin3DBlock) Output() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetOutput sets the Output property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#output
func (s *SimplexPerlin3DBlock) SetOutput(output *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	s.p.Set("output", output.JSObject())
	return s
}

// Seed returns the Seed property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#seed
func (s *SimplexPerlin3DBlock) Seed() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("seed")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetSeed sets the Seed property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#seed
func (s *SimplexPerlin3DBlock) SetSeed(seed *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	s.p.Set("seed", seed.JSObject())
	return s
}
