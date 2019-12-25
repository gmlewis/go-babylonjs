// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRAnisotropicConfiguration represents a babylon.js PBRAnisotropicConfiguration.
// Define the code related to the anisotropic parameters of the pbr material.
type PBRAnisotropicConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRAnisotropicConfiguration) JSObject() js.Value { return p.p }

// PBRAnisotropicConfiguration returns a PBRAnisotropicConfiguration JavaScript class.
func (ba *Babylon) PBRAnisotropicConfiguration() *PBRAnisotropicConfiguration {
	p := ba.ctx.Get("PBRAnisotropicConfiguration")
	return PBRAnisotropicConfigurationFromJSObject(p, ba.ctx)
}

// PBRAnisotropicConfigurationFromJSObject returns a wrapped PBRAnisotropicConfiguration JavaScript class.
func PBRAnisotropicConfigurationFromJSObject(p js.Value, ctx js.Value) *PBRAnisotropicConfiguration {
	return &PBRAnisotropicConfiguration{p: p, ctx: ctx}
}

// PBRAnisotropicConfigurationArrayToJSArray returns a JavaScript Array for the wrapped array.
func PBRAnisotropicConfigurationArrayToJSArray(array []*PBRAnisotropicConfiguration) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPBRAnisotropicConfiguration returns a new PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#constructor
func (ba *Babylon) NewPBRAnisotropicConfiguration(markAllSubMeshesAsTexturesDirty JSFunc) *PBRAnisotropicConfiguration {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(markAllSubMeshesAsTexturesDirty))

	p := ba.ctx.Get("PBRAnisotropicConfiguration").New(args...)
	return PBRAnisotropicConfigurationFromJSObject(p, ba.ctx)
}

// AddFallbacks calls the AddFallbacks method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#addfallbacks
func (p *PBRAnisotropicConfiguration) AddFallbacks(defines js.Value, fallbacks *EffectFallbacks, currentRank float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, defines)

	if fallbacks == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, fallbacks.JSObject())
	}

	args = append(args, currentRank)

	retVal := p.p.Call("AddFallbacks", args...)
	return retVal.Float()
}

// AddSamplers calls the AddSamplers method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#addsamplers
func (p *PBRAnisotropicConfiguration) AddSamplers(samplers []string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, samplers)

	p.p.Call("AddSamplers", args...)
}

// AddUniforms calls the AddUniforms method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#adduniforms
func (p *PBRAnisotropicConfiguration) AddUniforms(uniforms []string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uniforms)

	p.p.Call("AddUniforms", args...)
}

// BindForSubMesh calls the BindForSubMesh method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#bindforsubmesh
func (p *PBRAnisotropicConfiguration) BindForSubMesh(uniformBuffer *UniformBuffer, scene *Scene, isFrozen bool) {

	args := make([]interface{}, 0, 3+0)

	if uniformBuffer == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, uniformBuffer.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, isFrozen)

	p.p.Call("bindForSubMesh", args...)
}

// CopyTo calls the CopyTo method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#copyto
func (p *PBRAnisotropicConfiguration) CopyTo(anisotropicConfiguration *PBRAnisotropicConfiguration) {

	args := make([]interface{}, 0, 1+0)

	if anisotropicConfiguration == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, anisotropicConfiguration.JSObject())
	}

	p.p.Call("copyTo", args...)
}

// PBRAnisotropicConfigurationDisposeOpts contains optional parameters for PBRAnisotropicConfiguration.Dispose.
type PBRAnisotropicConfigurationDisposeOpts struct {
	ForceDisposeTextures *bool
}

// Dispose calls the Dispose method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#dispose
func (p *PBRAnisotropicConfiguration) Dispose(opts *PBRAnisotropicConfigurationDisposeOpts) {
	if opts == nil {
		opts = &PBRAnisotropicConfigurationDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeTextures)
	}

	p.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#getactivetextures
func (p *PBRAnisotropicConfiguration) GetActiveTextures(activeTextures []*BaseTexture) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, BaseTextureArrayToJSArray(activeTextures))

	p.p.Call("getActiveTextures", args...)
}

// GetAnimatables calls the GetAnimatables method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#getanimatables
func (p *PBRAnisotropicConfiguration) GetAnimatables(animatables []*IAnimatable) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, IAnimatableArrayToJSArray(animatables))

	p.p.Call("getAnimatables", args...)
}

// GetClassName calls the GetClassName method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#getclassname
func (p *PBRAnisotropicConfiguration) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#hastexture
func (p *PBRAnisotropicConfiguration) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := p.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#isreadyforsubmesh
func (p *PBRAnisotropicConfiguration) IsReadyForSubMesh(defines js.Value, scene *Scene) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := p.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// Parse calls the Parse method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#parse
func (p *PBRAnisotropicConfiguration) Parse(source JSObject, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	p.p.Call("parse", args...)
}

// PrepareDefines calls the PrepareDefines method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#preparedefines
func (p *PBRAnisotropicConfiguration) PrepareDefines(defines js.Value, mesh *AbstractMesh, scene *Scene) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, defines)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	p.p.Call("prepareDefines", args...)
}

// PrepareUniformBuffer calls the PrepareUniformBuffer method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#prepareuniformbuffer
func (p *PBRAnisotropicConfiguration) PrepareUniformBuffer(uniformBuffer *UniformBuffer) {

	args := make([]interface{}, 0, 1+0)

	if uniformBuffer == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, uniformBuffer.JSObject())
	}

	p.p.Call("PrepareUniformBuffer", args...)
}

// Serialize calls the Serialize method on the PBRAnisotropicConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#serialize
func (p *PBRAnisotropicConfiguration) Serialize() js.Value {

	retVal := p.p.Call("serialize")
	return retVal
}

// Direction returns the Direction property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#direction
func (p *PBRAnisotropicConfiguration) Direction() *Vector2 {
	retVal := p.p.Get("direction")
	return Vector2FromJSObject(retVal, p.ctx)
}

// SetDirection sets the Direction property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#direction
func (p *PBRAnisotropicConfiguration) SetDirection(direction *Vector2) *PBRAnisotropicConfiguration {
	p.p.Set("direction", direction.JSObject())
	return p
}

// Intensity returns the Intensity property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#intensity
func (p *PBRAnisotropicConfiguration) Intensity() float64 {
	retVal := p.p.Get("intensity")
	return retVal.Float()
}

// SetIntensity sets the Intensity property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#intensity
func (p *PBRAnisotropicConfiguration) SetIntensity(intensity float64) *PBRAnisotropicConfiguration {
	p.p.Set("intensity", intensity)
	return p
}

// IsEnabled returns the IsEnabled property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#isenabled
func (p *PBRAnisotropicConfiguration) IsEnabled() bool {
	retVal := p.p.Get("isEnabled")
	return retVal.Bool()
}

// SetIsEnabled sets the IsEnabled property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#isenabled
func (p *PBRAnisotropicConfiguration) SetIsEnabled(isEnabled bool) *PBRAnisotropicConfiguration {
	p.p.Set("isEnabled", isEnabled)
	return p
}

// Texture returns the Texture property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#texture
func (p *PBRAnisotropicConfiguration) Texture() *BaseTexture {
	retVal := p.p.Get("texture")
	return BaseTextureFromJSObject(retVal, p.ctx)
}

// SetTexture sets the Texture property of class PBRAnisotropicConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbranisotropicconfiguration#texture
func (p *PBRAnisotropicConfiguration) SetTexture(texture *BaseTexture) *PBRAnisotropicConfiguration {
	p.p.Set("texture", texture.JSObject())
	return p
}
