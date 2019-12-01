// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GradientBlock represents a babylon.js GradientBlock.
// Block used to return a color from a gradient based on an input value between 0 and 1
type GradientBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GradientBlock) JSObject() js.Value { return g.p }

// GradientBlock returns a GradientBlock JavaScript class.
func (ba *Babylon) GradientBlock() *GradientBlock {
	p := ba.ctx.Get("GradientBlock")
	return GradientBlockFromJSObject(p, ba.ctx)
}

// GradientBlockFromJSObject returns a wrapped GradientBlock JavaScript class.
func GradientBlockFromJSObject(p js.Value, ctx js.Value) *GradientBlock {
	return &GradientBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// GradientBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func GradientBlockArrayToJSArray(array []*GradientBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGradientBlock returns a new GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock
func (ba *Babylon) NewGradientBlock(name string) *GradientBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("GradientBlock").New(args...)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#autoconfigure
func (g *GradientBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	g.p.Call("autoConfigure", args...)
}

// GradientBlockBindOpts contains optional parameters for GradientBlock.Bind.
type GradientBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#bind
func (g *GradientBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *GradientBlockBindOpts) {
	if opts == nil {
		opts = &GradientBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	g.p.Call("bind", args...)
}

// Build calls the Build method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#build
func (g *GradientBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := g.p.Call("build", args...)
	return retVal.Bool()
}

// GradientBlockCloneOpts contains optional parameters for GradientBlock.Clone.
type GradientBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#clone
func (g *GradientBlock) Clone(scene *Scene, opts *GradientBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &GradientBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := g.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, g.ctx)
}

// GradientBlockConnectToOpts contains optional parameters for GradientBlock.ConnectTo.
type GradientBlockConnectToOpts struct {
	Options js.Value
}

// ConnectTo calls the ConnectTo method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#connectto
func (g *GradientBlock) ConnectTo(other *NodeMaterialBlock, opts *GradientBlockConnectToOpts) *GradientBlock {
	if opts == nil {
		opts = &GradientBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := g.p.Call("connectTo", args...)
	return GradientBlockFromJSObject(retVal, g.ctx)
}

// Dispose calls the Dispose method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#dispose
func (g *GradientBlock) Dispose() {

	g.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getclassname
func (g *GradientBlock) GetClassName() string {

	retVal := g.p.Call("getClassName")
	return retVal.String()
}

// GradientBlockGetFirstAvailableInputOpts contains optional parameters for GradientBlock.GetFirstAvailableInput.
type GradientBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getfirstavailableinput
func (g *GradientBlock) GetFirstAvailableInput(opts *GradientBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &GradientBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := g.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, g.ctx)
}

// GradientBlockGetFirstAvailableOutputOpts contains optional parameters for GradientBlock.GetFirstAvailableOutput.
type GradientBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getfirstavailableoutput
func (g *GradientBlock) GetFirstAvailableOutput(opts *GradientBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &GradientBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := g.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, g.ctx)
}

// GetInputByName calls the GetInputByName method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getinputbyname
func (g *GradientBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := g.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, g.ctx)
}

// GetOutputByName calls the GetOutputByName method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getoutputbyname
func (g *GradientBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := g.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, g.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#getsiblingoutput
func (g *GradientBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := g.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, g.ctx)
}

// Initialize calls the Initialize method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#initialize
func (g *GradientBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	g.p.Call("initialize", args...)
}

// GradientBlockInitializeDefinesOpts contains optional parameters for GradientBlock.InitializeDefines.
type GradientBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#initializedefines
func (g *GradientBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *GradientBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &GradientBlockInitializeDefinesOpts{}
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

	g.p.Call("initializeDefines", args...)
}

// GradientBlockIsReadyOpts contains optional parameters for GradientBlock.IsReady.
type GradientBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isready
func (g *GradientBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *GradientBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &GradientBlockIsReadyOpts{}
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

	retVal := g.p.Call("isReady", args...)
	return retVal.Bool()
}

// GradientBlockPrepareDefinesOpts contains optional parameters for GradientBlock.PrepareDefines.
type GradientBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#preparedefines
func (g *GradientBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *GradientBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &GradientBlockPrepareDefinesOpts{}
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

	g.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#providefallbacks
func (g *GradientBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	g.p.Call("provideFallbacks", args...)
}

// GradientBlockRegisterInputOpts contains optional parameters for GradientBlock.RegisterInput.
type GradientBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#registerinput
func (g *GradientBlock) RegisterInput(name string, jsType js.Value, opts *GradientBlockRegisterInputOpts) *GradientBlock {
	if opts == nil {
		opts = &GradientBlockRegisterInputOpts{}
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

	retVal := g.p.Call("registerInput", args...)
	return GradientBlockFromJSObject(retVal, g.ctx)
}

// GradientBlockRegisterOutputOpts contains optional parameters for GradientBlock.RegisterOutput.
type GradientBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#registeroutput
func (g *GradientBlock) RegisterOutput(name string, jsType js.Value, opts *GradientBlockRegisterOutputOpts) *GradientBlock {
	if opts == nil {
		opts = &GradientBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}

	retVal := g.p.Call("registerOutput", args...)
	return GradientBlockFromJSObject(retVal, g.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#replacerepeatablecontent
func (g *GradientBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	g.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#serialize
func (g *GradientBlock) Serialize() interface{} {

	retVal := g.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#updateuniformsandsamples
func (g *GradientBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	g.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#_deserialize
func (g *GradientBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	g.p.Call("_deserialize", args...)
}

/*

// BuildId returns the BuildId property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#buildid
func (g *GradientBlock) BuildId(buildId float64) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(buildId)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#buildid
func (g *GradientBlock) SetBuildId(buildId float64) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(buildId)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// ColorSteps returns the ColorSteps property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#colorsteps
func (g *GradientBlock) ColorSteps(colorSteps *GradientBlockColorStep) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(colorSteps.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetColorSteps sets the ColorSteps property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#colorsteps
func (g *GradientBlock) SetColorSteps(colorSteps *GradientBlockColorStep) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(colorSteps.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#comments
func (g *GradientBlock) Comments(comments string) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(comments)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#comments
func (g *GradientBlock) SetComments(comments string) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(comments)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Gradient returns the Gradient property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#gradient
func (g *GradientBlock) Gradient(gradient *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(gradient.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetGradient sets the Gradient property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#gradient
func (g *GradientBlock) SetGradient(gradient *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(gradient.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#inputs
func (g *GradientBlock) Inputs(inputs *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(inputs.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#inputs
func (g *GradientBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(inputs.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isfinalmerger
func (g *GradientBlock) IsFinalMerger(isFinalMerger bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isFinalMerger)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isfinalmerger
func (g *GradientBlock) SetIsFinalMerger(isFinalMerger bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isFinalMerger)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isinput
func (g *GradientBlock) IsInput(isInput bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isInput)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isinput
func (g *GradientBlock) SetIsInput(isInput bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isInput)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isunique
func (g *GradientBlock) IsUnique(isUnique bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isUnique)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#isunique
func (g *GradientBlock) SetIsUnique(isUnique bool) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(isUnique)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#name
func (g *GradientBlock) Name(name string) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(name)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#name
func (g *GradientBlock) SetName(name string) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(name)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#output
func (g *GradientBlock) Output(output *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(output.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#output
func (g *GradientBlock) SetOutput(output *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(output.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#outputs
func (g *GradientBlock) Outputs(outputs *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(outputs.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#outputs
func (g *GradientBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(outputs.JSObject())
	return GradientBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#target
func (g *GradientBlock) Target(target js.Value) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(target)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#target
func (g *GradientBlock) SetTarget(target js.Value) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(target)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#uniqueid
func (g *GradientBlock) UniqueId(uniqueId float64) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(uniqueId)
	return GradientBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class GradientBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock#uniqueid
func (g *GradientBlock) SetUniqueId(uniqueId float64) *GradientBlock {
	p := ba.ctx.Get("GradientBlock").New(uniqueId)
	return GradientBlockFromJSObject(p, ba.ctx)
}

*/
