// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GizmoManager represents a babylon.js GizmoManager.
// Helps setup gizmo&amp;#39;s in the scene to rotate/scale/position meshes
type GizmoManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GizmoManager) JSObject() js.Value { return g.p }

// GizmoManager returns a GizmoManager JavaScript class.
func (ba *Babylon) GizmoManager() *GizmoManager {
	p := ba.ctx.Get("GizmoManager")
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// GizmoManagerFromJSObject returns a wrapped GizmoManager JavaScript class.
func GizmoManagerFromJSObject(p js.Value, ctx js.Value) *GizmoManager {
	return &GizmoManager{p: p, ctx: ctx}
}

// GizmoManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GizmoManagerArrayToJSArray(array []*GizmoManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGizmoManager returns a new GizmoManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager
func (ba *Babylon) NewGizmoManager(scene *Scene) *GizmoManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("GizmoManager").New(args...)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// AttachToMesh calls the AttachToMesh method on the GizmoManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#attachtomesh
func (g *GizmoManager) AttachToMesh(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("attachToMesh", args...)
}

// Dispose calls the Dispose method on the GizmoManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#dispose
func (g *GizmoManager) Dispose() {

	g.p.Call("dispose")
}

/*

// AttachableMeshes returns the AttachableMeshes property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#attachablemeshes
func (g *GizmoManager) AttachableMeshes(attachableMeshes []*AbstractMesh) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(attachableMeshes)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetAttachableMeshes sets the AttachableMeshes property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#attachablemeshes
func (g *GizmoManager) SetAttachableMeshes(attachableMeshes []*AbstractMesh) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(attachableMeshes)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// BoundingBoxDragBehavior returns the BoundingBoxDragBehavior property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#boundingboxdragbehavior
func (g *GizmoManager) BoundingBoxDragBehavior(boundingBoxDragBehavior *SixDofDragBehavior) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(boundingBoxDragBehavior.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetBoundingBoxDragBehavior sets the BoundingBoxDragBehavior property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#boundingboxdragbehavior
func (g *GizmoManager) SetBoundingBoxDragBehavior(boundingBoxDragBehavior *SixDofDragBehavior) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(boundingBoxDragBehavior.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// BoundingBoxGizmoEnabled returns the BoundingBoxGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#boundingboxgizmoenabled
func (g *GizmoManager) BoundingBoxGizmoEnabled(boundingBoxGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(boundingBoxGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetBoundingBoxGizmoEnabled sets the BoundingBoxGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#boundingboxgizmoenabled
func (g *GizmoManager) SetBoundingBoxGizmoEnabled(boundingBoxGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(boundingBoxGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// ClearGizmoOnEmptyPointerEvent returns the ClearGizmoOnEmptyPointerEvent property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#cleargizmoonemptypointerevent
func (g *GizmoManager) ClearGizmoOnEmptyPointerEvent(clearGizmoOnEmptyPointerEvent bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(clearGizmoOnEmptyPointerEvent)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetClearGizmoOnEmptyPointerEvent sets the ClearGizmoOnEmptyPointerEvent property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#cleargizmoonemptypointerevent
func (g *GizmoManager) SetClearGizmoOnEmptyPointerEvent(clearGizmoOnEmptyPointerEvent bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(clearGizmoOnEmptyPointerEvent)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// Gizmos returns the Gizmos property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#gizmos
func (g *GizmoManager) Gizmos(gizmos js.Value) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(gizmos)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetGizmos sets the Gizmos property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#gizmos
func (g *GizmoManager) SetGizmos(gizmos js.Value) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(gizmos)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// KeepDepthUtilityLayer returns the KeepDepthUtilityLayer property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#keepdepthutilitylayer
func (g *GizmoManager) KeepDepthUtilityLayer(keepDepthUtilityLayer *UtilityLayerRenderer) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(keepDepthUtilityLayer.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetKeepDepthUtilityLayer sets the KeepDepthUtilityLayer property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#keepdepthutilitylayer
func (g *GizmoManager) SetKeepDepthUtilityLayer(keepDepthUtilityLayer *UtilityLayerRenderer) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(keepDepthUtilityLayer.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// OnAttachedToMeshObservable returns the OnAttachedToMeshObservable property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#onattachedtomeshobservable
func (g *GizmoManager) OnAttachedToMeshObservable(onAttachedToMeshObservable *Observable) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(onAttachedToMeshObservable.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetOnAttachedToMeshObservable sets the OnAttachedToMeshObservable property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#onattachedtomeshobservable
func (g *GizmoManager) SetOnAttachedToMeshObservable(onAttachedToMeshObservable *Observable) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(onAttachedToMeshObservable.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// PositionGizmoEnabled returns the PositionGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#positiongizmoenabled
func (g *GizmoManager) PositionGizmoEnabled(positionGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(positionGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetPositionGizmoEnabled sets the PositionGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#positiongizmoenabled
func (g *GizmoManager) SetPositionGizmoEnabled(positionGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(positionGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// RotationGizmoEnabled returns the RotationGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#rotationgizmoenabled
func (g *GizmoManager) RotationGizmoEnabled(rotationGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(rotationGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetRotationGizmoEnabled sets the RotationGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#rotationgizmoenabled
func (g *GizmoManager) SetRotationGizmoEnabled(rotationGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(rotationGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// ScaleGizmoEnabled returns the ScaleGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#scalegizmoenabled
func (g *GizmoManager) ScaleGizmoEnabled(scaleGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(scaleGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetScaleGizmoEnabled sets the ScaleGizmoEnabled property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#scalegizmoenabled
func (g *GizmoManager) SetScaleGizmoEnabled(scaleGizmoEnabled bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(scaleGizmoEnabled)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// UsePointerToAttachGizmos returns the UsePointerToAttachGizmos property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#usepointertoattachgizmos
func (g *GizmoManager) UsePointerToAttachGizmos(usePointerToAttachGizmos bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(usePointerToAttachGizmos)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetUsePointerToAttachGizmos sets the UsePointerToAttachGizmos property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#usepointertoattachgizmos
func (g *GizmoManager) SetUsePointerToAttachGizmos(usePointerToAttachGizmos bool) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(usePointerToAttachGizmos)
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// UtilityLayer returns the UtilityLayer property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#utilitylayer
func (g *GizmoManager) UtilityLayer(utilityLayer *UtilityLayerRenderer) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(utilityLayer.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

// SetUtilityLayer sets the UtilityLayer property of class GizmoManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gizmomanager#utilitylayer
func (g *GizmoManager) SetUtilityLayer(utilityLayer *UtilityLayerRenderer) *GizmoManager {
	p := ba.ctx.Get("GizmoManager").New(utilityLayer.JSObject())
	return GizmoManagerFromJSObject(p, ba.ctx)
}

*/
