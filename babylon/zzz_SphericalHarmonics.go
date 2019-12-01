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

// SphericalHarmonicsArrayToJSArray returns a JavaScript Array for the wrapped array.
func SphericalHarmonicsArrayToJSArray(array []*SphericalHarmonics) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddLight calls the AddLight method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#addlight
func (s *SphericalHarmonics) AddLight(direction *Vector3, color *Color3, deltaSolidAngle float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, direction.JSObject())
	args = append(args, color.JSObject())
	args = append(args, deltaSolidAngle)

	s.p.Call("addLight", args...)
}

// ConvertIncidentRadianceToIrradiance calls the ConvertIncidentRadianceToIrradiance method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#convertincidentradiancetoirradiance
func (s *SphericalHarmonics) ConvertIncidentRadianceToIrradiance() {

	s.p.Call("convertIncidentRadianceToIrradiance")
}

// ConvertIrradianceToLambertianRadiance calls the ConvertIrradianceToLambertianRadiance method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#convertirradiancetolambertianradiance
func (s *SphericalHarmonics) ConvertIrradianceToLambertianRadiance() {

	s.p.Call("convertIrradianceToLambertianRadiance")
}

// FromArray calls the FromArray method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#fromarray
func (s *SphericalHarmonics) FromArray(data js.Value) *SphericalHarmonics {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := s.p.Call("FromArray", args...)
	return SphericalHarmonicsFromJSObject(retVal, s.ctx)
}

// FromPolynomial calls the FromPolynomial method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#frompolynomial
func (s *SphericalHarmonics) FromPolynomial(polynomial *SphericalPolynomial) *SphericalHarmonics {

	args := make([]interface{}, 0, 1+0)

	args = append(args, polynomial.JSObject())

	retVal := s.p.Call("FromPolynomial", args...)
	return SphericalHarmonicsFromJSObject(retVal, s.ctx)
}

// PreScaleForRendering calls the PreScaleForRendering method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#prescaleforrendering
func (s *SphericalHarmonics) PreScaleForRendering() {

	s.p.Call("preScaleForRendering")
}

// ScaleInPlace calls the ScaleInPlace method on the SphericalHarmonics object.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#scaleinplace
func (s *SphericalHarmonics) ScaleInPlace(scale float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scale)

	s.p.Call("scaleInPlace", args...)
}

/*

// L00 returns the L00 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l00
func (s *SphericalHarmonics) L00(l00 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l00.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL00 sets the L00 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l00
func (s *SphericalHarmonics) SetL00(l00 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l00.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L10 returns the L10 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l10
func (s *SphericalHarmonics) L10(l10 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l10.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL10 sets the L10 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l10
func (s *SphericalHarmonics) SetL10(l10 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l10.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L11 returns the L11 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l11
func (s *SphericalHarmonics) L11(l11 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l11.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL11 sets the L11 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l11
func (s *SphericalHarmonics) SetL11(l11 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l11.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L1_1 returns the L1_1 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l1_1
func (s *SphericalHarmonics) L1_1(l1_1 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l1_1.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL1_1 sets the L1_1 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l1_1
func (s *SphericalHarmonics) SetL1_1(l1_1 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l1_1.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L20 returns the L20 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l20
func (s *SphericalHarmonics) L20(l20 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l20.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL20 sets the L20 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l20
func (s *SphericalHarmonics) SetL20(l20 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l20.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L21 returns the L21 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l21
func (s *SphericalHarmonics) L21(l21 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l21.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL21 sets the L21 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l21
func (s *SphericalHarmonics) SetL21(l21 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l21.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L22 returns the L22 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l22
func (s *SphericalHarmonics) L22(l22 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l22.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL22 sets the L22 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l22
func (s *SphericalHarmonics) SetL22(l22 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l22.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L2_1 returns the L2_1 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l2_1
func (s *SphericalHarmonics) L2_1(l2_1 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l2_1.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL2_1 sets the L2_1 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l2_1
func (s *SphericalHarmonics) SetL2_1(l2_1 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l2_1.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// L2_2 returns the L2_2 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l2_2
func (s *SphericalHarmonics) L2_2(l2_2 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l2_2.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetL2_2 sets the L2_2 property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#l2_2
func (s *SphericalHarmonics) SetL2_2(l2_2 *Vector3) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(l2_2.JSObject())
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// PreScaled returns the PreScaled property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#prescaled
func (s *SphericalHarmonics) PreScaled(preScaled bool) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(preScaled)
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

// SetPreScaled sets the PreScaled property of class SphericalHarmonics.
//
// https://doc.babylonjs.com/api/classes/babylon.sphericalharmonics#prescaled
func (s *SphericalHarmonics) SetPreScaled(preScaled bool) *SphericalHarmonics {
	p := ba.ctx.Get("SphericalHarmonics").New(preScaled)
	return SphericalHarmonicsFromJSObject(p, ba.ctx)
}

*/
