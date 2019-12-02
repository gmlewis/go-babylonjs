// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IAccessorSparse represents a babylon.js IAccessorSparse.
// Sparse storage of attributes that deviate from their initialization value
type IAccessorSparse struct {
	*IProperty
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IAccessorSparse) JSObject() js.Value { return i.p }

// IAccessorSparse returns a IAccessorSparse JavaScript class.
func (ba *Babylon) IAccessorSparse() *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse")
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// IAccessorSparseFromJSObject returns a wrapped IAccessorSparse JavaScript class.
func IAccessorSparseFromJSObject(p js.Value, ctx js.Value) *IAccessorSparse {
	return &IAccessorSparse{IProperty: IPropertyFromJSObject(p, ctx), ctx: ctx}
}

// IAccessorSparseArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAccessorSparseArrayToJSArray(array []*IAccessorSparse) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Count returns the Count property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#count
func (i *IAccessorSparse) Count(count float64) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(count)
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// SetCount sets the Count property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#count
func (i *IAccessorSparse) SetCount(count float64) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(count)
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// Indices returns the Indices property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#indices
func (i *IAccessorSparse) Indices(indices *IAccessorSparseIndices) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(indices.JSObject())
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// SetIndices sets the Indices property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#indices
func (i *IAccessorSparse) SetIndices(indices *IAccessorSparseIndices) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(indices.JSObject())
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// Values returns the Values property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#values
func (i *IAccessorSparse) Values(values *IAccessorSparseValues) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(values.JSObject())
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

// SetValues sets the Values property of class IAccessorSparse.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparse#values
func (i *IAccessorSparse) SetValues(values *IAccessorSparseValues) *IAccessorSparse {
	p := ba.ctx.Get("IAccessorSparse").New(values.JSObject())
	return IAccessorSparseFromJSObject(p, ba.ctx)
}

*/
