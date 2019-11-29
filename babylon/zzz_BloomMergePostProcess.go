// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BloomMergePostProcess represents a babylon.js BloomMergePostProcess.
// The BloomMergePostProcess merges blurred images with the original based on the values of the circle of confusion.
type BloomMergePostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (b *BloomMergePostProcess) JSObject() js.Value { return b.p }

// BloomMergePostProcess returns a BloomMergePostProcess JavaScript class.
func (b *Babylon) BloomMergePostProcess() *BloomMergePostProcess {
	p := b.ctx.Get("BloomMergePostProcess")
	return BloomMergePostProcessFromJSObject(p)
}

// BloomMergePostProcessFromJSObject returns a wrapped BloomMergePostProcess JavaScript class.
func BloomMergePostProcessFromJSObject(p js.Value) *BloomMergePostProcess {
	return &BloomMergePostProcess{PostProcessFromJSObject(p)}
}

// NewBloomMergePostProcess returns a new BloomMergePostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.bloommergepostprocess
func (b *Babylon) NewBloomMergePostProcess(todo parameters) *BloomMergePostProcess {
	p := b.ctx.Get("BloomMergePostProcess").New(todo)
	return BloomMergePostProcessFromJSObject(p)
}

// TODO: methods
