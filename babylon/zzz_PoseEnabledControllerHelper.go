// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PoseEnabledControllerHelper represents a babylon.js PoseEnabledControllerHelper.
// Defines the PoseEnabledControllerHelper object that is used initialize a gamepad as the controller type it is specified as (eg. windows mixed reality controller)
type PoseEnabledControllerHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PoseEnabledControllerHelper) JSObject() js.Value { return p.p }

// PoseEnabledControllerHelper returns a PoseEnabledControllerHelper JavaScript class.
func (ba *Babylon) PoseEnabledControllerHelper() *PoseEnabledControllerHelper {
	p := ba.ctx.Get("PoseEnabledControllerHelper")
	return PoseEnabledControllerHelperFromJSObject(p, ba.ctx)
}

// PoseEnabledControllerHelperFromJSObject returns a wrapped PoseEnabledControllerHelper JavaScript class.
func PoseEnabledControllerHelperFromJSObject(p js.Value, ctx js.Value) *PoseEnabledControllerHelper {
	return &PoseEnabledControllerHelper{p: p, ctx: ctx}
}

// TODO: methods
