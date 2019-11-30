// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphericalHarmonics represents a babylon.js SphericalHarmonics.
// Class representing spherical harmonics coefficients to the 3rd degree
type SphericalHarmonics struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SphericalHarmonics) JSObject() js.Value { return s.p }

// SphericalHarmonics returns a SphericalHarmonics JavaScript class.
func (ba *Babylon) SphericalHarmonics() *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics")
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SphericalHarmonicsFromJSObject returns a wrapped SphericalHarmonics JavaScript class.
func SphericalHarmonicsFromJSObject(p js.Value, ctx js.Value) *SphericalHarmonics {
	return &SphericalHarmonics{p: p, ctx: ctx}
}

// TODO: methods
