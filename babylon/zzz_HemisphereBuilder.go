// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphereBuilder represents a babylon.js HemisphereBuilder.
// Class containing static functions to help procedurally build meshes
type HemisphereBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HemisphereBuilder) JSObject() js.Value { return h.p }

// HemisphereBuilder returns a HemisphereBuilder JavaScript class.
func (ba *Babylon) HemisphereBuilder() *HemisphereBuilder {
	p := ba.ctx.Get("HemisphereBuilder")
	return HemisphereBuilderFromJSObject(p, ba.ctx)
}

// HemisphereBuilderFromJSObject returns a wrapped HemisphereBuilder JavaScript class.
func HemisphereBuilderFromJSObject(p js.Value, ctx js.Value) *HemisphereBuilder {
	return &HemisphereBuilder{p: p, ctx: ctx}
}

// CreateHemisphere calls the CreateHemisphere method on the HemisphereBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemispherebuilder#createhemisphere
func (h *HemisphereBuilder) CreateHemisphere(name string, options js.Value, scene interface{}) *Mesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, options)
	args = append(args, scene)

	retVal := h.p.Call("CreateHemisphere", args...)
	return MeshFromJSObject(retVal, h.ctx)
}

/*

 */
