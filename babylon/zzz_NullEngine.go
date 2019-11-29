// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NullEngine represents a babylon.js NullEngine.
// The null engine class provides support for headless version of babylon.js.
// This can be used in server side scenario or for testing purposes
type NullEngine struct{ *Engine }

// JSObject returns the underlying js.Value.
func (n *NullEngine) JSObject() js.Value { return n.p }

// NullEngine returns a NullEngine JavaScript class.
func (b *Babylon) NullEngine() *NullEngine {
	p := b.ctx.Get("NullEngine")
	return NullEngineFromJSObject(p)
}

// NullEngineFromJSObject returns a wrapped NullEngine JavaScript class.
func NullEngineFromJSObject(p js.Value) *NullEngine {
	return &NullEngine{EngineFromJSObject(p)}
}

// NewNullEngine returns a new NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine
func (b *Babylon) NewNullEngine(todo parameters) *NullEngine {
	p := b.ctx.Get("NullEngine").New(todo)
	return NullEngineFromJSObject(p)
}

// TODO: methods
