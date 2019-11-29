// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StarfieldProceduralTexture represents a babylon.js StarfieldProceduralTexture.
//
type StarfieldProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (s *StarfieldProceduralTexture) JSObject() js.Value { return s.p }

// StarfieldProceduralTexture returns a StarfieldProceduralTexture JavaScript class.
func (b *Babylon) StarfieldProceduralTexture() *StarfieldProceduralTexture {
	p := b.ctx.Get("StarfieldProceduralTexture")
	return StarfieldProceduralTextureFromJSObject(p)
}

// StarfieldProceduralTextureFromJSObject returns a wrapped StarfieldProceduralTexture JavaScript class.
func StarfieldProceduralTextureFromJSObject(p js.Value) *StarfieldProceduralTexture {
	return &StarfieldProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewStarfieldProceduralTexture returns a new StarfieldProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.starfieldproceduraltexture
func (b *Babylon) NewStarfieldProceduralTexture(todo parameters) *StarfieldProceduralTexture {
	p := b.ctx.Get("StarfieldProceduralTexture").New(todo)
	return StarfieldProceduralTextureFromJSObject(p)
}

// TODO: methods
