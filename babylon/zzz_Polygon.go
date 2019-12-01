// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Polygon represents a babylon.js Polygon.
// Polygon
//
// See: https://doc.babylonjs.com/how_to/parametric_shapes#non-regular-polygon
type Polygon struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *Polygon) JSObject() js.Value { return p.p }

// Polygon returns a Polygon JavaScript class.
func (ba *Babylon) Polygon() *Polygon {
	p := ba.ctx.Get("Polygon")
	return PolygonFromJSObject(p, ba.ctx)
}

// PolygonFromJSObject returns a wrapped Polygon JavaScript class.
func PolygonFromJSObject(p js.Value, ctx js.Value) *Polygon {
	return &Polygon{p: p, ctx: ctx}
}

// PolygonCircleOpts contains optional parameters for Polygon.Circle.
type PolygonCircleOpts struct {
	Cx            *float64
	Cy            *float64
	NumberOfSides *float64
}

// Circle calls the Circle method on the Polygon object.
//
// https://doc.babylonjs.com/api/classes/babylon.polygon#circle
func (p *Polygon) Circle(radius float64, opts *PolygonCircleOpts) *Vector2 {
	if opts == nil {
		opts = &PolygonCircleOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, radius)

	if opts.Cx == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Cx)
	}
	if opts.Cy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Cy)
	}
	if opts.NumberOfSides == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NumberOfSides)
	}

	retVal := p.p.Call("Circle", args...)
	return Vector2FromJSObject(retVal, p.ctx)
}

// Parse calls the Parse method on the Polygon object.
//
// https://doc.babylonjs.com/api/classes/babylon.polygon#parse
func (p *Polygon) Parse(input string) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, input)

	retVal := p.p.Call("Parse", args...)
	return Vector2FromJSObject(retVal, p.ctx)
}

// Rectangle calls the Rectangle method on the Polygon object.
//
// https://doc.babylonjs.com/api/classes/babylon.polygon#rectangle
func (p *Polygon) Rectangle(xmin float64, ymin float64, xmax float64, ymax float64) *Vector2 {

	args := make([]interface{}, 0, 4+0)

	args = append(args, xmin)
	args = append(args, ymin)
	args = append(args, xmax)
	args = append(args, ymax)

	retVal := p.p.Call("Rectangle", args...)
	return Vector2FromJSObject(retVal, p.ctx)
}

// StartingAt calls the StartingAt method on the Polygon object.
//
// https://doc.babylonjs.com/api/classes/babylon.polygon#startingat
func (p *Polygon) StartingAt(x float64, y float64) *Path2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)
	args = append(args, y)

	retVal := p.p.Call("StartingAt", args...)
	return Path2FromJSObject(retVal, p.ctx)
}

/*

 */
