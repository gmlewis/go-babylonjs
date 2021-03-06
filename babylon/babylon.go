package babylon

import (
	"syscall/js"
)

// JSObject can return its underlying js.Value.
type JSObject interface {
	JSObject() js.Value
}

// JSFunc represents a JavaScript function callback.
type JSFunc func(this js.Value, args []js.Value) interface{}

// Babylon represents the babylon.js API.
type Babylon struct {
	ctx js.Value
}

// New returns a new JavaScript object that binds to babylon.js.
func New() *Babylon { return &Babylon{ctx: js.Global().Get("BABYLON")} }

// JSObject returns the underlying *js.JavaScript class.
func (b *Babylon) JSObject() js.Value { return b.ctx }

// Bool returns the pointer to the provided bool.
func Bool(v bool) *bool {
	return &v
}

// Float64 returns the pointer to the provided float64.
func Float64(v float64) *float64 {
	return &v
}
