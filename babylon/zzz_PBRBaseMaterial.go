// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRBaseMaterial represents a babylon.js PBRBaseMaterial.
// The Physically based material base class of BJS.
//
// This offers the main features of a standard PBR material.
// For more information, please refer to the documentation :
// <a href="https://doc.babylonjs.com/how_to/physically_based_rendering">https://doc.babylonjs.com/how_to/physically_based_rendering</a>
type PBRBaseMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRBaseMaterial) JSObject() js.Value { return p.p }

// PBRBaseMaterial returns a PBRBaseMaterial JavaScript class.
func (ba *Babylon) PBRBaseMaterial() *PBRBaseMaterial {
	p := ba.ctx.Get("PBRBaseMaterial")
	return PBRBaseMaterialFromJSObject(p, ba.ctx)
}

// PBRBaseMaterialFromJSObject returns a wrapped PBRBaseMaterial JavaScript class.
func PBRBaseMaterialFromJSObject(p js.Value, ctx js.Value) *PBRBaseMaterial {
	return &PBRBaseMaterial{p: p, ctx: ctx}
}

// PBRBaseMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func PBRBaseMaterialArrayToJSArray(array []*PBRBaseMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPBRBaseMaterial returns a new PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#constructor
func (ba *Babylon) NewPBRBaseMaterial(name string, scene *Scene) *PBRBaseMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("PBRBaseMaterial").New(args...)
	return PBRBaseMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#bindforsubmesh
func (p *PBRBaseMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	p.p.Call("bindForSubMesh", args...)
}

// BuildUniformLayout calls the BuildUniformLayout method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#builduniformlayout
func (p *PBRBaseMaterial) BuildUniformLayout() {

	p.p.Call("buildUniformLayout")
}

// PBRBaseMaterialDisposeOpts contains optional parameters for PBRBaseMaterial.Dispose.
type PBRBaseMaterialDisposeOpts struct {
	ForceDisposeEffect   *bool
	ForceDisposeTextures *bool
}

// Dispose calls the Dispose method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#dispose
func (p *PBRBaseMaterial) Dispose(opts *PBRBaseMaterialDisposeOpts) {
	if opts == nil {
		opts = &PBRBaseMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}
	if opts.ForceDisposeTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeTextures)
	}

	p.p.Call("dispose", args...)
}

// PBRBaseMaterialForceCompilationOpts contains optional parameters for PBRBaseMaterial.ForceCompilation.
type PBRBaseMaterialForceCompilationOpts struct {
	OnCompiled JSFunc
	Options    js.Value
}

// ForceCompilation calls the ForceCompilation method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#forcecompilation
func (p *PBRBaseMaterial) ForceCompilation(mesh *AbstractMesh, opts *PBRBaseMaterialForceCompilationOpts) {
	if opts == nil {
		opts = &PBRBaseMaterialForceCompilationOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnCompiled) /* never freed! */)
	}
	args = append(args, opts.Options)

	p.p.Call("forceCompilation", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#getactivetextures
func (p *PBRBaseMaterial) GetActiveTextures() []*BaseTexture {

	retVal := p.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#getalphatesttexture
func (p *PBRBaseMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := p.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, p.ctx)
}

// GetAnimatables calls the GetAnimatables method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#getanimatables
func (p *PBRBaseMaterial) GetAnimatables() []*IAnimatable {

	retVal := p.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#getclassname
func (p *PBRBaseMaterial) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#hastexture
func (p *PBRBaseMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := p.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// IsMetallicWorkflow calls the IsMetallicWorkflow method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#ismetallicworkflow
func (p *PBRBaseMaterial) IsMetallicWorkflow() bool {

	retVal := p.p.Call("isMetallicWorkflow")
	return retVal.Bool()
}

// PBRBaseMaterialIsReadyForSubMeshOpts contains optional parameters for PBRBaseMaterial.IsReadyForSubMesh.
type PBRBaseMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#isreadyforsubmesh
func (p *PBRBaseMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *PBRBaseMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &PBRBaseMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := p.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#needalphablending
func (p *PBRBaseMaterial) NeedAlphaBlending() bool {

	retVal := p.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaBlendingForMesh calls the NeedAlphaBlendingForMesh method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#needalphablendingformesh
func (p *PBRBaseMaterial) NeedAlphaBlendingForMesh(mesh *AbstractMesh) bool {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	retVal := p.p.Call("needAlphaBlendingForMesh", args...)
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#needalphatesting
func (p *PBRBaseMaterial) NeedAlphaTesting() bool {

	retVal := p.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Unbind calls the Unbind method on the PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#unbind
func (p *PBRBaseMaterial) Unbind() {

	p.p.Call("unbind")
}

// Anisotropy returns the Anisotropy property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#anisotropy
func (p *PBRBaseMaterial) Anisotropy() *PBRAnisotropicConfiguration {
	retVal := p.p.Get("anisotropy")
	return PBRAnisotropicConfigurationFromJSObject(retVal, p.ctx)
}

// SetAnisotropy sets the Anisotropy property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#anisotropy
func (p *PBRBaseMaterial) SetAnisotropy(anisotropy *PBRAnisotropicConfiguration) *PBRBaseMaterial {
	p.p.Set("anisotropy", anisotropy.JSObject())
	return p
}

// Brdf returns the Brdf property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#brdf
func (p *PBRBaseMaterial) Brdf() *PBRBRDFConfiguration {
	retVal := p.p.Get("brdf")
	return PBRBRDFConfigurationFromJSObject(retVal, p.ctx)
}

// SetBrdf sets the Brdf property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#brdf
func (p *PBRBaseMaterial) SetBrdf(brdf *PBRBRDFConfiguration) *PBRBaseMaterial {
	p.p.Set("brdf", brdf.JSObject())
	return p
}

// ClearCoat returns the ClearCoat property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#clearcoat
func (p *PBRBaseMaterial) ClearCoat() *PBRClearCoatConfiguration {
	retVal := p.p.Get("clearCoat")
	return PBRClearCoatConfigurationFromJSObject(retVal, p.ctx)
}

// SetClearCoat sets the ClearCoat property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#clearcoat
func (p *PBRBaseMaterial) SetClearCoat(clearCoat *PBRClearCoatConfiguration) *PBRBaseMaterial {
	p.p.Set("clearCoat", clearCoat.JSObject())
	return p
}

// CustomShaderNameResolve returns the CustomShaderNameResolve property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#customshadernameresolve
func (p *PBRBaseMaterial) CustomShaderNameResolve() js.Value {
	retVal := p.p.Get("customShaderNameResolve")
	return retVal
}

// SetCustomShaderNameResolve sets the CustomShaderNameResolve property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#customshadernameresolve
func (p *PBRBaseMaterial) SetCustomShaderNameResolve(customShaderNameResolve JSFunc) *PBRBaseMaterial {
	p.p.Set("customShaderNameResolve", js.FuncOf(customShaderNameResolve))
	return p
}

// DEFAULT_AO_ON_ANALYTICAL_LIGHTS returns the DEFAULT_AO_ON_ANALYTICAL_LIGHTS property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#default_ao_on_analytical_lights
func (p *PBRBaseMaterial) DEFAULT_AO_ON_ANALYTICAL_LIGHTS() float64 {
	retVal := p.p.Get("DEFAULT_AO_ON_ANALYTICAL_LIGHTS")
	return retVal.Float()
}

// SetDEFAULT_AO_ON_ANALYTICAL_LIGHTS sets the DEFAULT_AO_ON_ANALYTICAL_LIGHTS property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#default_ao_on_analytical_lights
func (p *PBRBaseMaterial) SetDEFAULT_AO_ON_ANALYTICAL_LIGHTS(DEFAULT_AO_ON_ANALYTICAL_LIGHTS float64) *PBRBaseMaterial {
	p.p.Set("DEFAULT_AO_ON_ANALYTICAL_LIGHTS", DEFAULT_AO_ON_ANALYTICAL_LIGHTS)
	return p
}

// HasRenderTargetTextures returns the HasRenderTargetTextures property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#hasrendertargettextures
func (p *PBRBaseMaterial) HasRenderTargetTextures() bool {
	retVal := p.p.Get("hasRenderTargetTextures")
	return retVal.Bool()
}

// SetHasRenderTargetTextures sets the HasRenderTargetTextures property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#hasrendertargettextures
func (p *PBRBaseMaterial) SetHasRenderTargetTextures(hasRenderTargetTextures bool) *PBRBaseMaterial {
	p.p.Set("hasRenderTargetTextures", hasRenderTargetTextures)
	return p
}

// LIGHTFALLOFF_GLTF returns the LIGHTFALLOFF_GLTF property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_gltf
func (p *PBRBaseMaterial) LIGHTFALLOFF_GLTF() float64 {
	retVal := p.p.Get("LIGHTFALLOFF_GLTF")
	return retVal.Float()
}

// SetLIGHTFALLOFF_GLTF sets the LIGHTFALLOFF_GLTF property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_gltf
func (p *PBRBaseMaterial) SetLIGHTFALLOFF_GLTF(LIGHTFALLOFF_GLTF float64) *PBRBaseMaterial {
	p.p.Set("LIGHTFALLOFF_GLTF", LIGHTFALLOFF_GLTF)
	return p
}

// LIGHTFALLOFF_PHYSICAL returns the LIGHTFALLOFF_PHYSICAL property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_physical
func (p *PBRBaseMaterial) LIGHTFALLOFF_PHYSICAL() float64 {
	retVal := p.p.Get("LIGHTFALLOFF_PHYSICAL")
	return retVal.Float()
}

// SetLIGHTFALLOFF_PHYSICAL sets the LIGHTFALLOFF_PHYSICAL property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_physical
func (p *PBRBaseMaterial) SetLIGHTFALLOFF_PHYSICAL(LIGHTFALLOFF_PHYSICAL float64) *PBRBaseMaterial {
	p.p.Set("LIGHTFALLOFF_PHYSICAL", LIGHTFALLOFF_PHYSICAL)
	return p
}

// LIGHTFALLOFF_STANDARD returns the LIGHTFALLOFF_STANDARD property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_standard
func (p *PBRBaseMaterial) LIGHTFALLOFF_STANDARD() float64 {
	retVal := p.p.Get("LIGHTFALLOFF_STANDARD")
	return retVal.Float()
}

// SetLIGHTFALLOFF_STANDARD sets the LIGHTFALLOFF_STANDARD property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#lightfalloff_standard
func (p *PBRBaseMaterial) SetLIGHTFALLOFF_STANDARD(LIGHTFALLOFF_STANDARD float64) *PBRBaseMaterial {
	p.p.Set("LIGHTFALLOFF_STANDARD", LIGHTFALLOFF_STANDARD)
	return p
}

// PBRMATERIAL_ALPHABLEND returns the PBRMATERIAL_ALPHABLEND property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphablend
func (p *PBRBaseMaterial) PBRMATERIAL_ALPHABLEND() float64 {
	retVal := p.p.Get("PBRMATERIAL_ALPHABLEND")
	return retVal.Float()
}

// SetPBRMATERIAL_ALPHABLEND sets the PBRMATERIAL_ALPHABLEND property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphablend
func (p *PBRBaseMaterial) SetPBRMATERIAL_ALPHABLEND(PBRMATERIAL_ALPHABLEND float64) *PBRBaseMaterial {
	p.p.Set("PBRMATERIAL_ALPHABLEND", PBRMATERIAL_ALPHABLEND)
	return p
}

// PBRMATERIAL_ALPHATEST returns the PBRMATERIAL_ALPHATEST property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphatest
func (p *PBRBaseMaterial) PBRMATERIAL_ALPHATEST() float64 {
	retVal := p.p.Get("PBRMATERIAL_ALPHATEST")
	return retVal.Float()
}

// SetPBRMATERIAL_ALPHATEST sets the PBRMATERIAL_ALPHATEST property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphatest
func (p *PBRBaseMaterial) SetPBRMATERIAL_ALPHATEST(PBRMATERIAL_ALPHATEST float64) *PBRBaseMaterial {
	p.p.Set("PBRMATERIAL_ALPHATEST", PBRMATERIAL_ALPHATEST)
	return p
}

// PBRMATERIAL_ALPHATESTANDBLEND returns the PBRMATERIAL_ALPHATESTANDBLEND property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphatestandblend
func (p *PBRBaseMaterial) PBRMATERIAL_ALPHATESTANDBLEND() float64 {
	retVal := p.p.Get("PBRMATERIAL_ALPHATESTANDBLEND")
	return retVal.Float()
}

// SetPBRMATERIAL_ALPHATESTANDBLEND sets the PBRMATERIAL_ALPHATESTANDBLEND property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_alphatestandblend
func (p *PBRBaseMaterial) SetPBRMATERIAL_ALPHATESTANDBLEND(PBRMATERIAL_ALPHATESTANDBLEND float64) *PBRBaseMaterial {
	p.p.Set("PBRMATERIAL_ALPHATESTANDBLEND", PBRMATERIAL_ALPHATESTANDBLEND)
	return p
}

// PBRMATERIAL_OPAQUE returns the PBRMATERIAL_OPAQUE property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_opaque
func (p *PBRBaseMaterial) PBRMATERIAL_OPAQUE() float64 {
	retVal := p.p.Get("PBRMATERIAL_OPAQUE")
	return retVal.Float()
}

// SetPBRMATERIAL_OPAQUE sets the PBRMATERIAL_OPAQUE property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#pbrmaterial_opaque
func (p *PBRBaseMaterial) SetPBRMATERIAL_OPAQUE(PBRMATERIAL_OPAQUE float64) *PBRBaseMaterial {
	p.p.Set("PBRMATERIAL_OPAQUE", PBRMATERIAL_OPAQUE)
	return p
}

// Sheen returns the Sheen property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#sheen
func (p *PBRBaseMaterial) Sheen() *PBRSheenConfiguration {
	retVal := p.p.Get("sheen")
	return PBRSheenConfigurationFromJSObject(retVal, p.ctx)
}

// SetSheen sets the Sheen property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#sheen
func (p *PBRBaseMaterial) SetSheen(sheen *PBRSheenConfiguration) *PBRBaseMaterial {
	p.p.Set("sheen", sheen.JSObject())
	return p
}

// SubSurface returns the SubSurface property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#subsurface
func (p *PBRBaseMaterial) SubSurface() *PBRSubSurfaceConfiguration {
	retVal := p.p.Get("subSurface")
	return PBRSubSurfaceConfigurationFromJSObject(retVal, p.ctx)
}

// SetSubSurface sets the SubSurface property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#subsurface
func (p *PBRBaseMaterial) SetSubSurface(subSurface *PBRSubSurfaceConfiguration) *PBRBaseMaterial {
	p.p.Set("subSurface", subSurface.JSObject())
	return p
}

// TransparencyMode returns the TransparencyMode property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#transparencymode
func (p *PBRBaseMaterial) TransparencyMode() float64 {
	retVal := p.p.Get("transparencyMode")
	return retVal.Float()
}

// SetTransparencyMode sets the TransparencyMode property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#transparencymode
func (p *PBRBaseMaterial) SetTransparencyMode(transparencyMode float64) *PBRBaseMaterial {
	p.p.Set("transparencyMode", transparencyMode)
	return p
}

// UseLogarithmicDepth returns the UseLogarithmicDepth property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#uselogarithmicdepth
func (p *PBRBaseMaterial) UseLogarithmicDepth() bool {
	retVal := p.p.Get("useLogarithmicDepth")
	return retVal.Bool()
}

// SetUseLogarithmicDepth sets the UseLogarithmicDepth property of class PBRBaseMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial#uselogarithmicdepth
func (p *PBRBaseMaterial) SetUseLogarithmicDepth(useLogarithmicDepth bool) *PBRBaseMaterial {
	p.p.Set("useLogarithmicDepth", useLogarithmicDepth)
	return p
}
