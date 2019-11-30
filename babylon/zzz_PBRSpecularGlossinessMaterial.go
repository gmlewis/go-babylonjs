// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRSpecularGlossinessMaterial represents a babylon.js PBRSpecularGlossinessMaterial.
// The PBR material of BJS following the specular glossiness convention.
//
// This fits to the PBR convention in the GLTF definition:
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/tree/2.0/extensions/Khronos/KHR_materials_pbrSpecularGlossiness&#34;&gt;https://github.com/KhronosGroup/glTF/tree/2.0/extensions/Khronos/KHR_materials_pbrSpecularGlossiness&lt;/a&gt;
type PBRSpecularGlossinessMaterial struct{ *PBRBaseSimpleMaterial }

// JSObject returns the underlying js.Value.
func (p *PBRSpecularGlossinessMaterial) JSObject() js.Value { return p.p }

// PBRSpecularGlossinessMaterial returns a PBRSpecularGlossinessMaterial JavaScript class.
func (b *Babylon) PBRSpecularGlossinessMaterial() *PBRSpecularGlossinessMaterial {
	p := b.ctx.Get("PBRSpecularGlossinessMaterial")
	return PBRSpecularGlossinessMaterialFromJSObject(p)
}

// PBRSpecularGlossinessMaterialFromJSObject returns a wrapped PBRSpecularGlossinessMaterial JavaScript class.
func PBRSpecularGlossinessMaterialFromJSObject(p js.Value) *PBRSpecularGlossinessMaterial {
	return &PBRSpecularGlossinessMaterial{PBRBaseSimpleMaterialFromJSObject(p)}
}

// NewPBRSpecularGlossinessMaterial returns a new PBRSpecularGlossinessMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrspecularglossinessmaterial
func (b *Babylon) NewPBRSpecularGlossinessMaterial(name string, scene *Scene) *PBRSpecularGlossinessMaterial {
	p := b.ctx.Get("PBRSpecularGlossinessMaterial").New(name, scene.JSObject())
	return PBRSpecularGlossinessMaterialFromJSObject(p)
}

// TODO: methods
