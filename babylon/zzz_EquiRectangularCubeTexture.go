// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EquiRectangularCubeTexture represents a babylon.js EquiRectangularCubeTexture.
// This represents a texture coming from an equirectangular image supported by the web browser canvas.
type EquiRectangularCubeTexture struct{ *BaseTexture }

// JSObject returns the underlying js.Value.
func (e *EquiRectangularCubeTexture) JSObject() js.Value { return e.p }

// EquiRectangularCubeTexture returns a EquiRectangularCubeTexture JavaScript class.
func (b *Babylon) EquiRectangularCubeTexture() *EquiRectangularCubeTexture {
	p := b.ctx.Get("EquiRectangularCubeTexture")
	return EquiRectangularCubeTextureFromJSObject(p)
}

// EquiRectangularCubeTextureFromJSObject returns a wrapped EquiRectangularCubeTexture JavaScript class.
func EquiRectangularCubeTextureFromJSObject(p js.Value) *EquiRectangularCubeTexture {
	return &EquiRectangularCubeTexture{BaseTextureFromJSObject(p)}
}

// NewEquiRectangularCubeTextureOpts contains optional parameters for NewEquiRectangularCubeTexture.
type NewEquiRectangularCubeTextureOpts struct {
	NoMipmap *bool

	GammaSpace *bool

	OnLoad *func()

	OnError *func()
}

// NewEquiRectangularCubeTexture returns a new EquiRectangularCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture
func (b *Babylon) NewEquiRectangularCubeTexture(url string, scene *Scene, size float64, opts *NewEquiRectangularCubeTextureOpts) *EquiRectangularCubeTexture {
	if opts == nil {
		opts = &NewEquiRectangularCubeTextureOpts{}
	}

	p := b.ctx.Get("EquiRectangularCubeTexture").New(url, scene.JSObject(), size, opts.NoMipmap.JSObject(), opts.GammaSpace.JSObject(), opts.OnLoad, opts.OnError)
	return EquiRectangularCubeTextureFromJSObject(p)
}

// TODO: methods
