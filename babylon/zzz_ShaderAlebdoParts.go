// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShaderAlebdoParts represents a babylon.js ShaderAlebdoParts.
//
type ShaderAlebdoParts struct{}

// JSObject returns the underlying js.Value.
func (s *ShaderAlebdoParts) JSObject() js.Value { return s.p }

// ShaderAlebdoParts returns a ShaderAlebdoParts JavaScript class.
func (b *Babylon) ShaderAlebdoParts() *ShaderAlebdoParts {
	p := b.ctx.Get("ShaderAlebdoParts")
	return ShaderAlebdoPartsFromJSObject(p)
}

// ShaderAlebdoPartsFromJSObject returns a wrapped ShaderAlebdoParts JavaScript class.
func ShaderAlebdoPartsFromJSObject(p js.Value) *ShaderAlebdoParts {
	return &ShaderAlebdoParts{p: p}
}

// NewShaderAlebdoParts returns a new ShaderAlebdoParts object.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts
func (b *Babylon) NewShaderAlebdoParts(todo parameters) *ShaderAlebdoParts {
	p := b.ctx.Get("ShaderAlebdoParts").New(todo)
	return ShaderAlebdoPartsFromJSObject(p)
}

// TODO: methods