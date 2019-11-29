// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiviewRenderTarget represents a babylon.js MultiviewRenderTarget.
// Renders to multiple views with a single draw call

//
// See: https://www.khronos.org/registry/webgl/extensions/WEBGL_multiview/
type MultiviewRenderTarget struct{ *RenderTargetTexture }

// JSObject returns the underlying js.Value.
func (m *MultiviewRenderTarget) JSObject() js.Value { return m.p }

// MultiviewRenderTarget returns a MultiviewRenderTarget JavaScript class.
func (b *Babylon) MultiviewRenderTarget() *MultiviewRenderTarget {
	p := b.ctx.Get("MultiviewRenderTarget")
	return MultiviewRenderTargetFromJSObject(p)
}

// MultiviewRenderTargetFromJSObject returns a wrapped MultiviewRenderTarget JavaScript class.
func MultiviewRenderTargetFromJSObject(p js.Value) *MultiviewRenderTarget {
	return &MultiviewRenderTarget{RenderTargetTextureFromJSObject(p)}
}

// NewMultiviewRenderTarget returns a new MultiviewRenderTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.multiviewrendertarget
func (b *Babylon) NewMultiviewRenderTarget(todo parameters) *MultiviewRenderTarget {
	p := b.ctx.Get("MultiviewRenderTarget").New(todo)
	return MultiviewRenderTargetFromJSObject(p)
}

// TODO: methods