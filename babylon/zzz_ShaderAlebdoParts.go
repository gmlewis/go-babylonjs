// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShaderAlebdoParts represents a babylon.js ShaderAlebdoParts.
//
type ShaderAlebdoParts struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ShaderAlebdoParts) JSObject() js.Value { return s.p }

// ShaderAlebdoParts returns a ShaderAlebdoParts JavaScript class.
func (ba *Babylon) ShaderAlebdoParts() *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts")
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// ShaderAlebdoPartsFromJSObject returns a wrapped ShaderAlebdoParts JavaScript class.
func ShaderAlebdoPartsFromJSObject(p js.Value, ctx js.Value) *ShaderAlebdoParts {
	return &ShaderAlebdoParts{p: p, ctx: ctx}
}

// NewShaderAlebdoParts returns a new ShaderAlebdoParts object.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts
func (ba *Babylon) NewShaderAlebdoParts() *ShaderAlebdoParts {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("ShaderAlebdoParts").New(args...)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

/*

// Fragment_Before_Fog returns the Fragment_Before_Fog property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_fog
func (s *ShaderAlebdoParts) Fragment_Before_Fog(Fragment_Before_Fog string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_Fog)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Before_Fog sets the Fragment_Before_Fog property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_fog
func (s *ShaderAlebdoParts) SetFragment_Before_Fog(Fragment_Before_Fog string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_Fog)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Before_FragColor returns the Fragment_Before_FragColor property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_fragcolor
func (s *ShaderAlebdoParts) Fragment_Before_FragColor(Fragment_Before_FragColor string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_FragColor)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Before_FragColor sets the Fragment_Before_FragColor property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_fragcolor
func (s *ShaderAlebdoParts) SetFragment_Before_FragColor(Fragment_Before_FragColor string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_FragColor)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Before_Lights returns the Fragment_Before_Lights property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_lights
func (s *ShaderAlebdoParts) Fragment_Before_Lights(Fragment_Before_Lights string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_Lights)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Before_Lights sets the Fragment_Before_Lights property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_before_lights
func (s *ShaderAlebdoParts) SetFragment_Before_Lights(Fragment_Before_Lights string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Before_Lights)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Begin returns the Fragment_Begin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_begin
func (s *ShaderAlebdoParts) Fragment_Begin(Fragment_Begin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Begin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Begin sets the Fragment_Begin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_begin
func (s *ShaderAlebdoParts) SetFragment_Begin(Fragment_Begin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Begin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Custom_Albedo returns the Fragment_Custom_Albedo property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_albedo
func (s *ShaderAlebdoParts) Fragment_Custom_Albedo(Fragment_Custom_Albedo string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_Albedo)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Custom_Albedo sets the Fragment_Custom_Albedo property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_albedo
func (s *ShaderAlebdoParts) SetFragment_Custom_Albedo(Fragment_Custom_Albedo string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_Albedo)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Custom_Alpha returns the Fragment_Custom_Alpha property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_alpha
func (s *ShaderAlebdoParts) Fragment_Custom_Alpha(Fragment_Custom_Alpha string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_Alpha)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Custom_Alpha sets the Fragment_Custom_Alpha property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_alpha
func (s *ShaderAlebdoParts) SetFragment_Custom_Alpha(Fragment_Custom_Alpha string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_Alpha)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Custom_MetallicRoughness returns the Fragment_Custom_MetallicRoughness property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_metallicroughness
func (s *ShaderAlebdoParts) Fragment_Custom_MetallicRoughness(Fragment_Custom_MetallicRoughness string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_MetallicRoughness)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Custom_MetallicRoughness sets the Fragment_Custom_MetallicRoughness property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_metallicroughness
func (s *ShaderAlebdoParts) SetFragment_Custom_MetallicRoughness(Fragment_Custom_MetallicRoughness string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_MetallicRoughness)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Custom_MicroSurface returns the Fragment_Custom_MicroSurface property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_microsurface
func (s *ShaderAlebdoParts) Fragment_Custom_MicroSurface(Fragment_Custom_MicroSurface string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_MicroSurface)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Custom_MicroSurface sets the Fragment_Custom_MicroSurface property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_custom_microsurface
func (s *ShaderAlebdoParts) SetFragment_Custom_MicroSurface(Fragment_Custom_MicroSurface string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Custom_MicroSurface)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_Definitions returns the Fragment_Definitions property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_definitions
func (s *ShaderAlebdoParts) Fragment_Definitions(Fragment_Definitions string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Definitions)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_Definitions sets the Fragment_Definitions property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_definitions
func (s *ShaderAlebdoParts) SetFragment_Definitions(Fragment_Definitions string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_Definitions)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Fragment_MainBegin returns the Fragment_MainBegin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_mainbegin
func (s *ShaderAlebdoParts) Fragment_MainBegin(Fragment_MainBegin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_MainBegin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetFragment_MainBegin sets the Fragment_MainBegin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#fragment_mainbegin
func (s *ShaderAlebdoParts) SetFragment_MainBegin(Fragment_MainBegin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Fragment_MainBegin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_Before_NormalUpdated returns the Vertex_Before_NormalUpdated property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_before_normalupdated
func (s *ShaderAlebdoParts) Vertex_Before_NormalUpdated(Vertex_Before_NormalUpdated string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Before_NormalUpdated)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_Before_NormalUpdated sets the Vertex_Before_NormalUpdated property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_before_normalupdated
func (s *ShaderAlebdoParts) SetVertex_Before_NormalUpdated(Vertex_Before_NormalUpdated string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Before_NormalUpdated)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_Before_PositionUpdated returns the Vertex_Before_PositionUpdated property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_before_positionupdated
func (s *ShaderAlebdoParts) Vertex_Before_PositionUpdated(Vertex_Before_PositionUpdated string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Before_PositionUpdated)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_Before_PositionUpdated sets the Vertex_Before_PositionUpdated property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_before_positionupdated
func (s *ShaderAlebdoParts) SetVertex_Before_PositionUpdated(Vertex_Before_PositionUpdated string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Before_PositionUpdated)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_Begin returns the Vertex_Begin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_begin
func (s *ShaderAlebdoParts) Vertex_Begin(Vertex_Begin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Begin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_Begin sets the Vertex_Begin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_begin
func (s *ShaderAlebdoParts) SetVertex_Begin(Vertex_Begin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Begin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_Definitions returns the Vertex_Definitions property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_definitions
func (s *ShaderAlebdoParts) Vertex_Definitions(Vertex_Definitions string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Definitions)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_Definitions sets the Vertex_Definitions property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_definitions
func (s *ShaderAlebdoParts) SetVertex_Definitions(Vertex_Definitions string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_Definitions)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_MainBegin returns the Vertex_MainBegin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_mainbegin
func (s *ShaderAlebdoParts) Vertex_MainBegin(Vertex_MainBegin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_MainBegin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_MainBegin sets the Vertex_MainBegin property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_mainbegin
func (s *ShaderAlebdoParts) SetVertex_MainBegin(Vertex_MainBegin string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_MainBegin)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// Vertex_MainEnd returns the Vertex_MainEnd property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_mainend
func (s *ShaderAlebdoParts) Vertex_MainEnd(Vertex_MainEnd string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_MainEnd)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

// SetVertex_MainEnd sets the Vertex_MainEnd property of class ShaderAlebdoParts.
//
// https://doc.babylonjs.com/api/classes/babylon.shaderalebdoparts#vertex_mainend
func (s *ShaderAlebdoParts) SetVertex_MainEnd(Vertex_MainEnd string) *ShaderAlebdoParts {
	p := ba.ctx.Get("ShaderAlebdoParts").New(Vertex_MainEnd)
	return ShaderAlebdoPartsFromJSObject(p, ba.ctx)
}

*/
