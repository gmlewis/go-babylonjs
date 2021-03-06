// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Gizmo represents a babylon.js Gizmo.
// Renders gizmos on top of an existing scene which provide controls for position, rotation, etc.
type Gizmo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *Gizmo) JSObject() js.Value { return g.p }

// Gizmo returns a Gizmo JavaScript class.
func (ba *Babylon) Gizmo() *Gizmo {
	p := ba.ctx.Get("Gizmo")
	return GizmoFromJSObject(p, ba.ctx)
}

// GizmoFromJSObject returns a wrapped Gizmo JavaScript class.
func GizmoFromJSObject(p js.Value, ctx js.Value) *Gizmo {
	return &Gizmo{p: p, ctx: ctx}
}

// GizmoArrayToJSArray returns a JavaScript Array for the wrapped array.
func GizmoArrayToJSArray(array []*Gizmo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGizmoOpts contains optional parameters for NewGizmo.
type NewGizmoOpts struct {
	GizmoLayer *UtilityLayerRenderer
}

// NewGizmo returns a new Gizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#constructor
func (ba *Babylon) NewGizmo(opts *NewGizmoOpts) *Gizmo {
	if opts == nil {
		opts = &NewGizmoOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.GizmoLayer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.GizmoLayer.JSObject())
	}

	p := ba.ctx.Get("Gizmo").New(args...)
	return GizmoFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the Gizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#dispose
func (g *Gizmo) Dispose() {

	g.p.Call("dispose")
}

// SetCustomMesh calls the SetCustomMesh method on the Gizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#setcustommesh
func (g *Gizmo) SetCustomMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	g.p.Call("setCustomMesh", args...)
}

// AttachedMesh returns the AttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#attachedmesh
func (g *Gizmo) AttachedMesh() *AbstractMesh {
	retVal := g.p.Get("attachedMesh")
	return AbstractMeshFromJSObject(retVal, g.ctx)
}

// SetAttachedMesh sets the AttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#attachedmesh
func (g *Gizmo) SetAttachedMesh(attachedMesh *AbstractMesh) *Gizmo {
	g.p.Set("attachedMesh", attachedMesh.JSObject())
	return g
}

// GizmoLayer returns the GizmoLayer property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#gizmolayer
func (g *Gizmo) GizmoLayer() *UtilityLayerRenderer {
	retVal := g.p.Get("gizmoLayer")
	return UtilityLayerRendererFromJSObject(retVal, g.ctx)
}

// SetGizmoLayer sets the GizmoLayer property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#gizmolayer
func (g *Gizmo) SetGizmoLayer(gizmoLayer *UtilityLayerRenderer) *Gizmo {
	g.p.Set("gizmoLayer", gizmoLayer.JSObject())
	return g
}

// ScaleRatio returns the ScaleRatio property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#scaleratio
func (g *Gizmo) ScaleRatio() float64 {
	retVal := g.p.Get("scaleRatio")
	return retVal.Float()
}

// SetScaleRatio sets the ScaleRatio property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#scaleratio
func (g *Gizmo) SetScaleRatio(scaleRatio float64) *Gizmo {
	g.p.Set("scaleRatio", scaleRatio)
	return g
}

// UpdateGizmoPositionToMatchAttachedMesh returns the UpdateGizmoPositionToMatchAttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updategizmopositiontomatchattachedmesh
func (g *Gizmo) UpdateGizmoPositionToMatchAttachedMesh() bool {
	retVal := g.p.Get("updateGizmoPositionToMatchAttachedMesh")
	return retVal.Bool()
}

// SetUpdateGizmoPositionToMatchAttachedMesh sets the UpdateGizmoPositionToMatchAttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updategizmopositiontomatchattachedmesh
func (g *Gizmo) SetUpdateGizmoPositionToMatchAttachedMesh(updateGizmoPositionToMatchAttachedMesh bool) *Gizmo {
	g.p.Set("updateGizmoPositionToMatchAttachedMesh", updateGizmoPositionToMatchAttachedMesh)
	return g
}

// UpdateGizmoRotationToMatchAttachedMesh returns the UpdateGizmoRotationToMatchAttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updategizmorotationtomatchattachedmesh
func (g *Gizmo) UpdateGizmoRotationToMatchAttachedMesh() bool {
	retVal := g.p.Get("updateGizmoRotationToMatchAttachedMesh")
	return retVal.Bool()
}

// SetUpdateGizmoRotationToMatchAttachedMesh sets the UpdateGizmoRotationToMatchAttachedMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updategizmorotationtomatchattachedmesh
func (g *Gizmo) SetUpdateGizmoRotationToMatchAttachedMesh(updateGizmoRotationToMatchAttachedMesh bool) *Gizmo {
	g.p.Set("updateGizmoRotationToMatchAttachedMesh", updateGizmoRotationToMatchAttachedMesh)
	return g
}

// UpdateScale returns the UpdateScale property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updatescale
func (g *Gizmo) UpdateScale() bool {
	retVal := g.p.Get("updateScale")
	return retVal.Bool()
}

// SetUpdateScale sets the UpdateScale property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#updatescale
func (g *Gizmo) SetUpdateScale(updateScale bool) *Gizmo {
	g.p.Set("updateScale", updateScale)
	return g
}

// _rootMesh returns the _rootMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#_rootmesh
func (g *Gizmo) _rootMesh() *Mesh {
	retVal := g.p.Get("_rootMesh")
	return MeshFromJSObject(retVal, g.ctx)
}

// Set_rootMesh sets the _rootMesh property of class Gizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmo#_rootmesh
func (g *Gizmo) Set_rootMesh(_rootMesh *Mesh) *Gizmo {
	g.p.Set("_rootMesh", _rootMesh.JSObject())
	return g
}
