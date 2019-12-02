// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderPanel represents a babylon.js CylinderPanel.
// Class used to create a container panel deployed on the surface of a cylinder
type CylinderPanel struct {
	*VolumeBasedPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CylinderPanel) JSObject() js.Value { return c.p }

// CylinderPanel returns a CylinderPanel JavaScript class.
func (ba *Babylon) CylinderPanel() *CylinderPanel {
	p := ba.ctx.Get("CylinderPanel")
	return CylinderPanelFromJSObject(p, ba.ctx)
}

// CylinderPanelFromJSObject returns a wrapped CylinderPanel JavaScript class.
func CylinderPanelFromJSObject(p js.Value, ctx js.Value) *CylinderPanel {
	return &CylinderPanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// CylinderPanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func CylinderPanelArrayToJSArray(array []*CylinderPanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCylinderPanel returns a new CylinderPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderpanel
func (ba *Babylon) NewCylinderPanel() *CylinderPanel {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("CylinderPanel").New(args...)
	return CylinderPanelFromJSObject(p, ba.ctx)
}

/*

// Radius returns the Radius property of class CylinderPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderpanel#radius
func (c *CylinderPanel) Radius(radius float64) *CylinderPanel {
	p := ba.ctx.Get("CylinderPanel").New(radius)
	return CylinderPanelFromJSObject(p, ba.ctx)
}

// SetRadius sets the Radius property of class CylinderPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderpanel#radius
func (c *CylinderPanel) SetRadius(radius float64) *CylinderPanel {
	p := ba.ctx.Get("CylinderPanel").New(radius)
	return CylinderPanelFromJSObject(p, ba.ctx)
}

*/
