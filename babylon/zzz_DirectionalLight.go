// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DirectionalLight represents a babylon.js DirectionalLight.
// A directional light is defined by a direction (what a surprise!).
// The light is emitted from everywhere in the specified direction, and has an infinite range.
// An example of a directional light is when a distance planet is lit by the apparently parallel lines of light from its sun. Light in a downward direction will light the top of an object.
// Documentation: <a href="https://doc.babylonjs.com/babylon101/lights">https://doc.babylonjs.com/babylon101/lights</a>
type DirectionalLight struct {
	*ShadowLight
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DirectionalLight) JSObject() js.Value { return d.p }

// DirectionalLight returns a DirectionalLight JavaScript class.
func (ba *Babylon) DirectionalLight() *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight")
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// DirectionalLightFromJSObject returns a wrapped DirectionalLight JavaScript class.
func DirectionalLightFromJSObject(p js.Value, ctx js.Value) *DirectionalLight {
	return &DirectionalLight{ShadowLight: ShadowLightFromJSObject(p, ctx), ctx: ctx}
}

// DirectionalLightArrayToJSArray returns a JavaScript Array for the wrapped array.
func DirectionalLightArrayToJSArray(array []*DirectionalLight) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDirectionalLight returns a new DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight
func (ba *Babylon) NewDirectionalLight(name string, direction *Vector3, scene *Scene) *DirectionalLight {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, direction.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("DirectionalLight").New(args...)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#getclassname
func (d *DirectionalLight) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// GetDepthMaxZ calls the GetDepthMaxZ method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#getdepthmaxz
func (d *DirectionalLight) GetDepthMaxZ(activeCamera *Camera) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, activeCamera.JSObject())

	retVal := d.p.Call("getDepthMaxZ", args...)
	return retVal.Float()
}

// GetDepthMinZ calls the GetDepthMinZ method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#getdepthminz
func (d *DirectionalLight) GetDepthMinZ(activeCamera *Camera) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, activeCamera.JSObject())

	retVal := d.p.Call("getDepthMinZ", args...)
	return retVal.Float()
}

// GetTypeID calls the GetTypeID method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#gettypeid
func (d *DirectionalLight) GetTypeID() float64 {

	retVal := d.p.Call("getTypeID")
	return retVal.Float()
}

// PrepareLightSpecificDefines calls the PrepareLightSpecificDefines method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#preparelightspecificdefines
func (d *DirectionalLight) PrepareLightSpecificDefines(defines interface{}, lightIndex float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)
	args = append(args, lightIndex)

	d.p.Call("prepareLightSpecificDefines", args...)
}

// TransferToEffect calls the TransferToEffect method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#transfertoeffect
func (d *DirectionalLight) TransferToEffect(effect *Effect, lightIndex string) *DirectionalLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightIndex)

	retVal := d.p.Call("transferToEffect", args...)
	return DirectionalLightFromJSObject(retVal, d.ctx)
}

// TransferToNodeMaterialEffect calls the TransferToNodeMaterialEffect method on the DirectionalLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#transfertonodematerialeffect
func (d *DirectionalLight) TransferToNodeMaterialEffect(effect *Effect, lightDataUniformName string) *Light {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightDataUniformName)

	retVal := d.p.Call("transferToNodeMaterialEffect", args...)
	return LightFromJSObject(retVal, d.ctx)
}

/*

// AutoUpdateExtends returns the AutoUpdateExtends property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#autoupdateextends
func (d *DirectionalLight) AutoUpdateExtends(autoUpdateExtends bool) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(autoUpdateExtends)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// SetAutoUpdateExtends sets the AutoUpdateExtends property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#autoupdateextends
func (d *DirectionalLight) SetAutoUpdateExtends(autoUpdateExtends bool) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(autoUpdateExtends)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// ShadowFrustumSize returns the ShadowFrustumSize property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#shadowfrustumsize
func (d *DirectionalLight) ShadowFrustumSize(shadowFrustumSize float64) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(shadowFrustumSize)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// SetShadowFrustumSize sets the ShadowFrustumSize property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#shadowfrustumsize
func (d *DirectionalLight) SetShadowFrustumSize(shadowFrustumSize float64) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(shadowFrustumSize)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// ShadowOrthoScale returns the ShadowOrthoScale property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#shadoworthoscale
func (d *DirectionalLight) ShadowOrthoScale(shadowOrthoScale float64) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(shadowOrthoScale)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

// SetShadowOrthoScale sets the ShadowOrthoScale property of class DirectionalLight.
//
// https://doc.babylonjs.com/api/classes/babylon.directionallight#shadoworthoscale
func (d *DirectionalLight) SetShadowOrthoScale(shadowOrthoScale float64) *DirectionalLight {
	p := ba.ctx.Get("DirectionalLight").New(shadowOrthoScale)
	return DirectionalLightFromJSObject(p, ba.ctx)
}

*/
