// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneLoader represents a babylon.js SceneLoader.
// Class used to load scene from various file formats using registered plugins
//
// See: http://doc.babylonjs.com/how_to/load_from_any_file_type
type SceneLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SceneLoader) JSObject() js.Value { return s.p }

// SceneLoader returns a SceneLoader JavaScript class.
func (ba *Babylon) SceneLoader() *SceneLoader {
	p := ba.ctx.Get("SceneLoader")
	return SceneLoaderFromJSObject(p, ba.ctx)
}

// SceneLoaderFromJSObject returns a wrapped SceneLoader JavaScript class.
func SceneLoaderFromJSObject(p js.Value, ctx js.Value) *SceneLoader {
	return &SceneLoader{p: p, ctx: ctx}
}

// SceneLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func SceneLoaderArrayToJSArray(array []*SceneLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// SceneLoaderAppendOpts contains optional parameters for SceneLoader.Append.
type SceneLoaderAppendOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnSuccess       func()
	OnProgress      func()
	OnError         func()
	PluginExtension *string
}

// Append calls the Append method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#append
func (s *SceneLoader) Append(rootUrl string, opts *SceneLoaderAppendOpts) *ISceneLoaderPlugin {
	if opts == nil {
		opts = &SceneLoaderAppendOpts{}
	}

	args := make([]interface{}, 0, 1+6)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnSuccess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnSuccess(); return nil }) /* never freed! */)
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnError(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("Append", args...)
	return ISceneLoaderPluginFromJSObject(retVal, s.ctx)
}

// SceneLoaderAppendAsyncOpts contains optional parameters for SceneLoader.AppendAsync.
type SceneLoaderAppendAsyncOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnProgress      func()
	PluginExtension *string
}

// AppendAsync calls the AppendAsync method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#appendasync
func (s *SceneLoader) AppendAsync(rootUrl string, opts *SceneLoaderAppendAsyncOpts) *Promise {
	if opts == nil {
		opts = &SceneLoaderAppendAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("AppendAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// GetPluginForExtension calls the GetPluginForExtension method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#getpluginforextension
func (s *SceneLoader) GetPluginForExtension(extension string) *ISceneLoaderPlugin {

	args := make([]interface{}, 0, 1+0)

	args = append(args, extension)

	retVal := s.p.Call("GetPluginForExtension", args...)
	return ISceneLoaderPluginFromJSObject(retVal, s.ctx)
}

// SceneLoaderImportMeshOpts contains optional parameters for SceneLoader.ImportMesh.
type SceneLoaderImportMeshOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnSuccess       func()
	OnProgress      func()
	OnError         func()
	PluginExtension *string
}

// ImportMesh calls the ImportMesh method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#importmesh
func (s *SceneLoader) ImportMesh(meshNames interface{}, rootUrl string, opts *SceneLoaderImportMeshOpts) *ISceneLoaderPlugin {
	if opts == nil {
		opts = &SceneLoaderImportMeshOpts{}
	}

	args := make([]interface{}, 0, 2+6)

	args = append(args, meshNames)
	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnSuccess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnSuccess(); return nil }) /* never freed! */)
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnError(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("ImportMesh", args...)
	return ISceneLoaderPluginFromJSObject(retVal, s.ctx)
}

// SceneLoaderImportMeshAsyncOpts contains optional parameters for SceneLoader.ImportMeshAsync.
type SceneLoaderImportMeshAsyncOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnProgress      func()
	PluginExtension *string
}

// ImportMeshAsync calls the ImportMeshAsync method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#importmeshasync
func (s *SceneLoader) ImportMeshAsync(meshNames interface{}, rootUrl string, opts *SceneLoaderImportMeshAsyncOpts) *Promise {
	if opts == nil {
		opts = &SceneLoaderImportMeshAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+4)

	args = append(args, meshNames)
	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("ImportMeshAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// IsPluginForExtensionAvailable calls the IsPluginForExtensionAvailable method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#ispluginforextensionavailable
func (s *SceneLoader) IsPluginForExtensionAvailable(extension string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, extension)

	retVal := s.p.Call("IsPluginForExtensionAvailable", args...)
	return retVal.Bool()
}

// SceneLoaderLoadOpts contains optional parameters for SceneLoader.Load.
type SceneLoaderLoadOpts struct {
	SceneFilename   *string
	Engine          *Engine
	OnSuccess       func()
	OnProgress      func()
	OnError         func()
	PluginExtension *string
}

// Load calls the Load method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#load
func (s *SceneLoader) Load(rootUrl string, opts *SceneLoaderLoadOpts) *ISceneLoaderPlugin {
	if opts == nil {
		opts = &SceneLoaderLoadOpts{}
	}

	args := make([]interface{}, 0, 1+6)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.OnSuccess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnSuccess(); return nil }) /* never freed! */)
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnError(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("Load", args...)
	return ISceneLoaderPluginFromJSObject(retVal, s.ctx)
}

// SceneLoaderLoadAssetContainerOpts contains optional parameters for SceneLoader.LoadAssetContainer.
type SceneLoaderLoadAssetContainerOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnSuccess       func()
	OnProgress      func()
	OnError         func()
	PluginExtension *string
}

// LoadAssetContainer calls the LoadAssetContainer method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#loadassetcontainer
func (s *SceneLoader) LoadAssetContainer(rootUrl string, opts *SceneLoaderLoadAssetContainerOpts) *ISceneLoaderPlugin {
	if opts == nil {
		opts = &SceneLoaderLoadAssetContainerOpts{}
	}

	args := make([]interface{}, 0, 1+6)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnSuccess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnSuccess(); return nil }) /* never freed! */)
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnError(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("LoadAssetContainer", args...)
	return ISceneLoaderPluginFromJSObject(retVal, s.ctx)
}

// SceneLoaderLoadAssetContainerAsyncOpts contains optional parameters for SceneLoader.LoadAssetContainerAsync.
type SceneLoaderLoadAssetContainerAsyncOpts struct {
	SceneFilename   *string
	Scene           *Scene
	OnProgress      func()
	PluginExtension *string
}

// LoadAssetContainerAsync calls the LoadAssetContainerAsync method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#loadassetcontainerasync
func (s *SceneLoader) LoadAssetContainerAsync(rootUrl string, opts *SceneLoaderLoadAssetContainerAsyncOpts) *Promise {
	if opts == nil {
		opts = &SceneLoaderLoadAssetContainerAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("LoadAssetContainerAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// SceneLoaderLoadAsyncOpts contains optional parameters for SceneLoader.LoadAsync.
type SceneLoaderLoadAsyncOpts struct {
	SceneFilename   *string
	Engine          *Engine
	OnProgress      func()
	PluginExtension *string
}

// LoadAsync calls the LoadAsync method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#loadasync
func (s *SceneLoader) LoadAsync(rootUrl string, opts *SceneLoaderLoadAsyncOpts) *Promise {
	if opts == nil {
		opts = &SceneLoaderLoadAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, rootUrl)

	if opts.SceneFilename == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SceneFilename)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnProgress(); return nil }) /* never freed! */)
	}
	if opts.PluginExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PluginExtension)
	}

	retVal := s.p.Call("LoadAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// RegisterPlugin calls the RegisterPlugin method on the SceneLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#registerplugin
func (s *SceneLoader) RegisterPlugin(plugin *ISceneLoaderPlugin) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, plugin.JSObject())

	s.p.Call("RegisterPlugin", args...)
}

// CleanBoneMatrixWeights returns the CleanBoneMatrixWeights property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#cleanbonematrixweights
func (s *SceneLoader) CleanBoneMatrixWeights() bool {
	retVal := s.p.Get("CleanBoneMatrixWeights")
	return retVal.Bool()
}

// SetCleanBoneMatrixWeights sets the CleanBoneMatrixWeights property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#cleanbonematrixweights
func (s *SceneLoader) SetCleanBoneMatrixWeights(CleanBoneMatrixWeights bool) *SceneLoader {
	s.p.Set("CleanBoneMatrixWeights", CleanBoneMatrixWeights)
	return s
}

// DETAILED_LOGGING returns the DETAILED_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#detailed_logging
func (s *SceneLoader) DETAILED_LOGGING() float64 {
	retVal := s.p.Get("DETAILED_LOGGING")
	return retVal.Float()
}

// SetDETAILED_LOGGING sets the DETAILED_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#detailed_logging
func (s *SceneLoader) SetDETAILED_LOGGING(DETAILED_LOGGING float64) *SceneLoader {
	s.p.Set("DETAILED_LOGGING", DETAILED_LOGGING)
	return s
}

// ForceFullSceneLoadingForIncremental returns the ForceFullSceneLoadingForIncremental property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#forcefullsceneloadingforincremental
func (s *SceneLoader) ForceFullSceneLoadingForIncremental() bool {
	retVal := s.p.Get("ForceFullSceneLoadingForIncremental")
	return retVal.Bool()
}

// SetForceFullSceneLoadingForIncremental sets the ForceFullSceneLoadingForIncremental property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#forcefullsceneloadingforincremental
func (s *SceneLoader) SetForceFullSceneLoadingForIncremental(ForceFullSceneLoadingForIncremental bool) *SceneLoader {
	s.p.Set("ForceFullSceneLoadingForIncremental", ForceFullSceneLoadingForIncremental)
	return s
}

// LoggingLevel returns the LoggingLevel property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#logginglevel
func (s *SceneLoader) LoggingLevel() float64 {
	retVal := s.p.Get("loggingLevel")
	return retVal.Float()
}

// SetLoggingLevel sets the LoggingLevel property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#logginglevel
func (s *SceneLoader) SetLoggingLevel(loggingLevel float64) *SceneLoader {
	s.p.Set("loggingLevel", loggingLevel)
	return s
}

// MINIMAL_LOGGING returns the MINIMAL_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#minimal_logging
func (s *SceneLoader) MINIMAL_LOGGING() float64 {
	retVal := s.p.Get("MINIMAL_LOGGING")
	return retVal.Float()
}

// SetMINIMAL_LOGGING sets the MINIMAL_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#minimal_logging
func (s *SceneLoader) SetMINIMAL_LOGGING(MINIMAL_LOGGING float64) *SceneLoader {
	s.p.Set("MINIMAL_LOGGING", MINIMAL_LOGGING)
	return s
}

// NO_LOGGING returns the NO_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#no_logging
func (s *SceneLoader) NO_LOGGING() float64 {
	retVal := s.p.Get("NO_LOGGING")
	return retVal.Float()
}

// SetNO_LOGGING sets the NO_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#no_logging
func (s *SceneLoader) SetNO_LOGGING(NO_LOGGING float64) *SceneLoader {
	s.p.Set("NO_LOGGING", NO_LOGGING)
	return s
}

// OnPluginActivatedObservable returns the OnPluginActivatedObservable property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#onpluginactivatedobservable
func (s *SceneLoader) OnPluginActivatedObservable() *Observable {
	retVal := s.p.Get("OnPluginActivatedObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnPluginActivatedObservable sets the OnPluginActivatedObservable property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#onpluginactivatedobservable
func (s *SceneLoader) SetOnPluginActivatedObservable(OnPluginActivatedObservable *Observable) *SceneLoader {
	s.p.Set("OnPluginActivatedObservable", OnPluginActivatedObservable.JSObject())
	return s
}

// SUMMARY_LOGGING returns the SUMMARY_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#summary_logging
func (s *SceneLoader) SUMMARY_LOGGING() float64 {
	retVal := s.p.Get("SUMMARY_LOGGING")
	return retVal.Float()
}

// SetSUMMARY_LOGGING sets the SUMMARY_LOGGING property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#summary_logging
func (s *SceneLoader) SetSUMMARY_LOGGING(SUMMARY_LOGGING float64) *SceneLoader {
	s.p.Set("SUMMARY_LOGGING", SUMMARY_LOGGING)
	return s
}

// ShowLoadingScreen returns the ShowLoadingScreen property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#showloadingscreen
func (s *SceneLoader) ShowLoadingScreen() bool {
	retVal := s.p.Get("ShowLoadingScreen")
	return retVal.Bool()
}

// SetShowLoadingScreen sets the ShowLoadingScreen property of class SceneLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloader#showloadingscreen
func (s *SceneLoader) SetShowLoadingScreen(ShowLoadingScreen bool) *SceneLoader {
	s.p.Set("ShowLoadingScreen", ShowLoadingScreen)
	return s
}
