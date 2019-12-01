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
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock
func (ba *Babylon) NewSimplexPerlin3DBlock(name string) *SimplexPerlin3DBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SimplexPerlin3DBlock").New(args...)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#autoconfigure
func (s *SimplexPerlin3DBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	s.p.Call("autoConfigure", args...)
}

// SimplexPerlin3DBlockBindOpts contains optional parameters for SimplexPerlin3DBlock.Bind.
type SimplexPerlin3DBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#bind
func (s *SimplexPerlin3DBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *SimplexPerlin3DBlockBindOpts) {
	if opts == nil {
		opts = &SimplexPerlin3DBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	s.p.Call("bind", args...)
}

// Build calls the Build method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#build
func (s *SimplexPerlin3DBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := s.p.Call("build", args...)
	return retVal.Bool()
}

// SimplexPerlin3DBlockCloneOpts contains optional parameters for SimplexPerlin3DBlock.Clone.
type SimplexPerlin3DBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#clone
func (s *SimplexPerlin3DBlock) Clone(scene *Scene, opts *SimplexPerlin3DBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &SimplexPerlin3DBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := s.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, s.ctx)
}

// SimplexPerlin3DBlockConnectToOpts contains optional parameters for SimplexPerlin3DBlock.ConnectTo.
type SimplexPerlin3DBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#connectto
func (s *SimplexPerlin3DBlock) ConnectTo(other *NodeMaterialBlock, opts *SimplexPerlin3DBlockConnectToOpts) *SimplexPerlin3DBlock {
	if opts == nil {
		opts = &SimplexPerlin3DBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("connectTo", args...)
	return SimplexPerlin3DBlockFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#dispose
func (s *SimplexPerlin3DBlock) Dispose() {

	s.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getclassname
func (s *SimplexPerlin3DBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// SimplexPerlin3DBlockGetFirstAvailableInputOpts contains optional parameters for SimplexPerlin3DBlock.GetFirstAvailableInput.
type SimplexPerlin3DBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getfirstavailableinput
func (s *SimplexPerlin3DBlock) GetFirstAvailableInput(opts *SimplexPerlin3DBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &SimplexPerlin3DBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := s.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SimplexPerlin3DBlockGetFirstAvailableOutputOpts contains optional parameters for SimplexPerlin3DBlock.GetFirstAvailableOutput.
type SimplexPerlin3DBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getfirstavailableoutput
func (s *SimplexPerlin3DBlock) GetFirstAvailableOutput(opts *SimplexPerlin3DBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &SimplexPerlin3DBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := s.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetInputByName calls the GetInputByName method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getinputbyname
func (s *SimplexPerlin3DBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetOutputByName calls the GetOutputByName method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getoutputbyname
func (s *SimplexPerlin3DBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#getsiblingoutput
func (s *SimplexPerlin3DBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := s.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// Initialize calls the Initialize method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#initialize
func (s *SimplexPerlin3DBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	s.p.Call("initialize", args...)
}

// SimplexPerlin3DBlockInitializeDefinesOpts contains optional parameters for SimplexPerlin3DBlock.InitializeDefines.
type SimplexPerlin3DBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#initializedefines
func (s *SimplexPerlin3DBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SimplexPerlin3DBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &SimplexPerlin3DBlockInitializeDefinesOpts{}
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

	s.p.Call("initializeDefines", args...)
}

// SimplexPerlin3DBlockIsReadyOpts contains optional parameters for SimplexPerlin3DBlock.IsReady.
type SimplexPerlin3DBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isready
func (s *SimplexPerlin3DBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SimplexPerlin3DBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &SimplexPerlin3DBlockIsReadyOpts{}
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

	retVal := s.p.Call("isReady", args...)
	return retVal.Bool()
}

// SimplexPerlin3DBlockPrepareDefinesOpts contains optional parameters for SimplexPerlin3DBlock.PrepareDefines.
type SimplexPerlin3DBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#preparedefines
func (s *SimplexPerlin3DBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SimplexPerlin3DBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &SimplexPerlin3DBlockPrepareDefinesOpts{}
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

	s.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#providefallbacks
func (s *SimplexPerlin3DBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	s.p.Call("provideFallbacks", args...)
}

// SimplexPerlin3DBlockRegisterInputOpts contains optional parameters for SimplexPerlin3DBlock.RegisterInput.
type SimplexPerlin3DBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#registerinput
func (s *SimplexPerlin3DBlock) RegisterInput(name string, jsType js.Value, opts *SimplexPerlin3DBlockRegisterInputOpts) *SimplexPerlin3DBlock {
	if opts == nil {
		opts = &SimplexPerlin3DBlockRegisterInputOpts{}
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

	retVal := s.p.Call("registerInput", args...)
	return SimplexPerlin3DBlockFromJSObject(retVal, s.ctx)
}

// SimplexPerlin3DBlockRegisterOutputOpts contains optional parameters for SimplexPerlin3DBlock.RegisterOutput.
type SimplexPerlin3DBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#registeroutput
func (s *SimplexPerlin3DBlock) RegisterOutput(name string, jsType js.Value, opts *SimplexPerlin3DBlockRegisterOutputOpts) *SimplexPerlin3DBlock {
	if opts == nil {
		opts = &SimplexPerlin3DBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := s.p.Call("registerOutput", args...)
	return SimplexPerlin3DBlockFromJSObject(retVal, s.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#replacerepeatablecontent
func (s *SimplexPerlin3DBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	s.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#serialize
func (s *SimplexPerlin3DBlock) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the SimplexPerlin3DBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#updateuniformsandsamples
func (s *SimplexPerlin3DBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	s.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#buildid
func (s *SimplexPerlin3DBlock) BuildId(buildId float64) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(buildId)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#buildid
func (s *SimplexPerlin3DBlock) SetBuildId(buildId float64) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(buildId)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#comments
func (s *SimplexPerlin3DBlock) Comments(comments string) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(comments)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#comments
func (s *SimplexPerlin3DBlock) SetComments(comments string) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(comments)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#inputs
func (s *SimplexPerlin3DBlock) Inputs(inputs *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(inputs.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#inputs
func (s *SimplexPerlin3DBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(inputs.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isfinalmerger
func (s *SimplexPerlin3DBlock) IsFinalMerger(isFinalMerger bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isFinalMerger)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isfinalmerger
func (s *SimplexPerlin3DBlock) SetIsFinalMerger(isFinalMerger bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isFinalMerger)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isinput
func (s *SimplexPerlin3DBlock) IsInput(isInput bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isInput)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isinput
func (s *SimplexPerlin3DBlock) SetIsInput(isInput bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isInput)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isunique
func (s *SimplexPerlin3DBlock) IsUnique(isUnique bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isUnique)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#isunique
func (s *SimplexPerlin3DBlock) SetIsUnique(isUnique bool) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(isUnique)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#name
func (s *SimplexPerlin3DBlock) Name(name string) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(name)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#name
func (s *SimplexPerlin3DBlock) SetName(name string) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(name)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#output
func (s *SimplexPerlin3DBlock) Output(output *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(output.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#output
func (s *SimplexPerlin3DBlock) SetOutput(output *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(output.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#outputs
func (s *SimplexPerlin3DBlock) Outputs(outputs *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(outputs.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#outputs
func (s *SimplexPerlin3DBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(outputs.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Seed returns the Seed property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#seed
func (s *SimplexPerlin3DBlock) Seed(seed *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(seed.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetSeed sets the Seed property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#seed
func (s *SimplexPerlin3DBlock) SetSeed(seed *NodeMaterialConnectionPoint) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(seed.JSObject())
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#target
func (s *SimplexPerlin3DBlock) Target(target js.Value) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(target)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#target
func (s *SimplexPerlin3DBlock) SetTarget(target js.Value) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(target)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#uniqueid
func (s *SimplexPerlin3DBlock) UniqueId(uniqueId float64) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(uniqueId)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class SimplexPerlin3DBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.simplexperlin3dblock#uniqueid
func (s *SimplexPerlin3DBlock) SetUniqueId(uniqueId float64) *SimplexPerlin3DBlock {
	p := ba.ctx.Get("SimplexPerlin3DBlock").New(uniqueId)
	return SimplexPerlin3DBlockFromJSObject(p, ba.ctx)
}

*/
