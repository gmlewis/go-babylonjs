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
// https://doc.babylonjs.com/api/classes/babylon.bonesblock
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

	args = append(args, material.JSObject())

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

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	b.p.Call("bind", args...)
}

// Build calls the Build method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#build
func (b *BonesBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := b.p.Call("build", args...)
	return retVal.Bool()
}

// BonesBlockCloneOpts contains optional parameters for BonesBlock.Clone.
type BonesBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#clone
func (b *BonesBlock) Clone(scene *Scene, opts *BonesBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &BonesBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := b.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, b.ctx)
}

// BonesBlockConnectToOpts contains optional parameters for BonesBlock.ConnectTo.
type BonesBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#connectto
func (b *BonesBlock) ConnectTo(other *NodeMaterialBlock, opts *BonesBlockConnectToOpts) *BonesBlock {
	if opts == nil {
		opts = &BonesBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := b.p.Call("connectTo", args...)
	return BonesBlockFromJSObject(retVal, b.ctx)
}

// Dispose calls the Dispose method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#dispose
func (b *BonesBlock) Dispose() {

	b.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getclassname
func (b *BonesBlock) GetClassName() string {

	retVal := b.p.Call("getClassName")
	return retVal.String()
}

// BonesBlockGetFirstAvailableInputOpts contains optional parameters for BonesBlock.GetFirstAvailableInput.
type BonesBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getfirstavailableinput
func (b *BonesBlock) GetFirstAvailableInput(opts *BonesBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &BonesBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := b.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// BonesBlockGetFirstAvailableOutputOpts contains optional parameters for BonesBlock.GetFirstAvailableOutput.
type BonesBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getfirstavailableoutput
func (b *BonesBlock) GetFirstAvailableOutput(opts *BonesBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &BonesBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := b.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// GetInputByName calls the GetInputByName method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getinputbyname
func (b *BonesBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := b.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// GetOutputByName calls the GetOutputByName method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getoutputbyname
func (b *BonesBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := b.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#getsiblingoutput
func (b *BonesBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := b.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, b.ctx)
}

// Initialize calls the Initialize method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#initialize
func (b *BonesBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	b.p.Call("initialize", args...)
}

// BonesBlockInitializeDefinesOpts contains optional parameters for BonesBlock.InitializeDefines.
type BonesBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#initializedefines
func (b *BonesBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *BonesBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &BonesBlockInitializeDefinesOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	b.p.Call("initializeDefines", args...)
}

// BonesBlockIsReadyOpts contains optional parameters for BonesBlock.IsReady.
type BonesBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isready
func (b *BonesBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *BonesBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &BonesBlockIsReadyOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := b.p.Call("isReady", args...)
	return retVal.Bool()
}

// PrepareDefines calls the PrepareDefines method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#preparedefines
func (b *BonesBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	b.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#providefallbacks
func (b *BonesBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	b.p.Call("provideFallbacks", args...)
}

// BonesBlockRegisterInputOpts contains optional parameters for BonesBlock.RegisterInput.
type BonesBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#registerinput
func (b *BonesBlock) RegisterInput(name string, jsType js.Value, opts *BonesBlockRegisterInputOpts) *BonesBlock {
	if opts == nil {
		opts = &BonesBlockRegisterInputOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, name)
	args = append(args, jsType)

	if opts.IsOptional == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsOptional)
	}
	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := b.p.Call("registerInput", args...)
	return BonesBlockFromJSObject(retVal, b.ctx)
}

// BonesBlockRegisterOutputOpts contains optional parameters for BonesBlock.RegisterOutput.
type BonesBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#registeroutput
func (b *BonesBlock) RegisterOutput(name string, jsType js.Value, opts *BonesBlockRegisterOutputOpts) *BonesBlock {
	if opts == nil {
		opts = &BonesBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := b.p.Call("registerOutput", args...)
	return BonesBlockFromJSObject(retVal, b.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#replacerepeatablecontent
func (b *BonesBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	b.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#serialize
func (b *BonesBlock) Serialize() interface{} {

	retVal := b.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#updateuniformsandsamples
func (b *BonesBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	b.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#buildid
func (b *BonesBlock) BuildId(buildId float64) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(buildId)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#buildid
func (b *BonesBlock) SetBuildId(buildId float64) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(buildId)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#comments
func (b *BonesBlock) Comments(comments string) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(comments)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#comments
func (b *BonesBlock) SetComments(comments string) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(comments)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#inputs
func (b *BonesBlock) Inputs(inputs *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(inputs.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#inputs
func (b *BonesBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(inputs.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isfinalmerger
func (b *BonesBlock) IsFinalMerger(isFinalMerger bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isFinalMerger)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isfinalmerger
func (b *BonesBlock) SetIsFinalMerger(isFinalMerger bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isFinalMerger)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isinput
func (b *BonesBlock) IsInput(isInput bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isInput)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isinput
func (b *BonesBlock) SetIsInput(isInput bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isInput)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isunique
func (b *BonesBlock) IsUnique(isUnique bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isUnique)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#isunique
func (b *BonesBlock) SetIsUnique(isUnique bool) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(isUnique)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// MatricesIndices returns the MatricesIndices property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindices
func (b *BonesBlock) MatricesIndices(matricesIndices *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesIndices.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetMatricesIndices sets the MatricesIndices property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindices
func (b *BonesBlock) SetMatricesIndices(matricesIndices *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesIndices.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// MatricesIndicesExtra returns the MatricesIndicesExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindicesextra
func (b *BonesBlock) MatricesIndicesExtra(matricesIndicesExtra *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesIndicesExtra.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetMatricesIndicesExtra sets the MatricesIndicesExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesindicesextra
func (b *BonesBlock) SetMatricesIndicesExtra(matricesIndicesExtra *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesIndicesExtra.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// MatricesWeights returns the MatricesWeights property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweights
func (b *BonesBlock) MatricesWeights(matricesWeights *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesWeights.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetMatricesWeights sets the MatricesWeights property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweights
func (b *BonesBlock) SetMatricesWeights(matricesWeights *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesWeights.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// MatricesWeightsExtra returns the MatricesWeightsExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweightsextra
func (b *BonesBlock) MatricesWeightsExtra(matricesWeightsExtra *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesWeightsExtra.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetMatricesWeightsExtra sets the MatricesWeightsExtra property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#matricesweightsextra
func (b *BonesBlock) SetMatricesWeightsExtra(matricesWeightsExtra *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(matricesWeightsExtra.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#name
func (b *BonesBlock) Name(name string) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(name)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#name
func (b *BonesBlock) SetName(name string) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(name)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#output
func (b *BonesBlock) Output(output *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(output.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#output
func (b *BonesBlock) SetOutput(output *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(output.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#outputs
func (b *BonesBlock) Outputs(outputs *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(outputs.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#outputs
func (b *BonesBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(outputs.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#target
func (b *BonesBlock) Target(target js.Value) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(target)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#target
func (b *BonesBlock) SetTarget(target js.Value) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(target)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#uniqueid
func (b *BonesBlock) UniqueId(uniqueId float64) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(uniqueId)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#uniqueid
func (b *BonesBlock) SetUniqueId(uniqueId float64) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(uniqueId)
	return BonesBlockFromJSObject(p, ba.ctx)
}

// World returns the World property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#world
func (b *BonesBlock) World(world *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(world.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

// SetWorld sets the World property of class BonesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock#world
func (b *BonesBlock) SetWorld(world *NodeMaterialConnectionPoint) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(world.JSObject())
	return BonesBlockFromJSObject(p, ba.ctx)
}

*/
