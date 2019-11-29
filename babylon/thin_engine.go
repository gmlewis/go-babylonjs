package babylon

import "syscall/js"

// ThinEngine represents a Babylon ThinEngine.
type ThinEngine struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *ThinEngine) JSObject() js.Value { return e.p }

// ThinEngine returns a ThinEngine JavaScript class.
func (b *Babylon) ThinEngine() *ThinEngine {
	p := b.ctx.Get("ThinEngine")
	return ThinEngineFromJSObject(p)
}

// ThinEngineFromJSObject returns a wrapped ThinEngine JavaScript class.
func ThinEngineFromJSObject(p js.Value) *ThinEngine {
	return &ThinEngine{p: p}
}

func (b *Babylon) NewThinEngine(canvas Canvas, antialias bool) *ThinEngine {
	p := b.ctx.Get("ThinEngine").New(canvas, antialias)
	return ThinEngineFromJSObject(p)
}

func (e ThinEngine) RunRenderLoop(f func()) {
	f()
}

func (e ThinEngine) Resize() {}
