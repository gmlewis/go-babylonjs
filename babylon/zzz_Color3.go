// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Color3 represents a babylon.js Color3.
// Class used to hold a RBG color
type Color3 struct{}

// JSObject returns the underlying js.Value.
func (c *Color3) JSObject() js.Value { return c.p }

// Color3 returns a Color3 JavaScript class.
func (b *Babylon) Color3() *Color3 {
	p := b.ctx.Get("Color3")
	return Color3FromJSObject(p)
}

// Color3FromJSObject returns a wrapped Color3 JavaScript class.
func Color3FromJSObject(p js.Value) *Color3 {
	return &Color3{p: p}
}

// NewColor3 returns a new Color3 object.
//
// https://doc.babylonjs.com/api/classes/babylon.color3
func (b *Babylon) NewColor3(todo parameters) *Color3 {
	p := b.ctx.Get("Color3").New(todo)
	return Color3FromJSObject(p)
}

// TODO: methods