// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StepBlock represents a babylon.js StepBlock.
// Block used to step a value
type StepBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StepBlock) JSObject() js.Value { return s.p }

// StepBlock returns a StepBlock JavaScript class.
func (ba *Babylon) StepBlock() *StepBlock {
	p := ba.ctx.Get("StepBlock")
	return StepBlockFromJSObject(p, ba.ctx)
}

// StepBlockFromJSObject returns a wrapped StepBlock JavaScript class.
func StepBlockFromJSObject(p js.Value, ctx js.Value) *StepBlock {
	return &StepBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// StepBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func StepBlockArrayToJSArray(array []*StepBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStepBlock returns a new StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock
func (ba *Babylon) NewStepBlock(name string) *StepBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("StepBlock").New(args...)
	return StepBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#autoconfigure
func (s *StepBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	s.p.Call("autoConfigure", args...)
}

// StepBlockBindOpts contains optional parameters for StepBlock.Bind.
type StepBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#bind
func (s *StepBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *StepBlockBindOpts) {
	if opts == nil {
		opts = &StepBlockBindOpts{}
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

// Build calls the Build method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#build
func (s *StepBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := s.p.Call("build", args...)
	return retVal.Bool()
}

// StepBlockCloneOpts contains optional parameters for StepBlock.Clone.
type StepBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#clone
func (s *StepBlock) Clone(scene *Scene, opts *StepBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &StepBlockCloneOpts{}
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

// StepBlockConnectToOpts contains optional parameters for StepBlock.ConnectTo.
type StepBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#connectto
func (s *StepBlock) ConnectTo(other *NodeMaterialBlock, opts *StepBlockConnectToOpts) *StepBlock {
	if opts == nil {
		opts = &StepBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("connectTo", args...)
	return StepBlockFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#dispose
func (s *StepBlock) Dispose() {

	s.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getclassname
func (s *StepBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// StepBlockGetFirstAvailableInputOpts contains optional parameters for StepBlock.GetFirstAvailableInput.
type StepBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getfirstavailableinput
func (s *StepBlock) GetFirstAvailableInput(opts *StepBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &StepBlockGetFirstAvailableInputOpts{}
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

// StepBlockGetFirstAvailableOutputOpts contains optional parameters for StepBlock.GetFirstAvailableOutput.
type StepBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getfirstavailableoutput
func (s *StepBlock) GetFirstAvailableOutput(opts *StepBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &StepBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getinputbyname
func (s *StepBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetOutputByName calls the GetOutputByName method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getoutputbyname
func (s *StepBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#getsiblingoutput
func (s *StepBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := s.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// Initialize calls the Initialize method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#initialize
func (s *StepBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	s.p.Call("initialize", args...)
}

// StepBlockInitializeDefinesOpts contains optional parameters for StepBlock.InitializeDefines.
type StepBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#initializedefines
func (s *StepBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *StepBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &StepBlockInitializeDefinesOpts{}
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

// StepBlockIsReadyOpts contains optional parameters for StepBlock.IsReady.
type StepBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isready
func (s *StepBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *StepBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &StepBlockIsReadyOpts{}
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

// StepBlockPrepareDefinesOpts contains optional parameters for StepBlock.PrepareDefines.
type StepBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#preparedefines
func (s *StepBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *StepBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &StepBlockPrepareDefinesOpts{}
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

// ProvideFallbacks calls the ProvideFallbacks method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#providefallbacks
func (s *StepBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	s.p.Call("provideFallbacks", args...)
}

// StepBlockRegisterInputOpts contains optional parameters for StepBlock.RegisterInput.
type StepBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#registerinput
func (s *StepBlock) RegisterInput(name string, jsType js.Value, opts *StepBlockRegisterInputOpts) *StepBlock {
	if opts == nil {
		opts = &StepBlockRegisterInputOpts{}
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
	return StepBlockFromJSObject(retVal, s.ctx)
}

// StepBlockRegisterOutputOpts contains optional parameters for StepBlock.RegisterOutput.
type StepBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#registeroutput
func (s *StepBlock) RegisterOutput(name string, jsType js.Value, opts *StepBlockRegisterOutputOpts) *StepBlock {
	if opts == nil {
		opts = &StepBlockRegisterOutputOpts{}
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
	return StepBlockFromJSObject(retVal, s.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#replacerepeatablecontent
func (s *StepBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	s.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#serialize
func (s *StepBlock) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#updateuniformsandsamples
func (s *StepBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	s.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#buildid
func (s *StepBlock) BuildId(buildId float64) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(buildId)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#buildid
func (s *StepBlock) SetBuildId(buildId float64) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(buildId)
	return StepBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#comments
func (s *StepBlock) Comments(comments string) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(comments)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#comments
func (s *StepBlock) SetComments(comments string) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(comments)
	return StepBlockFromJSObject(p, ba.ctx)
}

// Edge returns the Edge property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#edge
func (s *StepBlock) Edge(edge *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(edge.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetEdge sets the Edge property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#edge
func (s *StepBlock) SetEdge(edge *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(edge.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#inputs
func (s *StepBlock) Inputs(inputs *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(inputs.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#inputs
func (s *StepBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(inputs.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isfinalmerger
func (s *StepBlock) IsFinalMerger(isFinalMerger bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isFinalMerger)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isfinalmerger
func (s *StepBlock) SetIsFinalMerger(isFinalMerger bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isFinalMerger)
	return StepBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isinput
func (s *StepBlock) IsInput(isInput bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isInput)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isinput
func (s *StepBlock) SetIsInput(isInput bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isInput)
	return StepBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isunique
func (s *StepBlock) IsUnique(isUnique bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isUnique)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#isunique
func (s *StepBlock) SetIsUnique(isUnique bool) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(isUnique)
	return StepBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#name
func (s *StepBlock) Name(name string) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(name)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#name
func (s *StepBlock) SetName(name string) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(name)
	return StepBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#output
func (s *StepBlock) Output(output *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(output.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#output
func (s *StepBlock) SetOutput(output *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(output.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#outputs
func (s *StepBlock) Outputs(outputs *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(outputs.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#outputs
func (s *StepBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(outputs.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#target
func (s *StepBlock) Target(target js.Value) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(target)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#target
func (s *StepBlock) SetTarget(target js.Value) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(target)
	return StepBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#uniqueid
func (s *StepBlock) UniqueId(uniqueId float64) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(uniqueId)
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#uniqueid
func (s *StepBlock) SetUniqueId(uniqueId float64) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(uniqueId)
	return StepBlockFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#value
func (s *StepBlock) Value(value *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(value.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class StepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock#value
func (s *StepBlock) SetValue(value *NodeMaterialConnectionPoint) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(value.JSObject())
	return StepBlockFromJSObject(p, ba.ctx)
}

*/
