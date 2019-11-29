// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GroundMesh represents a babylon.js GroundMesh.
// Mesh representing the gorund
type GroundMesh struct{ *Mesh }

// JSObject returns the underlying js.Value.
func (g *GroundMesh) JSObject() js.Value { return g.p }

// GroundMesh returns a GroundMesh JavaScript class.
func (b *Babylon) GroundMesh() *GroundMesh {
	p := b.ctx.Get("GroundMesh")
	return GroundMeshFromJSObject(p)
}

// GroundMeshFromJSObject returns a wrapped GroundMesh JavaScript class.
func GroundMeshFromJSObject(p js.Value) *GroundMesh {
	return &GroundMesh{MeshFromJSObject(p)}
}

// NewGroundMesh returns a new GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh
func (b *Babylon) NewGroundMesh(todo parameters) *GroundMesh {
	p := b.ctx.Get("GroundMesh").New(todo)
	return GroundMeshFromJSObject(p)
}

// TODO: methods