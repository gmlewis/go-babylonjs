// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TargetedAnimation represents a babylon.js TargetedAnimation.
// This class defines the direct association between an animation and a target
type TargetedAnimation struct{}

// JSObject returns the underlying js.Value.
func (t *TargetedAnimation) JSObject() js.Value { return t.p }

// TargetedAnimation returns a TargetedAnimation JavaScript class.
func (b *Babylon) TargetedAnimation() *TargetedAnimation {
	p := b.ctx.Get("TargetedAnimation")
	return TargetedAnimationFromJSObject(p)
}

// TargetedAnimationFromJSObject returns a wrapped TargetedAnimation JavaScript class.
func TargetedAnimationFromJSObject(p js.Value) *TargetedAnimation {
	return &TargetedAnimation{p: p}
}

// NewTargetedAnimation returns a new TargetedAnimation object.
//
// https://doc.babylonjs.com/api/classes/babylon.targetedanimation
func (b *Babylon) NewTargetedAnimation(todo parameters) *TargetedAnimation {
	p := b.ctx.Get("TargetedAnimation").New(todo)
	return TargetedAnimationFromJSObject(p)
}

// TODO: methods