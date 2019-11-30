// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorGradingTexture represents a babylon.js ColorGradingTexture.
// This represents a color grading texture. This acts as a lookup table LUT, useful during post process
// It can help converting any input color in a desired output one. This can then be used to create effects
// from sepia, black and white to sixties or futuristic rendering...
//
// The only supported format is currently 3dl.
// More information on LUT: &lt;a href=&#34;https://en.wikipedia.org/wiki/3D_lookup_table&#34;&gt;https://en.wikipedia.org/wiki/3D_lookup_table&lt;/a&gt;
type ColorGradingTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ColorGradingTexture) JSObject() js.Value { return c.p }

// ColorGradingTexture returns a ColorGradingTexture JavaScript class.
func (ba *Babylon) ColorGradingTexture() *ColorGradingTexture {
	p := ba.ctx.Get("ColorGradingTexture")
	return ColorGradingTextureFromJSObject(p, ba.ctx)
}

// ColorGradingTextureFromJSObject returns a wrapped ColorGradingTexture JavaScript class.
func ColorGradingTextureFromJSObject(p js.Value, ctx js.Value) *ColorGradingTexture {
	return &ColorGradingTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// NewColorGradingTexture returns a new ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture
func (ba *Babylon) NewColorGradingTexture(url string, scene *Scene) *ColorGradingTexture {
	p := ba.ctx.Get("ColorGradingTexture").New(url, scene.JSObject())
	return ColorGradingTextureFromJSObject(p, ba.ctx)
}

// TODO: methods
