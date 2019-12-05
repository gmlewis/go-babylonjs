package babylon

import (
	"syscall/js"
)

// GUI represents the babylon.js GUI namespace.
type GUI struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *GUI) JSObject() js.Value { return a.p }

// GUI returns a GUI JavaScript class.
func (ba *Babylon) GUI() *GUI {
	p := ba.ctx.Get("GUI")
	return GUIFromJSObject(p, p)
}

// GUIFromJSObject returns a wrapped GUI JavaScript class.
func GUIFromJSObject(p js.Value, ctx js.Value) *GUI {
	return &GUI{p: p, ctx: ctx}
}
