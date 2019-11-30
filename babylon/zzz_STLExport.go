// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// STLExport represents a babylon.js STLExport.
// Class for generating STL data from a Babylon scene.
type STLExport struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *STLExport) JSObject() js.Value { return s.p }

// STLExport returns a STLExport JavaScript class.
func (ba *Babylon) STLExport() *STLExport {
	p := ba.ctx.Get("STLExport")
	return STLExportFromJSObject(p, ba.ctx)
}

// STLExportFromJSObject returns a wrapped STLExport JavaScript class.
func STLExportFromJSObject(p js.Value, ctx js.Value) *STLExport {
	return &STLExport{p: p, ctx: ctx}
}

// TODO: methods
