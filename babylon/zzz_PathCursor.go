// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PathCursor represents a babylon.js PathCursor.
// A cursor which tracks a point on a path
type PathCursor struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PathCursor) JSObject() js.Value { return p.p }

// PathCursor returns a PathCursor JavaScript class.
func (ba *Babylon) PathCursor() *PathCursor {
	p := ba.ctx.Get("PathCursor")
	return PathCursorFromJSObject(p, ba.ctx)
}

// PathCursorFromJSObject returns a wrapped PathCursor JavaScript class.
func PathCursorFromJSObject(p js.Value, ctx js.Value) *PathCursor {
	return &PathCursor{p: p, ctx: ctx}
}

// PathCursorArrayToJSArray returns a JavaScript Array for the wrapped array.
func PathCursorArrayToJSArray(array []*PathCursor) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPathCursor returns a new PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor
func (ba *Babylon) NewPathCursor(path *Path2) *PathCursor {

	args := make([]interface{}, 0, 1+0)

	args = append(args, path.JSObject())

	p := ba.ctx.Get("PathCursor").New(args...)
	return PathCursorFromJSObject(p, ba.ctx)
}

// GetPoint calls the GetPoint method on the PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#getpoint
func (p *PathCursor) GetPoint() *Vector3 {

	retVal := p.p.Call("getPoint")
	return Vector3FromJSObject(retVal, p.ctx)
}

// Move calls the Move method on the PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#move
func (p *PathCursor) Move(step float64) *PathCursor {

	args := make([]interface{}, 0, 1+0)

	args = append(args, step)

	retVal := p.p.Call("move", args...)
	return PathCursorFromJSObject(retVal, p.ctx)
}

// PathCursorMoveAheadOpts contains optional parameters for PathCursor.MoveAhead.
type PathCursorMoveAheadOpts struct {
	Step *float64
}

// MoveAhead calls the MoveAhead method on the PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#moveahead
func (p *PathCursor) MoveAhead(opts *PathCursorMoveAheadOpts) *PathCursor {
	if opts == nil {
		opts = &PathCursorMoveAheadOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Step == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Step)
	}

	retVal := p.p.Call("moveAhead", args...)
	return PathCursorFromJSObject(retVal, p.ctx)
}

// PathCursorMoveBackOpts contains optional parameters for PathCursor.MoveBack.
type PathCursorMoveBackOpts struct {
	Step *float64
}

// MoveBack calls the MoveBack method on the PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#moveback
func (p *PathCursor) MoveBack(opts *PathCursorMoveBackOpts) *PathCursor {
	if opts == nil {
		opts = &PathCursorMoveBackOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Step == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Step)
	}

	retVal := p.p.Call("moveBack", args...)
	return PathCursorFromJSObject(retVal, p.ctx)
}

// Onchange calls the Onchange method on the PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#onchange
func (p *PathCursor) Onchange(f func()) *PathCursor {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { f(); return nil }))

	retVal := p.p.Call("onchange", args...)
	return PathCursorFromJSObject(retVal, p.ctx)
}

// Animations returns the Animations property of class PathCursor.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#animations
func (p *PathCursor) Animations() []*Animation {
	retVal := p.p.Get("animations")
	result := []*Animation{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AnimationFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetAnimations sets the Animations property of class PathCursor.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#animations
func (p *PathCursor) SetAnimations(animations []*Animation) *PathCursor {
	p.p.Set("animations", animations)
	return p
}

// Value returns the Value property of class PathCursor.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#value
func (p *PathCursor) Value() float64 {
	retVal := p.p.Get("value")
	return retVal.Float()
}

// SetValue sets the Value property of class PathCursor.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor#value
func (p *PathCursor) SetValue(value float64) *PathCursor {
	p.p.Set("value", value)
	return p
}
