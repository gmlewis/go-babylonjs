// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ProceduralTexture represents a babylon.js ProceduralTexture.
// Procedural texturing is a way to programmatically create a texture. There are 2 types of procedural textures: code-only, and code that references some classic 2D images, sometimes calmpler&amp;#39; images.
// This is the base class of any Procedural texture and contains most of the shareable code.

//
// See: http://doc.babylonjs.com/how_to/how_to_use_procedural_textures
type ProceduralTexture struct{ *Texture }

// JSObject returns the underlying js.Value.
func (p *ProceduralTexture) JSObject() js.Value { return p.p }

// ProceduralTexture returns a ProceduralTexture JavaScript class.
func (b *Babylon) ProceduralTexture() *ProceduralTexture {
	p := b.ctx.Get("ProceduralTexture")
	return ProceduralTextureFromJSObject(p)
}

// ProceduralTextureFromJSObject returns a wrapped ProceduralTexture JavaScript class.
func ProceduralTextureFromJSObject(p js.Value) *ProceduralTexture {
	return &ProceduralTexture{TextureFromJSObject(p)}
}

// NewProceduralTexture returns a new ProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.proceduraltexture
func (b *Babylon) NewProceduralTexture(todo parameters) *ProceduralTexture {
	p := b.ctx.Get("ProceduralTexture").New(todo)
	return ProceduralTextureFromJSObject(p)
}

// TODO: methods