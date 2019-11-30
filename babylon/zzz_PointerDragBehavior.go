// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerDragBehavior represents a babylon.js PointerDragBehavior.
// A behavior that when attached to a mesh will allow the mesh to be dragged around the screen based on pointer events
type PointerDragBehavior struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PointerDragBehavior) JSObject() js.Value { return p.p }

// PointerDragBehavior returns a PointerDragBehavior JavaScript class.
func (ba *Babylon) PointerDragBehavior() *PointerDragBehavior {
	p := ba.ctx.Get("PointerDragBehavior")
	return PointerDragBehaviorFromJSObject(p, ba.ctx)
}

// PointerDragBehaviorFromJSObject returns a wrapped PointerDragBehavior JavaScript class.
func PointerDragBehaviorFromJSObject(p js.Value, ctx js.Value) *PointerDragBehavior {
	return &PointerDragBehavior{p: p, ctx: ctx}
}

// NewPointerDragBehaviorOpts contains optional parameters for NewPointerDragBehavior.
type NewPointerDragBehaviorOpts struct {
	Options *JSValue
}

// NewPointerDragBehavior returns a new PointerDragBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerdragbehavior
func (ba *Babylon) NewPointerDragBehavior(opts *NewPointerDragBehaviorOpts) *PointerDragBehavior {
	if opts == nil {
		opts = &NewPointerDragBehaviorOpts{}
	}

	p := ba.ctx.Get("PointerDragBehavior").New(opts.Options.JSObject())
	return PointerDragBehaviorFromJSObject(p, ba.ctx)
}

// TODO: methods
