// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRSubSurfaceConfiguration represents a babylon.js PBRSubSurfaceConfiguration.
// Define the code related to the sub surface parameters of the pbr material.
type PBRSubSurfaceConfiguration struct{}

// JSObject returns the underlying js.Value.
func (p *PBRSubSurfaceConfiguration) JSObject() js.Value { return p.p }

// PBRSubSurfaceConfiguration returns a PBRSubSurfaceConfiguration JavaScript class.
func (b *Babylon) PBRSubSurfaceConfiguration() *PBRSubSurfaceConfiguration {
	p := b.ctx.Get("PBRSubSurfaceConfiguration")
	return PBRSubSurfaceConfigurationFromJSObject(p)
}

// PBRSubSurfaceConfigurationFromJSObject returns a wrapped PBRSubSurfaceConfiguration JavaScript class.
func PBRSubSurfaceConfigurationFromJSObject(p js.Value) *PBRSubSurfaceConfiguration {
	return &PBRSubSurfaceConfiguration{p: p}
}

// NewPBRSubSurfaceConfiguration returns a new PBRSubSurfaceConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrsubsurfaceconfiguration
func (b *Babylon) NewPBRSubSurfaceConfiguration(todo parameters) *PBRSubSurfaceConfiguration {
	p := b.ctx.Get("PBRSubSurfaceConfiguration").New(todo)
	return PBRSubSurfaceConfigurationFromJSObject(p)
}

// TODO: methods
