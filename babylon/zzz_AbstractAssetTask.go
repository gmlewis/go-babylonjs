// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AbstractAssetTask represents a babylon.js AbstractAssetTask.
// Define an abstract asset task used with a AssetsManager class to load assets into a scene
type AbstractAssetTask struct{}

// JSObject returns the underlying js.Value.
func (a *AbstractAssetTask) JSObject() js.Value { return a.p }

// AbstractAssetTask returns a AbstractAssetTask JavaScript class.
func (b *Babylon) AbstractAssetTask() *AbstractAssetTask {
	p := b.ctx.Get("AbstractAssetTask")
	return AbstractAssetTaskFromJSObject(p)
}

// AbstractAssetTaskFromJSObject returns a wrapped AbstractAssetTask JavaScript class.
func AbstractAssetTaskFromJSObject(p js.Value) *AbstractAssetTask {
	return &AbstractAssetTask{p: p}
}

// NewAbstractAssetTask returns a new AbstractAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractassettask
func (b *Babylon) NewAbstractAssetTask(todo parameters) *AbstractAssetTask {
	p := b.ctx.Get("AbstractAssetTask").New(todo)
	return AbstractAssetTaskFromJSObject(p)
}

// TODO: methods
