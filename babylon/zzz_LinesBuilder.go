// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LinesBuilder represents a babylon.js LinesBuilder.
// Class containing static functions to help procedurally build meshes
type LinesBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LinesBuilder) JSObject() js.Value { return l.p }

// LinesBuilder returns a LinesBuilder JavaScript class.
func (ba *Babylon) LinesBuilder() *LinesBuilder {
	p := ba.ctx.Get("LinesBuilder")
	return LinesBuilderFromJSObject(p, ba.ctx)
}

// LinesBuilderFromJSObject returns a wrapped LinesBuilder JavaScript class.
func LinesBuilderFromJSObject(p js.Value, ctx js.Value) *LinesBuilder {
	return &LinesBuilder{p: p, ctx: ctx}
}

// TODO: methods
