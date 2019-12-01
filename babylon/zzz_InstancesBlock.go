// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InstancesBlock represents a babylon.js InstancesBlock.
// Block used to add support for instances
//
// See: https://doc.babylonjs.com/how_to/how_to_use_instances
type InstancesBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InstancesBlock) JSObject() js.Value { return i.p }

// InstancesBlock returns a InstancesBlock JavaScript class.
func (ba *Babylon) InstancesBlock() *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock")
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// InstancesBlockFromJSObject returns a wrapped InstancesBlock JavaScript class.
func InstancesBlockFromJSObject(p js.Value, ctx js.Value) *InstancesBlock {
	return &InstancesBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// InstancesBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func InstancesBlockArrayToJSArray(array []*InstancesBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewInstancesBlock returns a new InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock
func (ba *Babylon) NewInstancesBlock(name string) *InstancesBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("InstancesBlock").New(args...)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#autoconfigure
func (i *InstancesBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	i.p.Call("autoConfigure", args...)
}

// InstancesBlockBindOpts contains optional parameters for InstancesBlock.Bind.
type InstancesBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#bind
func (i *InstancesBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *InstancesBlockBindOpts) {
	if opts == nil {
		opts = &InstancesBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	i.p.Call("bind", args...)
}

// Build calls the Build method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#build
func (i *InstancesBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := i.p.Call("build", args...)
	return retVal.Bool()
}

// InstancesBlockCloneOpts contains optional parameters for InstancesBlock.Clone.
type InstancesBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#clone
func (i *InstancesBlock) Clone(scene *Scene, opts *InstancesBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &InstancesBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := i.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, i.ctx)
}

// InstancesBlockConnectToOpts contains optional parameters for InstancesBlock.ConnectTo.
type InstancesBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#connectto
func (i *InstancesBlock) ConnectTo(other *NodeMaterialBlock, opts *InstancesBlockConnectToOpts) *InstancesBlock {
	if opts == nil {
		opts = &InstancesBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := i.p.Call("connectTo", args...)
	return InstancesBlockFromJSObject(retVal, i.ctx)
}

// Dispose calls the Dispose method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#dispose
func (i *InstancesBlock) Dispose() {

	i.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getclassname
func (i *InstancesBlock) GetClassName() string {

	retVal := i.p.Call("getClassName")
	return retVal.String()
}

// InstancesBlockGetFirstAvailableInputOpts contains optional parameters for InstancesBlock.GetFirstAvailableInput.
type InstancesBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getfirstavailableinput
func (i *InstancesBlock) GetFirstAvailableInput(opts *InstancesBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &InstancesBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := i.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// InstancesBlockGetFirstAvailableOutputOpts contains optional parameters for InstancesBlock.GetFirstAvailableOutput.
type InstancesBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getfirstavailableoutput
func (i *InstancesBlock) GetFirstAvailableOutput(opts *InstancesBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &InstancesBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := i.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// GetInputByName calls the GetInputByName method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getinputbyname
func (i *InstancesBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := i.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// GetOutputByName calls the GetOutputByName method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getoutputbyname
func (i *InstancesBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := i.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#getsiblingoutput
func (i *InstancesBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := i.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// Initialize calls the Initialize method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#initialize
func (i *InstancesBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	i.p.Call("initialize", args...)
}

// InstancesBlockInitializeDefinesOpts contains optional parameters for InstancesBlock.InitializeDefines.
type InstancesBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#initializedefines
func (i *InstancesBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InstancesBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &InstancesBlockInitializeDefinesOpts{}
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

	i.p.Call("initializeDefines", args...)
}

// InstancesBlockIsReadyOpts contains optional parameters for InstancesBlock.IsReady.
type InstancesBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isready
func (i *InstancesBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InstancesBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &InstancesBlockIsReadyOpts{}
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

	retVal := i.p.Call("isReady", args...)
	return retVal.Bool()
}

// InstancesBlockPrepareDefinesOpts contains optional parameters for InstancesBlock.PrepareDefines.
type InstancesBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#preparedefines
func (i *InstancesBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InstancesBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &InstancesBlockPrepareDefinesOpts{}
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

	i.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#providefallbacks
func (i *InstancesBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	i.p.Call("provideFallbacks", args...)
}

// InstancesBlockRegisterInputOpts contains optional parameters for InstancesBlock.RegisterInput.
type InstancesBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#registerinput
func (i *InstancesBlock) RegisterInput(name string, jsType js.Value, opts *InstancesBlockRegisterInputOpts) *InstancesBlock {
	if opts == nil {
		opts = &InstancesBlockRegisterInputOpts{}
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

	retVal := i.p.Call("registerInput", args...)
	return InstancesBlockFromJSObject(retVal, i.ctx)
}

// InstancesBlockRegisterOutputOpts contains optional parameters for InstancesBlock.RegisterOutput.
type InstancesBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#registeroutput
func (i *InstancesBlock) RegisterOutput(name string, jsType js.Value, opts *InstancesBlockRegisterOutputOpts) *InstancesBlock {
	if opts == nil {
		opts = &InstancesBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := i.p.Call("registerOutput", args...)
	return InstancesBlockFromJSObject(retVal, i.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#replacerepeatablecontent
func (i *InstancesBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	i.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#serialize
func (i *InstancesBlock) Serialize() interface{} {

	retVal := i.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the InstancesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#updateuniformsandsamples
func (i *InstancesBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	i.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#buildid
func (i *InstancesBlock) BuildId(buildId float64) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(buildId)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#buildid
func (i *InstancesBlock) SetBuildId(buildId float64) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(buildId)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#comments
func (i *InstancesBlock) Comments(comments string) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(comments)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#comments
func (i *InstancesBlock) SetComments(comments string) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(comments)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#inputs
func (i *InstancesBlock) Inputs(inputs *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(inputs.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#inputs
func (i *InstancesBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(inputs.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isfinalmerger
func (i *InstancesBlock) IsFinalMerger(isFinalMerger bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isFinalMerger)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isfinalmerger
func (i *InstancesBlock) SetIsFinalMerger(isFinalMerger bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isFinalMerger)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isinput
func (i *InstancesBlock) IsInput(isInput bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isInput)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isinput
func (i *InstancesBlock) SetIsInput(isInput bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isInput)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isunique
func (i *InstancesBlock) IsUnique(isUnique bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isUnique)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#isunique
func (i *InstancesBlock) SetIsUnique(isUnique bool) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(isUnique)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#name
func (i *InstancesBlock) Name(name string) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(name)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#name
func (i *InstancesBlock) SetName(name string) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(name)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#output
func (i *InstancesBlock) Output(output *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(output.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#output
func (i *InstancesBlock) SetOutput(output *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(output.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#outputs
func (i *InstancesBlock) Outputs(outputs *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(outputs.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#outputs
func (i *InstancesBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(outputs.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#target
func (i *InstancesBlock) Target(target js.Value) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(target)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#target
func (i *InstancesBlock) SetTarget(target js.Value) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(target)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#uniqueid
func (i *InstancesBlock) UniqueId(uniqueId float64) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(uniqueId)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#uniqueid
func (i *InstancesBlock) SetUniqueId(uniqueId float64) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(uniqueId)
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// World returns the World property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world
func (i *InstancesBlock) World(world *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetWorld sets the World property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world
func (i *InstancesBlock) SetWorld(world *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// World0 returns the World0 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world0
func (i *InstancesBlock) World0(world0 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world0.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetWorld0 sets the World0 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world0
func (i *InstancesBlock) SetWorld0(world0 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world0.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// World1 returns the World1 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world1
func (i *InstancesBlock) World1(world1 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world1.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetWorld1 sets the World1 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world1
func (i *InstancesBlock) SetWorld1(world1 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world1.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// World2 returns the World2 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world2
func (i *InstancesBlock) World2(world2 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world2.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetWorld2 sets the World2 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world2
func (i *InstancesBlock) SetWorld2(world2 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world2.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// World3 returns the World3 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world3
func (i *InstancesBlock) World3(world3 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world3.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

// SetWorld3 sets the World3 property of class InstancesBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.instancesblock#world3
func (i *InstancesBlock) SetWorld3(world3 *NodeMaterialConnectionPoint) *InstancesBlock {
	p := ba.ctx.Get("InstancesBlock").New(world3.JSObject())
	return InstancesBlockFromJSObject(p, ba.ctx)
}

*/
