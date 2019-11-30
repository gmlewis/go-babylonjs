// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CloudProceduralTexture represents a babylon.js CloudProceduralTexture.
//
type CloudProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (c *CloudProceduralTexture) JSObject() js.Value { return c.p }

// CloudProceduralTexture returns a CloudProceduralTexture JavaScript class.
func (b *Babylon) CloudProceduralTexture() *CloudProceduralTexture {
	p := b.ctx.Get("CloudProceduralTexture")
	return CloudProceduralTextureFromJSObject(p)
}

// CloudProceduralTextureFromJSObject returns a wrapped CloudProceduralTexture JavaScript class.
func CloudProceduralTextureFromJSObject(p js.Value) *CloudProceduralTexture {
	return &CloudProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewCloudProceduralTextureOpts contains optional parameters for NewCloudProceduralTexture.
type NewCloudProceduralTextureOpts struct {
	FallbackTexture *Texture

	GenerateMipMaps *bool
}

// NewCloudProceduralTexture returns a new CloudProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture
func (b *Babylon) NewCloudProceduralTexture(name string, size float64, scene *Scene, opts *NewCloudProceduralTextureOpts) *CloudProceduralTexture {
	if opts == nil {
		opts = &NewCloudProceduralTextureOpts{}
	}

	p := b.ctx.Get("CloudProceduralTexture").New(name, size, scene.JSObject(), opts.FallbackTexture.JSObject(), opts.GenerateMipMaps.JSObject())
	return CloudProceduralTextureFromJSObject(p)
}

// TODO: methods
