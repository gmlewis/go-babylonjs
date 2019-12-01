// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SmoothStepBlock represents a babylon.js SmoothStepBlock.
// Block used to smooth step a value
type SmoothStepBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SmoothStepBlock) JSObject() js.Value { return s.p }

// SmoothStepBlock returns a SmoothStepBlock JavaScript class.
func (ba *Babylon) SmoothStepBlock() *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock")
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SmoothStepBlockFromJSObject returns a wrapped SmoothStepBlock JavaScript class.
func SmoothStepBlockFromJSObject(p js.Value, ctx js.Value) *SmoothStepBlock {
	return &SmoothStepBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// SmoothStepBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func SmoothStepBlockArrayToJSArray(array []*SmoothStepBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSmoothStepBlock returns a new SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock
func (ba *Babylon) NewSmoothStepBlock(name string) *SmoothStepBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SmoothStepBlock").New(args...)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#autoconfigure
func (s *SmoothStepBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	s.p.Call("autoConfigure", args...)
}

// SmoothStepBlockBindOpts contains optional parameters for SmoothStepBlock.Bind.
type SmoothStepBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#bind
func (s *SmoothStepBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *SmoothStepBlockBindOpts) {
	if opts == nil {
		opts = &SmoothStepBlockBindOpts{}
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

// Build calls the Build method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#build
func (s *SmoothStepBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := s.p.Call("build", args...)
	return retVal.Bool()
}

// SmoothStepBlockCloneOpts contains optional parameters for SmoothStepBlock.Clone.
type SmoothStepBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#clone
func (s *SmoothStepBlock) Clone(scene *Scene, opts *SmoothStepBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &SmoothStepBlockCloneOpts{}
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

// SmoothStepBlockConnectToOpts contains optional parameters for SmoothStepBlock.ConnectTo.
type SmoothStepBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#connectto
func (s *SmoothStepBlock) ConnectTo(other *NodeMaterialBlock, opts *SmoothStepBlockConnectToOpts) *SmoothStepBlock {
	if opts == nil {
		opts = &SmoothStepBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("connectTo", args...)
	return SmoothStepBlockFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#dispose
func (s *SmoothStepBlock) Dispose() {

	s.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getclassname
func (s *SmoothStepBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// SmoothStepBlockGetFirstAvailableInputOpts contains optional parameters for SmoothStepBlock.GetFirstAvailableInput.
type SmoothStepBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getfirstavailableinput
func (s *SmoothStepBlock) GetFirstAvailableInput(opts *SmoothStepBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &SmoothStepBlockGetFirstAvailableInputOpts{}
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

// SmoothStepBlockGetFirstAvailableOutputOpts contains optional parameters for SmoothStepBlock.GetFirstAvailableOutput.
type SmoothStepBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getfirstavailableoutput
func (s *SmoothStepBlock) GetFirstAvailableOutput(opts *SmoothStepBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &SmoothStepBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getinputbyname
func (s *SmoothStepBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetOutputByName calls the GetOutputByName method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getoutputbyname
func (s *SmoothStepBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getsiblingoutput
func (s *SmoothStepBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := s.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// Initialize calls the Initialize method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#initialize
func (s *SmoothStepBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	s.p.Call("initialize", args...)
}

// SmoothStepBlockInitializeDefinesOpts contains optional parameters for SmoothStepBlock.InitializeDefines.
type SmoothStepBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#initializedefines
func (s *SmoothStepBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SmoothStepBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &SmoothStepBlockInitializeDefinesOpts{}
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

// SmoothStepBlockIsReadyOpts contains optional parameters for SmoothStepBlock.IsReady.
type SmoothStepBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isready
func (s *SmoothStepBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SmoothStepBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &SmoothStepBlockIsReadyOpts{}
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

// SmoothStepBlockPrepareDefinesOpts contains optional parameters for SmoothStepBlock.PrepareDefines.
type SmoothStepBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#preparedefines
func (s *SmoothStepBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *SmoothStepBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &SmoothStepBlockPrepareDefinesOpts{}
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

// ProvideFallbacks calls the ProvideFallbacks method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#providefallbacks
func (s *SmoothStepBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	s.p.Call("provideFallbacks", args...)
}

// SmoothStepBlockRegisterInputOpts contains optional parameters for SmoothStepBlock.RegisterInput.
type SmoothStepBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#registerinput
func (s *SmoothStepBlock) RegisterInput(name string, jsType js.Value, opts *SmoothStepBlockRegisterInputOpts) *SmoothStepBlock {
	if opts == nil {
		opts = &SmoothStepBlockRegisterInputOpts{}
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
	return SmoothStepBlockFromJSObject(retVal, s.ctx)
}

// SmoothStepBlockRegisterOutputOpts contains optional parameters for SmoothStepBlock.RegisterOutput.
type SmoothStepBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#registeroutput
func (s *SmoothStepBlock) RegisterOutput(name string, jsType js.Value, opts *SmoothStepBlockRegisterOutputOpts) *SmoothStepBlock {
	if opts == nil {
		opts = &SmoothStepBlockRegisterOutputOpts{}
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
	return SmoothStepBlockFromJSObject(retVal, s.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#replacerepeatablecontent
func (s *SmoothStepBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	s.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#serialize
func (s *SmoothStepBlock) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#updateuniformsandsamples
func (s *SmoothStepBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	s.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#buildid
func (s *SmoothStepBlock) BuildId(buildId float64) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(buildId)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#buildid
func (s *SmoothStepBlock) SetBuildId(buildId float64) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(buildId)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#comments
func (s *SmoothStepBlock) Comments(comments string) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(comments)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#comments
func (s *SmoothStepBlock) SetComments(comments string) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(comments)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Edge0 returns the Edge0 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge0
func (s *SmoothStepBlock) Edge0(edge0 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(edge0.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetEdge0 sets the Edge0 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge0
func (s *SmoothStepBlock) SetEdge0(edge0 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(edge0.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Edge1 returns the Edge1 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge1
func (s *SmoothStepBlock) Edge1(edge1 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(edge1.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetEdge1 sets the Edge1 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge1
func (s *SmoothStepBlock) SetEdge1(edge1 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(edge1.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#inputs
func (s *SmoothStepBlock) Inputs(inputs *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(inputs.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#inputs
func (s *SmoothStepBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(inputs.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isfinalmerger
func (s *SmoothStepBlock) IsFinalMerger(isFinalMerger bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isFinalMerger)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isfinalmerger
func (s *SmoothStepBlock) SetIsFinalMerger(isFinalMerger bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isFinalMerger)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isinput
func (s *SmoothStepBlock) IsInput(isInput bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isInput)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isinput
func (s *SmoothStepBlock) SetIsInput(isInput bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isInput)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isunique
func (s *SmoothStepBlock) IsUnique(isUnique bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isUnique)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#isunique
func (s *SmoothStepBlock) SetIsUnique(isUnique bool) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(isUnique)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#name
func (s *SmoothStepBlock) Name(name string) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(name)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#name
func (s *SmoothStepBlock) SetName(name string) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(name)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#output
func (s *SmoothStepBlock) Output(output *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(output.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#output
func (s *SmoothStepBlock) SetOutput(output *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(output.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#outputs
func (s *SmoothStepBlock) Outputs(outputs *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(outputs.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#outputs
func (s *SmoothStepBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(outputs.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#target
func (s *SmoothStepBlock) Target(target js.Value) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(target)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#target
func (s *SmoothStepBlock) SetTarget(target js.Value) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(target)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#uniqueid
func (s *SmoothStepBlock) UniqueId(uniqueId float64) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(uniqueId)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#uniqueid
func (s *SmoothStepBlock) SetUniqueId(uniqueId float64) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(uniqueId)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#value
func (s *SmoothStepBlock) Value(value *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(value.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#value
func (s *SmoothStepBlock) SetValue(value *NodeMaterialConnectionPoint) *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock").New(value.JSObject())
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

*/
