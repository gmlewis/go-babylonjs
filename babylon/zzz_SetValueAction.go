// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SetValueAction represents a babylon.js SetValueAction.
// This defines an action responsible to set a property of the target
// to a desired value once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type SetValueAction struct{ *Action }

// JSObject returns the underlying js.Value.
func (s *SetValueAction) JSObject() js.Value { return s.p }

// SetValueAction returns a SetValueAction JavaScript class.
func (b *Babylon) SetValueAction() *SetValueAction {
	p := b.ctx.Get("SetValueAction")
	return SetValueActionFromJSObject(p)
}

// SetValueActionFromJSObject returns a wrapped SetValueAction JavaScript class.
func SetValueActionFromJSObject(p js.Value) *SetValueAction {
	return &SetValueAction{ActionFromJSObject(p)}
}

// NewSetValueActionOpts contains optional parameters for NewSetValueAction.
type NewSetValueActionOpts struct {
	Condition *Condition
}

// NewSetValueAction returns a new SetValueAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setvalueaction
func (b *Babylon) NewSetValueAction(triggerOptions interface{}, target interface{}, propertyPath string, value interface{}, opts *NewSetValueActionOpts) *SetValueAction {
	if opts == nil {
		opts = &NewSetValueActionOpts{}
	}

	p := b.ctx.Get("SetValueAction").New(triggerOptions, target, propertyPath, value, opts.Condition.JSObject())
	return SetValueActionFromJSObject(p)
}

// TODO: methods
