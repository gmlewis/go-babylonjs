// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BRDFTextureTools represents a babylon.js BRDFTextureTools.
// Class used to host texture specific utilities
type BRDFTextureTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BRDFTextureTools) JSObject() js.Value { return b.p }

// BRDFTextureTools returns a BRDFTextureTools JavaScript class.
func (ba *Babylon) BRDFTextureTools() *BRDFTextureTools {
	p := ba.ctx.Get("BRDFTextureTools")
	return BRDFTextureToolsFromJSObject(p, ba.ctx)
}

// BRDFTextureToolsFromJSObject returns a wrapped BRDFTextureTools JavaScript class.
func BRDFTextureToolsFromJSObject(p js.Value, ctx js.Value) *BRDFTextureTools {
	return &BRDFTextureTools{p: p, ctx: ctx}
}

// GetEnvironmentBRDFTexture calls the GetEnvironmentBRDFTexture method on the BRDFTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.brdftexturetools#getenvironmentbrdftexture
func (b *BRDFTextureTools) GetEnvironmentBRDFTexture(scene *Scene) *BaseTexture {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	retVal := b.p.Call("GetEnvironmentBRDFTexture", args...)
	return BaseTextureFromJSObject(retVal, b.ctx)
}

/*

 */
