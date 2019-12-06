// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Curve3 represents a babylon.js Curve3.
// A Curve3 object is a logical object, so not a mesh, to handle curves in the 3D geometric space.
// A Curve3 is designed from a series of successive Vector3.
//
// See: https://doc.babylonjs.com/how_to/how_to_use_curve3
type Curve3 struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Curve3) JSObject() js.Value { return c.p }

// Curve3 returns a Curve3 JavaScript class.
func (ba *Babylon) Curve3() *Curve3 {
	p := ba.ctx.Get("Curve3")
	return Curve3FromJSObject(p, ba.ctx)
}

// Curve3FromJSObject returns a wrapped Curve3 JavaScript class.
func Curve3FromJSObject(p js.Value, ctx js.Value) *Curve3 {
	return &Curve3{p: p, ctx: ctx}
}

// Curve3ArrayToJSArray returns a JavaScript Array for the wrapped array.
func Curve3ArrayToJSArray(array []*Curve3) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCurve3 returns a new Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3
func (ba *Babylon) NewCurve3(points []*Vector3) *Curve3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, points)

	p := ba.ctx.Get("Curve3").New(args...)
	return Curve3FromJSObject(p, ba.ctx)
}

// Continue calls the Continue method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#continue
func (c *Curve3) Continue(curve *Curve3) *Curve3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, curve.JSObject())

	retVal := c.p.Call("continue", args...)
	return Curve3FromJSObject(retVal, c.ctx)
}

// Curve3CreateCatmullRomSplineOpts contains optional parameters for Curve3.CreateCatmullRomSpline.
type Curve3CreateCatmullRomSplineOpts struct {
	Closed *bool
}

// CreateCatmullRomSpline calls the CreateCatmullRomSpline method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#createcatmullromspline
func (c *Curve3) CreateCatmullRomSpline(points []*Vector3, nbPoints float64, opts *Curve3CreateCatmullRomSplineOpts) *Curve3 {
	if opts == nil {
		opts = &Curve3CreateCatmullRomSplineOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, Vector3ArrayToJSArray(points))
	args = append(args, nbPoints)

	if opts.Closed == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Closed)
	}

	retVal := c.p.Call("CreateCatmullRomSpline", args...)
	return Curve3FromJSObject(retVal, c.ctx)
}

// CreateCubicBezier calls the CreateCubicBezier method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#createcubicbezier
func (c *Curve3) CreateCubicBezier(v0 *Vector3, v1 *Vector3, v2 *Vector3, v3 *Vector3, nbPoints float64) *Curve3 {

	args := make([]interface{}, 0, 5+0)

	args = append(args, v0.JSObject())
	args = append(args, v1.JSObject())
	args = append(args, v2.JSObject())
	args = append(args, v3.JSObject())
	args = append(args, nbPoints)

	retVal := c.p.Call("CreateCubicBezier", args...)
	return Curve3FromJSObject(retVal, c.ctx)
}

// CreateHermiteSpline calls the CreateHermiteSpline method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#createhermitespline
func (c *Curve3) CreateHermiteSpline(p1 *Vector3, t1 *Vector3, p2 *Vector3, t2 *Vector3, nbPoints float64) *Curve3 {

	args := make([]interface{}, 0, 5+0)

	args = append(args, p1.JSObject())
	args = append(args, t1.JSObject())
	args = append(args, p2.JSObject())
	args = append(args, t2.JSObject())
	args = append(args, nbPoints)

	retVal := c.p.Call("CreateHermiteSpline", args...)
	return Curve3FromJSObject(retVal, c.ctx)
}

// CreateQuadraticBezier calls the CreateQuadraticBezier method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#createquadraticbezier
func (c *Curve3) CreateQuadraticBezier(v0 *Vector3, v1 *Vector3, v2 *Vector3, nbPoints float64) *Curve3 {

	args := make([]interface{}, 0, 4+0)

	args = append(args, v0.JSObject())
	args = append(args, v1.JSObject())
	args = append(args, v2.JSObject())
	args = append(args, nbPoints)

	retVal := c.p.Call("CreateQuadraticBezier", args...)
	return Curve3FromJSObject(retVal, c.ctx)
}

// GetPoints calls the GetPoints method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#getpoints
func (c *Curve3) GetPoints() []*Vector3 {

	retVal := c.p.Call("getPoints")
	result := []*Vector3{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, Vector3FromJSObject(retVal.Index(ri), c.ctx))
	}
	return result
}

// Length calls the Length method on the Curve3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.curve3#length
func (c *Curve3) Length() float64 {

	retVal := c.p.Call("length")
	return retVal.Float()
}
