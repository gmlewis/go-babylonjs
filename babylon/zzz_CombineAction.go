// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CombineAction represents a babylon.js CombineAction.
// This defines an action responsible to trigger several actions once triggered.

//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type CombineAction struct{ *Action }

// JSObject returns the underlying js.Value.
func (c *CombineAction) JSObject() js.Value { return c.p }

// CombineAction returns a CombineAction JavaScript class.
func (b *Babylon) CombineAction() *CombineAction {
	p := b.ctx.Get("CombineAction")
	return CombineActionFromJSObject(p)
}

// CombineActionFromJSObject returns a wrapped CombineAction JavaScript class.
func CombineActionFromJSObject(p js.Value) *CombineAction {
	return &CombineAction{ActionFromJSObject(p)}
}

// NewCombineAction returns a new CombineAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.combineaction
func (b *Babylon) NewCombineAction(todo parameters) *CombineAction {
	p := b.ctx.Get("CombineAction").New(todo)
	return CombineActionFromJSObject(p)
}

// TODO: methods