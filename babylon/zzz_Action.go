// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Action represents a babylon.js Action.
// The action to be carried out following a trigger
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions#available-actions
type Action struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *Action) JSObject() js.Value { return a.p }

// Action returns a Action JavaScript class.
func (b *Babylon) Action() *Action {
	p := b.ctx.Get("Action")
	return ActionFromJSObject(p)
}

// ActionFromJSObject returns a wrapped Action JavaScript class.
func ActionFromJSObject(p js.Value) *Action {
	return &Action{p: p}
}

// NewActionOpts contains optional parameters for NewAction.
type NewActionOpts struct {
	Condition *Condition
}

// NewAction returns a new Action object.
//
// https://doc.babylonjs.com/api/classes/babylon.action
func (b *Babylon) NewAction(triggerOptions interface{}, opts *NewActionOpts) *Action {
	if opts == nil {
		opts = &NewActionOpts{}
	}

	p := b.ctx.Get("Action").New(triggerOptions, opts.Condition.JSObject())
	return ActionFromJSObject(p)
}

// TODO: methods
