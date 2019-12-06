// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpherePanel represents a babylon.js SpherePanel.
// Class used to create a container panel deployed on the surface of a sphere
type SpherePanel struct {
	*VolumeBasedPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SpherePanel) JSObject() js.Value { return s.p }

// SpherePanel returns a SpherePanel JavaScript class.
func (gui *GUI) SpherePanel() *SpherePanel {
	p := gui.ctx.Get("SpherePanel")
	return SpherePanelFromJSObject(p, gui.ctx)
}

// SpherePanelFromJSObject returns a wrapped SpherePanel JavaScript class.
func SpherePanelFromJSObject(p js.Value, ctx js.Value) *SpherePanel {
	return &SpherePanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// SpherePanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func SpherePanelArrayToJSArray(array []*SpherePanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSpherePanel returns a new SpherePanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.spherepanel
func (gui *GUI) NewSpherePanel() *SpherePanel {

	args := make([]interface{}, 0, 0+0)

	p := gui.ctx.Get("SpherePanel").New(args...)
	return SpherePanelFromJSObject(p, gui.ctx)
}

// Radius returns the Radius property of class SpherePanel.
//
// https://doc.babylonjs.com/api/classes/babylon.spherepanel#radius
func (s *SpherePanel) Radius() float64 {
	retVal := s.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class SpherePanel.
//
// https://doc.babylonjs.com/api/classes/babylon.spherepanel#radius
func (s *SpherePanel) SetRadius(radius float64) *SpherePanel {
	s.p.Set("radius", radius)
	return s
}
