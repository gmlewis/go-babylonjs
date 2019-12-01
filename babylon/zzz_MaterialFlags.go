// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MaterialFlags represents a babylon.js MaterialFlags.
// This groups all the flags used to control the materials channel.
type MaterialFlags struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MaterialFlags) JSObject() js.Value { return m.p }

// MaterialFlags returns a MaterialFlags JavaScript class.
func (ba *Babylon) MaterialFlags() *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags")
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// MaterialFlagsFromJSObject returns a wrapped MaterialFlags JavaScript class.
func MaterialFlagsFromJSObject(p js.Value, ctx js.Value) *MaterialFlags {
	return &MaterialFlags{p: p, ctx: ctx}
}

/*

// AmbientTextureEnabled returns the AmbientTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#ambienttextureenabled
func (m *MaterialFlags) AmbientTextureEnabled(AmbientTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(AmbientTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetAmbientTextureEnabled sets the AmbientTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#ambienttextureenabled
func (m *MaterialFlags) SetAmbientTextureEnabled(AmbientTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(AmbientTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// AnisotropicTextureEnabled returns the AnisotropicTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#anisotropictextureenabled
func (m *MaterialFlags) AnisotropicTextureEnabled(AnisotropicTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(AnisotropicTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetAnisotropicTextureEnabled sets the AnisotropicTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#anisotropictextureenabled
func (m *MaterialFlags) SetAnisotropicTextureEnabled(AnisotropicTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(AnisotropicTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// BumpTextureEnabled returns the BumpTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#bumptextureenabled
func (m *MaterialFlags) BumpTextureEnabled(BumpTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(BumpTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetBumpTextureEnabled sets the BumpTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#bumptextureenabled
func (m *MaterialFlags) SetBumpTextureEnabled(BumpTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(BumpTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ClearCoatBumpTextureEnabled returns the ClearCoatBumpTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoatbumptextureenabled
func (m *MaterialFlags) ClearCoatBumpTextureEnabled(ClearCoatBumpTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatBumpTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetClearCoatBumpTextureEnabled sets the ClearCoatBumpTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoatbumptextureenabled
func (m *MaterialFlags) SetClearCoatBumpTextureEnabled(ClearCoatBumpTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatBumpTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ClearCoatTextureEnabled returns the ClearCoatTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoattextureenabled
func (m *MaterialFlags) ClearCoatTextureEnabled(ClearCoatTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetClearCoatTextureEnabled sets the ClearCoatTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoattextureenabled
func (m *MaterialFlags) SetClearCoatTextureEnabled(ClearCoatTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ClearCoatTintTextureEnabled returns the ClearCoatTintTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoattinttextureenabled
func (m *MaterialFlags) ClearCoatTintTextureEnabled(ClearCoatTintTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatTintTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetClearCoatTintTextureEnabled sets the ClearCoatTintTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#clearcoattinttextureenabled
func (m *MaterialFlags) SetClearCoatTintTextureEnabled(ClearCoatTintTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ClearCoatTintTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ColorGradingTextureEnabled returns the ColorGradingTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#colorgradingtextureenabled
func (m *MaterialFlags) ColorGradingTextureEnabled(ColorGradingTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ColorGradingTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetColorGradingTextureEnabled sets the ColorGradingTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#colorgradingtextureenabled
func (m *MaterialFlags) SetColorGradingTextureEnabled(ColorGradingTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ColorGradingTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// DiffuseTextureEnabled returns the DiffuseTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#diffusetextureenabled
func (m *MaterialFlags) DiffuseTextureEnabled(DiffuseTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(DiffuseTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetDiffuseTextureEnabled sets the DiffuseTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#diffusetextureenabled
func (m *MaterialFlags) SetDiffuseTextureEnabled(DiffuseTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(DiffuseTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// EmissiveTextureEnabled returns the EmissiveTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#emissivetextureenabled
func (m *MaterialFlags) EmissiveTextureEnabled(EmissiveTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(EmissiveTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetEmissiveTextureEnabled sets the EmissiveTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#emissivetextureenabled
func (m *MaterialFlags) SetEmissiveTextureEnabled(EmissiveTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(EmissiveTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// FresnelEnabled returns the FresnelEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#fresnelenabled
func (m *MaterialFlags) FresnelEnabled(FresnelEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(FresnelEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetFresnelEnabled sets the FresnelEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#fresnelenabled
func (m *MaterialFlags) SetFresnelEnabled(FresnelEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(FresnelEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// LightmapTextureEnabled returns the LightmapTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#lightmaptextureenabled
func (m *MaterialFlags) LightmapTextureEnabled(LightmapTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(LightmapTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetLightmapTextureEnabled sets the LightmapTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#lightmaptextureenabled
func (m *MaterialFlags) SetLightmapTextureEnabled(LightmapTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(LightmapTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// OpacityTextureEnabled returns the OpacityTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#opacitytextureenabled
func (m *MaterialFlags) OpacityTextureEnabled(OpacityTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(OpacityTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetOpacityTextureEnabled sets the OpacityTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#opacitytextureenabled
func (m *MaterialFlags) SetOpacityTextureEnabled(OpacityTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(OpacityTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ReflectionTextureEnabled returns the ReflectionTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#reflectiontextureenabled
func (m *MaterialFlags) ReflectionTextureEnabled(ReflectionTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ReflectionTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetReflectionTextureEnabled sets the ReflectionTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#reflectiontextureenabled
func (m *MaterialFlags) SetReflectionTextureEnabled(ReflectionTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ReflectionTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// RefractionTextureEnabled returns the RefractionTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#refractiontextureenabled
func (m *MaterialFlags) RefractionTextureEnabled(RefractionTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(RefractionTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetRefractionTextureEnabled sets the RefractionTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#refractiontextureenabled
func (m *MaterialFlags) SetRefractionTextureEnabled(RefractionTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(RefractionTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SheenTextureEnabled returns the SheenTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#sheentextureenabled
func (m *MaterialFlags) SheenTextureEnabled(SheenTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(SheenTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetSheenTextureEnabled sets the SheenTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#sheentextureenabled
func (m *MaterialFlags) SetSheenTextureEnabled(SheenTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(SheenTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SpecularTextureEnabled returns the SpecularTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#speculartextureenabled
func (m *MaterialFlags) SpecularTextureEnabled(SpecularTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(SpecularTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetSpecularTextureEnabled sets the SpecularTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#speculartextureenabled
func (m *MaterialFlags) SetSpecularTextureEnabled(SpecularTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(SpecularTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// ThicknessTextureEnabled returns the ThicknessTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#thicknesstextureenabled
func (m *MaterialFlags) ThicknessTextureEnabled(ThicknessTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ThicknessTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

// SetThicknessTextureEnabled sets the ThicknessTextureEnabled property of class MaterialFlags.
//
// https://doc.babylonjs.com/api/classes/babylon.materialflags#thicknesstextureenabled
func (m *MaterialFlags) SetThicknessTextureEnabled(ThicknessTextureEnabled bool) *MaterialFlags {
	p := ba.ctx.Get("MaterialFlags").New(ThicknessTextureEnabled)
	return MaterialFlagsFromJSObject(p, ba.ctx)
}

*/
