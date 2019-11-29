// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CubeTexture represents a babylon.js CubeTexture.
// Class for creating a cube texture
type CubeTexture struct{ *BaseTexture }

// JSObject returns the underlying js.Value.
func (c *CubeTexture) JSObject() js.Value { return c.p }

// CubeTexture returns a CubeTexture JavaScript class.
func (b *Babylon) CubeTexture() *CubeTexture {
	p := b.ctx.Get("CubeTexture")
	return CubeTextureFromJSObject(p)
}

// CubeTextureFromJSObject returns a wrapped CubeTexture JavaScript class.
func CubeTextureFromJSObject(p js.Value) *CubeTexture {
	return &CubeTexture{BaseTextureFromJSObject(p)}
}

// NewCubeTexture returns a new CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture
func (b *Babylon) NewCubeTexture(todo parameters) *CubeTexture {
	p := b.ctx.Get("CubeTexture").New(todo)
	return CubeTextureFromJSObject(p)
}

// TODO: methods
