// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RefractionPostProcess represents a babylon.js RefractionPostProcess.
// Post process which applies a refractin texture

//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses#refraction
type RefractionPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (r *RefractionPostProcess) JSObject() js.Value { return r.p }

// RefractionPostProcess returns a RefractionPostProcess JavaScript class.
func (b *Babylon) RefractionPostProcess() *RefractionPostProcess {
	p := b.ctx.Get("RefractionPostProcess")
	return RefractionPostProcessFromJSObject(p)
}

// RefractionPostProcessFromJSObject returns a wrapped RefractionPostProcess JavaScript class.
func RefractionPostProcessFromJSObject(p js.Value) *RefractionPostProcess {
	return &RefractionPostProcess{PostProcessFromJSObject(p)}
}

// NewRefractionPostProcess returns a new RefractionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess
func (b *Babylon) NewRefractionPostProcess(todo parameters) *RefractionPostProcess {
	p := b.ctx.Get("RefractionPostProcess").New(todo)
	return RefractionPostProcessFromJSObject(p)
}

// TODO: methods
