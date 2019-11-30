// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRSubSurfaceConfiguration represents a babylon.js PBRSubSurfaceConfiguration.
// Define the code related to the sub surface parameters of the pbr material.
type PBRSubSurfaceConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRSubSurfaceConfiguration) JSObject() js.Value { return p.p }

// PBRSubSurfaceConfiguration returns a PBRSubSurfaceConfiguration JavaScript class.
func (ba *Babylon) PBRSubSurfaceConfiguration() *PBRSubSurfaceConfiguration {
	p := ba.ctx.Get("PBRSubSurfaceConfiguration")
	return PBRSubSurfaceConfigurationFromJSObject(p, ba.ctx)
}

// PBRSubSurfaceConfigurationFromJSObject returns a wrapped PBRSubSurfaceConfiguration JavaScript class.
func PBRSubSurfaceConfigurationFromJSObject(p js.Value, ctx js.Value) *PBRSubSurfaceConfiguration {
	return &PBRSubSurfaceConfiguration{p: p, ctx: ctx}
}

// NewPBRSubSurfaceConfiguration returns a new PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration
func (ba *Babylon) NewPBRSubSurfaceConfiguration(markAllSubMeshesAsTexturesDirty func()) *PBRSubSurfaceConfiguration {
	p := ba.ctx.Get("PBRSubSurfaceConfiguration").New(markAllSubMeshesAsTexturesDirty)
	return PBRSubSurfaceConfigurationFromJSObject(p, ba.ctx)
}

// TODO: methods
