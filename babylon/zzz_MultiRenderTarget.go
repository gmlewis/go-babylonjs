// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiRenderTarget represents a babylon.js MultiRenderTarget.
// A multi render target, like a render target provides the ability to render to a texture.
// Unlike the render target, it can render to several draw buffers in one draw.
// This is specially interesting in deferred rendering or for any effects requiring more than
// just one color from a single pass.
type MultiRenderTarget struct{ *RenderTargetTexture }

// JSObject returns the underlying js.Value.
func (m *MultiRenderTarget) JSObject() js.Value { return m.p }

// MultiRenderTarget returns a MultiRenderTarget JavaScript class.
func (b *Babylon) MultiRenderTarget() *MultiRenderTarget {
	p := b.ctx.Get("MultiRenderTarget")
	return MultiRenderTargetFromJSObject(p)
}

// MultiRenderTargetFromJSObject returns a wrapped MultiRenderTarget JavaScript class.
func MultiRenderTargetFromJSObject(p js.Value) *MultiRenderTarget {
	return &MultiRenderTarget{RenderTargetTextureFromJSObject(p)}
}

// NewMultiRenderTargetOpts contains optional parameters for NewMultiRenderTarget.
type NewMultiRenderTargetOpts struct {
	Options js.Value
}

// NewMultiRenderTarget returns a new MultiRenderTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.multirendertarget
func (b *Babylon) NewMultiRenderTarget(name string, size interface{}, count float64, scene *Scene, opts *NewMultiRenderTargetOpts) *MultiRenderTarget {
	if opts == nil {
		opts = &NewMultiRenderTargetOpts{}
	}

	p := b.ctx.Get("MultiRenderTarget").New(name, size, count, scene.JSObject(), opts.Options)
	return MultiRenderTargetFromJSObject(p)
}

// TODO: methods
