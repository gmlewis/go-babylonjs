// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BonesBlock represents a babylon.js BonesBlock.
// Block used to add support for vertex skinning (bones)
type BonesBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BonesBlock) JSObject() js.Value { return b.p }

// BonesBlock returns a BonesBlock JavaScript class.
func (ba *Babylon) BonesBlock() *BonesBlock {
	p := ba.ctx.Get("BonesBlock")
	return BonesBlockFromJSObject(p, ba.ctx)
}

// BonesBlockFromJSObject returns a wrapped BonesBlock JavaScript class.
func BonesBlockFromJSObject(p js.Value, ctx js.Value) *BonesBlock {
	return &BonesBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// BonesBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func BonesBlockArrayToJSArray(array []*BonesBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBonesBlock returns a new BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#constructor
func (ba *Babylon) NewBonesBlock(name string) *BonesBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("BonesBlock").New(args...)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#autoconfigure
func (b *BonesBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	if material == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, material.JSObject())
	}

	b.p.Call("autoConfigure", args...)
}

// BonesBlockBindOpts contains optional parameters for BonesBlock.Bind.
type BonesBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#bind
func (b *BonesBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *BonesBlockBindOpts) {
	if opts == nil {
		opts = &BonesBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	b.p.Call("bind", args...)
}

// GetClassName calls the GetClassName method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getclassname
func (b *BonesBlock) GetClassName() string {

	retVal := b.p.Call("getClassName")
	return retVal.String()
}

// Initialize calls the Initialize method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#initialize
func (b *BonesBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	if state == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, state.JSObject())
	}

	b.p.Call("initialize", args...)
}

// PrepareDefines calls the PrepareDefines method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#preparedefines
func (b *BonesBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

	args := make([]interface{}, 0, 3+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	args = append(args, defines)

	b.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#providefallbacks
func (b *BonesBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if fallbacks == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, fallbacks.JSObject())
	}

	b.p.Call("provideFallbacks", args...)
}

// MatricesIndices returns the MatricesIndices property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindices
func (b *BonesBlock) MatricesIndices() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("matricesIndices")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetMatricesIndices sets the MatricesIndices property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindices
func (b *BonesBlock) SetMatricesIndices(matricesIndices *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("matricesIndices", matricesIndices.JSObject())
	return b
}

// MatricesIndicesExtra returns the MatricesIndicesExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindicesextra
func (b *BonesBlock) MatricesIndicesExtra() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("matricesIndicesExtra")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetMatricesIndicesExtra sets the MatricesIndicesExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindicesextra
func (b *BonesBlock) SetMatricesIndicesExtra(matricesIndicesExtra *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("matricesIndicesExtra", matricesIndicesExtra.JSObject())
	return b
}

// MatricesWeights returns the MatricesWeights property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweights
func (b *BonesBlock) MatricesWeights() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("matricesWeights")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetMatricesWeights sets the MatricesWeights property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweights
func (b *BonesBlock) SetMatricesWeights(matricesWeights *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("matricesWeights", matricesWeights.JSObject())
	return b
}

// MatricesWeightsExtra returns the MatricesWeightsExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweightsextra
func (b *BonesBlock) MatricesWeightsExtra() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("matricesWeightsExtra")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetMatricesWeightsExtra sets the MatricesWeightsExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweightsextra
func (b *BonesBlock) SetMatricesWeightsExtra(matricesWeightsExtra *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("matricesWeightsExtra", matricesWeightsExtra.JSObject())
	return b
}

// Output returns the Output property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#output
func (b *BonesBlock) Output() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetOutput sets the Output property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#output
func (b *BonesBlock) SetOutput(output *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("output", output.JSObject())
	return b
}

// World returns the World property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#world
func (b *BonesBlock) World() *NodeMaterialConnectionPoint {
	retVal := b.p.Get("world")
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// SetWorld sets the World property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#world
func (b *BonesBlock) SetWorld(world *NodeMaterialConnectionPoint) *BonesBlock {
	b.p.Set("world", world.JSObject())
	return b
}
