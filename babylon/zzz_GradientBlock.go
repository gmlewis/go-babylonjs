// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GradientBlock represents a babylon.js GradientBlock.
// Block used to return a color from a gradient based on an input value between 0 and 1
type GradientBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (g *GradientBlock) JSObject() js.Value { return g.p }

// GradientBlock returns a GradientBlock JavaScript class.
func (b *Babylon) GradientBlock() *GradientBlock {
	p := b.ctx.Get("GradientBlock")
	return GradientBlockFromJSObject(p)
}

// GradientBlockFromJSObject returns a wrapped GradientBlock JavaScript class.
func GradientBlockFromJSObject(p js.Value) *GradientBlock {
	return &GradientBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewGradientBlock returns a new GradientBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradientblock
func (b *Babylon) NewGradientBlock(name string) *GradientBlock {
	p := b.ctx.Get("GradientBlock").New(name)
	return GradientBlockFromJSObject(p)
}

// TODO: methods
