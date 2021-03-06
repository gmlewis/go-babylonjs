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
func (gui *GUI) PlanePanel() *PlanePanel {
	p := gui.ctx.Get("PlanePanel")
	return PlanePanelFromJSObject(p, gui.ctx)
}

// PlanePanelFromJSObject returns a wrapped PlanePanel JavaScript class.
func PlanePanelFromJSObject(p js.Value, ctx js.Value) *PlanePanel {
	return &PlanePanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// PlanePanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func PlanePanelArrayToJSArray(array []*PlanePanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPlanePanel returns a new PlanePanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.planepanel#constructor
func (gui *GUI) NewPlanePanel() *PlanePanel {

	args := make([]interface{}, 0, 0+0)

	p := gui.ctx.Get("PlanePanel").New(args...)
	return PlanePanelFromJSObject(p, gui.ctx)
}
