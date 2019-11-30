// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorSplitterBlock represents a babylon.js ColorSplitterBlock.
// Block used to expand a Color3/4 into 4 outputs (one for each component)
type ColorSplitterBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (c *ColorSplitterBlock) JSObject() js.Value { return c.p }

// ColorSplitterBlock returns a ColorSplitterBlock JavaScript class.
func (b *Babylon) ColorSplitterBlock() *ColorSplitterBlock {
	p := b.ctx.Get("ColorSplitterBlock")
	return ColorSplitterBlockFromJSObject(p)
}

// ColorSplitterBlockFromJSObject returns a wrapped ColorSplitterBlock JavaScript class.
func ColorSplitterBlockFromJSObject(p js.Value) *ColorSplitterBlock {
	return &ColorSplitterBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewColorSplitterBlock returns a new ColorSplitterBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorsplitterblock
func (b *Babylon) NewColorSplitterBlock(name string) *ColorSplitterBlock {
	p := b.ctx.Get("ColorSplitterBlock").New(name)
	return ColorSplitterBlockFromJSObject(p)
}

// TODO: methods
