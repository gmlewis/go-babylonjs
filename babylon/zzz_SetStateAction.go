// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SetStateAction represents a babylon.js SetStateAction.
// This defines an action responsible to set a the state field of the target
// to a desired value once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type SetStateAction struct {
	*Action
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SetStateAction) JSObject() js.Value { return s.p }

// SetStateAction returns a SetStateAction JavaScript class.
func (ba *Babylon) SetStateAction() *SetStateAction {
	p := ba.ctx.Get("SetStateAction")
	return SetStateActionFromJSObject(p, ba.ctx)
}

// SetStateActionFromJSObject returns a wrapped SetStateAction JavaScript class.
func SetStateActionFromJSObject(p js.Value, ctx js.Value) *SetStateAction {
	return &SetStateAction{Action: ActionFromJSObject(p, ctx), ctx: ctx}
}

// SetStateActionArrayToJSArray returns a JavaScript Array for the wrapped array.
func SetStateActionArrayToJSArray(array []*SetStateAction) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSetStateActionOpts contains optional parameters for NewSetStateAction.
type NewSetStateActionOpts struct {
	Condition *Condition
}

// NewSetStateAction returns a new SetStateAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setstateaction
func (ba *Babylon) NewSetStateAction(triggerOptions interface{}, target interface{}, value string, opts *NewSetStateActionOpts) *SetStateAction {
	if opts == nil {
		opts = &NewSetStateActionOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, triggerOptions)
	args = append(args, target)
	args = append(args, value)

	if opts.Condition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Condition.JSObject())
	}

	p := ba.ctx.Get("SetStateAction").New(args...)
	return SetStateActionFromJSObject(p, ba.ctx)
}

// Execute calls the Execute method on the SetStateAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setstateaction#execute
func (s *SetStateAction) Execute() {

	s.p.Call("execute")
}

// Serialize calls the Serialize method on the SetStateAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.setstateaction#serialize
func (s *SetStateAction) Serialize(parent interface{}) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, parent)

	retVal := s.p.Call("serialize", args...)
	return retVal
}

/*

// Value returns the Value property of class SetStateAction.
//
// https://doc.babylonjs.com/api/classes/babylon.setstateaction#value
func (s *SetStateAction) Value(value string) *SetStateAction {
	p := ba.ctx.Get("SetStateAction").New(value)
	return SetStateActionFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class SetStateAction.
//
// https://doc.babylonjs.com/api/classes/babylon.setstateaction#value
func (s *SetStateAction) SetValue(value string) *SetStateAction {
	p := ba.ctx.Get("SetStateAction").New(value)
	return SetStateActionFromJSObject(p, ba.ctx)
}

*/
