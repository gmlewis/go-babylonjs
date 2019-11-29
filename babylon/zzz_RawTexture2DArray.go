// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RawTexture2DArray represents a babylon.js RawTexture2DArray.
// Class used to store 2D array textures containing user data
type RawTexture2DArray struct{ *Texture }

// JSObject returns the underlying js.Value.
func (r *RawTexture2DArray) JSObject() js.Value { return r.p }

// RawTexture2DArray returns a RawTexture2DArray JavaScript class.
func (b *Babylon) RawTexture2DArray() *RawTexture2DArray {
	p := b.ctx.Get("RawTexture2DArray")
	return RawTexture2DArrayFromJSObject(p)
}

// RawTexture2DArrayFromJSObject returns a wrapped RawTexture2DArray JavaScript class.
func RawTexture2DArrayFromJSObject(p js.Value) *RawTexture2DArray {
	return &RawTexture2DArray{TextureFromJSObject(p)}
}

// NewRawTexture2DArray returns a new RawTexture2DArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture2darray
func (b *Babylon) NewRawTexture2DArray(todo parameters) *RawTexture2DArray {
	p := b.ctx.Get("RawTexture2DArray").New(todo)
	return RawTexture2DArrayFromJSObject(p)
}

// TODO: methods