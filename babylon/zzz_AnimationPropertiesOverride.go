// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnimationPropertiesOverride represents a babylon.js AnimationPropertiesOverride.
// Class used to override all child animations of a given target
type AnimationPropertiesOverride struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnimationPropertiesOverride) JSObject() js.Value { return a.p }

// AnimationPropertiesOverride returns a AnimationPropertiesOverride JavaScript class.
func (ba *Babylon) AnimationPropertiesOverride() *AnimationPropertiesOverride {
	p := ba.ctx.Get("AnimationPropertiesOverride")
	return AnimationPropertiesOverrideFromJSObject(p, ba.ctx)
}

// AnimationPropertiesOverrideFromJSObject returns a wrapped AnimationPropertiesOverride JavaScript class.
func AnimationPropertiesOverrideFromJSObject(p js.Value, ctx js.Value) *AnimationPropertiesOverride {
	return &AnimationPropertiesOverride{p: p, ctx: ctx}
}

// TODO: methods
