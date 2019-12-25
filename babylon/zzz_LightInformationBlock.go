// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LightInformationBlock represents a babylon.js LightInformationBlock.
// Block used to get data information from a light
type LightInformationBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LightInformationBlock) JSObject() js.Value { return l.p }

// LightInformationBlock returns a LightInformationBlock JavaScript class.
func (ba *Babylon) LightInformationBlock() *LightInformationBlock {
	p := ba.ctx.Get("LightInformationBlock")
	return LightInformationBlockFromJSObject(p, ba.ctx)
}

// LightInformationBlockFromJSObject returns a wrapped LightInformationBlock JavaScript class.
func LightInformationBlockFromJSObject(p js.Value, ctx js.Value) *LightInformationBlock {
	return &LightInformationBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// LightInformationBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func LightInformationBlockArrayToJSArray(array []*LightInformationBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLightInformationBlock returns a new LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#constructor
func (ba *Babylon) NewLightInformationBlock(name string) *LightInformationBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("LightInformationBlock").New(args...)
	return LightInformationBlockFromJSObject(p, ba.ctx)
}

// LightInformationBlockBindOpts contains optional parameters for LightInformationBlock.Bind.
type LightInformationBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#bind
func (l *LightInformationBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *LightInformationBlockBindOpts) {
	if opts == nil {
		opts = &LightInformationBlockBindOpts{}
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

	l.p.Call("bind", args...)
}

// GetClassName calls the GetClassName method on the LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#getclassname
func (l *LightInformationBlock) GetClassName() string {

	retVal := l.p.Call("getClassName")
	return retVal.String()
}

// PrepareDefines calls the PrepareDefines method on the LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#preparedefines
func (l *LightInformationBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

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

	l.p.Call("prepareDefines", args...)
}

// Serialize calls the Serialize method on the LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#serialize
func (l *LightInformationBlock) Serialize() js.Value {

	retVal := l.p.Call("serialize")
	return retVal
}

// _deserialize calls the _deserialize method on the LightInformationBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#_deserialize
func (l *LightInformationBlock) _deserialize(serializationObject JSObject, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	l.p.Call("_deserialize", args...)
}

// Color returns the Color property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#color
func (l *LightInformationBlock) Color() *NodeMaterialConnectionPoint {
	retVal := l.p.Get("color")
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// SetColor sets the Color property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#color
func (l *LightInformationBlock) SetColor(color *NodeMaterialConnectionPoint) *LightInformationBlock {
	l.p.Set("color", color.JSObject())
	return l
}

// Direction returns the Direction property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#direction
func (l *LightInformationBlock) Direction() *NodeMaterialConnectionPoint {
	retVal := l.p.Get("direction")
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// SetDirection sets the Direction property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#direction
func (l *LightInformationBlock) SetDirection(direction *NodeMaterialConnectionPoint) *LightInformationBlock {
	l.p.Set("direction", direction.JSObject())
	return l
}

// Intensity returns the Intensity property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#intensity
func (l *LightInformationBlock) Intensity() *NodeMaterialConnectionPoint {
	retVal := l.p.Get("intensity")
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// SetIntensity sets the Intensity property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#intensity
func (l *LightInformationBlock) SetIntensity(intensity *NodeMaterialConnectionPoint) *LightInformationBlock {
	l.p.Set("intensity", intensity.JSObject())
	return l
}

// Light returns the Light property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#light
func (l *LightInformationBlock) Light() *Light {
	retVal := l.p.Get("light")
	return LightFromJSObject(retVal, l.ctx)
}

// SetLight sets the Light property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#light
func (l *LightInformationBlock) SetLight(light *Light) *LightInformationBlock {
	l.p.Set("light", light.JSObject())
	return l
}

// WorldPosition returns the WorldPosition property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#worldposition
func (l *LightInformationBlock) WorldPosition() *NodeMaterialConnectionPoint {
	retVal := l.p.Get("worldPosition")
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// SetWorldPosition sets the WorldPosition property of class LightInformationBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lightinformationblock#worldposition
func (l *LightInformationBlock) SetWorldPosition(worldPosition *NodeMaterialConnectionPoint) *LightInformationBlock {
	l.p.Set("worldPosition", worldPosition.JSObject())
	return l
}
