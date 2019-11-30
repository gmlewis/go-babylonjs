// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoneLookController represents a babylon.js BoneLookController.
// Class used to make a bone look toward a point in space
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons#bonelookcontroller
type BoneLookController struct{ p js.Value }

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

// NewBoneLookControllerOpts contains optional parameters for NewBoneLookController.
type NewBoneLookControllerOpts struct {
	Options js.Value
}

// NewBoneLookController returns a new BoneLookController object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller
func (b *Babylon) NewBoneLookController(mesh *AbstractMesh, bone *Bone, target *Vector3, opts *NewBoneLookControllerOpts) *BoneLookController {
	if opts == nil {
		opts = &NewBoneLookControllerOpts{}
	}

	p := b.ctx.Get("BoneLookController").New(mesh.JSObject(), bone.JSObject(), target.JSObject(), opts.Options)
	return BoneLookControllerFromJSObject(p)
}

// TODO: methods
