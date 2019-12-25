// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TerrainMaterial represents a babylon.js TerrainMaterial.
//
type TerrainMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TerrainMaterial) JSObject() js.Value { return t.p }

// TerrainMaterial returns a TerrainMaterial JavaScript class.
func (ba *Babylon) TerrainMaterial() *TerrainMaterial {
	p := ba.ctx.Get("TerrainMaterial")
	return TerrainMaterialFromJSObject(p, ba.ctx)
}

// TerrainMaterialFromJSObject returns a wrapped TerrainMaterial JavaScript class.
func TerrainMaterialFromJSObject(p js.Value, ctx js.Value) *TerrainMaterial {
	return &TerrainMaterial{p: p, ctx: ctx}
}

// TerrainMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func TerrainMaterialArrayToJSArray(array []*TerrainMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTerrainMaterial returns a new TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#constructor
func (ba *Babylon) NewTerrainMaterial(name string, scene *Scene) *TerrainMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("TerrainMaterial").New(args...)
	return TerrainMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bindforsubmesh
func (t *TerrainMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

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

	t.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#clone
func (t *TerrainMaterial) Clone(name string) *TerrainMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("clone", args...)
	return TerrainMaterialFromJSObject(retVal, t.ctx)
}

// TerrainMaterialDisposeOpts contains optional parameters for TerrainMaterial.Dispose.
type TerrainMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#dispose
func (t *TerrainMaterial) Dispose(opts *TerrainMaterialDisposeOpts) {
	if opts == nil {
		opts = &TerrainMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	t.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#getactivetextures
func (t *TerrainMaterial) GetActiveTextures() []*BaseTexture {

	retVal := t.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), t.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#getalphatesttexture
func (t *TerrainMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := t.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, t.ctx)
}

// GetAnimatables calls the GetAnimatables method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#getanimatables
func (t *TerrainMaterial) GetAnimatables() []*IAnimatable {

	retVal := t.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), t.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#getclassname
func (t *TerrainMaterial) GetClassName() string {

	retVal := t.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#hastexture
func (t *TerrainMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := t.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// TerrainMaterialIsReadyForSubMeshOpts contains optional parameters for TerrainMaterial.IsReadyForSubMesh.
type TerrainMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#isreadyforsubmesh
func (t *TerrainMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *TerrainMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &TerrainMaterialIsReadyForSubMeshOpts{}
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

	retVal := t.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#needalphablending
func (t *TerrainMaterial) NeedAlphaBlending() bool {

	retVal := t.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#needalphatesting
func (t *TerrainMaterial) NeedAlphaTesting() bool {

	retVal := t.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#parse
func (t *TerrainMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *TerrainMaterial {

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

	retVal := t.p.Call("Parse", args...)
	return TerrainMaterialFromJSObject(retVal, t.ctx)
}

// Serialize calls the Serialize method on the TerrainMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#serialize
func (t *TerrainMaterial) Serialize() js.Value {

	retVal := t.p.Call("serialize")
	return retVal
}

// BumpTexture1 returns the BumpTexture1 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture1
func (t *TerrainMaterial) BumpTexture1() *Texture {
	retVal := t.p.Get("bumpTexture1")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetBumpTexture1 sets the BumpTexture1 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture1
func (t *TerrainMaterial) SetBumpTexture1(bumpTexture1 *Texture) *TerrainMaterial {
	t.p.Set("bumpTexture1", bumpTexture1.JSObject())
	return t
}

// BumpTexture2 returns the BumpTexture2 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture2
func (t *TerrainMaterial) BumpTexture2() *Texture {
	retVal := t.p.Get("bumpTexture2")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetBumpTexture2 sets the BumpTexture2 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture2
func (t *TerrainMaterial) SetBumpTexture2(bumpTexture2 *Texture) *TerrainMaterial {
	t.p.Set("bumpTexture2", bumpTexture2.JSObject())
	return t
}

// BumpTexture3 returns the BumpTexture3 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture3
func (t *TerrainMaterial) BumpTexture3() *Texture {
	retVal := t.p.Get("bumpTexture3")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetBumpTexture3 sets the BumpTexture3 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#bumptexture3
func (t *TerrainMaterial) SetBumpTexture3(bumpTexture3 *Texture) *TerrainMaterial {
	t.p.Set("bumpTexture3", bumpTexture3.JSObject())
	return t
}

// DiffuseColor returns the DiffuseColor property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusecolor
func (t *TerrainMaterial) DiffuseColor() *Color3 {
	retVal := t.p.Get("diffuseColor")
	return Color3FromJSObject(retVal, t.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusecolor
func (t *TerrainMaterial) SetDiffuseColor(diffuseColor *Color3) *TerrainMaterial {
	t.p.Set("diffuseColor", diffuseColor.JSObject())
	return t
}

// DiffuseTexture1 returns the DiffuseTexture1 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture1
func (t *TerrainMaterial) DiffuseTexture1() *Texture {
	retVal := t.p.Get("diffuseTexture1")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetDiffuseTexture1 sets the DiffuseTexture1 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture1
func (t *TerrainMaterial) SetDiffuseTexture1(diffuseTexture1 *Texture) *TerrainMaterial {
	t.p.Set("diffuseTexture1", diffuseTexture1.JSObject())
	return t
}

// DiffuseTexture2 returns the DiffuseTexture2 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture2
func (t *TerrainMaterial) DiffuseTexture2() *Texture {
	retVal := t.p.Get("diffuseTexture2")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetDiffuseTexture2 sets the DiffuseTexture2 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture2
func (t *TerrainMaterial) SetDiffuseTexture2(diffuseTexture2 *Texture) *TerrainMaterial {
	t.p.Set("diffuseTexture2", diffuseTexture2.JSObject())
	return t
}

// DiffuseTexture3 returns the DiffuseTexture3 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture3
func (t *TerrainMaterial) DiffuseTexture3() *Texture {
	retVal := t.p.Get("diffuseTexture3")
	return TextureFromJSObject(retVal, t.ctx)
}

// SetDiffuseTexture3 sets the DiffuseTexture3 property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#diffusetexture3
func (t *TerrainMaterial) SetDiffuseTexture3(diffuseTexture3 *Texture) *TerrainMaterial {
	t.p.Set("diffuseTexture3", diffuseTexture3.JSObject())
	return t
}

// DisableLighting returns the DisableLighting property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#disablelighting
func (t *TerrainMaterial) DisableLighting() bool {
	retVal := t.p.Get("disableLighting")
	return retVal.Bool()
}

// SetDisableLighting sets the DisableLighting property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#disablelighting
func (t *TerrainMaterial) SetDisableLighting(disableLighting bool) *TerrainMaterial {
	t.p.Set("disableLighting", disableLighting)
	return t
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#maxsimultaneouslights
func (t *TerrainMaterial) MaxSimultaneousLights() float64 {
	retVal := t.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#maxsimultaneouslights
func (t *TerrainMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *TerrainMaterial {
	t.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return t
}

// MixTexture returns the MixTexture property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#mixtexture
func (t *TerrainMaterial) MixTexture() *BaseTexture {
	retVal := t.p.Get("mixTexture")
	return BaseTextureFromJSObject(retVal, t.ctx)
}

// SetMixTexture sets the MixTexture property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#mixtexture
func (t *TerrainMaterial) SetMixTexture(mixTexture *BaseTexture) *TerrainMaterial {
	t.p.Set("mixTexture", mixTexture.JSObject())
	return t
}

// SpecularColor returns the SpecularColor property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#specularcolor
func (t *TerrainMaterial) SpecularColor() *Color3 {
	retVal := t.p.Get("specularColor")
	return Color3FromJSObject(retVal, t.ctx)
}

// SetSpecularColor sets the SpecularColor property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#specularcolor
func (t *TerrainMaterial) SetSpecularColor(specularColor *Color3) *TerrainMaterial {
	t.p.Set("specularColor", specularColor.JSObject())
	return t
}

// SpecularPower returns the SpecularPower property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#specularpower
func (t *TerrainMaterial) SpecularPower() float64 {
	retVal := t.p.Get("specularPower")
	return retVal.Float()
}

// SetSpecularPower sets the SpecularPower property of class TerrainMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.terrainmaterial#specularpower
func (t *TerrainMaterial) SetSpecularPower(specularPower float64) *TerrainMaterial {
	t.p.Set("specularPower", specularPower)
	return t
}
