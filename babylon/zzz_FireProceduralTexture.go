// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FireProceduralTexture represents a babylon.js FireProceduralTexture.
//
type FireProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (f *FireProceduralTexture) JSObject() js.Value { return f.p }

// FireProceduralTexture returns a FireProceduralTexture JavaScript class.
func (b *Babylon) FireProceduralTexture() *FireProceduralTexture {
	p := b.ctx.Get("FireProceduralTexture")
	return FireProceduralTextureFromJSObject(p)
}

// FireProceduralTextureFromJSObject returns a wrapped FireProceduralTexture JavaScript class.
func FireProceduralTextureFromJSObject(p js.Value) *FireProceduralTexture {
	return &FireProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewFireProceduralTextureOpts contains optional parameters for NewFireProceduralTexture.
type NewFireProceduralTextureOpts struct {
	FallbackTexture *Texture

	GenerateMipMaps *bool
}

// NewFireProceduralTexture returns a new FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture
func (b *Babylon) NewFireProceduralTexture(name string, size float64, scene *Scene, opts *NewFireProceduralTextureOpts) *FireProceduralTexture {
	if opts == nil {
		opts = &NewFireProceduralTextureOpts{}
	}

	p := b.ctx.Get("FireProceduralTexture").New(name, size, scene.JSObject(), opts.FallbackTexture.JSObject(), opts.GenerateMipMaps.JSObject())
	return FireProceduralTextureFromJSObject(p)
}

// TODO: methods
