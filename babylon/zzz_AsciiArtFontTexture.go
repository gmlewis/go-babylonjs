// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsciiArtFontTexture represents a babylon.js AsciiArtFontTexture.
// AsciiArtFontTexture is the helper class used to easily create your ascii art font texture.
//
// It basically takes care rendering the font front the given font size to a texture.
// This is used later on in the postprocess.
type AsciiArtFontTexture struct{ *BaseTexture }

// JSObject returns the underlying js.Value.
func (a *AsciiArtFontTexture) JSObject() js.Value { return a.p }

// AsciiArtFontTexture returns a AsciiArtFontTexture JavaScript class.
func (b *Babylon) AsciiArtFontTexture() *AsciiArtFontTexture {
	p := b.ctx.Get("AsciiArtFontTexture")
	return AsciiArtFontTextureFromJSObject(p)
}

// AsciiArtFontTextureFromJSObject returns a wrapped AsciiArtFontTexture JavaScript class.
func AsciiArtFontTextureFromJSObject(p js.Value) *AsciiArtFontTexture {
	return &AsciiArtFontTexture{BaseTextureFromJSObject(p)}
}

// NewAsciiArtFontTextureOpts contains optional parameters for NewAsciiArtFontTexture.
type NewAsciiArtFontTextureOpts struct {
	Scene *Scene
}

// NewAsciiArtFontTexture returns a new AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture
func (b *Babylon) NewAsciiArtFontTexture(name string, font string, text string, opts *NewAsciiArtFontTextureOpts) *AsciiArtFontTexture {
	if opts == nil {
		opts = &NewAsciiArtFontTextureOpts{}
	}

	p := b.ctx.Get("AsciiArtFontTexture").New(name, font, text, opts.Scene.JSObject())
	return AsciiArtFontTextureFromJSObject(p)
}

// TODO: methods
