// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScreenshotTools represents a babylon.js ScreenshotTools.
// Class containing a set of static utilities functions for screenshots
type ScreenshotTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScreenshotTools) JSObject() js.Value { return s.p }

// ScreenshotTools returns a ScreenshotTools JavaScript class.
func (ba *Babylon) ScreenshotTools() *ScreenshotTools {
	p := ba.ctx.Get("ScreenshotTools")
	return ScreenshotToolsFromJSObject(p, ba.ctx)
}

// ScreenshotToolsFromJSObject returns a wrapped ScreenshotTools JavaScript class.
func ScreenshotToolsFromJSObject(p js.Value, ctx js.Value) *ScreenshotTools {
	return &ScreenshotTools{p: p, ctx: ctx}
}

// TODO: methods
