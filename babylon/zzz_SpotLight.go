// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpotLight represents a babylon.js SpotLight.
// A spot light is defined by a position, a direction, an angle, and an exponent.
// These values define a cone of light starting from the position, emitting toward the direction.
// The angle, in radians, defines the size (field of illumination) of the spotlight&#39;s conical beam,
// and the exponent defines the speed of the decay of the light with distance (reach).
// Documentation: <a href="https://doc.babylonjs.com/babylon101/lights">https://doc.babylonjs.com/babylon101/lights</a>
type SpotLight struct {
	*ShadowLight
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SpotLight) JSObject() js.Value { return s.p }

// SpotLight returns a SpotLight JavaScript class.
func (ba *Babylon) SpotLight() *SpotLight {
	p := ba.ctx.Get("SpotLight")
	return SpotLightFromJSObject(p, ba.ctx)
}

// SpotLightFromJSObject returns a wrapped SpotLight JavaScript class.
func SpotLightFromJSObject(p js.Value, ctx js.Value) *SpotLight {
	return &SpotLight{ShadowLight: ShadowLightFromJSObject(p, ctx), ctx: ctx}
}

// SpotLightArrayToJSArray returns a JavaScript Array for the wrapped array.
func SpotLightArrayToJSArray(array []*SpotLight) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSpotLight returns a new SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight
func (ba *Babylon) NewSpotLight(name string, position *Vector3, direction *Vector3, angle float64, exponent float64, scene *Scene) *SpotLight {

	args := make([]interface{}, 0, 6+0)

	args = append(args, name)
	args = append(args, position.JSObject())
	args = append(args, direction.JSObject())
	args = append(args, angle)
	args = append(args, exponent)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("SpotLight").New(args...)
	return SpotLightFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#dispose
func (s *SpotLight) Dispose() {

	s.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#getclassname
func (s *SpotLight) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// GetTypeID calls the GetTypeID method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#gettypeid
func (s *SpotLight) GetTypeID() float64 {

	retVal := s.p.Call("getTypeID")
	return retVal.Float()
}

// PrepareLightSpecificDefines calls the PrepareLightSpecificDefines method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#preparelightspecificdefines
func (s *SpotLight) PrepareLightSpecificDefines(defines interface{}, lightIndex float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)
	args = append(args, lightIndex)

	s.p.Call("prepareLightSpecificDefines", args...)
}

// TransferTexturesToEffect calls the TransferTexturesToEffect method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#transfertexturestoeffect
func (s *SpotLight) TransferTexturesToEffect(effect *Effect, lightIndex string) *Light {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightIndex)

	retVal := s.p.Call("transferTexturesToEffect", args...)
	return LightFromJSObject(retVal, s.ctx)
}

// TransferToEffect calls the TransferToEffect method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#transfertoeffect
func (s *SpotLight) TransferToEffect(effect *Effect, lightIndex string) *SpotLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightIndex)

	retVal := s.p.Call("transferToEffect", args...)
	return SpotLightFromJSObject(retVal, s.ctx)
}

// TransferToNodeMaterialEffect calls the TransferToNodeMaterialEffect method on the SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#transfertonodematerialeffect
func (s *SpotLight) TransferToNodeMaterialEffect(effect *Effect, lightDataUniformName string) *SpotLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightDataUniformName)

	retVal := s.p.Call("transferToNodeMaterialEffect", args...)
	return SpotLightFromJSObject(retVal, s.ctx)
}

/*

// Angle returns the Angle property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#angle
func (s *SpotLight) Angle(angle float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(angle)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetAngle sets the Angle property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#angle
func (s *SpotLight) SetAngle(angle float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(angle)
	return SpotLightFromJSObject(p, ba.ctx)
}

// Exponent returns the Exponent property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#exponent
func (s *SpotLight) Exponent(exponent float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(exponent)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetExponent sets the Exponent property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#exponent
func (s *SpotLight) SetExponent(exponent float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(exponent)
	return SpotLightFromJSObject(p, ba.ctx)
}

// InnerAngle returns the InnerAngle property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#innerangle
func (s *SpotLight) InnerAngle(innerAngle float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(innerAngle)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetInnerAngle sets the InnerAngle property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#innerangle
func (s *SpotLight) SetInnerAngle(innerAngle float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(innerAngle)
	return SpotLightFromJSObject(p, ba.ctx)
}

// ProjectionTexture returns the ProjectionTexture property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexture
func (s *SpotLight) ProjectionTexture(projectionTexture *BaseTexture) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTexture.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetProjectionTexture sets the ProjectionTexture property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexture
func (s *SpotLight) SetProjectionTexture(projectionTexture *BaseTexture) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTexture.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// ProjectionTextureLightFar returns the ProjectionTextureLightFar property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturelightfar
func (s *SpotLight) ProjectionTextureLightFar(projectionTextureLightFar float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureLightFar)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetProjectionTextureLightFar sets the ProjectionTextureLightFar property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturelightfar
func (s *SpotLight) SetProjectionTextureLightFar(projectionTextureLightFar float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureLightFar)
	return SpotLightFromJSObject(p, ba.ctx)
}

// ProjectionTextureLightNear returns the ProjectionTextureLightNear property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturelightnear
func (s *SpotLight) ProjectionTextureLightNear(projectionTextureLightNear float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureLightNear)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetProjectionTextureLightNear sets the ProjectionTextureLightNear property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturelightnear
func (s *SpotLight) SetProjectionTextureLightNear(projectionTextureLightNear float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureLightNear)
	return SpotLightFromJSObject(p, ba.ctx)
}

// ProjectionTextureMatrix returns the ProjectionTextureMatrix property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturematrix
func (s *SpotLight) ProjectionTextureMatrix(projectionTextureMatrix *Matrix) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureMatrix.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetProjectionTextureMatrix sets the ProjectionTextureMatrix property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontexturematrix
func (s *SpotLight) SetProjectionTextureMatrix(projectionTextureMatrix *Matrix) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureMatrix.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// ProjectionTextureUpDirection returns the ProjectionTextureUpDirection property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontextureupdirection
func (s *SpotLight) ProjectionTextureUpDirection(projectionTextureUpDirection *Vector3) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureUpDirection.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetProjectionTextureUpDirection sets the ProjectionTextureUpDirection property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#projectiontextureupdirection
func (s *SpotLight) SetProjectionTextureUpDirection(projectionTextureUpDirection *Vector3) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(projectionTextureUpDirection.JSObject())
	return SpotLightFromJSObject(p, ba.ctx)
}

// ShadowAngleScale returns the ShadowAngleScale property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#shadowanglescale
func (s *SpotLight) ShadowAngleScale(shadowAngleScale float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(shadowAngleScale)
	return SpotLightFromJSObject(p, ba.ctx)
}

// SetShadowAngleScale sets the ShadowAngleScale property of class SpotLight.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight#shadowanglescale
func (s *SpotLight) SetShadowAngleScale(shadowAngleScale float64) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(shadowAngleScale)
	return SpotLightFromJSObject(p, ba.ctx)
}

*/
