// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AxisScaleGizmo represents a babylon.js AxisScaleGizmo.
// Single axis scale gizmo
type AxisScaleGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (a *AxisScaleGizmo) JSObject() js.Value { return a.p }

// AxisScaleGizmo returns a AxisScaleGizmo JavaScript class.
func (b *Babylon) AxisScaleGizmo() *AxisScaleGizmo {
	p := b.ctx.Get("AxisScaleGizmo")
	return AxisScaleGizmoFromJSObject(p)
}

// AxisScaleGizmoFromJSObject returns a wrapped AxisScaleGizmo JavaScript class.
func AxisScaleGizmoFromJSObject(p js.Value) *AxisScaleGizmo {
	return &AxisScaleGizmo{GizmoFromJSObject(p)}
}

// NewAxisScaleGizmo returns a new AxisScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo
func (b *Babylon) NewAxisScaleGizmo(todo parameters) *AxisScaleGizmo {
	p := b.ctx.Get("AxisScaleGizmo").New(todo)
	return AxisScaleGizmoFromJSObject(p)
}

// TODO: methods
