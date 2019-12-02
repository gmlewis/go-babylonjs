// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISceneLoaderPluginAsync represents a babylon.js ISceneLoaderPluginAsync.
// Interface used to define an async SceneLoader plugin
type ISceneLoaderPluginAsync struct {
	*ISceneLoaderPluginBase
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISceneLoaderPluginAsync) JSObject() js.Value { return i.p }

// ISceneLoaderPluginAsync returns a ISceneLoaderPluginAsync JavaScript class.
func (ba *Babylon) ISceneLoaderPluginAsync() *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync")
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// ISceneLoaderPluginAsyncFromJSObject returns a wrapped ISceneLoaderPluginAsync JavaScript class.
func ISceneLoaderPluginAsyncFromJSObject(p js.Value, ctx js.Value) *ISceneLoaderPluginAsync {
	return &ISceneLoaderPluginAsync{ISceneLoaderPluginBase: ISceneLoaderPluginBaseFromJSObject(p, ctx), ctx: ctx}
}

// ISceneLoaderPluginAsyncArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneLoaderPluginAsyncArrayToJSArray(array []*ISceneLoaderPluginAsync) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ISceneLoaderPluginAsyncImportMeshAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.ImportMeshAsync.
type ISceneLoaderPluginAsyncImportMeshAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// ImportMeshAsync calls the ImportMeshAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#importmeshasync
func (i *ISceneLoaderPluginAsync) ImportMeshAsync(meshesNames interface{}, scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncImportMeshAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncImportMeshAsyncOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, meshesNames)
	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("importMeshAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.LoadAssetContainerAsync.
type ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// LoadAssetContainerAsync calls the LoadAssetContainerAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#loadassetcontainerasync
func (i *ISceneLoaderPluginAsync) LoadAssetContainerAsync(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("loadAssetContainerAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginAsyncLoadAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.LoadAsync.
type ISceneLoaderPluginAsyncLoadAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// LoadAsync calls the LoadAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#loadasync
func (i *ISceneLoaderPluginAsync) LoadAsync(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncLoadAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncLoadAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("loadAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

/*

 */
