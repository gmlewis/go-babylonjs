// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DebugLayer represents a babylon.js DebugLayer.
// The debug layer (aka Inspector) is the go to tool in order to better understand
// what is happening in your scene
//
// See: http://doc.babylonjs.com/features/playground_debuglayer
type DebugLayer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (d *DebugLayer) JSObject() js.Value { return d.p }

// DebugLayer returns a DebugLayer JavaScript class.
func (b *Babylon) DebugLayer() *DebugLayer {
	p := b.ctx.Get("DebugLayer")
	return DebugLayerFromJSObject(p)
}

// DebugLayerFromJSObject returns a wrapped DebugLayer JavaScript class.
func DebugLayerFromJSObject(p js.Value) *DebugLayer {
	return &DebugLayer{p: p}
}

// NewDebugLayer returns a new DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer
func (b *Babylon) NewDebugLayer(scene *Scene) *DebugLayer {
	p := b.ctx.Get("DebugLayer").New(scene.JSObject())
	return DebugLayerFromJSObject(p)
}

// TODO: methods
