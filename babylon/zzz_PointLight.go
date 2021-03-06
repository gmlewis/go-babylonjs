// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointLight represents a babylon.js PointLight.
// A point light is a light defined by an unique point in world space.
// The light is emitted in every direction from this point.
// A good example of a point light is a standard light bulb.
// Documentation: <a href="https://doc.babylonjs.com/babylon101/lights">https://doc.babylonjs.com/babylon101/lights</a>
type PointLight struct {
	*ShadowLight
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PointLight) JSObject() js.Value { return p.p }

// PointLight returns a PointLight JavaScript class.
func (ba *Babylon) PointLight() *PointLight {
	p := ba.ctx.Get("PointLight")
	return PointLightFromJSObject(p, ba.ctx)
}

// PointLightFromJSObject returns a wrapped PointLight JavaScript class.
func PointLightFromJSObject(p js.Value, ctx js.Value) *PointLight {
	return &PointLight{ShadowLight: ShadowLightFromJSObject(p, ctx), ctx: ctx}
}

// PointLightArrayToJSArray returns a JavaScript Array for the wrapped array.
func PointLightArrayToJSArray(array []*PointLight) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPointLight returns a new PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#constructor
func (ba *Babylon) NewPointLight(name string, position *Vector3, scene *Scene) *PointLight {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("PointLight").New(args...)
	return PointLightFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#getclassname
func (p *PointLight) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// PointLightGetShadowDirectionOpts contains optional parameters for PointLight.GetShadowDirection.
type PointLightGetShadowDirectionOpts struct {
	FaceIndex *float64
}

// GetShadowDirection calls the GetShadowDirection method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#getshadowdirection
func (p *PointLight) GetShadowDirection(opts *PointLightGetShadowDirectionOpts) *Vector3 {
	if opts == nil {
		opts = &PointLightGetShadowDirectionOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.FaceIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FaceIndex)
	}

	retVal := p.p.Call("getShadowDirection", args...)
	return Vector3FromJSObject(retVal, p.ctx)
}

// GetTypeID calls the GetTypeID method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#gettypeid
func (p *PointLight) GetTypeID() float64 {

	retVal := p.p.Call("getTypeID")
	return retVal.Float()
}

// NeedCube calls the NeedCube method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#needcube
func (p *PointLight) NeedCube() bool {

	retVal := p.p.Call("needCube")
	return retVal.Bool()
}

// PrepareLightSpecificDefines calls the PrepareLightSpecificDefines method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#preparelightspecificdefines
func (p *PointLight) PrepareLightSpecificDefines(defines JSObject, lightIndex float64) {

	args := make([]interface{}, 0, 2+0)

	if defines == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, defines.JSObject())
	}

	args = append(args, lightIndex)

	p.p.Call("prepareLightSpecificDefines", args...)
}

// TransferToEffect calls the TransferToEffect method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#transfertoeffect
func (p *PointLight) TransferToEffect(effect *Effect, lightIndex string) *PointLight {

	args := make([]interface{}, 0, 2+0)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	args = append(args, lightIndex)

	retVal := p.p.Call("transferToEffect", args...)
	return PointLightFromJSObject(retVal, p.ctx)
}

// TransferToNodeMaterialEffect calls the TransferToNodeMaterialEffect method on the PointLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#transfertonodematerialeffect
func (p *PointLight) TransferToNodeMaterialEffect(effect *Effect, lightDataUniformName string) *PointLight {

	args := make([]interface{}, 0, 2+0)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	args = append(args, lightDataUniformName)

	retVal := p.p.Call("transferToNodeMaterialEffect", args...)
	return PointLightFromJSObject(retVal, p.ctx)
}

// Direction returns the Direction property of class PointLight.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#direction
func (p *PointLight) Direction() *Vector3 {
	retVal := p.p.Get("direction")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetDirection sets the Direction property of class PointLight.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#direction
func (p *PointLight) SetDirection(direction *Vector3) *PointLight {
	p.p.Set("direction", direction.JSObject())
	return p
}

// ShadowAngle returns the ShadowAngle property of class PointLight.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#shadowangle
func (p *PointLight) ShadowAngle() float64 {
	retVal := p.p.Get("shadowAngle")
	return retVal.Float()
}

// SetShadowAngle sets the ShadowAngle property of class PointLight.
//
// https://doc.babylonjs.com/api/classes/babylon.pointlight#shadowangle
func (p *PointLight) SetShadowAngle(shadowAngle float64) *PointLight {
	p.p.Set("shadowAngle", shadowAngle)
	return p
}
