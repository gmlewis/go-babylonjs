// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TrailMesh represents a babylon.js TrailMesh.
// Class used to create a trail following a mesh
type TrailMesh struct{ *Mesh }

// JSObject returns the underlying js.Value.
func (t *TrailMesh) JSObject() js.Value { return t.p }

// TrailMesh returns a TrailMesh JavaScript class.
func (b *Babylon) TrailMesh() *TrailMesh {
	p := b.ctx.Get("TrailMesh")
	return TrailMeshFromJSObject(p)
}

// TrailMeshFromJSObject returns a wrapped TrailMesh JavaScript class.
func TrailMeshFromJSObject(p js.Value) *TrailMesh {
	return &TrailMesh{MeshFromJSObject(p)}
}

// NewTrailMesh returns a new TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh
func (b *Babylon) NewTrailMesh(todo parameters) *TrailMesh {
	p := b.ctx.Get("TrailMesh").New(todo)
	return TrailMeshFromJSObject(p)
}

// TODO: methods