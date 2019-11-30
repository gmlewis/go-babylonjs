// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EngineInstrumentation represents a babylon.js EngineInstrumentation.
// This class can be used to get instrumentation data from a Babylon engine
//
// See: http://doc.babylonjs.com/how_to/optimizing_your_scene#engineinstrumentation
type EngineInstrumentation struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EngineInstrumentation) JSObject() js.Value { return e.p }

// EngineInstrumentation returns a EngineInstrumentation JavaScript class.
func (b *Babylon) EngineInstrumentation() *EngineInstrumentation {
	p := b.ctx.Get("EngineInstrumentation")
	return EngineInstrumentationFromJSObject(p)
}

// EngineInstrumentationFromJSObject returns a wrapped EngineInstrumentation JavaScript class.
func EngineInstrumentationFromJSObject(p js.Value) *EngineInstrumentation {
	return &EngineInstrumentation{p: p}
}

// NewEngineInstrumentation returns a new EngineInstrumentation object.
//
// https://doc.babylonjs.com/api/classes/babylon.engineinstrumentation
func (b *Babylon) NewEngineInstrumentation(engine *Engine) *EngineInstrumentation {
	p := b.ctx.Get("EngineInstrumentation").New(engine.JSObject())
	return EngineInstrumentationFromJSObject(p)
}

// TODO: methods
