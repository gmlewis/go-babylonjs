// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RecastJSCrowd represents a babylon.js RecastJSCrowd.
// Recast detour crowd implementation
type RecastJSCrowd struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RecastJSCrowd) JSObject() js.Value { return r.p }

// RecastJSCrowd returns a RecastJSCrowd JavaScript class.
func (ba *Babylon) RecastJSCrowd() *RecastJSCrowd {
	p := ba.ctx.Get("RecastJSCrowd")
	return RecastJSCrowdFromJSObject(p, ba.ctx)
}

// RecastJSCrowdFromJSObject returns a wrapped RecastJSCrowd JavaScript class.
func RecastJSCrowdFromJSObject(p js.Value, ctx js.Value) *RecastJSCrowd {
	return &RecastJSCrowd{p: p, ctx: ctx}
}

// NewRecastJSCrowd returns a new RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd
func (ba *Babylon) NewRecastJSCrowd(plugin *RecastJSPlugin, maxAgents float64, maxAgentRadius float64, scene *Scene) *RecastJSCrowd {
	p := ba.ctx.Get("RecastJSCrowd").New(plugin.JSObject(), maxAgents, maxAgentRadius, scene.JSObject())
	return RecastJSCrowdFromJSObject(p, ba.ctx)
}

// TODO: methods
