// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FilesInput represents a babylon.js FilesInput.
// Class used to help managing file picking and drag&amp;#39;n&amp;#39;drop
type FilesInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FilesInput) JSObject() js.Value { return f.p }

// FilesInput returns a FilesInput JavaScript class.
func (ba *Babylon) FilesInput() *FilesInput {
	p := ba.ctx.Get("FilesInput")
	return FilesInputFromJSObject(p, ba.ctx)
}

// FilesInputFromJSObject returns a wrapped FilesInput JavaScript class.
func FilesInputFromJSObject(p js.Value, ctx js.Value) *FilesInput {
	return &FilesInput{p: p, ctx: ctx}
}

// NewFilesInput returns a new FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput
func (ba *Babylon) NewFilesInput(engine *Engine, scene *Scene, sceneLoadedCallback func(), progressCallback func(), additionalRenderLoopLogicCallback func(), textureLoadingCallback func(), startingProcessingFilesCallback func(), onReloadCallback func(), errorCallback func()) *FilesInput {
	p := ba.ctx.Get("FilesInput").New(engine.JSObject(), scene.JSObject(), sceneLoadedCallback, progressCallback, additionalRenderLoopLogicCallback, textureLoadingCallback, startingProcessingFilesCallback, onReloadCallback, errorCallback)
	return FilesInputFromJSObject(p, ba.ctx)
}

// TODO: methods
