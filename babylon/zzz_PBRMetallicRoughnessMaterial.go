// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRMetallicRoughnessMaterial represents a babylon.js PBRMetallicRoughnessMaterial.
// The PBR material of BJS following the metal roughness convention.
//
// This fits to the PBR convention in the GLTF definition:
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/tree/2.0/specification/2.0&#34;&gt;https://github.com/KhronosGroup/glTF/tree/2.0/specification/2.0&lt;/a&gt;
type PBRMetallicRoughnessMaterial struct {
	*PBRBaseSimpleMaterial
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRMetallicRoughnessMaterial) JSObject() js.Value { return p.p }

// PBRMetallicRoughnessMaterial returns a PBRMetallicRoughnessMaterial JavaScript class.
func (ba *Babylon) PBRMetallicRoughnessMaterial() *PBRMetallicRoughnessMaterial {
	p := ba.ctx.Get("PBRMetallicRoughnessMaterial")
	return PBRMetallicRoughnessMaterialFromJSObject(p, ba.ctx)
}

// PBRMetallicRoughnessMaterialFromJSObject returns a wrapped PBRMetallicRoughnessMaterial JavaScript class.
func PBRMetallicRoughnessMaterialFromJSObject(p js.Value, ctx js.Value) *PBRMetallicRoughnessMaterial {
	return &PBRMetallicRoughnessMaterial{PBRBaseSimpleMaterial: PBRBaseSimpleMaterialFromJSObject(p, ctx), ctx: ctx}
}

// NewPBRMetallicRoughnessMaterial returns a new PBRMetallicRoughnessMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrmetallicroughnessmaterial
func (ba *Babylon) NewPBRMetallicRoughnessMaterial(name string, scene *Scene) *PBRMetallicRoughnessMaterial {
	p := ba.ctx.Get("PBRMetallicRoughnessMaterial").New(name, scene.JSObject())
	return PBRMetallicRoughnessMaterialFromJSObject(p, ba.ctx)
}

// TODO: methods
