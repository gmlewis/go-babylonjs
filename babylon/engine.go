package babylon

import (
	"syscall/js"
)

/*
// Engine represents a Babylon Engine.
type Engine struct{ *ThinEngine }

// JSObject returns the underlying js.Value.
func (e *Engine) JSObject() js.Value { return e.p }

// Engine returns a Engine JavaScript class.
func (b *Babylon) Engine() *Engine {
	p := b.ctx.Get("Engine")
	return EngineFromJSObject(p)
}

// EngineFromJSObject returns a wrapped Engine JavaScript class.
func EngineFromJSObject(p js.Value) *Engine {
	return &Engine{ThinEngineFromJSObject(p)}
}

// NewEngine returns a new Engine object.
//
// https://doc.babylonjs.com/api/classes/babylon.engine
func (b *Babylon) NewEngine(canvas js.Value, antialias bool) *Engine {
	p := b.ctx.Get("Engine").New(canvas, antialias)
	return EngineFromJSObject(p)
}
*/

// RunRenderLoop calls the JavaScript method of the same name.
func (e *Engine) RunRenderLoop(f func()) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	// Note that "cb" is never released since it is the render loop.
	e.p.Call("runRenderLoop", cb)
}

// Resize calls the JavaScript method of the same name.
func (e *Engine) Resize() {
	e.p.Call("resize")
}
