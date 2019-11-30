// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PlanePanel represents a babylon.js PlanePanel.
// Class used to create a container panel deployed on the surface of a plane
type PlanePanel struct {
	*VolumeBasedPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PlanePanel) JSObject() js.Value { return p.p }

// PlanePanel returns a PlanePanel JavaScript class.
func (ba *Babylon) PlanePanel() *PlanePanel {
	p := ba.ctx.Get("PlanePanel")
	return PlanePanelFromJSObject(p, ba.ctx)
}

// PlanePanelFromJSObject returns a wrapped PlanePanel JavaScript class.
func PlanePanelFromJSObject(p js.Value, ctx js.Value) *PlanePanel {
	return &PlanePanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// NewPlanePanel returns a new PlanePanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.planepanel
func (ba *Babylon) NewPlanePanel() *PlanePanel {
	p := ba.ctx.Get("PlanePanel").New()
	return PlanePanelFromJSObject(p, ba.ctx)
}

// TODO: methods
