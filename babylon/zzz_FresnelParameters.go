// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FresnelParameters represents a babylon.js FresnelParameters.
// This represents all the required information to add a fresnel effect on a material:
//
// See: http://doc.babylonjs.com/how_to/how_to_use_fresnelparameters
type FresnelParameters struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FresnelParameters) JSObject() js.Value { return f.p }

// FresnelParameters returns a FresnelParameters JavaScript class.
func (ba *Babylon) FresnelParameters() *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters")
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// FresnelParametersFromJSObject returns a wrapped FresnelParameters JavaScript class.
func FresnelParametersFromJSObject(p js.Value, ctx js.Value) *FresnelParameters {
	return &FresnelParameters{p: p, ctx: ctx}
}

// FresnelParametersArrayToJSArray returns a JavaScript Array for the wrapped array.
func FresnelParametersArrayToJSArray(array []*FresnelParameters) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Clone calls the Clone method on the FresnelParameters object.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#clone
func (f *FresnelParameters) Clone() *FresnelParameters {

	retVal := f.p.Call("clone")
	return FresnelParametersFromJSObject(retVal, f.ctx)
}

// Parse calls the Parse method on the FresnelParameters object.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#parse
func (f *FresnelParameters) Parse(parsedFresnelParameters interface{}) *FresnelParameters {

	args := make([]interface{}, 0, 1+0)

	args = append(args, parsedFresnelParameters)

	retVal := f.p.Call("Parse", args...)
	return FresnelParametersFromJSObject(retVal, f.ctx)
}

// Serialize calls the Serialize method on the FresnelParameters object.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#serialize
func (f *FresnelParameters) Serialize() interface{} {

	retVal := f.p.Call("serialize")
	return retVal
}

/*

// Bias returns the Bias property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#bias
func (f *FresnelParameters) Bias(bias float64) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(bias)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// SetBias sets the Bias property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#bias
func (f *FresnelParameters) SetBias(bias float64) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(bias)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// IsEnabled returns the IsEnabled property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#isenabled
func (f *FresnelParameters) IsEnabled(isEnabled bool) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(isEnabled)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// SetIsEnabled sets the IsEnabled property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#isenabled
func (f *FresnelParameters) SetIsEnabled(isEnabled bool) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(isEnabled)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// LeftColor returns the LeftColor property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#leftcolor
func (f *FresnelParameters) LeftColor(leftColor *Color3) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(leftColor.JSObject())
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// SetLeftColor sets the LeftColor property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#leftcolor
func (f *FresnelParameters) SetLeftColor(leftColor *Color3) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(leftColor.JSObject())
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// Power returns the Power property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#power
func (f *FresnelParameters) Power(power float64) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(power)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// SetPower sets the Power property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#power
func (f *FresnelParameters) SetPower(power float64) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(power)
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// RightColor returns the RightColor property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#rightcolor
func (f *FresnelParameters) RightColor(rightColor *Color3) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(rightColor.JSObject())
	return FresnelParametersFromJSObject(p, ba.ctx)
}

// SetRightColor sets the RightColor property of class FresnelParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelparameters#rightcolor
func (f *FresnelParameters) SetRightColor(rightColor *Color3) *FresnelParameters {
	p := ba.ctx.Get("FresnelParameters").New(rightColor.JSObject())
	return FresnelParametersFromJSObject(p, ba.ctx)
}

*/
