// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HighlightsPostProcess represents a babylon.js HighlightsPostProcess.
// Extracts highlights from the image

//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses
type HighlightsPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (h *HighlightsPostProcess) JSObject() js.Value { return h.p }

// HighlightsPostProcess returns a HighlightsPostProcess JavaScript class.
func (b *Babylon) HighlightsPostProcess() *HighlightsPostProcess {
	p := b.ctx.Get("HighlightsPostProcess")
	return HighlightsPostProcessFromJSObject(p)
}

// HighlightsPostProcessFromJSObject returns a wrapped HighlightsPostProcess JavaScript class.
func HighlightsPostProcessFromJSObject(p js.Value) *HighlightsPostProcess {
	return &HighlightsPostProcess{PostProcessFromJSObject(p)}
}

// NewHighlightsPostProcess returns a new HighlightsPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.highlightspostprocess
func (b *Babylon) NewHighlightsPostProcess(todo parameters) *HighlightsPostProcess {
	p := b.ctx.Get("HighlightsPostProcess").New(todo)
	return HighlightsPostProcessFromJSObject(p)
}

// TODO: methods