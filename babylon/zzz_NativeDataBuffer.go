// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NativeDataBuffer represents a babylon.js NativeDataBuffer.
// Container for accessors for natively-stored mesh data buffers.
type NativeDataBuffer struct{ *DataBuffer }

// JSObject returns the underlying js.Value.
func (n *NativeDataBuffer) JSObject() js.Value { return n.p }

// NativeDataBuffer returns a NativeDataBuffer JavaScript class.
func (b *Babylon) NativeDataBuffer() *NativeDataBuffer {
	p := b.ctx.Get("NativeDataBuffer")
	return NativeDataBufferFromJSObject(p)
}

// NativeDataBufferFromJSObject returns a wrapped NativeDataBuffer JavaScript class.
func NativeDataBufferFromJSObject(p js.Value) *NativeDataBuffer {
	return &NativeDataBuffer{DataBufferFromJSObject(p)}
}

// NewNativeDataBuffer returns a new NativeDataBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.nativedatabuffer
func (b *Babylon) NewNativeDataBuffer(todo parameters) *NativeDataBuffer {
	p := b.ctx.Get("NativeDataBuffer").New(todo)
	return NativeDataBufferFromJSObject(p)
}

// TODO: methods
