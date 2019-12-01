// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WaveBlock represents a babylon.js WaveBlock.
// Block used to apply wave operation to floats
type WaveBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WaveBlock) JSObject() js.Value { return w.p }

// WaveBlock returns a WaveBlock JavaScript class.
func (ba *Babylon) WaveBlock() *WaveBlock {
	p := ba.ctx.Get("WaveBlock")
	return WaveBlockFromJSObject(p, ba.ctx)
}

// WaveBlockFromJSObject returns a wrapped WaveBlock JavaScript class.
func WaveBlockFromJSObject(p js.Value, ctx js.Value) *WaveBlock {
	return &WaveBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewWaveBlock returns a new WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock
func (ba *Babylon) NewWaveBlock(name string) *WaveBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("WaveBlock").New(args...)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#autoconfigure
func (w *WaveBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	w.p.Call("autoConfigure", args...)
}

// WaveBlockBindOpts contains optional parameters for WaveBlock.Bind.
type WaveBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#bind
func (w *WaveBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *WaveBlockBindOpts) {
	if opts == nil {
		opts = &WaveBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	w.p.Call("bind", args...)
}

// Build calls the Build method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#build
func (w *WaveBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := w.p.Call("build", args...)
	return retVal.Bool()
}

// WaveBlockCloneOpts contains optional parameters for WaveBlock.Clone.
type WaveBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#clone
func (w *WaveBlock) Clone(scene *Scene, opts *WaveBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &WaveBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := w.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, w.ctx)
}

// WaveBlockConnectToOpts contains optional parameters for WaveBlock.ConnectTo.
type WaveBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#connectto
func (w *WaveBlock) ConnectTo(other *NodeMaterialBlock, opts *WaveBlockConnectToOpts) *WaveBlock {
	if opts == nil {
		opts = &WaveBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := w.p.Call("connectTo", args...)
	return WaveBlockFromJSObject(retVal, w.ctx)
}

// Dispose calls the Dispose method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#dispose
func (w *WaveBlock) Dispose() {

	args := make([]interface{}, 0, 0+0)

	w.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getclassname
func (w *WaveBlock) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := w.p.Call("getClassName", args...)
	return retVal.String()
}

// WaveBlockGetFirstAvailableInputOpts contains optional parameters for WaveBlock.GetFirstAvailableInput.
type WaveBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getfirstavailableinput
func (w *WaveBlock) GetFirstAvailableInput(opts *WaveBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &WaveBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := w.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, w.ctx)
}

// WaveBlockGetFirstAvailableOutputOpts contains optional parameters for WaveBlock.GetFirstAvailableOutput.
type WaveBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getfirstavailableoutput
func (w *WaveBlock) GetFirstAvailableOutput(opts *WaveBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &WaveBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := w.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, w.ctx)
}

// GetInputByName calls the GetInputByName method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getinputbyname
func (w *WaveBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := w.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, w.ctx)
}

// GetOutputByName calls the GetOutputByName method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getoutputbyname
func (w *WaveBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := w.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, w.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#getsiblingoutput
func (w *WaveBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := w.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, w.ctx)
}

// Initialize calls the Initialize method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#initialize
func (w *WaveBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	w.p.Call("initialize", args...)
}

// WaveBlockInitializeDefinesOpts contains optional parameters for WaveBlock.InitializeDefines.
type WaveBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#initializedefines
func (w *WaveBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *WaveBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &WaveBlockInitializeDefinesOpts{}
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

	w.p.Call("initializeDefines", args...)
}

// WaveBlockIsReadyOpts contains optional parameters for WaveBlock.IsReady.
type WaveBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isready
func (w *WaveBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *WaveBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &WaveBlockIsReadyOpts{}
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

	retVal := w.p.Call("isReady", args...)
	return retVal.Bool()
}

// WaveBlockPrepareDefinesOpts contains optional parameters for WaveBlock.PrepareDefines.
type WaveBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#preparedefines
func (w *WaveBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *WaveBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &WaveBlockPrepareDefinesOpts{}
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

	w.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#providefallbacks
func (w *WaveBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	w.p.Call("provideFallbacks", args...)
}

// WaveBlockRegisterInputOpts contains optional parameters for WaveBlock.RegisterInput.
type WaveBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#registerinput
func (w *WaveBlock) RegisterInput(name string, jsType js.Value, opts *WaveBlockRegisterInputOpts) *WaveBlock {
	if opts == nil {
		opts = &WaveBlockRegisterInputOpts{}
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

	retVal := w.p.Call("registerInput", args...)
	return WaveBlockFromJSObject(retVal, w.ctx)
}

// WaveBlockRegisterOutputOpts contains optional parameters for WaveBlock.RegisterOutput.
type WaveBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#registeroutput
func (w *WaveBlock) RegisterOutput(name string, jsType js.Value, opts *WaveBlockRegisterOutputOpts) *WaveBlock {
	if opts == nil {
		opts = &WaveBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := w.p.Call("registerOutput", args...)
	return WaveBlockFromJSObject(retVal, w.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#replacerepeatablecontent
func (w *WaveBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	w.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#serialize
func (w *WaveBlock) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := w.p.Call("serialize", args...)
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#updateuniformsandsamples
func (w *WaveBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	w.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#_deserialize
func (w *WaveBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	w.p.Call("_deserialize", args...)
}

/*

// BuildId returns the BuildId property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#buildid
func (w *WaveBlock) BuildId(buildId float64) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(buildId)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#buildid
func (w *WaveBlock) SetBuildId(buildId float64) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(buildId)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#comments
func (w *WaveBlock) Comments(comments string) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(comments)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#comments
func (w *WaveBlock) SetComments(comments string) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(comments)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Input returns the Input property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#input
func (w *WaveBlock) Input(input *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(input.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetInput sets the Input property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#input
func (w *WaveBlock) SetInput(input *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(input.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#inputs
func (w *WaveBlock) Inputs(inputs *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(inputs.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#inputs
func (w *WaveBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(inputs.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isfinalmerger
func (w *WaveBlock) IsFinalMerger(isFinalMerger bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isFinalMerger)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isfinalmerger
func (w *WaveBlock) SetIsFinalMerger(isFinalMerger bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isFinalMerger)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isinput
func (w *WaveBlock) IsInput(isInput bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isInput)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isinput
func (w *WaveBlock) SetIsInput(isInput bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isInput)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isunique
func (w *WaveBlock) IsUnique(isUnique bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isUnique)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#isunique
func (w *WaveBlock) SetIsUnique(isUnique bool) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(isUnique)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Kind returns the Kind property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#kind
func (w *WaveBlock) Kind(kind *WaveBlockKind) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(kind.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetKind sets the Kind property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#kind
func (w *WaveBlock) SetKind(kind *WaveBlockKind) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(kind.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#name
func (w *WaveBlock) Name(name string) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(name)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#name
func (w *WaveBlock) SetName(name string) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(name)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#output
func (w *WaveBlock) Output(output *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(output.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#output
func (w *WaveBlock) SetOutput(output *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(output.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#outputs
func (w *WaveBlock) Outputs(outputs *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(outputs.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#outputs
func (w *WaveBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(outputs.JSObject())
	return WaveBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#target
func (w *WaveBlock) Target(target js.Value) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(target)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#target
func (w *WaveBlock) SetTarget(target js.Value) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(target)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#uniqueid
func (w *WaveBlock) UniqueId(uniqueId float64) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(uniqueId)
	return WaveBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class WaveBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock#uniqueid
func (w *WaveBlock) SetUniqueId(uniqueId float64) *WaveBlock {
	p := ba.ctx.Get("WaveBlock").New(uniqueId)
	return WaveBlockFromJSObject(p, ba.ctx)
}

*/
