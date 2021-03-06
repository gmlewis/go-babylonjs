// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageProcessingBlock represents a babylon.js ImageProcessingBlock.
// Block used to add image processing support to fragment shader
type ImageProcessingBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ImageProcessingBlock) JSObject() js.Value { return i.p }

// ImageProcessingBlock returns a ImageProcessingBlock JavaScript class.
func (ba *Babylon) ImageProcessingBlock() *ImageProcessingBlock {
	p := ba.ctx.Get("ImageProcessingBlock")
	return ImageProcessingBlockFromJSObject(p, ba.ctx)
}

// ImageProcessingBlockFromJSObject returns a wrapped ImageProcessingBlock JavaScript class.
func ImageProcessingBlockFromJSObject(p js.Value, ctx js.Value) *ImageProcessingBlock {
	return &ImageProcessingBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ImageProcessingBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ImageProcessingBlockArrayToJSArray(array []*ImageProcessingBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewImageProcessingBlock returns a new ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#constructor
func (ba *Babylon) NewImageProcessingBlock(name string) *ImageProcessingBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ImageProcessingBlock").New(args...)
	return ImageProcessingBlockFromJSObject(p, ba.ctx)
}

// ImageProcessingBlockBindOpts contains optional parameters for ImageProcessingBlock.Bind.
type ImageProcessingBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#bind
func (i *ImageProcessingBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *ImageProcessingBlockBindOpts) {
	if opts == nil {
		opts = &ImageProcessingBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	i.p.Call("bind", args...)
}

// GetClassName calls the GetClassName method on the ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#getclassname
func (i *ImageProcessingBlock) GetClassName() string {

	retVal := i.p.Call("getClassName")
	return retVal.String()
}

// Initialize calls the Initialize method on the ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#initialize
func (i *ImageProcessingBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	if state == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, state.JSObject())
	}

	i.p.Call("initialize", args...)
}

// IsReady calls the IsReady method on the ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#isready
func (i *ImageProcessingBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) bool {

	args := make([]interface{}, 0, 3+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	args = append(args, defines)

	retVal := i.p.Call("isReady", args...)
	return retVal.Bool()
}

// PrepareDefines calls the PrepareDefines method on the ImageProcessingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#preparedefines
func (i *ImageProcessingBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

	args := make([]interface{}, 0, 3+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	args = append(args, defines)

	i.p.Call("prepareDefines", args...)
}

// Color returns the Color property of class ImageProcessingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#color
func (i *ImageProcessingBlock) Color() *NodeMaterialConnectionPoint {
	retVal := i.p.Get("color")
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// SetColor sets the Color property of class ImageProcessingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#color
func (i *ImageProcessingBlock) SetColor(color *NodeMaterialConnectionPoint) *ImageProcessingBlock {
	i.p.Set("color", color.JSObject())
	return i
}

// Output returns the Output property of class ImageProcessingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#output
func (i *ImageProcessingBlock) Output() *NodeMaterialConnectionPoint {
	retVal := i.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, i.ctx)
}

// SetOutput sets the Output property of class ImageProcessingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingblock#output
func (i *ImageProcessingBlock) SetOutput(output *NodeMaterialConnectionPoint) *ImageProcessingBlock {
	i.p.Set("output", output.JSObject())
	return i
}
