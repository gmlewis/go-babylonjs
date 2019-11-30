// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PositionGizmo represents a babylon.js PositionGizmo.
// Gizmo that enables dragging a mesh along 3 axis
type PositionGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (p *PositionGizmo) JSObject() js.Value { return p.p }

// PositionGizmo returns a PositionGizmo JavaScript class.
func (b *Babylon) PositionGizmo() *PositionGizmo {
	p := b.ctx.Get("PositionGizmo")
	return PositionGizmoFromJSObject(p)
}

// PositionGizmoFromJSObject returns a wrapped PositionGizmo JavaScript class.
func PositionGizmoFromJSObject(p js.Value) *PositionGizmo {
	return &PositionGizmo{GizmoFromJSObject(p)}
}

// NewPositionGizmoOpts contains optional parameters for NewPositionGizmo.
type NewPositionGizmoOpts struct {
	GizmoLayer *UtilityLayerRenderer
}

// NewPositionGizmo returns a new PositionGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.positiongizmo
func (b *Babylon) NewPositionGizmo(opts *NewPositionGizmoOpts) *PositionGizmo {
	if opts == nil {
		opts = &NewPositionGizmoOpts{}
	}

	p := b.ctx.Get("PositionGizmo").New(opts.GizmoLayer.JSObject())
	return PositionGizmoFromJSObject(p)
}

// TODO: methods
