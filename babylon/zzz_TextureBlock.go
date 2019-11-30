// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextureBlock represents a babylon.js TextureBlock.
// Block used to read a texture from a sampler
type TextureBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TextureBlock) JSObject() js.Value { return t.p }

// TextureBlock returns a TextureBlock JavaScript class.
func (ba *Babylon) TextureBlock() *TextureBlock {
	p := ba.ctx.Get("TextureBlock")
	return TextureBlockFromJSObject(p, ba.ctx)
}

// TextureBlockFromJSObject returns a wrapped TextureBlock JavaScript class.
func TextureBlockFromJSObject(p js.Value, ctx js.Value) *TextureBlock {
	return &TextureBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewTextureBlock returns a new TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock
func (ba *Babylon) NewTextureBlock(name string) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(name)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
