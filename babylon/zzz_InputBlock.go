// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputBlock represents a babylon.js InputBlock.
// Block used to expose an input value
type InputBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InputBlock) JSObject() js.Value { return i.p }

// InputBlock returns a InputBlock JavaScript class.
func (ba *Babylon) InputBlock() *InputBlock {
	p := ba.ctx.Get("InputBlock")
	return InputBlockFromJSObject(p, ba.ctx)
}

// InputBlockFromJSObject returns a wrapped InputBlock JavaScript class.
func InputBlockFromJSObject(p js.Value, ctx js.Value) *InputBlock {
	return &InputBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewInputBlockOpts contains optional parameters for NewInputBlock.
type NewInputBlockOpts struct {
	Target js.Value
	Type   js.Value
}

// NewInputBlock returns a new InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock
func (ba *Babylon) NewInputBlock(name string, opts *NewInputBlockOpts) *InputBlock {
	if opts == nil {
		opts = &NewInputBlockOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, name)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Type)
	}

	p := ba.ctx.Get("InputBlock").New(args...)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Animate calls the Animate method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#animate
func (i *InputBlock) Animate(scene *Scene) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	i.p.Call("animate", args...)
}

// AutoConfigure calls the AutoConfigure method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#autoconfigure
func (i *InputBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	i.p.Call("autoConfigure", args...)
}

// InputBlockBindOpts contains optional parameters for InputBlock.Bind.
type InputBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#bind
func (i *InputBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *InputBlockBindOpts) {
	if opts == nil {
		opts = &InputBlockBindOpts{}
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

// Build calls the Build method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#build
func (i *InputBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := i.p.Call("build", args...)
	return retVal.Bool()
}

// InputBlockCloneOpts contains optional parameters for InputBlock.Clone.
type InputBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#clone
func (i *InputBlock) Clone(scene *Scene, opts *InputBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &InputBlockCloneOpts{}
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

// InputBlockConnectToOpts contains optional parameters for InputBlock.ConnectTo.
type InputBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#connectto
func (i *InputBlock) ConnectTo(other *NodeMaterialBlock, opts *InputBlockConnectToOpts) *InputBlock {
	if opts == nil {
		opts = &InputBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := i.p.Call("connectTo", args...)
	return InputBlockFromJSObject(retVal, i.ctx)
}

// Dispose calls the Dispose method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#dispose
func (i *InputBlock) Dispose() {

	args := make([]interface{}, 0, 0+0)

	i.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getclassname
func (i *InputBlock) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := i.p.Call("getClassName", args...)
	return retVal.String()
}

// InputBlockGetFirstAvailableInputOpts contains optional parameters for InputBlock.GetFirstAvailableInput.
type InputBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getfirstavailableinput
func (i *InputBlock) GetFirstAvailableInput(opts *InputBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &InputBlockGetFirstAvailableInputOpts{}
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

// InputBlockGetFirstAvailableOutputOpts contains optional parameters for InputBlock.GetFirstAvailableOutput.
type InputBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getfirstavailableoutput
func (i *InputBlock) GetFirstAvailableOutput(opts *InputBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &InputBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getinputbyname
func (i *InputBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := i.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// GetOutputByName calls the GetOutputByName method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getoutputbyname
func (i *InputBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := i.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#getsiblingoutput
func (i *InputBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := i.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// Initialize calls the Initialize method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#initialize
func (i *InputBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	i.p.Call("initialize", args...)
}

// InputBlockInitializeDefinesOpts contains optional parameters for InputBlock.InitializeDefines.
type InputBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#initializedefines
func (i *InputBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InputBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &InputBlockInitializeDefinesOpts{}
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

// InputBlockIsReadyOpts contains optional parameters for InputBlock.IsReady.
type InputBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isready
func (i *InputBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InputBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &InputBlockIsReadyOpts{}
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

// InputBlockPrepareDefinesOpts contains optional parameters for InputBlock.PrepareDefines.
type InputBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#preparedefines
func (i *InputBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *InputBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &InputBlockPrepareDefinesOpts{}
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

// ProvideFallbacks calls the ProvideFallbacks method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#providefallbacks
func (i *InputBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	i.p.Call("provideFallbacks", args...)
}

// InputBlockRegisterInputOpts contains optional parameters for InputBlock.RegisterInput.
type InputBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#registerinput
func (i *InputBlock) RegisterInput(name string, jsType js.Value, opts *InputBlockRegisterInputOpts) *InputBlock {
	if opts == nil {
		opts = &InputBlockRegisterInputOpts{}
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
	return InputBlockFromJSObject(retVal, i.ctx)
}

// InputBlockRegisterOutputOpts contains optional parameters for InputBlock.RegisterOutput.
type InputBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#registeroutput
func (i *InputBlock) RegisterOutput(name string, jsType js.Value, opts *InputBlockRegisterOutputOpts) *InputBlock {
	if opts == nil {
		opts = &InputBlockRegisterOutputOpts{}
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
	return InputBlockFromJSObject(retVal, i.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#replacerepeatablecontent
func (i *InputBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	i.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#serialize
func (i *InputBlock) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := i.p.Call("serialize", args...)
	return retVal
}

// InputBlockSetAsAttributeOpts contains optional parameters for InputBlock.SetAsAttribute.
type InputBlockSetAsAttributeOpts struct {
	AttributeName *string
}

// SetAsAttribute calls the SetAsAttribute method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#setasattribute
func (i *InputBlock) SetAsAttribute(opts *InputBlockSetAsAttributeOpts) *InputBlock {
	if opts == nil {
		opts = &InputBlockSetAsAttributeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.AttributeName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AttributeName)
	}

	retVal := i.p.Call("setAsAttribute", args...)
	return InputBlockFromJSObject(retVal, i.ctx)
}

// SetAsSystemValue calls the SetAsSystemValue method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#setassystemvalue
func (i *InputBlock) SetAsSystemValue(value *NodeMaterialSystemValues) *InputBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value.JSObject())

	retVal := i.p.Call("setAsSystemValue", args...)
	return InputBlockFromJSObject(retVal, i.ctx)
}

// SetDefaultValue calls the SetDefaultValue method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#setdefaultvalue
func (i *InputBlock) SetDefaultValue() {

	args := make([]interface{}, 0, 0+0)

	i.p.Call("setDefaultValue", args...)
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#updateuniformsandsamples
func (i *InputBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	i.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the InputBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#_deserialize
func (i *InputBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	i.p.Call("_deserialize", args...)
}

/*

// AnimationType returns the AnimationType property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#animationtype
func (i *InputBlock) AnimationType(animationType *AnimatedInputBlockTypes) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(animationType.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetAnimationType sets the AnimationType property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#animationtype
func (i *InputBlock) SetAnimationType(animationType *AnimatedInputBlockTypes) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(animationType.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// AssociatedVariableName returns the AssociatedVariableName property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#associatedvariablename
func (i *InputBlock) AssociatedVariableName(associatedVariableName string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(associatedVariableName)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetAssociatedVariableName sets the AssociatedVariableName property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#associatedvariablename
func (i *InputBlock) SetAssociatedVariableName(associatedVariableName string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(associatedVariableName)
	return InputBlockFromJSObject(p, ba.ctx)
}

// BuildId returns the BuildId property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#buildid
func (i *InputBlock) BuildId(buildId float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(buildId)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#buildid
func (i *InputBlock) SetBuildId(buildId float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(buildId)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#comments
func (i *InputBlock) Comments(comments string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(comments)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#comments
func (i *InputBlock) SetComments(comments string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(comments)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#inputs
func (i *InputBlock) Inputs(inputs *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(inputs.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#inputs
func (i *InputBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(inputs.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsAttribute returns the IsAttribute property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isattribute
func (i *InputBlock) IsAttribute(isAttribute bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isAttribute)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsAttribute sets the IsAttribute property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isattribute
func (i *InputBlock) SetIsAttribute(isAttribute bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isAttribute)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsConstant returns the IsConstant property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isconstant
func (i *InputBlock) IsConstant(isConstant bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isConstant)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsConstant sets the IsConstant property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isconstant
func (i *InputBlock) SetIsConstant(isConstant bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isConstant)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isfinalmerger
func (i *InputBlock) IsFinalMerger(isFinalMerger bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isFinalMerger)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isfinalmerger
func (i *InputBlock) SetIsFinalMerger(isFinalMerger bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isFinalMerger)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isinput
func (i *InputBlock) IsInput(isInput bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isInput)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isinput
func (i *InputBlock) SetIsInput(isInput bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isInput)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsSystemValue returns the IsSystemValue property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#issystemvalue
func (i *InputBlock) IsSystemValue(isSystemValue bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isSystemValue)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsSystemValue sets the IsSystemValue property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#issystemvalue
func (i *InputBlock) SetIsSystemValue(isSystemValue bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isSystemValue)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsUndefined returns the IsUndefined property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isundefined
func (i *InputBlock) IsUndefined(isUndefined bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUndefined)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsUndefined sets the IsUndefined property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isundefined
func (i *InputBlock) SetIsUndefined(isUndefined bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUndefined)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsUniform returns the IsUniform property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isuniform
func (i *InputBlock) IsUniform(isUniform bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUniform)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsUniform sets the IsUniform property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isuniform
func (i *InputBlock) SetIsUniform(isUniform bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUniform)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isunique
func (i *InputBlock) IsUnique(isUnique bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUnique)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isunique
func (i *InputBlock) SetIsUnique(isUnique bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isUnique)
	return InputBlockFromJSObject(p, ba.ctx)
}

// IsVarying returns the IsVarying property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isvarying
func (i *InputBlock) IsVarying(isVarying bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isVarying)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetIsVarying sets the IsVarying property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#isvarying
func (i *InputBlock) SetIsVarying(isVarying bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(isVarying)
	return InputBlockFromJSObject(p, ba.ctx)
}

// MatrixMode returns the MatrixMode property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#matrixmode
func (i *InputBlock) MatrixMode(matrixMode float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(matrixMode)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetMatrixMode sets the MatrixMode property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#matrixmode
func (i *InputBlock) SetMatrixMode(matrixMode float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(matrixMode)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Max returns the Max property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#max
func (i *InputBlock) Max(max float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(max)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetMax sets the Max property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#max
func (i *InputBlock) SetMax(max float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(max)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Min returns the Min property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#min
func (i *InputBlock) Min(min float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(min)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetMin sets the Min property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#min
func (i *InputBlock) SetMin(min float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(min)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#name
func (i *InputBlock) Name(name string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(name)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#name
func (i *InputBlock) SetName(name string) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(name)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#output
func (i *InputBlock) Output(output *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(output.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#output
func (i *InputBlock) SetOutput(output *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(output.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#outputs
func (i *InputBlock) Outputs(outputs *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(outputs.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#outputs
func (i *InputBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(outputs.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SystemValue returns the SystemValue property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#systemvalue
func (i *InputBlock) SystemValue(systemValue *NodeMaterialSystemValues) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(systemValue.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetSystemValue sets the SystemValue property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#systemvalue
func (i *InputBlock) SetSystemValue(systemValue *NodeMaterialSystemValues) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(systemValue.JSObject())
	return InputBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#target
func (i *InputBlock) Target(target js.Value) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(target)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#target
func (i *InputBlock) SetTarget(target js.Value) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(target)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#type
func (i *InputBlock) Type(jsType js.Value) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(jsType)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#type
func (i *InputBlock) SetType(jsType js.Value) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(jsType)
	return InputBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#uniqueid
func (i *InputBlock) UniqueId(uniqueId float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(uniqueId)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#uniqueid
func (i *InputBlock) SetUniqueId(uniqueId float64) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(uniqueId)
	return InputBlockFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#value
func (i *InputBlock) Value(value interface{}) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(value)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#value
func (i *InputBlock) SetValue(value interface{}) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(value)
	return InputBlockFromJSObject(p, ba.ctx)
}

// ValueCallback returns the ValueCallback property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#valuecallback
func (i *InputBlock) ValueCallback(valueCallback func()) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(valueCallback)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetValueCallback sets the ValueCallback property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#valuecallback
func (i *InputBlock) SetValueCallback(valueCallback func()) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(valueCallback)
	return InputBlockFromJSObject(p, ba.ctx)
}

// VisibleInInspector returns the VisibleInInspector property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#visibleininspector
func (i *InputBlock) VisibleInInspector(visibleInInspector bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(visibleInInspector)
	return InputBlockFromJSObject(p, ba.ctx)
}

// SetVisibleInInspector sets the VisibleInInspector property of class InputBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.inputblock#visibleininspector
func (i *InputBlock) SetVisibleInInspector(visibleInInspector bool) *InputBlock {
	p := ba.ctx.Get("InputBlock").New(visibleInInspector)
	return InputBlockFromJSObject(p, ba.ctx)
}

*/
