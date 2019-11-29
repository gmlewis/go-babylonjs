// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PredicateCondition represents a babylon.js PredicateCondition.
// Defines a predicate condition as an extension of Condition
type PredicateCondition struct{ *Condition }

// JSObject returns the underlying js.Value.
func (p *PredicateCondition) JSObject() js.Value { return p.p }

// PredicateCondition returns a PredicateCondition JavaScript class.
func (b *Babylon) PredicateCondition() *PredicateCondition {
	p := b.ctx.Get("PredicateCondition")
	return PredicateConditionFromJSObject(p)
}

// PredicateConditionFromJSObject returns a wrapped PredicateCondition JavaScript class.
func PredicateConditionFromJSObject(p js.Value) *PredicateCondition {
	return &PredicateCondition{ConditionFromJSObject(p)}
}

// NewPredicateCondition returns a new PredicateCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition
func (b *Babylon) NewPredicateCondition(todo parameters) *PredicateCondition {
	p := b.ctx.Get("PredicateCondition").New(todo)
	return PredicateConditionFromJSObject(p)
}

// TODO: methods
