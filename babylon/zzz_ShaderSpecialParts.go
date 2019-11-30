// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShaderSpecialParts represents a babylon.js ShaderSpecialParts.
//
type ShaderSpecialParts struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ShaderSpecialParts) JSObject() js.Value { return s.p }

// ShaderSpecialParts returns a ShaderSpecialParts JavaScript class.
func (ba *Babylon) ShaderSpecialParts() *ShaderSpecialParts {
	p := ba.ctx.Get("ShaderSpecialParts")
	return ShaderSpecialPartsFromJSObject(p, ba.ctx)
}

// ShaderSpecialPartsFromJSObject returns a wrapped ShaderSpecialParts JavaScript class.
func ShaderSpecialPartsFromJSObject(p js.Value, ctx js.Value) *ShaderSpecialParts {
	return &ShaderSpecialParts{p: p, ctx: ctx}
}

// NewShaderSpecialParts returns a new ShaderSpecialParts object.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderspecialparts
func (ba *Babylon) NewShaderSpecialParts() *ShaderSpecialParts {
	p := ba.ctx.Get("ShaderSpecialParts").New()
	return ShaderSpecialPartsFromJSObject(p, ba.ctx)
}

// TODO: methods
