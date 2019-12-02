// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SimpleMaterial represents a babylon.js SimpleMaterial.
//
type SimpleMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SimpleMaterial) JSObject() js.Value { return s.p }

// SimpleMaterial returns a SimpleMaterial JavaScript class.
func (ba *Babylon) SimpleMaterial() *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial")
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// SimpleMaterialFromJSObject returns a wrapped SimpleMaterial JavaScript class.
func SimpleMaterialFromJSObject(p js.Value, ctx js.Value) *SimpleMaterial {
	return &SimpleMaterial{p: p, ctx: ctx}
}

// SimpleMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func SimpleMaterialArrayToJSArray(array []*SimpleMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSimpleMaterial returns a new SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial
func (ba *Babylon) NewSimpleMaterial(name string, scene *Scene) *SimpleMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("SimpleMaterial").New(args...)
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#bindforsubmesh
func (s *SimpleMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, world.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, subMesh.JSObject())

	s.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#clone
func (s *SimpleMaterial) Clone(name string) *SimpleMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("clone", args...)
	return SimpleMaterialFromJSObject(retVal, s.ctx)
}

// SimpleMaterialDisposeOpts contains optional parameters for SimpleMaterial.Dispose.
type SimpleMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#dispose
func (s *SimpleMaterial) Dispose(opts *SimpleMaterialDisposeOpts) {
	if opts == nil {
		opts = &SimpleMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	s.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#getactivetextures
func (s *SimpleMaterial) GetActiveTextures() *BaseTexture {

	retVal := s.p.Call("getActiveTextures")
	return BaseTextureFromJSObject(retVal, s.ctx)
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#getalphatesttexture
func (s *SimpleMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := s.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, s.ctx)
}

// GetAnimatables calls the GetAnimatables method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#getanimatables
func (s *SimpleMaterial) GetAnimatables() *IAnimatable {

	retVal := s.p.Call("getAnimatables")
	return IAnimatableFromJSObject(retVal, s.ctx)
}

// GetClassName calls the GetClassName method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#getclassname
func (s *SimpleMaterial) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#hastexture
func (s *SimpleMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, texture.JSObject())

	retVal := s.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// SimpleMaterialIsReadyForSubMeshOpts contains optional parameters for SimpleMaterial.IsReadyForSubMesh.
type SimpleMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#isreadyforsubmesh
func (s *SimpleMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *SimpleMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &SimpleMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, mesh.JSObject())
	args = append(args, subMesh.JSObject())

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := s.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#needalphablending
func (s *SimpleMaterial) NeedAlphaBlending() bool {

	retVal := s.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#needalphatesting
func (s *SimpleMaterial) NeedAlphaTesting() bool {

	retVal := s.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#parse
func (s *SimpleMaterial) Parse(source interface{}, scene *Scene, rootUrl string) *SimpleMaterial {

	args := make([]interface{}, 0, 3+0)

	args = append(args, source)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := s.p.Call("Parse", args...)
	return SimpleMaterialFromJSObject(retVal, s.ctx)
}

// Serialize calls the Serialize method on the SimpleMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#serialize
func (s *SimpleMaterial) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

/*

// DiffuseColor returns the DiffuseColor property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#diffusecolor
func (s *SimpleMaterial) DiffuseColor(diffuseColor *Color3) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(diffuseColor.JSObject())
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#diffusecolor
func (s *SimpleMaterial) SetDiffuseColor(diffuseColor *Color3) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(diffuseColor.JSObject())
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// DiffuseTexture returns the DiffuseTexture property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#diffusetexture
func (s *SimpleMaterial) DiffuseTexture(diffuseTexture *BaseTexture) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(diffuseTexture.JSObject())
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// SetDiffuseTexture sets the DiffuseTexture property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#diffusetexture
func (s *SimpleMaterial) SetDiffuseTexture(diffuseTexture *BaseTexture) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(diffuseTexture.JSObject())
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// DisableLighting returns the DisableLighting property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#disablelighting
func (s *SimpleMaterial) DisableLighting(disableLighting bool) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(disableLighting)
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// SetDisableLighting sets the DisableLighting property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#disablelighting
func (s *SimpleMaterial) SetDisableLighting(disableLighting bool) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(disableLighting)
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#maxsimultaneouslights
func (s *SimpleMaterial) MaxSimultaneousLights(maxSimultaneousLights float64) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(maxSimultaneousLights)
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class SimpleMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.simplematerial#maxsimultaneouslights
func (s *SimpleMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *SimpleMaterial {
	p := ba.ctx.Get("SimpleMaterial").New(maxSimultaneousLights)
	return SimpleMaterialFromJSObject(p, ba.ctx)
}

*/
