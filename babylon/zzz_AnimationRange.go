// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnimationRange represents a babylon.js AnimationRange.
// Represents the range of an animation
type AnimationRange struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnimationRange) JSObject() js.Value { return a.p }

// AnimationRange returns a AnimationRange JavaScript class.
func (ba *Babylon) AnimationRange() *AnimationRange {
	p := ba.ctx.Get("AnimationRange")
	return AnimationRangeFromJSObject(p, ba.ctx)
}

// AnimationRangeFromJSObject returns a wrapped AnimationRange JavaScript class.
func AnimationRangeFromJSObject(p js.Value, ctx js.Value) *AnimationRange {
	return &AnimationRange{p: p, ctx: ctx}
}

// AnimationRangeArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnimationRangeArrayToJSArray(array []*AnimationRange) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnimationRange returns a new AnimationRange object.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#constructor
func (ba *Babylon) NewAnimationRange(name string, from float64, to float64) *AnimationRange {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, from)
	args = append(args, to)

	p := ba.ctx.Get("AnimationRange").New(args...)
	return AnimationRangeFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the AnimationRange object.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#clone
func (a *AnimationRange) Clone() *AnimationRange {

	retVal := a.p.Call("clone")
	return AnimationRangeFromJSObject(retVal, a.ctx)
}

// From returns the From property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#from
func (a *AnimationRange) From() float64 {
	retVal := a.p.Get("from")
	return retVal.Float()
}

// SetFrom sets the From property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#from
func (a *AnimationRange) SetFrom(from float64) *AnimationRange {
	a.p.Set("from", from)
	return a
}

// Name returns the Name property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#name
func (a *AnimationRange) Name() string {
	retVal := a.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#name
func (a *AnimationRange) SetName(name string) *AnimationRange {
	a.p.Set("name", name)
	return a
}

// To returns the To property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#to
func (a *AnimationRange) To() float64 {
	retVal := a.p.Get("to")
	return retVal.Float()
}

// SetTo sets the To property of class AnimationRange.
//
// https://doc.babylonjs.com/api/classes/babylon.animationrange#to
func (a *AnimationRange) SetTo(to float64) *AnimationRange {
	a.p.Set("to", to)
	return a
}
