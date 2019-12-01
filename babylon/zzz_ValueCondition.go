// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ValueCondition represents a babylon.js ValueCondition.
// Defines specific conditional operators as extensions of Condition
type ValueCondition struct {
	*Condition
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *ValueCondition) JSObject() js.Value { return v.p }

// ValueCondition returns a ValueCondition JavaScript class.
func (ba *Babylon) ValueCondition() *ValueCondition {
	p := ba.ctx.Get("ValueCondition")
	return ValueConditionFromJSObject(p, ba.ctx)
}

// ValueConditionFromJSObject returns a wrapped ValueCondition JavaScript class.
func ValueConditionFromJSObject(p js.Value, ctx js.Value) *ValueCondition {
	return &ValueCondition{Condition: ConditionFromJSObject(p, ctx), ctx: ctx}
}

// NewValueConditionOpts contains optional parameters for NewValueCondition.
type NewValueConditionOpts struct {
	Operator *float64
}

// NewValueCondition returns a new ValueCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition
func (ba *Babylon) NewValueCondition(actionManager *ActionManager, target interface{}, propertyPath string, value interface{}, opts *NewValueConditionOpts) *ValueCondition {
	if opts == nil {
		opts = &NewValueConditionOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, actionManager.JSObject())
	args = append(args, target)
	args = append(args, propertyPath)
	args = append(args, value)

	if opts.Operator == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Operator)
	}

	p := ba.ctx.Get("ValueCondition").New(args...)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// GetOperatorName calls the GetOperatorName method on the ValueCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#getoperatorname
func (v *ValueCondition) GetOperatorName(operator float64) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, operator)

	retVal := v.p.Call("GetOperatorName", args...)
	return retVal.String()
}

// IsValid calls the IsValid method on the ValueCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isvalid
func (v *ValueCondition) IsValid() bool {

	args := make([]interface{}, 0, 0+0)

	retVal := v.p.Call("isValid", args...)
	return retVal.Bool()
}

// Serialize calls the Serialize method on the ValueCondition object.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#serialize
func (v *ValueCondition) Serialize() interface{} {

	args := make([]interface{}, 0, 0+0)

	retVal := v.p.Call("serialize", args...)
	return retVal
}

/*

// IsDifferent returns the IsDifferent property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isdifferent
func (v *ValueCondition) IsDifferent(IsDifferent float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsDifferent)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetIsDifferent sets the IsDifferent property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isdifferent
func (v *ValueCondition) SetIsDifferent(IsDifferent float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsDifferent)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// IsEqual returns the IsEqual property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isequal
func (v *ValueCondition) IsEqual(IsEqual float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsEqual)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetIsEqual sets the IsEqual property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isequal
func (v *ValueCondition) SetIsEqual(IsEqual float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsEqual)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// IsGreater returns the IsGreater property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isgreater
func (v *ValueCondition) IsGreater(IsGreater float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsGreater)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetIsGreater sets the IsGreater property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#isgreater
func (v *ValueCondition) SetIsGreater(IsGreater float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsGreater)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// IsLesser returns the IsLesser property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#islesser
func (v *ValueCondition) IsLesser(IsLesser float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsLesser)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetIsLesser sets the IsLesser property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#islesser
func (v *ValueCondition) SetIsLesser(IsLesser float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(IsLesser)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// Operator returns the Operator property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#operator
func (v *ValueCondition) Operator(operator float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(operator)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetOperator sets the Operator property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#operator
func (v *ValueCondition) SetOperator(operator float64) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(operator)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// PropertyPath returns the PropertyPath property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#propertypath
func (v *ValueCondition) PropertyPath(propertyPath string) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(propertyPath)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetPropertyPath sets the PropertyPath property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#propertypath
func (v *ValueCondition) SetPropertyPath(propertyPath string) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(propertyPath)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#value
func (v *ValueCondition) Value(value interface{}) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(value)
	return ValueConditionFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class ValueCondition.
//
// https://doc.babylonjs.com/api/classes/babylon.valuecondition#value
func (v *ValueCondition) SetValue(value interface{}) *ValueCondition {
	p := ba.ctx.Get("ValueCondition").New(value)
	return ValueConditionFromJSObject(p, ba.ctx)
}

*/
