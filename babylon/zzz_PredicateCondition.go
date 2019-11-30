// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PredicateCondition represents a babylon.js PredicateCondition.
// Defines a predicate condition as an extension of Condition
type PredicateCondition struct {
	*Condition
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PredicateCondition) JSObject() js.Value { return p.p }

// PredicateCondition returns a PredicateCondition JavaScript class.
func (ba *Babylon) PredicateCondition() *PredicateCondition {
	p := ba.ctx.Get("PredicateCondition")
	return PredicateConditionFromJSObject(p, ba.ctx)
}

// PredicateConditionFromJSObject returns a wrapped PredicateCondition JavaScript class.
func PredicateConditionFromJSObject(p js.Value, ctx js.Value) *PredicateCondition {
	return &PredicateCondition{Condition: ConditionFromJSObject(p, ctx), ctx: ctx}
}

// NewPredicateCondition returns a new PredicateCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition
func (ba *Babylon) NewPredicateCondition(actionManager *ActionManager, predicate func()) *PredicateCondition {
	p := ba.ctx.Get("PredicateCondition").New(actionManager.JSObject(), predicate)
	return PredicateConditionFromJSObject(p, ba.ctx)
}

// TODO: methods
