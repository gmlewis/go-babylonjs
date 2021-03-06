// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PlaneBuilder represents a babylon.js PlaneBuilder.
// Class containing static functions to help procedurally build meshes
type PlaneBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PlaneBuilder) JSObject() js.Value { return p.p }

// PlaneBuilder returns a PlaneBuilder JavaScript class.
func (ba *Babylon) PlaneBuilder() *PlaneBuilder {
	p := ba.ctx.Get("PlaneBuilder")
	return PlaneBuilderFromJSObject(p, ba.ctx)
}

// PlaneBuilderFromJSObject returns a wrapped PlaneBuilder JavaScript class.
func PlaneBuilderFromJSObject(p js.Value, ctx js.Value) *PlaneBuilder {
	return &PlaneBuilder{p: p, ctx: ctx}
}

// PlaneBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func PlaneBuilderArrayToJSArray(array []*PlaneBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// PlaneBuilderCreatePlaneOpts contains optional parameters for PlaneBuilder.CreatePlane.
type PlaneBuilderCreatePlaneOpts struct {
	Scene *Scene
}

// CreatePlane calls the CreatePlane method on the PlaneBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.planebuilder#createplane
func (p *PlaneBuilder) CreatePlane(name string, options js.Value, opts *PlaneBuilderCreatePlaneOpts) *Mesh {
	if opts == nil {
		opts = &PlaneBuilderCreatePlaneOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)

	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := p.p.Call("CreatePlane", args...)
	return MeshFromJSObject(retVal, p.ctx)
}
