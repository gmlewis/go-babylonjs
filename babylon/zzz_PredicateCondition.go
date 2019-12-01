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

// PredicateConditionArrayToJSArray returns a JavaScript Array for the wrapped array.
func PredicateConditionArrayToJSArray(array []*PredicateCondition) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPredicateCondition returns a new PredicateCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition
func (ba *Babylon) NewPredicateCondition(actionManager *ActionManager, predicate func()) *PredicateCondition {

	args := make([]interface{}, 0, 2+0)

	args = append(args, actionManager.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { predicate(); return nil }))

	p := ba.ctx.Get("PredicateCondition").New(args...)
	return PredicateConditionFromJSObject(p, ba.ctx)
}

// IsValid calls the IsValid method on the PredicateCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition#isvalid
func (p *PredicateCondition) IsValid() bool {

	retVal := p.p.Call("isValid")
	return retVal.Bool()
}

// Serialize calls the Serialize method on the PredicateCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition#serialize
func (p *PredicateCondition) Serialize() interface{} {

	retVal := p.p.Call("serialize")
	return retVal
}

/*

// Predicate returns the Predicate property of class PredicateCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition#predicate
func (p *PredicateCondition) Predicate(predicate func()) *PredicateCondition {
	p := ba.ctx.Get("PredicateCondition").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {predicate(); return nil}))
	return PredicateConditionFromJSObject(p, ba.ctx)
}

// SetPredicate sets the Predicate property of class PredicateCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.predicatecondition#predicate
func (p *PredicateCondition) SetPredicate(predicate func()) *PredicateCondition {
	p := ba.ctx.Get("PredicateCondition").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {predicate(); return nil}))
	return PredicateConditionFromJSObject(p, ba.ctx)
}

*/
