// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IAnimatable represents a babylon.js IAnimatable.
// Interface containing an array of animations
type IAnimatable struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IAnimatable) JSObject() js.Value { return i.p }

// IAnimatable returns a IAnimatable JavaScript class.
func (ba *Babylon) IAnimatable() *IAnimatable {
	p := ba.ctx.Get("IAnimatable")
	return IAnimatableFromJSObject(p, ba.ctx)
}

// IAnimatableFromJSObject returns a wrapped IAnimatable JavaScript class.
func IAnimatableFromJSObject(p js.Value, ctx js.Value) *IAnimatable {
	return &IAnimatable{p: p, ctx: ctx}
}

// IAnimatableArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAnimatableArrayToJSArray(array []*IAnimatable) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Animations returns the Animations property of class IAnimatable.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimatable#animations
func (i *IAnimatable) Animations() []*Animation {
	retVal := i.p.Get("animations")
	result := []*Animation{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AnimationFromJSObject(retVal.Index(ri), i.ctx))
	}
	return result
}

// SetAnimations sets the Animations property of class IAnimatable.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimatable#animations
func (i *IAnimatable) SetAnimations(animations []*Animation) *IAnimatable {
	i.p.Set("animations", animations)
	return i
}
