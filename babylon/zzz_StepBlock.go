// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StepBlock represents a babylon.js StepBlock.
// Block used to step a value
type StepBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StepBlock) JSObject() js.Value { return s.p }

// StepBlock returns a StepBlock JavaScript class.
func (ba *Babylon) StepBlock() *StepBlock {
	p := ba.ctx.Get("StepBlock")
	return StepBlockFromJSObject(p, ba.ctx)
}

// StepBlockFromJSObject returns a wrapped StepBlock JavaScript class.
func StepBlockFromJSObject(p js.Value, ctx js.Value) *StepBlock {
	return &StepBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NewStepBlock returns a new StepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.stepblock
func (ba *Babylon) NewStepBlock(name string) *StepBlock {
	p := ba.ctx.Get("StepBlock").New(name)
	return StepBlockFromJSObject(p, ba.ctx)
}

// TODO: methods
