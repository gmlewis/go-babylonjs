// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FilesInput represents a babylon.js FilesInput.
// Class used to help managing file picking and drag&amp;#39;n&amp;#39;drop
type FilesInput struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (f *FilesInput) JSObject() js.Value { return f.p }

// FilesInput returns a FilesInput JavaScript class.
func (b *Babylon) FilesInput() *FilesInput {
	p := b.ctx.Get("FilesInput")
	return FilesInputFromJSObject(p)
}

// FilesInputFromJSObject returns a wrapped FilesInput JavaScript class.
func FilesInputFromJSObject(p js.Value) *FilesInput {
	return &FilesInput{p: p}
}

// NewFilesInput returns a new FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput
func (b *Babylon) NewFilesInput(engine *Engine, scene *Scene, sceneLoadedCallback func(), progressCallback func(), additionalRenderLoopLogicCallback func(), textureLoadingCallback func(), startingProcessingFilesCallback func(), onReloadCallback func(), errorCallback func()) *FilesInput {
	p := b.ctx.Get("FilesInput").New(engine.JSObject(), scene.JSObject(), sceneLoadedCallback, progressCallback, additionalRenderLoopLogicCallback, textureLoadingCallback, startingProcessingFilesCallback, onReloadCallback, errorCallback)
	return FilesInputFromJSObject(p)
}

// TODO: methods
