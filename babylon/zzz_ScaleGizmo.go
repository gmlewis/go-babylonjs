// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScaleGizmo represents a babylon.js ScaleGizmo.
// Gizmo that enables scaling a mesh along 3 axis
type ScaleGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (s *ScaleGizmo) JSObject() js.Value { return s.p }

// ScaleGizmo returns a ScaleGizmo JavaScript class.
func (b *Babylon) ScaleGizmo() *ScaleGizmo {
	p := b.ctx.Get("ScaleGizmo")
	return ScaleGizmoFromJSObject(p)
}

// ScaleGizmoFromJSObject returns a wrapped ScaleGizmo JavaScript class.
func ScaleGizmoFromJSObject(p js.Value) *ScaleGizmo {
	return &ScaleGizmo{GizmoFromJSObject(p)}
}

// NewScaleGizmo returns a new ScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalegizmo
func (b *Babylon) NewScaleGizmo(todo parameters) *ScaleGizmo {
	p := b.ctx.Get("ScaleGizmo").New(todo)
	return ScaleGizmoFromJSObject(p)
}

// TODO: methods
