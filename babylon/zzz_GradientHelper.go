// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GradientHelper represents a babylon.js GradientHelper.
// Helper used to simplify some generic gradient tasks
type GradientHelper struct{}

// JSObject returns the underlying js.Value.
func (g *GradientHelper) JSObject() js.Value { return g.p }

// GradientHelper returns a GradientHelper JavaScript class.
func (b *Babylon) GradientHelper() *GradientHelper {
	p := b.ctx.Get("GradientHelper")
	return GradientHelperFromJSObject(p)
}

// GradientHelperFromJSObject returns a wrapped GradientHelper JavaScript class.
func GradientHelperFromJSObject(p js.Value) *GradientHelper {
	return &GradientHelper{p: p}
}

// NewGradientHelper returns a new GradientHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradienthelper
func (b *Babylon) NewGradientHelper(todo parameters) *GradientHelper {
	p := b.ctx.Get("GradientHelper").New(todo)
	return GradientHelperFromJSObject(p)
}

// TODO: methods
