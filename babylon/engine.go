package babylon

import (
	"syscall/js"
)

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
