// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EngineView represents a babylon.js EngineView.
// Class used to define an additional view for the engine

//
// See: https://doc.babylonjs.com/how_to/multi_canvases
type EngineView struct{}

// JSObject returns the underlying js.Value.
func (e *EngineView) JSObject() js.Value { return e.p }

// EngineView returns a EngineView JavaScript class.
func (b *Babylon) EngineView() *EngineView {
	p := b.ctx.Get("EngineView")
	return EngineViewFromJSObject(p)
}

// EngineViewFromJSObject returns a wrapped EngineView JavaScript class.
func EngineViewFromJSObject(p js.Value) *EngineView {
	return &EngineView{p: p}
}

// NewEngineView returns a new EngineView object.
//
// https://doc.babylonjs.com/api/classes/babylon.engineview
func (b *Babylon) NewEngineView(todo parameters) *EngineView {
	p := b.ctx.Get("EngineView").New(todo)
	return EngineViewFromJSObject(p)
}

// TODO: methods