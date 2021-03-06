// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VideoTexture represents a babylon.js VideoTexture.
// If you want to display a video in your scene, this is the special texture for that.
// This special texture works similar to other textures, with the exception of a few parameters.
//
// See: https://doc.babylonjs.com/how_to/video_texture
type VideoTexture struct {
	*Texture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VideoTexture) JSObject() js.Value { return v.p }

// VideoTexture returns a VideoTexture JavaScript class.
func (ba *Babylon) VideoTexture() *VideoTexture {
	p := ba.ctx.Get("VideoTexture")
	return VideoTextureFromJSObject(p, ba.ctx)
}

// VideoTextureFromJSObject returns a wrapped VideoTexture JavaScript class.
func VideoTextureFromJSObject(p js.Value, ctx js.Value) *VideoTexture {
	return &VideoTexture{Texture: TextureFromJSObject(p, ctx), ctx: ctx}
}

// VideoTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func VideoTextureArrayToJSArray(array []*VideoTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVideoTextureOpts contains optional parameters for NewVideoTexture.
type NewVideoTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Settings        js.Value
}

// NewVideoTexture returns a new VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#constructor
func (ba *Babylon) NewVideoTexture(name string, src []string, scene *Scene, opts *NewVideoTextureOpts) *VideoTexture {
	if opts == nil {
		opts = &NewVideoTextureOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, name)
	args = append(args, src)
	args = append(args, scene.JSObject())

	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}
	if opts.InvertY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InvertY)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	args = append(args, opts.Settings)

	p := ba.ctx.Get("VideoTexture").New(args...)
	return VideoTextureFromJSObject(p, ba.ctx)
}

// CreateFromStreamAsync calls the CreateFromStreamAsync method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#createfromstreamasync
func (v *VideoTexture) CreateFromStreamAsync(scene *Scene, stream js.Value) *Promise {

	args := make([]interface{}, 0, 2+0)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, stream)

	retVal := v.p.Call("CreateFromStreamAsync", args...)
	return PromiseFromJSObject(retVal, v.ctx)
}

// VideoTextureCreateFromWebCamOpts contains optional parameters for VideoTexture.CreateFromWebCam.
type VideoTextureCreateFromWebCamOpts struct {
	AudioConstaints js.Value
}

// CreateFromWebCam calls the CreateFromWebCam method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#createfromwebcam
func (v *VideoTexture) CreateFromWebCam(scene *Scene, onReady JSFunc, constraints js.Value, opts *VideoTextureCreateFromWebCamOpts) {
	if opts == nil {
		opts = &VideoTextureCreateFromWebCamOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, js.FuncOf(onReady))

	args = append(args, constraints)

	args = append(args, opts.AudioConstaints)

	v.p.Call("CreateFromWebCam", args...)
}

// VideoTextureCreateFromWebCamAsyncOpts contains optional parameters for VideoTexture.CreateFromWebCamAsync.
type VideoTextureCreateFromWebCamAsyncOpts struct {
	AudioConstaints js.Value
}

// CreateFromWebCamAsync calls the CreateFromWebCamAsync method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#createfromwebcamasync
func (v *VideoTexture) CreateFromWebCamAsync(scene *Scene, constraints js.Value, opts *VideoTextureCreateFromWebCamAsyncOpts) *Promise {
	if opts == nil {
		opts = &VideoTextureCreateFromWebCamAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, constraints)

	args = append(args, opts.AudioConstaints)

	retVal := v.p.Call("CreateFromWebCamAsync", args...)
	return PromiseFromJSObject(retVal, v.ctx)
}

// Dispose calls the Dispose method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#dispose
func (v *VideoTexture) Dispose() {

	v.p.Call("dispose")
}

// Update calls the Update method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#update
func (v *VideoTexture) Update() {

	v.p.Call("update")
}

// UpdateTexture calls the UpdateTexture method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#updatetexture
func (v *VideoTexture) UpdateTexture(isVisible bool) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, isVisible)

	v.p.Call("updateTexture", args...)
}

// UpdateURL calls the UpdateURL method on the VideoTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#updateurl
func (v *VideoTexture) UpdateURL(url string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, url)

	v.p.Call("updateURL", args...)
}

// AutoUpdateTexture returns the AutoUpdateTexture property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#autoupdatetexture
func (v *VideoTexture) AutoUpdateTexture() bool {
	retVal := v.p.Get("autoUpdateTexture")
	return retVal.Bool()
}

// SetAutoUpdateTexture sets the AutoUpdateTexture property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#autoupdatetexture
func (v *VideoTexture) SetAutoUpdateTexture(autoUpdateTexture bool) *VideoTexture {
	v.p.Set("autoUpdateTexture", autoUpdateTexture)
	return v
}

// OnUserActionRequestedObservable returns the OnUserActionRequestedObservable property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#onuseractionrequestedobservable
func (v *VideoTexture) OnUserActionRequestedObservable() *Observable {
	retVal := v.p.Get("onUserActionRequestedObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnUserActionRequestedObservable sets the OnUserActionRequestedObservable property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#onuseractionrequestedobservable
func (v *VideoTexture) SetOnUserActionRequestedObservable(onUserActionRequestedObservable *Observable) *VideoTexture {
	v.p.Set("onUserActionRequestedObservable", onUserActionRequestedObservable.JSObject())
	return v
}

// Video returns the Video property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#video
func (v *VideoTexture) Video() js.Value {
	retVal := v.p.Get("video")
	return retVal
}

// SetVideo sets the Video property of class VideoTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexture#video
func (v *VideoTexture) SetVideo(video js.Value) *VideoTexture {
	v.p.Set("video", video)
	return v
}
