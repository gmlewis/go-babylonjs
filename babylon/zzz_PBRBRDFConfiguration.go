// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRBRDFConfiguration represents a babylon.js PBRBRDFConfiguration.
// Define the code related to the BRDF parameters of the pbr material.
type PBRBRDFConfiguration struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PBRBRDFConfiguration) JSObject() js.Value { return p.p }

// PBRBRDFConfiguration returns a PBRBRDFConfiguration JavaScript class.
func (ba *Babylon) PBRBRDFConfiguration() *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration")
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// PBRBRDFConfigurationFromJSObject returns a wrapped PBRBRDFConfiguration JavaScript class.
func PBRBRDFConfigurationFromJSObject(p js.Value, ctx js.Value) *PBRBRDFConfiguration {
	return &PBRBRDFConfiguration{p: p, ctx: ctx}
}

// PBRBRDFConfigurationArrayToJSArray returns a JavaScript Array for the wrapped array.
func PBRBRDFConfigurationArrayToJSArray(array []*PBRBRDFConfiguration) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPBRBRDFConfiguration returns a new PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration
func (ba *Babylon) NewPBRBRDFConfiguration(markAllSubMeshesAsMiscDirty func()) *PBRBRDFConfiguration {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { markAllSubMeshesAsMiscDirty(); return nil }))

	p := ba.ctx.Get("PBRBRDFConfiguration").New(args...)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// CopyTo calls the CopyTo method on the PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#copyto
func (p *PBRBRDFConfiguration) CopyTo(brdfConfiguration *PBRBRDFConfiguration) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, brdfConfiguration.JSObject())

	p.p.Call("copyTo", args...)
}

// GetClassName calls the GetClassName method on the PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#getclassname
func (p *PBRBRDFConfiguration) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// Parse calls the Parse method on the PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#parse
func (p *PBRBRDFConfiguration) Parse(source interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, source)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	p.p.Call("parse", args...)
}

// PrepareDefines calls the PrepareDefines method on the PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#preparedefines
func (p *PBRBRDFConfiguration) PrepareDefines(defines js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, defines)

	p.p.Call("prepareDefines", args...)
}

// Serialize calls the Serialize method on the PBRBRDFConfiguration object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#serialize
func (p *PBRBRDFConfiguration) Serialize() interface{} {

	retVal := p.p.Call("serialize")
	return retVal
}

/*

// DEFAULT_USE_ENERGY_CONSERVATION returns the DEFAULT_USE_ENERGY_CONSERVATION property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_energy_conservation
func (p *PBRBRDFConfiguration) DEFAULT_USE_ENERGY_CONSERVATION(DEFAULT_USE_ENERGY_CONSERVATION bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_ENERGY_CONSERVATION)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetDEFAULT_USE_ENERGY_CONSERVATION sets the DEFAULT_USE_ENERGY_CONSERVATION property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_energy_conservation
func (p *PBRBRDFConfiguration) SetDEFAULT_USE_ENERGY_CONSERVATION(DEFAULT_USE_ENERGY_CONSERVATION bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_ENERGY_CONSERVATION)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED returns the DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_smith_visibility_height_correlated
func (p *PBRBRDFConfiguration) DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED(DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetDEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED sets the DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_smith_visibility_height_correlated
func (p *PBRBRDFConfiguration) SetDEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED(DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SMITH_VISIBILITY_HEIGHT_CORRELATED)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION returns the DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_specular_glossiness_input_energy_conservation
func (p *PBRBRDFConfiguration) DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION(DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetDEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION sets the DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_specular_glossiness_input_energy_conservation
func (p *PBRBRDFConfiguration) SetDEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION(DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SPECULAR_GLOSSINESS_INPUT_ENERGY_CONSERVATION)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// DEFAULT_USE_SPHERICAL_HARMONICS returns the DEFAULT_USE_SPHERICAL_HARMONICS property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_spherical_harmonics
func (p *PBRBRDFConfiguration) DEFAULT_USE_SPHERICAL_HARMONICS(DEFAULT_USE_SPHERICAL_HARMONICS bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SPHERICAL_HARMONICS)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetDEFAULT_USE_SPHERICAL_HARMONICS sets the DEFAULT_USE_SPHERICAL_HARMONICS property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#default_use_spherical_harmonics
func (p *PBRBRDFConfiguration) SetDEFAULT_USE_SPHERICAL_HARMONICS(DEFAULT_USE_SPHERICAL_HARMONICS bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(DEFAULT_USE_SPHERICAL_HARMONICS)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// UseEnergyConservation returns the UseEnergyConservation property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#useenergyconservation
func (p *PBRBRDFConfiguration) UseEnergyConservation(useEnergyConservation bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useEnergyConservation)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetUseEnergyConservation sets the UseEnergyConservation property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#useenergyconservation
func (p *PBRBRDFConfiguration) SetUseEnergyConservation(useEnergyConservation bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useEnergyConservation)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// UseSmithVisibilityHeightCorrelated returns the UseSmithVisibilityHeightCorrelated property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usesmithvisibilityheightcorrelated
func (p *PBRBRDFConfiguration) UseSmithVisibilityHeightCorrelated(useSmithVisibilityHeightCorrelated bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSmithVisibilityHeightCorrelated)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetUseSmithVisibilityHeightCorrelated sets the UseSmithVisibilityHeightCorrelated property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usesmithvisibilityheightcorrelated
func (p *PBRBRDFConfiguration) SetUseSmithVisibilityHeightCorrelated(useSmithVisibilityHeightCorrelated bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSmithVisibilityHeightCorrelated)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// UseSpecularGlossinessInputEnergyConservation returns the UseSpecularGlossinessInputEnergyConservation property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usespecularglossinessinputenergyconservation
func (p *PBRBRDFConfiguration) UseSpecularGlossinessInputEnergyConservation(useSpecularGlossinessInputEnergyConservation bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSpecularGlossinessInputEnergyConservation)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetUseSpecularGlossinessInputEnergyConservation sets the UseSpecularGlossinessInputEnergyConservation property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usespecularglossinessinputenergyconservation
func (p *PBRBRDFConfiguration) SetUseSpecularGlossinessInputEnergyConservation(useSpecularGlossinessInputEnergyConservation bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSpecularGlossinessInputEnergyConservation)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// UseSphericalHarmonics returns the UseSphericalHarmonics property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usesphericalharmonics
func (p *PBRBRDFConfiguration) UseSphericalHarmonics(useSphericalHarmonics bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSphericalHarmonics)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

// SetUseSphericalHarmonics sets the UseSphericalHarmonics property of class PBRBRDFConfiguration.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbrdfconfiguration#usesphericalharmonics
func (p *PBRBRDFConfiguration) SetUseSphericalHarmonics(useSphericalHarmonics bool) *PBRBRDFConfiguration {
	p := ba.ctx.Get("PBRBRDFConfiguration").New(useSphericalHarmonics)
	return PBRBRDFConfigurationFromJSObject(p, ba.ctx)
}

*/
