// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SetParentAction represents a babylon.js SetParentAction.
// This defines an action responsible to set the parent property of the target once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type SetParentAction struct {
	*Action
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SetParentAction) JSObject() js.Value { return s.p }

// SetParentAction returns a SetParentAction JavaScript class.
func (ba *Babylon) SetParentAction() *SetParentAction {
	p := ba.ctx.Get("SetParentAction")
	return SetParentActionFromJSObject(p, ba.ctx)
}

// SetParentActionFromJSObject returns a wrapped SetParentAction JavaScript class.
func SetParentActionFromJSObject(p js.Value, ctx js.Value) *SetParentAction {
	return &SetParentAction{Action: ActionFromJSObject(p, ctx), ctx: ctx}
}

// SetParentActionArrayToJSArray returns a JavaScript Array for the wrapped array.
func SetParentActionArrayToJSArray(array []*SetParentAction) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSetParentActionOpts contains optional parameters for NewSetParentAction.
type NewSetParentActionOpts struct {
	Condition *Condition
}

// NewSetParentAction returns a new SetParentAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setparentaction#constructor
func (ba *Babylon) NewSetParentAction(triggerOptions JSObject, target JSObject, parent JSObject, opts *NewSetParentActionOpts) *SetParentAction {
	if opts == nil {
		opts = &NewSetParentActionOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, triggerOptions.JSObject())
	args = append(args, target.JSObject())
	args = append(args, parent.JSObject())

	if opts.Condition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Condition.JSObject())
	}

	p := ba.ctx.Get("SetParentAction").New(args...)
	return SetParentActionFromJSObject(p, ba.ctx)
}

// Execute calls the Execute method on the SetParentAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setparentaction#execute
func (s *SetParentAction) Execute() {

	s.p.Call("execute")
}

// Serialize calls the Serialize method on the SetParentAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setparentaction#serialize
func (s *SetParentAction) Serialize(parent JSObject) js.Value {

	args := make([]interface{}, 0, 1+0)

	if parent == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parent.JSObject())
	}

	retVal := s.p.Call("serialize", args...)
	return retVal
}
