// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HDRCubeTextureAssetTask represents a babylon.js HDRCubeTextureAssetTask.
// Define a task used by AssetsManager to load HDR cube textures
type HDRCubeTextureAssetTask struct {
	*AbstractAssetTask
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HDRCubeTextureAssetTask) JSObject() js.Value { return h.p }

// HDRCubeTextureAssetTask returns a HDRCubeTextureAssetTask JavaScript class.
func (ba *Babylon) HDRCubeTextureAssetTask() *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask")
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// HDRCubeTextureAssetTaskFromJSObject returns a wrapped HDRCubeTextureAssetTask JavaScript class.
func HDRCubeTextureAssetTaskFromJSObject(p js.Value, ctx js.Value) *HDRCubeTextureAssetTask {
	return &HDRCubeTextureAssetTask{AbstractAssetTask: AbstractAssetTaskFromJSObject(p, ctx), ctx: ctx}
}

// HDRCubeTextureAssetTaskArrayToJSArray returns a JavaScript Array for the wrapped array.
func HDRCubeTextureAssetTaskArrayToJSArray(array []*HDRCubeTextureAssetTask) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHDRCubeTextureAssetTaskOpts contains optional parameters for NewHDRCubeTextureAssetTask.
type NewHDRCubeTextureAssetTaskOpts struct {
	NoMipmap          *bool
	GenerateHarmonics *bool
	GammaSpace        *bool
	Reserved          *bool
}

// NewHDRCubeTextureAssetTask returns a new HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask
func (ba *Babylon) NewHDRCubeTextureAssetTask(name string, url string, size float64, opts *NewHDRCubeTextureAssetTaskOpts) *HDRCubeTextureAssetTask {
	if opts == nil {
		opts = &NewHDRCubeTextureAssetTaskOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, name)
	args = append(args, url)
	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GenerateHarmonics == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateHarmonics)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}
	if opts.Reserved == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reserved)
	}

	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(args...)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Reset calls the Reset method on the HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#reset
func (h *HDRCubeTextureAssetTask) Reset() {

	h.p.Call("reset")
}

// Run calls the Run method on the HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#run
func (h *HDRCubeTextureAssetTask) Run(scene *Scene, onSuccess func(), onError func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))

	h.p.Call("run", args...)
}

// RunTask calls the RunTask method on the HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#runtask
func (h *HDRCubeTextureAssetTask) RunTask(scene *Scene, onSuccess func(), onError func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))

	h.p.Call("runTask", args...)
}

/*

// ErrorObject returns the ErrorObject property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#errorobject
func (h *HDRCubeTextureAssetTask) ErrorObject(errorObject js.Value) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(errorObject)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetErrorObject sets the ErrorObject property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#errorobject
func (h *HDRCubeTextureAssetTask) SetErrorObject(errorObject js.Value) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(errorObject)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// GammaSpace returns the GammaSpace property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#gammaspace
func (h *HDRCubeTextureAssetTask) GammaSpace(gammaSpace bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(gammaSpace)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetGammaSpace sets the GammaSpace property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#gammaspace
func (h *HDRCubeTextureAssetTask) SetGammaSpace(gammaSpace bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(gammaSpace)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// GenerateHarmonics returns the GenerateHarmonics property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#generateharmonics
func (h *HDRCubeTextureAssetTask) GenerateHarmonics(generateHarmonics bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(generateHarmonics)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetGenerateHarmonics sets the GenerateHarmonics property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#generateharmonics
func (h *HDRCubeTextureAssetTask) SetGenerateHarmonics(generateHarmonics bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(generateHarmonics)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// IsCompleted returns the IsCompleted property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#iscompleted
func (h *HDRCubeTextureAssetTask) IsCompleted(isCompleted bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(isCompleted)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetIsCompleted sets the IsCompleted property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#iscompleted
func (h *HDRCubeTextureAssetTask) SetIsCompleted(isCompleted bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(isCompleted)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#name
func (h *HDRCubeTextureAssetTask) Name(name string) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(name)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#name
func (h *HDRCubeTextureAssetTask) SetName(name string) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(name)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// NoMipmap returns the NoMipmap property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#nomipmap
func (h *HDRCubeTextureAssetTask) NoMipmap(noMipmap bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(noMipmap)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetNoMipmap sets the NoMipmap property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#nomipmap
func (h *HDRCubeTextureAssetTask) SetNoMipmap(noMipmap bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(noMipmap)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// OnError returns the OnError property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#onerror
func (h *HDRCubeTextureAssetTask) OnError(onError func()) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onError(); return nil}))
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetOnError sets the OnError property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#onerror
func (h *HDRCubeTextureAssetTask) SetOnError(onError func()) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onError(); return nil}))
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// OnSuccess returns the OnSuccess property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#onsuccess
func (h *HDRCubeTextureAssetTask) OnSuccess(onSuccess func()) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSuccess(); return nil}))
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetOnSuccess sets the OnSuccess property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#onsuccess
func (h *HDRCubeTextureAssetTask) SetOnSuccess(onSuccess func()) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSuccess(); return nil}))
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Reserved returns the Reserved property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#reserved
func (h *HDRCubeTextureAssetTask) Reserved(reserved bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(reserved)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetReserved sets the Reserved property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#reserved
func (h *HDRCubeTextureAssetTask) SetReserved(reserved bool) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(reserved)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Size returns the Size property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#size
func (h *HDRCubeTextureAssetTask) Size(size float64) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(size)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetSize sets the Size property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#size
func (h *HDRCubeTextureAssetTask) SetSize(size float64) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(size)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// TaskState returns the TaskState property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#taskstate
func (h *HDRCubeTextureAssetTask) TaskState(taskState *AssetTaskState) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(taskState.JSObject())
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetTaskState sets the TaskState property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#taskstate
func (h *HDRCubeTextureAssetTask) SetTaskState(taskState *AssetTaskState) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(taskState.JSObject())
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Texture returns the Texture property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#texture
func (h *HDRCubeTextureAssetTask) Texture(texture *HDRCubeTexture) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(texture.JSObject())
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetTexture sets the Texture property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#texture
func (h *HDRCubeTextureAssetTask) SetTexture(texture *HDRCubeTexture) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(texture.JSObject())
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// Url returns the Url property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#url
func (h *HDRCubeTextureAssetTask) Url(url string) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(url)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// SetUrl sets the Url property of class HDRCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask#url
func (h *HDRCubeTextureAssetTask) SetUrl(url string) *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(url)
	return HDRCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

*/
