// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Path2 represents a babylon.js Path2.
// Represents a 2D path made up of multiple 2D points
type Path2 struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *Path2) JSObject() js.Value { return p.p }

// Path2 returns a Path2 JavaScript class.
func (ba *Babylon) Path2() *Path2 {
	p := ba.ctx.Get("Path2")
	return Path2FromJSObject(p, ba.ctx)
}

// Path2FromJSObject returns a wrapped Path2 JavaScript class.
func Path2FromJSObject(p js.Value, ctx js.Value) *Path2 {
	return &Path2{p: p, ctx: ctx}
}

// Path2ArrayToJSArray returns a JavaScript Array for the wrapped array.
func Path2ArrayToJSArray(array []*Path2) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPath2 returns a new Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2
func (ba *Babylon) NewPath2(x float64, y float64) *Path2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)
	args = append(args, y)

	p := ba.ctx.Get("Path2").New(args...)
	return Path2FromJSObject(p, ba.ctx)
}

// Path2AddArcToOpts contains optional parameters for Path2.AddArcTo.
type Path2AddArcToOpts struct {
	NumberOfSegments *float64
}

// AddArcTo calls the AddArcTo method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#addarcto
func (p *Path2) AddArcTo(midX float64, midY float64, endX float64, endY float64, opts *Path2AddArcToOpts) *Path2 {
	if opts == nil {
		opts = &Path2AddArcToOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, midX)
	args = append(args, midY)
	args = append(args, endX)
	args = append(args, endY)

	if opts.NumberOfSegments == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NumberOfSegments)
	}

	retVal := p.p.Call("addArcTo", args...)
	return Path2FromJSObject(retVal, p.ctx)
}

// AddLineTo calls the AddLineTo method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#addlineto
func (p *Path2) AddLineTo(x float64, y float64) *Path2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)
	args = append(args, y)

	retVal := p.p.Call("addLineTo", args...)
	return Path2FromJSObject(retVal, p.ctx)
}

// Close calls the Close method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#close
func (p *Path2) Close() *Path2 {

	retVal := p.p.Call("close")
	return Path2FromJSObject(retVal, p.ctx)
}

// GetPointAtLengthPosition calls the GetPointAtLengthPosition method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#getpointatlengthposition
func (p *Path2) GetPointAtLengthPosition(normalizedLengthPosition float64) *Vector2 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, normalizedLengthPosition)

	retVal := p.p.Call("getPointAtLengthPosition", args...)
	return Vector2FromJSObject(retVal, p.ctx)
}

// GetPoints calls the GetPoints method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#getpoints
func (p *Path2) GetPoints() []*Vector2 {

	retVal := p.p.Call("getPoints")
	result := []*Vector2{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, Vector2FromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// Length calls the Length method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#length
func (p *Path2) Length() float64 {

	retVal := p.p.Call("length")
	return retVal.Float()
}

// StartingAt calls the StartingAt method on the Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#startingat
func (p *Path2) StartingAt(x float64, y float64) *Path2 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)
	args = append(args, y)

	retVal := p.p.Call("StartingAt", args...)
	return Path2FromJSObject(retVal, p.ctx)
}

// Closed returns the Closed property of class Path2.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#closed
func (p *Path2) Closed() bool {
	retVal := p.p.Get("closed")
	return retVal.Bool()
}

// SetClosed sets the Closed property of class Path2.
//
// https://doc.babylonjs.com/api/classes/babylon.path2#closed
func (p *Path2) SetClosed(closed bool) *Path2 {
	p.p.Set("closed", closed)
	return p
}
