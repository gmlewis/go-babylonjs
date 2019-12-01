// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ClampBlock represents a babylon.js ClampBlock.
// Block used to clamp a float
type ClampBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ClampBlock) JSObject() js.Value { return c.p }

// ClampBlock returns a ClampBlock JavaScript class.
func (ba *Babylon) ClampBlock() *ClampBlock {
	p := ba.ctx.Get("ClampBlock")
	return ClampBlockFromJSObject(p, ba.ctx)
}

// ClampBlockFromJSObject returns a wrapped ClampBlock JavaScript class.
func ClampBlockFromJSObject(p js.Value, ctx js.Value) *ClampBlock {
	return &ClampBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewClampBlock returns a new ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock
func (ba *Babylon) NewClampBlock(name string) *ClampBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ClampBlock").New(args...)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#autoconfigure
func (c *ClampBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	c.p.Call("autoConfigure", args...)
}

// ClampBlockBindOpts contains optional parameters for ClampBlock.Bind.
type ClampBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#bind
func (c *ClampBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *ClampBlockBindOpts) {
	if opts == nil {
		opts = &ClampBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	c.p.Call("bind", args...)
}

// Build calls the Build method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#build
func (c *ClampBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := c.p.Call("build", args...)
	return retVal.Bool()
}

// ClampBlockCloneOpts contains optional parameters for ClampBlock.Clone.
type ClampBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#clone
func (c *ClampBlock) Clone(scene *Scene, opts *ClampBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &ClampBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := c.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, c.ctx)
}

// ClampBlockConnectToOpts contains optional parameters for ClampBlock.ConnectTo.
type ClampBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#connectto
func (c *ClampBlock) ConnectTo(other *NodeMaterialBlock, opts *ClampBlockConnectToOpts) *ClampBlock {
	if opts == nil {
		opts = &ClampBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := c.p.Call("connectTo", args...)
	return ClampBlockFromJSObject(retVal, c.ctx)
}

// Dispose calls the Dispose method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#dispose
func (c *ClampBlock) Dispose() {

	args := make([]interface{}, 0, 0+0)

	c.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getclassname
func (c *ClampBlock) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("getClassName", args...)
	return retVal.String()
}

// ClampBlockGetFirstAvailableInputOpts contains optional parameters for ClampBlock.GetFirstAvailableInput.
type ClampBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getfirstavailableinput
func (c *ClampBlock) GetFirstAvailableInput(opts *ClampBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ClampBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := c.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// ClampBlockGetFirstAvailableOutputOpts contains optional parameters for ClampBlock.GetFirstAvailableOutput.
type ClampBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getfirstavailableoutput
func (c *ClampBlock) GetFirstAvailableOutput(opts *ClampBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ClampBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := c.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// GetInputByName calls the GetInputByName method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getinputbyname
func (c *ClampBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := c.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// GetOutputByName calls the GetOutputByName method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getoutputbyname
func (c *ClampBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := c.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#getsiblingoutput
func (c *ClampBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := c.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// Initialize calls the Initialize method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#initialize
func (c *ClampBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	c.p.Call("initialize", args...)
}

// ClampBlockInitializeDefinesOpts contains optional parameters for ClampBlock.InitializeDefines.
type ClampBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#initializedefines
func (c *ClampBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ClampBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &ClampBlockInitializeDefinesOpts{}
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

	c.p.Call("initializeDefines", args...)
}

// ClampBlockIsReadyOpts contains optional parameters for ClampBlock.IsReady.
type ClampBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isready
func (c *ClampBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ClampBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &ClampBlockIsReadyOpts{}
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

	retVal := c.p.Call("isReady", args...)
	return retVal.Bool()
}

// ClampBlockPrepareDefinesOpts contains optional parameters for ClampBlock.PrepareDefines.
type ClampBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#preparedefines
func (c *ClampBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ClampBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &ClampBlockPrepareDefinesOpts{}
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

	c.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#providefallbacks
func (c *ClampBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	c.p.Call("provideFallbacks", args...)
}

// ClampBlockRegisterInputOpts contains optional parameters for ClampBlock.RegisterInput.
type ClampBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#registerinput
func (c *ClampBlock) RegisterInput(name string, jsType js.Value, opts *ClampBlockRegisterInputOpts) *ClampBlock {
	if opts == nil {
		opts = &ClampBlockRegisterInputOpts{}
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

	retVal := c.p.Call("registerInput", args...)
	return ClampBlockFromJSObject(retVal, c.ctx)
}

// ClampBlockRegisterOutputOpts contains optional parameters for ClampBlock.RegisterOutput.
type ClampBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#registeroutput
func (c *ClampBlock) RegisterOutput(name string, jsType js.Value, opts *ClampBlockRegisterOutputOpts) *ClampBlock {
	if opts == nil {
		opts = &ClampBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := c.p.Call("registerOutput", args...)
	return ClampBlockFromJSObject(retVal, c.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#replacerepeatablecontent
func (c *ClampBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	c.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#serialize
func (c *ClampBlock) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("serialize", args...)
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#updateuniformsandsamples
func (c *ClampBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	c.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the ClampBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#_deserialize
func (c *ClampBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	c.p.Call("_deserialize", args...)
}

/*

// BuildId returns the BuildId property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#buildid
func (c *ClampBlock) BuildId(buildId float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(buildId)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#buildid
func (c *ClampBlock) SetBuildId(buildId float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(buildId)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#comments
func (c *ClampBlock) Comments(comments string) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(comments)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#comments
func (c *ClampBlock) SetComments(comments string) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(comments)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#inputs
func (c *ClampBlock) Inputs(inputs *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(inputs.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#inputs
func (c *ClampBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(inputs.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isfinalmerger
func (c *ClampBlock) IsFinalMerger(isFinalMerger bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isFinalMerger)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isfinalmerger
func (c *ClampBlock) SetIsFinalMerger(isFinalMerger bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isFinalMerger)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isinput
func (c *ClampBlock) IsInput(isInput bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isInput)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isinput
func (c *ClampBlock) SetIsInput(isInput bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isInput)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isunique
func (c *ClampBlock) IsUnique(isUnique bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isUnique)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#isunique
func (c *ClampBlock) SetIsUnique(isUnique bool) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(isUnique)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Maximum returns the Maximum property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#maximum
func (c *ClampBlock) Maximum(maximum float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(maximum)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetMaximum sets the Maximum property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#maximum
func (c *ClampBlock) SetMaximum(maximum float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(maximum)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Minimum returns the Minimum property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#minimum
func (c *ClampBlock) Minimum(minimum float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(minimum)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetMinimum sets the Minimum property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#minimum
func (c *ClampBlock) SetMinimum(minimum float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(minimum)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#name
func (c *ClampBlock) Name(name string) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(name)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#name
func (c *ClampBlock) SetName(name string) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(name)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#output
func (c *ClampBlock) Output(output *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(output.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#output
func (c *ClampBlock) SetOutput(output *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(output.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#outputs
func (c *ClampBlock) Outputs(outputs *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(outputs.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#outputs
func (c *ClampBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(outputs.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#target
func (c *ClampBlock) Target(target js.Value) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(target)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#target
func (c *ClampBlock) SetTarget(target js.Value) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(target)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#uniqueid
func (c *ClampBlock) UniqueId(uniqueId float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(uniqueId)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#uniqueid
func (c *ClampBlock) SetUniqueId(uniqueId float64) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(uniqueId)
	return ClampBlockFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#value
func (c *ClampBlock) Value(value *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(value.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class ClampBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.clampblock#value
func (c *ClampBlock) SetValue(value *NodeMaterialConnectionPoint) *ClampBlock {
	p := ba.ctx.Get("ClampBlock").New(value.JSObject())
	return ClampBlockFromJSObject(p, ba.ctx)
}

*/
