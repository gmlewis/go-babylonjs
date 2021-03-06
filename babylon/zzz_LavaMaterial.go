// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LavaMaterial represents a babylon.js LavaMaterial.
//
type LavaMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LavaMaterial) JSObject() js.Value { return l.p }

// LavaMaterial returns a LavaMaterial JavaScript class.
func (ba *Babylon) LavaMaterial() *LavaMaterial {
	p := ba.ctx.Get("LavaMaterial")
	return LavaMaterialFromJSObject(p, ba.ctx)
}

// LavaMaterialFromJSObject returns a wrapped LavaMaterial JavaScript class.
func LavaMaterialFromJSObject(p js.Value, ctx js.Value) *LavaMaterial {
	return &LavaMaterial{p: p, ctx: ctx}
}

// LavaMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func LavaMaterialArrayToJSArray(array []*LavaMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLavaMaterial returns a new LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#constructor
func (ba *Babylon) NewLavaMaterial(name string, scene *Scene) *LavaMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("LavaMaterial").New(args...)
	return LavaMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#bindforsubmesh
func (l *LavaMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

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

	l.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#clone
func (l *LavaMaterial) Clone(name string) *LavaMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := l.p.Call("clone", args...)
	return LavaMaterialFromJSObject(retVal, l.ctx)
}

// LavaMaterialDisposeOpts contains optional parameters for LavaMaterial.Dispose.
type LavaMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#dispose
func (l *LavaMaterial) Dispose(opts *LavaMaterialDisposeOpts) {
	if opts == nil {
		opts = &LavaMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	l.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#getactivetextures
func (l *LavaMaterial) GetActiveTextures() []*BaseTexture {

	retVal := l.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), l.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#getalphatesttexture
func (l *LavaMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := l.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, l.ctx)
}

// GetAnimatables calls the GetAnimatables method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#getanimatables
func (l *LavaMaterial) GetAnimatables() []*IAnimatable {

	retVal := l.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), l.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#getclassname
func (l *LavaMaterial) GetClassName() string {

	retVal := l.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#hastexture
func (l *LavaMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := l.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// LavaMaterialIsReadyForSubMeshOpts contains optional parameters for LavaMaterial.IsReadyForSubMesh.
type LavaMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#isreadyforsubmesh
func (l *LavaMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *LavaMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &LavaMaterialIsReadyForSubMeshOpts{}
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

	retVal := l.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#needalphablending
func (l *LavaMaterial) NeedAlphaBlending() bool {

	retVal := l.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#needalphatesting
func (l *LavaMaterial) NeedAlphaTesting() bool {

	retVal := l.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#parse
func (l *LavaMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *LavaMaterial {

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

	retVal := l.p.Call("Parse", args...)
	return LavaMaterialFromJSObject(retVal, l.ctx)
}

// Serialize calls the Serialize method on the LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#serialize
func (l *LavaMaterial) Serialize() js.Value {

	retVal := l.p.Call("serialize")
	return retVal
}

// DiffuseColor returns the DiffuseColor property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#diffusecolor
func (l *LavaMaterial) DiffuseColor() *Color3 {
	retVal := l.p.Get("diffuseColor")
	return Color3FromJSObject(retVal, l.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#diffusecolor
func (l *LavaMaterial) SetDiffuseColor(diffuseColor *Color3) *LavaMaterial {
	l.p.Set("diffuseColor", diffuseColor.JSObject())
	return l
}

// DiffuseTexture returns the DiffuseTexture property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#diffusetexture
func (l *LavaMaterial) DiffuseTexture() *BaseTexture {
	retVal := l.p.Get("diffuseTexture")
	return BaseTextureFromJSObject(retVal, l.ctx)
}

// SetDiffuseTexture sets the DiffuseTexture property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#diffusetexture
func (l *LavaMaterial) SetDiffuseTexture(diffuseTexture *BaseTexture) *LavaMaterial {
	l.p.Set("diffuseTexture", diffuseTexture.JSObject())
	return l
}

// DisableLighting returns the DisableLighting property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#disablelighting
func (l *LavaMaterial) DisableLighting() bool {
	retVal := l.p.Get("disableLighting")
	return retVal.Bool()
}

// SetDisableLighting sets the DisableLighting property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#disablelighting
func (l *LavaMaterial) SetDisableLighting(disableLighting bool) *LavaMaterial {
	l.p.Set("disableLighting", disableLighting)
	return l
}

// FogColor returns the FogColor property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#fogcolor
func (l *LavaMaterial) FogColor() *Color3 {
	retVal := l.p.Get("fogColor")
	return Color3FromJSObject(retVal, l.ctx)
}

// SetFogColor sets the FogColor property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#fogcolor
func (l *LavaMaterial) SetFogColor(fogColor *Color3) *LavaMaterial {
	l.p.Set("fogColor", fogColor.JSObject())
	return l
}

// FogDensity returns the FogDensity property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#fogdensity
func (l *LavaMaterial) FogDensity() float64 {
	retVal := l.p.Get("fogDensity")
	return retVal.Float()
}

// SetFogDensity sets the FogDensity property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#fogdensity
func (l *LavaMaterial) SetFogDensity(fogDensity float64) *LavaMaterial {
	l.p.Set("fogDensity", fogDensity)
	return l
}

// LowFrequencySpeed returns the LowFrequencySpeed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#lowfrequencyspeed
func (l *LavaMaterial) LowFrequencySpeed() float64 {
	retVal := l.p.Get("lowFrequencySpeed")
	return retVal.Float()
}

// SetLowFrequencySpeed sets the LowFrequencySpeed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#lowfrequencyspeed
func (l *LavaMaterial) SetLowFrequencySpeed(lowFrequencySpeed float64) *LavaMaterial {
	l.p.Set("lowFrequencySpeed", lowFrequencySpeed)
	return l
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#maxsimultaneouslights
func (l *LavaMaterial) MaxSimultaneousLights() float64 {
	retVal := l.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#maxsimultaneouslights
func (l *LavaMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *LavaMaterial {
	l.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return l
}

// MovingSpeed returns the MovingSpeed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#movingspeed
func (l *LavaMaterial) MovingSpeed() float64 {
	retVal := l.p.Get("movingSpeed")
	return retVal.Float()
}

// SetMovingSpeed sets the MovingSpeed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#movingspeed
func (l *LavaMaterial) SetMovingSpeed(movingSpeed float64) *LavaMaterial {
	l.p.Set("movingSpeed", movingSpeed)
	return l
}

// NoiseTexture returns the NoiseTexture property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#noisetexture
func (l *LavaMaterial) NoiseTexture() *BaseTexture {
	retVal := l.p.Get("noiseTexture")
	return BaseTextureFromJSObject(retVal, l.ctx)
}

// SetNoiseTexture sets the NoiseTexture property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#noisetexture
func (l *LavaMaterial) SetNoiseTexture(noiseTexture *BaseTexture) *LavaMaterial {
	l.p.Set("noiseTexture", noiseTexture.JSObject())
	return l
}

// Speed returns the Speed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#speed
func (l *LavaMaterial) Speed() float64 {
	retVal := l.p.Get("speed")
	return retVal.Float()
}

// SetSpeed sets the Speed property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#speed
func (l *LavaMaterial) SetSpeed(speed float64) *LavaMaterial {
	l.p.Set("speed", speed)
	return l
}

// Unlit returns the Unlit property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#unlit
func (l *LavaMaterial) Unlit() bool {
	retVal := l.p.Get("unlit")
	return retVal.Bool()
}

// SetUnlit sets the Unlit property of class LavaMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial#unlit
func (l *LavaMaterial) SetUnlit(unlit bool) *LavaMaterial {
	l.p.Set("unlit", unlit)
	return l
}
