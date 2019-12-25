// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Vector2 represents a babylon.js Vector2.
// Class representing a vector containing 2 coordinates
type Vector2 struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *Vector2) JSObject() js.Value { return v.p }

// Vector2 returns a Vector2 JavaScript class.
func (ba *Babylon) Vector2() *Vector2 {
	p := ba.ctx.Get("Vector2")
	return Vector2FromJSObject(p, ba.ctx)
}

// Vector2FromJSObject returns a wrapped Vector2 JavaScript class.
func Vector2FromJSObject(p js.Value, ctx js.Value) *Vector2 {
	return &Vector2{p: p, ctx: ctx}
}

// Vector2ArrayToJSArray returns a JavaScript Array for the wrapped array.
func Vector2ArrayToJSArray(array []*Vector2) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVector2 returns a new Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#constructor
func (ba *Babylon) NewVector2(x float64, y float64) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)
	args = append(args, y)

	p := ba.ctx.Get("Vector2").New(args...)
	return Vector2FromJSObject(p, ba.ctx)
}

// Add calls the Add method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#add
func (v *Vector2) Add(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("add", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// AddInPlace calls the AddInPlace method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#addinplace
func (v *Vector2) AddInPlace(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("addInPlace", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// AddToRef calls the AddToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#addtoref
func (v *Vector2) AddToRef(otherVector *Vector2, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("addToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// AddVector3 calls the AddVector3 method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#addvector3
func (v *Vector2) AddVector3(otherVector *Vector3) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("addVector3", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// AsArray calls the AsArray method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#asarray
func (v *Vector2) AsArray() []float64 {

	retVal := v.p.Call("asArray")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// CatmullRom calls the CatmullRom method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#catmullrom
func (v *Vector2) CatmullRom(value1 *Vector2, value2 *Vector2, value3 *Vector2, value4 *Vector2, amount float64) *Vector2 {

	args := make([]interface{}, 0, 5+0)

	if value1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value1.JSObject())
	}

	if value2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value2.JSObject())
	}

	if value3 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value3.JSObject())
	}

	if value4 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value4.JSObject())
	}

	args = append(args, amount)

	retVal := v.p.Call("CatmullRom", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Center calls the Center method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#center
func (v *Vector2) Center(value1 *Vector2, value2 *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if value1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value1.JSObject())
	}

	if value2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value2.JSObject())
	}

	retVal := v.p.Call("Center", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Clamp calls the Clamp method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#clamp
func (v *Vector2) Clamp(value *Vector2, min *Vector2, max *Vector2) *Vector2 {

	args := make([]interface{}, 0, 3+0)

	if value == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value.JSObject())
	}

	if min == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, min.JSObject())
	}

	if max == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, max.JSObject())
	}

	retVal := v.p.Call("Clamp", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Clone calls the Clone method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#clone
func (v *Vector2) Clone() *Vector2 {

	retVal := v.p.Call("clone")
	return Vector2FromJSObject(retVal, v.ctx)
}

// CopyFrom calls the CopyFrom method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#copyfrom
func (v *Vector2) CopyFrom(source *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	retVal := v.p.Call("copyFrom", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// CopyFromFloats calls the CopyFromFloats method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#copyfromfloats
func (v *Vector2) CopyFromFloats(x float64, y float64) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, y)

	retVal := v.p.Call("copyFromFloats", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Distance calls the Distance method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#distance
func (v *Vector2) Distance(value1 *Vector2, value2 *Vector2) float64 {

	args := make([]interface{}, 0, 2+0)

	if value1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value1.JSObject())
	}

	if value2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value2.JSObject())
	}

	retVal := v.p.Call("Distance", args...)
	return retVal.Float()
}

// DistanceOfPointFromSegment calls the DistanceOfPointFromSegment method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#distanceofpointfromsegment
func (v *Vector2) DistanceOfPointFromSegment(p *Vector2, segA *Vector2, segB *Vector2) float64 {

	args := make([]interface{}, 0, 3+0)

	if p == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, p.JSObject())
	}

	if segA == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, segA.JSObject())
	}

	if segB == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, segB.JSObject())
	}

	retVal := v.p.Call("DistanceOfPointFromSegment", args...)
	return retVal.Float()
}

// DistanceSquared calls the DistanceSquared method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#distancesquared
func (v *Vector2) DistanceSquared(value1 *Vector2, value2 *Vector2) float64 {

	args := make([]interface{}, 0, 2+0)

	if value1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value1.JSObject())
	}

	if value2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value2.JSObject())
	}

	retVal := v.p.Call("DistanceSquared", args...)
	return retVal.Float()
}

// Divide calls the Divide method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#divide
func (v *Vector2) Divide(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("divide", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// DivideInPlace calls the DivideInPlace method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#divideinplace
func (v *Vector2) DivideInPlace(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("divideInPlace", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// DivideToRef calls the DivideToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#dividetoref
func (v *Vector2) DivideToRef(otherVector *Vector2, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("divideToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Dot calls the Dot method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#dot
func (v *Vector2) Dot(left *Vector2, right *Vector2) float64 {

	args := make([]interface{}, 0, 2+0)

	if left == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, left.JSObject())
	}

	if right == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, right.JSObject())
	}

	retVal := v.p.Call("Dot", args...)
	return retVal.Float()
}

// Equals calls the Equals method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#equals
func (v *Vector2) Equals(otherVector *Vector2) bool {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("equals", args...)
	return retVal.Bool()
}

// Vector2EqualsWithEpsilonOpts contains optional parameters for Vector2.EqualsWithEpsilon.
type Vector2EqualsWithEpsilonOpts struct {
	Epsilon *float64
}

// EqualsWithEpsilon calls the EqualsWithEpsilon method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#equalswithepsilon
func (v *Vector2) EqualsWithEpsilon(otherVector *Vector2, opts *Vector2EqualsWithEpsilonOpts) bool {
	if opts == nil {
		opts = &Vector2EqualsWithEpsilonOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	if opts.Epsilon == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Epsilon)
	}

	retVal := v.p.Call("equalsWithEpsilon", args...)
	return retVal.Bool()
}

// Floor calls the Floor method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#floor
func (v *Vector2) Floor() *Vector2 {

	retVal := v.p.Call("floor")
	return Vector2FromJSObject(retVal, v.ctx)
}

// Fract calls the Fract method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#fract
func (v *Vector2) Fract() *Vector2 {

	retVal := v.p.Call("fract")
	return Vector2FromJSObject(retVal, v.ctx)
}

// Vector2FromArrayOpts contains optional parameters for Vector2.FromArray.
type Vector2FromArrayOpts struct {
	Offset *float64
}

// FromArray calls the FromArray method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#fromarray
func (v *Vector2) FromArray(array js.Value, opts *Vector2FromArrayOpts) *Vector2 {
	if opts == nil {
		opts = &Vector2FromArrayOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, array)

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	retVal := v.p.Call("FromArray", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// FromArrayToRef calls the FromArrayToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#fromarraytoref
func (v *Vector2) FromArrayToRef(array js.Value, offset float64, result *Vector2) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, array)

	args = append(args, offset)

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	v.p.Call("FromArrayToRef", args...)
}

// GetClassName calls the GetClassName method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#getclassname
func (v *Vector2) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// GetHashCode calls the GetHashCode method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#gethashcode
func (v *Vector2) GetHashCode() float64 {

	retVal := v.p.Call("getHashCode")
	return retVal.Float()
}

// Hermite calls the Hermite method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#hermite
func (v *Vector2) Hermite(value1 *Vector2, tangent1 *Vector2, value2 *Vector2, tangent2 *Vector2, amount float64) *Vector2 {

	args := make([]interface{}, 0, 5+0)

	if value1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value1.JSObject())
	}

	if tangent1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, tangent1.JSObject())
	}

	if value2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value2.JSObject())
	}

	if tangent2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, tangent2.JSObject())
	}

	args = append(args, amount)

	retVal := v.p.Call("Hermite", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Length calls the Length method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#length
func (v *Vector2) Length() float64 {

	retVal := v.p.Call("length")
	return retVal.Float()
}

// LengthSquared calls the LengthSquared method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#lengthsquared
func (v *Vector2) LengthSquared() float64 {

	retVal := v.p.Call("lengthSquared")
	return retVal.Float()
}

// Lerp calls the Lerp method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#lerp
func (v *Vector2) Lerp(start *Vector2, end *Vector2, amount float64) *Vector2 {

	args := make([]interface{}, 0, 3+0)

	if start == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, start.JSObject())
	}

	if end == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, end.JSObject())
	}

	args = append(args, amount)

	retVal := v.p.Call("Lerp", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Maximize calls the Maximize method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#maximize
func (v *Vector2) Maximize(left *Vector2, right *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if left == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, left.JSObject())
	}

	if right == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, right.JSObject())
	}

	retVal := v.p.Call("Maximize", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Minimize calls the Minimize method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#minimize
func (v *Vector2) Minimize(left *Vector2, right *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if left == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, left.JSObject())
	}

	if right == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, right.JSObject())
	}

	retVal := v.p.Call("Minimize", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Multiply calls the Multiply method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#multiply
func (v *Vector2) Multiply(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("multiply", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// MultiplyByFloats calls the MultiplyByFloats method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#multiplybyfloats
func (v *Vector2) MultiplyByFloats(x float64, y float64) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, y)

	retVal := v.p.Call("multiplyByFloats", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// MultiplyInPlace calls the MultiplyInPlace method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#multiplyinplace
func (v *Vector2) MultiplyInPlace(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("multiplyInPlace", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// MultiplyToRef calls the MultiplyToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#multiplytoref
func (v *Vector2) MultiplyToRef(otherVector *Vector2, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("multiplyToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Negate calls the Negate method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#negate
func (v *Vector2) Negate() *Vector2 {

	retVal := v.p.Call("negate")
	return Vector2FromJSObject(retVal, v.ctx)
}

// Normalize calls the Normalize method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#normalize
func (v *Vector2) Normalize(vector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if vector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vector.JSObject())
	}

	retVal := v.p.Call("Normalize", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// One calls the One method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#one
func (v *Vector2) One() *Vector2 {

	retVal := v.p.Call("One")
	return Vector2FromJSObject(retVal, v.ctx)
}

// PointInTriangle calls the PointInTriangle method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#pointintriangle
func (v *Vector2) PointInTriangle(p *Vector2, p0 *Vector2, p1 *Vector2, p2 *Vector2) bool {

	args := make([]interface{}, 0, 4+0)

	if p == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, p.JSObject())
	}

	if p0 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, p0.JSObject())
	}

	if p1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, p1.JSObject())
	}

	if p2 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, p2.JSObject())
	}

	retVal := v.p.Call("PointInTriangle", args...)
	return retVal.Bool()
}

// Scale calls the Scale method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#scale
func (v *Vector2) Scale(scale float64) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scale)

	retVal := v.p.Call("scale", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// ScaleAndAddToRef calls the ScaleAndAddToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#scaleandaddtoref
func (v *Vector2) ScaleAndAddToRef(scale float64, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scale)

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("scaleAndAddToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// ScaleInPlace calls the ScaleInPlace method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#scaleinplace
func (v *Vector2) ScaleInPlace(scale float64) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scale)

	retVal := v.p.Call("scaleInPlace", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// ScaleToRef calls the ScaleToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#scaletoref
func (v *Vector2) ScaleToRef(scale float64, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scale)

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("scaleToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Set calls the Set method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#set
func (v *Vector2) Set(x float64, y float64) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, y)

	retVal := v.p.Call("set", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Subtract calls the Subtract method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#subtract
func (v *Vector2) Subtract(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("subtract", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// SubtractInPlace calls the SubtractInPlace method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#subtractinplace
func (v *Vector2) SubtractInPlace(otherVector *Vector2) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	retVal := v.p.Call("subtractInPlace", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// SubtractToRef calls the SubtractToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#subtracttoref
func (v *Vector2) SubtractToRef(otherVector *Vector2, result *Vector2) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if otherVector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, otherVector.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	retVal := v.p.Call("subtractToRef", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// Vector2ToArrayOpts contains optional parameters for Vector2.ToArray.
type Vector2ToArrayOpts struct {
	Index *float64
}

// ToArray calls the ToArray method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#toarray
func (v *Vector2) ToArray(array js.Value, opts *Vector2ToArrayOpts) *Vector2 {
	if opts == nil {
		opts = &Vector2ToArrayOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, array)

	if opts.Index == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Index)
	}

	retVal := v.p.Call("toArray", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// ToString calls the ToString method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#tostring
func (v *Vector2) ToString() string {

	retVal := v.p.Call("toString")
	return retVal.String()
}

// Transform calls the Transform method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#transform
func (v *Vector2) Transform(vector *Vector2, transformation *Matrix) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if vector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vector.JSObject())
	}

	if transformation == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transformation.JSObject())
	}

	retVal := v.p.Call("Transform", args...)
	return Vector2FromJSObject(retVal, v.ctx)
}

// TransformToRef calls the TransformToRef method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#transformtoref
func (v *Vector2) TransformToRef(vector *Vector2, transformation *Matrix, result *Vector2) {

	args := make([]interface{}, 0, 3+0)

	if vector == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, vector.JSObject())
	}

	if transformation == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transformation.JSObject())
	}

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	v.p.Call("TransformToRef", args...)
}

// Zero calls the Zero method on the Vector2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#zero
func (v *Vector2) Zero() *Vector2 {

	retVal := v.p.Call("Zero")
	return Vector2FromJSObject(retVal, v.ctx)
}

// X returns the X property of class Vector2.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#x
func (v *Vector2) X() float64 {
	retVal := v.p.Get("x")
	return retVal.Float()
}

// SetX sets the X property of class Vector2.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#x
func (v *Vector2) SetX(x float64) *Vector2 {
	v.p.Set("x", x)
	return v
}

// Y returns the Y property of class Vector2.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#y
func (v *Vector2) Y() float64 {
	retVal := v.p.Get("y")
	return retVal.Float()
}

// SetY sets the Y property of class Vector2.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2#y
func (v *Vector2) SetY(y float64) *Vector2 {
	v.p.Set("y", y)
	return v
}
