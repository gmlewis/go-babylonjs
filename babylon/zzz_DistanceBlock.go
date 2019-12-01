// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DistanceBlock represents a babylon.js DistanceBlock.
// Block used to get the distance between 2 values
type DistanceBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DistanceBlock) JSObject() js.Value { return d.p }

// DistanceBlock returns a DistanceBlock JavaScript class.
func (ba *Babylon) DistanceBlock() *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock")
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// DistanceBlockFromJSObject returns a wrapped DistanceBlock JavaScript class.
func DistanceBlockFromJSObject(p js.Value, ctx js.Value) *DistanceBlock {
	return &DistanceBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// DistanceBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func DistanceBlockArrayToJSArray(array []*DistanceBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDistanceBlock returns a new DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock
func (ba *Babylon) NewDistanceBlock(name string) *DistanceBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("DistanceBlock").New(args...)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#autoconfigure
func (d *DistanceBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	d.p.Call("autoConfigure", args...)
}

// DistanceBlockBindOpts contains optional parameters for DistanceBlock.Bind.
type DistanceBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#bind
func (d *DistanceBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *DistanceBlockBindOpts) {
	if opts == nil {
		opts = &DistanceBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	d.p.Call("bind", args...)
}

// Build calls the Build method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#build
func (d *DistanceBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := d.p.Call("build", args...)
	return retVal.Bool()
}

// DistanceBlockCloneOpts contains optional parameters for DistanceBlock.Clone.
type DistanceBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#clone
func (d *DistanceBlock) Clone(scene *Scene, opts *DistanceBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &DistanceBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := d.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, d.ctx)
}

// DistanceBlockConnectToOpts contains optional parameters for DistanceBlock.ConnectTo.
type DistanceBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#connectto
func (d *DistanceBlock) ConnectTo(other *NodeMaterialBlock, opts *DistanceBlockConnectToOpts) *DistanceBlock {
	if opts == nil {
		opts = &DistanceBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := d.p.Call("connectTo", args...)
	return DistanceBlockFromJSObject(retVal, d.ctx)
}

// Dispose calls the Dispose method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#dispose
func (d *DistanceBlock) Dispose() {

	d.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getclassname
func (d *DistanceBlock) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// DistanceBlockGetFirstAvailableInputOpts contains optional parameters for DistanceBlock.GetFirstAvailableInput.
type DistanceBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getfirstavailableinput
func (d *DistanceBlock) GetFirstAvailableInput(opts *DistanceBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &DistanceBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := d.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// DistanceBlockGetFirstAvailableOutputOpts contains optional parameters for DistanceBlock.GetFirstAvailableOutput.
type DistanceBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getfirstavailableoutput
func (d *DistanceBlock) GetFirstAvailableOutput(opts *DistanceBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &DistanceBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := d.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetInputByName calls the GetInputByName method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getinputbyname
func (d *DistanceBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := d.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetOutputByName calls the GetOutputByName method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getoutputbyname
func (d *DistanceBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := d.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#getsiblingoutput
func (d *DistanceBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := d.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// Initialize calls the Initialize method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#initialize
func (d *DistanceBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	d.p.Call("initialize", args...)
}

// DistanceBlockInitializeDefinesOpts contains optional parameters for DistanceBlock.InitializeDefines.
type DistanceBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#initializedefines
func (d *DistanceBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DistanceBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &DistanceBlockInitializeDefinesOpts{}
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

	d.p.Call("initializeDefines", args...)
}

// DistanceBlockIsReadyOpts contains optional parameters for DistanceBlock.IsReady.
type DistanceBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isready
func (d *DistanceBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DistanceBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &DistanceBlockIsReadyOpts{}
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

	retVal := d.p.Call("isReady", args...)
	return retVal.Bool()
}

// DistanceBlockPrepareDefinesOpts contains optional parameters for DistanceBlock.PrepareDefines.
type DistanceBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#preparedefines
func (d *DistanceBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DistanceBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &DistanceBlockPrepareDefinesOpts{}
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

	d.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#providefallbacks
func (d *DistanceBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	d.p.Call("provideFallbacks", args...)
}

// DistanceBlockRegisterInputOpts contains optional parameters for DistanceBlock.RegisterInput.
type DistanceBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#registerinput
func (d *DistanceBlock) RegisterInput(name string, jsType js.Value, opts *DistanceBlockRegisterInputOpts) *DistanceBlock {
	if opts == nil {
		opts = &DistanceBlockRegisterInputOpts{}
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

	retVal := d.p.Call("registerInput", args...)
	return DistanceBlockFromJSObject(retVal, d.ctx)
}

// DistanceBlockRegisterOutputOpts contains optional parameters for DistanceBlock.RegisterOutput.
type DistanceBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#registeroutput
func (d *DistanceBlock) RegisterOutput(name string, jsType js.Value, opts *DistanceBlockRegisterOutputOpts) *DistanceBlock {
	if opts == nil {
		opts = &DistanceBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := d.p.Call("registerOutput", args...)
	return DistanceBlockFromJSObject(retVal, d.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#replacerepeatablecontent
func (d *DistanceBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	d.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#serialize
func (d *DistanceBlock) Serialize() interface{} {

	retVal := d.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the DistanceBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#updateuniformsandsamples
func (d *DistanceBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	d.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#buildid
func (d *DistanceBlock) BuildId(buildId float64) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(buildId)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#buildid
func (d *DistanceBlock) SetBuildId(buildId float64) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(buildId)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#comments
func (d *DistanceBlock) Comments(comments string) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(comments)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#comments
func (d *DistanceBlock) SetComments(comments string) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(comments)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#inputs
func (d *DistanceBlock) Inputs(inputs *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(inputs.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#inputs
func (d *DistanceBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(inputs.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isfinalmerger
func (d *DistanceBlock) IsFinalMerger(isFinalMerger bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isFinalMerger)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isfinalmerger
func (d *DistanceBlock) SetIsFinalMerger(isFinalMerger bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isFinalMerger)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isinput
func (d *DistanceBlock) IsInput(isInput bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isInput)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isinput
func (d *DistanceBlock) SetIsInput(isInput bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isInput)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isunique
func (d *DistanceBlock) IsUnique(isUnique bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isUnique)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#isunique
func (d *DistanceBlock) SetIsUnique(isUnique bool) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(isUnique)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Left returns the Left property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#left
func (d *DistanceBlock) Left(left *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(left.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetLeft sets the Left property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#left
func (d *DistanceBlock) SetLeft(left *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(left.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#name
func (d *DistanceBlock) Name(name string) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(name)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#name
func (d *DistanceBlock) SetName(name string) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(name)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#output
func (d *DistanceBlock) Output(output *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(output.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#output
func (d *DistanceBlock) SetOutput(output *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(output.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#outputs
func (d *DistanceBlock) Outputs(outputs *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(outputs.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#outputs
func (d *DistanceBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(outputs.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Right returns the Right property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#right
func (d *DistanceBlock) Right(right *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(right.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetRight sets the Right property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#right
func (d *DistanceBlock) SetRight(right *NodeMaterialConnectionPoint) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(right.JSObject())
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#target
func (d *DistanceBlock) Target(target js.Value) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(target)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#target
func (d *DistanceBlock) SetTarget(target js.Value) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(target)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#uniqueid
func (d *DistanceBlock) UniqueId(uniqueId float64) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(uniqueId)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class DistanceBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.distanceblock#uniqueid
func (d *DistanceBlock) SetUniqueId(uniqueId float64) *DistanceBlock {
	p := ba.ctx.Get("DistanceBlock").New(uniqueId)
	return DistanceBlockFromJSObject(p, ba.ctx)
}

*/
