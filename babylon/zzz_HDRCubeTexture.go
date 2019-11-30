// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HDRCubeTexture represents a babylon.js HDRCubeTexture.
// This represents a texture coming from an HDR input.
//
// The only supported format is currently panorama picture stored in RGBE format.
// Example of such files can be found on HDRLib: &lt;a href=&#34;http://hdrlib.com/&#34;&gt;http://hdrlib.com/&lt;/a&gt;
type HDRCubeTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HDRCubeTexture) JSObject() js.Value { return h.p }

// HDRCubeTexture returns a HDRCubeTexture JavaScript class.
func (ba *Babylon) HDRCubeTexture() *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture")
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// HDRCubeTextureFromJSObject returns a wrapped HDRCubeTexture JavaScript class.
func HDRCubeTextureFromJSObject(p js.Value, ctx js.Value) *HDRCubeTexture {
	return &HDRCubeTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// NewHDRCubeTextureOpts contains optional parameters for NewHDRCubeTexture.
type NewHDRCubeTextureOpts struct {
	NoMipmap *JSBool

	GenerateHarmonics *JSBool

	GammaSpace *JSBool

	Reserved *JSBool

	OnLoad *func()

	OnError *func()
}

// NewHDRCubeTexture returns a new HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture
func (ba *Babylon) NewHDRCubeTexture(url string, scene *Scene, size float64, opts *NewHDRCubeTextureOpts) *HDRCubeTexture {
	if opts == nil {
		opts = &NewHDRCubeTextureOpts{}
	}

	p := ba.ctx.Get("HDRCubeTexture").New(url, scene.JSObject(), size, opts.NoMipmap.JSObject(), opts.GenerateHarmonics.JSObject(), opts.GammaSpace.JSObject(), opts.Reserved.JSObject(), opts.OnLoad, opts.OnError)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// TODO: methods
