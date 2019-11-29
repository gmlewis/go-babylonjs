// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageAssetTask represents a babylon.js ImageAssetTask.
// Define a task used by AssetsManager to load images
type ImageAssetTask struct{ *AbstractAssetTask }

// JSObject returns the underlying js.Value.
func (i *ImageAssetTask) JSObject() js.Value { return i.p }

// ImageAssetTask returns a ImageAssetTask JavaScript class.
func (b *Babylon) ImageAssetTask() *ImageAssetTask {
	p := b.ctx.Get("ImageAssetTask")
	return ImageAssetTaskFromJSObject(p)
}

// ImageAssetTaskFromJSObject returns a wrapped ImageAssetTask JavaScript class.
func ImageAssetTaskFromJSObject(p js.Value) *ImageAssetTask {
	return &ImageAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewImageAssetTask returns a new ImageAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageassettask
func (b *Babylon) NewImageAssetTask(todo parameters) *ImageAssetTask {
	p := b.ctx.Get("ImageAssetTask").New(todo)
	return ImageAssetTaskFromJSObject(p)
}

// TODO: methods