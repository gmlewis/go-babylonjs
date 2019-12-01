// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PrecisionDate represents a babylon.js PrecisionDate.
// Class containing a set of static utilities functions for precision date
type PrecisionDate struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PrecisionDate) JSObject() js.Value { return p.p }

// PrecisionDate returns a PrecisionDate JavaScript class.
func (ba *Babylon) PrecisionDate() *PrecisionDate {
	p := ba.ctx.Get("PrecisionDate")
	return PrecisionDateFromJSObject(p, ba.ctx)
}

// PrecisionDateFromJSObject returns a wrapped PrecisionDate JavaScript class.
func PrecisionDateFromJSObject(p js.Value, ctx js.Value) *PrecisionDate {
	return &PrecisionDate{p: p, ctx: ctx}
}

// PrecisionDateArrayToJSArray returns a JavaScript Array for the wrapped array.
func PrecisionDateArrayToJSArray(array []*PrecisionDate) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Now returns the Now property of class PrecisionDate.
//
// https://doc.babylonjs.com/api/classes/babylon.precisiondate#now
func (p *PrecisionDate) Now(Now float64) *PrecisionDate {
	p := ba.ctx.Get("PrecisionDate").New(Now)
	return PrecisionDateFromJSObject(p, ba.ctx)
}

// SetNow sets the Now property of class PrecisionDate.
//
// https://doc.babylonjs.com/api/classes/babylon.precisiondate#now
func (p *PrecisionDate) SetNow(Now float64) *PrecisionDate {
	p := ba.ctx.Get("PrecisionDate").New(Now)
	return PrecisionDateFromJSObject(p, ba.ctx)
}

*/
