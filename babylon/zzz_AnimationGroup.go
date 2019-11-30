// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnimationGroup represents a babylon.js AnimationGroup.
// Use this class to create coordinated animations on multiple targets
type AnimationGroup struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnimationGroup) JSObject() js.Value { return a.p }

// AnimationGroup returns a AnimationGroup JavaScript class.
func (ba *Babylon) AnimationGroup() *AnimationGroup {
	p := ba.ctx.Get("AnimationGroup")
	return AnimationGroupFromJSObject(p, ba.ctx)
}

// AnimationGroupFromJSObject returns a wrapped AnimationGroup JavaScript class.
func AnimationGroupFromJSObject(p js.Value, ctx js.Value) *AnimationGroup {
	return &AnimationGroup{p: p, ctx: ctx}
}

// NewAnimationGroupOpts contains optional parameters for NewAnimationGroup.
type NewAnimationGroupOpts struct {
	Scene *Scene
}

// NewAnimationGroup returns a new AnimationGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.animationgroup
func (ba *Babylon) NewAnimationGroup(name string, opts *NewAnimationGroupOpts) *AnimationGroup {
	if opts == nil {
		opts = &NewAnimationGroupOpts{}
	}

	p := ba.ctx.Get("AnimationGroup").New(name, opts.Scene.JSObject())
	return AnimationGroupFromJSObject(p, ba.ctx)
}

// TODO: methods
