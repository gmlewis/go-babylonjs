// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectRenderer represents a babylon.js EffectRenderer.
// Helper class to render one or more effects
type EffectRenderer struct{}

// JSObject returns the underlying js.Value.
func (e *EffectRenderer) JSObject() js.Value { return e.p }

// EffectRenderer returns a EffectRenderer JavaScript class.
func (b *Babylon) EffectRenderer() *EffectRenderer {
	p := b.ctx.Get("EffectRenderer")
	return EffectRendererFromJSObject(p)
}

// EffectRendererFromJSObject returns a wrapped EffectRenderer JavaScript class.
func EffectRendererFromJSObject(p js.Value) *EffectRenderer {
	return &EffectRenderer{p: p}
}

// NewEffectRenderer returns a new EffectRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectrenderer
func (b *Babylon) NewEffectRenderer(todo parameters) *EffectRenderer {
	p := b.ctx.Get("EffectRenderer").New(todo)
	return EffectRendererFromJSObject(p)
}

// TODO: methods
