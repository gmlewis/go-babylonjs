// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ExecuteCodeAction represents a babylon.js ExecuteCodeAction.
// This defines an action responsible to run code (external event) once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type ExecuteCodeAction struct {
	*Action
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *ExecuteCodeAction) JSObject() js.Value { return e.p }

// ExecuteCodeAction returns a ExecuteCodeAction JavaScript class.
func (ba *Babylon) ExecuteCodeAction() *ExecuteCodeAction {
	p := ba.ctx.Get("ExecuteCodeAction")
	return ExecuteCodeActionFromJSObject(p, ba.ctx)
}

// ExecuteCodeActionFromJSObject returns a wrapped ExecuteCodeAction JavaScript class.
func ExecuteCodeActionFromJSObject(p js.Value, ctx js.Value) *ExecuteCodeAction {
	return &ExecuteCodeAction{Action: ActionFromJSObject(p, ctx), ctx: ctx}
}

// ExecuteCodeActionArrayToJSArray returns a JavaScript Array for the wrapped array.
func ExecuteCodeActionArrayToJSArray(array []*ExecuteCodeAction) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewExecuteCodeActionOpts contains optional parameters for NewExecuteCodeAction.
type NewExecuteCodeActionOpts struct {
	Condition *Condition
}

// NewExecuteCodeAction returns a new ExecuteCodeAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.executecodeaction#constructor
func (ba *Babylon) NewExecuteCodeAction(triggerOptions JSObject, jsFunc JSFunc, opts *NewExecuteCodeActionOpts) *ExecuteCodeAction {
	if opts == nil {
		opts = &NewExecuteCodeActionOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, triggerOptions.JSObject())
	args = append(args, js.FuncOf(jsFunc))

	if opts.Condition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Condition.JSObject())
	}

	p := ba.ctx.Get("ExecuteCodeAction").New(args...)
	return ExecuteCodeActionFromJSObject(p, ba.ctx)
}

// Execute calls the Execute method on the ExecuteCodeAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.executecodeaction#execute
func (e *ExecuteCodeAction) Execute(evt *ActionEvent) {

	args := make([]interface{}, 0, 1+0)

	if evt == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, evt.JSObject())
	}

	e.p.Call("execute", args...)
}

// Func returns the Func property of class ExecuteCodeAction.
//
// https://doc.babylonjs.com/api/classes/babylon.executecodeaction#func
func (e *ExecuteCodeAction) Func() js.Value {
	retVal := e.p.Get("func")
	return retVal
}

// SetFunc sets the Func property of class ExecuteCodeAction.
//
// https://doc.babylonjs.com/api/classes/babylon.executecodeaction#func
func (e *ExecuteCodeAction) SetFunc(jsFunc JSFunc) *ExecuteCodeAction {
	e.p.Set("func", js.FuncOf(jsFunc))
	return e
}
