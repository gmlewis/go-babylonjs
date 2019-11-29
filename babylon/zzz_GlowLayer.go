// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GlowLayer represents a babylon.js GlowLayer.
// The glow layer Helps adding a glow effect around the emissive parts of a mesh.

//
// Documentation: &lt;a href=&#34;https://doc.babylonjs.com/how_to/glow_layer&#34;&gt;https://doc.babylonjs.com/how_to/glow_layer&lt;/a&gt;
type GlowLayer struct{ *EffectLayer }

// JSObject returns the underlying js.Value.
func (g *GlowLayer) JSObject() js.Value { return g.p }

// GlowLayer returns a GlowLayer JavaScript class.
func (b *Babylon) GlowLayer() *GlowLayer {
	p := b.ctx.Get("GlowLayer")
	return GlowLayerFromJSObject(p)
}

// GlowLayerFromJSObject returns a wrapped GlowLayer JavaScript class.
func GlowLayerFromJSObject(p js.Value) *GlowLayer {
	return &GlowLayer{EffectLayerFromJSObject(p)}
}

// NewGlowLayer returns a new GlowLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.glowlayer
func (b *Babylon) NewGlowLayer(todo parameters) *GlowLayer {
	p := b.ctx.Get("GlowLayer").New(todo)
	return GlowLayerFromJSObject(p)
}

// TODO: methods