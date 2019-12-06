// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// STLFileLoader represents a babylon.js STLFileLoader.
// STL file type loader.
// This is a babylon scene loader plugin.
type STLFileLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *STLFileLoader) JSObject() js.Value { return s.p }

// STLFileLoader returns a STLFileLoader JavaScript class.
func (ba *Babylon) STLFileLoader() *STLFileLoader {
	p := ba.ctx.Get("STLFileLoader")
	return STLFileLoaderFromJSObject(p, ba.ctx)
}

// STLFileLoaderFromJSObject returns a wrapped STLFileLoader JavaScript class.
func STLFileLoaderFromJSObject(p js.Value, ctx js.Value) *STLFileLoader {
	return &STLFileLoader{p: p, ctx: ctx}
}

// STLFileLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func STLFileLoaderArrayToJSArray(array []*STLFileLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ImportMesh calls the ImportMesh method on the STLFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#importmesh
func (s *STLFileLoader) ImportMesh(meshesNames interface{}, scene *Scene, data interface{}, rootUrl string, meshes []*AbstractMesh, particleSystems []*IParticleSystem, skeletons []*Skeleton) bool {

	args := make([]interface{}, 0, 7+0)

	args = append(args, meshesNames)
	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)
	args = append(args, AbstractMeshArrayToJSArray(meshes))
	args = append(args, IParticleSystemArrayToJSArray(particleSystems))
	args = append(args, SkeletonArrayToJSArray(skeletons))

	retVal := s.p.Call("importMesh", args...)
	return retVal.Bool()
}

// Load calls the Load method on the STLFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#load
func (s *STLFileLoader) Load(scene *Scene, data interface{}, rootUrl string) bool {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	retVal := s.p.Call("load", args...)
	return retVal.Bool()
}

// STLFileLoaderLoadAssetContainerOpts contains optional parameters for STLFileLoader.LoadAssetContainer.
type STLFileLoaderLoadAssetContainerOpts struct {
	OnError func()
}

// LoadAssetContainer calls the LoadAssetContainer method on the STLFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#loadassetcontainer
func (s *STLFileLoader) LoadAssetContainer(scene *Scene, data string, rootUrl string, opts *STLFileLoaderLoadAssetContainerOpts) *AssetContainer {
	if opts == nil {
		opts = &STLFileLoaderLoadAssetContainerOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnError(); return nil }) /* never freed! */)
	}

	retVal := s.p.Call("loadAssetContainer", args...)
	return AssetContainerFromJSObject(retVal, s.ctx)
}

// Extensions returns the Extensions property of class STLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#extensions
func (s *STLFileLoader) Extensions() *ISceneLoaderPluginExtensions {
	retVal := s.p.Get("extensions")
	return ISceneLoaderPluginExtensionsFromJSObject(retVal, s.ctx)
}

// SetExtensions sets the Extensions property of class STLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#extensions
func (s *STLFileLoader) SetExtensions(extensions *ISceneLoaderPluginExtensions) *STLFileLoader {
	s.p.Set("extensions", extensions.JSObject())
	return s
}

// Name returns the Name property of class STLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#name
func (s *STLFileLoader) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class STLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.stlfileloader#name
func (s *STLFileLoader) SetName(name string) *STLFileLoader {
	s.p.Set("name", name)
	return s
}
