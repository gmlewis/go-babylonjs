// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RawCubeTexture represents a babylon.js RawCubeTexture.
// Raw cube texture where the raw buffers are passed in
type RawCubeTexture struct{ *CubeTexture }

// JSObject returns the underlying js.Value.
func (r *RawCubeTexture) JSObject() js.Value { return r.p }

// RawCubeTexture returns a RawCubeTexture JavaScript class.
func (b *Babylon) RawCubeTexture() *RawCubeTexture {
	p := b.ctx.Get("RawCubeTexture")
	return RawCubeTextureFromJSObject(p)
}

// RawCubeTextureFromJSObject returns a wrapped RawCubeTexture JavaScript class.
func RawCubeTextureFromJSObject(p js.Value) *RawCubeTexture {
	return &RawCubeTexture{CubeTextureFromJSObject(p)}
}

// NewRawCubeTextureOpts contains optional parameters for NewRawCubeTexture.
type NewRawCubeTextureOpts struct {
	Format *float64

	Type *float64

	GenerateMipMaps *bool

	InvertY *bool

	SamplingMode *float64

	Compression *string
}

// NewRawCubeTexture returns a new RawCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawcubetexture
func (b *Babylon) NewRawCubeTexture(scene *Scene, data ArrayBufferView, size float64, opts *NewRawCubeTextureOpts) *RawCubeTexture {
	if opts == nil {
		opts = &NewRawCubeTextureOpts{}
	}

	p := b.ctx.Get("RawCubeTexture").New(scene.JSObject(), data.JSObject(), size, opts.Format, opts.Type, opts.GenerateMipMaps.JSObject(), opts.InvertY.JSObject(), opts.SamplingMode, opts.Compression)
	return RawCubeTextureFromJSObject(p)
}

// TODO: methods
