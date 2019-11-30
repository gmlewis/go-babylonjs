// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RGBDTextureTools represents a babylon.js RGBDTextureTools.
// Class used to host RGBD texture specific utilities
type RGBDTextureTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RGBDTextureTools) JSObject() js.Value { return r.p }

// RGBDTextureTools returns a RGBDTextureTools JavaScript class.
func (ba *Babylon) RGBDTextureTools() *RGBDTextureTools {
	p := ba.ctx.Get("RGBDTextureTools")
	return RGBDTextureToolsFromJSObject(p, ba.ctx)
}

// RGBDTextureToolsFromJSObject returns a wrapped RGBDTextureTools JavaScript class.
func RGBDTextureToolsFromJSObject(p js.Value, ctx js.Value) *RGBDTextureTools {
	return &RGBDTextureTools{p: p, ctx: ctx}
}

// TODO: methods
