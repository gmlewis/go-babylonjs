// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StopAnimationAction represents a babylon.js StopAnimationAction.
// This defines an action responsible to stop an animation once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type StopAnimationAction struct {
	*Action
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StopAnimationAction) JSObject() js.Value { return s.p }

// StopAnimationAction returns a StopAnimationAction JavaScript class.
func (ba *Babylon) StopAnimationAction() *StopAnimationAction {
	p := ba.ctx.Get("StopAnimationAction")
	return StopAnimationActionFromJSObject(p, ba.ctx)
}

// StopAnimationActionFromJSObject returns a wrapped StopAnimationAction JavaScript class.
func StopAnimationActionFromJSObject(p js.Value, ctx js.Value) *StopAnimationAction {
	return &StopAnimationAction{Action: ActionFromJSObject(p, ctx), ctx: ctx}
}

// StopAnimationActionArrayToJSArray returns a JavaScript Array for the wrapped array.
func StopAnimationActionArrayToJSArray(array []*StopAnimationAction) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStopAnimationActionOpts contains optional parameters for NewStopAnimationAction.
type NewStopAnimationActionOpts struct {
	Condition *Condition
}

// NewStopAnimationAction returns a new StopAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.stopanimationaction#constructor
func (ba *Babylon) NewStopAnimationAction(triggerOptions JSObject, target JSObject, opts *NewStopAnimationActionOpts) *StopAnimationAction {
	if opts == nil {
		opts = &NewStopAnimationActionOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, triggerOptions.JSObject())
	args = append(args, target.JSObject())

	if opts.Condition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Condition.JSObject())
	}

	p := ba.ctx.Get("StopAnimationAction").New(args...)
	return StopAnimationActionFromJSObject(p, ba.ctx)
}

// Execute calls the Execute method on the StopAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.stopanimationaction#execute
func (s *StopAnimationAction) Execute() {

	s.p.Call("execute")
}

// Serialize calls the Serialize method on the StopAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.stopanimationaction#serialize
func (s *StopAnimationAction) Serialize(parent JSObject) js.Value {

	args := make([]interface{}, 0, 1+0)

	if parent == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parent.JSObject())
	}

	retVal := s.p.Call("serialize", args...)
	return retVal
}
