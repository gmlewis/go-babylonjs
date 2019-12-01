// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PlayAnimationAction represents a babylon.js PlayAnimationAction.
// This defines an action responsible to start an animation once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type PlayAnimationAction struct {
	*Action
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PlayAnimationAction) JSObject() js.Value { return p.p }

// PlayAnimationAction returns a PlayAnimationAction JavaScript class.
func (ba *Babylon) PlayAnimationAction() *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction")
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// PlayAnimationActionFromJSObject returns a wrapped PlayAnimationAction JavaScript class.
func PlayAnimationActionFromJSObject(p js.Value, ctx js.Value) *PlayAnimationAction {
	return &PlayAnimationAction{Action: ActionFromJSObject(p, ctx), ctx: ctx}
}

// NewPlayAnimationActionOpts contains optional parameters for NewPlayAnimationAction.
type NewPlayAnimationActionOpts struct {
	Loop      *bool
	Condition *Condition
}

// NewPlayAnimationAction returns a new PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction
func (ba *Babylon) NewPlayAnimationAction(triggerOptions interface{}, target interface{}, from float64, to float64, opts *NewPlayAnimationActionOpts) *PlayAnimationAction {
	if opts == nil {
		opts = &NewPlayAnimationActionOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, triggerOptions)
	args = append(args, target)
	args = append(args, from)
	args = append(args, to)

	if opts.Loop == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Loop)
	}
	if opts.Condition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Condition.JSObject())
	}

	p := ba.ctx.Get("PlayAnimationAction").New(args...)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// Execute calls the Execute method on the PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#execute
func (p *PlayAnimationAction) Execute() {

	args := make([]interface{}, 0, 0+0)

	p.p.Call("execute", args...)
}

// GetTriggerParameter calls the GetTriggerParameter method on the PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#gettriggerparameter
func (p *PlayAnimationAction) GetTriggerParameter() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := p.p.Call("getTriggerParameter", args...)
	return retVal
}

// Serialize calls the Serialize method on the PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#serialize
func (p *PlayAnimationAction) Serialize(parent interface{}) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, parent)

	retVal := p.p.Call("serialize", args...)
	return retVal
}

// SkipToNextActiveAction calls the SkipToNextActiveAction method on the PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#skiptonextactiveaction
func (p *PlayAnimationAction) SkipToNextActiveAction() {

	args := make([]interface{}, 0, 0+0)

	p.p.Call("skipToNextActiveAction", args...)
}

// Then calls the Then method on the PlayAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#then
func (p *PlayAnimationAction) Then(action *Action) *Action {

	args := make([]interface{}, 0, 1+0)

	args = append(args, action.JSObject())

	retVal := p.p.Call("then", args...)
	return ActionFromJSObject(retVal, p.ctx)
}

/*

// From returns the From property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#from
func (p *PlayAnimationAction) From(from float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(from)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetFrom sets the From property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#from
func (p *PlayAnimationAction) SetFrom(from float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(from)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// Loop returns the Loop property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#loop
func (p *PlayAnimationAction) Loop(loop bool) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(loop)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetLoop sets the Loop property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#loop
func (p *PlayAnimationAction) SetLoop(loop bool) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(loop)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// OnBeforeExecuteObservable returns the OnBeforeExecuteObservable property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#onbeforeexecuteobservable
func (p *PlayAnimationAction) OnBeforeExecuteObservable(onBeforeExecuteObservable *Observable) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(onBeforeExecuteObservable.JSObject())
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetOnBeforeExecuteObservable sets the OnBeforeExecuteObservable property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#onbeforeexecuteobservable
func (p *PlayAnimationAction) SetOnBeforeExecuteObservable(onBeforeExecuteObservable *Observable) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(onBeforeExecuteObservable.JSObject())
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// To returns the To property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#to
func (p *PlayAnimationAction) To(to float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(to)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetTo sets the To property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#to
func (p *PlayAnimationAction) SetTo(to float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(to)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// Trigger returns the Trigger property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#trigger
func (p *PlayAnimationAction) Trigger(trigger float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(trigger)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetTrigger sets the Trigger property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#trigger
func (p *PlayAnimationAction) SetTrigger(trigger float64) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(trigger)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// TriggerOptions returns the TriggerOptions property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#triggeroptions
func (p *PlayAnimationAction) TriggerOptions(triggerOptions interface{}) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(triggerOptions)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

// SetTriggerOptions sets the TriggerOptions property of class PlayAnimationAction.
//
// https://doc.babylonjs.com/api/classes/babylon.playanimationaction#triggeroptions
func (p *PlayAnimationAction) SetTriggerOptions(triggerOptions interface{}) *PlayAnimationAction {
	p := ba.ctx.Get("PlayAnimationAction").New(triggerOptions)
	return PlayAnimationActionFromJSObject(p, ba.ctx)
}

*/
