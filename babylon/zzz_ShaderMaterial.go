// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShaderMaterial represents a babylon.js ShaderMaterial.
// The ShaderMaterial object has the necessary methods to pass data from your scene to the Vertex and Fragment Shaders and returns a material that can be applied to any mesh.

//
// This returned material effects how the mesh will look based on the code in the shaders.

//
// See: http://doc.babylonjs.com/how_to/shader_material
type ShaderMaterial struct{ *Material }

// JSObject returns the underlying js.Value.
func (s *ShaderMaterial) JSObject() js.Value { return s.p }

// ShaderMaterial returns a ShaderMaterial JavaScript class.
func (b *Babylon) ShaderMaterial() *ShaderMaterial {
	p := b.ctx.Get("ShaderMaterial")
	return ShaderMaterialFromJSObject(p)
}

// ShaderMaterialFromJSObject returns a wrapped ShaderMaterial JavaScript class.
func ShaderMaterialFromJSObject(p js.Value) *ShaderMaterial {
	return &ShaderMaterial{MaterialFromJSObject(p)}
}

// NewShaderMaterial returns a new ShaderMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadermaterial
func (b *Babylon) NewShaderMaterial(todo parameters) *ShaderMaterial {
	p := b.ctx.Get("ShaderMaterial").New(todo)
	return ShaderMaterialFromJSObject(p)
}

// TODO: methods
