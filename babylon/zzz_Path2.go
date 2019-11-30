// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Path2 represents a babylon.js Path2.
// Represents a 2D path made up of multiple 2D points
type Path2 struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *Path2) JSObject() js.Value { return p.p }

// Path2 returns a Path2 JavaScript class.
func (b *Babylon) Path2() *Path2 {
	p := b.ctx.Get("Path2")
	return Path2FromJSObject(p)
}

// Path2FromJSObject returns a wrapped Path2 JavaScript class.
func Path2FromJSObject(p js.Value) *Path2 {
	return &Path2{p: p}
}

// NewPath2 returns a new Path2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.path2
func (b *Babylon) NewPath2(x float64, y float64) *Path2 {
	p := b.ctx.Get("Path2").New(x, y)
	return Path2FromJSObject(p)
}

// TODO: methods
