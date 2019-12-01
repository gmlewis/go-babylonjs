// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeepAssets represents a babylon.js KeepAssets.
// Set of assets to keep when moving a scene into an asset container.
type KeepAssets struct {
	*AbstractScene
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KeepAssets) JSObject() js.Value { return k.p }

// KeepAssets returns a KeepAssets JavaScript class.
func (ba *Babylon) KeepAssets() *KeepAssets {
	p := ba.ctx.Get("KeepAssets")
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// KeepAssetsFromJSObject returns a wrapped KeepAssets JavaScript class.
func KeepAssetsFromJSObject(p js.Value, ctx js.Value) *KeepAssets {
	return &KeepAssets{AbstractScene: AbstractSceneFromJSObject(p, ctx), ctx: ctx}
}

// KeepAssetsArrayToJSArray returns a JavaScript Array for the wrapped array.
func KeepAssetsArrayToJSArray(array []*KeepAssets) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddEffectLayer calls the AddEffectLayer method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#addeffectlayer
func (k *KeepAssets) AddEffectLayer(newEffectLayer *EffectLayer) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newEffectLayer.JSObject())

	k.p.Call("addEffectLayer", args...)
}

// AddIndividualParser calls the AddIndividualParser method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#addindividualparser
func (k *KeepAssets) AddIndividualParser(name string, parser js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, parser)

	k.p.Call("AddIndividualParser", args...)
}

// AddLensFlareSystem calls the AddLensFlareSystem method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#addlensflaresystem
func (k *KeepAssets) AddLensFlareSystem(newLensFlareSystem *LensFlareSystem) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newLensFlareSystem.JSObject())

	k.p.Call("addLensFlareSystem", args...)
}

// AddParser calls the AddParser method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#addparser
func (k *KeepAssets) AddParser(name string, parser js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, parser)

	k.p.Call("AddParser", args...)
}

// AddReflectionProbe calls the AddReflectionProbe method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#addreflectionprobe
func (k *KeepAssets) AddReflectionProbe(newReflectionProbe *ReflectionProbe) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newReflectionProbe.JSObject())

	k.p.Call("addReflectionProbe", args...)
}

// GetGlowLayerByName calls the GetGlowLayerByName method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#getglowlayerbyname
func (k *KeepAssets) GetGlowLayerByName(name string) *GlowLayer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := k.p.Call("getGlowLayerByName", args...)
	return GlowLayerFromJSObject(retVal, k.ctx)
}

// GetHighlightLayerByName calls the GetHighlightLayerByName method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#gethighlightlayerbyname
func (k *KeepAssets) GetHighlightLayerByName(name string) *HighlightLayer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := k.p.Call("getHighlightLayerByName", args...)
	return HighlightLayerFromJSObject(retVal, k.ctx)
}

// GetIndividualParser calls the GetIndividualParser method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#getindividualparser
func (k *KeepAssets) GetIndividualParser(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := k.p.Call("GetIndividualParser", args...)
	return retVal
}

// GetLensFlareSystemByID calls the GetLensFlareSystemByID method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#getlensflaresystembyid
func (k *KeepAssets) GetLensFlareSystemByID(id string) *LensFlareSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := k.p.Call("getLensFlareSystemByID", args...)
	return LensFlareSystemFromJSObject(retVal, k.ctx)
}

// GetLensFlareSystemByName calls the GetLensFlareSystemByName method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#getlensflaresystembyname
func (k *KeepAssets) GetLensFlareSystemByName(name string) *LensFlareSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := k.p.Call("getLensFlareSystemByName", args...)
	return LensFlareSystemFromJSObject(retVal, k.ctx)
}

// GetParser calls the GetParser method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#getparser
func (k *KeepAssets) GetParser(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := k.p.Call("GetParser", args...)
	return retVal
}

// Parse calls the Parse method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#parse
func (k *KeepAssets) Parse(jsonData interface{}, scene *Scene, container *AssetContainer, rootUrl string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, jsonData)
	args = append(args, scene.JSObject())
	args = append(args, container.JSObject())
	args = append(args, rootUrl)

	k.p.Call("Parse", args...)
}

// RemoveEffectLayer calls the RemoveEffectLayer method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#removeeffectlayer
func (k *KeepAssets) RemoveEffectLayer(toRemove *EffectLayer) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := k.p.Call("removeEffectLayer", args...)
	return retVal.Float()
}

// RemoveLensFlareSystem calls the RemoveLensFlareSystem method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#removelensflaresystem
func (k *KeepAssets) RemoveLensFlareSystem(toRemove *LensFlareSystem) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := k.p.Call("removeLensFlareSystem", args...)
	return retVal.Float()
}

// RemoveReflectionProbe calls the RemoveReflectionProbe method on the KeepAssets object.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#removereflectionprobe
func (k *KeepAssets) RemoveReflectionProbe(toRemove *ReflectionProbe) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := k.p.Call("removeReflectionProbe", args...)
	return retVal.Float()
}

/*

// ActionManagers returns the ActionManagers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#actionmanagers
func (k *KeepAssets) ActionManagers(actionManagers *AbstractActionManager) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(actionManagers.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetActionManagers sets the ActionManagers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#actionmanagers
func (k *KeepAssets) SetActionManagers(actionManagers *AbstractActionManager) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(actionManagers.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// AnimationGroups returns the AnimationGroups property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#animationgroups
func (k *KeepAssets) AnimationGroups(animationGroups *AnimationGroup) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(animationGroups.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetAnimationGroups sets the AnimationGroups property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#animationgroups
func (k *KeepAssets) SetAnimationGroups(animationGroups *AnimationGroup) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(animationGroups.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#animations
func (k *KeepAssets) Animations(animations *Animation) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(animations.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#animations
func (k *KeepAssets) SetAnimations(animations *Animation) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(animations.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Cameras returns the Cameras property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#cameras
func (k *KeepAssets) Cameras(cameras *Camera) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(cameras.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetCameras sets the Cameras property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#cameras
func (k *KeepAssets) SetCameras(cameras *Camera) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(cameras.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// EffectLayers returns the EffectLayers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#effectlayers
func (k *KeepAssets) EffectLayers(effectLayers []*EffectLayer) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(effectLayers)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetEffectLayers sets the EffectLayers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#effectlayers
func (k *KeepAssets) SetEffectLayers(effectLayers []*EffectLayer) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(effectLayers)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// EnvironmentTexture returns the EnvironmentTexture property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#environmenttexture
func (k *KeepAssets) EnvironmentTexture(environmentTexture *BaseTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(environmentTexture.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetEnvironmentTexture sets the EnvironmentTexture property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#environmenttexture
func (k *KeepAssets) SetEnvironmentTexture(environmentTexture *BaseTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(environmentTexture.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Geometries returns the Geometries property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#geometries
func (k *KeepAssets) Geometries(geometries *Geometry) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(geometries.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetGeometries sets the Geometries property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#geometries
func (k *KeepAssets) SetGeometries(geometries *Geometry) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(geometries.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Layers returns the Layers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#layers
func (k *KeepAssets) Layers(layers []*Layer) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(layers)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetLayers sets the Layers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#layers
func (k *KeepAssets) SetLayers(layers []*Layer) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(layers)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// LensFlareSystems returns the LensFlareSystems property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#lensflaresystems
func (k *KeepAssets) LensFlareSystems(lensFlareSystems []*LensFlareSystem) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(lensFlareSystems)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetLensFlareSystems sets the LensFlareSystems property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#lensflaresystems
func (k *KeepAssets) SetLensFlareSystems(lensFlareSystems []*LensFlareSystem) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(lensFlareSystems)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Lights returns the Lights property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#lights
func (k *KeepAssets) Lights(lights *Light) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(lights.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetLights sets the Lights property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#lights
func (k *KeepAssets) SetLights(lights *Light) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(lights.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Materials returns the Materials property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#materials
func (k *KeepAssets) Materials(materials *Material) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(materials.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetMaterials sets the Materials property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#materials
func (k *KeepAssets) SetMaterials(materials *Material) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(materials.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Meshes returns the Meshes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#meshes
func (k *KeepAssets) Meshes(meshes *AbstractMesh) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(meshes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetMeshes sets the Meshes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#meshes
func (k *KeepAssets) SetMeshes(meshes *AbstractMesh) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(meshes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// MorphTargetManagers returns the MorphTargetManagers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#morphtargetmanagers
func (k *KeepAssets) MorphTargetManagers(morphTargetManagers *MorphTargetManager) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(morphTargetManagers.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetMorphTargetManagers sets the MorphTargetManagers property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#morphtargetmanagers
func (k *KeepAssets) SetMorphTargetManagers(morphTargetManagers *MorphTargetManager) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(morphTargetManagers.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// MultiMaterials returns the MultiMaterials property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#multimaterials
func (k *KeepAssets) MultiMaterials(multiMaterials *MultiMaterial) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(multiMaterials.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetMultiMaterials sets the MultiMaterials property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#multimaterials
func (k *KeepAssets) SetMultiMaterials(multiMaterials *MultiMaterial) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(multiMaterials.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// ParticleSystems returns the ParticleSystems property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#particlesystems
func (k *KeepAssets) ParticleSystems(particleSystems js.Value) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(particleSystems)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetParticleSystems sets the ParticleSystems property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#particlesystems
func (k *KeepAssets) SetParticleSystems(particleSystems js.Value) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(particleSystems)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// ProceduralTextures returns the ProceduralTextures property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#proceduraltextures
func (k *KeepAssets) ProceduralTextures(proceduralTextures []*ProceduralTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(proceduralTextures)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetProceduralTextures sets the ProceduralTextures property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#proceduraltextures
func (k *KeepAssets) SetProceduralTextures(proceduralTextures []*ProceduralTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(proceduralTextures)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// ReflectionProbes returns the ReflectionProbes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#reflectionprobes
func (k *KeepAssets) ReflectionProbes(reflectionProbes []*ReflectionProbe) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(reflectionProbes)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetReflectionProbes sets the ReflectionProbes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#reflectionprobes
func (k *KeepAssets) SetReflectionProbes(reflectionProbes []*ReflectionProbe) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(reflectionProbes)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// RootNodes returns the RootNodes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#rootnodes
func (k *KeepAssets) RootNodes(rootNodes *Node) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(rootNodes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetRootNodes sets the RootNodes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#rootnodes
func (k *KeepAssets) SetRootNodes(rootNodes *Node) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(rootNodes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Skeletons returns the Skeletons property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#skeletons
func (k *KeepAssets) Skeletons(skeletons *Skeleton) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(skeletons.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetSkeletons sets the Skeletons property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#skeletons
func (k *KeepAssets) SetSkeletons(skeletons *Skeleton) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(skeletons.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Sounds returns the Sounds property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#sounds
func (k *KeepAssets) Sounds(sounds []*Sound) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(sounds)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetSounds sets the Sounds property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#sounds
func (k *KeepAssets) SetSounds(sounds []*Sound) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(sounds)
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// Textures returns the Textures property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#textures
func (k *KeepAssets) Textures(textures *BaseTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(textures.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetTextures sets the Textures property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#textures
func (k *KeepAssets) SetTextures(textures *BaseTexture) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(textures.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// TransformNodes returns the TransformNodes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#transformnodes
func (k *KeepAssets) TransformNodes(transformNodes *TransformNode) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(transformNodes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// SetTransformNodes sets the TransformNodes property of class KeepAssets.
//
// https://doc.babylonjs.com/api/classes/babylon.keepassets#transformnodes
func (k *KeepAssets) SetTransformNodes(transformNodes *TransformNode) *KeepAssets {
	p := ba.ctx.Get("KeepAssets").New(transformNodes.JSObject())
	return KeepAssetsFromJSObject(p, ba.ctx)
}

*/
