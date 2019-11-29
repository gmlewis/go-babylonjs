package babylon

import "syscall/js"

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

func (b *Babylon) NewEngine(canvas Canvas, antialias bool) *Engine {
	p := b.ctx.Get("Engine").New(canvas, antialias)
	return EngineFromJSObject(p)
}

func (e Engine) RunRenderLoop(f func()) {
	f()
}

func (e Engine) Resize() {}
