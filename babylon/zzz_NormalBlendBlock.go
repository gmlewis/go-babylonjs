// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NormalBlendBlock represents a babylon.js NormalBlendBlock.
// Block used to blend normals
type NormalBlendBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NormalBlendBlock) JSObject() js.Value { return n.p }

// NormalBlendBlock returns a NormalBlendBlock JavaScript class.
func (ba *Babylon) NormalBlendBlock() *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock")
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// NormalBlendBlockFromJSObject returns a wrapped NormalBlendBlock JavaScript class.
func NormalBlendBlockFromJSObject(p js.Value, ctx js.Value) *NormalBlendBlock {
	return &NormalBlendBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NormalBlendBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func NormalBlendBlockArrayToJSArray(array []*NormalBlendBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNormalBlendBlock returns a new NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock
func (ba *Babylon) NewNormalBlendBlock(name string) *NormalBlendBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("NormalBlendBlock").New(args...)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#autoconfigure
func (n *NormalBlendBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	n.p.Call("autoConfigure", args...)
}

// NormalBlendBlockBindOpts contains optional parameters for NormalBlendBlock.Bind.
type NormalBlendBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#bind
func (n *NormalBlendBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *NormalBlendBlockBindOpts) {
	if opts == nil {
		opts = &NormalBlendBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	n.p.Call("bind", args...)
}

// Build calls the Build method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#build
func (n *NormalBlendBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := n.p.Call("build", args...)
	return retVal.Bool()
}

// NormalBlendBlockCloneOpts contains optional parameters for NormalBlendBlock.Clone.
type NormalBlendBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#clone
func (n *NormalBlendBlock) Clone(scene *Scene, opts *NormalBlendBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &NormalBlendBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := n.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, n.ctx)
}

// NormalBlendBlockConnectToOpts contains optional parameters for NormalBlendBlock.ConnectTo.
type NormalBlendBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#connectto
func (n *NormalBlendBlock) ConnectTo(other *NodeMaterialBlock, opts *NormalBlendBlockConnectToOpts) *NormalBlendBlock {
	if opts == nil {
		opts = &NormalBlendBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := n.p.Call("connectTo", args...)
	return NormalBlendBlockFromJSObject(retVal, n.ctx)
}

// Dispose calls the Dispose method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#dispose
func (n *NormalBlendBlock) Dispose() {

	n.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getclassname
func (n *NormalBlendBlock) GetClassName() string {

	retVal := n.p.Call("getClassName")
	return retVal.String()
}

// NormalBlendBlockGetFirstAvailableInputOpts contains optional parameters for NormalBlendBlock.GetFirstAvailableInput.
type NormalBlendBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getfirstavailableinput
func (n *NormalBlendBlock) GetFirstAvailableInput(opts *NormalBlendBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &NormalBlendBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := n.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// NormalBlendBlockGetFirstAvailableOutputOpts contains optional parameters for NormalBlendBlock.GetFirstAvailableOutput.
type NormalBlendBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getfirstavailableoutput
func (n *NormalBlendBlock) GetFirstAvailableOutput(opts *NormalBlendBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &NormalBlendBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := n.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetInputByName calls the GetInputByName method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getinputbyname
func (n *NormalBlendBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetOutputByName calls the GetOutputByName method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getoutputbyname
func (n *NormalBlendBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#getsiblingoutput
func (n *NormalBlendBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := n.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// Initialize calls the Initialize method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#initialize
func (n *NormalBlendBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	n.p.Call("initialize", args...)
}

// NormalBlendBlockInitializeDefinesOpts contains optional parameters for NormalBlendBlock.InitializeDefines.
type NormalBlendBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#initializedefines
func (n *NormalBlendBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NormalBlendBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &NormalBlendBlockInitializeDefinesOpts{}
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

	n.p.Call("initializeDefines", args...)
}

// NormalBlendBlockIsReadyOpts contains optional parameters for NormalBlendBlock.IsReady.
type NormalBlendBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isready
func (n *NormalBlendBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NormalBlendBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &NormalBlendBlockIsReadyOpts{}
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

	retVal := n.p.Call("isReady", args...)
	return retVal.Bool()
}

// NormalBlendBlockPrepareDefinesOpts contains optional parameters for NormalBlendBlock.PrepareDefines.
type NormalBlendBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#preparedefines
func (n *NormalBlendBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NormalBlendBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &NormalBlendBlockPrepareDefinesOpts{}
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

	n.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#providefallbacks
func (n *NormalBlendBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	n.p.Call("provideFallbacks", args...)
}

// NormalBlendBlockRegisterInputOpts contains optional parameters for NormalBlendBlock.RegisterInput.
type NormalBlendBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#registerinput
func (n *NormalBlendBlock) RegisterInput(name string, jsType js.Value, opts *NormalBlendBlockRegisterInputOpts) *NormalBlendBlock {
	if opts == nil {
		opts = &NormalBlendBlockRegisterInputOpts{}
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

	retVal := n.p.Call("registerInput", args...)
	return NormalBlendBlockFromJSObject(retVal, n.ctx)
}

// NormalBlendBlockRegisterOutputOpts contains optional parameters for NormalBlendBlock.RegisterOutput.
type NormalBlendBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#registeroutput
func (n *NormalBlendBlock) RegisterOutput(name string, jsType js.Value, opts *NormalBlendBlockRegisterOutputOpts) *NormalBlendBlock {
	if opts == nil {
		opts = &NormalBlendBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := n.p.Call("registerOutput", args...)
	return NormalBlendBlockFromJSObject(retVal, n.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#replacerepeatablecontent
func (n *NormalBlendBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	n.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#serialize
func (n *NormalBlendBlock) Serialize() interface{} {

	retVal := n.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#updateuniformsandsamples
func (n *NormalBlendBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	n.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#buildid
func (n *NormalBlendBlock) BuildId(buildId float64) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(buildId)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#buildid
func (n *NormalBlendBlock) SetBuildId(buildId float64) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(buildId)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#comments
func (n *NormalBlendBlock) Comments(comments string) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(comments)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#comments
func (n *NormalBlendBlock) SetComments(comments string) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(comments)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#inputs
func (n *NormalBlendBlock) Inputs(inputs *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(inputs.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#inputs
func (n *NormalBlendBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(inputs.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isfinalmerger
func (n *NormalBlendBlock) IsFinalMerger(isFinalMerger bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isFinalMerger)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isfinalmerger
func (n *NormalBlendBlock) SetIsFinalMerger(isFinalMerger bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isFinalMerger)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isinput
func (n *NormalBlendBlock) IsInput(isInput bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isInput)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isinput
func (n *NormalBlendBlock) SetIsInput(isInput bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isInput)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isunique
func (n *NormalBlendBlock) IsUnique(isUnique bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isUnique)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#isunique
func (n *NormalBlendBlock) SetIsUnique(isUnique bool) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(isUnique)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#name
func (n *NormalBlendBlock) Name(name string) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(name)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#name
func (n *NormalBlendBlock) SetName(name string) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(name)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// NormalMap0 returns the NormalMap0 property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#normalmap0
func (n *NormalBlendBlock) NormalMap0(normalMap0 *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(normalMap0.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetNormalMap0 sets the NormalMap0 property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#normalmap0
func (n *NormalBlendBlock) SetNormalMap0(normalMap0 *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(normalMap0.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// NormalMap1 returns the NormalMap1 property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#normalmap1
func (n *NormalBlendBlock) NormalMap1(normalMap1 *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(normalMap1.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetNormalMap1 sets the NormalMap1 property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#normalmap1
func (n *NormalBlendBlock) SetNormalMap1(normalMap1 *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(normalMap1.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#output
func (n *NormalBlendBlock) Output(output *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(output.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#output
func (n *NormalBlendBlock) SetOutput(output *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(output.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#outputs
func (n *NormalBlendBlock) Outputs(outputs *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(outputs.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#outputs
func (n *NormalBlendBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(outputs.JSObject())
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#target
func (n *NormalBlendBlock) Target(target js.Value) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(target)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#target
func (n *NormalBlendBlock) SetTarget(target js.Value) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(target)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#uniqueid
func (n *NormalBlendBlock) UniqueId(uniqueId float64) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(uniqueId)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class NormalBlendBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock#uniqueid
func (n *NormalBlendBlock) SetUniqueId(uniqueId float64) *NormalBlendBlock {
	p := ba.ctx.Get("NormalBlendBlock").New(uniqueId)
	return NormalBlendBlockFromJSObject(p, ba.ctx)
}

*/
