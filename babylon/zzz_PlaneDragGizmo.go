// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PlaneDragGizmo represents a babylon.js PlaneDragGizmo.
// Single plane drag gizmo
type PlaneDragGizmo struct {
	*Gizmo
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PlaneDragGizmo) JSObject() js.Value { return p.p }

// PlaneDragGizmo returns a PlaneDragGizmo JavaScript class.
func (ba *Babylon) PlaneDragGizmo() *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo")
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// PlaneDragGizmoFromJSObject returns a wrapped PlaneDragGizmo JavaScript class.
func PlaneDragGizmoFromJSObject(p js.Value, ctx js.Value) *PlaneDragGizmo {
	return &PlaneDragGizmo{Gizmo: GizmoFromJSObject(p, ctx), ctx: ctx}
}

// PlaneDragGizmoArrayToJSArray returns a JavaScript Array for the wrapped array.
func PlaneDragGizmoArrayToJSArray(array []*PlaneDragGizmo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPlaneDragGizmoOpts contains optional parameters for NewPlaneDragGizmo.
type NewPlaneDragGizmoOpts struct {
	Color      *Color3
	GizmoLayer *UtilityLayerRenderer
	Parent     *PositionGizmo
}

// NewPlaneDragGizmo returns a new PlaneDragGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo
func (ba *Babylon) NewPlaneDragGizmo(dragPlaneNormal *Vector3, opts *NewPlaneDragGizmoOpts) *PlaneDragGizmo {
	if opts == nil {
		opts = &NewPlaneDragGizmoOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, dragPlaneNormal.JSObject())

	if opts.Color == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Color.JSObject())
	}
	if opts.GizmoLayer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.GizmoLayer.JSObject())
	}
	if opts.Parent == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Parent.JSObject())
	}

	p := ba.ctx.Get("PlaneDragGizmo").New(args...)
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the PlaneDragGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#dispose
func (p *PlaneDragGizmo) Dispose() {

	p.p.Call("dispose")
}

/*

// DragBehavior returns the DragBehavior property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#dragbehavior
func (p *PlaneDragGizmo) DragBehavior(dragBehavior *PointerDragBehavior) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(dragBehavior.JSObject())
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// SetDragBehavior sets the DragBehavior property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#dragbehavior
func (p *PlaneDragGizmo) SetDragBehavior(dragBehavior *PointerDragBehavior) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(dragBehavior.JSObject())
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// IsEnabled returns the IsEnabled property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#isenabled
func (p *PlaneDragGizmo) IsEnabled(isEnabled bool) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(isEnabled)
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// SetIsEnabled sets the IsEnabled property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#isenabled
func (p *PlaneDragGizmo) SetIsEnabled(isEnabled bool) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(isEnabled)
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// OnSnapObservable returns the OnSnapObservable property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#onsnapobservable
func (p *PlaneDragGizmo) OnSnapObservable(onSnapObservable *Observable) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(onSnapObservable.JSObject())
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// SetOnSnapObservable sets the OnSnapObservable property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#onsnapobservable
func (p *PlaneDragGizmo) SetOnSnapObservable(onSnapObservable *Observable) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(onSnapObservable.JSObject())
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// SnapDistance returns the SnapDistance property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#snapdistance
func (p *PlaneDragGizmo) SnapDistance(snapDistance float64) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(snapDistance)
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

// SetSnapDistance sets the SnapDistance property of class PlaneDragGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.planedraggizmo#snapdistance
func (p *PlaneDragGizmo) SetSnapDistance(snapDistance float64) *PlaneDragGizmo {
	p := ba.ctx.Get("PlaneDragGizmo").New(snapDistance)
	return PlaneDragGizmoFromJSObject(p, ba.ctx)
}

*/
