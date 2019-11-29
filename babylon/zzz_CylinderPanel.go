// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderPanel represents a babylon.js CylinderPanel.
// Class used to create a container panel deployed on the surface of a cylinder
type CylinderPanel struct{ *VolumeBasedPanel }

// JSObject returns the underlying js.Value.
func (c *CylinderPanel) JSObject() js.Value { return c.p }

// CylinderPanel returns a CylinderPanel JavaScript class.
func (b *Babylon) CylinderPanel() *CylinderPanel {
	p := b.ctx.Get("CylinderPanel")
	return CylinderPanelFromJSObject(p)
}

// CylinderPanelFromJSObject returns a wrapped CylinderPanel JavaScript class.
func CylinderPanelFromJSObject(p js.Value) *CylinderPanel {
	return &CylinderPanel{VolumeBasedPanelFromJSObject(p)}
}

// NewCylinderPanel returns a new CylinderPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderpanel
func (b *Babylon) NewCylinderPanel(todo parameters) *CylinderPanel {
	p := b.ctx.Get("CylinderPanel").New(todo)
	return CylinderPanelFromJSObject(p)
}

// TODO: methods
