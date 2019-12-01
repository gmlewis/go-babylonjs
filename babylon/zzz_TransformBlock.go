// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TransformBlock represents a babylon.js TransformBlock.
// Block used to transform a vector (2, 3 or 4) with a matrix. It will generate a Vector4
type TransformBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TransformBlock) JSObject() js.Value { return t.p }

// TransformBlock returns a TransformBlock JavaScript class.
func (ba *Babylon) TransformBlock() *TransformBlock {
	p := ba.ctx.Get("TransformBlock")
	return TransformBlockFromJSObject(p, ba.ctx)
}

// TransformBlockFromJSObject returns a wrapped TransformBlock JavaScript class.
func TransformBlockFromJSObject(p js.Value, ctx js.Value) *TransformBlock {
	return &TransformBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewTransformBlock returns a new TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock
func (ba *Babylon) NewTransformBlock(name string) *TransformBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("TransformBlock").New(args...)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#autoconfigure
func (t *TransformBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	t.p.Call("autoConfigure", args...)
}

// TransformBlockBindOpts contains optional parameters for TransformBlock.Bind.
type TransformBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#bind
func (t *TransformBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *TransformBlockBindOpts) {
	if opts == nil {
		opts = &TransformBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	t.p.Call("bind", args...)
}

// Build calls the Build method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#build
func (t *TransformBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := t.p.Call("build", args...)
	return retVal.Bool()
}

// TransformBlockCloneOpts contains optional parameters for TransformBlock.Clone.
type TransformBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#clone
func (t *TransformBlock) Clone(scene *Scene, opts *TransformBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &TransformBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := t.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, t.ctx)
}

// TransformBlockConnectToOpts contains optional parameters for TransformBlock.ConnectTo.
type TransformBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#connectto
func (t *TransformBlock) ConnectTo(other *NodeMaterialBlock, opts *TransformBlockConnectToOpts) *TransformBlock {
	if opts == nil {
		opts = &TransformBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := t.p.Call("connectTo", args...)
	return TransformBlockFromJSObject(retVal, t.ctx)
}

// Dispose calls the Dispose method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#dispose
func (t *TransformBlock) Dispose() {

	args := make([]interface{}, 0, 0+0)

	t.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getclassname
func (t *TransformBlock) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := t.p.Call("getClassName", args...)
	return retVal.String()
}

// TransformBlockGetFirstAvailableInputOpts contains optional parameters for TransformBlock.GetFirstAvailableInput.
type TransformBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getfirstavailableinput
func (t *TransformBlock) GetFirstAvailableInput(opts *TransformBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TransformBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := t.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// TransformBlockGetFirstAvailableOutputOpts contains optional parameters for TransformBlock.GetFirstAvailableOutput.
type TransformBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getfirstavailableoutput
func (t *TransformBlock) GetFirstAvailableOutput(opts *TransformBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TransformBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := t.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetInputByName calls the GetInputByName method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getinputbyname
func (t *TransformBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetOutputByName calls the GetOutputByName method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getoutputbyname
func (t *TransformBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#getsiblingoutput
func (t *TransformBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := t.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// Initialize calls the Initialize method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#initialize
func (t *TransformBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	t.p.Call("initialize", args...)
}

// TransformBlockInitializeDefinesOpts contains optional parameters for TransformBlock.InitializeDefines.
type TransformBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#initializedefines
func (t *TransformBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TransformBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &TransformBlockInitializeDefinesOpts{}
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

	t.p.Call("initializeDefines", args...)
}

// TransformBlockIsReadyOpts contains optional parameters for TransformBlock.IsReady.
type TransformBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isready
func (t *TransformBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TransformBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &TransformBlockIsReadyOpts{}
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

	retVal := t.p.Call("isReady", args...)
	return retVal.Bool()
}

// TransformBlockPrepareDefinesOpts contains optional parameters for TransformBlock.PrepareDefines.
type TransformBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#preparedefines
func (t *TransformBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TransformBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &TransformBlockPrepareDefinesOpts{}
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

	t.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#providefallbacks
func (t *TransformBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	t.p.Call("provideFallbacks", args...)
}

// TransformBlockRegisterInputOpts contains optional parameters for TransformBlock.RegisterInput.
type TransformBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#registerinput
func (t *TransformBlock) RegisterInput(name string, jsType js.Value, opts *TransformBlockRegisterInputOpts) *TransformBlock {
	if opts == nil {
		opts = &TransformBlockRegisterInputOpts{}
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

	retVal := t.p.Call("registerInput", args...)
	return TransformBlockFromJSObject(retVal, t.ctx)
}

// TransformBlockRegisterOutputOpts contains optional parameters for TransformBlock.RegisterOutput.
type TransformBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#registeroutput
func (t *TransformBlock) RegisterOutput(name string, jsType js.Value, opts *TransformBlockRegisterOutputOpts) *TransformBlock {
	if opts == nil {
		opts = &TransformBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := t.p.Call("registerOutput", args...)
	return TransformBlockFromJSObject(retVal, t.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#replacerepeatablecontent
func (t *TransformBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	t.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#serialize
func (t *TransformBlock) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := t.p.Call("serialize", args...)
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#updateuniformsandsamples
func (t *TransformBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	t.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the TransformBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#_deserialize
func (t *TransformBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	t.p.Call("_deserialize", args...)
}

/*

// BuildId returns the BuildId property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#buildid
func (t *TransformBlock) BuildId(buildId float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(buildId)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#buildid
func (t *TransformBlock) SetBuildId(buildId float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(buildId)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#comments
func (t *TransformBlock) Comments(comments string) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(comments)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#comments
func (t *TransformBlock) SetComments(comments string) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(comments)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// ComplementW returns the ComplementW property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#complementw
func (t *TransformBlock) ComplementW(complementW float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(complementW)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetComplementW sets the ComplementW property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#complementw
func (t *TransformBlock) SetComplementW(complementW float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(complementW)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// ComplementZ returns the ComplementZ property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#complementz
func (t *TransformBlock) ComplementZ(complementZ float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(complementZ)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetComplementZ sets the ComplementZ property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#complementz
func (t *TransformBlock) SetComplementZ(complementZ float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(complementZ)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#inputs
func (t *TransformBlock) Inputs(inputs *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(inputs.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#inputs
func (t *TransformBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(inputs.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isfinalmerger
func (t *TransformBlock) IsFinalMerger(isFinalMerger bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isFinalMerger)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isfinalmerger
func (t *TransformBlock) SetIsFinalMerger(isFinalMerger bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isFinalMerger)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isinput
func (t *TransformBlock) IsInput(isInput bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isInput)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isinput
func (t *TransformBlock) SetIsInput(isInput bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isInput)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isunique
func (t *TransformBlock) IsUnique(isUnique bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isUnique)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#isunique
func (t *TransformBlock) SetIsUnique(isUnique bool) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(isUnique)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#name
func (t *TransformBlock) Name(name string) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(name)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#name
func (t *TransformBlock) SetName(name string) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(name)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#output
func (t *TransformBlock) Output(output *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(output.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#output
func (t *TransformBlock) SetOutput(output *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(output.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#outputs
func (t *TransformBlock) Outputs(outputs *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(outputs.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#outputs
func (t *TransformBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(outputs.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#target
func (t *TransformBlock) Target(target js.Value) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(target)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#target
func (t *TransformBlock) SetTarget(target js.Value) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(target)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Transform returns the Transform property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#transform
func (t *TransformBlock) Transform(transform *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(transform.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetTransform sets the Transform property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#transform
func (t *TransformBlock) SetTransform(transform *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(transform.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#uniqueid
func (t *TransformBlock) UniqueId(uniqueId float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(uniqueId)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#uniqueid
func (t *TransformBlock) SetUniqueId(uniqueId float64) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(uniqueId)
	return TransformBlockFromJSObject(p, ba.ctx)
}

// Vector returns the Vector property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#vector
func (t *TransformBlock) Vector(vector *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(vector.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

// SetVector sets the Vector property of class TransformBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.transformblock#vector
func (t *TransformBlock) SetVector(vector *NodeMaterialConnectionPoint) *TransformBlock {
	p := ba.ctx.Get("TransformBlock").New(vector.JSObject())
	return TransformBlockFromJSObject(p, ba.ctx)
}

*/
