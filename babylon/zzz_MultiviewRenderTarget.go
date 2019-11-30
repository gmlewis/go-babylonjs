// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiviewRenderTarget represents a babylon.js MultiviewRenderTarget.
// Renders to multiple views with a single draw call
//
// See: https://www.khronos.org/registry/webgl/extensions/WEBGL_multiview/
type MultiviewRenderTarget struct {
	*RenderTargetTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiviewRenderTarget) JSObject() js.Value { return m.p }

// MultiviewRenderTarget returns a MultiviewRenderTarget JavaScript class.
func (ba *Babylon) MultiviewRenderTarget() *MultiviewRenderTarget {
	p := ba.ctx.Get("MultiviewRenderTarget")
	return MultiviewRenderTargetFromJSObject(p, ba.ctx)
}

// MultiviewRenderTargetFromJSObject returns a wrapped MultiviewRenderTarget JavaScript class.
func MultiviewRenderTargetFromJSObject(p js.Value, ctx js.Value) *MultiviewRenderTarget {
	return &MultiviewRenderTarget{RenderTargetTexture: RenderTargetTextureFromJSObject(p, ctx), ctx: ctx}
}

// NewMultiviewRenderTargetOpts contains optional parameters for NewMultiviewRenderTarget.
type NewMultiviewRenderTargetOpts struct {
	Size *JSFloat64
}

// NewMultiviewRenderTarget returns a new MultiviewRenderTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.multiviewrendertarget
func (ba *Babylon) NewMultiviewRenderTarget(scene *Scene, opts *NewMultiviewRenderTargetOpts) *MultiviewRenderTarget {
	if opts == nil {
		opts = &NewMultiviewRenderTargetOpts{}
	}

	p := ba.ctx.Get("MultiviewRenderTarget").New(scene.JSObject(), opts.Size.JSObject())
	return MultiviewRenderTargetFromJSObject(p, ba.ctx)
}

// TODO: methods
