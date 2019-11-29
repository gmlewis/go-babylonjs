// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoneLookController represents a babylon.js BoneLookController.
// Class used to make a bone look toward a point in space

//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons#bonelookcontroller
type BoneLookController struct{}

// JSObject returns the underlying js.Value.
func (b *BoneLookController) JSObject() js.Value { return b.p }

// BoneLookController returns a BoneLookController JavaScript class.
func (b *Babylon) BoneLookController() *BoneLookController {
	p := b.ctx.Get("BoneLookController")
	return BoneLookControllerFromJSObject(p)
}

// BoneLookControllerFromJSObject returns a wrapped BoneLookController JavaScript class.
func BoneLookControllerFromJSObject(p js.Value) *BoneLookController {
	return &BoneLookController{p: p}
}

// NewBoneLookController returns a new BoneLookController object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller
func (b *Babylon) NewBoneLookController(todo parameters) *BoneLookController {
	p := b.ctx.Get("BoneLookController").New(todo)
	return BoneLookControllerFromJSObject(p)
}

// TODO: methods
