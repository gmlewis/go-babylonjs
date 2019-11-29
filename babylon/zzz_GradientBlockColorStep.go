// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GradientBlockColorStep represents a babylon.js GradientBlockColorStep.
// Class used to store a color step for the GradientBlock
type GradientBlockColorStep struct{}

// JSObject returns the underlying js.Value.
func (g *GradientBlockColorStep) JSObject() js.Value { return g.p }

// GradientBlockColorStep returns a GradientBlockColorStep JavaScript class.
func (b *Babylon) GradientBlockColorStep() *GradientBlockColorStep {
	p := b.ctx.Get("GradientBlockColorStep")
	return GradientBlockColorStepFromJSObject(p)
}

// GradientBlockColorStepFromJSObject returns a wrapped GradientBlockColorStep JavaScript class.
func GradientBlockColorStepFromJSObject(p js.Value) *GradientBlockColorStep {
	return &GradientBlockColorStep{p: p}
}

// NewGradientBlockColorStep returns a new GradientBlockColorStep object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblockcolorstep
func (b *Babylon) NewGradientBlockColorStep(todo parameters) *GradientBlockColorStep {
	p := b.ctx.Get("GradientBlockColorStep").New(todo)
	return GradientBlockColorStepFromJSObject(p)
}

// TODO: methods
