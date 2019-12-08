package babylon

import "syscall/js"

// ResizeFunc is a convenience function to reduce typing.
func (e *Engine) ResizeFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.Resize()
		return nil
	})
}
